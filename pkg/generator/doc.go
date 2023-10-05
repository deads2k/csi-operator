// Package generator contains the code for generating YAML files of a CSI driver operator.
//
// The generator is intended to be used *before* building the operator, using `make update`.
// It generates YAML files based on the configuration in CSIDriverGeneratorConfig.
//
// The YAML files are generated from base assets (in assets/base directory and assets/drivers/*/base) that are patched
// heavily using strategic merge patches or JSON patches. These patches come from assets/patches and
// assets/drivers/*/patches.
//
// The generated assets are typically conditional on the cluster flavour e.g. a YAML file will be included only on
// HyperShift, only on standalone or on both. Similarly, patches are typically applied conditionally too - e.g. a patch
// can add an HyperShift-only specific parameters to a YAML file.
//
// The generated assets are saved in assets/generated/*/* directories and are part of the git repository. I.e.
// any PR that changes the way the assets are generated will include also the generated result, so it's
// easy to spot the differences there.
//
// There is as little go code as possible, however, e.g. adding extra cmdline arguments to CSI sidecars is impossible
// with strategic merge patch and ugly with JSON patch, so there is a small go code here and there that does such
// changes nicer.
//
// ## Terminology
//
// YAMLWithHistory (with no adjective) - a single YAML file that is part of the operator assets (i.e. inside assets/ directory).
//
// AssetName. *AssetName - path to a file with the YAML file, relative to assets/ directory. E.g.
// `base/controller_sa.yaml`.
//
// PatchAssetName, *Patch* - path to a file with a patch to apply. It follows AssetName rules, i.e. it's relative to
// assets/ directory. If it has `.patch` suffix, it is RFC 6902 JSON-patch encoded as YAML. Otherwise, it's used as
// strategic merge patch.
//
// generated YAMLWithHistory - a single YAML file that is generated by the generator. It is saved in `assets/generated/*/*`
// directory, and it is input of the CSI driver operator in runtime.
//
// GeneratedAssetName - name of a generated YAMLWithHistory file. It is relative to its generated YAMLWithHistory directory
// (e.g. assets/generated/aws-ebs/hypershift/)!
// It usually does not have any path component. Name of the generated assets are:
//
//   - controller.yaml - for the controller Deployment.
//   - node.yaml - for the node DaemonSet.
//   - service.yaml - for the metrics Service.
//   - servicemonitor.yaml - for the metrics ServiceMonitor.
//   - base(sourceAssetName) - for all other assets. E.g. `controller_sa.yaml`, generated out of `base/controller_sa.yaml`
//     and any patches that modify it.
//
// The reason for different base path is that we want the generated assets in a single directory
// (e.g. assets/generated/aws-ebs/hypershift/), using their complete paths with `base/` or `drivers/aws-ebs/base` prefix
// would be quite ugly.
//
// ClusterFlavour - type of the cluster installation. E.g. standalone or HyperShift. Standalone covers also single-node
// clusters.
//
// (Please file a PR where the terminology is not applied correctly.)
package generator
