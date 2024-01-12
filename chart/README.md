# lifeboat

![Version: 0.1.0](https://img.shields.io/badge/Version-0.1.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.1.0](https://img.shields.io/badge/AppVersion-0.1.0-informational?style=flat-square)

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| affinity | object | `{}` | [Kubernetes affinity and anti-affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#affinity-and-anti-affinity) allows defining rules that determine on which nodes the pod should be run preferentially. |
| annotations | object | `{}` | [Additional cronjob annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/) |
| concurrencyPolicy | string | `"Forbid"` | Specifies how to treat concurrent executions of a Job. Valid values are: - "Allow": allows CronJobs to run concurrently; - "Forbid": forbids concurrent runs, skipping next run if previous run hasn't finished yet; - "Replace": cancels currently running job and replaces it with a new one |
| configuration | object | `{}` | Lifeboat configuration |
| env | object | `{}` | Extra environment variables that will be pass onto deployment pods. |
| envConfigMapNames | list | `[]` | [Kubernetes ConfigMap Resource](https://kubernetes.io/docs/concepts/configuration/configmap/) names to load environment variables from. |
| envSecretNames | list | `[]` | [Kubernetes Secret Resource](https://kubernetes.io/docs/concepts/configuration/secret/) names to load environment variables from. |
| envValueFrom | object | `{}` | ["valueFrom" environment variable references](https://kubernetes.io/docs/tasks/inject-data-application/define-environment-variable-container/) that will be added to deployment pods. Name is templated. |
| failedJobsHistoryLimit | int | `1` | The number of failed finished jobs to retain. Value must be non-negative integer. |
| fullnameOverride | string | `""` |  |
| image.pullPolicy | string | `"IfNotPresent"` | The pull policy for the container image. |
| image.repository | string | `"cluetec/lifeboat"` | The repository path of the container image. |
| imagePullSecrets | list | `[]` | Container registry secret names as an array |
| jobAnnotations | object | `{}` | [Additional job annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/) |
| nameOverride | string | `""` |  |
| nodeSelector | object | `{}` | [Kubernetes node selector](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/) allows to select specific Kubernetes nodes (nodes) on which the pod should be scheduled. |
| podAnnotations | object | `{}` | [Additional pod annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/) |
| podSecurityContext | object | `{"fsGroup":2000,"runAsGroup":3000,"runAsUser":1000}` | [Pod security context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod) |
| resources | object | `{}` |  |
| restartPolicy | string | `"Never"` |  |
| schedule | string | `"0 3 * * *"` | The schedule in Cron format, see <https://en.wikipedia.org/wiki/Cron>. Default is everyday at 3am |
| securityContext | object | `{"capabilities":{"drop":["ALL"]},"readOnlyRootFilesystem":true,"runAsNonRoot":true,"runAsUser":1000}` | [Container security context](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-container) |
| serviceAccount.annotations | object | `{}` | Annotations to add to the service account |
| serviceAccount.enabled | bool | `false` | Specifies whether a service account should be created |
| serviceAccount.imagePullSecrets | list | `[]` | List of image pull secret names which should be used by the service account |
| serviceAccount.name | string | `""` | The name of the service account to use. If not set and create is true, a name is generated using the fullname template |
| startingDeadlineSeconds | int | `0` | Optional deadline in seconds for starting the job if it misses scheduled time for any reason. Missed jobs executions will be counted as failed ones. |
| storage.accessModes | list | `["ReadWriteOnce"]` | PV Access Mode |
| storage.annotation | list | `[]` | [Additional pvc annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/) |
| storage.enabled | bool | `false` | Enables/Disables the provisioning of an persistent volume claim |
| storage.existingClaim | string | `""` | Provide an existing `PersistentVolumeClaim`. If defined, PVC must be created manually before volume will be bound |
| storage.mountPath | string | `"/backups"` |  |
| storage.size | string | `"10Gi"` | PVC Storage Request for the backup volume |
| storage.storageClass | string | `nil` | PVC Storage Class for the backup volume If defined, storageClassName: <storageClass> If undefined (the default) or set to null, no storageClassName spec is set, choosing the default provisioner. |
| successfulJobsHistoryLimit | int | `3` | The number of successful finished jobs to retain. Value must be non-negative integer. |
| suspend | bool | `false` | This flag tells the controller to suspend subsequent executions, it does not apply to already started executions. |
| timeZone | string | `""` | The time zone name for the given schedule, see <https://en.wikipedia.org/wiki/List_of_tz_database_time_zones>. If not specified, this will default to the time zone of the kube-controller-manager process. The set of valid time zone names and the time zone offset is loaded from the system-wide time zone database by the API server during CronJob validation and the controller manager during execution. If no system-wide time zone database can be found a bundled version of the database is used instead. If the time zone name becomes invalid during the lifetime of a CronJob or due to a change in host configuration, the controller will stop creating new new Jobs and will create a system event with the reason UnknownTimeZone. More information can be found in <https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/#time-zones> |
| tolerations | list | `[]` | [Kubernetes tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/) allow the scheduler to schedule pods with matching taints (constraints). |
| volumeMounts | list | `[]` | Additional volumeMounts to the backend container |
| volumes | list | `[]` | Additional volumes to the backend pod |

## License

The project is licensed under the "Apache-2.0" license. Details can be found in the `LICENSE` file within the helm chart
or in the [GitHub repository](https://github.com/cluetec/lifeboat).

----------------------------------------------
Autogenerated from chart metadata using [helm-docs](https://github.com/norwoodj/helm-docs/)
