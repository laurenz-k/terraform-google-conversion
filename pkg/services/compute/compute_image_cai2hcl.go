// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/compute/Image.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/tgc_next/cai2hcl/resource_converter.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package compute

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/converters/utils"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/cai2hcl/models"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/caiasset"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tgcresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/tpgresource"
	"github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/transport"
	transport_tpg "github.com/GoogleCloudPlatform/terraform-google-conversion/v6/pkg/transport"
)

type ComputeImageCai2hclConverter struct {
	name   string
	schema map[string]*schema.Schema
}

func NewComputeImageCai2hclConverter(provider *schema.Provider) models.Cai2hclConverter {
	schema := provider.ResourcesMap[ComputeImageSchemaName].Schema

	return &ComputeImageCai2hclConverter{
		name:   ComputeImageSchemaName,
		schema: schema,
	}
}

// Convert converts asset to HCL resource blocks.
func (c *ComputeImageCai2hclConverter) Convert(asset caiasset.Asset) ([]*models.TerraformResourceBlock, error) {
	var blocks []*models.TerraformResourceBlock
	block, err := c.convertResourceData(asset)
	if err != nil {
		return nil, err
	}
	blocks = append(blocks, block)
	return blocks, nil
}

func (c *ComputeImageCai2hclConverter) convertResourceData(asset caiasset.Asset) (*models.TerraformResourceBlock, error) {
	if asset.Resource == nil || asset.Resource.Data == nil {
		return nil, fmt.Errorf("asset resource data is nil")
	}

	var err error
	res := asset.Resource.Data
	config := transport.NewConfig()
	d := &schema.ResourceData{}

	assetNameParts := strings.Split(asset.Name, "/")
	hclBlockName := assetNameParts[len(assetNameParts)-1]

	hclData := make(map[string]interface{})

	outputFields := map[string]struct{}{"archive_size_bytes": struct{}{}, "creation_timestamp": struct{}{}, "effective_labels": struct{}{}, "label_fingerprint": struct{}{}, "terraform_labels": struct{}{}}
	utils.ParseUrlParamValuesFromAssetName(asset.Name, "//compute.googleapis.com/projects/{{project}}/global/images/{{name}}", outputFields, hclData)

	hclData["description"] = flattenComputeImageDescription(res["description"], d, config)
	hclData["storage_locations"] = flattenComputeImageStorageLocations(res["storageLocations"], d, config)
	hclData["disk_size_gb"] = flattenComputeImageDiskSizeGb(res["diskSizeGb"], d, config)
	hclData["family"] = flattenComputeImageFamily(res["family"], d, config)
	hclData["guest_os_features"] = flattenComputeImageGuestOsFeatures(res["guestOsFeatures"], d, config)
	hclData["image_encryption_key"] = flattenComputeImageImageEncryptionKey(res["imageEncryptionKey"], d, config)
	hclData["labels"] = flattenComputeImageLabels(res["labels"], d, config)
	hclData["licenses"] = flattenComputeImageLicenses(res["licenses"], d, config)
	hclData["name"] = flattenComputeImageName(res["name"], d, config)
	hclData["raw_disk"] = flattenComputeImageRawDisk(res["rawDisk"], d, config)
	hclData["source_disk"] = flattenComputeImageSourceDisk(res["sourceDisk"], d, config)
	hclData["source_disk_encryption_key"] = flattenComputeImageSourceDiskEncryptionKey(res["sourceDiskEncryptionKey"], d, config)
	hclData["source_image"] = flattenComputeImageSourceImage(res["sourceImage"], d, config)
	hclData["source_image_encryption_key"] = flattenComputeImageSourceImageEncryptionKey(res["sourceImageEncryptionKey"], d, config)
	hclData["source_snapshot"] = flattenComputeImageSourceSnapshot(res["sourceSnapshot"], d, config)
	hclData["shielded_instance_initial_state"] = flattenComputeImageShieldedInstanceInitialState(res["shieldedInstanceInitialState"], d, config)
	hclData["source_snapshot_encryption_key"] = flattenComputeImageSourceSnapshotEncryptionKey(res["sourceSnapshotEncryptionKey"], d, config)

	ctyVal, err := utils.MapToCtyValWithSchema(hclData, c.schema)
	if err != nil {
		return nil, err
	}
	return &models.TerraformResourceBlock{
		Labels: []string{c.name, hclBlockName},
		Value:  ctyVal,
	}, nil
}

func flattenComputeImageDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageStorageLocations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageDiskSizeGb(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenComputeImageFamily(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageGuestOsFeatures(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := schema.NewSet(schema.HashResource(computeImageGuestOsFeaturesSchema()), []interface{}{})
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed.Add(map[string]interface{}{
			"type": flattenComputeImageGuestOsFeaturesType(original["type"], d, config),
		})
	}
	return transformed
}

func flattenComputeImageGuestOsFeaturesType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageImageEncryptionKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_self_link"] =
		flattenComputeImageImageEncryptionKeyKmsKeySelfLink(original["kmsKeyName"], d, config)
	transformed["kms_key_service_account"] =
		flattenComputeImageImageEncryptionKeyKmsKeyServiceAccount(original["kmsKeyServiceAccount"], d, config)
	transformed["raw_key"] =
		flattenComputeImageImageEncryptionKeyRawKey(original["rawKey"], d, config)
	transformed["rsa_encrypted_key"] =
		flattenComputeImageImageEncryptionKeyRsaEncryptedKey(original["rsaEncryptedKey"], d, config)
	return []interface{}{transformed}
}

