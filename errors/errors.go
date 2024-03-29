package errors

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// start of Placementgroup Messages

const GetPlacementGroupOperationFailed = "failed to perform Get Placement Group Operation for placement group %s with error %w"
const CreatePlacementGroupOperationFailed = "failed to perform Create Placement Group Operation for cloud instance %s with error  %w"
const DeletePlacementGroupOperationFailed = "failed to perform Delete Placement Group Operation for placement group %s with error %w"
const AddMemberPlacementGroupOperationFailed = "failed to perform Add Member Operation for instance %s and placement group %s with error %w"
const DeleteMemberPlacementGroupOperationFailed = "failed to perform Delete Member Operation for instance %s and placement group %s with error %w"

// start of Job Messages
const GetJobOperationFailed = "failed to perform get Job operation for job id %s with error %w"
const GetAllJobsOperationFailed = "failed to perform get all jobs operation with error %w"
const DeleteJobsOperationFailed = "failed to perform delete Job operation for job id %s with error %w"

// start of DHCP Messages
const GetDhcpOperationFailed = "failed to perform Get DHCP Operation for dhcp id %s with error %w"
const CreateDchpOperationFailed = "failed to perform Create DHCP Operation for cloud instance %s with error %w"
const DeleteDhcpOperationFailed = "failed to perform Delete DHCP Operation for dhcp id %s with error %w"

// start of System-Pools Messages
const GetSystemPoolsOperationFailed = "failed to perform Get System Pools Operation for cloud instance %s with error %w"

// start of Image Messages
const GetImageOperationFailed = "failed to perform Get Image Operation for image %s with error %w"
const CreateImageOperationFailed = "failed to perform Create Image Operation for cloud instance %s with error  %w"

// Start of Network Messages
const GetNetworkOperationFailed = "failed to perform Get Network Operation for Network id %s with error %w"
const CreateNetworkOperationFailed = "failed to perform Create Network Operation for Network %s with error %w"
const CreateNetworkPortOperationFailed = "failed to perform Create Network Port Operation for Network %s with error %w"

// start of Volume Messages
const DeleteVolumeOperationFailed = "failed to perform Delete Volume Operation for volume %s with error %w"
const UpdateVolumeOperationFailed = "failed to perform Update Volume Operation for volume %s with error %w"
const GetVolumeOperationFailed = "failed to perform the Get Volume Operation for volume %s with error %w"
const CreateVolumeOperationFailed = "failed to perform the Create volume Operation for volume %s with error %w"
const CreateVolumeV2OperationFailed = "failed to perform the Create volume Operation V2 for volume %s with error %w"
const AttachVolumeOperationFailed = "failed to perform the Attach volume Operation for volume %s with error %w"
const DetachVolumeOperationFailed = "failed to perform the Detach volume Operation for volume %s with error %w"
const GetVolumeRemoteCopyRelationshipsOperationFailed = "failed to Get remote copy relationships of a volume %s for the cloud instance %s with error %w"
const GetVolumeFlashCopyMappingOperationFailed = "failed to Get flash copy mapping of a volume %s for the cloud instance %s with error %w"

// start of Clone Messages
const StartCloneOperationFailed = "failed to start the clone operation for volumes-clone %s with error %w"
const PrepareCloneOperationFailed = "failed to prepare the clone operation for volumes-clone %s with error %w"
const DeleteCloneOperationFailed = "failed to perform delete clone operation %w"
const GetCloneOperationFailed = "failed to get the volumes-clone %s for the cloud instance %s with error %w"
const CreateCloneOperationFailed = "failed to perform the create clone operation %w"

// start of Cloud Instance Messages
const GetCloudInstanceOperationFailed = "failed to Get Cloud Instance %s with error %w"
const UpdateCloudInstanceOperationFailed = "failed to update the Cloud instance %s with error %w"
const DeleteCloudInstanceOperationFailed = "failed to delete the Cloud instance %s with error %w"

// start of PPC Key Messages
const GetPPCKeyOperationFailed = "failed to Get PPC Key %s with error %w"
const CreatePPCKeyOperationFailed = "failed to Create PPC Key with error %w"
const DeletePPCKeyOperationFailed = "failed to Delete PPC Key %s with error %w"

// start of Volume Groups
const GetVolumeGroupOperationFailed = "failed to Get volume-group %s for the cloud instance %s with error %w"
const GetVolumeGroupDetailsOperationFailed = "failed to Get volume-group %s details for the cloud instance %s with error %w"
const CreateVolumeGroupOperationFailed = "failed to perform the Create volume-group Operation for the cloud instance %s with error %w"
const DeleteVolumeGroupOperationFailed = "failed to perform Delete volume-group Operation for volume-group %s with error %w"
const GetLiveVolumeGroupDetailsOperationFailed = "failed to Get live details of volume-group %s for the cloud instance %s with error %w"
const VolumeGroupActionOperationFailed = "failed to perform action on volume-group %s for the cloud instance %s with error %w"
const GetVolumeGroupRemoteCopyRelationshipsOperationFailed = "failed to Get remote copy relationships of the volumes belonging to volume group %s for the cloud instance %s with error %w"

// ErrorTarget ...
type ErrorTarget struct {
	Name string
	Type string
}

// SingleError ...
type SingleError struct {
	Code     string
	Message  string
	MoreInfo string
	Target   ErrorTarget
}

// PowerError ...
type Error struct {
	Payload *models.Error
}

func (e Error) Error() string {
	b, _ := json.Marshal(e.Payload)
	return string(b)
}

// ToError ...
func ToError(err error) error {
	if err == nil {
		return nil
	}

	// check if its ours
	kind := reflect.TypeOf(err).Kind()
	if kind != reflect.Ptr {
		return err
	}

	// next follow pointer
	errstruct := reflect.TypeOf(err).Elem()
	if errstruct.Kind() != reflect.Struct {
		return err
	}

	n := errstruct.NumField()
	found := false
	for i := 0; i < n; i++ {
		if errstruct.Field(i).Name == "Payload" {
			found = true
			break
		}
	}

	if !found {
		return err
	}

	// check if a payload field exists
	payloadValue := reflect.ValueOf(err).Elem().FieldByName("Payload")
	if payloadValue.Interface() == nil {
		return err
	}

	payloadIntf := payloadValue.Elem().Interface()
	payload, parsed := payloadIntf.(models.Error)
	if !parsed {
		return err
	}

	var reterr = Error{
		Payload: &payload,
	}

	return reterr
}

// Retrieve wrapped error from err.
// When does not contain wrapped error returns nil.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
