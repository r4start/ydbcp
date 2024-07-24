package queries

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	table_types "github.com/ydb-platform/ydb-go-sdk/v3/table/types"
	"testing"
	"time"
	"ydbcp/internal/types"
)

func TestQueryBuilder_UpdateUpdate(t *testing.T) {
	const (
		queryString = `DECLARE $id_0 AS Uuid;
DECLARE $status_0 AS String;
UPDATE Backups SET status = $status_0 WHERE id = $id_0;
DECLARE $id_1 AS Uuid;
DECLARE $status_1 AS String;
DECLARE $message_1 AS String;
UPDATE Operations SET status = $status_1, message = $message_1 WHERE id = $id_1`
	)
	opId := types.GenerateObjectID()
	backupId := types.GenerateObjectID()
	op := types.TakeBackupOperation{
		Id:      opId,
		State:   "Done",
		Message: "Abcde",
	}
	backup := types.Backup{
		ID:     backupId,
		Status: "Available",
	}
	builder := NewWriteTableQuery().
		WithUpdateBackup(backup).
		WithUpdateOperation(&op)
	var (
		queryParams = table.NewQueryParameters(
			table.ValueParam("$id_0", table_types.UUIDValue(backupId)),
			table.ValueParam("$status_0", table_types.StringValueFromString("Available")),
			table.ValueParam("$id_1", table_types.UUIDValue(opId)),
			table.ValueParam("$status_1", table_types.StringValueFromString("Done")),
			table.ValueParam("$message_1", table_types.StringValueFromString("Abcde")),
		)
	)
	query, err := builder.FormatQuery(context.Background())
	assert.Empty(t, err)
	assert.Equal(
		t, queryString, query.QueryText,
		"bad query format",
	)
	assert.Equal(t, queryParams, query.QueryParams, "bad query params")
}

func TestQueryBuilder_CreateCreate(t *testing.T) {
	const (
		queryString = `DECLARE $id_0 AS Uuid;
DECLARE $container_id_0 AS String;
DECLARE $database_0 AS String;
DECLARE $initiated_0 AS String;
DECLARE $s3_endpoint_0 AS String;
DECLARE $s3_region_0 AS String;
DECLARE $s3_bucket_0 AS String;
DECLARE $s3_path_prefix_0 AS String;
DECLARE $status_0 AS String;
UPSERT INTO Backups (id, container_id, database, initiated, s3_endpoint, s3_region, s3_bucket, s3_path_prefix, status) VALUES ($id_0, $container_id_0, $database_0, $initiated_0, $s3_endpoint_0, $s3_region_0, $s3_bucket_0, $s3_path_prefix_0, $status_0);
DECLARE $id_1 AS Uuid;
DECLARE $type_1 AS String;
DECLARE $status_1 AS String;
DECLARE $container_id_1 AS String;
DECLARE $database_1 AS String;
DECLARE $backup_id_1 AS Uuid;
DECLARE $initiated_1 AS String;
DECLARE $created_at_1 AS Timestamp;
DECLARE $operation_id_1 AS String;
UPSERT INTO Operations (id, type, status, container_id, database, backup_id, initiated, created_at, operation_id) VALUES ($id_1, $type_1, $status_1, $container_id_1, $database_1, $backup_id_1, $initiated_1, $created_at_1, $operation_id_1)`
	)
	opId := types.GenerateObjectID()
	backupId := types.GenerateObjectID()
	tbOp := types.TakeBackupOperation{
		Id:          opId,
		ContainerID: "a",
		BackupId:    backupId,
		State:       "PENDING",
		Message:     "Message",
		YdbConnectionParams: types.YdbConnectionParams{
			Endpoint:     "",
			DatabaseName: "dbname",
		},
		YdbOperationId:      "1234",
		SourcePaths:         nil,
		SourcePathToExclude: nil,
		CreatedAt:           time.Unix(0, 0),
	}
	backup := types.Backup{
		ID:           backupId,
		ContainerID:  "a",
		DatabaseName: "b",
		S3Endpoint:   "c",
		S3Region:     "d",
		S3Bucket:     "e",
		S3PathPrefix: "f",
		Status:       "Available",
	}
	builder := NewWriteTableQuery().
		WithCreateBackup(backup).
		WithCreateOperation(&tbOp)
	var (
		queryParams = table.NewQueryParameters(
			table.ValueParam("$id_0", table_types.UUIDValue(backupId)),
			table.ValueParam("$container_id_0", table_types.StringValueFromString("a")),
			table.ValueParam("$database_0", table_types.StringValueFromString("b")),
			table.ValueParam("$initiated_0", table_types.StringValueFromString("")),
			table.ValueParam("$s3_endpoint_0", table_types.StringValueFromString("c")),
			table.ValueParam("$s3_region_0", table_types.StringValueFromString("d")),
			table.ValueParam("$s3_bucket_0", table_types.StringValueFromString("e")),
			table.ValueParam("$s3_path_prefix_0", table_types.StringValueFromString("f")),
			table.ValueParam("$status_0", table_types.StringValueFromString("Available")),
			table.ValueParam("$id_1", table_types.UUIDValue(opId)),
			table.ValueParam("$type_1", table_types.StringValueFromString("TB")),
			table.ValueParam(
				"$status_1", table_types.StringValueFromString(string(tbOp.State)),
			),
			table.ValueParam(
				"$container_id_1", table_types.StringValueFromString(tbOp.ContainerID),
			),
			table.ValueParam(
				"$database_1",
				table_types.StringValueFromString(tbOp.YdbConnectionParams.DatabaseName),
			),
			table.ValueParam(
				"$backup_id_1",
				table_types.UUIDValue(tbOp.BackupId),
			),
			table.ValueParam(
				"$initiated_1",
				table_types.StringValueFromString(""),
			),
			table.ValueParam(
				"$created_at_1",
				table_types.TimestampValueFromTime(tbOp.CreatedAt),
			),
			table.ValueParam(
				"$operation_id_1",
				table_types.StringValueFromString(tbOp.YdbOperationId),
			),
		)
	)
	query, err := builder.FormatQuery(context.Background())
	assert.Empty(t, err)
	assert.Equal(
		t, queryString, query.QueryText,
		"bad query format",
	)
	assert.Equal(t, queryParams, query.QueryParams, "bad query params")
}