func flattenComputeImageImageEncryptionKeyKmsKeySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	vStr := v.(string)
	return strings.Split(vStr, "/cryptoKeyVersions/")[0]
}

func flattenComputeImageImageEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageImageEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("image_encryption_key.0.raw_key")
}

func flattenComputeImageImageEncryptionKeyRsaEncryptedKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("image_encryption_key.0.rsa_encrypted_key")
}

func flattenComputeImageLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return tgcresource.RemoveTerraformAttributionLabel(v)
}
func flattenComputeImageLicenses(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	return tpgresource.ConvertAndMapStringArr(v.([]interface{}), tpgresource.ConvertSelfLinkToV1)
}

func flattenComputeImageName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageRawDisk(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("raw_disk")
}

func flattenComputeImageRawDiskContainerType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageRawDiskSha1(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageRawDiskSource(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceDisk(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	relative, err := tpgresource.GetRelativePath(v.(string))
	if err != nil {
		return v
	}
	return relative
}

func flattenComputeImageSourceDiskEncryptionKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("source_disk_encryption_key")
}

func flattenComputeImageSourceDiskEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceDiskEncryptionKeyRsaEncryptedKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceDiskEncryptionKeyKmsKeySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	vStr := v.(string)
	return strings.Split(vStr, "/cryptoKeyVersions/")[0]
}

func flattenComputeImageSourceDiskEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceImage(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	relative, err := tpgresource.GetRelativePath(v.(string))
	if err != nil {
		return v
	}
	return relative
}

func flattenComputeImageSourceImageEncryptionKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("source_image_encryption_key")
}

func flattenComputeImageSourceImageEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceImageEncryptionKeyRsaEncryptedKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceImageEncryptionKeyKmsKeySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	vStr := v.(string)
	return strings.Split(vStr, "/cryptoKeyVersions/")[0]
}

func flattenComputeImageSourceImageEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceSnapshot(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	relative, err := tpgresource.GetRelativePath(v.(string))
	if err != nil {
		return v
	}
	return relative
}

func flattenComputeImageShieldedInstanceInitialState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["pk"] =
		flattenComputeImageShieldedInstanceInitialStatePk(original["pk"], d, config)
	transformed["keks"] =
		flattenComputeImageShieldedInstanceInitialStateKeks(original["keks"], d, config)
	transformed["dbs"] =
		flattenComputeImageShieldedInstanceInitialStateDbs(original["dbs"], d, config)
	transformed["dbxs"] =
		flattenComputeImageShieldedInstanceInitialStateDbxs(original["dbxs"], d, config)
	return []interface{}{transformed}
}

func flattenComputeImageShieldedInstanceInitialStatePk(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["content"] =
		flattenComputeImageShieldedInstanceInitialStatePkContent(original["content"], d, config)
	transformed["file_type"] =
		flattenComputeImageShieldedInstanceInitialStatePkFileType(original["fileType"], d, config)
	return []interface{}{transformed}
}

func flattenComputeImageShieldedInstanceInitialStatePkContent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageShieldedInstanceInitialStatePkFileType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageShieldedInstanceInitialStateKeks(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"content":   flattenComputeImageShieldedInstanceInitialStateKeksContent(original["content"], d, config),
			"file_type": flattenComputeImageShieldedInstanceInitialStateKeksFileType(original["fileType"], d, config),
		})
	}
	return transformed
}

func flattenComputeImageShieldedInstanceInitialStateKeksContent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageShieldedInstanceInitialStateKeksFileType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageShieldedInstanceInitialStateDbs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"content":   flattenComputeImageShieldedInstanceInitialStateDbsContent(original["content"], d, config),
			"file_type": flattenComputeImageShieldedInstanceInitialStateDbsFileType(original["fileType"], d, config),
		})
	}
	return transformed
}

func flattenComputeImageShieldedInstanceInitialStateDbsContent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageShieldedInstanceInitialStateDbsFileType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageShieldedInstanceInitialStateDbxs(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"content":   flattenComputeImageShieldedInstanceInitialStateDbxsContent(original["content"], d, config),
			"file_type": flattenComputeImageShieldedInstanceInitialStateDbxsFileType(original["fileType"], d, config),
		})
	}
	return transformed
}

func flattenComputeImageShieldedInstanceInitialStateDbxsContent(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageShieldedInstanceInitialStateDbxsFileType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceSnapshotEncryptionKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return d.Get("source_snapshot_encryption_key")
}

func flattenComputeImageSourceSnapshotEncryptionKeyRawKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceSnapshotEncryptionKeyRsaEncryptedKey(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenComputeImageSourceSnapshotEncryptionKeyKmsKeySelfLink(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	vStr := v.(string)
	return strings.Split(vStr, "/cryptoKeyVersions/")[0]
}

func flattenComputeImageSourceSnapshotEncryptionKeyKmsKeyServiceAccount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}
