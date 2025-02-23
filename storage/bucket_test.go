// Copyright 2017 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"testing"
	"time"

	"cloud.google.com/go/internal/testutil"
	"github.com/google/go-cmp/cmp"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	raw "google.golang.org/api/storage/v1"
)

// TODO(#6539): re-enable tests after breaking change is released and AgeInDays is a int64*
func TestBucketAttrsToRawBucket(t *testing.T) {
	t.Skip("TestBucketAttrsToRawBucket skipped: https://github.com/googleapis/google-cloud-go/issues/6539")
	t.Parallel()
	attrs := &BucketAttrs{
		Name: "name",
		ACL:  []ACLRule{{Entity: "bob@example.com", Role: RoleOwner, Domain: "d", Email: "e"}},
		DefaultObjectACL: []ACLRule{{Entity: AllUsers, Role: RoleReader, EntityID: "eid",
			ProjectTeam: &ProjectTeam{ProjectNumber: "17", Team: "t"}}},
		Etag:         "Zkyw9ACJZUvcYmlFaKGChzhmtnE/dt1zHSfweiWpwzdGsqXwuJZqiD0",
		Location:     "loc",
		StorageClass: "class",
		RetentionPolicy: &RetentionPolicy{
			RetentionPeriod: 3 * time.Second,
		},
		BucketPolicyOnly:         BucketPolicyOnly{Enabled: true},
		UniformBucketLevelAccess: UniformBucketLevelAccess{Enabled: true},
		PublicAccessPrevention:   PublicAccessPreventionEnforced,
		VersioningEnabled:        false,
		RPO:                      RPOAsyncTurbo,
		// should be ignored:
		MetaGeneration: 39,
		Created:        time.Now(),
		Labels:         map[string]string{"label": "value"},
		CORS: []CORS{
			{
				MaxAge:          time.Hour,
				Methods:         []string{"GET", "POST"},
				Origins:         []string{"*"},
				ResponseHeaders: []string{"FOO"},
			},
		},
		Encryption: &BucketEncryption{DefaultKMSKeyName: "key"},
		Logging:    &BucketLogging{LogBucket: "lb", LogObjectPrefix: "p"},
		Website:    &BucketWebsite{MainPageSuffix: "mps", NotFoundPage: "404"},
		Lifecycle: Lifecycle{
			Rules: []LifecycleRule{{
				Action: LifecycleAction{
					Type:         SetStorageClassAction,
					StorageClass: "NEARLINE",
				},
				Condition: LifecycleCondition{
					AgeInDays:             10,
					Liveness:              Live,
					CreatedBefore:         time.Date(2017, 1, 2, 3, 4, 5, 6, time.UTC),
					MatchesStorageClasses: []string{"STANDARD"},
					NumNewerVersions:      3,
				},
			}, {
				Action: LifecycleAction{
					Type:         SetStorageClassAction,
					StorageClass: "ARCHIVE",
				},
				Condition: LifecycleCondition{
					CustomTimeBefore:      time.Date(2020, 1, 2, 3, 0, 0, 0, time.UTC),
					DaysSinceCustomTime:   100,
					Liveness:              Live,
					MatchesStorageClasses: []string{"STANDARD"},
				},
			}, {
				Action: LifecycleAction{
					Type: DeleteAction,
				},
				Condition: LifecycleCondition{
					DaysSinceNoncurrentTime: 30,
					Liveness:                Live,
					NoncurrentTimeBefore:    time.Date(2017, 1, 2, 3, 4, 5, 6, time.UTC),
					MatchesStorageClasses:   []string{"NEARLINE"},
					NumNewerVersions:        10,
				},
			}, {
				Action: LifecycleAction{
					Type: DeleteAction,
				},
				Condition: LifecycleCondition{
					AgeInDays:        10,
					MatchesPrefix:    []string{"testPrefix"},
					MatchesSuffix:    []string{"testSuffix"},
					NumNewerVersions: 3,
				},
			}, {
				Action: LifecycleAction{
					Type: DeleteAction,
				},
				Condition: LifecycleCondition{
					Liveness: Archived,
				},
			}, {
				Action: LifecycleAction{
					Type: AbortIncompleteMPUAction,
				},
				Condition: LifecycleCondition{
					AgeInDays: 20,
				},
			}},
		},
	}
	got := attrs.toRawBucket()
	want := &raw.Bucket{
		Name: "name",
		Acl: []*raw.BucketAccessControl{
			{Entity: "bob@example.com", Role: "OWNER"}, // other fields ignored on create/update
		},
		DefaultObjectAcl: []*raw.ObjectAccessControl{
			{Entity: "allUsers", Role: "READER"}, // other fields ignored on create/update
		},
		Location:     "loc",
		StorageClass: "class",
		RetentionPolicy: &raw.BucketRetentionPolicy{
			RetentionPeriod: 3,
		},
		IamConfiguration: &raw.BucketIamConfiguration{
			UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled: true,
			},
			PublicAccessPrevention: "enforced",
		},
		Versioning: nil, // ignore VersioningEnabled if false
		Rpo:        rpoAsyncTurbo,
		Labels:     map[string]string{"label": "value"},
		Cors: []*raw.BucketCors{
			{
				MaxAgeSeconds:  3600,
				Method:         []string{"GET", "POST"},
				Origin:         []string{"*"},
				ResponseHeader: []string{"FOO"},
			},
		},
		Encryption: &raw.BucketEncryption{DefaultKmsKeyName: "key"},
		Logging:    &raw.BucketLogging{LogBucket: "lb", LogObjectPrefix: "p"},
		Website:    &raw.BucketWebsite{MainPageSuffix: "mps", NotFoundPage: "404"},
		Lifecycle: &raw.BucketLifecycle{
			Rule: []*raw.BucketLifecycleRule{{
				Action: &raw.BucketLifecycleRuleAction{
					Type:         SetStorageClassAction,
					StorageClass: "NEARLINE",
				},
				Condition: &raw.BucketLifecycleRuleCondition{
					//Age:                 10,
					IsLive:              googleapi.Bool(true),
					CreatedBefore:       "2017-01-02",
					MatchesStorageClass: []string{"STANDARD"},
					NumNewerVersions:    3,
				},
			},
				{
					Action: &raw.BucketLifecycleRuleAction{
						StorageClass: "ARCHIVE",
						Type:         SetStorageClassAction,
					},
					Condition: &raw.BucketLifecycleRuleCondition{
						IsLive:              googleapi.Bool(true),
						CustomTimeBefore:    "2020-01-02",
						DaysSinceCustomTime: 100,
						MatchesStorageClass: []string{"STANDARD"},
					},
				},
				{
					Action: &raw.BucketLifecycleRuleAction{
						Type: DeleteAction,
					},
					Condition: &raw.BucketLifecycleRuleCondition{
						DaysSinceNoncurrentTime: 30,
						IsLive:                  googleapi.Bool(true),
						NoncurrentTimeBefore:    "2017-01-02",
						MatchesStorageClass:     []string{"NEARLINE"},
						NumNewerVersions:        10,
					},
				},
				{
					Action: &raw.BucketLifecycleRuleAction{
						Type: DeleteAction,
					},
					Condition: &raw.BucketLifecycleRuleCondition{
						//Age:              10,
						MatchesPrefix:    []string{"testPrefix"},
						MatchesSuffix:    []string{"testSuffix"},
						NumNewerVersions: 3,
					},
				},
				{
					Action: &raw.BucketLifecycleRuleAction{
						Type: DeleteAction,
					},
					Condition: &raw.BucketLifecycleRuleCondition{
						IsLive: googleapi.Bool(false),
					},
				},
				{
					Action: &raw.BucketLifecycleRuleAction{
						Type: AbortIncompleteMPUAction,
					},
					Condition: &raw.BucketLifecycleRuleCondition{
						//Age: 20,
					},
				},
			},
		},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Error(msg)
	}

	attrs.VersioningEnabled = true
	attrs.RequesterPays = true
	got = attrs.toRawBucket()
	want.Versioning = &raw.BucketVersioning{Enabled: true}
	want.Billing = &raw.BucketBilling{RequesterPays: true}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Error(msg)
	}

	// Test that setting either of BucketPolicyOnly or UniformBucketLevelAccess
	// will enable UniformBucketLevelAccess.
	// Set UBLA.Enabled = true --> UBLA should be set to enabled in the proto.
	attrs.BucketPolicyOnly = BucketPolicyOnly{}
	attrs.UniformBucketLevelAccess = UniformBucketLevelAccess{Enabled: true}
	got = attrs.toRawBucket()
	want.IamConfiguration = &raw.BucketIamConfiguration{
		UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
			Enabled: true,
		},
		PublicAccessPrevention: "enforced",
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Set BucketPolicyOnly.Enabled = true --> UBLA should be set to enabled in
	// the proto.
	attrs.BucketPolicyOnly = BucketPolicyOnly{Enabled: true}
	attrs.UniformBucketLevelAccess = UniformBucketLevelAccess{}
	got = attrs.toRawBucket()
	want.IamConfiguration = &raw.BucketIamConfiguration{
		UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
			Enabled: true,
		},
		PublicAccessPrevention: "enforced",
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Set both BucketPolicyOnly.Enabled = true and
	// UniformBucketLevelAccess.Enabled=true --> UBLA should be set to enabled
	// in the proto.
	attrs.BucketPolicyOnly = BucketPolicyOnly{Enabled: true}
	attrs.UniformBucketLevelAccess = UniformBucketLevelAccess{Enabled: true}
	got = attrs.toRawBucket()
	want.IamConfiguration = &raw.BucketIamConfiguration{
		UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
			Enabled: true,
		},
		PublicAccessPrevention: "enforced",
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Set UBLA.Enabled=false and BucketPolicyOnly.Enabled=false --> UBLA
	// should be disabled in the proto.
	attrs.BucketPolicyOnly = BucketPolicyOnly{}
	attrs.UniformBucketLevelAccess = UniformBucketLevelAccess{}
	got = attrs.toRawBucket()
	want.IamConfiguration = &raw.BucketIamConfiguration{
		PublicAccessPrevention: "enforced",
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Test that setting PublicAccessPrevention to "unspecified" leads to the
	// inherited setting being propagated in the proto.
	attrs.PublicAccessPrevention = PublicAccessPreventionUnspecified
	got = attrs.toRawBucket()
	want.IamConfiguration = &raw.BucketIamConfiguration{
		PublicAccessPrevention: "inherited",
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Test that setting PublicAccessPrevention to "inherited" leads to the
	// setting being propagated in the proto.
	attrs.PublicAccessPrevention = PublicAccessPreventionInherited
	got = attrs.toRawBucket()
	want.IamConfiguration = &raw.BucketIamConfiguration{
		PublicAccessPrevention: "inherited",
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Test that setting RPO to default is propagated in the proto.
	attrs.RPO = RPODefault
	got = attrs.toRawBucket()
	want.Rpo = rpoDefault
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Re-enable UBLA and confirm that it does not affect the PAP setting.
	attrs.UniformBucketLevelAccess = UniformBucketLevelAccess{Enabled: true}
	got = attrs.toRawBucket()
	want.IamConfiguration = &raw.BucketIamConfiguration{
		UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
			Enabled: true,
		},
		PublicAccessPrevention: "inherited",
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Disable UBLA and reset PAP to default. Confirm that the IAM config is set
	// to nil in the proto.
	attrs.UniformBucketLevelAccess = UniformBucketLevelAccess{Enabled: false}
	attrs.PublicAccessPrevention = PublicAccessPreventionUnknown
	got = attrs.toRawBucket()
	want.IamConfiguration = nil
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}
}

// TODO(#6539): re-enable tests after breaking change is released and AgeInDays is a int64*
func TestBucketAttrsToUpdateToRawBucket(t *testing.T) {
	t.Skip("TestBucketAttrsToUpdateToRawBucket skipped: https://github.com/googleapis/google-cloud-go/issues/6539")
	t.Parallel()
	au := &BucketAttrsToUpdate{
		VersioningEnabled:        false,
		RequesterPays:            false,
		BucketPolicyOnly:         &BucketPolicyOnly{Enabled: false},
		UniformBucketLevelAccess: &UniformBucketLevelAccess{Enabled: false},
		DefaultEventBasedHold:    false,
		RetentionPolicy:          &RetentionPolicy{RetentionPeriod: time.Hour},
		Encryption:               &BucketEncryption{DefaultKMSKeyName: "key2"},
		Lifecycle: &Lifecycle{
			Rules: []LifecycleRule{
				{
					Action:    LifecycleAction{Type: "Delete"},
					Condition: LifecycleCondition{AgeInDays: 30},
				},
				{
					Action:    LifecycleAction{Type: AbortIncompleteMPUAction},
					Condition: LifecycleCondition{AgeInDays: 13},
				},
			},
		},
		Logging:      &BucketLogging{LogBucket: "lb", LogObjectPrefix: "p"},
		Website:      &BucketWebsite{MainPageSuffix: "mps", NotFoundPage: "404"},
		StorageClass: "NEARLINE",
	}
	au.SetLabel("a", "foo")
	au.DeleteLabel("b")
	au.SetLabel("c", "")
	got := au.toRawBucket()
	want := &raw.Bucket{
		Versioning: &raw.BucketVersioning{
			Enabled:         false,
			ForceSendFields: []string{"Enabled"},
		},
		Labels: map[string]string{
			"a": "foo",
			"c": "",
		},
		Billing: &raw.BucketBilling{
			RequesterPays:   false,
			ForceSendFields: []string{"RequesterPays"},
		},
		DefaultEventBasedHold: false,
		RetentionPolicy:       &raw.BucketRetentionPolicy{RetentionPeriod: 3600},
		IamConfiguration: &raw.BucketIamConfiguration{
			UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled:         false,
				ForceSendFields: []string{"Enabled"},
			},
		},
		Encryption: &raw.BucketEncryption{DefaultKmsKeyName: "key2"},
		NullFields: []string{"Labels.b"},
		Lifecycle: &raw.BucketLifecycle{
			Rule: []*raw.BucketLifecycleRule{
				{
					Action: &raw.BucketLifecycleRuleAction{Type: "Delete"},
					//Condition: &raw.BucketLifecycleRuleCondition{Age: 30},
				},
				{
					Action: &raw.BucketLifecycleRuleAction{Type: AbortIncompleteMPUAction},
					//Condition: &raw.BucketLifecycleRuleCondition{Age: 13},
				},
			},
		},
		Logging:         &raw.BucketLogging{LogBucket: "lb", LogObjectPrefix: "p"},
		Website:         &raw.BucketWebsite{MainPageSuffix: "mps", NotFoundPage: "404"},
		StorageClass:    "NEARLINE",
		ForceSendFields: []string{"DefaultEventBasedHold", "Lifecycle"},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Error(msg)
	}

	var au2 BucketAttrsToUpdate
	au2.DeleteLabel("b")
	got = au2.toRawBucket()
	want = &raw.Bucket{
		Labels:          map[string]string{},
		ForceSendFields: []string{"Labels"},
		NullFields:      []string{"Labels.b"},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Error(msg)
	}

	// Test nulls.
	au3 := &BucketAttrsToUpdate{
		RetentionPolicy: &RetentionPolicy{},
		Encryption:      &BucketEncryption{},
		Logging:         &BucketLogging{},
		Website:         &BucketWebsite{},
	}
	got = au3.toRawBucket()
	want = &raw.Bucket{
		NullFields: []string{"RetentionPolicy", "Encryption", "Logging", "Website"},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Error(msg)
	}

	// Test that setting either of BucketPolicyOnly or UniformBucketLevelAccess
	// will enable UniformBucketLevelAccess.
	// Set UBLA.Enabled = true --> UBLA should be set to enabled in the proto.
	au4 := &BucketAttrsToUpdate{
		UniformBucketLevelAccess: &UniformBucketLevelAccess{Enabled: true},
	}
	got = au4.toRawBucket()
	want = &raw.Bucket{
		IamConfiguration: &raw.BucketIamConfiguration{
			UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled:         true,
				ForceSendFields: []string{"Enabled"},
			},
		},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Set BucketPolicyOnly.Enabled = true --> UBLA should be set to enabled in
	// the proto.
	au5 := &BucketAttrsToUpdate{
		BucketPolicyOnly: &BucketPolicyOnly{Enabled: true},
	}
	got = au5.toRawBucket()
	want = &raw.Bucket{
		IamConfiguration: &raw.BucketIamConfiguration{
			UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled:         true,
				ForceSendFields: []string{"Enabled"},
			},
		},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Set both BucketPolicyOnly.Enabled = true and
	// UniformBucketLevelAccess.Enabled=true --> UBLA should be set to enabled
	// in the proto.
	au6 := &BucketAttrsToUpdate{
		BucketPolicyOnly:         &BucketPolicyOnly{Enabled: true},
		UniformBucketLevelAccess: &UniformBucketLevelAccess{Enabled: true},
	}
	got = au6.toRawBucket()
	want = &raw.Bucket{
		IamConfiguration: &raw.BucketIamConfiguration{
			UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled:         true,
				ForceSendFields: []string{"Enabled"},
			},
		},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Set UBLA.Enabled=false and BucketPolicyOnly.Enabled=false --> UBLA
	// should be disabled in the proto.
	au7 := &BucketAttrsToUpdate{
		BucketPolicyOnly:         &BucketPolicyOnly{Enabled: false},
		UniformBucketLevelAccess: &UniformBucketLevelAccess{Enabled: false},
	}
	got = au7.toRawBucket()
	want = &raw.Bucket{
		IamConfiguration: &raw.BucketIamConfiguration{
			UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled:         false,
				ForceSendFields: []string{"Enabled"},
			},
		},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// UBLA.Enabled will have precedence above BucketPolicyOnly.Enabled if both
	// are set with different values.
	au8 := &BucketAttrsToUpdate{
		BucketPolicyOnly:         &BucketPolicyOnly{Enabled: true},
		UniformBucketLevelAccess: &UniformBucketLevelAccess{Enabled: false},
	}
	got = au8.toRawBucket()
	want = &raw.Bucket{
		IamConfiguration: &raw.BucketIamConfiguration{
			UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled:         false,
				ForceSendFields: []string{"Enabled"},
			},
		},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}

	// Set an empty Lifecycle and verify that it will be sent.
	au9 := &BucketAttrsToUpdate{
		Lifecycle: &Lifecycle{},
	}
	got = au9.toRawBucket()
	want = &raw.Bucket{
		Lifecycle: &raw.BucketLifecycle{
			ForceSendFields: []string{"Rule"},
		},
		ForceSendFields: []string{"Lifecycle"},
	}
	if msg := testutil.Diff(got, want); msg != "" {
		t.Errorf(msg)
	}
}

