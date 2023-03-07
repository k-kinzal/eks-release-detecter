
# EKS Release Detector

![CI](https://github.com/k-kinzal/eks-release-detecter/actions/workflows/scheduler.yaml/badge.svg)

This repository monitors [Amazon EKS Kubernetes versions](https://docs.aws.amazon.com/eks/latest/userguide/kubernetes-versions.html) page.
If there are any new version it will update `latest` and `versions` and create [tags](https://github.com/k-kinzal/eks-release-detecter/tags) and [releases](https://github.com/k-kinzal/eks-release-detecter/releases).

## Use cases

Renovate's [Github Releases Datasource](https://docs.renovatebot.com/modules/datasource/github-releases/) and [Github Tags Datasource](https://docs.renovatebot.com/modules/datasource/github-tags/) are used to automatically update the EKS.

**example**

TBD