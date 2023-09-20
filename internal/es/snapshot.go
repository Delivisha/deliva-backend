package es

import "fmt"

type (
	Snapshot interface {
		SnapshotName() string
	}
	SnapshotApplier interface {
		ApplySnapshot(snaphot Snapshot) error
	}
	Snapshotter interface {
		SnapshotApplier
		ToSnapshot() Snapshot
	}
)

func LoadSnapshot(v interface{}, snapshot Snapshot, version int) error {
	type loader interface {
		SnapshotApplier
		VersionSetter
	}
	agg, ok := v.(loader)
	if !ok {
		return fmt.Errorf("%T does not have methods implemented to load snapshots")
	}
	if err := agg.ApplySnapshot(snapshot); err != nil {
		return err
	}
	agg.setVersion(version)
	return nil
}