func TestAgeConditionBackwardCompat(t *testing.T) {
	var ti int64
	var want int64 = 100
	setAgeCondition(want, &ti)
	if getAgeCondition(ti) != want {
		t.Fatalf("got %v, want %v", getAgeCondition(ti), want)
	}

	var tp *int64
	want = 10
	setAgeCondition(want, &tp)
	if getAgeCondition(tp) != want {
		t.Fatalf("got %v, want %v", getAgeCondition(tp), want)
	}

}

// TODO(#6539): re-enable tests after breaking change is released and AgeInDays is a int64*
func TestNewBucket(t *testing.T) {
	t.Skip("TestNewBucket skipped: https://github.com/googleapis/google-cloud-go/issues/6539")
	labels := map[string]string{"a": "b"}
	matchClasses := []string{"STANDARD"}
	aTime := time.Date(2017, 1, 2, 0, 0, 0, 0, time.UTC)
	rb := &raw.Bucket{
		Name:                  "name",
		Location:              "loc",
		DefaultEventBasedHold: true,
		Metageneration:        3,
		StorageClass:          "sc",
		TimeCreated:           "2017-10-23T04:05:06Z",
		Versioning:            &raw.BucketVersioning{Enabled: true},
		Labels:                labels,
		Billing:               &raw.BucketBilling{RequesterPays: true},
		Etag:                  "Zkyw9ACJZUvcYmlFaKGChzhmtnE/dt1zHSfweiWpwzdGsqXwuJZqiD0",
		Lifecycle: &raw.BucketLifecycle{
			Rule: []*raw.BucketLifecycleRule{{
				Action: &raw.BucketLifecycleRuleAction{
					Type:         "SetStorageClass",
					StorageClass: "NEARLINE",
				},
				Condition: &raw.BucketLifecycleRuleCondition{
					// Age:                 10,
					IsLive:              googleapi.Bool(true),
					CreatedBefore:       "2017-01-02",
					MatchesStorageClass: matchClasses,
					NumNewerVersions:    3,
				},
			}},
		},
		RetentionPolicy: &raw.BucketRetentionPolicy{
			RetentionPeriod: 3,
			EffectiveTime:   aTime.Format(time.RFC3339),
		},
		IamConfiguration: &raw.BucketIamConfiguration{
			BucketPolicyOnly: &raw.BucketIamConfigurationBucketPolicyOnly{
				Enabled:    true,
				LockedTime: aTime.Format(time.RFC3339),
			},
			UniformBucketLevelAccess: &raw.BucketIamConfigurationUniformBucketLevelAccess{
				Enabled:    true,
				LockedTime: aTime.Format(time.RFC3339),
			},
		},
		Cors: []*raw.BucketCors{
			{
				MaxAgeSeconds:  3600,
				Method:         []string{"GET", "POST"},
				Origin:         []string{"*"},
				ResponseHeader: []string{"FOO"},
			},
		},
		Acl: []*raw.BucketAccessControl{
			{Bucket: "name", Role: "READER", Email: "joe@example.com", Entity: "allUsers"},
		},
		LocationType:  "dual-region",
		Encryption:    &raw.BucketEncryption{DefaultKmsKeyName: "key"},
		Logging:       &raw.BucketLogging{LogBucket: "lb", LogObjectPrefix: "p"},
		Website:       &raw.BucketWebsite{MainPageSuffix: "mps", NotFoundPage: "404"},
		ProjectNumber: 123231313,
	}
	want := &BucketAttrs{
		Name:                  "name",
		Location:              "loc",
		DefaultEventBasedHold: true,
		MetaGeneration:        3,
		StorageClass:          "sc",
		Created:               time.Date(2017, 10, 23, 4, 5, 6, 0, time.UTC),
		VersioningEnabled:     true,
		Labels:                labels,
		Etag:                  "Zkyw9ACJZUvcYmlFaKGChzhmtnE/dt1zHSfweiWpwzdGsqXwuJZqiD0",
		RequesterPays:         true,
		Lifecycle: Lifecycle{
			Rules: []LifecycleRule{
				{
					Action: LifecycleAction{
						Type:         SetStorageClassAction,
						StorageClass: "NEARLINE",
					},
					Condition: LifecycleCondition{
						AgeInDays:             10,
						Liveness:              Live,
						CreatedBefore:         time.Date(2017, 1, 2, 0, 0, 0, 0, time.UTC),
						MatchesStorageClasses: matchClasses,
						NumNewerVersions:      3,
					},
				},
			},
		},
		RetentionPolicy: &RetentionPolicy{
			EffectiveTime:   aTime,
			RetentionPeriod: 3 * time.Second,
		},
		BucketPolicyOnly:         BucketPolicyOnly{Enabled: true, LockedTime: aTime},
		UniformBucketLevelAccess: UniformBucketLevelAccess{Enabled: true, LockedTime: aTime},
		CORS: []CORS{
			{
				MaxAge:          time.Hour,
				Methods:         []string{"GET", "POST"},
				Origins:         []string{"*"},
				ResponseHeaders: []string{"FOO"},
			},
		},
		Encryption:       &BucketEncryption{DefaultKMSKeyName: "key"},
		Logging:          &BucketLogging{LogBucket: "lb", LogObjectPrefix: "p"},
		Website:          &BucketWebsite{MainPageSuffix: "mps", NotFoundPage: "404"},
		ACL:              []ACLRule{{Entity: "allUsers", Role: RoleReader, Email: "joe@example.com"}},
		DefaultObjectACL: nil,
		LocationType:     "dual-region",
		ProjectNumber:    123231313,
	}
	got, err := newBucket(rb)
	if err != nil {
		t.Fatal(err)
	}
	if diff := testutil.Diff(got, want); diff != "" {
		t.Errorf("got=-, want=+:\n%s", diff)
	}
}