func TestQueryBuilder_UpdateCreate(t *testing.T) {
	const (
		queryString = `DECLARE $id_0 AS Uuid;
DECLARE $status_0 AS String;
UPDATE Backups SET status = $status_0 WHERE id = $id_0;
DECLARE $id_1 AS Uuid;
DECLARE $type_1 AS String;
DECLARE $status_1 AS String;
DECLARE $container_id_1 AS String;
DECLARE $database_1 AS String;
DECLARE $backup_id_1 AS Uuid;
DECLARE $initiated_1 AS String;
DECLARE $created_at_1 AS Timestamp;
DECLARE $operation_id_1 AS String;
UPSERT INTO Operations (id, type, status, container_id, database, backup_id, initiated, created_at, operation_id) VALUES ($id_1, $type_1, $status_1, $container_id_1, $database_1, $backup_id_1, $initiated_1, $created_at_1, $operation_id_1)`
	)
	opId := types.GenerateObjectID()
	backupId := types.GenerateObjectID()
	tbOp := types.TakeBackupOperation{
		Id:          opId,
		ContainerID: "a",
		BackupId:    backupId,
		State:       "PENDING",
		Message:     "Message",
		YdbConnectionParams: types.YdbConnectionParams{
			Endpoint:     "",
			DatabaseName: "dbname",
		},
		YdbOperationId:      "1234",
		SourcePaths:         nil,
		SourcePathToExclude: nil,
		CreatedAt:           time.Unix(0, 0),
	}
	backup := types.Backup{
		ID:     backupId,
		Status: "Available",
	}
	builder := NewWriteTableQuery().
		WithUpdateBackup(backup).
		WithCreateOperation(&tbOp)
	var (
		queryParams = table.NewQueryParameters(
			table.ValueParam("$id_0", table_types.UUIDValue(backupId)),
			table.ValueParam("$status_0", table_types.StringValueFromString("Available")),
			table.ValueParam("$id_1", table_types.UUIDValue(opId)),
			table.ValueParam("$type_1", table_types.StringValueFromString("TB")),
			table.ValueParam(
				"$status_1", table_types.StringValueFromString(string(tbOp.State)),
			),
			table.ValueParam(
				"$container_id_1", table_types.StringValueFromString(tbOp.ContainerID),
			),
			table.ValueParam(
				"$database_1",
				table_types.StringValueFromString(tbOp.YdbConnectionParams.DatabaseName),
			),
			table.ValueParam(
				"$backup_id_1",
				table_types.UUIDValue(tbOp.BackupId),
			),
			table.ValueParam(
				"$initiated_1",
				table_types.StringValueFromString(""),
			),
			table.ValueParam(
				"$created_at_1",
				table_types.TimestampValueFromTime(tbOp.CreatedAt),
			),
			table.ValueParam(
				"$operation_id_1",
				table_types.StringValueFromString(tbOp.YdbOperationId),
			),
		)
	)
	query, err := builder.FormatQuery(context.Background())
	assert.Empty(t, err)
	assert.Equal(
		t, queryString, query.QueryText,
		"bad query format",
	)
	assert.Equal(t, queryParams, query.QueryParams, "bad query params")
}