func TestBucketRetryer(t *testing.T) {
	testCases := []struct {
		name string
		call func(b *BucketHandle) *BucketHandle
		want *retryConfig
	}{
		{
			name: "all defaults",
			call: func(b *BucketHandle) *BucketHandle {
				return b.Retryer()
			},
			want: &retryConfig{},
		},
		{
			name: "set all options",
			call: func(b *BucketHandle) *BucketHandle {
				return b.Retryer(
					WithBackoff(gax.Backoff{
						Initial:    2 * time.Second,
						Max:        30 * time.Second,
						Multiplier: 3,
					}),
					WithPolicy(RetryAlways),
					WithErrorFunc(func(err error) bool { return false }))
			},
			want: &retryConfig{
				backoff: &gax.Backoff{
					Initial:    2 * time.Second,
					Max:        30 * time.Second,
					Multiplier: 3,
				},
				policy:      RetryAlways,
				shouldRetry: func(err error) bool { return false },
			},
		},
		{
			name: "set some backoff options",
			call: func(b *BucketHandle) *BucketHandle {
				return b.Retryer(
					WithBackoff(gax.Backoff{
						Multiplier: 3,
					}))
			},
			want: &retryConfig{
				backoff: &gax.Backoff{
					Multiplier: 3,
				}},
		},
		{
			name: "set policy only",
			call: func(b *BucketHandle) *BucketHandle {
				return b.Retryer(WithPolicy(RetryNever))
			},
			want: &retryConfig{
				policy: RetryNever,
			},
		},
		{
			name: "set ErrorFunc only",
			call: func(b *BucketHandle) *BucketHandle {
				return b.Retryer(
					WithErrorFunc(func(err error) bool { return false }))
			},
			want: &retryConfig{
				shouldRetry: func(err error) bool { return false },
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(s *testing.T) {
			b := tc.call(&BucketHandle{})
			if diff := cmp.Diff(
				b.retry,
				tc.want,
				cmp.AllowUnexported(retryConfig{}, gax.Backoff{}),
				// ErrorFunc cannot be compared directly, but we check if both are
				// either nil or non-nil.
				cmp.Comparer(func(a, b func(err error) bool) bool {
					return (a == nil && b == nil) || (a != nil && b != nil)
				}),
			); diff != "" {
				s.Fatalf("retry not configured correctly: %v", diff)
			}
		})
	}
}
