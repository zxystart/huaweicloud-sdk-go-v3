# 0.1.83 2024-02-22

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSubCustomers**
    - 请求参数变更
      - `+ customer_id`

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateDashboardWidgets**
    - 请求参数变更
      - `* body: list<BaseWidgetInfo> -> list<object>`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **GeneratePocketFile**
    - 请求参数变更
      - `* padding: int32 -> float`
  - **ShowAdmetJob**
    - 响应参数变更
      - `- models.value_range.lower_inclusive`
      - `- models.value_range.upper_inclusive`
      - `* models.value_range.lower: number -> float`
      - `* models.value_range.upper: number -> float`
      - `* models.value_range: object<ValueRange> -> object<ValueRange2>`
  - **CreatePocketMolDesignJob**
    - 请求参数变更
      - `* receptor.bounding_box.padding: int32 -> float`
  - **ShowPocketMolDesignJob**
    - 响应参数变更
      - `* receptor.bounding_box.padding: int32 -> float`
      - `- model_list.value_range.lower_inclusive`
      - `- model_list.value_range.upper_inclusive`
      - `* model_list.value_range.lower: number -> float`
      - `* model_list.value_range.upper: number -> float`
      - `* model_list.value_range: object<ValueRange> -> object<ValueRange2>`
  - **ParseDrugReceptorInfo**
    - 请求参数变更
      - `+ add_hydrogen`
      - `* body: object<ReceptorDrugFile> -> object<ReceptorDrugFileReq>`
  - **CreateOptmJob**
    - 请求参数变更
      - `* binding_site.bounding_box.padding: int32 -> float`
  - **ShowOptmJob**
    - 响应参数变更
      - `* binding_site.bounding_box.padding: int32 -> float`
      - `- models.value_range.lower_inclusive`
      - `- models.value_range.upper_inclusive`
      - `* models.value_range.lower: number -> float`
      - `* models.value_range.upper: number -> float`
      - `* models.value_range: object<ValueRange> -> object<ValueRange2>`
  - **CreateDockingJob**
    - 请求参数变更
      - `* receptors.bounding_box.padding: int32 -> float`
  - **ShowDockingJob**
    - 响应参数变更
      - `* receptors.bounding_box.padding: int32 -> float`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ModifyGaussMySqlProxyRouteMode**
    - 请求参数变更
      - `+ new_node_auto_add_status`
      - `+ new_node_weight`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListRecycleInstances**
    - 响应参数变更
      - `+ instances.engine_name`
      - `+ instances.volume_size`
      - `+ instances.enterprise_project_name`
      - `+ instances.backup_level`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListImages**
    - 响应参数变更
      - `+ images.__image_displayname`
  - **UpdateImage**
    - 响应参数变更
      - `+ __image_displayname`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeTrainTicket**
    - 响应参数变更
      - `+ result.invoice_style`
      - `+ result.issue_date`
      - `+ result.discount_mark`
      - `+ result.serial_number`
      - `+ result.tax_amount`
      - `+ result.tax_rate`
      - `+ result.air_conditioning`
      - `+ result.original_invoice_number`
      - `+ result.unified_social_credit_code`
      - `+ result.buyer_name`
      - `+ result.total_amount_excluding_tax`
      - `+ result.invoice_number`
      - `+ result.seal_mark`
      - `+ result.title`
      - `+ result.area`
      - `+ result.receipt_number`
      - `+ result.amount_in_figures`
      - `+ result.amount_in_words`

# 0.1.82 2024-02-20

### HuaweiCloud SDK IAMAccessAnalyzer

- _新增特性_
  - 支持访问分析服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ListClients`、`ScanClients`、`HangUpClients`、`HangUpKillAllClients`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DLI

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSqlJobs**
    - 响应参数变更
      - `* jobs.duration: int32 -> int64`
      - `* jobs.result_count: int32 -> int64`

# 0.1.81 2024-02-07

### HuaweiCloud SDK GEIP

- _新增特性_
  - 支持全域弹性公网IP服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持接口`CreateClusterJob`、`GeneratePocketFile`、`GenerateSurfacePoints`、`GenerateComplexCombine`
- _解决问题_
  - 无
- _特性变更_
  - **CreateProject**
    - 请求参数变更
      - `+ is_new_bucket`
      - `+ bucket_name`
  - **ListComputingResourceFlavors**
    - 响应参数变更
      - `+ flavors.gpu_info`
  - **ListDrugJob**
    - 响应参数变更
      - `* jobs.expect_charge_num: int32 -> double`
      - `* jobs.real_charge_num: int32 -> double`
  - **ShowSynthesisJob**
    - 响应参数变更
      - `* basic_info.expect_charge_num: int32 -> double`
      - `* basic_info.real_charge_num: int32 -> double`
  - **CreateFepJob**
    - 请求参数变更
      - `+ params.pre_equilibrium_time`
      - `+ params.equilibrium_time`
  - **ShowFepJob**
    - 响应参数变更
      - `+ params.pre_equilibrium_time`
      - `+ params.equilibrium_time`
      - `* basic_info.expect_charge_num: int32 -> double`
      - `* basic_info.real_charge_num: int32 -> double`
  - **ShowPocketDetectionJob**
    - 响应参数变更
      - `* basic_info.expect_charge_num: int32 -> double`
      - `* basic_info.real_charge_num: int32 -> double`
  - **ShowAdmetJob**
    - 响应参数变更
      - `+ cluster_result`
      - `* basic_info.expect_charge_num: int32 -> double`
      - `* basic_info.real_charge_num: int32 -> double`
      - `+ models.value_range.lower_inclusive`
      - `+ models.value_range.upper_inclusive`
      - `* models.value_range.lower: float -> number`
      - `* models.value_range.upper: float -> number`
      - `* models.value_range: object<ValueRange2> -> object<ValueRange>`
  - **CreatePocketMolDesignJob**
    - 请求参数变更
      - `+ num_trials`
      - `+ optimization_mode: enum value [scaffold_hopping]`
      - `+ receptor.remove_ligand`
      - `+ receptor.add_hydrogen`
      - `+ ligands.edited`
      - `+ ligands.label_sites`
      - `* ligands: list<DrugFile> -> list<PocketFragment>`
  - **ShowPocketMolDesignJob**
    - 响应参数变更
      - `+ cluster_result`
      - `+ num_trials`
      - `+ optimization_mode: enum value [scaffold_hopping]`
      - `* basic_info.expect_charge_num: int32 -> double`
      - `* basic_info.real_charge_num: int32 -> double`
      - `+ receptor.remove_ligand`
      - `+ receptor.add_hydrogen`
      - `+ ligands.edited`
      - `+ ligands.label_sites`
      - `* ligands: list<DrugFile> -> list<PocketFragment>`
      - `+ model_list.value_range.lower_inclusive`
      - `+ model_list.value_range.upper_inclusive`
      - `* model_list.value_range.lower: float -> number`
      - `* model_list.value_range.upper: float -> number`
      - `* model_list.value_range: object<ValueRange2> -> object<ValueRange>`
  - **RunDrugReceptorPreprocess**
    - 请求参数变更
      - `+ add_hydrogen`
      - `+ file.add_hydrogen`
      - `* body: object<ReceptorDrugFile> -> object<ReceptorDrugFileReq>`
  - **RecognizeDrugReceptorPocket**
    - 请求参数变更
      - `+ receptor_file.add_hydrogen`
      - `* body: object<ReceptorDrugFile> -> object<ReceptorDrugFileReq>`
  - **CreateDrugLigandInteraction2dSvg**
    - 请求参数变更
      - `+ receptor_file.add_hydrogen`
      - `* body: object<ReceptorDrugFile> -> object<ReceptorDrugFileReq>`
  - **CreateOptmJob**
    - 请求参数变更
      - `+ binding_site.add_hydrogen`
  - **ShowOptmJob**
    - 响应参数变更
      - `+ cluster_result`
      - `* basic_info.expect_charge_num: int32 -> double`
      - `* basic_info.real_charge_num: int32 -> double`
      - `+ binding_site.add_hydrogen`
      - `+ models.value_range.lower_inclusive`
      - `+ models.value_range.upper_inclusive`
      - `* models.value_range.lower: float -> number`
      - `* models.value_range.upper: float -> number`
      - `* models.value_range: object<ValueRange2> -> object<ValueRange>`
  - **CreateDockingJob**
    - 请求参数变更
      - `+ receptors.add_hydrogen`
  - **ShowDockingJob**
    - 响应参数变更
      - `+ cluster_result`
      - `* basic_info.expect_charge_num: int32 -> double`
      - `* basic_info.real_charge_num: int32 -> double`
      - `+ receptors.add_hydrogen`
  - **ListDrugModel**
    - 响应参数变更
      - `+ models.losses`
      - `+ models.metrics`
      - `+ models.ModelMetric`

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持以下接口：
    - `ListTenantVpcIgws`
    - `CreateTenantVpcIgw`
    - `ShowInternalVpcIgw`
    - `UpdateTenantVpcIgw`
    - `DeleteTenantVpcIgw`
    - `ListProjectGeipBindings`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`UpdateNewNodeAutoAddSwitch`
- _解决问题_
  - 无
- _特性变更_
  - **CreateGaussMySqlProxy**
    - 请求参数变更
      - `+ new_node_auto_add_status`
      - `+ new_node_weight`
  - **ShowGaussMySqlProxyList**
    - 响应参数变更
      - `+ proxy_list.proxy.new_node_auto_add_status`
      - `+ proxy_list.proxy.new_node_weight`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持接口`StopBackup`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Meeting

- _新增特性_
  - 支持接口`SetProfileImage`、`SetUserProfileImage`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持以下接口：
    - `UpdateDatabaseOwner`
    - `ExecutePrivilegeDatabaseUserRole`
    - `ExecuteRevokeDatabaseUserRole`
    - `ListDatabaseUserRole`
    - `UpdatePostgresqlExtension`
- _解决问题_
  - 无
- _特性变更_
  - **ListPostgresqlExtension**
    - 响应参数变更
      - `+ extensions.version_update`

### HuaweiCloud SDK VPN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateVgwCertificate**
    - 请求参数变更
      - `* certificate: object<VpnGatewayCertificateRequestBodyContent> -> object<UpdateVpnGatewayCertificateRequestBodyContent>`
      - `* body: object<VpnGatewayCertificateRequestBody> -> object<UpdateVpnGatewayCertificateRequestBody>`
  - **CreateVgw**
    - 响应参数变更
      - `+ vpn_gateway.tags`
  - **UpdateVgw**
    - 响应参数变更
      - `+ vpn_gateway.tags`

# 0.1.80 2024-02-01

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListEvents**
    - 响应参数变更
      - `* events.annotations: map<string, string> -> map<string, object>`
      - `* events.attach_rule: map<string, string> -> map<string, object>`
  - **PushEvents**
    - 请求参数变更
      - `* events.annotations: map<string, string> -> map<string, object>`
      - `* events.attach_rule: map<string, string> -> map<string, object>`
  - **CreatePromInstance**
    - 请求参数变更
      - `+ region`

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`CreatePrivateProvider`、`CreatePrivateProviderVersion`
- _解决问题_
  - 无
- _特性变更_
  - **GetExecutionPlan**
    - 响应参数变更
      - `+ execution_plan_items.imported`
  - **GetExecutionPlanMetadata**
    - 响应参数变更
      - `+ summary.resource_import`

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowSubCustomerBudget**
    - 响应参数变更
      - `+ budget_type`
  - **UpdateSubCustomerBudget**
    - 请求参数变更
      - `+ budget_type`
  - **ListCustomerselfResourceRecordDetails**
    - 请求参数变更
      - `+ query_type`
      - `+ bill_cycle_begin`
      - `+ bill_cycle_end`

### HuaweiCloud SDK CC

- _新增特性_
  - 支持以下接口：
    - `ListGcbResourceTags`
    - `CreateGcbResourceTag`
    - `DeleteGcbResourceTag`
    - `BatchCreateGcbResourceTags`
    - `BatchDeleteGcbResourceTags`
    - `CountGcbResourceByTag`
    - `ListGcbResourceByTag`
    - `ListGcbTenantTags`
    - `ListGlobalConnectionBandwidths`
    - `CreateGlobalConnectionBandwidth`
    - `ShowGlobalConnectionBandwidth`
    - `UpdateGlobalConnectionBandwidth`
    - `DeleteGlobalConnectionBandwidth`
    - `AssociateGlobalConnectionBandwidthInstance`
    - `DisassociateGlobalConnectionBandwidthInstance`
    - `ListSupportBindingConnectionBandwidths`
    - `ListGlobalConnectionBandwidthConfigs`
    - `ListGlobalConnectionBandwidthSpecCodes`
    - `ListGlobalConnectionBandwidthSites`
    - `ListGlobalConnectionBandwidthLineLevels`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ShowLogs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DAS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListProcesses**
    - 请求参数变更
      - `+ node_id`

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateFactoryJob**
    - 请求参数变更
      - `+ single_node_job_type: enum value [NormalJob,OneClick]`
      - `- single_node_job_type: enum value [DataMigration]`
      - `+ nodes.type: enum value [DataMigration]`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateJob**
    - 请求参数变更
      - `+ singleNodeJobType: enum value [NormalJob,OneClick]`
      - `- singleNodeJobType: enum value [DataMigration]`
      - `+ nodes.type: enum value [DataMigration]`
  - **ShowJob**
    - 响应参数变更
      - `+ nodes.type: enum value [DataMigration]`
  - **UpdateJob**
    - 请求参数变更
      - `+ singleNodeJobType: enum value [NormalJob,OneClick]`
      - `- singleNodeJobType: enum value [DataMigration]`
      - `+ nodes.type: enum value [DataMigration]`
  - **CreateSupplementdata**
    - 请求参数变更
      - `+ dependJobs.singleNodeJobType: enum value [NormalJob,OneClick]`
      - `- dependJobs.singleNodeJobType: enum value [DataMigration]`
      - `+ dependJobs.nodes.type: enum value [DataMigration]`

### HuaweiCloud SDK DLI

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSparkJobs**
    - 请求参数变更
      - `+ job_name`
      - `- job-name`

### HuaweiCloud SDK DWS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDatabaseUser**
    - 响应参数变更
      - `+ user_type`
      - `+ name`
      - `+ logical_cluster`
      - `+ desc`
  - **ResizeClusterWithExistedNodes**
    - 请求参数变更
      - `- resize`
      - `- create_node_only`
      - `- is_scheduler_build_mode`
      - `- build_task_info`
      - `- order_id`
      - `- redis_conf.schedule_conf`
      - `- redis_conf.parallel_job`
      - `* redis_conf: object<RedisConf> -> object<RedisConfReq>`
  - **UpdateDatabaseAuthority**
    - 请求参数变更
      - `+ all_object`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateGaussMySqlInstance**
    - 请求参数变更
      - `+ tde_info`
  - **ShowGaussMySqlInstanceInfoUnifyStatus**
    - 响应参数变更
      - `+ instance.tde_info`
  - **ListGaussMySqlInstanceDetailInfoUnifyStatus**
    - 响应参数变更
      - `+ instances.tde_info`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CreateShrinkageJob`、`ShowShrinkCheckResult`、`UpdateSinkTaskQuota`
  - **CreateInstanceTopic**
    - 响应参数变更
      - `- id`

### HuaweiCloud SDK KooMessage

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAimTemplateReports**
    - 请求参数变更
      - `+ factory_type`
  - **CreateVmsTemplate**
    - 请求参数变更
      - `+ is_draft`
  - **ListVmsTemplateStatus**
    - 请求参数变更
      - `* offset: required -> optional`
      - `* limit: required -> optional`

### HuaweiCloud SDK MetaStudio

- _新增特性_
  - 支持接口`ListSmartLiveJobs`、`CopyVideoScripts`
- _解决问题_
  - 无
- _特性变更_
  - **ListRobot**
    - 响应参数变更
      - `+ llm_url`
      - `+ chat_rounds`
      - `+ is_stream`
      - `+ language`
      - `+ data.language`
      - `+ data.llm_url`
      - `+ data.is_stream`
      - `+ data.chat_rounds`
  - **CreateRobot**
    - 请求参数变更
      - `+ language`
      - `+ third_party_model_config`
  - **ShowRobot**
    - 响应参数变更
      - `+ llm_url`
      - `+ chat_rounds`
      - `+ is_stream`
      - `+ language`
  - **UpdateRobot**
    - 请求参数变更
      - `+ language`
      - `+ third_party_model_config`
  - **ListAssetSummary**
    - 响应参数变更
      - `+ asset_list.thumbnail_url`
  - **ListDigitalHumanVideo**
    - 响应参数变更
      - `+ jobs.job_type`
  - **CreateSmartChatRoom**
    - 请求参数变更
      - `+ layer_config.asset_id`
      - `+ layer_config.video_config.loop_count`
  - **ShowSmartChatRoom**
    - 响应参数变更
      - `+ layer_config.asset_id`
      - `+ layer_config.video_config.loop_count`
  - **UpdateSmartChatRoom**
    - 请求参数变更
      - `+ layer_config.asset_id`
      - `+ layer_config.video_config.loop_count`
    - 响应参数变更
      - `+ layer_config.asset_id`
      - `+ layer_config.video_config.loop_count`
  - **ShowSmartLive**
    - 响应参数变更
      - `+ room_id`
      - `+ cover_url`
      - `+ room_name`
  - **CheckTextLanguage**
    - 请求参数变更
      - `+ shoot_script.audio_config.asset_id`
  - **ShowAsset**
    - 响应参数变更
      - `+ app_user_id`
      - `- is_need_generate_cover`
      - `+ asset_extra_meta.voice_model_meta.order`
      - `+ asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.human_model_2d_meta.model_version: enum value [V3_2]`
  - **UpdateDigitalAsset**
    - 请求参数变更
      - `+ shared_config`
      - `- is_need_generate_cover`
      - `+ asset_extra_meta.voice_model_meta.order`
      - `+ asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.human_model_2d_meta.model_version: enum value [V3_2]`
    - 响应参数变更
      - `+ app_user_id`
      - `- is_need_generate_cover`
      - `+ asset_extra_meta.voice_model_meta.order`
      - `+ asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.human_model_2d_meta.model_version: enum value [V3_2]`
  - **StartSmartLive**
    - 请求参数变更
      - `+ view_mode`
      - `+ play_policy.play_mode: enum value [NO_PRESET]`
  - **ListSmartLive**
    - 响应参数变更
      - `+ room_id`
      - `+ cover_url`
      - `+ room_name`
      - `+ smart_live_jobs.room_id`
      - `+ smart_live_jobs.room_name`
      - `+ smart_live_jobs.cover_url`
  - **CreateDigitalAsset**
    - 请求参数变更
      - `+ shared_config`
      - `- is_need_generate_cover`
      - `+ asset_extra_meta.voice_model_meta.order`
      - `+ asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.human_model_2d_meta.model_version: enum value [V3_2]`
  - **ListAssets**
    - 请求参数变更
      - `+ is_movable`
      - `+ voice_provider`
      - `+ role`
    - 响应参数变更
      - `+ app_user_id`
      - `- is_need_generate_cover`
      - `+ assets.app_user_id`
      - `- assets.is_need_generate_cover`
      - `+ assets.asset_extra_meta.voice_model_meta.order`
      - `+ assets.asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ assets.asset_extra_meta.human_model_2d_meta.model_version: enum value [V3_2]`
  - **Create2DDigitalHumanVideo**
    - 请求参数变更
      - `+ shoot_scripts.shoot_script.layer_config.asset_id`
      - `+ shoot_scripts.shoot_script.layer_config.video_config.loop_count`
  - **Show2DDigitalHumanVideo**
    - 响应参数变更
      - `+ job_type`
      - `+ shoot_scripts.shoot_script.layer_config.asset_id`
      - `+ shoot_scripts.shoot_script.layer_config.video_config.loop_count`
  - **CreatePhotoDigitalHumanVideo**
    - 请求参数变更
      - `+ shoot_scripts.shoot_script.layer_config.asset_id`
      - `+ shoot_scripts.shoot_script.layer_config.video_config.loop_count`
  - **ShowPhotoDigitalHumanVideo**
    - 响应参数变更
      - `+ job_type`
      - `+ shoot_scripts.shoot_script.layer_config.asset_id`
      - `+ shoot_scripts.shoot_script.layer_config.video_config.loop_count`
  - **ListSmartLiveRooms**
    - 请求参数变更
      - `+ template_own_type`
    - 响应参数变更
      - `+ smart_live_rooms.project_id`
      - `+ smart_live_rooms.shared_config`
  - **CreateSmartLiveRoom**
    - 请求参数变更
      - `+ view_mode`
      - `+ play_policy.play_mode: enum value [NO_PRESET]`
      - `+ scene_scripts.layer_config.asset_id`
      - `+ scene_scripts.layer_config.video_config.loop_count`
      - `+ scene_scripts.shoot_scripts.audio_config.asset_id`
      - `+ interaction_rules.trigger.layer_config.asset_id`
  - **ShowSmartLiveRoom**
    - 响应参数变更
      - `+ view_mode`
      - `+ play_policy.play_mode: enum value [NO_PRESET]`
      - `+ scene_scripts.layer_config.asset_id`
      - `+ scene_scripts.layer_config.video_config.loop_count`
      - `+ scene_scripts.shoot_scripts.audio_config.asset_id`
      - `+ interaction_rules.trigger.layer_config.asset_id`
  - **UpdateSmartLiveRoom**
    - 请求参数变更
      - `+ view_mode`
      - `+ play_policy.play_mode: enum value [NO_PRESET]`
      - `+ scene_scripts.layer_config.asset_id`
      - `+ scene_scripts.layer_config.video_config.loop_count`
      - `+ scene_scripts.shoot_scripts.audio_config.asset_id`
      - `+ interaction_rules.trigger.layer_config.asset_id`
    - 响应参数变更
      - `+ view_mode`
      - `+ play_policy.play_mode: enum value [NO_PRESET]`
      - `+ scene_scripts.layer_config.asset_id`
      - `+ scene_scripts.layer_config.video_config.loop_count`
      - `+ scene_scripts.shoot_scripts.audio_config.asset_id`
      - `+ interaction_rules.trigger.layer_config.asset_id`
  - **UpdateInteractionRuleGroup**
    - 请求参数变更
      - `+ interaction_rules.trigger.layer_config.asset_id`
    - 响应参数变更
      - `+ interaction_rules.trigger.layer_config.asset_id`
  - **CreateVideoScripts**
    - 请求参数变更
      - `+ shoot_scripts.shoot_script.layer_config.asset_id`
      - `+ shoot_scripts.shoot_script.layer_config.video_config.loop_count`
  - **ShowVideoScript**
    - 响应参数变更
      - `+ shoot_scripts.shoot_script.layer_config.asset_id`
      - `+ shoot_scripts.shoot_script.layer_config.video_config.loop_count`
  - **UpdateVideoScript**
    - 请求参数变更
      - `+ shoot_scripts.shoot_script.layer_config.asset_id`
      - `+ shoot_scripts.shoot_script.layer_config.video_config.loop_count`
  - **CreateInteractionRuleGroup**
    - 请求参数变更
      - `+ interaction_rules.trigger.layer_config.asset_id`
  - **ListInteractionRuleGroups**
    - 响应参数变更
      - `+ interaction_rule_groups.interaction_rules.trigger.layer_config.asset_id`

### HuaweiCloud SDK NAT

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListNatGatewayDnatRules**
    - 请求参数变更
      - `+ marker`
  - **ListNatGatewaySnatRules**
    - 请求参数变更
      - `+ marker`
  - **ListNatGateways**
    - 请求参数变更
      - `+ marker`

### HuaweiCloud SDK OMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **StartTask**
    - 请求参数变更
      - `+ json_auth_file`
  - **StartTaskGroup**
    - 请求参数变更
      - `+ json_auth_file`
  - **RetryTaskGroup**
    - 请求参数变更
      - `+ json_auth_file`
  - **ShowBucketList**
    - 请求参数变更
      - `+ json_auth_file`
  - **ShowBucketRegion**
    - 请求参数变更
      - `+ json_auth_file`
  - **CheckPrefix**
    - 请求参数变更
      - `+ json_auth_file`
  - **ShowBucketObjects**
    - 请求参数变更
      - `+ json_auth_file`
  - **ShowCdnInfo**
    - 请求参数变更
      - `+ source_cdn.authentication_type: enum value [AZURE_SAS_TOKEN]`
  - **CreateSyncTask**
    - 请求参数变更
      - `+ dst_storage_policy`
      - `+ source_cdn.authentication_type: enum value [AZURE_SAS_TOKEN]`
  - **ShowTaskGroup**
    - 响应参数变更
      - `+ dst_storage_policy`
      - `+ src_node.cloud_type: enum value [GOOGLE]`
      - `+ src_node.list_file.list_file_num`
  - **CreateTask**
    - 请求参数变更
      - `+ dst_storage_policy`
      - `+ source_cdn.authentication_type: enum value [AZURE_SAS_TOKEN]`
      - `+ smn_config.message_template_name`
      - `+ src_node.json_auth_file`
      - `+ src_node.list_file.list_file_num`
  - **ListTasks**
    - 响应参数变更
      - `+ dst_storage_policy`
      - `+ tasks.dst_storage_policy`
      - `+ tasks.src_node.cloud_type: enum value [Google]`
      - `+ tasks.src_node.list_file.list_file_num`
  - **CreateTaskGroup**
    - 请求参数变更
      - `+ dst_storage_policy`
      - `+ smn_config.message_template_name`
      - `+ source_cdn.authentication_type: enum value [AZURE_SAS_TOKEN]`
      - `+ src_node.json_auth_file`
      - `+ src_node.list_file.list_file_num`
  - **ListTaskGroup**
    - 响应参数变更
      - `+ dst_storage_policy`
      - `+ taskgroups.dst_storage_policy`
      - `+ taskgroups.src_node.cloud_type: enum value [GOOGLE]`
      - `+ taskgroups.src_node.list_file.list_file_num`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持以下接口：
    - `StartInstanceReduceVolumeAction`
    - `UpdateHostPrivilege`
    - `ShowIncreBackupPolicy1`
    - `UpdateIncreBackupPolicy1`
    - `ListRdSforMySqlProxy`
    - `DeleteRdSforMySqlProxy`
    - `ModifyRdSforMySqlProxyRouteMode`
    - `RestartRdSforMysqlProxy`
    - `ListRdSforMysqlProxyFlavors`
    - `CreateRdSforMySqlProxy`
- _解决问题_
  - 无
- _特性变更_
  - **SetAuditlogPolicy**
    - 请求参数变更
      - `+ audit_types`
  - **CreateDbUser**
    - 请求参数变更
      - `+ is_privilege`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowGroup**
    - 响应参数变更
      - `* retry_max_time: number -> int32`
  - **UpdateConsumerGroup**
    - 请求参数变更
      - `* retry_max_time: number -> int32`
  - **ResetConsumeOffset**
    - 请求参数变更
      - `* timestamp: number -> string`
  - **CreateConsumerGroupOrBatchDeleteConsumerGroup**
    - 请求参数变更
      - `+ createdAt`
      - `+ permissions`
      - `+ consume_orderly`
      - `- from_beginning`
      - `* retry_max_time: number -> int32`
  - **ListInstanceConsumerGroups**
    - 响应参数变更
      - `+ groups.createdAt`
      - `+ groups.permissions`
      - `+ groups.consume_orderly`
      - `- groups.from_beginning`
      - `* groups.retry_max_time: number -> int32`
  - **BatchUpdateConsumerGroup**
    - 请求参数变更
      - `+ groups.createdAt`
      - `+ groups.permissions`
      - `+ groups.consume_orderly`
      - `- groups.from_beginning`
      - `* groups.retry_max_time: number -> int32`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListEvent**
    - 请求参数变更
      - `+ X-Language`
  - **ShowEvent**
    - 请求参数变更
      - `+ X-Language`

# 0.1.79 2024-01-25

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **DeleteserviceDiscoveryRules**
    - 响应参数变更
      - `+ id`

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDetailsOfApiV2**
    - 响应参数变更
      - `+ policy_https.req_protocol: enum value [GRPC]`
      - `+ backend_api.req_protocol: enum value [GRPC]`
  - **UpdateApiV2**
    - 请求参数变更
      - `+ policy_https.req_protocol: enum value [GRPC]`
      - `+ backend_api.req_protocol: enum value [GRPC]`
    - 响应参数变更
      - `+ policy_https.req_protocol: enum value [GRPC]`
      - `+ backend_api.req_protocol: enum value [GRPC]`
  - **ListApiVersionDetailV2**
    - 响应参数变更
      - `+ policy_https.req_protocol: enum value [GRPC]`
      - `+ backend_api.req_protocol: enum value [GRPC]`
  - **CreateApiV2**
    - 请求参数变更
      - `+ policy_https.req_protocol: enum value [GRPC]`
      - `+ backend_api.req_protocol: enum value [GRPC]`
    - 响应参数变更
      - `+ policy_https.req_protocol: enum value [GRPC]`
      - `+ backend_api.req_protocol: enum value [GRPC]`
  - **ListApisV2**
    - 响应参数变更
      - `+ apis.backend_api.req_protocol: enum value [GRPC]`

### HuaweiCloud SDK CSE

- _新增特性_
  - 支持以下接口：
    - `ShowPlugins`
    - `CreatePlugin`
    - `ShowSinglePlugin`
    - `ModifyPlugin`
    - `DeletePlugin`
    - `ShowHttp2Rpcs`
    - `CreateHttp2Rpc`
    - `ModifyHttp2Rpc`
    - `DeleteHttp2Rpc`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateFactoryJob**
    - 请求参数变更
      - `+ nodes.type: enum value [MRSFlinkJob]`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSystemTasks**
    - 响应参数变更
      - `+ subtasks`
      - `- subTasks`
      - `* startTime: string -> int64`
      - `* endTime: string -> int64`
      - `* lastUpdate: string -> int64`
  - **CreateJob**
    - 请求参数变更
      - `- singleNodeJobFlag`
      - `- approvers`
      - `- params`
      - `- directory`
      - `- singleNodeJobType`
      - `- schedule`
      - `- nodes`
      - `- logPath`
      - `- basicConfig`
      - `- name`
      - `- lastUpdateUser`
      - `- processType`
      - `- targetStatus`
      - `+ basicConfig.tags`
      - `+ nodes.type: enum value [MRSFlinkJob]`
  - **ShowJob**
    - 响应参数变更
      - `+ cleanOverdueDays`
      - `+ description`
      - `+ version`
      - `+ cleanWaitingJob`
      - `+ createTime`
      - `+ id`
      - `+ emptyRunningJob`
      - `- approvers`
      - `- targetStatus`
      - `+ basicConfig.tags`
      - `+ nodes.type: enum value [MRSFlinkJob]`
  - **UpdateJob**
    - 请求参数变更
      - `- singleNodeJobFlag`
      - `- approvers`
      - `- params`
      - `- directory`
      - `- singleNodeJobType`
      - `- schedule`
      - `- nodes`
      - `- logPath`
      - `- basicConfig`
      - `- name`
      - `- lastUpdateUser`
      - `- processType`
      - `- targetStatus`
      - `+ basicConfig.tags`
      - `+ nodes.type: enum value [MRSFlinkJob]`
  - **CreateSupplementdata**
    - 请求参数变更
      - `- singleNodeJobFlag`
      - `- approvers`
      - `- params`
      - `- directory`
      - `- singleNodeJobType`
      - `- schedule`
      - `- nodes`
      - `- logPath`
      - `- basicConfig`
      - `- name`
      - `- lastUpdateUser`
      - `- processType`
      - `- targetStatus`
      - `+ dependJobs.basicConfig.tags`
      - `+ dependJobs.nodes.type: enum value [MRSFlinkJob]`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateGaussMySqlReadonlyNode**
    - 请求参数变更
      - `+ availability_zones`

### HuaweiCloud SDK IEC

- _新增特性_
  - 支持接口`ListVolume`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeAutoClassification**
    - 请求参数变更
      - `+ detect_seal`
    - 响应参数变更
      - `+ result.seal_mark`
  - **RecognizeVatInvoice**
    - 响应参数变更
      - `+ result.province`
      - `+ result.city`
      - `+ result.belong_buyer_name`
      - `+ result.belong_seller_name`
      - `+ result.belong_vat_code`
      - `+ result.belong_number`
      - `+ result.belong_pages`
      - `+ result.belong_current_page`
      - `+ result.belong_remarks`
      - `+ result.belong_issue_date`
      - `+ result.sales_mark`
      - `+ result.belong_sum_amount`
      - `+ result.belong_sum_tax`
      - `+ result.belong_subtotal_amount`
      - `+ result.belong_subtotal_tax`
      - `+ result.belong_discount_amount`
      - `+ result.belong_discount_tax`
      - `+ result.belong_item_list`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListFlavors**
    - 请求参数变更
      - `+ is_serverless`
  - **ListInstances**
    - 响应参数变更
      - `+ instances.serverless_info`
  - **CreateInstance**
    - 请求参数变更
      - `+ serverless_info`
  - **CreateRestoreInstance**
    - 请求参数变更
      - `+ serverless_info`

### HuaweiCloud SDK SMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateTask**
    - 请求参数变更
      - `+ source_server.start_type`
  - **RegisterServer**
    - 请求参数变更
      - `+ start_type`

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateAssetByFileUpload**
    - 请求参数变更
      - `* category_id: string -> int32`
  - **CreateAssetProcessTask**
    - 请求参数变更
      - `+ hls_storage_type`
  - **ListTopStatistics**
    - 响应参数变更
      - `- top_urls.duration_ms`
  - **ListAssetList**
    - 响应参数变更
      - `- assets.duration_ms`
  - **ShowTakeOverAssetDetails**
    - 响应参数变更
      - `- base_info.meta_data.duration_ms`
  - **PublishAssets**
    - 响应参数变更
      - `- asset_info_array.base_info.meta_data.duration_ms`
  - **UnpublishAssets**
    - 响应参数变更
      - `- asset_info_array.base_info.meta_data.duration_ms`
  - **ShowAssetMeta**
    - 响应参数变更
      - `- asset_info_array.base_info.meta_data.duration_ms`
  - **ShowAssetDetail**
    - 响应参数变更
      - `- base_info.meta_data.duration_ms`
  - **ShowTakeOverTaskDetails**
    - 响应参数变更
      - `- assets.base_info.meta_data.duration_ms`

# 0.1.78 2024-01-18

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPromInstance**
    - 请求参数变更
      - `+ prom_status: enum value [DELETED,NORMAL,ALL]`
      - `+ prom_status: enum value [true,false]`
    - 响应参数变更
      - `- prometheus.cce_spec`
      - `- prometheus.prom_config`
      - `- prometheus.application`
      - `+ prometheus.prom_status: enum value [DELETED,NORMAL,ALL]`
  - **CreatePromInstance**
    - 请求参数变更
      - `- prom_id`
      - `- cce_spec`
      - `- prom_config`
      - `- prom_create_timestamp`
      - `- prom_update_timestamp`
      - `- prom_status`
      - `- is_deleted_tag`
      - `- deleted_time`
      - `- prom_spec_config`
      - `- cce_spec_config`
      - `- application`
      - `- prom_type: enum value [DEFAULT]`
      - `* body: object<PromInstanceEpsModel> -> object<PromInstanceRequestModel>`
    - 响应参数变更
      - `- prometheus.cce_spec`
      - `- prometheus.prom_config`
      - `- prometheus.is_deleted_tag`
      - `- prometheus.application`
      - `- prometheus.prom_type: enum value [DEFAULT]`
      - `+ prometheus.prom_status: enum value [DELETED,NORMAL,ALL]`
      - `* prometheus: list<PromInstanceEpsModel> -> list<PromInstanceEpsCreateModel>`
  - **AddOrUpdateServiceDiscoveryRules**
    - 响应参数变更
      - `+ id`
      - `+ results`

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDetailsOfInstanceV2**
    - 响应参数变更
      - `- public_ipv6_ips`
      - `- unreliable_ips`
      - `- node_ipv6_ips`
      - `- enable_fullstack_ipv6`
      - `- publicips.ipv6_address`
  - **UpdateInstanceV2**
    - 响应参数变更
      - `- public_ipv6_ips`
      - `- unreliable_ips`
      - `- node_ipv6_ips`
      - `- enable_fullstack_ipv6`
      - `- publicips.ipv6_address`
  - **CreateInstanceV2**
    - 请求参数变更
      - `+ spec_id: enum value [BASIC_IPV6,PROFESSIONAL_IPV6,ENTERPRISE_IPV6,PLATINUM_IPV6]`
  - **CreateOrder**
    - 请求参数变更
      - `+ instance_info.spec_id: enum value [BASIC_IPV6,PROFESSIONAL_IPV6,ENTERPRISE_IPV6,PLATINUM_IPV6]`

### HuaweiCloud SDK CAE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListComponentConfigurations**
    - 响应参数变更
      - `+ items.data.spec.items.ports.paths`
      - `- items.data.spec.items.ports.path`
      - `+ items.data.spec.volumes.umask`
  - **CreateComponentConfiguration**
    - 请求参数变更
      - `+ items.data.spec.items.ports.paths`
      - `- items.data.spec.items.ports.path`
      - `+ items.data.spec.volumes.umask`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowCluster**
    - 响应参数变更
      - `+ spec.enableAutopilot`
  - **UpdateCluster**
    - 响应参数变更
      - `+ spec.enableAutopilot`
  - **DeleteCluster**
    - 响应参数变更
      - `+ spec.enableAutopilot`
  - **CreateCluster**
    - 请求参数变更
      - `+ spec.enableAutopilot`
    - 响应参数变更
      - `+ spec.enableAutopilot`
  - **ListClusters**
    - 响应参数变更
      - `+ items.spec.enableAutopilot`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowTopUrl**
    - 请求参数变更
      - `+ service_area: enum value [global]`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ListCdnDomainTopRefers`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAgentConfig**
    - 响应参数变更
      - `+ result.extensions`
      - `- result.elasticsearch_enable`
      - `- result.elasticsearch_shadow_type`
      - `- result.elasticsearch_shadow_repository`
  - **UpdateAgentHealthStatus**
    - 响应参数变更
      - `+ result.result.extensions`
      - `- result.result.elasticsearch_enable`
      - `- result.result.elasticsearch_shadow_type`
      - `- result.result.elasticsearch_shadow_repository`
  - **ShowTask**
    - 响应参数变更
      - `+ taskInfo.related_temp_running_data.content_method_url`
  - **UpdateTask**
    - 请求参数变更
      - `+ related_temp_running_data.content_method_url`
    - 响应参数变更
      - `+ taskInfo.related_temp_running_data.content_method_url`
  - **UpdateTaskRelatedTestCase**
    - 响应参数变更
      - `+ taskInfo.related_temp_running_data.content_method_url`

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 支持以下接口：
    - `ShowDataDetail`
    - `ShowDatamapLineage`
    - `ShowLineageBulk`
    - `ShowNodes`
    - `ShowInstanceInfos`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ChangeMasterStandbyAsync`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowScript**
    - 响应参数变更
      - `+ owner`
      - `* configuration: string -> map<string, object>`
  - **UpdateScript**
    - 请求参数变更
      - `+ owner`
      - `* configuration: string -> map<string, object>`
  - **CreateScript**
    - 请求参数变更
      - `+ owner`
      - `* configuration: string -> map<string, object>`
  - **ListScripts**
    - 响应参数变更
      - `+ owner`
      - `+ scripts.owner`
      - `* scripts.configuration: string -> map<string, object>`
  - **CreateJob**
    - 请求参数变更
      - `+ basicConfig.agency`
      - `+ basicConfig.isIgnoreWaiting`
  - **ShowJob**
    - 响应参数变更
      - `+ basicConfig.agency`
      - `+ basicConfig.isIgnoreWaiting`
  - **UpdateJob**
    - 请求参数变更
      - `+ basicConfig.agency`
      - `+ basicConfig.isIgnoreWaiting`
  - **CreateSupplementdata**
    - 请求参数变更
      - `+ dependJobs.basicConfig.agency`
      - `+ dependJobs.basicConfig.isIgnoreWaiting`

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持接口`DeleteScanJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持接口`ListWorkloadPlans`
- _解决问题_
  - 无
- _特性变更_
  - **ShowWorkloadPlan**
    - 响应参数变更
      - `* workload_plan.status: string -> int32`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateListener**
    - 请求参数变更
      - `+ listener.protection_status`
      - `+ listener.protection_reason`
  - **ListLoadbalancers**
    - 响应参数变更
      - `+ loadbalancers.billing_info`
      - `+ loadbalancers.protection_status`
      - `+ loadbalancers.protection_reason`
  - **CreateLoadbalancer**
    - 响应参数变更
      - `+ loadbalancer_id`
      - `+ order_id`
      - `+ loadbalancer.billing_info`
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
  - **ShowLoadbalancer**
    - 响应参数变更
      - `+ loadbalancer.billing_info`
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
  - **UpdateLoadbalancer**
    - 响应参数变更
      - `+ loadbalancer.billing_info`
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`

### HuaweiCloud SDK MetaStudio

- _新增特性_
  - 支持以下接口：
    - `CreateDialogUrl`
    - `StartSmartChatJob`
    - `StopSmartChatJob`
    - `ShowSmartChatJob`
    - `CreateOnceCode`
    - `ListRobot`
    - `CreateRobot`
    - `ShowRobot`
    - `UpdateRobot`
    - `DeleteRobot`
    - `ListSmartChatRooms`
    - `CreateSmartChatRoom`
    - `ShowSmartChatRoom`
    - `UpdateSmartChatRoom`
    - `DeleteSmartChatRoom`
- _解决问题_
  - 无
- _特性变更_
  - **ListDigitalHumanVideo**
    - 请求参数变更
      - `+ job_type`
  - **ListSmartLiveRooms**
    - 响应参数变更
      - `+ smart_live_rooms.last_job_status: enum value [BLOCKED]`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeSeal`
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeFlightItinerary**
    - 响应参数变更
      - `+ result.tax`
      - `+ result.tax_rate`
      - `+ result.buyer_name`
      - `+ result.buyer_id`
      - `+ result.number`
      - `+ result.international_flag`
      - `+ result.issue_status`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAuditlogPolicy**
    - 响应参数变更
      - `+ all_audit_log_action`
      - `+ audit_types`
  - **ListRecycleInstances**
    - 响应参数变更
      - `+ instances.is_serverless`

### HuaweiCloud SDK RGC

- _新增特性_
  - 支持接口`ReRegisterOrganizationalUnit`、`EnrollAccount`、`UnEnrollAccount`
- _解决问题_
  - 无
- _特性变更_
  - **CheckLaunch**
    - 请求参数变更
      - `- X-Security-Token`
  - **ShowHomeRegion**
    - 请求参数变更
      - `- X-Security-Token`
  - **ShowAvailableUpdates**
    - 请求参数变更
      - `- X-Security-Token`
  - **ShowOperation**
    - 请求参数变更
      - `- X-Security-Token`
  - **ShowLandingZoneStatus**
    - 请求参数变更
      - `- X-Security-Token`
  - **ShowLandingZoneIdentityCenter**
    - 请求参数变更
      - `- X-Security-Token`
  - **ListControls**
    - 请求参数变更
      - `- X-Security-Token`
  - **ListControlViolations**
    - 请求参数变更
      - `- X-Security-Token`
  - **ShowControl**
    - 请求参数变更
      - `- X-Security-Token`
  - **ListEnabledControls**
    - 请求参数变更
      - `- X-Security-Token`
  - **ListDriftDetails**
    - 请求参数变更
      - `- X-Security-Token`
  - **ListManagedOrganizationalUnits**
    - 请求参数变更
      - `- X-Security-Token`
    - 响应参数变更
      - `+ landing_zone_version`
      - `+ message`
      - `+ managed_organization_units.landing_zone_version`
      - `+ managed_organization_units.message`
  - **CreateAccount**
    - 请求参数变更
      - `- X-Security-Token`
  - **ListManagedAccounts**
    - 请求参数变更
      - `- X-Security-Token`
    - 响应参数变更
      - `+ landing_zone_version`
      - `- account_email`
      - `- phone`
      - `- identity_store_email_name`
      - `+ managed_accounts.landing_zone_version`
      - `- managed_accounts.phone`
      - `- managed_accounts.account_email`
      - `- managed_accounts.identity_store_email_name`
  - **SetupLandingZone**
    - 请求参数变更
      - `- X-Security-Token`
      - `+ organization_structure.accounts.account_id`
  - **ShowLandingZoneConfiguration**
    - 请求参数变更
      - `- X-Security-Token`
    - 响应参数变更
      - `- identity_store_email`
      - `+ organization_structure.accounts.account_id`
      - `- organization_structure.accounts.phone`
      - `* organization_structure.accounts: list<AccountBaseline> -> list<AccountBaselineRsp>`
      - `* organization_structure: list<OrganizationStructureBaseLine> -> list<OrganizationStructureBaseLineRsp>`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ValidateConsumedMessage**
    - 请求参数变更
      - `+ topic`

### HuaweiCloud SDK VPN

- _新增特性_
  - 支持以下接口：
    - `BatchCreateResourceTags`
    - `BatchDeleteResourceTags`
    - `ListResourcesByTags`
    - `CountResourcesByTags`
    - `ShowResourceTags`
    - `ListProjectTags`
- _解决问题_
  - 无
- _特性变更_
  - **ListCgws**
    - 响应参数变更
      - `+ customer_gateways.id_type`
      - `+ customer_gateways.id_value`
      - `+ customer_gateways.tags`
      - `- customer_gateways.route_mode`
      - `- customer_gateways.ip`
      - `+ customer_gateways.ca_certificate.id`
  - **CreateCgw**
    - 请求参数变更
      - `+ customer_gateway.id_type`
      - `+ customer_gateway.id_value`
      - `+ customer_gateway.tags`
      - `- customer_gateway.route_mode`
      - `- customer_gateway.ip`
      - `+ customer_gateway.ca_certificate.id`
    - 响应参数变更
      - `+ customer_gateway.id_type`
      - `+ customer_gateway.id_value`
      - `+ customer_gateway.tags`
      - `- customer_gateway.route_mode`
      - `- customer_gateway.ip`
      - `+ customer_gateway.ca_certificate.id`
  - **ShowCgw**
    - 响应参数变更
      - `+ customer_gateway.id_type`
      - `+ customer_gateway.id_value`
      - `+ customer_gateway.tags`
      - `- customer_gateway.route_mode`
      - `- customer_gateway.ip`
      - `+ customer_gateway.ca_certificate.id`
  - **UpdateCgw**
    - 请求参数变更
      - `+ customer_gateway.ca_certificate`
    - 响应参数变更
      - `+ customer_gateway.id_type`
      - `+ customer_gateway.id_value`
      - `+ customer_gateway.tags`
      - `- customer_gateway.route_mode`
      - `- customer_gateway.ip`
      - `+ customer_gateway.ca_certificate.id`
  - **ListVgws**
    - 响应参数变更
      - `+ vpn_gateways.policy_template`
      - `+ vpn_gateways.supported_flavors`
      - `+ vpn_gateways.tags`
      - `- vpn_gateways.access_private_ips`
      - `- vpn_gateways.master_eip`
      - `- vpn_gateways.slave_eip`
  - **CreateVgw**
    - 请求参数变更
      - `+ vpn_gateway.access_private_ip_1`
      - `+ vpn_gateway.access_private_ip_2`
      - `+ vpn_gateway.tags`
      - `- vpn_gateway.master_eip`
      - `- vpn_gateway.slave_eip`
      - `+ vpn_gateway.flavor: enum value [Professional1-NonFixedIP,Professional2-NonFixedIP]`
    - 响应参数变更
      - `+ vpn_gateway.certificate_id`
      - `+ vpn_gateway.ha_mode`
      - `+ vpn_gateway.policy_template`
  - **ListAvailabilityZones**
    - 响应参数变更
      - `+ availability_zones.Professional1-NonFixedIP`
      - `+ availability_zones.Professional2-NonFixedIP`
  - **ShowVgw**
    - 响应参数变更
      - `+ vpn_gateway.policy_template`
      - `+ vpn_gateway.supported_flavors`
      - `+ vpn_gateway.tags`
      - `- vpn_gateway.access_private_ips`
      - `- vpn_gateway.master_eip`
      - `- vpn_gateway.slave_eip`
  - **UpdateVgw**
    - 请求参数变更
      - `+ vpn_gateway.policy_template`
      - `- vpn_gateway.master_eip_id`
      - `- vpn_gateway.slave_eip_id`
    - 响应参数变更
      - `+ vpn_gateway.certificate_id`
      - `+ vpn_gateway.policy_template`
      - `- vpn_gateway.access_private_ips`
      - `- vpn_gateway.master_eip`
      - `- vpn_gateway.slave_eip`
  - **ListVpnConnections**
    - 响应参数变更
      - `+ vpn_connections.tags`
  - **CreateVpnConnection**
    - 请求参数变更
      - `+ vpn_connection.tags`
    - 响应参数变更
      - `+ vpn_connection.tags`
  - **ShowVpnConnection**
    - 响应参数变更
      - `+ vpn_connection.tags`
  - **UpdateVpnConnection**
    - 响应参数变更
      - `+ vpn_connection.tags`

### HuaweiCloud SDK Workspace

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListItaSubJobs**
    - 响应参数变更
      - `+ jobs.process`
      - `- jobs.desktop_name`
      - `- jobs.ip_address`
      - `- jobs.mac_address`
      - `+ jobs.entities.desktop_name`
      - `+ jobs.entities.ip_address`

### HuaweiCloud SDK WorkspaceApp

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListServers**
    - 响应参数变更
      - `+ items.job_type: enum value [COLLECT_HDA_LOG]`
  - **ShowJobDetail**
    - 响应参数变更
      - `+ job_type: enum value [COLLECT_HDA_LOG]`

# 0.1.77 2024-01-11

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDetailsOfInstanceV2**
    - 响应参数变更
      - `+ public_ipv6_ips`
      - `+ unreliable_ips`
      - `+ node_ipv6_ips`
      - `+ enable_fullstack_ipv6`
      - `+ publicips.ipv6_address`
  - **UpdateInstanceV2**
    - 响应参数变更
      - `+ public_ipv6_ips`
      - `+ unreliable_ips`
      - `+ node_ipv6_ips`
      - `+ enable_fullstack_ipv6`
      - `+ publicips.ipv6_address`
  - **CreateInstanceV2**
    - 请求参数变更
      - `- spec_id: enum value [BASIC_IPV6,PROFESSIONAL_IPV6,ENTERPRISE_IPV6,PLATINUM_IPV6]`
  - **CreateOrder**
    - 请求参数变更
      - `- instance_info.spec_id: enum value [BASIC_IPV6,PROFESSIONAL_IPV6,ENTERPRISE_IPV6,PLATINUM_IPV6]`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateNodePool**
    - 响应参数变更
      - `- status.jobId`
      - `* status: object<NodePoolStatus> -> object<UpdateNodePoolStatus>`
  - **CreateNodePool**
    - 响应参数变更
      - `- status.jobId`
      - `* status: object<NodePoolStatus> -> object<CreateNodePoolStatus>`

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 支持接口`ImportLineage`、`ListAllTables`、`ImportCatalogs`、`ImportModels`
- _解决问题_
  - 无
- _特性变更_
  - **ShowEntities**
    - 响应参数变更
      - `* entities.connection: list<Connection> -> object<Connection>`
  - **ShowMetricAssets**
    - 响应参数变更
      - `* entities.connection: list<Connection> -> object<Connection>`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchValidateConnections**
    - 请求参数变更
      - `+ jobs.customized_dns`

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持以下接口：
    - `BatchCreateTags`
    - `BatchDeleteTags`
    - `ListInstanceByTags`
    - `CountInstanceByTags`
    - `ListInstanceTags`
    - `ListTags`
    - `UpdateJobConfigurations`
    - `ListJobParameters`
    - `ListJobHistoryParameters`
- _解决问题_
  - 无
- _特性变更_
  - **ListLinks**
    - 响应参数变更
      - `+ job_links.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
  - **ListJobs**
    - 请求参数变更
      - `+ engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
    - 响应参数变更
      - `+ jobs.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
      - `+ jobs.children.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
  - **CreateJob**
    - 请求参数变更
      - `+ job.base_info.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
      - `+ job.source_endpoint.customized_dns`
      - `+ job.source_endpoint.db_type: enum value [redis,rediscluster,gaussredis]`
      - `+ job.source_endpoint.endpoint.endpoint_name: enum value [redis,ecs_redis,rediscluster,ecs_rediscluster,cloud_gaussdb_redis]`
      - `+ job.source_endpoint.config.node_num`
  - **BatchCreateJobsAsync**
    - 请求参数变更
      - `+ jobs.base_info.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
      - `+ jobs.source_endpoint.customized_dns`
      - `+ jobs.source_endpoint.db_type: enum value [redis,rediscluster,gaussredis]`
      - `+ jobs.source_endpoint.endpoint.endpoint_name: enum value [redis,ecs_redis,rediscluster,ecs_rediscluster,cloud_gaussdb_redis]`
      - `+ jobs.source_endpoint.config.node_num`
  - **ListAsyncJobDetail**
    - 响应参数变更
      - `+ jobs.base_info.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
      - `+ jobs.source_endpoint.customized_dns`
      - `+ jobs.source_endpoint.db_type: enum value [redis,rediscluster,gaussredis]`
      - `+ jobs.source_endpoint.endpoint.endpoint_name: enum value [redis,ecs_redis,rediscluster,ecs_rediscluster,cloud_gaussdb_redis]`
      - `+ jobs.source_endpoint.config.node_num`
  - **UpdateBatchAsyncJobs**
    - 请求参数变更
      - `+ jobs.params.base_info.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
      - `+ jobs.params.source_endpoint.customized_dns`
      - `+ jobs.params.source_endpoint.db_type: enum value [redis,rediscluster,gaussredis]`
      - `+ jobs.params.source_endpoint.endpoint.endpoint_name: enum value [redis,ecs_redis,rediscluster,ecs_rediscluster,cloud_gaussdb_redis]`
      - `+ jobs.params.source_endpoint.config.node_num`
  - **ShowJobDetail**
    - 响应参数变更
      - `+ job.base_info.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
      - `+ job.source_endpoint.customized_dns`
      - `+ job.source_endpoint.db_type: enum value [redis,rediscluster,gaussredis]`
      - `+ job.source_endpoint.endpoint.endpoint_name: enum value [redis,ecs_redis,rediscluster,ecs_rediscluster,cloud_gaussdb_redis]`
      - `+ job.source_endpoint.config.node_num`
  - **UpdateJob**
    - 请求参数变更
      - `+ job.params.base_info.engine_type: enum value [redis-to-gaussredis,rediscluster-to-gaussredis]`
      - `+ job.params.source_endpoint.customized_dns`
      - `+ job.params.source_endpoint.db_type: enum value [redis,rediscluster,gaussredis]`
      - `+ job.params.source_endpoint.endpoint.endpoint_name: enum value [redis,ecs_redis,rediscluster,ecs_rediscluster,cloud_gaussdb_redis]`
      - `+ job.params.source_endpoint.config.node_num`
  - **ExecuteJobAction**
    - 请求参数变更
      - `+ job.action_name: enum value [column_limit,reload_parameters]`
      - `+ job.action_params.endpoints.customized_dns`
      - `+ job.action_params.endpoints.db_type: enum value [redis,rediscluster,gaussredis]`
      - `+ job.action_params.endpoints.endpoint.endpoint_name: enum value [redis,ecs_redis,rediscluster,ecs_rediscluster,cloud_gaussdb_redis]`
      - `+ job.action_params.endpoints.config.node_num`
  - **BatchExecuteJobActions**
    - 请求参数变更
      - `+ jobs.action_name: enum value [column_limit,reload_parameters]`
      - `+ jobs.action_params.endpoints.customized_dns`
      - `+ jobs.action_params.endpoints.db_type: enum value [redis,rediscluster,gaussredis]`
      - `+ jobs.action_params.endpoints.endpoint.endpoint_name: enum value [redis,ecs_redis,rediscluster,ecs_rediscluster,cloud_gaussdb_redis]`
      - `+ jobs.action_params.endpoints.config.node_num`

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持以下接口：
    - `ListWorkloadQueueUsers`
    - `ShowWorkloadQueue`
    - `ShowWorkloadPlanStage`
    - `DeleteWorkloadPlanStage`
    - `ListPlanExecLogs`
    - `DeleteQueueUserList`
    - `SwitchPlanStage`
    - `AddQueueUserList`
    - `UpdateQueueResources`
    - `ListSchemas`
    - `UpdateSchemas`
    - `ShowResourceStatistics`
    - `AddWorkloadPlanStage`
- _解决问题_
  - 无
- _特性变更_
  - **CreateClusterV2**
    - 请求参数变更
      - `* cluster.tags: object<Tags> -> list<Tags>`

### HuaweiCloud SDK EC

- _新增特性_
  - 支持接口`ListEcnWithVpc`、`AddEcnWithVpc`、`UpdateEcnWithVpc`、`DeleteEcnWithVpc`
- _解决问题_
  - 无
- _特性变更_
  - **ShowEcnInfo**
    - 响应参数变更
      - `+ vpc_ids`
  - **UpdateEcn**
    - 响应参数变更
      - `+ vpc_ids`
  - **UpdateEcnAccessPoint**
    - 响应参数变更
      - `+ attach_vpc_count`
  - **ShowEquipmentOspf**
    - 响应参数变更
      - `+ cloud_subnet_list`
  - **UpdateEquipmentOspf**
    - 请求参数变更
      - `+ cloud_subnet_list`
    - 响应参数变更
      - `+ cloud_subnet_list`
  - **ListEcn**
    - 响应参数变更
      - `+ vpc_ids`
      - `+ enterprise_connect_networks.vpc_ids`
  - **CreateEcnAccessPoint**
    - 响应参数变更
      - `+ attach_vpc_count`
  - **ListEcnAccessPointByEcnId**
    - 响应参数变更
      - `+ attach_vpc_count`
      - `+ access_points.attach_vpc_count`

### HuaweiCloud SDK ECS

- _新增特性_
  - 支持接口`NovaShowServerInterface`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateServer**
    - 响应参数变更
      - `+ server.OS-EXT-SRV-ATTR:user_data`

### HuaweiCloud SDK EG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowEventStreaming**
    - 响应参数变更
      - `+ id`
      - `- streaming_id`
  - **ListEventStreaming**
    - 响应参数变更
      - `+ id`
      - `- streaming_id`
      - `+ items.id`
      - `- items.streaming_id`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowCertificate**
    - 响应参数变更
      - `+ tenant_id`
      - `+ create_time`
      - `+ certificate`
      - `+ description`
      - `+ expire_time`
      - `+ private_key`
      - `+ type`
      - `+ update_time`
      - `+ admin_state_up`
      - `+ domain`
      - `+ name`
      - `+ id`
  - **UpdateCertificate**
    - 响应参数变更
      - `+ tenant_id`
      - `+ create_time`
      - `+ certificate`
      - `+ description`
      - `+ expire_time`
      - `+ private_key`
      - `+ type`
      - `+ update_time`
      - `+ admin_state_up`
      - `+ domain`
      - `+ name`
      - `+ id`
  - **CreateCertificate**
    - 响应参数变更
      - `+ tenant_id`
      - `+ create_time`
      - `+ certificate`
      - `+ description`
      - `+ expire_time`
      - `+ private_key`
      - `+ type`
      - `+ update_time`
      - `+ admin_state_up`
      - `+ domain`
      - `+ name`
      - `+ id`
  - **ListHealthmonitors**
    - 响应参数变更
      - `+ healthmonitors`
  - **CreateHealthmonitor**
    - 响应参数变更
      - `+ healthmonitor`
  - **ShowHealthmonitors**
    - 响应参数变更
      - `+ healthmonitor`
  - **UpdateHealthmonitor**
    - 响应参数变更
      - `+ healthmonitor`
  - **ListMembers**
    - 响应参数变更
      - `+ members`
  - **CreateMember**
    - 响应参数变更
      - `+ member`
  - **ShowMember**
    - 响应参数变更
      - `+ member`
  - **UpdateMember**
    - 响应参数变更
      - `+ member`
  - **ShowLoadbalancer**
    - 响应参数变更
      - `+ loadbalancer`
  - **UpdateLoadbalancer**
    - 请求参数变更
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
    - 响应参数变更
      - `+ loadbalancer`
  - **ListL7rules**
    - 响应参数变更
      - `+ rules`
  - **CreateL7rule**
    - 响应参数变更
      - `+ rule`
  - **ShowL7rule**
    - 响应参数变更
      - `+ rule`
  - **UpdateL7rule**
    - 响应参数变更
      - `+ rule`
  - **ListPools**
    - 响应参数变更
      - `+ pools`
  - **CreatePool**
    - 请求参数变更
      - `+ pool.protection_status`
      - `+ pool.protection_reason`
    - 响应参数变更
      - `+ pool`
  - **ShowPool**
    - 响应参数变更
      - `+ pool`
  - **UpdatePool**
    - 请求参数变更
      - `+ pool.protection_status`
      - `+ pool.protection_reason`
    - 响应参数变更
      - `+ pool`
  - **ListListeners**
    - 响应参数变更
      - `+ listeners`
  - **CreateListener**
    - 响应参数变更
      - `+ listener`
  - **ShowListener**
    - 响应参数变更
      - `+ listener`
  - **UpdateListener**
    - 请求参数变更
      - `+ listener.protection_status`
      - `+ listener.protection_reason`
    - 响应参数变更
      - `+ listener`
  - **CreateLoadbalancer**
    - 请求参数变更
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
    - 响应参数变更
      - `+ loadbalancer`
  - **ListLoadbalancers**
    - 响应参数变更
      - `+ loadbalancers.publicips`
      - `+ loadbalancers.charge_mode`
      - `+ loadbalancers.frozen_scene`
  - **ListL7policies**
    - 响应参数变更
      - `+ l7policies`
  - **CreateL7policy**
    - 响应参数变更
      - `+ l7policy`
  - **UpdateL7policies**
    - 响应参数变更
      - `+ l7policy`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPools**
    - 请求参数变更
      - `+ connection_drain`
    - 响应参数变更
      - `+ pools.connection_drain`
  - **CreatePool**
    - 请求参数变更
      - `+ pool.connection_drain`
    - 响应参数变更
      - `+ pool.connection_drain`
  - **ShowPool**
    - 响应参数变更
      - `+ pool.connection_drain`
  - **UpdatePool**
    - 请求参数变更
      - `+ pool.connection_drain`
    - 响应参数变更
      - `+ pool.connection_drain`
  - **ListMasterSlavePools**
    - 请求参数变更
      - `+ connection_drain`
    - 响应参数变更
      - `+ pools.connection_drain`
  - **CreateMasterSlavePool**
    - 请求参数变更
      - `+ pool.connection_drain`
    - 响应参数变更
      - `+ pool.connection_drain`
  - **ShowMasterSlavePool**
    - 响应参数变更
      - `+ pool.connection_drain`

### HuaweiCloud SDK GES

- _新增特性_
  - 支持接口`ExportBackup2`、`ImportBackup2`
- _解决问题_
  - 无
- _特性变更_
  - **ListGraphs2**
    - 响应参数变更
      - `+ graphs.capacity_ratio`
  - **CreateGraph2**
    - 请求参数变更
      - `+ graph.capacity_ratio`
  - **ShowGraph2**
    - 响应参数变更
      - `+ graph.capacity_ratio`

### HuaweiCloud SDK IEC

- _新增特性_
  - 支持接口`AttachVipBandwidth`、`DetachVipBandwidth`
- _解决问题_
  - 无
- _特性变更_
  - **AddNics**
    - 请求参数变更
      - `+ ipv6_enable`
      - `+ ipv6_bandwidth`
  - **CreateRoutes**
    - 请求参数变更
      - `+ routes.type: enum value [cc,igw]`
  - **UpdateRoutes**
    - 请求参数变更
      - `+ routes.type: enum value [cc,igw]`
  - **ListImages**
    - 响应参数变更
      - `+ images.__support_amd`
  - **ListFlavors**
    - 响应参数变更
      - `+ flavors.os_extra_specs.cond:network`
  - **ListSites**
    - 响应参数变更
      - `+ sites.pools.allow_share_bandwidth_types`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateReassignmentTask**
    - 请求参数变更
      - `+ is_schedule`
      - `+ execute_at`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListTranscodingTask**
    - 响应参数变更
      - `+ task_array.progress`
  - **ListTranscodeDetail**
    - 响应参数变更
      - `+ task_array.progress`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeIdCard**
    - 请求参数变更
      - `+ detect_tampering`
      - `+ detect_border_integrity`
      - `+ detect_blocking_within_border`
      - `+ detect_blur`
      - `+ detect_interim`
      - `+ detect_glare`
    - 响应参数变更
      - `+ result.detect_tampering_result`
      - `+ result.detect_border_integrity_result`
      - `+ result.detect_blocking_within_border_result`
      - `+ result.detect_blur_result`
      - `+ result.detect_interim_result`
      - `+ result.detect_glare_result`
      - `+ result.score_info`
      - `+ result.front`
      - `+ result.back`
      - `+ result.verification_result.valid_validity_period`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstances**
    - 请求参数变更
      - `+ eps_id`
  - **BatchRestoreDatabase**
    - 请求参数变更
      - `+ instances.is_fast_restore`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAddressGroup**
    - 响应参数变更
      - `+ address_group.ip_extra_set`
  - **UpdateAddressGroup**
    - 请求参数变更
      - `+ address_group.ip_extra_set`
    - 响应参数变更
      - `+ address_group.ip_extra_set`
  - **ListAddressGroup**
    - 响应参数变更
      - `+ address_groups.ip_extra_set`
  - **CreateAddressGroup**
    - 请求参数变更
      - `+ address_group.ip_extra_set`
    - 响应参数变更
      - `+ address_group.ip_extra_set`

### HuaweiCloud SDK Workspace

- _新增特性_
  - 支持接口`ListDesktopNamePolicy`、`CreateDesktopNamePolicy`、`UpdateDesktopNamePolicy`、`BatchDeleteDesktopNamePolicy`
- _解决问题_
  - 无
- _特性变更_
  - **ListAccessPolicies**
    - 响应参数变更
      - `- policies.policy_name: enum value [PRIVATE_ACCESS]`
  - **CreateAccessPolicy**
    - 请求参数变更
      - `- policy.policy_name: enum value [PRIVATE_ACCESS]`
  - **CreateDesktop**
    - 请求参数变更
      - `+ desktop_name_policy_id`
  - **ListUsedDesktopInfo**
    - 请求参数变更
      - `+ group_by_type`
  - **ListWorkspaces**
    - 响应参数变更
      - `+ authorized_collect_log`
      - `+ authorized_hda_upgrade`
  - **UpdateWorkspace**
    - 请求参数变更
      - `+ authorized_collect_log`
      - `+ authorized_hda_upgrade`
      - `+ apply_dedicated_standby_network_param`

# 0.1.76 2024-01-04

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListNotifiedHistories**
    - 响应参数变更
      - `+ notified_histories`
      - `- event_sn`
      - `- notifications`
  - **ListEvent2alarmRule**
    - 响应参数变更
      - `+ event_names`
      - `+ migrated`
      - `+ topics`
      - `+ trigger_policies.trigger_type: enum value [immediately]`
      - `- trigger_policies.trigger_type: enum value [notification]`
      - `+ alarm_type: enum value [notification,denoising]`
      - `* : list<event2alarmRuleBody> -> list<Event2alarmRuleBody>`
  - **UpdateEventRule**
    - 请求参数变更
      - `+ event_names`
      - `+ migrated`
      - `+ topics`
      - `+ trigger_policies.trigger_type: enum value [immediately]`
      - `- trigger_policies.trigger_type: enum value [notification]`
      - `+ alarm_type: enum value [notification,denoising]`
      - `* body: object<event2alarmRuleBody> -> object<Event2alarmRuleBody>`
  - **AddEvent2alarmRule**
    - 请求参数变更
      - `+ event_names`
      - `+ migrated`
      - `+ topics`
      - `+ trigger_policies.trigger_type: enum value [immediately]`
      - `- trigger_policies.trigger_type: enum value [notification]`
      - `+ alarm_type: enum value [notification,denoising]`
      - `* body: object<event2alarmRuleBody> -> object<Event2alarmRuleBody>`
  - **ListMuteRule**
    - 响应参数变更
      - `* mute_config.scope: list<string> -> list<integer>`
  - **UpdateMuteRule**
    - 请求参数变更
      - `* mute_config.scope: list<string> -> list<integer>`
  - **AddMuteRules**
    - 请求参数变更
      - `* mute_config.scope: list<string> -> list<integer>`
  - **ShowActionRule**
    - 响应参数变更
      - `+ type: enum value [1,2]`
      - `- type: enum value ["1","2"]`
  - **ListEvents**
    - 响应参数变更
      - `* events.metadata: object -> map<string, string>`
      - `* events.annotations: object -> map<string, string>`
      - `* events.attach_rule: object -> map<string, string>`
  - **PushEvents**
    - 请求参数变更
      - `* events.metadata: object -> map<string, string>`
      - `* events.annotations: object -> map<string, string>`
      - `* events.attach_rule: object -> map<string, string>`
  - **UpdateActionRule**
    - 请求参数变更
      - `+ type: enum value [1,2]`
      - `- type: enum value ["1","2"]`
  - **AddActionRule**
    - 请求参数变更
      - `+ type: enum value [1,2]`
      - `- type: enum value ["1","2"]`
  - **ListActionRule**
    - 响应参数变更
      - `+ action_rules.type: enum value [1,2]`
      - `- action_rules.type: enum value ["1","2"]`
  - **ListMetricItems**
    - 响应参数变更
      - `+ metaData.offset`
      - `+ metaData.nextToken`
      - `- metaData.start`
      - `* metaData: object<metaData> -> object<MetaDataSeries>`
  - **UpdateAlarmRule**
    - 请求参数变更
      - `+ is_turn_on`
      - `- id_turn_on`
      - `+ comparison_operator: enum value [<,>,<=,>=]`
      - `+ period: enum value [60000,300000,900000,36000000]`
  - **AddAlarmRule**
    - 请求参数变更
      - `+ is_turn_on`
      - `- id_turn_on`
      - `+ comparison_operator: enum value [<,>,<=,>=]`
      - `+ period: enum value [60000,300000,900000,36000000]`
  - **ListServiceDiscoveryRules**
    - 响应参数变更
      - `+ id`
      - `+ appRules.desc`
      - `+ appRules.spec.dataSource`
      - `+ appRules.spec.editable`
      - `+ appRules.spec.aom_metric_relabel_configs`
  - **AddOrUpdateServiceDiscoveryRules**
    - 请求参数变更
      - `+ appRules.desc`
      - `+ appRules.spec.dataSource`
      - `+ appRules.spec.editable`
      - `+ appRules.spec.aom_metric_relabel_configs`

### HuaweiCloud SDK CodeArtsPipeline

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `ShowProjectOpenSourceStrategy`
    - `ListProjectOpenSourceStrategy`
    - `CreateOpenSourceStrategy`
    - `UpdateOpenSourceStrategy`
    - `ShowOpenSourceStrategy`
    - `ListOpenSourceStrategy`
    - `DeleteOpenSourceStrategy`
    - `SwitchOpenSourceStrategy`

### HuaweiCloud SDK CSE

- _新增特性_
  - 支持接口`ShowEngineQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 支持接口`ListFactoryTaskOverview`、`ListFactoryTaskCompletion`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateJob**
    - 请求参数变更
      - `* nodes.cronTrigger.dependJobs: list<DependJobs> -> object<DependJob>`
      - `+ schedule.cron.dependJobs.sameWorkSpaceJobs`
      - `+ schedule.cron.dependJobs.otherWorkSpaceJobs`
      - `+ schedule.cron.dependJobs.dependFailPolicy: enum value [FAIL,IGNORE,SUSPEND]`
      - `* schedule.cron.dependJobs: object<DependJobs> -> object<DependJob>`
  - **ShowJob**
    - 响应参数变更
      - `* nodes.cronTrigger.dependJobs: list<DependJobs> -> object<DependJob>`
      - `+ schedule.cron.dependJobs.sameWorkSpaceJobs`
      - `+ schedule.cron.dependJobs.otherWorkSpaceJobs`
      - `+ schedule.cron.dependJobs.dependFailPolicy: enum value [FAIL,IGNORE,SUSPEND]`
      - `* schedule.cron.dependJobs: object<DependJobs> -> object<DependJob>`
  - **UpdateJob**
    - 请求参数变更
      - `* nodes.cronTrigger.dependJobs: list<DependJobs> -> object<DependJob>`
      - `+ schedule.cron.dependJobs.sameWorkSpaceJobs`
      - `+ schedule.cron.dependJobs.otherWorkSpaceJobs`
      - `+ schedule.cron.dependJobs.dependFailPolicy: enum value [FAIL,IGNORE,SUSPEND]`
      - `* schedule.cron.dependJobs: object<DependJobs> -> object<DependJob>`
  - **CreateSupplementdata**
    - 请求参数变更
      - `* dependJobs.nodes.cronTrigger.dependJobs: list<DependJobs> -> object<DependJob>`
      - `+ dependJobs.schedule.cron.dependJobs.sameWorkSpaceJobs`
      - `+ dependJobs.schedule.cron.dependJobs.otherWorkSpaceJobs`
      - `+ dependJobs.schedule.cron.dependJobs.dependFailPolicy: enum value [FAIL,IGNORE,SUSPEND]`
      - `* dependJobs.schedule.cron.dependJobs: object<DependJobs> -> object<DependJob>`

### HuaweiCloud SDK EG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListEventStreaming**
    - 响应参数变更
      - `+ total`
      - `+ size`
      - `+ items`
      - `- body`
  - **CreateEventStreaming**
    - 请求参数变更
      - `* rule_config.filter: string -> object`
      - `+ source.source_kafka.addr`
  - **ShowEventStreaming**
    - 响应参数变更
      - `* rule_config.filter: string -> object`
      - `+ source.source_kafka.addr`
  - **UpdateEventStreaming**
    - 请求参数变更
      - `* rule_config.filter: string -> object`
      - `+ source.source_kafka.addr`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateAccessCode**
    - 请求参数变更
      - `+ force_disconnect`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`ResetMessageOffsetWithEngine`
- _解决问题_
  - 无
- _特性变更_
  - **ShowGroups**
    - 响应参数变更
      - `* group.group_message_offsets.message_current_offset: int32 -> int64`
      - `* group.group_message_offsets.message_log_end_offset: int32 -> int64`
  - **ResetMessageOffset**
    - 请求参数变更
      - `* message_offset: int32 -> int64`
      - `* timestamp: int32 -> int64`
  - **ShowMessages**
    - 响应参数变更
      - `* messages.message_offset: int32 -> int64`
  - **ShowInstanceMessages**
    - 响应参数变更
      - `* messages.message_offset: int32 -> int64`

### HuaweiCloud SDK MPC

- _新增特性_
  - 支持接口`ShowTenantAccessInfo`、`UpdateTenantAccessInfo`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`ShowMrsFlavors`、`ShowSyncIamUser`、`UpdateSyncIamUser`、`CancelSyncIamUser`
- _解决问题_
  - 无
- _特性变更_
  - **CreateCluster**
    - 请求参数变更
      - `+ smn_notify`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchRestartOrDeleteInstances**
    - 请求参数变更
      - `- action: enum value [restart]`
  - **ListInstancesDetails**
    - 请求参数变更
      - `+ status: enum value [UPGRADING,UPGRADINGFAILED]`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **SetOffSiteBackupPolicy**
    - 请求参数变更
      - `* policy_para: list<OffSiteBackupPolicy> -> object<OffSiteBackupPolicy>`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ResizeInstance**
    - 请求参数变更
      - `+ publicip_id`

### HuaweiCloud SDK SecMaster

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAlerts**
    - 响应参数变更
      - `+ data.data_object.data_source`
      - `+ data.data_object.severity`
      - `+ data.data_object.creator`
      - `- data.data_object.datasource`
      - `- data.data_object.serverity`
      - `- data.data_object.cteator`

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 支持以下接口：
    - `ShowLdapConfig`
    - `UpdateLdapConfig`
    - `CreateLdapConfig`
    - `DeleteLdapConfig`
    - `ShowJobDetail`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SMN

- _新增特性_
  - 支持以下接口：
    - `PublishHttpDetect`
    - `ShowHttpDetectResult`
    - `BatchUpdateSubscriptionsFilterPolices`
    - `BatchCreateSubscriptionsFilterPolices`
    - `BatchDeleteSubscriptionsFilterPolices`
    - `AddSubscriptionFromSubscriptionUser`
- _解决问题_
  - 无
- _特性变更_
  - 废弃以下接口：
    - `ListApplicationAttributes`
    - `UpdateApplication`
    - `DeleteApplication`
    - `ListApplicationEndpointAttributes`
    - `UpdateApplicationEndpoint`
    - `DeleteApplicationEndpoint`
    - `PublishAppMessage`
    - `CreateApplication`
    - `ListApplications`
    - `CreateApplicationEndpoint`
    - `ListApplicationEndpoints`
  - **PublishMessage**
    - 请求参数变更
      - `+ message_attributes`
  - **ListTopics**
    - 请求参数变更
      - `+ fuzzy_display_name`
  - **ListSubscriptions**
    - 请求参数变更
      - `+ fuzzy_remark`
    - 响应参数变更
      - `+ subscriptions.filter_polices`
  - **ListSubscriptionsByTopic**
    - 请求参数变更
      - `+ fuzzy_remark`
    - 响应参数变更
      - `+ subscriptions.filter_polices`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPorts**
    - 响应参数变更
      - `+ ports.ipv6_bandwidth_id`
  - **CreatePort**
    - 响应参数变更
      - `+ port.ipv6_bandwidth_id`
  - **ShowPort**
    - 响应参数变更
      - `+ port.ipv6_bandwidth_id`
  - **UpdatePort**
    - 响应参数变更
      - `+ port.ipv6_bandwidth_id`

# 0.1.75 2023-12-27

### HuaweiCloud SDK AAD

- _新增特性_
  - 支持以下接口：
    - `CreateAadDomain`
    - `CreateCertificate`
    - `ModifyDomainWebSwitch`
    - `ListSourceIps`
    - `AddBlackWhiteIpList`
    - `DeleteBlackWhiteIpList`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListMetricData**
    - 响应参数变更
      - `* datapoints.timestamp: int32 -> int64`
  - **ListApisV2**
    - 请求参数变更
      - `+ return_data_mode: enum value [brief,include_group,include_group_backend]`

### HuaweiCloud SDK CodeArtsArtifact

- _新增特性_
  - 支持以下接口：
    - `BatchRestoreRepo`
    - `BatchDeleteTrashes`
    - `CreateMavenRepo`
    - `ShowProjectList`
    - `ModifyRepository`
    - `ShowRepositoryInfo`
    - `CreateDockerRepositories`
    - `DeleteRepository`
    - `ShowStorage`
    - `ShowMavenInfo`
    - `CreateProjectRelatedRepository`
    - `SearchByChecksum`
    - `SearchArtifacts`
    - `ResetUserPassword`
    - `ShowFileTree`
    - `ListArtifactoryComponent`
    - `ListAllRepositories`
    - `ShowAudit`
    - `ShowRepository`
    - `ListArtifactoryStorageStatistic`
    - `CreateAttention`
    - `ListAttentions`
    - `UpdateArtifactory`
    - `CreateArtifactory`
    - `DeleteArtifactFile`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CSMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListResourceInstances**
    - 响应参数变更
      - `+ resources.sys_tags`

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 支持接口`ListFactoryJobs`、`CreateFactoryJob`、`ListFactoryAlarmInfo`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ShowInstanceSslDetail`、`UpdateSslSwitch`、`DownloadSslCert`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持接口`ShowClientNetwork`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateJob**
    - 请求参数变更
      - `+ singleNodeJobFlag`
      - `+ singleNodeJobType`
  - **ShowJob**
    - 响应参数变更
      - `+ singleNodeJobFlag`
      - `+ singleNodeJobType`
  - **UpdateJob**
    - 请求参数变更
      - `+ singleNodeJobFlag`
      - `+ singleNodeJobType`
  - **CreateSupplementdata**
    - 请求参数变更
      - `+ singleNodeJobFlag`
      - `+ singleNodeJobType`
      - `+ dependJobs.singleNodeJobFlag`
      - `+ dependJobs.singleNodeJobType`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持接口`ListJobs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK HSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAssetStatistic**
    - 响应参数变更
      - `+ environment_num`
      - `+ core_conf_file_num`
  - **ListPorts**
    - 响应参数变更
      - `+ data_list.agent_id`
      - `+ data_list.container_id`
  - **ListSwrImageRepository**
    - 响应参数变更
      - `+ data_list.instance_name`
      - `+ data_list.instance_id`
      - `+ data_list.instance_url`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`CreateShrinkageJob`、`ShowShrinkCheckResult`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MetaStudio

- _新增特性_
  - 支持以下接口：
    - `ListDigitalHumanVideo`
    - `ListInteractionRuleGroups`
    - `CreateInteractionRuleGroup`
    - `UpdateInteractionRuleGroup`
    - `DeleteInteractionRuleGroup`
    - `CheckTextLanguage`
    - `CreateFacialAnimations`
    - `ListFacialAnimationsData`
- _解决问题_
  - 无
- _特性变更_
  - **CreateFile**
    - 响应参数变更
      - `- file_id`
      - `- upload_url`
  - **ExecuteSmartLiveCommand**
    - 请求参数变更
      - `+ review_config`
      - `+ command: enum value [GET_CURRENT_PLAYING_SCRIPTS]`
  - **CreatePictureModelingByUrlJob**
    - 请求参数变更
      - `- X-User-Privilege`
  - **ListAssetSummary**
    - 响应参数变更
      - `+ asset_list.asset_type: enum value [AUDIO]`
  - **Create2DDigitalHumanVideo**
    - 请求参数变更
      - `+ review_config`
      - `+ callback_config`
      - `+ video_config.subtitle_config`
      - `- video_config.disable_system_watermark`
      - `+ video_config.codec: enum value [VP9]`
  - **Show2DDigitalHumanVideo**
    - 响应参数变更
      - `+ video_config.subtitle_config`
      - `- video_config.disable_system_watermark`
      - `+ video_config.codec: enum value [VP9]`
  - **CreatePhotoDigitalHumanVideo**
    - 请求参数变更
      - `+ review_config`
      - `- video_config.disable_system_watermark`
  - **ShowPhotoDigitalHumanVideo**
    - 响应参数变更
      - `- video_config.disable_system_watermark`
  - **LiveEventReport**
    - 请求参数变更
      - `+ review_config`
  - **CreateTtsa**
    - 请求参数变更
      - `- X-User-Privilege`
      - `+ script_type`
      - `+ audio_file_download_url`
      - `+ job_type`
      - `- parent_job_id`
      - `- auto_motion`
  - **ListTtsaJobs**
    - 响应参数变更
      - `+ ttsa_jobs.job_type`
  - **ListTtsaData**
    - 响应参数变更
      - `+ start_time`
      - `+ end_time`
      - `+ is_tail`
  - **ListStyles**
    - 响应参数变更
      - `- styles.extra_meta.edit_value_items`
      - `- styles.extra_meta.edit_color_items`
      - `- styles.extra_meta.edit_components`
      - `- styles.extra_meta.modelling_algorithm`
  - **CreateDigitalHumanBusinessCard**
    - 请求参数变更
      - `+ introduction_type`
      - `+ introduction_audio_asset_id`
      - `+ review_config`
    - 响应参数变更
      - `- job_id`
  - **UpdateDigitalHumanBusinessCard**
    - 请求参数变更
      - `+ introduction_type`
      - `+ introduction_audio_asset_id`
      - `+ review_config`
    - 响应参数变更
      - `- job_id`
  - **ShowDigitalHumanBusinessCard**
    - 响应参数变更
      - `+ introduction_audio_asset_id`
      - `+ introduction_type`
  - **ShowSmartLive**
    - 响应参数变更
      - `+ stream_duration`
      - `+ block_reason`
      - `+ live_event_callback_config`
      - `+ state: enum value [BLOCKED]`
      - `+ rtc_room_info.users.user_type: enum value [INFERENCE_USER,END_USER]`
  - **ListVideoScripts**
    - 请求参数变更
      - `+ name`
      - `+ script_catalog`
      - `+ view_mode`
    - 响应参数变更
      - `+ video_scripts.script_cover_url`
      - `+ video_scripts.script_type`
      - `+ video_scripts.text`
      - `- video_scripts.video_making_type`
      - `- video_scripts.human_image`
  - **ShowVideoScript**
    - 响应参数变更
      - `+ script_cover_url`
      - `+ review_config`
      - `- video_making_type`
      - `- human_image`
      - `+ video_config.subtitle_config`
      - `- video_config.disable_system_watermark`
      - `+ video_config.codec: enum value [VP9]`
      - `* shoot_scripts: list<ShootScriptItem> -> list<ShootScriptShowItem>`
  - **CreatePictureModelingJob**
    - 请求参数变更
      - `- X-User-Privilege`
    - 响应参数变更
      - `- model_asset_id`
      - `- job_id`
  - **ShowVideoMotionCaptureJob**
    - 响应参数变更
      - `+ input_info.rtc_room_info.users.user_type: enum value [INFERENCE_USER,END_USER]`
  - **ShowAsset**
    - 响应参数变更
      - `+ reason`
      - `+ is_need_generate_cover`
      - `+ fail_type`
      - `+ asset_type: enum value [AUDIO]`
      - `+ system_properties.key: enum value [MATERIAL_IMG,MATERIAL_VIDEO,BUSSINESS_CARD_VIDEO,TO_BE_TRANSLATED_VIDEO]`
      - `+ files.state`
      - `+ files.reason`
      - `+ asset_extra_meta.voice_model_meta.speed_ratio`
      - `+ asset_extra_meta.voice_model_meta.volume_ratio`
      - `- asset_extra_meta.voice_model_meta.tts_mode`
      - `- asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.voice_model_meta.language: enum value [GER,fr,Kr,por,JPN,Ita,ESP,DBH,GT,GXH,HBH,SXH,SCH,YY,Russian,Filipino,Dutch,Indonesian,Vietnamese,Arabic,Turkish,Malay,Thai,Finnish]`
      - `+ asset_extra_meta.human_model_2d_meta.model_resolution`
      - `- asset_extra_meta.human_model_2d_meta.is_realtime_matting`
      - `+ asset_extra_meta.ppt_meta.error_info`
  - **UpdateDigitalAsset**
    - 请求参数变更
      - `+ is_need_generate_cover`
      - `+ review_config`
      - `+ asset_type: enum value [AUDIO]`
      - `+ system_properties.key: enum value [MATERIAL_IMG,MATERIAL_VIDEO,BUSSINESS_CARD_VIDEO,TO_BE_TRANSLATED_VIDEO]`
      - `+ asset_extra_meta.voice_model_meta.speed_ratio`
      - `+ asset_extra_meta.voice_model_meta.volume_ratio`
      - `- asset_extra_meta.voice_model_meta.tts_mode`
      - `- asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.voice_model_meta.language: enum value [GER,fr,Kr,por,JPN,Ita,ESP,DBH,GT,GXH,HBH,SXH,SCH,YY,Russian,Filipino,Dutch,Indonesian,Vietnamese,Arabic,Turkish,Malay,Thai,Finnish]`
      - `+ asset_extra_meta.human_model_2d_meta.model_resolution`
      - `- asset_extra_meta.human_model_2d_meta.is_realtime_matting`
      - `+ asset_extra_meta.ppt_meta.error_info`
    - 响应参数变更
      - `+ reason`
      - `+ is_need_generate_cover`
      - `+ fail_type`
      - `+ asset_type: enum value [AUDIO]`
      - `+ system_properties.key: enum value [MATERIAL_IMG,MATERIAL_VIDEO,BUSSINESS_CARD_VIDEO,TO_BE_TRANSLATED_VIDEO]`
      - `+ files.state`
      - `+ files.reason`
      - `+ asset_extra_meta.voice_model_meta.speed_ratio`
      - `+ asset_extra_meta.voice_model_meta.volume_ratio`
      - `- asset_extra_meta.voice_model_meta.tts_mode`
      - `- asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.voice_model_meta.language: enum value [GER,fr,Kr,por,JPN,Ita,ESP,DBH,GT,GXH,HBH,SXH,SCH,YY,Russian,Filipino,Dutch,Indonesian,Vietnamese,Arabic,Turkish,Malay,Thai,Finnish]`
      - `+ asset_extra_meta.human_model_2d_meta.model_resolution`
      - `- asset_extra_meta.human_model_2d_meta.is_realtime_matting`
      - `+ asset_extra_meta.ppt_meta.error_info`
  - **ListSmartLiveRooms**
    - 请求参数变更
      - `+ room_type`
    - 响应参数变更
      - `+ smart_live_rooms.room_type`
      - `+ smart_live_rooms.room_state`
      - `+ smart_live_rooms.error_info`
      - `+ smart_live_rooms.model_infos.backup_model_asset_ids`
  - **CreateSmartLiveRoom**
    - 请求参数变更
      - `+ stream_keys`
      - `+ backup_model_asset_ids`
      - `+ live_event_callback_config`
      - `+ review_config`
      - `+ shared_config`
      - `+ room_type: enum value [TEMPLATE]`
      - `+ play_policy.random_play_mode`
      - `+ video_config.subtitle_config`
      - `- video_config.disable_system_watermark`
      - `+ video_config.codec: enum value [VP9]`
      - `+ scene_scripts.layer_config.group_id`
      - `+ scene_scripts.layer_config.layer_type: enum value [TEXT]`
      - `+ interaction_rules.rule_index`
      - `+ interaction_rules.review_config`
      - `+ interaction_rules.trigger.layer_config`
      - `+ interaction_rules.trigger.reply_audios`
      - `+ interaction_rules.trigger.reply_mode: enum value [CALLBACK,SHOW_LAYER]`
  - **ShowSmartLiveRoom**
    - 响应参数变更
      - `+ backup_model_asset_ids`
      - `+ error_info`
      - `+ stream_keys`
      - `+ shared_config`
      - `+ live_event_callback_config`
      - `+ room_state`
      - `+ review_config`
      - `+ room_type: enum value [TEMPLATE]`
      - `+ play_policy.random_play_mode`
      - `+ video_config.subtitle_config`
      - `- video_config.disable_system_watermark`
      - `+ video_config.codec: enum value [VP9]`
      - `+ scene_scripts.layer_config.group_id`
      - `+ scene_scripts.layer_config.layer_type: enum value [TEXT]`
      - `+ interaction_rules.rule_index`
      - `+ interaction_rules.review_config`
      - `+ interaction_rules.trigger.layer_config`
      - `+ interaction_rules.trigger.reply_audios`
      - `+ interaction_rules.trigger.reply_mode: enum value [CALLBACK,SHOW_LAYER]`
  - **UpdateSmartLiveRoom**
    - 请求参数变更
      - `+ stream_keys`
      - `+ backup_model_asset_ids`
      - `+ live_event_callback_config`
      - `+ review_config`
      - `+ shared_config`
      - `+ room_type: enum value [TEMPLATE]`
      - `+ play_policy.random_play_mode`
      - `+ video_config.subtitle_config`
      - `- video_config.disable_system_watermark`
      - `+ video_config.codec: enum value [VP9]`
      - `+ scene_scripts.layer_config.group_id`
      - `+ scene_scripts.layer_config.layer_type: enum value [TEXT]`
      - `+ interaction_rules.rule_index`
      - `+ interaction_rules.review_config`
      - `+ interaction_rules.trigger.layer_config`
      - `+ interaction_rules.trigger.reply_audios`
      - `+ interaction_rules.trigger.reply_mode: enum value [CALLBACK,SHOW_LAYER]`
    - 响应参数变更
      - `+ backup_model_asset_ids`
      - `+ error_info`
      - `+ stream_keys`
      - `+ shared_config`
      - `+ live_event_callback_config`
      - `+ room_state`
      - `+ review_config`
      - `+ room_type: enum value [TEMPLATE]`
      - `+ play_policy.random_play_mode`
      - `+ video_config.subtitle_config`
      - `- video_config.disable_system_watermark`
      - `+ video_config.codec: enum value [VP9]`
      - `+ scene_scripts.layer_config.group_id`
      - `+ scene_scripts.layer_config.layer_type: enum value [TEXT]`
      - `+ interaction_rules.rule_index`
      - `+ interaction_rules.review_config`
      - `+ interaction_rules.trigger.layer_config`
      - `+ interaction_rules.trigger.reply_audios`
      - `+ interaction_rules.trigger.reply_mode: enum value [CALLBACK,SHOW_LAYER]`
  - **StartSmartLive**
    - 请求参数变更
      - `+ stream_keys`
      - `+ interaction_callback_url`
      - `+ live_event_callback_config`
      - `+ video_config.subtitle_config`
      - `- video_config.disable_system_watermark`
      - `+ video_config.codec: enum value [VP9]`
      - `+ play_policy.random_play_mode`
    - 响应参数变更
      - `+ live_warning_info`
      - `+ live_event_callback_config`
      - `+ rtc_room_info.users.user_type: enum value [INFERENCE_USER,END_USER]`
  - **ListSmartLive**
    - 响应参数变更
      - `+ stream_duration`
      - `+ block_reason`
      - `+ live_event_callback_config`
      - `+ smart_live_jobs.live_event_callback_config`
      - `+ smart_live_jobs.stream_duration`
      - `+ smart_live_jobs.block_reason`
      - `+ smart_live_jobs.state: enum value [BLOCKED]`
      - `+ smart_live_jobs.rtc_room_info.users.user_type: enum value [INFERENCE_USER,END_USER]`
  - **CreateVideoMotionCaptureJob**
    - 请求参数变更
      - `+ input_info.rtc_room_info.users.user_type: enum value [INFERENCE_USER,END_USER]`
    - 响应参数变更
      - `- rtc_room_info`
      - `- job_id`
      - `+ rtc_room_info.users.user_type: enum value [INFERENCE_USER,END_USER]`
  - **ListVideoMotionCaptureJobs**
    - 响应参数变更
      - `+ video_motion_capture_jobs.input_info.rtc_room_info.users.user_type: enum value [INFERENCE_USER,END_USER]`
  - **CreateDigitalAsset**
    - 请求参数变更
      - `+ is_need_generate_cover`
      - `+ review_config`
      - `+ asset_type: enum value [AUDIO]`
      - `+ system_properties.key: enum value [MATERIAL_IMG,MATERIAL_VIDEO,BUSSINESS_CARD_VIDEO,TO_BE_TRANSLATED_VIDEO]`
      - `+ asset_extra_meta.voice_model_meta.speed_ratio`
      - `+ asset_extra_meta.voice_model_meta.volume_ratio`
      - `- asset_extra_meta.voice_model_meta.tts_mode`
      - `- asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.voice_model_meta.language: enum value [GER,fr,Kr,por,JPN,Ita,ESP,DBH,GT,GXH,HBH,SXH,SCH,YY,Russian,Filipino,Dutch,Indonesian,Vietnamese,Arabic,Turkish,Malay,Thai,Finnish]`
      - `+ asset_extra_meta.human_model_2d_meta.model_resolution`
      - `- asset_extra_meta.human_model_2d_meta.is_realtime_matting`
      - `+ asset_extra_meta.ppt_meta.error_info`
  - **ListAssets**
    - 请求参数变更
      - `- asset_manage_type`
      - `- X-User-MePrivilege`
    - 响应参数变更
      - `+ reason`
      - `+ is_need_generate_cover`
      - `+ fail_type`
      - `+ assets.fail_type`
      - `+ assets.reason`
      - `+ assets.is_need_generate_cover`
      - `+ assets.asset_type: enum value [AUDIO]`
      - `+ assets.system_properties.key: enum value [MATERIAL_IMG,MATERIAL_VIDEO,BUSSINESS_CARD_VIDEO,TO_BE_TRANSLATED_VIDEO]`
      - `+ assets.files.state`
      - `+ assets.files.reason`
      - `+ assets.asset_extra_meta.voice_model_meta.speed_ratio`
      - `+ assets.asset_extra_meta.voice_model_meta.volume_ratio`
      - `- assets.asset_extra_meta.voice_model_meta.tts_mode`
      - `- assets.asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ assets.asset_extra_meta.voice_model_meta.language: enum value [GER,fr,Kr,por,JPN,Ita,ESP,DBH,GT,GXH,HBH,SXH,SCH,YY,Russian,Filipino,Dutch,Indonesian,Vietnamese,Arabic,Turkish,Malay,Thai,Finnish]`
      - `+ assets.asset_extra_meta.human_model_2d_meta.model_resolution`
      - `- assets.asset_extra_meta.human_model_2d_meta.is_realtime_matting`
      - `+ assets.asset_extra_meta.ppt_meta.error_info`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`SetInstancesNewDbShrink`、`StopBackup`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.74 2023-12-22

### HuaweiCloud SDK Config

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowTrackerConfig**
    - 响应参数变更
      - `+ channel.obs.bucket_prefix`
  - **CreateTrackerConfig**
    - 请求参数变更
      - `+ channel.obs.bucket_prefix`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowGroups**
    - 响应参数变更
      - `* group.group_message_offsets.lag: int32 -> int64`
  - **ShowInstanceTopicDetail**
    - 响应参数变更
      - `* partitions.replicas.lag: int32 -> int64`

# 0.1.73 2023-12-21

### HuaweiCloud SDK AS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateScalingConfig**
    - 请求参数变更
      - `+ instance_config.disk.iops`
      - `+ instance_config.disk.throughput`
      - `+ instance_config.disk.volume_type: enum value [GPSSD2,ESSD2]`
  - **ListScalingConfigs**
    - 响应参数变更
      - `+ scaling_configurations.instance_config.disk.iops`
      - `+ scaling_configurations.instance_config.disk.throughput`
      - `+ scaling_configurations.instance_config.disk.volume_type: enum value [GPSSD2,ESSD2]`
  - **ShowScalingConfig**
    - 响应参数变更
      - `+ scaling_configuration.instance_config.disk.iops`
      - `+ scaling_configuration.instance_config.disk.throughput`
      - `+ scaling_configuration.instance_config.disk.volume_type: enum value [GPSSD2,ESSD2]`

### HuaweiCloud SDK CloudPond

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListEdgeSites**
    - 响应参数变更
      - `+ edge_sites.location.zone_code`
      - `+ edge_sites.location.address`
  - **CreateEdgeSite**
    - 请求参数变更
      - `+ edge_site.location.address`
      - `+ edge_site.location.zone_code`
    - 响应参数变更
      - `+ edge_site.location.zone_code`
      - `+ edge_site.location.address`
  - **ShowEdgeSite**
    - 响应参数变更
      - `+ edge_site.location.zone_code`
      - `+ edge_site.location.address`
  - **UpdateEdgeSite**
    - 响应参数变更
      - `+ edge_site.location.zone_code`
      - `+ edge_site.location.address`

### HuaweiCloud SDK CloudTable

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateCloudTableCluster**
    - 请求参数变更
      - `+ cluster_name`
      - `+ datastore`
      - `+ availability_zone`
      - `+ nics`
      - `+ cluster_info`
      - `+ enterprise_project_id`
      - `+ vpc_id`
      - `+ dbuser`
      - `+ dbuserpwd`
      - `- cluster`
      - `* body: object<CreateClusterRequestBody> -> object<CreateClusterReqBody>`
    - 响应参数变更
      - `+ jobId`
      - `+ getJobEndpoint`
  - **CreateCluster**
    - 请求参数变更
      - `* cluster.instance.nics: list<Nics> -> list<nic>`

### HuaweiCloud SDK CodeArtsPipeline

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowPipelineRunDetail**
    - 响应参数变更
      - `* current_system_time: string -> int64`
      - `* stages.pre.endpoint_ids: string -> list<string>`

### HuaweiCloud SDK DC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateVifPeer**
    - 响应参数变更
      - `+ vif_peer.enable_nqa`
      - `+ vif_peer.enable_bfd`
  - **CreateVifPeer**
    - 响应参数变更
      - `+ vif_peer.enable_nqa`
      - `+ vif_peer.enable_bfd`
  - **ShowDirectConnect**
    - 响应参数变更
      - `+ direct_connect.signed_agreement_time`
      - `+ direct_connect.locales`
      - `+ direct_connect.support_feature`
      - `+ direct_connect.ies_id`
      - `+ direct_connect.reason`
      - `+ direct_connect.email`
      - `+ direct_connect.onestop_product_id`
      - `+ direct_connect.building_line_product_id`
      - `+ direct_connect.last_onestop_product_id`
      - `+ direct_connect.last_building_line_product_id`
      - `+ direct_connect.modified_bandwidth`
      - `+ direct_connect.change_mode`
      - `+ direct_connect.onestopdc_status`
      - `+ direct_connect.public_border_group`
      - `+ direct_connect.auto_renew`
      - `+ direct_connect.ratio_95peak`
      - `+ direct_connect.type: enum value [onestop_standard,onestop_hosted]`
      - `- direct_connect.status: enum value [ORDERING,ACCEPT,REJECTED]`
  - **UpdateDirectConnect**
    - 响应参数变更
      - `+ direct_connect.signed_agreement_time`
      - `+ direct_connect.locales`
      - `+ direct_connect.support_feature`
      - `+ direct_connect.ies_id`
      - `+ direct_connect.reason`
      - `+ direct_connect.email`
      - `+ direct_connect.onestop_product_id`
      - `+ direct_connect.building_line_product_id`
      - `+ direct_connect.last_onestop_product_id`
      - `+ direct_connect.last_building_line_product_id`
      - `+ direct_connect.modified_bandwidth`
      - `+ direct_connect.change_mode`
      - `+ direct_connect.onestopdc_status`
      - `+ direct_connect.public_border_group`
      - `+ direct_connect.auto_renew`
      - `+ direct_connect.ratio_95peak`
      - `+ direct_connect.type: enum value [onestop_standard,onestop_hosted]`
      - `- direct_connect.status: enum value [ORDERING,ACCEPT,REJECTED]`
  - **ListDirectConnects**
    - 响应参数变更
      - `+ direct_connects.signed_agreement_time`
      - `+ direct_connects.locales`
      - `+ direct_connects.support_feature`
      - `+ direct_connects.ies_id`
      - `+ direct_connects.reason`
      - `+ direct_connects.email`
      - `+ direct_connects.onestop_product_id`
      - `+ direct_connects.building_line_product_id`
      - `+ direct_connects.last_onestop_product_id`
      - `+ direct_connects.last_building_line_product_id`
      - `+ direct_connects.modified_bandwidth`
      - `+ direct_connects.change_mode`
      - `+ direct_connects.onestopdc_status`
      - `+ direct_connects.public_border_group`
      - `+ direct_connects.auto_renew`
      - `+ direct_connects.ratio_95peak`
      - `+ direct_connects.type: enum value [onestop_standard,onestop_hosted]`
      - `- direct_connects.status: enum value [ORDERING,ACCEPT,REJECTED]`
  - **ListHostedDirectConnects**
    - 响应参数变更
      - `+ hosted_connects.port_type`
      - `+ hosted_connects.type`
      - `+ hosted_connects.status: enum value [PENDING_UPDATE,PENDING_CREATE]`
      - `- hosted_connects.status: enum value [PAID,APPLY,PENDING_SURVEY,DELETED,DENY,PENDING_PAY,ORDERING,ACCEPT,REJECTED]`
  - **CreateHostedDirectConnect**
    - 响应参数变更
      - `+ hosted_connect.port_type`
      - `+ hosted_connect.type`
      - `+ hosted_connect.status: enum value [PENDING_UPDATE,PENDING_CREATE]`
      - `- hosted_connect.status: enum value [PAID,APPLY,PENDING_SURVEY,DELETED,DENY,PENDING_PAY,ORDERING,ACCEPT,REJECTED]`
  - **ShowHostedDirectConnect**
    - 响应参数变更
      - `+ hosted_connect.port_type`
      - `+ hosted_connect.type`
      - `+ hosted_connect.status: enum value [PENDING_UPDATE,PENDING_CREATE]`
      - `- hosted_connect.status: enum value [PAID,APPLY,PENDING_SURVEY,DELETED,DENY,PENDING_PAY,ORDERING,ACCEPT,REJECTED]`
  - **UpdateHostedDirectConnect**
    - 响应参数变更
      - `+ hosted_connect.port_type`
      - `+ hosted_connect.type`
      - `+ hosted_connect.status: enum value [PENDING_UPDATE,PENDING_CREATE]`
      - `- hosted_connect.status: enum value [PAID,APPLY,PENDING_SURVEY,DELETED,DENY,PENDING_PAY,ORDERING,ACCEPT,REJECTED]`
  - **ShowVirtualGateway**
    - 响应参数变更
      - `+ virtual_gateway.device_id`
      - `+ virtual_gateway.redundant_device_id`
      - `+ virtual_gateway.public_border_group`
  - **UpdateVirtualGateway**
    - 响应参数变更
      - `+ virtual_gateway.device_id`
      - `+ virtual_gateway.redundant_device_id`
      - `+ virtual_gateway.public_border_group`
  - **ListVirtualGateways**
    - 响应参数变更
      - `+ virtual_gateways.device_id`
      - `+ virtual_gateways.redundant_device_id`
      - `+ virtual_gateways.public_border_group`
  - **CreateVirtualGateway**
    - 响应参数变更
      - `+ virtual_gateway.device_id`
      - `+ virtual_gateway.redundant_device_id`
      - `+ virtual_gateway.public_border_group`
  - **ShowVirtualInterface**
    - 响应参数变更
      - `+ virtual_interface.local_gateway_v4_ip`
      - `+ virtual_interface.remote_gateway_v4_ip`
      - `+ virtual_interface.ies_id`
      - `+ virtual_interface.reason`
      - `+ virtual_interface.rate_limit`
      - `+ virtual_interface.address_family`
      - `+ virtual_interface.local_gateway_v6_ip`
      - `+ virtual_interface.remote_gateway_v6_ip`
      - `+ virtual_interface.lgw_id`
      - `+ virtual_interface.gateway_id`
      - `+ virtual_interface.remote_ep_group`
      - `+ virtual_interface.service_ep_group`
      - `+ virtual_interface.bgp_route_limit`
      - `+ virtual_interface.priority`
      - `+ virtual_interface.vif_peers.enable_nqa`
      - `+ virtual_interface.vif_peers.enable_bfd`
  - **UpdateVirtualInterface**
    - 响应参数变更
      - `+ virtual_interface.local_gateway_v4_ip`
      - `+ virtual_interface.remote_gateway_v4_ip`
      - `+ virtual_interface.ies_id`
      - `+ virtual_interface.reason`
      - `+ virtual_interface.rate_limit`
      - `+ virtual_interface.address_family`
      - `+ virtual_interface.local_gateway_v6_ip`
      - `+ virtual_interface.remote_gateway_v6_ip`
      - `+ virtual_interface.lgw_id`
      - `+ virtual_interface.gateway_id`
      - `+ virtual_interface.remote_ep_group`
      - `+ virtual_interface.service_ep_group`
      - `+ virtual_interface.bgp_route_limit`
      - `+ virtual_interface.priority`
      - `+ virtual_interface.vif_peers.enable_nqa`
      - `+ virtual_interface.vif_peers.enable_bfd`
  - **ListVirtualInterfaces**
    - 响应参数变更
      - `+ virtual_interfaces.local_gateway_v4_ip`
      - `+ virtual_interfaces.remote_gateway_v4_ip`
      - `+ virtual_interfaces.ies_id`
      - `+ virtual_interfaces.reason`
      - `+ virtual_interfaces.rate_limit`
      - `+ virtual_interfaces.address_family`
      - `+ virtual_interfaces.local_gateway_v6_ip`
      - `+ virtual_interfaces.remote_gateway_v6_ip`
      - `+ virtual_interfaces.lgw_id`
      - `+ virtual_interfaces.gateway_id`
      - `+ virtual_interfaces.remote_ep_group`
      - `+ virtual_interfaces.service_ep_group`
      - `+ virtual_interfaces.bgp_route_limit`
      - `+ virtual_interfaces.priority`
      - `+ virtual_interfaces.vif_peers.enable_nqa`
      - `+ virtual_interfaces.vif_peers.enable_bfd`
  - **CreateVirtualInterface**
    - 响应参数变更
      - `+ virtual_interface.local_gateway_v4_ip`
      - `+ virtual_interface.remote_gateway_v4_ip`
      - `+ virtual_interface.ies_id`
      - `+ virtual_interface.reason`
      - `+ virtual_interface.rate_limit`
      - `+ virtual_interface.address_family`
      - `+ virtual_interface.local_gateway_v6_ip`
      - `+ virtual_interface.remote_gateway_v6_ip`
      - `+ virtual_interface.lgw_id`
      - `+ virtual_interface.gateway_id`
      - `+ virtual_interface.remote_ep_group`
      - `+ virtual_interface.service_ep_group`
      - `+ virtual_interface.bgp_route_limit`
      - `+ virtual_interface.priority`
      - `+ virtual_interface.vif_peers.enable_nqa`
      - `+ virtual_interface.vif_peers.enable_bfd`

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ValidateDeletableReplica`
- _解决问题_
  - 无
- _特性变更_
  - **ShowExpireKeyScanInfo**
    - 请求参数变更
      - `+ offset`
      - `+ limit`
      - `+ status`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateJob**
    - 请求参数变更
      - `+ nodes.execTimeOutRetry`
  - **ShowJob**
    - 响应参数变更
      - `+ nodes.execTimeOutRetry`
  - **UpdateJob**
    - 请求参数变更
      - `+ nodes.execTimeOutRetry`
  - **CreateSupplementdata**
    - 请求参数变更
      - `+ dependJobs.nodes.execTimeOutRetry`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchCreateJobs**
    - 请求参数变更
      - `+ jobs.engine_type: enum value [mysql-to-kafka,taurus-to-kafka,gaussdbv5ha-to-kafka]`
      - `+ jobs.source_endpoint.db_type: enum value [taurus]`
  - **BatchValidateConnections**
    - 请求参数变更
      - `+ jobs.db_type: enum value [taurus]`
  - **ShowJobList**
    - 请求参数变更
      - `+ engine_type: enum value [gaussdbv5,postgresql,mysql-to-kafka,taurus-to-kafka,gaussdbv5ha-to-kafka]`
  - **BatchUpdateJob**
    - 请求参数变更
      - `+ jobs.engine_type: enum value [gaussdbv5,postgresql,mysql-to-kafka,taurus-to-kafka,gaussdbv5ha-to-kafka]`
      - `+ jobs.source_endpoint.db_type: enum value [taurus]`
  - **BatchListJobDetails**
    - 响应参数变更
      - `+ results.source_endpoint.db_type: enum value [taurus]`
      - `+ results.inst_info.engine_type: enum value [gaussdbv5,postgresql,mysql-to-kafka,taurus-to-kafka,gaussdbv5ha-to-kafka]`
  - **ShowJobDetail**
    - 请求参数变更
      - `+ type: enum value [compare]`
      - `+ type: enum value [comapre]`
      - `+ query_type: enum value [diff]`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchCreateServerTags**
    - 请求参数变更
      - `* tags: list<ServerTag> -> list<BatchAddServerTag>`
  - **UpdateServer**
    - 请求参数变更
      - `+ server.user_data`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`UpdateDisassociatePublicip`、`UpdateAssociatePublicip`

### HuaweiCloud SDK HSS

- _新增特性_
  - 支持以下接口：
    - `ListProcessesHost`
    - `ListPortHost`
    - `ChangeCheckRuleAction`
    - `ListVulScanTask`
    - `CreateVulnerabilityScanTask`
    - `ListVulScanTaskHost`
- _解决问题_
  - 无
- _特性变更_
  - **ShowAssetStatistic**
    - 请求参数变更
      - `+ category`
    - 响应参数变更
      - `+ web_app_num`
      - `+ database_num`
      - `+ web_service_num`
  - **ChangeVulScanPolicy**
    - 请求参数变更
      - `+ scan_vul_types`
  - **ListJarPackageHostInfo**
    - 请求参数变更
      - `+ part_match`
  - **ListUserStatistics**
    - 请求参数变更
      - `+ category`
  - **ListPortStatistics**
    - 请求参数变更
      - `+ category`
  - **ListProcessStatistics**
    - 请求参数变更
      - `+ category`
  - **ListAppStatistics**
    - 请求参数变更
      - `+ category`
  - **ListUsers**
    - 请求参数变更
      - `+ category`
      - `+ part_match`
    - 响应参数变更
      - `+ data_list.container_id`
      - `+ data_list.container_name`
  - **ListPorts**
    - 请求参数变更
      - `+ category`
  - **ListApps**
    - 请求参数变更
      - `+ category`
      - `+ part_match`
    - 响应参数变更
      - `+ data_list.container_id`
      - `+ data_list.container_name`
  - **ListAutoLaunchs**
    - 请求参数变更
      - `+ part_match`
  - **ListProtectionServer**
    - 响应参数变更
      - `+ data_list.agent_version`
  - **ListContainerNodes**
    - 请求参数变更
      - `+ container_tags`
    - 响应参数变更
      - `+ data_list.protect_interrupt`
      - `+ data_list.container_tags`
      - `+ data_list.private_ip`
      - `+ data_list.public_ip`
      - `+ data_list.resource_id`
      - `+ data_list.group_name`
      - `+ data_list.enterprise_project_name`
      - `+ data_list.detect_result`
      - `+ data_list.asset`
      - `+ data_list.vulnerability`
      - `+ data_list.intrusion`
      - `+ data_list.policy_group_id`
      - `+ data_list.policy_group_name`
  - **ListHostStatus**
    - 响应参数变更
      - `+ data_list.expire_time`
      - `+ data_list.protect_interrupt`
  - **BatchScanSwrImage**
    - 请求参数变更
      - `+ image_size`
      - `+ start_latest_update_time`
      - `+ end_latest_update_time`
      - `+ start_latest_scan_time`
      - `+ end_latest_scan_time`
      - `+ image_info_list.instance_id`
      - `+ image_info_list.instance_url`
  - **ListImageVulnerabilities**
    - 请求参数变更
      - `+ type`
    - 响应参数变更
      - `+ data_list.app_path`
  - **ListImageRiskConfigs**
    - 请求参数变更
      - `+ instance_id`
  - **ListImageRiskConfigRules**
    - 请求参数变更
      - `+ instance_id`
  - **ShowImageCheckRuleDetail**
    - 请求参数变更
      - `+ instance_id`
  - **ListAlarmWhiteList**
    - 响应参数变更
      - `+ data_list.white_field`
      - `+ data_list.field_value`
      - `+ data_list.judge_type`
  - **ListSwrImageRepository**
    - 请求参数变更
      - `+ instance_name`
      - `+ image_size`
      - `+ start_latest_update_time`
      - `+ end_latest_update_time`
      - `+ start_latest_scan_time`
      - `+ end_latest_scan_time`
      - `+ has_malicious_file`
      - `+ has_unsafe_setting`
      - `+ has_vul`
      - `+ instance_id`
    - 响应参数变更
      - `+ data_list.scan_failed_desc`
  - **ListSecurityEvents**
    - 请求参数变更
      - `+ public_ip`
      - `+ event_name`
    - 响应参数变更
      - `+ data_list.event_count`
  - **ChangeEvent**
    - 请求参数变更
      - `+ event_white_rule_list`

### HuaweiCloud SDK IEC

- _新增特性_
  - 支持接口`ListBandwidthTypes`、`CreateSubnet`、`CreateInstance`
- _解决问题_
  - 无
- _特性变更_
  - **ListSubnets**
    - 响应参数变更
      - `+ subnets.cidr_v6`
      - `+ subnets.ipv6_enable`
      - `+ subnets.pool_id`
      - `+ subnets.neutron_subnet_id_v6`
      - `+ subnets.gateway_ip_v6`
  - **ShowSubnet**
    - 响应参数变更
      - `+ subnet.cidr_v6`
      - `+ subnet.ipv6_enable`
      - `+ subnet.pool_id`
      - `+ subnet.neutron_subnet_id_v6`
      - `+ subnet.gateway_ip_v6`
  - **UpdateSubnet**
    - 请求参数变更
      - `+ subnet.ipv6_enable`
      - `+ subnet.pool_id`
    - 响应参数变更
      - `+ subnet.ipv6_enable`
      - `+ subnet.neutron_subnet_id_v6`
  - **CreateSecurityGroupRule**
    - 请求参数变更
      - `+ security_group_rule.ethertype: enum value [IPv6]`
  - **ListPorts**
    - 响应参数变更
      - `+ ports.ipv6_bandwidth_id`
      - `+ ports.binding:profile`
  - **CreatePort**
    - 响应参数变更
      - `+ port.ipv6_bandwidth_id`
      - `+ port.binding:profile`
  - **ShowPort**
    - 响应参数变更
      - `+ port.ipv6_bandwidth_id`
      - `+ port.binding:profile`
  - **UpdatePort**
    - 响应参数变更
      - `+ port.ipv6_bandwidth_id`
      - `+ port.binding:profile`
  - **CreatePublicIp**
    - 请求参数变更
      - `+ bandwidth`
  - **ShowEdgeCloud**
    - 响应参数变更
      - `+ stacks.resources.net_config.allowed_address_pairs`
      - `+ coverage.coverage_sites.demands.pool_id_v6`
      - `+ coverage.coverage_sites.demands.ipv6_bandwidth_enable`
  - **ListEdgeCloud**
    - 响应参数变更
      - `+ edgeclouds.coverage.coverage_sites.demands.pool_id_v6`
      - `+ edgeclouds.coverage.coverage_sites.demands.ipv6_bandwidth_enable`
  - **CreateDeployment**
    - 请求参数变更
      - `+ edgecloud.coverage.coverage_sites.demands.bandwidth_type`
      - `+ edgecloud.coverage.coverage_sites.demands.pool_id_v6`
      - `+ edgecloud.coverage.coverage_sites.demands.ipv6_bandwidth_enable`
      - `+ edgecloud.coverage.coverage_sites.demands.ipv6_bandwidth_type`
      - `+ edgecloud.stack.resources.net_config.allowed_address_pairs`
    - 响应参数变更
      - `+ locations.ipv6_enable`
      - `+ locations.ipv6_bandwidth_enable`
      - `+ locations.pool_id_v6`
  - **ListDeployments**
    - 响应参数变更
      - `+ deployments.distribution.ipv6_enable`
      - `+ deployments.distribution.ipv6_bandwidth_enable`
      - `+ deployments.distribution.pool_id_v6`
      - `+ deployments.edgecloud.stacks.resources.net_config.allowed_address_pairs`
      - `+ deployments.edgecloud.coverage.coverage_sites.demands.bandwidth_type`
      - `+ deployments.edgecloud.coverage.coverage_sites.demands.pool_id_v6`
      - `+ deployments.edgecloud.coverage.coverage_sites.demands.ipv6_bandwidth_enable`
      - `+ deployments.edgecloud.coverage.coverage_sites.demands.ipv6_bandwidth_type`

### HuaweiCloud SDK IVS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **DetectStandardByNameAndId**
    - 请求参数变更
      - `+ data.req_data.detail`
      - `+ data.req_data.crop`
  - **DetectStandardByIdCardImage**
    - 请求参数变更
      - `+ data.req_data.detail`
      - `+ data.req_data.crop`
  - **DetectStandardByVideoAndIdCardImage**
    - 请求参数变更
      - `+ data.req_data.detail`
  - **DetectStandardByVideoAndNameAndId**
    - 请求参数变更
      - `+ data.req_data.detail`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`SendKafkaMessage`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CreatePartition`
  - **UpdateInstanceTopic**
    - 请求参数变更
      - `+ topics.new_partition_brokers`
  - **ListInstanceConsumerGroups**
    - 响应参数变更
      - `* groups.lag: int32 -> int64`
  - **ListInstances**
    - 请求参数变更
      - `+ status: enum value [UPGRADING,UPGRADINGFAILED]`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateTranscodingTask**
    - 请求参数变更
      - `+ video_process.hls_storage_type`

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListFsTasks**
    - 响应参数变更
      - `* tasks: list<object> -> list<OneFsTaskResp>`
  - **ShowShare**
    - 响应参数变更
      - `+ tags`
      - `+ enterprise_project_id`
  - **DeleteBackendTarget**
    - 响应参数变更
      - `+ lifecycle: enum value [FAILED]`
  - **CreateShare**
    - 请求参数变更
      - `+ share.tags`
  - **ListShares**
    - 响应参数变更
      - `+ tags`
      - `+ enterprise_project_id`
      - `+ shares.enterprise_project_id`
      - `+ shares.tags`

### HuaweiCloud SDK VOD

- _新增特性_
  - 支持接口`ListAssetDailySummaryLog`、`UpdateStorageMode`、`ShowVodRetrieval`
- _解决问题_
  - 无
- _特性变更_
  - **ShowTakeOverAssetDetails**
    - 响应参数变更
      - `+ transcode_info.output.group_id`
      - `+ transcode_info.output.group_name`
  - **PublishAssets**
    - 响应参数变更
      - `+ asset_info_array.is_multi_transcode`
      - `+ asset_info_array.play_info_array.group_id`
      - `+ asset_info_array.play_info_array.group_name`
  - **UnpublishAssets**
    - 响应参数变更
      - `+ asset_info_array.is_multi_transcode`
      - `+ asset_info_array.play_info_array.group_id`
      - `+ asset_info_array.play_info_array.group_name`
  - **ShowAssetMeta**
    - 响应参数变更
      - `+ asset_info_array.is_multi_transcode`
      - `+ asset_info_array.play_info_array.group_id`
      - `+ asset_info_array.play_info_array.group_name`
  - **ShowAssetDetail**
    - 响应参数变更
      - `+ transcode_info.output.group_id`
      - `+ transcode_info.output.group_name`
  - **ShowTakeOverTaskDetails**
    - 响应参数变更
      - `+ assets.transcode_info.output.group_id`
      - `+ assets.transcode_info.output.group_name`

# 0.1.72 2023-12-14

### HuaweiCloud SDK BMS

- _新增特性_
  - 支持接口`DeleteBaremetalServer`
- _解决问题_
  - 无
- _特性变更_
  - **CreateBareMetalServers**
    - 请求参数变更
      - `+ server.root_volume.volumetype: enum value [GPSSD]`
      - `+ server.data_volumes.volumetype: enum value [GPSSD]`

### HuaweiCloud SDK CAE

- _新增特性_
  - 支持接口`ShowMonitorSystem`、`CreateMonitorSystem`、`UpdateMonitorSystem`
- _解决问题_
  - 无
- _特性变更_
  - **ListComponentConfigurations**
    - 响应参数变更
      - `+ items.data.spec.instrumentation`
      - `+ items.data.spec.apm_application`
      - `+ items.data.spec.type`
      - `+ items.data.spec.app_name`
      - `+ items.data.spec.instance_name`
      - `+ items.data.spec.env_name`
      - `- items.data.spec.image_pull_policy: enum value [Always,IfNotPresent]`
  - **CreateComponentConfiguration**
    - 请求参数变更
      - `+ items.data.spec.instrumentation`
      - `+ items.data.spec.apm_application`
      - `+ items.data.spec.type`
      - `+ items.data.spec.app_name`
      - `+ items.data.spec.instance_name`
      - `+ items.data.spec.env_name`
      - `- items.data.spec.image_pull_policy: enum value [Always,IfNotPresent]`

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAlarmTemplates**
    - 请求参数变更
      - `+ template_type: enum value [system_event,custom_event,system_custom_event]`
  - **CreateAlarmTemplate**
    - 请求参数变更
      - `+ template_type`
      - `+ policies.period: enum value [0]`
  - **UpdateAlarmTemplate**
    - 请求参数变更
      - `+ policies.period: enum value [0]`

### HuaweiCloud SDK CFW

- _新增特性_
  - 支持接口`CreateFirewall`、`ListJob`、`DeleteFirewall`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudTable

- _新增特性_
  - 支持接口`CreateCloudTableCluster`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeArtsPipeline

- _新增特性_
  - 支持接口`ShowPipelineDetail`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Config

- _新增特性_
  - 支持以下接口：
    - `ListTrackedResources`
    - `CountTrackedResources`
    - `ListTrackedResourceTags`
    - `CollectTrackedResourcesSummary`
    - `ShowTrackedResourceDetail`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DLI

- _新增特性_
  - 支持接口`DeleteRouteFromEnhancedConnection`、`CreateRouteToEnhancedConnection`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持以下接口：
    - `ListLogicalClusterVolumes`
    - `ResizeClusterWithExistedNodes`
    - `RestartLogicalCluster`
    - `ListTopoRings`
    - `UpdateLogicalCluster`
    - `DeleteLogicalCluster`
    - `EnableLogicalCluster`
    - `ListClusterNodes`
    - `ConvertToLogicalCluster`
    - `RestoreRedistribution`
    - `StopRedistribution`
    - `ListLogicalClusterTasks`
    - `ListLogicalClusters`
    - `CreateLogicalCluster`
    - `DeleteClusterNodes`
    - `ListLogicalClusterRings`
    - `ListLtsLogs`
    - `ListQueries`
    - `ListTablesStatistic`
    - `ShowQueryDetail`
    - `DisableLtsLogs`
    - `EnableLtsLogs`
- _解决问题_
  - 无
- _特性变更_
  - **ListHostDisk**
    - 请求参数变更
      - `+ instance_id`
    - 响应参数变更
      - `+ instance_id`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowResInstanceInfo**
    - 响应参数变更
      - `+ resources.resource_detail.detailId`
      - `- resources.resource_detail.resource_id`
      - `- resources.resource_detail.func_urn`
      - `- resources.resource_detail.func_name`
      - `- resources.resource_detail.domain_id`
      - `- resources.resource_detail.namespace`
      - `- resources.resource_detail.project_name`
      - `- resources.resource_detail.package`
      - `- resources.resource_detail.runtime`
      - `- resources.resource_detail.timeout`
      - `- resources.resource_detail.handler`
      - `- resources.resource_detail.memory_size`
      - `- resources.resource_detail.gpu_memory`
      - `- resources.resource_detail.cpu`
      - `- resources.resource_detail.code_type`
      - `- resources.resource_detail.code_url`
      - `- resources.resource_detail.code_filename`
      - `- resources.resource_detail.code_size`
      - `- resources.resource_detail.user_data`
      - `- resources.resource_detail.encrypted_user_data`
      - `- resources.resource_detail.digest`
      - `- resources.resource_detail.version`
      - `- resources.resource_detail.image_name`
      - `- resources.resource_detail.xrole`
      - `- resources.resource_detail.app_xrole`
      - `- resources.resource_detail.description`
      - `- resources.resource_detail.last_modified`
      - `- resources.resource_detail.func_vpc_id`
      - `- resources.resource_detail.strategy_config`
      - `- resources.resource_detail.extend_config`
      - `- resources.resource_detail.initializer_handler`
      - `- resources.resource_detail.initializer_timeout`
      - `- resources.resource_detail.pre_stop_handler`
      - `- resources.resource_detail.pre_stop_timeout`
      - `- resources.resource_detail.enterprise_project_id`
      - `- resources.resource_detail.long_time`
      - `- resources.resource_detail.log_group_id`
      - `- resources.resource_detail.log_stream_id`
      - `- resources.resource_detail.type`
      - `- resources.resource_detail.fail_count`
      - `- resources.resource_detail.is_bridge_function`
      - `- resources.resource_detail.bind_bridge_funcUrns`
      - `* resources.resource_detail: object<ListFunctionResult> -> object<ListEnterpriseResourceDetail>`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateAsyncCommand**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`
  - **ShowAsyncDeviceCommand**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`
  - **BroadcastMessage**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`
  - **CreateCommand**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`
  - **ListProperties**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`
    - 响应参数变更
      - `+ request_id`
  - **UpdateProperties**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`
    - 响应参数变更
      - `+ request_id`
  - **CloseDeviceTunnel**
    - 请求参数变更
      - `+ Sp-Auth-Token`
  - **DeleteDeviceTunnel**
    - 请求参数变更
      - `+ Sp-Auth-Token`
  - **ShowDeviceTunnel**
    - 请求参数变更
      - `+ Sp-Auth-Token`
  - **AddTunnel**
    - 请求参数变更
      - `+ Sp-Auth-Token`
  - **ListDeviceTunnels**
    - 请求参数变更
      - `+ Sp-Auth-Token`
  - **ShowDeviceMessage**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`
  - **CreateMessage**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`
  - **ListDeviceMessages**
    - 请求参数变更
      - `+ Stage-Auth-Token`
      - `+ Sp-Auth-Token`

### HuaweiCloud SDK Live

- _新增特性_
  - 支持以下接口：
    - `ListDelayConfig`
    - `UpdateDelayConfig`
    - `ShowPullSourcesConfig`
    - `UpdatePullSourcesConfig`
    - `ListGeoBlockingConfig`
    - `UpdateGeoBlockingConfig`
    - `CreateUrlAuthchain`
    - `ListIpAuthList`
    - `UpdateIpAuthList`
    - `ListPublishTemplate`
    - `UpdatePublishTemplate`
    - `DeletePublishTemplate`
- _解决问题_
  - 无
- _特性变更_
  - **ListRecordContents**
    - 请求参数变更
      - `+ record_type: enum value [PLAN_RECORD,ON_DEMAND_RECORD]`
    - 响应参数变更
      - `- record_contents.record_type: enum value [PLAN_RECORD,ON_DEMAND_RECORD]`

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持接口`CreateAgencyAccess`
- _解决问题_
  - 无
- _特性变更_
  - **ListSqlAlarmRules**
    - 响应参数变更
      - `+ sql_alarm_rules.is_css_sql`
      - `+ sql_alarm_rules.notification_frequency`
      - `+ sql_alarm_rules.alarm_action_rule_name`
      - `+ sql_alarm_rules.status: enum value [RUNNING 启用,STOPPING 停止]`
      - `- sql_alarm_rules.status: enum value [RUNNING,STOPPING]`
      - `* sql_alarm_rules.frequency: object<Frequency> -> object<FrequencyRespBody>`
  - **UpdateSqlAlarmRule**
    - 请求参数变更
      - `+ is_css_sql`
      - `+ notification_frequency`
      - `+ alarm_action_rule_name`
      - `* frequency: object<Frequency> -> object<CreateSqlAlarmRuleFrequency>`
    - 响应参数变更
      - `+ is_css_sql`
      - `+ alarm_action_rule_name`
      - `+ notification_frequency`
      - `+ language: enum value [zh-cn,en-us]`
      - `* frequency: object<Frequency> -> object<FrequencyRespBody>`
  - **CreateSqlAlarmRule**
    - 请求参数变更
      - `+ is_css_sql`
      - `+ notification_frequency`
      - `+ alarm_action_rule_name`
      - `* frequency: object<Frequency> -> object<CreateSqlAlarmRuleFrequency>`
  - **ListKeywordsAlarmRules**
    - 响应参数变更
      - `+ keywords_alarm_rules.notification_frequency`
      - `+ keywords_alarm_rules.alarm_action_rule_name`
      - `+ keywords_alarm_rules.status: enum value [RUNNING  启用,STOPPING  停止]`
      - `- keywords_alarm_rules.status: enum value [RUNNING,STOPPING]`
  - **UpdateKeywordsAlarmRule**
    - 请求参数变更
      - `+ notification_frequency`
      - `+ alarm_action_rule_name`
    - 响应参数变更
      - `+ alarm_action_rule_name`
      - `+ notification_frequency`
      - `+ language: enum value [zh-cn,en-us]`
      - `- keywords_requests.search_time_range_unit: enum value [minute]`
      - `* keywords_requests: list<KeywordsRequest> -> list<KeywordsResBody>`
      - `* frequency: object<Frequency> -> object<FrequencyRespBody>`
  - **CreateKeywordsAlarmRule**
    - 请求参数变更
      - `+ notification_frequency`
      - `+ alarm_action_rule_name`
      - `+ keywords_alarm_level: enum value [Critical]`
      - `- keywords_alarm_level: enum value [CRITICAL]`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`ShowMrsVersionMetadata`、`SwitchClusterTags`、`ShowTagStatus`、`ShowTagQuota`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK NAT

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateNatGatewayTag**
    - 请求参数变更
      - `+ tag.key`
      - `+ tag.value`
      - `* tag: object -> object<TagBody>`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListLogLtsConfigs`、`SetLogLtsConfigs`、`DeleteLogLtsConfigs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateRocketMqMigrationTask**
    - 请求参数变更
      - `+ topicConfigTable`
      - `+ subscriptionGroupTable`
      - `+ vhosts`
      - `+ queues`
      - `+ exchanges`
      - `+ bindings`
      - `* body: map<string, object> -> object<CreateRocketMqMigrationTaskReq>`

### HuaweiCloud SDK SecMaster

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAlerts**
    - 请求参数变更
      - `* condition.conditions.data: list<object> -> list<string>`
      - `* condition.logics: list<object> -> list<string>`
  - **ListIncidents**
    - 请求参数变更
      - `* condition.conditions.data: list<object> -> list<string>`
      - `* condition.logics: list<object> -> list<string>`
  - **ListIndicators**
    - 请求参数变更
      - `* condition: string -> object`
  - **ListDataobjectRelations**
    - 请求参数变更
      - `* condition.conditions.data: list<object> -> list<string>`
      - `* condition.logics: list<object> -> list<string>`

### HuaweiCloud SDK SMS

- _新增特性_
  - 支持接口`ShowConsistencyResult`、`UpdateConsistencyResult`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateTaskStatus**
    - 请求参数变更
      - `+ is_need_consistency_check`
  - **ListServers**
    - 请求参数变更
      - `+ is_consistency_result_exist`
    - 响应参数变更
      - `+ source_servers.is_consistency_result_exist`

# 0.1.71 2023-12-07

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`DeleteStackInstance`
- _解决问题_
  - 无
- _特性变更_
  - **ListStackSetOperations**
    - 响应参数变更
      - `+ stack_set_operations.action: enum value [UPDATE_STACK_INSTANCES]`
  - **ShowStackSetOperationMetadata**
    - 响应参数变更
      - `+ action: enum value [UPDATE_STACK_INSTANCES]`

### HuaweiCloud SDK APIG

- _新增特性_
  - 支持接口`CheckApiGroupsV2`
- _解决问题_
  - 无
- _特性变更_
  - **CreatePrepayResize**
    - 请求参数变更
      - `+ instance_id`
  - **ListPluginAttachableApis**
    - 请求参数变更
      - `* env_id: optional -> required`
  - **ListApisV2**
    - 请求参数变更
      - `+ return_data_mode`

### HuaweiCloud SDK CBH

- _新增特性_
  - 支持接口`LoginCbh`
- _解决问题_
  - 无
- _特性变更_
  - **ShowAvailableZoneInfo**
    - 响应参数变更
      - `* availability_zone: object<AvailabilityZones> -> list<AvailabilityZones>`

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 支持接口`ShowProjectDataDashboard`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeArtsPipeline

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowPipelineLog**
    - 请求参数变更
      - `- level`
      - `- job_run_id`

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 支持以下接口：
    - `SetFactoryJobTags`
    - `ListSecurityPermissionSets`
    - `CreateSecurityPermissionSet`
    - `ShowSecurityPermissionSet`
    - `UpdateSecurityPermissionSet`
    - `DeleteSecurityPermissionSet`
    - `ListSecurityPermissionSetMembers`
    - `CreateSecurityPermissionSetMember`
    - `BatchDeleteSecurityPermissionSetMembers`
    - `ListSecurityPermissionSetPermissions`
    - `CreateSecurityPermissionSetPermission`
    - `BatchDeleteSecurityPermissionSetPermissions`
    - `UpdateSecurityPermissionSetPermission`
    - `ListSecurityDataClassificationRules`
    - `CreateSecurityDataClassificationRule`
    - `ShowSecurityDataClassificationRule`
    - `UpdateSecurityDataClassificationRule`
    - `DeleteSecurityDataClassificationRule`
    - `BatchDeleteSecurityDataClassificationRule`
    - `UpdateSecurityRuleEnableStatus`
    - `ListSecurityDataClassificationRuleGroups`
    - `ShowSecurityDataClassificationRuleGroup`
- _解决问题_
  - 无
- _特性变更_
  - **DeleteApi**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **PublishApi**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowApplyDetail**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowMessageDetail**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowCatalogDetail**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **UpdateCatalog**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **CreateServiceCatalog**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **DeleteServiceCatalog**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **MigrateCatalog**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **MigrateApi**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **SearchIdByPath**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowPathById**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **PublishApiToInstance**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ExecuteApiToInstance**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **AuthorizeApiToInstance**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **AuthorizeActionApiToInstance**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **DeleteApp**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowAppInfo**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **UpdateApp**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowApisOverview**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowAppsOverview**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowApisDetail**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowAppsDetail**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **UpdateFactoryJobName**
    - 请求参数变更
      - `- x-Auth-Token`
  - **BatchApproveApply**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListApply**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ConfirmMessage**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListMessage**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListAllCatalogList**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListCatalogList**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowPathObjectById**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **DebugApi**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **SearchPublishInfo**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListInstanceList**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **SearchDebugInfo**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListApicInstances**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListApicGroups**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **CreateApp**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListApps**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListApisTop**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListAppsTop**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowApisDashboard**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowApiDashboard**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowAppsDashboard**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListApiTopN**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListApis**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **CreateApi**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ShowApi**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **UpdateApi**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **ListApiCatalogList**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **SearchAuthorizeApp**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`
  - **SearchBindApi**
    - 请求参数变更
      - `* Dlm-Type: required -> optional`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowJobInstance**
    - 响应参数变更
      - `* planTime: int32 -> int64`
      - `* startTime: int32 -> int64`
      - `* endTime: int32 -> int64`
      - `* executeTime: int32 -> int64`
      - `* nodes.planTime: int32 -> int64`
      - `* nodes.startTime: int32 -> int64`
      - `* nodes.endTime: int32 -> int64`
      - `* nodes.executeTime: int32 -> int64`
  - **ListJobs**
    - 请求参数变更
      - `+ tags`

### HuaweiCloud SDK DLI

- _新增特性_
  - 支持接口`ListJobAuthInfos`、`UpdateJobAuthInfo`、`CreateJobAuthInfo`、`DeleteJobAuthInfo`
- _解决问题_
  - 无
- _特性变更_
  - 废弃以下接口：
    - `ShowDliAgency`
    - `CreateDliAgency`
    - `DeleteResource`
    - `ShowResourceInfo`
    - `UpdateGroupOrResourceOwner`
    - `ShowBatchLog`
    - `ExportTable`
    - `ImportTable`
    - `ExportSqlJobResult`
    - `UpdateDatabaseOwner`
    - `DeleteDatabase`
    - `RegisterAuthorizedQueue`
    - `UpdateTableOwner`
    - `ShowTableContent`
    - `UpdateQueueCidr`
    - `BatchDeleteQueuePlans`
    - `ChangeQueuePlan`
    - `DeleteQueuePlan`
    - `DeleteAuthInfo`
    - `DeleteEnhancedConnectionRoutes`
    - `CreateEnhancedConnectionRoutes`
    - `RegisterBucket`
    - `CreateIefMessageChannel`
    - `UploadFiles`
    - `UploadPythonFiles`
    - `ListResources`
    - `UploadResources`
    - `UploadJars`
    - `ListDatabases`
    - `CreateDatabase`
    - `ListTableUsers`
    - `ChangeAuthorization`
    - `ListTablePrivileges`
    - `ListDatabaseUsers`
    - `ListQueueUsers`
    - `ListAllTables`
    - `CreateTable`
    - `DeleteTable`
    - `ShowDescribeTable`
    - `CreateQueuePlan`
    - `ListQueuePlans`
    - `UpdateAuthInfo`
    - `CreateAuthInfo`
    - `ListAuthInfo`
    - `ChangeFlinkJobStatusReport`
    - `RunIefJobActionCallBack`
    - `CreateIefSystemEvents`
    - `ListDatasourceConnections`
    - `CreateDatasourceConnection`
    - `DeleteDatasourceConnection`
    - `ShowDatasourceConnection`
    - `ShowSqlSampleTemplates`
    - `ShowPartitions`
    - `ShowFlinkMetric`

### HuaweiCloud SDK EC

- _新增特性_
  - 支持接口`ShowEquipmentWlan`、`UpdateEquipmentWlan`
- _解决问题_
  - 无
- _特性变更_
  - **ShowEquipmentInfo**
    - 响应参数变更
      - `+ type: enum value [soho]`
  - **UpdateEquipmentInfo**
    - 响应参数变更
      - `+ type: enum value [soho]`
  - **GenerateInitialConfiguration**
    - 响应参数变更
      - `+ url_config_content`
      - `+ script_config_content`
      - `- config_content`
  - **ShowEquipmentSpecificConfig**
    - 请求参数变更
      - `+ equipment_id`
      - `- equipment_type`
    - 响应参数变更
      - `- model`
  - **CreateEquipment**
    - 请求参数变更
      - `+ type: enum value [soho]`
    - 响应参数变更
      - `+ type: enum value [soho]`
  - **ListEquipments**
    - 响应参数变更
      - `+ equipments.type: enum value [soho]`
  - **ShowIegInfo**
    - 响应参数变更
      - `+ equipment_infos.type: enum value [soho]`
  - **UpdateIeg**
    - 响应参数变更
      - `+ equipment_infos.type: enum value [soho]`
  - **ListIeg**
    - 响应参数变更
      - `+ intelligent_enterprise_gateways.equipment_infos.type: enum value [soho]`

### HuaweiCloud SDK EG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDetailOfEvent**
    - 请求参数变更
      - `+ channel_id`
  - **DeleteChannel**
    - 请求参数变更
      - `+ enterprise_project_id`
  - **ShowDetailOfChannel**
    - 请求参数变更
      - `+ enterprise_project_id`
    - 响应参数变更
      - `+ eps_id`
  - **UpdateChannel**
    - 请求参数变更
      - `+ enterprise_project_id`
      - `+ eps_id`
      - `+ cross_account`
      - `+ policy`
    - 响应参数变更
      - `+ eps_id`
  - **CreateChannel**
    - 请求参数变更
      - `+ enterprise_project_id`
      - `+ eps_id`
      - `+ cross_account`
      - `+ policy`
    - 响应参数变更
      - `+ eps_id`
  - **ListChannels**
    - 请求参数变更
      - `+ eps_id`
    - 响应参数变更
      - `+ eps_id`
      - `+ items.eps_id`
  - **CreateSubscriptionTarget**
    - 请求参数变更
      - `+ eg_detail`
      - `- detail.url`
      - `- detail.agency_name`
      - `- detail.invocation_http_parameters`
      - `* detail: object<Detail> -> object`
  - **UpdateSubscriptionTarget**
    - 请求参数变更
      - `+ eg_detail`
      - `- detail.url`
      - `- detail.agency_name`
      - `- detail.invocation_http_parameters`
      - `* detail: object<Detail> -> object`
  - **ShowDetailOfConnection**
    - 响应参数变更
      - `+ kafka_detail.security_protocol`
  - **UpdateConnection**
    - 响应参数变更
      - `+ kafka_detail.security_protocol`
  - **UpdateSubscription**
    - 请求参数变更
      - `+ targets.eg_detail`
      - `- targets.detail.url`
      - `- targets.detail.agency_name`
      - `- targets.detail.invocation_http_parameters`
      - `* targets.detail: object<Detail> -> object`
  - **CreateConnection**
    - 请求参数变更
      - `+ kafka_detail.security_protocol`
    - 响应参数变更
      - `+ kafka_detail.security_protocol`
  - **ListConnections**
    - 响应参数变更
      - `+ items.kafka_detail.security_protocol`
  - **ShowEventStreaming**
    - 响应参数变更
      - `+ source.source_mobile_rocketmq`
      - `+ source.source_kafka.security_protocol`
      - `+ sink.sink_kafka`
      - `+ sink.name: enum value [HC.FunctionGraph,HC.Kafka]`
      - `- sink.name: enum value [HC.FG]`
  - **UpdateEventStreaming**
    - 请求参数变更
      - `+ source.source_mobile_rocketmq`
      - `+ source.source_kafka.security_protocol`
      - `+ sink.sink_kafka`
      - `+ sink.name: enum value [HC.FunctionGraph,HC.Kafka]`
      - `- sink.name: enum value [HC.FG]`
  - **CreateSubscription**
    - 请求参数变更
      - `+ targets.eg_detail`
      - `- targets.detail.url`
      - `- targets.detail.agency_name`
      - `- targets.detail.invocation_http_parameters`
      - `* targets.detail: object<Detail> -> object`
  - **CreateEventStreaming**
    - 请求参数变更
      - `+ source.source_mobile_rocketmq`
      - `+ source.source_kafka.security_protocol`
      - `+ sink.sink_kafka`
      - `+ sink.name: enum value [HC.FunctionGraph,HC.Kafka]`
      - `- sink.name: enum value [HC.FG]`
  - **ListEventStreaming**
    - 请求参数变更
      - `+ name`
      - `+ fuzzy_name`
    - 响应参数变更
      - `+ source.source_mobile_rocketmq`
      - `+ source.source_kafka.security_protocol`
      - `+ sink.sink_kafka`
      - `+ sink.name: enum value [HC.FunctionGraph,HC.Kafka]`
      - `- sink.name: enum value [HC.FG]`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `ShowDependcy`
    - `UpdateDependcy`
    - `DeleteDependency`
    - `AsyncInvokeReservedFunction`
    - `CreateDependency`
  - **ShowFuncReservedInstanceMetrics**
    - 请求参数变更
      - `+ marker`
      - `+ limit`
  - **ListFunctionApplications**
    - 请求参数变更
      - `+ limit`
      - `+ marker`
  - **ListStatistics**
    - 请求参数变更
      - `+ limit`
      - `+ marker`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ListGaussMySqlInstancesUnifyStatus`、`ShowGaussMySqlInstanceInfoUnifyStatus`、`ListGaussMySqlInstanceDetailInfoUnifyStatus`、`SwitchGaussMySqlProxySsl`
- _解决问题_
  - 无
- _特性变更_
  - **ShowGaussMySqlProxyList**
    - 响应参数变更
      - `+ proxy_list.proxy.ssl_option`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持接口`OfflineNodes`
- _解决问题_
  - 无
- _特性变更_
  - **ListLtsConfigs**
    - 响应参数变更
      - `* instance_lts_configs.instance.supported_log_types: string -> list<string>`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ResetDeviceSecret**
    - 请求参数变更
      - `+ secret_type`
    - 响应参数变更
      - `+ secret_type`
  - **ResetFingerprint**
    - 请求参数变更
      - `+ fingerprint_type`
    - 响应参数变更
      - `+ fingerprint_type`
  - **ShowDevice**
    - 响应参数变更
      - `+ auth_info.secondary_secret`
      - `+ auth_info.secondary_fingerprint`
      - `* auth_info: object<AuthInfo> -> object<AuthInfoRes>`
  - **UpdateDevice**
    - 响应参数变更
      - `+ auth_info.secondary_secret`
      - `+ auth_info.secondary_fingerprint`
      - `* auth_info: object<AuthInfo> -> object<AuthInfoRes>`
  - **AddDevice**
    - 响应参数变更
      - `+ auth_info.secondary_secret`
      - `+ auth_info.secondary_fingerprint`
      - `* auth_info: object<AuthInfo> -> object<AuthInfoRes>`

### HuaweiCloud SDK KooMessage

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateAimPersonalTemplate**
    - 请求参数变更
      - `+ pages.contents.button_type`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`UpdateAutoScalingPolicy`、`CreateAutoScalingPolicy`、`DeleteAutoScalingPolicy`
- _解决问题_
  - 无
- _特性变更_
  - **ShowAutoScalingPolicy**
    - 响应参数变更
      - `+ auto_scaling_policy.tags`
      - `- auto_scaling_policy.exec_scripts`
      - `* auto_scaling_policy: object<AutoScalingPolicy> -> object<AutoScalingPolicyInfo>`

### HuaweiCloud SDK NAT

- _新增特性_
  - 支持以下接口：
    - `ListNatGatewayByTag`
    - `BatchCreateDeleteNatGatewayTag`
    - `ShowNatGatewayTag`
    - `CreateNatGatewayTag`
    - `DeleteNatGatewayTag`
    - `ListNatGatewayTag`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OMS

- _新增特性_
  - 支持接口`BatchUpdateTasks`
- _解决问题_
  - 无
- _特性变更_
  - **ShowSyncTask**
    - 响应参数变更
      - `+ dst_storage_policy`
      - `+ object_overwrite_mode`
  - **ListSyncTasks**
    - 响应参数变更
      - `+ tasks.object_overwrite_mode`
      - `+ tasks.dst_storage_policy`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 支持接口`ShowRocketMqConfigs`、`UpdateRocketMqConfigs`
- _解决问题_
  - 无
- _特性变更_
  - **ListInstances**
    - 请求参数变更
      - `+ status: enum value [UPGRADING,UPGRADINGFAILED]`

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **SetHpcCacheBackend**
    - 请求参数变更
      - `* update_hpc_cache.data.nas.type: object -> string`
      - `* update_hpc_cache.data.nas.url: object -> string`

# 0.1.70 2023-11-30

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPermissions**
    - 响应参数变更
      - `* : map<string, AuthModel> -> string`
  - **ListAccessCode**
    - 响应参数变更
      - `- access_codes.status: enum value [enable,unenable]`

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSubCustomerBillDetail**
    - 响应参数变更
      - `* fee_records.usage_amount: double -> bigdecimal`
      - `* fee_records.free_resource_usage: double -> bigdecimal`
      - `* fee_records.ri_usage: double -> bigdecimal`
      - `* fee_records.official_amount: double -> bigdecimal`
      - `* fee_records.official_discount_amount: double -> bigdecimal`
      - `* fee_records.payment_amount: double -> bigdecimal`
      - `* fee_records.cash_amount: double -> bigdecimal`
      - `* fee_records.credit_amount: double -> bigdecimal`
      - `* fee_records.coupon_amount: double -> bigdecimal`
      - `* fee_records.flexipurchase_coupon_amount: double -> bigdecimal`
      - `* fee_records.stored_value_card_amount: double -> bigdecimal`
      - `* fee_records.debt_amount: double -> bigdecimal`
      - `* fee_records.writeoff_amount: double -> bigdecimal`

### HuaweiCloud SDK CC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListBandwidthPackageTags**
    - 响应参数变更
      - `+ tags.values`
      - `- tags.value`
      - `* tags: list<Tag> -> list<MultivaluedTag>`

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持以下接口：
    - `ShowClusterConfig`
    - `UpdateClusterLogConfig`
    - `ListPartitions`
    - `CreatePartition`
    - `ShowPartition`
    - `UpdatePartition`
    - `ShowNodePoolConfigurations`
    - `UpdateNodePoolConfiguration`
    - `ShowClusterConfigurationDetails`
    - `ListCharts`
    - `UploadChart`
    - `ShowChart`
    - `UpdateChart`
    - `DeleteChart`
    - `DownloadChart`
    - `ShowChartValues`
    - `ShowUserChartsQuotas`
    - `ListReleases`
    - `CreateRelease`
    - `ShowRelease`
    - `UpdateRelease`
    - `DeleteRelease`
    - `ShowReleaseHistory`
- _解决问题_
  - 无
- _特性变更_
  - **ResizeCluster**
    - 请求参数变更
      - `* extendParam: object<ResizeClusterRequestExtendParam> -> object`
  - **UpdateClusterEip**
    - 请求参数变更
      - `* spec: object -> object<MasterEIPRequestSpec>`
    - 响应参数变更
      - `* spec: object -> object<MasterEIPResponseSpec>`
  - **ShowClusterEndpoints**
    - 响应参数变更
      - `* spec: object -> object<OpenAPISpec>`
  - **ShowAddonInstance**
    - 响应参数变更
      - `- status.status: enum value [unknown]`
  - **UpdateAddonInstance**
    - 响应参数变更
      - `- status.status: enum value [unknown]`
  - **RollbackAddonInstance**
    - 响应参数变更
      - `- status.status: enum value [unknown]`
  - **ShowCluster**
    - 响应参数变更
      - `+ spec.enableMasterVolumeEncryption`
  - **UpdateCluster**
    - 响应参数变更
      - `+ spec.enableMasterVolumeEncryption`
  - **DeleteCluster**
    - 请求参数变更
      - `+ ondemand_node_policy`
      - `+ periodic_node_policy`
    - 响应参数变更
      - `+ spec.enableMasterVolumeEncryption`
  - **CreateAddonInstance**
    - 响应参数变更
      - `- status.status: enum value [unknown]`
  - **ListAddonInstances**
    - 响应参数变更
      - `- items.status.status: enum value [unknown]`
  - **CreateCluster**
    - 请求参数变更
      - `+ spec.enableMasterVolumeEncryption`
    - 响应参数变更
      - `+ spec.enableMasterVolumeEncryption`
  - **ListClusters**
    - 请求参数变更
      - `+ status: enum value [Hibernating,Hibernation,Awaking]`
    - 响应参数变更
      - `+ items.spec.enableMasterVolumeEncryption`
  - **ShowNode**
    - 响应参数变更
      - `+ spec.hostnameConfig`
      - `+ spec.extendParam.kubeReservedMem`
      - `+ spec.extendParam.systemReservedMem`
      - `+ spec.extendParam.init-node-password`
      - `- spec.extendParam.kube-reserved-mem`
      - `- spec.extendParam.system-reserved-mem`
  - **UpdateNode**
    - 响应参数变更
      - `+ spec.hostnameConfig`
      - `+ spec.extendParam.kubeReservedMem`
      - `+ spec.extendParam.systemReservedMem`
      - `+ spec.extendParam.init-node-password`
      - `- spec.extendParam.kube-reserved-mem`
      - `- spec.extendParam.system-reserved-mem`
  - **DeleteNode**
    - 响应参数变更
      - `+ spec.hostnameConfig`
      - `+ spec.extendParam.kubeReservedMem`
      - `+ spec.extendParam.systemReservedMem`
      - `+ spec.extendParam.init-node-password`
      - `- spec.extendParam.kube-reserved-mem`
      - `- spec.extendParam.system-reserved-mem`
  - **CreateNode**
    - 请求参数变更
      - `+ spec.hostnameConfig`
      - `+ spec.extendParam.kubeReservedMem`
      - `+ spec.extendParam.systemReservedMem`
      - `+ spec.extendParam.init-node-password`
      - `- spec.extendParam.kube-reserved-mem`
      - `- spec.extendParam.system-reserved-mem`
    - 响应参数变更
      - `+ spec.hostnameConfig`
      - `+ spec.extendParam.kubeReservedMem`
      - `+ spec.extendParam.systemReservedMem`
      - `+ spec.extendParam.init-node-password`
      - `- spec.extendParam.kube-reserved-mem`
      - `- spec.extendParam.system-reserved-mem`
  - **ListNodes**
    - 响应参数变更
      - `+ items.spec.hostnameConfig`
      - `+ items.spec.extendParam.kubeReservedMem`
      - `+ items.spec.extendParam.systemReservedMem`
      - `+ items.spec.extendParam.init-node-password`
      - `- items.spec.extendParam.kube-reserved-mem`
      - `- items.spec.extendParam.system-reserved-mem`
  - **ShowNodePool**
    - 响应参数变更
      - `+ spec.nodeTemplate.hostnameConfig`
      - `+ spec.nodeTemplate.extendParam.kubeReservedMem`
      - `+ spec.nodeTemplate.extendParam.systemReservedMem`
      - `+ spec.nodeTemplate.extendParam.init-node-password`
      - `- spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `- spec.nodeTemplate.extendParam.system-reserved-mem`
  - **UpdateNodePool**
    - 响应参数变更
      - `+ spec.nodeTemplate.hostnameConfig`
      - `+ spec.nodeTemplate.extendParam.kubeReservedMem`
      - `+ spec.nodeTemplate.extendParam.systemReservedMem`
      - `+ spec.nodeTemplate.extendParam.init-node-password`
      - `- spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `- spec.nodeTemplate.extendParam.system-reserved-mem`
  - **DeleteNodePool**
    - 响应参数变更
      - `+ spec.nodeTemplate.hostnameConfig`
      - `+ spec.nodeTemplate.extendParam.kubeReservedMem`
      - `+ spec.nodeTemplate.extendParam.systemReservedMem`
      - `+ spec.nodeTemplate.extendParam.init-node-password`
      - `- spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `- spec.nodeTemplate.extendParam.system-reserved-mem`
  - **AddNode**
    - 请求参数变更
      - `+ nodeList.spec.hostnameConfig`
  - **ResetNode**
    - 请求参数变更
      - `+ nodeList.spec.hostnameConfig`
  - **CreateNodePool**
    - 请求参数变更
      - `+ spec.nodeTemplate.hostnameConfig`
      - `+ spec.nodeTemplate.extendParam.kubeReservedMem`
      - `+ spec.nodeTemplate.extendParam.systemReservedMem`
      - `+ spec.nodeTemplate.extendParam.init-node-password`
      - `- spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `- spec.nodeTemplate.extendParam.system-reserved-mem`
    - 响应参数变更
      - `+ spec.nodeTemplate.hostnameConfig`
      - `+ spec.nodeTemplate.extendParam.kubeReservedMem`
      - `+ spec.nodeTemplate.extendParam.systemReservedMem`
      - `+ spec.nodeTemplate.extendParam.init-node-password`
      - `- spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `- spec.nodeTemplate.extendParam.system-reserved-mem`
  - **ListNodePools**
    - 响应参数变更
      - `+ items.spec.nodeTemplate.hostnameConfig`
      - `+ items.spec.nodeTemplate.extendParam.kubeReservedMem`
      - `+ items.spec.nodeTemplate.extendParam.systemReservedMem`
      - `+ items.spec.nodeTemplate.extendParam.init-node-password`
      - `- items.spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `- items.spec.nodeTemplate.extendParam.system-reserved-mem`

### HuaweiCloud SDK CodeArtsPipeline

- _新增特性_
  - 支持接口`ShowPipelineLog`
- _解决问题_
  - 无
- _特性变更_
  - **UpdatePluginBaseInfo**
    - 请求参数变更
      - `+ plugin_composition_type`
  - **CreatePublisher**
    - 请求参数变更
      - `+ description`
  - **ListPublisher**
    - 响应参数变更
      - `+ total`
      - `+ offset`
      - `+ data`
      - `+ limit`
      - `- tenant_id`
      - `- website`
      - `- logo_url`
      - `- description`
      - `- last_update_user_id`
      - `- source_url`
      - `- is_delete`
      - `- last_update_time`
      - `- support_url`
      - `- user_id`
      - `- name`
      - `- en_name`
      - `- auth_status`
      - `- publisher_unique_id`
      - `- last_update_user_name`
  - **ListBasePluginsNewPost**
    - 响应参数变更
      - `+ data.business_type`
      - `+ data.display_name`
      - `+ data.unique_id`
      - `- data.businessType`
      - `- data.displayName`
      - `- data.uniqueId`
      - `+ data.plugins_list.publisher_unique_id`
      - `+ data.plugins_list.manifest_version`
      - `- data.plugins_list.publisherUniqueId`
      - `- data.plugins_list.manifestVersion`
  - **ListPlugins**
    - 响应参数变更
      - `+ data.plugin_name`
      - `+ data.display_name`
      - `+ data.version_description`
      - `+ data.version_attribution`
      - `+ data.unique_id`
      - `+ data.op_user`
      - `+ data.op_time`
      - `+ data.plugin_composition_type`
      - `+ data.plugin_attribution`
      - `+ data.workspace_id`
      - `+ data.business_type`
      - `+ data.business_type_display_name`
      - `+ data.icon_url`
      - `+ data.refer_count`
      - `+ data.usage_count`
      - `+ data.runtime_attribution`
      - `- data.pluginName`
      - `- data.displayName`
      - `- data.versionDescription`
      - `- data.versionAttribution`
      - `- data.uniqueId`
      - `- data.opUser`
      - `- data.opTime`
      - `- data.pluginCompositionType`
      - `- data.pluginAttribution`
      - `- data.workspaceId`
      - `- data.businessType`
      - `- data.businessTypeDisplayName`
      - `- data.iconUrl`
      - `- data.referCount`
      - `- data.usageCount`
      - `- data.runtimeAttribution`
      - `* data: list<object> -> list<PluginBasicVO>`
  - **ListPLuginVersion**
    - 响应参数变更
      - `+ data.plugin_name`
      - `+ data.display_name`
      - `+ data.version_description`
      - `+ data.version_attribution`
      - `+ data.unique_id`
      - `+ data.op_user`
      - `+ data.op_time`
      - `+ data.plugin_composition_type`
      - `+ data.plugin_attribution`
      - `+ data.workspace_id`
      - `+ data.business_type`
      - `+ data.business_type_display_name`
      - `+ data.icon_url`
      - `+ data.refer_count`
      - `+ data.usage_count`
      - `+ data.runtime_attribution`
      - `- data.pluginName`
      - `- data.displayName`
      - `- data.versionDescription`
      - `- data.versionAttribution`
      - `- data.uniqueId`
      - `- data.opUser`
      - `- data.opTime`
      - `- data.pluginCompositionType`
      - `- data.pluginAttribution`
      - `- data.workspaceId`
      - `- data.businessType`
      - `- data.businessTypeDisplayName`
      - `- data.iconUrl`
      - `- data.referCount`
      - `- data.usageCount`
      - `- data.runtimeAttribution`
      - `* data: list<object> -> list<PluginBasicVO>`
  - **ShowPublisher**
    - 响应参数变更
      - `+ publisher_detail_map`
      - `- body`
  - **CreateBasicPlugin**
    - 请求参数变更
      - `+ plugin_composition_type`
  - **UpdateBasicPlugin**
    - 请求参数变更
      - `+ plugin_composition_type`
  - **CreateStrategy**
    - 请求参数变更
      - `- type`
      - `- rules.type`
      - `- rules.name`
      - `- rules.layout_content`
      - `- rules.plugin_id`
      - `- rules.plugin_name`
      - `- rules.plugin_version`
      - `- rules.content`
  - **UpdateStrategy**
    - 请求参数变更
      - `- parent_id`
      - `- rules.type`
      - `- rules.name`
      - `- rules.layout_content`
      - `- rules.plugin_id`
      - `- rules.plugin_name`
      - `- rules.plugin_version`
      - `- rules.content`

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **SearchAtomicIndexes**
    - 响应参数变更
      - `+ data.value`
  - **ShowAtomicIndexById**
    - 响应参数变更
      - `+ data.value`
  - **ListDerivativeIndexes**
    - 响应参数变更
      - `+ data.value`
  - **ShowDerivativeIndexById**
    - 响应参数变更
      - `+ data.value`
  - **ListCompoundMetrics**
    - 响应参数变更
      - `+ data.value`
  - **ShowCompoundMetricById**
    - 响应参数变更
      - `+ data.value`
  - **ListApiCatalogList**
    - 响应参数变更
      - `+ apis.publish_messages`
  - **ParseUserBehavior**
    - 请求参数变更
      - `+ params.filter.attribute: enum value [base.DataAsset.sourceType,typeName,classifications.name,tags.name,securityLevel.name,workspaceId]`
      - `+ params.filter.operator: enum value [IN,EQ]`
      - `* params.filter.value: object -> list<string>`
      - `* params.filter.condition: object<ConditionInfo> -> string`
  - **ShowDataSets**
    - 请求参数变更
      - `+ filter.attribute: enum value [base.DataAsset.sourceType,typeName,classifications.name,tags.name,securityLevel.name,workspaceId]`
      - `+ filter.operator: enum value [IN,EQ]`
      - `* filter.value: object -> list<string>`
      - `* filter.condition: object<ConditionInfo> -> string`
    - 响应参数变更
      - `* facets: object -> list<object>`
  - **ListApis**
    - 请求参数变更
      - `+ x-return-publish-messages`
    - 响应参数变更
      - `+ records.publish_messages`
  - **ShowApi**
    - 响应参数变更
      - `+ publish_messages`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSlowlog**
    - 响应参数变更
      - `+ total_num`

### HuaweiCloud SDK DLI

- _新增特性_
  - 支持接口`ShowQuota`
- _解决问题_
  - 无
- _特性变更_
  - **ListQueueProperties**
    - 请求参数变更
      - `+ offset`
      - `+ limit`
      - `- page`
      - `- page_size`

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持接口`BatchModifyBandwidth`、`ListEipBandwidths`、`ListBandwidthsLimit`、`UpdatePublicip`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持接口`ListInfluxdbSlowLogs`
- _解决问题_
  - 无
- _特性变更_
  - **ListLtsConfigs**
    - 响应参数变更
      - `+ instance_lts_configs.instance.supported_log_types`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RunQueryDocumentModerationJob**
    - 响应参数变更
      - `+ result.details.start_position`
      - `+ result.details.end_position`
      - `+ result.details.image_url`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`SetInstancesDbShrink`
- _解决问题_
  - 无
- _特性变更_
  - **UpgradeDbMajorVersion**
    - 响应参数变更
      - `+ job_id`
  - **ShowUpgradeDbMajorVersionStatus**
    - 响应参数变更
      - `+ check_expiration_time`
      - `- report_expiration_time`
  - **StartResizeFlavorAction**
    - 响应参数变更
      - `+ order_id`
  - **StartInstanceEnlargeVolumeAction**
    - 响应参数变更
      - `+ order_id`
  - **StartInstanceSingleToHaAction**
    - 响应参数变更
      - `+ order_id`
  - **ListHistoryDatabase**
    - 请求参数变更
      - `+ engine`
      - `- database_name`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 支持接口`ShowEngineInstanceExtendProductInfo`、`ResizeInstance`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeFlashAsr**
    - 请求参数变更
      - `* obs_bucket_name: optional -> required`
      - `* obs_object_key: optional -> required`

# 0.1.69 2023-11-23

### HuaweiCloud SDK CFW

- _新增特性_
  - 支持接口`ListLogConfig`、`UpdateLogConfig`、`AddLogConfig`、`CreateEastWestFirewall`
- _解决问题_
  - 无
- _特性变更_
  - **ListFlowLogs**
    - 响应参数变更
      - `* data.records.start_time: int32 -> int64`
      - `* data.records.end_time: int32 -> int64`
      - `* data.records.src_port: string -> int32`
      - `* data.records.dst_port: string -> int32`
  - **ListAccessControlLogs**
    - 响应参数变更
      - `* data.records.hit_time: int32 -> int64`
      - `* data.records.src_port: string -> int32`
      - `* data.records.dst_port: string -> int32`
  - **ChangeIpsSwitchStatus**
    - 请求参数变更
      - `+ X-Language`
  - **ListAttackLogs**
    - 响应参数变更
      - `* data.records.event_time: string -> int64`
      - `* data.records.attack_rule_id: int32 -> string`
      - `* data.records.packet: object<Packet> -> string`

# 0.1.68 2023-11-23

### HuaweiCloud SDK AOM

- _新增特性_
  - 支持以下接口：
    - `ListPromInstance`
    - `CreatePromInstance`
    - `DeletePromInstance`
    - `CreateRecordingRule`
    - `ListPermissions`
    - `ListAccessCode`
    - `ListAgents`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeArtsPipeline

- _新增特性_
  - 支持以下接口：
    - `UpdatePluginBaseInfo`
    - `CreatePluginDraft`
    - `PublishPluginDraft`
    - `CreatePluginVersion`
    - `UpdatePluginDraft`
    - `CreatePublisher`
- _解决问题_
  - 无
- _特性变更_
  - **SwitchStrategy**
    - 响应参数变更
      - `+ rule_set_id`
      - `- rule_template_instance_id`
  - **SwitchOpenSourceStrategy**
    - 响应参数变更
      - `+ rule_set_id`
      - `- rule_template_instance_id`
  - **ShowPublisher**
    - 响应参数变更
      - `+ body`
  - **CreatePipelineNew**
    - 请求参数变更
      - `+ group_id`
      - `+ id`
      - `* schedules.days_of_week: string -> list<integer>`
  - **UpdateStrategy**
    - 响应参数变更
      - `+ rule_set_id`
      - `- rule_template_instance_id`
  - **UpdateOpenSourceStrategy**
    - 响应参数变更
      - `+ rule_set_id`
      - `- rule_template_instance_id`

### HuaweiCloud SDK CSE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpgradeEngineConfig**
    - 请求参数变更
      - `+ authType`
      - `- version`
  - **ShowEngine**
    - 响应参数变更
      - `+ specType`
      - `+ enterpriseProjectId`
      - `+ externalEntrypoint`
      - `+ beDefault`
      - `+ engineAdditionalActions`
      - `+ cceSpec`
      - `+ vmIds`
      - `+ latestVersion`
      - `+ createTime`
      - `+ createUser`
      - `+ authType`
      - `+ latestJobId`
      - `+ projectId`
      - `+ enterpriseProjectName`
      - `- auth_type`
      - `- create_time`
      - `- be_default`
      - `- enterprise_project_name`
      - `- latest_job_id`
      - `- spec_type`
      - `- external_entrypoint`
      - `- cce_spec`
      - `- enterprise_project_id`
      - `- latest_version`
      - `- project_id`
      - `- vm_ids`
      - `- engine_additional_actions`
      - `- create_user`
      - `+ reference.azList`
      - `+ reference.networkId`
      - `+ reference.subnetCidr`
      - `+ reference.subnetCidrV6`
      - `+ reference.subnetGateway`
      - `+ reference.publicIpId`
      - `+ reference.serviceLimit`
      - `+ reference.instanceLimit`
      - `- reference.az_list`
      - `- reference.network_id`
      - `- reference.subnet_cidr`
      - `- reference.subnet_cidr_v6`
      - `- reference.subnet_gateway`
      - `- reference.public_ip_id`
      - `- reference.service_limit`
      - `- reference.instance_limit`
  - **CreateEngine**
    - 请求参数变更
      - `+ vpcId`
  - **ListEngines**
    - 响应参数变更
      - `+ data.enterpriseProjectId`
      - `+ data.enterpriseProjectName`
      - `+ data.authType`
      - `+ data.externalAddress`
      - `+ data.serviceEndpoint`
      - `+ data.publicAddress`
      - `+ data.publicServiceEndpoint`
      - `+ data.totalInstance`
      - `+ data.usedInstance`
      - `+ data.availableInstance`
      - `+ data.latestVersion`
      - `+ data.createTime`
      - `+ data.dueTo`
      - `+ data.latestJobId`
      - `+ data.engineAdditionalActions`
      - `+ data.specType`
      - `- data.enterpris_project_id`
      - `- data.enterprise_project_name`
      - `- data.auth_type`
      - `- data.external_address`
      - `- data.service_endpoint`
      - `- data.public_address`
      - `- data.public_service_endpoint`
      - `- data.total_instance`
      - `- data.used_instance`
      - `- data.available_instance`
      - `- data.latest_version`
      - `- data.create_time`
      - `- data.due_to`
      - `- data.latest_job_id`
      - `- data.engine_additional_actions`
      - `- data.spec_type`
      - `+ data.reference.azList`
      - `+ data.reference.networkId`
      - `+ data.reference.subnetCidr`
      - `+ data.reference.subnetCidrV6`
      - `+ data.reference.subnetGateway`
      - `+ data.reference.publicIpId`
      - `+ data.reference.serviceLimit`
      - `+ data.reference.instanceLimit`
      - `- data.reference.az_list`
      - `- data.reference.network_id`
      - `- data.reference.subnet_cidr`
      - `- data.reference.subnet_cidr_v6`
      - `- data.reference.subnet_gateway`
      - `- data.reference.public_ip_id`
      - `- data.reference.service_limit`
      - `- data.reference.instance_limit`
  - **ShowEngineJob**
    - 响应参数变更
      - `+ createUser`
      - `+ startTime`
      - `+ endTime`
      - `+ engineId`
      - `- start_time`
      - `- engine_id`
      - `- end_time`
      - `- create_user`
      - `+ tasks.taskName`
      - `+ tasks.taskNames`
      - `+ tasks.startTime`
      - `+ tasks.endTime`
      - `+ tasks.taskExecutorBrief`
      - `- tasks.task_name`
      - `- tasks.task_names`
      - `- tasks.start_time`
      - `- tasks.end_time`
      - `- tasks.task_executor_brief`
      - `+ tasks.tasks.jobId`
      - `+ tasks.tasks.taskName`
      - `+ tasks.tasks.engineName`
      - `+ tasks.tasks.taskOrder`
      - `+ tasks.tasks.startTime`
      - `+ tasks.tasks.endTime`
      - `+ tasks.tasks.createTime`
      - `+ tasks.tasks.updateTime`
      - `+ tasks.tasks.taskExecutorBrief`
      - `- tasks.tasks.job_id`
      - `- tasks.tasks.task_name`
      - `- tasks.tasks.engine_name`
      - `- tasks.tasks.task_order`
      - `- tasks.tasks.start_time`
      - `- tasks.tasks.end_time`
      - `- tasks.tasks.create_time`
      - `- tasks.tasks.update_time`
      - `- tasks.tasks.task_executor_brief`
  - **CreateMicroserviceRouteRule**
    - 请求参数变更
      - `+ match.headers.<header>`
      - `- match.headers.aaaa`
      - `+ route.tags.<tag>`
      - `- route.tags.version`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListCerts**
    - 响应参数变更
      - `* defaultCerts: object<DefaultCertsResource> -> list<DefaultCertsResource>`
      - `* customCerts: object<CustomCertsResource> -> list<CustomCertsResource>`

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`LogoffWebCli`
- _解决问题_
  - 无
- _特性变更_
  - **ListSlowlog**
    - 响应参数变更
      - `+ slowlogs.database_id`
      - `+ slowlogs.username`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持以下接口：
    - `ListLtsConfigs`
    - `UpdateLtsConfig`
    - `DeleteLtsConfig`
    - `ListLtsErrorLogs`
    - `ShowKillOpRuleRuleList`
    - `UpdateKillOpRule`
    - `CreateKillOpRule`
    - `DeleteKillOpRuleList`
    - `SwitchInstancePrimary`
    - `DeleteReadonlyNode`
    - `StopBackup`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchDeleteJobs**
    - 请求参数变更
      - `+ jobs.is_show_breakpoint_position`
  - **BatchSetPolicy**
    - 请求参数变更
      - `+ jobs.file_and_position`
      - `+ jobs.gtid_set`
  - **BatchListProgresses**
    - 响应参数变更
      - `+ results.incre_trans_delay_millis`
  - **ShowJobList**
    - 请求参数变更
      - `+ instance_ids`
      - `+ instance_ip`

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持接口`UploadJdbcDriver`、`ListJdbcDrivers`、`DeleteJdbcDriver`、`SyncJdbcDriver`
- _解决问题_
  - 无
- _特性变更_
  - **BatchCreateJobsAsync**
    - 请求参数变更
      - `+ jobs.policy_config.dml_types`
  - **ListAsyncJobDetail**
    - 响应参数变更
      - `+ jobs.connection_management`
      - `+ jobs.policy_config.dml_types`
  - **UpdateBatchAsyncJobs**
    - 请求参数变更
      - `+ jobs.params.policy_config.dml_types`
  - **ShowJobDetail**
    - 响应参数变更
      - `+ job.connection_management`
      - `+ job.policy_config.dml_types`
  - **UpdateJob**
    - 请求参数变更
      - `+ job.params.policy_config.dml_types`

### HuaweiCloud SDK ELB

- _新增特性_
  - 支持以下接口：
    - `BatchAddAvailableZones`
    - `BatchRemoveAvailableZones`
    - `ListMasterSlavePools`
    - `CreateMasterSlavePool`
    - `ShowMasterSlavePool`
    - `DeleteMasterSlavePool`
- _解决问题_
  - 无
- _特性变更_
  - **ShowLoadBalancer**
    - 响应参数变更
      - `+ loadbalancer.charge_mode`
      - `+ loadbalancer.log_group_id`
      - `+ loadbalancer.log_topic_id`
  - **UpdateLoadBalancer**
    - 请求参数变更
      - `+ loadbalancer.charge_mode`
      - `+ loadbalancer.ipv6_vip_address`
    - 响应参数变更
      - `+ loadbalancer.charge_mode`
      - `+ loadbalancer.log_group_id`
      - `+ loadbalancer.log_topic_id`
  - **ListListeners**
    - 响应参数变更
      - `+ listeners.ssl_early_data_enable`
  - **CreateListener**
    - 请求参数变更
      - `+ listener.ssl_early_data_enable`
    - 响应参数变更
      - `+ listener.ssl_early_data_enable`
  - **ShowListener**
    - 响应参数变更
      - `+ listener.ssl_early_data_enable`
  - **UpdateListener**
    - 请求参数变更
      - `+ listener.ssl_early_data_enable`
    - 响应参数变更
      - `+ listener.ssl_early_data_enable`
  - **ListLoadBalancers**
    - 请求参数变更
      - `+ log_topic_id`
      - `+ log_group_id`
    - 响应参数变更
      - `+ loadbalancers.charge_mode`
      - `+ loadbalancers.log_group_id`
      - `+ loadbalancers.log_topic_id`
  - **CreateLoadBalancer**
    - 请求参数变更
      - `+ loadbalancer.charge_mode`
      - `+ loadbalancer.ipv6_vip_address`
    - 响应参数变更
      - `+ loadbalancer.charge_mode`
      - `+ loadbalancer.log_group_id`
      - `+ loadbalancer.log_topic_id`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持以下接口：
    - `ListAppTemplates`
    - `ShowAppTemplate`
    - `ListFunctionApplications`
    - `CreateFunctionApp`
    - `ShowFunctionApp`
    - `DeleteFunctionApp`
    - `CreateCallbackWorkflow`
- _解决问题_
  - 无
- _特性变更_
  - **ImportFunction**
    - 响应参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
  - **ListFunctions**
    - 响应参数变更
      - `+ functions.pre_stop_handler`
      - `+ functions.pre_stop_timeout`
  - **CreateFunction**
    - 请求参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
    - 响应参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
  - **ShowFunctionConfig**
    - 响应参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
  - **UpdateFunctionConfig**
    - 请求参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
    - 响应参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
  - **UpdateFunctionMaxInstanceConfig**
    - 响应参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
  - **ListBridgeFunctions**
    - 响应参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
  - **ShowResInstanceInfo**
    - 响应参数变更
      - `+ resources.resource_detail.pre_stop_handler`
      - `+ resources.resource_detail.pre_stop_timeout`
  - **CreateFunctionVersion**
    - 响应参数变更
      - `+ pre_stop_handler`
      - `+ pre_stop_timeout`
  - **ListFunctionVersions**
    - 响应参数变更
      - `+ versions.pre_stop_handler`
      - `+ versions.pre_stop_timeout`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ShowRestoreTables`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持以下接口：
    - `ShowElbIpGroup`
    - `SwitchIpGroup`
    - `ListInstancesSession`
    - `DeleteInstancesSession`
    - `ListInstancesSessionStatistics`
    - `ResetParamGroupTemplate`
    - `ListRedisSlowLogs`
    - `ListCassandraSlowLogs`
    - `ListMongodbSlowLogs`
    - `ListLtsConfigs`
    - `SaveLtsConfigs`
    - `DeleteLtsConfigs`
    - `ListRestoreDatabases`
    - `ListRestoreTables`
    - `ListMongodbErrorLogs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`ShowKafkaUserClientQuota`、`UpdateKafkaUserClientQuotaTask`、`CreateKafkaUserClientQuotaTask`、`DeleteKafkaUserClientQuotaTask`
- _解决问题_
  - 无
- _特性变更_
  - **ListInstances**
    - 请求参数变更
      - `+ status: enum value [DELETING,ERROR,CREATEFAILED,FREEZING,EXTENDING,SHRINKING,EXTENDEDFAILED,CONFIGURING,UPGRADING,UPGRADINGFAILED,ROLLBACK,ROLLBACKFAILED,VOLUMETYPECHANGING]`
      - `+ status: enum value [FAULTY,RESIZING,RESIZING FAILED]`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeIdCard**
    - 请求参数变更
      - `+ return_portrait_location`
    - 响应参数变更
      - `+ result.portrait_location`

### HuaweiCloud SDK OMS

- _新增特性_
  - 支持以下接口：
    - `ListSyncTasks`
    - `CreateSyncTask`
    - `ShowSyncTask`
    - `DeleteSyncTask`
    - `ListSyncTaskStatistic`
    - `StopSyncTask`
    - `StartSyncTask`
    - `ShowBucketObjects`
    - `ShowCdnInfo`
    - `ShowCloudType`
    - `ShowRegionInfo`
    - `ShowBucketList`
    - `ShowBucketRegion`
    - `CheckPrefix`
- _解决问题_
  - 无
- _特性变更_
  - **ShowTask**
    - 响应参数变更
      - `+ source_cdn.authentication_type: enum value [TENCENT_COS_A,TENCENT_COS_B,TENCENT_COS_C,TENCENT_COS_D]`
  - **ShowTaskGroup**
    - 响应参数变更
      - `+ source_cdn.authentication_type: enum value [TENCENT_COS_A,TENCENT_COS_B,TENCENT_COS_C,TENCENT_COS_D]`
  - **CreateTask**
    - 请求参数变更
      - `+ source_cdn.authentication_type: enum value [TENCENT_COS_A,TENCENT_COS_B,TENCENT_COS_C,TENCENT_COS_D]`
  - **ListTasks**
    - 响应参数变更
      - `+ tasks.source_cdn.authentication_type: enum value [TENCENT_COS_A,TENCENT_COS_B,TENCENT_COS_C,TENCENT_COS_D]`
  - **CreateTaskGroup**
    - 请求参数变更
      - `+ source_cdn.authentication_type: enum value [TENCENT_COS_A,TENCENT_COS_B,TENCENT_COS_C,TENCENT_COS_D]`
  - **ListTaskGroup**
    - 响应参数变更
      - `+ taskgroups.source_cdn.authentication_type: enum value [TENCENT_COS_A,TENCENT_COS_B,TENCENT_COS_C,TENCENT_COS_D]`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstancesDetails**
    - 请求参数变更
      - `+ status: enum value [DELETING,FREEZING,EXTENDING,SHRINKING,EXTENDEDFAILED,CONFIGURING,UPGRADING,UPGRADINGFAILED,ROLLBACK,ROLLBACKFAILED,VOLUMETYPECHANGING]`
      - `+ status: enum value [STARTING,CLOSING]`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持以下接口：
    - `UpgradeDbMajorVersion`
    - `ShowAvailableVersion`
    - `UpgradeDbMajorVersionPreCheck`
    - `ListInspectionHistories`
    - `ListUpgradeHistories`
    - `ShowUpgradeDbMajorVersionStatus`
    - `UpdateTdeStatus`
    - `ShowTdeStatus`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstances**
    - 请求参数变更
      - `+ status: enum value [DELETING,ERROR,CREATEFAILED,FREEZING,EXTENDING,SHRINKING,EXTENDEDFAILED,CONFIGURING,UPGRADING,UPGRADINGFAILED,ROLLBACK,ROLLBACKFAILED,VOLUMETYPECHANGING]`
      - `+ status: enum value [FAULTY,RESIZING,RESIZING FAILED]`

### HuaweiCloud SDK SCM

- _新增特性_
  - 支持接口`DeployCertificate`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SecMaster

- _新增特性_
  - 支持以下接口：
    - `ListDataclass`
    - `ListDataclassFields`
    - `ListWorkflows`
    - `CreateDataspace`
    - `CreatePipe`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ServiceStage

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateInstance**
    - 请求参数变更
      - `+ configuration.container_spec`
  - **ChangeInstance**
    - 请求参数变更
      - `+ configuration.container_spec`

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateAssetByFileUpload**
    - 请求参数变更
      - `+ thumbnail.quantity`
      - `+ thumbnail.quantity_time`
      - `+ thumbnail.type: enum value [quantity]`
  - **PublishAssetFromObs**
    - 请求参数变更
      - `+ thumbnail.quantity`
      - `+ thumbnail.quantity_time`
      - `+ thumbnail.type: enum value [quantity]`
  - **CreateAssetProcessTask**
    - 请求参数变更
      - `+ thumbnail.quantity`
      - `+ thumbnail.quantity_time`
      - `+ thumbnail.type: enum value [quantity]`
  - **ListTopStatistics**
    - 响应参数变更
      - `+ top_urls.duration_ms`
  - **UploadMetaDataByUrl**
    - 请求参数变更
      - `+ upload_metadatas.thumbnail.quantity`
      - `+ upload_metadatas.thumbnail.quantity_time`
      - `+ upload_metadatas.thumbnail.type: enum value [quantity]`
  - **ListAssetList**
    - 响应参数变更
      - `+ assets.duration_ms`
  - **ShowTakeOverAssetDetails**
    - 响应参数变更
      - `+ base_info.meta_data.duration_ms`
  - **PublishAssets**
    - 响应参数变更
      - `+ asset_info_array.base_info.meta_data.duration_ms`
  - **UnpublishAssets**
    - 响应参数变更
      - `+ asset_info_array.base_info.meta_data.duration_ms`
  - **ShowAssetMeta**
    - 响应参数变更
      - `+ asset_info_array.base_info.meta_data.duration_ms`
  - **ShowAssetDetail**
    - 响应参数变更
      - `+ base_info.meta_data.duration_ms`
      - `+ thumbnail_info.quantity`
  - **ShowTakeOverTaskDetails**
    - 响应参数变更
      - `+ assets.base_info.meta_data.duration_ms`

# 0.1.67 2023-11-16

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDomainFullConfig**
    - 响应参数变更
      - `+ configs.business_type`
      - `+ configs.service_area`
      - `+ configs.remark`
      - `+ configs.flexible_origin`
      - `+ configs.slice_etag_status`
      - `+ configs.origin_receive_timeout`
      - `+ configs.remote_auth`
      - `+ configs.websocket`
      - `+ configs.video_seek`
      - `+ configs.request_limit_rules`
      - `+ configs.ip_frequency_limit`
      - `+ configs.hsts`
      - `+ configs.quic`
      - `+ configs.url_auth.sign_method`
      - `+ configs.url_auth.match_type`
      - `+ configs.url_auth.inherit_config`
      - `+ configs.url_auth.key`
      - `+ configs.url_auth.backup_key`
      - `+ configs.url_auth.sign_arg`
      - `+ configs.https.expire_time`
      - `+ configs.https.certificate_type`
      - `+ configs.https.ocsp_stapling_status`
      - `+ configs.sources.weight`
      - `+ configs.sources.obs_bucket_type`
      - `+ configs.sources.bucket_access_key`
      - `+ configs.sources.bucket_secret_key`
      - `+ configs.sources.bucket_region`
      - `+ configs.sources.bucket_name`
      - `+ configs.compress.file_type`
      - `+ configs.user_agent_filter.ua_list`
  - **UpdateDomainFullConfig**
    - 请求参数变更
      - `+ configs.sources.weight`
      - `+ configs.sources.obs_bucket_type`
      - `+ configs.sources.bucket_access_key`
      - `+ configs.sources.bucket_secret_key`
      - `+ configs.sources.bucket_region`
      - `+ configs.sources.bucket_name`
      - `+ configs.compress.file_type`
      - `+ configs.user_agent_filter.ua_list`

### HuaweiCloud SDK CodeArtsBuild

- _新增特性_
  - 支持以下接口：
    - `DownloadBuildLog`
    - `DownloadTaskLog`
    - `ShowFlowGraph`
    - `ShowRecordDetail`
    - `ShowOutputInfo`
    - `StopJob`
    - `CreateBuildJob`
    - `UpdateBuildJob`
    - `ListTemplates`
    - `CreateTemplates`
    - `DeleteTemplates`
    - `ListNotice`
    - `UpdateNotice`
    - `DisableNotice`
    - `ListJobConfig`
- _解决问题_
  - 无
- _特性变更_
  - **ShowRecordInfo**
    - 响应参数变更
      - `+ result.duration`

### HuaweiCloud SDK CodeArtsPipeline

- _新增特性_
  - 支持以下接口：
    - `CreatePipelineTemplate`
    - `ShowProjectOpenSourceStrategy`
    - `ListProjectStrategy`
    - `ListProjectOpenSourceStrategy`
    - `ShowProjectStrategy`
    - `ShowRule`
    - `ListRule`
    - `UpdateRule`
    - `DeleteRule`
    - `CreateRule`
    - `CreateStrategy`
    - `UpdateStrategy`
    - `ShowStrategy`
    - `ListStrategy`
    - `DeleteStrategy`
    - `SwitchStrategy`
    - `CreateOpenSourceStrategy`
    - `UpdateOpenSourceStrategy`
    - `ShowOpenSourceStrategy`
    - `ListOpenSourceStrategy`
    - `DeleteOpenSourceStrategy`
    - `SwitchOpenSourceStrategy`
    - `CreatePipelineGroup`
    - `UpdatePipelineGroup`
    - `DeletePipelineGroup`
    - `ShowPipelineGroupTree`
    - `BatchMovePipelineToGroup`
    - `PublishPlugin`
    - `PublishPluginBind`
    - `UploadPluginIcon`
    - `UploadPublisherIcon`
    - `DeletePluginDraft`
    - `ListPublisher`
    - `ListAvailablePublisher`
    - `ListStagePlugins`
    - `ListBasePlugins`
    - `ListBasePluginsNewPost`
    - `ListPlugins`
    - `ShowPluginMetrics`
    - `ShowPluginInputs`
    - `ShowPluginOutputs`
    - `ListPLuginVersion`
    - `ShowPluginVersion`
    - `ListPluginVersionNumber`
    - `DeletePublisher`
    - `ShowPublisher`
    - `CreateBasicPlugin`
    - `UpdateBasicPlugin`
    - `DeleteBasicPlugin`
    - `UploadBasicPlugin`
    - `ShowBasicPlugin`
    - `UpdatePipelineTemplate`
    - `DeletePipelineTemplate`
- _解决问题_
  - 无
- _特性变更_
  - **ListPipelines**
    - 响应参数变更
      - `+ pipelines.project_name`
  - **CreatePipelineNew**
    - 请求参数变更
      - `+ variables`
      - `+ schedules`
      - `+ triggers`
      - `+ manifest_version`
      - `+ definition`
      - `+ project_name`

### HuaweiCloud SDK CSE

- _新增特性_
  - 支持以下接口：
    - `ListGovernancePolicy`
    - `CreateGovernancePolicy`
    - `ListGovernancePolicyByPolicyId`
    - `UpdateGovernancePolicy`
    - `DeleteGovernancePolicy`
    - `ListMicroserviceRouteRule`
    - `CreateMicroserviceRouteRule`
    - `DeleteMicroserviceRouteRule`
    - `ListGovernancePolicys`
    - `UpgradeEngineConfig`
    - `ResizeEngine`
    - `ListNacosNamespaces`
    - `UpdateNacosNamespaces`
    - `CreateNacosNamespaces`
    - `DeleteNacosNamespaces`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CSMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSecretTags**
    - 响应参数变更
      - `+ sys_tags.value`
      - `- sys_tags.values`
  - **ListNotificationRecords**
    - 请求参数变更
      - `+ limit`
      - `+ marker`

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 支持接口`UpdateFactoryJobName`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateVifPeer**
    - 响应参数变更
      - `+ vif_peer.receive_route_num`
  - **CreateVifPeer**
    - 响应参数变更
      - `+ vif_peer.receive_route_num`
  - **ShowVirtualInterface**
    - 响应参数变更
      - `+ virtual_interface.vif_peers.receive_route_num`
  - **UpdateVirtualInterface**
    - 响应参数变更
      - `+ virtual_interface.vif_peers.receive_route_num`
  - **ListVirtualInterfaces**
    - 响应参数变更
      - `+ virtual_interfaces.vif_peers.receive_route_num`
  - **CreateVirtualInterface**
    - 响应参数变更
      - `+ virtual_interface.vif_peers.receive_route_num`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowScript**
    - 响应参数变更
      - `+ description`
      - `+ approvers`
      - `+ targetStatus`
      - `+ type: enum value [RDSSQL,PRESTO,ClickHouseSQL,HetuEngineSQL,PYTHON,ImpalaSQL,SparkPython]`
  - **UpdateScript**
    - 请求参数变更
      - `+ description`
      - `+ approvers`
      - `+ targetStatus`
      - `+ type: enum value [RDSSQL,PRESTO,ClickHouseSQL,HetuEngineSQL,PYTHON,ImpalaSQL,SparkPython]`
  - **CreateScript**
    - 请求参数变更
      - `+ description`
      - `+ approvers`
      - `+ targetStatus`
      - `+ type: enum value [RDSSQL,PRESTO,ClickHouseSQL,HetuEngineSQL,PYTHON,ImpalaSQL,SparkPython]`
  - **ListScripts**
    - 响应参数变更
      - `+ description`
      - `+ approvers`
      - `+ targetStatus`
      - `+ scripts.description`
      - `+ scripts.targetStatus`
      - `+ scripts.approvers`
      - `+ scripts.type: enum value [RDSSQL,PRESTO,ClickHouseSQL,HetuEngineSQL,PYTHON,ImpalaSQL,SparkPython]`

### HuaweiCloud SDK DLI

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowSqlJobStatus**
    - 响应参数变更
      - `+ result_format`
      - `+ result_path`

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreatePrivateZone**
    - 请求参数变更
      - `+ router.status`
  - **AssociateRouter**
    - 请求参数变更
      - `+ router.status`
  - **DisassociateRouter**
    - 请求参数变更
      - `+ router.status`
  - **ShowPrivateZone**
    - 响应参数变更
      - `+ routers.status`

### HuaweiCloud SDK EdgeSec

- _新增特性_
  - 支持以下接口：
    - `ListCertificates`
    - `CreateCertificate`
    - `ShowCertificate`
    - `UpdateCertificate`
    - `DeleteCertificate`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持以下接口：
    - `ListFunctionTags`
    - `ListBridgeFunctions`
    - `ListBridgeVersions`
    - `UpdateFunctionCollectState`
    - `ListFunctionTemplate`
    - `ShowFunctionTemplate`
    - `ShowFuncReservedInstanceMetrics`
    - `ShowFunctionMetrics`
    - `EnableAsyncStatusLog`
    - `ShowProjectAsyncStatusLogInfo`
- _解决问题_
  - 无
- _特性变更_
  - **ListFunctions**
    - 响应参数变更
      - `+ functions.resource_id`
  - **ShowFunctionConfig**
    - 响应参数变更
      - `+ func_id`
      - `+ resource_id`
  - **UpdateFunctionConfig**
    - 响应参数变更
      - `+ func_id`
      - `+ resource_id`
  - **ShowResInstanceInfo**
    - 响应参数变更
      - `+ resources.resource_detail.resource_id`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateGaussMySqlInstance**
    - 响应参数变更
      - `* instance.volume.size: string -> int32`

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListFlowBySimCards**
    - 请求参数变更
      - `+ sim_card_ids`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowRuleAction**
    - 响应参数变更
      - `+ channel_detail.dms_kafka_forwarding.security_protocol`
  - **UpdateRuleAction**
    - 请求参数变更
      - `+ channel_detail.dms_kafka_forwarding.security_protocol`
    - 响应参数变更
      - `+ channel_detail.dms_kafka_forwarding.security_protocol`
  - **CreateRuleAction**
    - 请求参数变更
      - `+ channel_detail.dms_kafka_forwarding.security_protocol`
    - 响应参数变更
      - `+ channel_detail.dms_kafka_forwarding.security_protocol`
  - **ListRuleActions**
    - 响应参数变更
      - `+ actions.channel_detail.dms_kafka_forwarding.security_protocol`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstanceConsumerGroups**
    - 响应参数变更
      - `* groups.createdAt: int32 -> int64`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListTopnTrafficStatistics**
    - 响应参数变更
      - `+ results.cold_storage`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`AddComponent`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizePeruIdCard`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowRabbitMqProductCores`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`RevokePostgresqlDbPrivilege`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ServiceStage

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ChangeInstance**
    - 请求参数变更
      - `+ configuration.env`
      - `+ configuration.storage`
      - `+ configuration.strategy`
      - `+ configuration.lifecycle`
      - `+ configuration.scheduler`
      - `+ configuration.probes`
      - `* configuration: map<string, object> -> object<InstanceConfiguration>`

### HuaweiCloud SDK Workspace

- _新增特性_
  - 支持接口`BatchAddDesktopsTags`、`BatchDeleteDesktopsTags`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.66 2023-11-13

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ShowNodesInformation`、`DeleteCenterTask`、`DeleteDiagnosisTask`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstanceTopics**
    - 请求参数变更
      - `+ offset`
      - `+ limit`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **Createfavorite**
    - 请求参数变更
      - `+ is_global`
    - 响应参数变更
      - `+ is_global`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 支持接口`ShowRabbitMqProductCores`、`ShowCesHierarchy`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListPostgresqlListHistoryTables`、`ListHistoryDatabase`、`BatchRestorePostgreSqlTables`、`BatchRestoreDatabase`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListInstancesResourceMetrics`、`ListInstancesRecommendation`

# 0.1.65 2023-11-09

### HuaweiCloud SDK VPN

- _新增特性_
  - 支持虚拟专用网络服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ASM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowMesh**
    - 响应参数变更
      - `- spec.region`
      - `- spec.extendParams.clusters.region`
  - **DeleteMesh**
    - 响应参数变更
      - `- spec.region`
      - `- spec.extendParams.clusters.region`
  - **CreateMesh**
    - 请求参数变更
      - `- spec.region`
      - `- spec.extendParams.clusters.region`
    - 响应参数变更
      - `- spec.region`
      - `- spec.extendParams.clusters.region`
  - **ListMeshes**
    - 响应参数变更
      - `- items.spec.region`
      - `- items.spec.extendParams.clusters.region`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAddonInstance**
    - 响应参数变更
      - `+ status.status: enum value [unknown]`
  - **UpdateAddonInstance**
    - 响应参数变更
      - `+ status.status: enum value [unknown]`
  - **RollbackAddonInstance**
    - 响应参数变更
      - `+ status.status: enum value [unknown]`
  - **ShowCluster**
    - 响应参数变更
      - `+ spec.serviceNetwork`
  - **UpdateCluster**
    - 响应参数变更
      - `+ spec.serviceNetwork`
  - **DeleteCluster**
    - 响应参数变更
      - `+ spec.serviceNetwork`
  - **CreateAddonInstance**
    - 响应参数变更
      - `+ status.status: enum value [unknown]`
  - **ListAddonInstances**
    - 响应参数变更
      - `+ items.status.status: enum value [unknown]`
  - **CreateCluster**
    - 请求参数变更
      - `+ spec.serviceNetwork`
    - 响应参数变更
      - `+ spec.serviceNetwork`
  - **ListClusters**
    - 响应参数变更
      - `+ items.spec.serviceNetwork`
  - **ShowNode**
    - 响应参数变更
      - `- status.phase: enum value [Installed,ShutDown]`
  - **UpdateNode**
    - 响应参数变更
      - `- status.phase: enum value [Installed,ShutDown]`
  - **DeleteNode**
    - 响应参数变更
      - `- status.phase: enum value [Installed,ShutDown]`
  - **CreateNode**
    - 响应参数变更
      - `- status.phase: enum value [Installed,ShutDown]`
  - **ListNodes**
    - 响应参数变更
      - `- items.status.phase: enum value [Installed,ShutDown]`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CheckMigrationConnectivity`
  - **ListBackupRecords**
    - 响应参数变更
      - `+ backup_record_response.backup_format`
      - `+ backup_record_response.execution_at`

### HuaweiCloud SDK DLI

- _新增特性_
  - 支持接口`ListQueueProperty`、`UpdateQueueProperty`、`CreateQueueProperty`、`DeleteQueueProperty`
- _解决问题_
  - 无
- _特性变更_
  - **ShowSqlJobStatus**
    - 响应参数变更
      - `+ user_conf`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowNextflowJob**
    - 响应参数变更
      - `+ priority`
  - **ListDrugJob**
    - 响应参数变更
      - `- jobs.priority`
  - **ShowSynthesisJob**
    - 响应参数变更
      - `- basic_info.priority`
  - **ShowFepJob**
    - 响应参数变更
      - `- basic_info.priority`
  - **ShowPocketDetectionJob**
    - 响应参数变更
      - `- basic_info.priority`
  - **ShowAdmetJob**
    - 响应参数变更
      - `- basic_info.priority`
      - `- models.value_range.lower_inclusive`
      - `- models.value_range.upper_inclusive`
      - `* models.value_range.lower: number -> float`
      - `* models.value_range.upper: number -> float`
      - `* models.value_range: object<ValueRange> -> object<ValueRange2>`
  - **ShowPocketMolDesignJob**
    - 响应参数变更
      - `- basic_info.priority`
      - `- model_list.value_range.lower_inclusive`
      - `- model_list.value_range.upper_inclusive`
      - `* model_list.value_range.lower: number -> float`
      - `* model_list.value_range.upper: number -> float`
      - `* model_list.value_range: object<ValueRange> -> object<ValueRange2>`
  - **ShowOptmJob**
    - 响应参数变更
      - `- basic_info.priority`
      - `- models.value_range.lower_inclusive`
      - `- models.value_range.upper_inclusive`
      - `* models.value_range.lower: number -> float`
      - `* models.value_range.upper: number -> float`
      - `* models.value_range: object<ValueRange> -> object<ValueRange2>`
  - **ShowDockingJob**
    - 响应参数变更
      - `- basic_info.priority`
  - **ListDrugModel**
    - 响应参数变更
      - `- models.value_range.lower_inclusive`
      - `- models.value_range.upper_inclusive`
      - `* models.value_range.lower: number -> float`
      - `* models.value_range.upper: number -> float`
      - `* models.value_range: object<ValueRange> -> object<ValueRange2>`

### HuaweiCloud SDK GES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListGraphs2**
    - 响应参数变更
      - `+ graphs.origin_graph_size_type_index`
      - `+ graphs.expand_time`
      - `+ graphs.resize_time`
      - `+ graphs.enable_multi_label`
  - **CreateGraph2**
    - 请求参数变更
      - `+ graph.enable_multi_label`
  - **ShowGraph2**
    - 响应参数变更
      - `+ graph.origin_graph_size_type_index`
      - `+ graph.expand_time`
      - `+ graph.resize_time`
      - `+ graph.enable_multi_label`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`ShowInstanceConfigs`、`ModifyInstanceConfigs`
- _解决问题_
  - 无
- _特性变更_
  - **BatchRestartOrDeleteInstances**
    - 请求参数变更
      - `+ all_failure`
      - `- allFailure`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchRestartOrDeleteInstances**
    - 请求参数变更
      - `+ all_failure`
      - `- allFailure`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchDeleteInstances**
    - 请求参数变更
      - `+ all_failure`
      - `- allFailure`
  - **DeleteRocketMqMigrationTask**
    - 请求参数变更
      - `+ task_ids`
      - `- taskIds`

### HuaweiCloud SDK SCM

- _新增特性_
  - 支持接口`BatchPushCertificate`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Workspace

- _新增特性_
  - 支持以下接口：
    - `BatchRebuildDesktopsSystemDisk`
    - `ShowDesktopNetwork`
    - `ChangeDesktopNetwork`
    - `ShowTagByDesktopId`
    - `CreateTag`
    - `DeleteTag`
    - `ListProjectTags`
    - `BatchChangeTags`
    - `ListDesktopByTags`
- _解决问题_
  - 无
- _特性变更_
  - **BatchDeleteDesktops**
    - 请求参数变更
      - `+ is_force_delete`
  - **ListDesktops**
    - 请求参数变更
      - `+ enterprise_project_id`
      - `+ desktop_type`
    - 响应参数变更
      - `+ desktops.attach_user_infos`
      - `+ desktops.enterprise_project_id`
      - `+ desktops.in_maintenance_mode`
  - **CreateDesktop**
    - 请求参数变更
      - `+ desktop_name`
      - `+ size`
      - `+ enterprise_project_id`
      - `+ desktop_type: enum value [SHARED]`
      - `+ desktops.user_phone`
  - **ApplyDesktopsInternet**
    - 请求参数变更
      - `+ enterprise_project_id`
  - **ListDesktopsEips**
    - 请求参数变更
      - `+ enterprise_project_id`
    - 响应参数变更
      - `+ eips.enterprise_project_id`
  - **ListUsersOfGroup**
    - 请求参数变更
      - `+ description`
      - `+ active_type`
    - 响应参数变更
      - `+ users.description`
  - **ListProducts**
    - 响应参数变更
      - `+ products.data_disk_size`
      - `+ products.default_desktop_num`
      - `+ products.max_apply_desktop_num`
  - **ListUsers**
    - 请求参数变更
      - `+ group_name`
  - **ListItaSubJobs**
    - 请求参数变更
      - `+ desktop_pool_id`
    - 响应参数变更
      - `+ jobs.desktop_name`
      - `+ jobs.ip_address`
      - `+ jobs.mac_address`
  - **ListWorkspaces**
    - 响应参数变更
      - `+ dc_vnc_ip`
  - **UpdateWorkspace**
    - 请求参数变更
      - `+ dc_vnc_ip`
    - 响应参数变更
      - `+ dc_vnc_ip`
  - **DeleteDesktop**
    - 请求参数变更
      - `+ is_force_delete`
  - **ShowDesktopDetail**
    - 响应参数变更
      - `+ desktop.user_list`
      - `+ desktop.user_group_list`
      - `+ desktop.attach_user_infos`
      - `+ desktop.attach_state`
      - `+ desktop.enterprise_project_id`
  - **ListDesktopsDetail**
    - 请求参数变更
      - `+ user_names`
      - `+ sort_field`
      - `+ sort_type`
      - `+ user_attached`
      - `+ enterprise_project_id`
      - `+ image_id`
      - `+ charge_mode`
      - `+ in_maintenance_mode`
      - `* desktop_id: string -> list<string>`
    - 响应参数变更
      - `+ desktops.user_list`
      - `+ desktops.user_group_list`
      - `+ desktops.attach_user_infos`
      - `+ desktops.attach_state`
      - `+ desktops.enterprise_project_id`

# 0.1.64 2023-11-02

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListApisV2**
    - 请求参数变更
      - `+ vpc_channel_name`

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListPostalAddress`、`UpdatePostal`、`CreatePostal`、`DeletePostal`
  - **ListCustomerselfResourceRecordDetails**
    - 响应参数变更
      - `+ monthly_records.id`

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListCustomerselfResourceRecordDetails**
    - 响应参数变更
      - `+ monthly_records.id`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`StartChat`、`SyncChat`、`SyncGetChatResult`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CSMS

- _新增特性_
  - 支持接口`RotateSecret`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持以下接口：
    - `ShowConfigHistoryDetail`
    - `UpdateClientIpTransparentTransmission`
    - `UpdateInstanceConfig`
    - `ListInstanceOperations`
    - `ExportInstancesTask`
    - `ExportExcelJob`
    - `CreateResizeOrder`
    - `ShowExpireAutoScanConfig`
    - `UpdateExpireAutoScanConfig`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowNodesInformation`、`ShowBackUpInfo`、`CreateOrUpdateBackUpInfo`、`CreateConnectivityTest`

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowPrivateZone**
    - 响应参数变更
      - `+ enterprise_project_id`
      - `+ proxy_pattern`

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持接口`CollectPositionAsync`、`ShowPositionResult`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持以下接口：
    - `ExecuteClusterUpgradeAction`
    - `ListUpdatableVersion`
    - `ListUpdateRecord`
    - `CheckTableRestore`
    - `RestoreTable`
    - `ListSnapshotCrossRegion`
    - `ListSnapshotCrossRegionPolicy`
    - `AddSnapshotCrossRegionPolicy`
    - `DeleteSnapshotCrossRegionPolicy`
    - `StopWorkloadPlan`
    - `ShowWorkloadPlan`
    - `DeleteWorkloadPlan`
    - `StartWorkloadPlan`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持以下接口：
    - `ListNodes`
    - `ListScalingHistory`
    - `ListPolicyEvents`
    - `CreatePocketDetectionJob`
    - `ShowPocketDetectionJob`
    - `CreateAdmetJob`
    - `ShowAdmetJob`
    - `CreatePocketMolDesignJob`
    - `ShowPocketMolDesignJob`
    - `ListDrugModel`
    - `CreateDrugModel`
    - `UpdateDrugModel`
    - `DeleteDrugModel`
- _解决问题_
  - 无
- _特性变更_
  - **CreateDrugLigandSvg**
    - 请求参数变更
      - `+ scaffold`
  - **CreateNextflowJob**
    - 请求参数变更
      - `+ priority`
  - **CreateDrugLigandSimilarityGraphTask**
    - 请求参数变更
      - `+ mode: enum value [FREE]`
  - **ListDrugJob**
    - 响应参数变更
      - `+ jobs.priority`
  - **ShowSynthesisJob**
    - 响应参数变更
      - `+ basic_info.priority`
  - **ShowFepJob**
    - 响应参数变更
      - `+ part_failed_reason`
      - `+ basic_info.priority`
  - **ParseDrugReceptorInfo**
    - 响应参数变更
      - `* ligands.bounding_box: object<BoundingBoxDto> -> object<DrugBoundingBoxDto>`
  - **RecognizeDrugReceptorPocket**
    - 响应参数变更
      - `* pockets: list<BoundingBoxDto> -> list<DrugBoundingBoxDto>`
  - **CreateOptmJob**
    - 请求参数变更
      - `+ molecule_file`
      - `+ sampler_mixin_weight`
      - `+ model_ids`
      - `+ strong_constraints.id`
      - `+ binding_site.bounding_box.padding`
      - `* body: object<DrugFile> -> object<ReceptorDrugFile>`
      - `+ weak_constraints.id`
  - **ShowOptmJob**
    - 响应参数变更
      - `+ models`
      - `+ sampler_mixin_weight`
      - `+ molecule_file`
      - `+ basic_info.priority`
      - `+ strong_constraints.id`
      - `+ binding_site.bounding_box.padding`
      - `* body: object<DrugFile> -> object<ReceptorDrugFile>`
      - `+ weak_constraints.id`
  - **CreateDockingJob**
    - 请求参数变更
      - `+ receptors.bounding_box.padding`
      - `* body: object<DrugFile> -> object<ReceptorDrugFile>`
  - **ShowDockingJob**
    - 响应参数变更
      - `+ part_failed_reason`
      - `+ basic_info.priority`
      - `+ receptors.bounding_box.padding`
      - `* body: object<DrugFile> -> object<ReceptorDrugFile>`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListListeners**
    - 请求参数变更
      - `+ proxy_protocol_enable`
    - 响应参数变更
      - `+ listeners.proxy_protocol_enable`
      - `+ listeners.insert_headers.X-Forwarded-Proto`
      - `+ listeners.insert_headers.X-Real-IP`
      - `+ listeners.insert_headers.X-Forwarded-ELB-ID`
      - `+ listeners.insert_headers.X-Forwarded-TLS-Certificate-ID`
      - `+ listeners.insert_headers.X-Forwarded-TLS-Protocol`
      - `+ listeners.insert_headers.X-Forwarded-TLS-Cipher`
  - **CreateListener**
    - 请求参数变更
      - `+ listener.proxy_protocol_enable`
      - `+ listener.insert_headers.X-Forwarded-Proto`
      - `+ listener.insert_headers.X-Real-IP`
      - `+ listener.insert_headers.X-Forwarded-ELB-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Certificate-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Protocol`
      - `+ listener.insert_headers.X-Forwarded-TLS-Cipher`
    - 响应参数变更
      - `+ listener.proxy_protocol_enable`
      - `+ listener.insert_headers.X-Forwarded-Proto`
      - `+ listener.insert_headers.X-Real-IP`
      - `+ listener.insert_headers.X-Forwarded-ELB-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Certificate-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Protocol`
      - `+ listener.insert_headers.X-Forwarded-TLS-Cipher`
  - **ShowListener**
    - 响应参数变更
      - `+ listener.proxy_protocol_enable`
      - `+ listener.insert_headers.X-Forwarded-Proto`
      - `+ listener.insert_headers.X-Real-IP`
      - `+ listener.insert_headers.X-Forwarded-ELB-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Certificate-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Protocol`
      - `+ listener.insert_headers.X-Forwarded-TLS-Cipher`
  - **UpdateListener**
    - 请求参数变更
      - `+ listener.proxy_protocol_enable`
      - `+ listener.insert_headers.X-Forwarded-Proto`
      - `+ listener.insert_headers.X-Real-IP`
      - `+ listener.insert_headers.X-Forwarded-ELB-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Certificate-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Protocol`
      - `+ listener.insert_headers.X-Forwarded-TLS-Cipher`
    - 响应参数变更
      - `+ listener.proxy_protocol_enable`
      - `+ listener.insert_headers.X-Forwarded-Proto`
      - `+ listener.insert_headers.X-Real-IP`
      - `+ listener.insert_headers.X-Forwarded-ELB-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Certificate-ID`
      - `+ listener.insert_headers.X-Forwarded-TLS-Protocol`
      - `+ listener.insert_headers.X-Forwarded-TLS-Cipher`
  - **CreatePool**
    - 请求参数变更
      - `+ pool.ip_version`
  - **UpdatePool**
    - 请求参数变更
      - `+ pool.any_port_enable`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持以下接口：
    - `ShowIntelligentDiagnosisAbnormalCountOfInstances`
    - `ShowIntelligentDiagnosisInstanceInfosPerMetric`
    - `ShrinkGaussMySqlProxy`
    - `ShowInstanceDatabaseVersion`
    - `CopyInstanceConfigurations`
    - `ShowAutoScalingPolicy`
    - `UpdateAutoScalingPolicy`
    - `CheckResource`
    - `UpdateInstanceConfigurations`
- _解决问题_
  - 无
- _特性变更_
  - **DeleteGaussMySqlBackup**
    - 响应参数变更
      - `+ backup_id`
      - `+ backup_name`
      - `- job_id`
  - **CreateGaussMySqlProxy**
    - 请求参数变更
      - `+ subnet_id`
  - **ShowGaussMySqlBackupList**
    - 请求参数变更
      - `+ name`
      - `+ instance_name`
    - 响应参数变更
      - `+ backups.instance_name`
      - `- backups.status: enum value [BUILDING,COMPLETED,FAILED,AVAILABLE]`
      - `- backups.type: enum value [auto,manual]`
  - **ShowGaussMySqlProxyList**
    - 响应参数变更
      - `+ proxy_list.proxy.subnet_id`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持接口`DeleteDatabase`
- _解决问题_
  - 无
- _特性变更_
  - **ListInstances**
    - 请求参数变更
      - `+ charge_mode`
  - **ListInstancesDetails**
    - 请求参数变更
      - `+ charge_mode`

### HuaweiCloud SDK HSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSwrImageRepository**
    - 响应参数变更
      - `+ data_list.scannable`
      - `- data_list.scanable`

### HuaweiCloud SDK KPS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ImportPrivateKey**
    - 响应参数变更
      - `+ keypair.user_id`
      - `+ keypair.key_protection`
      - `* keypair: object<KeypairBean> -> object<ImportPrivateKeyAction>`

### HuaweiCloud SDK NAT

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListNatGatewayDnatRules**
    - 响应参数变更
      - `+ dnat_rules.global_eip_id`
      - `+ dnat_rules.global_eip_address`
  - **CreateNatGatewayDnatRule**
    - 请求参数变更
      - `+ dnat_rule.global_eip_id`
    - 响应参数变更
      - `+ dnat_rule.global_eip_id`
      - `+ dnat_rule.global_eip_address`
  - **ShowNatGatewayDnatRule**
    - 响应参数变更
      - `+ dnat_rule.global_eip_id`
      - `+ dnat_rule.global_eip_address`
  - **UpdateNatGatewayDnatRule**
    - 请求参数变更
      - `+ dnat_rule.global_eip_id`
    - 响应参数变更
      - `+ dnat_rule.global_eip_id`
      - `+ dnat_rule.global_eip_address`
  - **BatchCreateNatGatewayDnatRules**
    - 请求参数变更
      - `+ dnat_rules.global_eip_id`
    - 响应参数变更
      - `+ dnat_rules.global_eip_id`
      - `+ dnat_rules.global_eip_address`
  - **ListNatGatewaySnatRules**
    - 响应参数变更
      - `+ snat_rules.global_eip_id`
      - `+ snat_rules.global_eip_address`
  - **CreateNatGatewaySnatRule**
    - 请求参数变更
      - `+ snat_rule.global_eip_id`
    - 响应参数变更
      - `+ snat_rule.global_eip_id`
      - `+ snat_rule.global_eip_address`
  - **ShowNatGatewaySnatRule**
    - 响应参数变更
      - `+ snat_rule.global_eip_id`
      - `+ snat_rule.global_eip_address`
  - **UpdateNatGatewaySnatRule**
    - 请求参数变更
      - `+ snat_rule.global_eip_id`
    - 响应参数变更
      - `+ snat_rule.global_eip_address`
      - `+ snat_rule.global_eip_id`
  - **ListNatGateways**
    - 响应参数变更
      - `+ nat_gateways.ngport_ip_address`
      - `+ nat_gateways.billing_info`
      - `+ nat_gateways.dnat_rules_limit`
      - `+ nat_gateways.snat_rule_public_ip_limit`
  - **CreateNatGateway**
    - 请求参数变更
      - `+ nat_gateway.ngport_ip_address`
    - 响应参数变更
      - `+ nat_gateway.ngport_ip_address`
      - `+ nat_gateway.billing_info`
      - `+ nat_gateway.dnat_rules_limit`
      - `+ nat_gateway.snat_rule_public_ip_limit`
  - **ShowNatGateway**
    - 响应参数变更
      - `+ nat_gateway.ngport_ip_address`
      - `+ nat_gateway.billing_info`
      - `+ nat_gateway.dnat_rules_limit`
      - `+ nat_gateway.snat_rule_public_ip_limit`
  - **UpdateNatGateway**
    - 响应参数变更
      - `+ nat_gateway.ngport_ip_address`
      - `+ nat_gateway.billing_info`
      - `+ nat_gateway.dnat_rules_limit`
      - `+ nat_gateway.snat_rule_public_ip_limit`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateRocketMqMigrationTask**
    - 请求参数变更
      - `* body: string -> map<string, object>`

### HuaweiCloud SDK SecMaster

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListIncidentTypes`
  - **ListAlertRuleMetrics**
    - 响应参数变更
      - `+ metric`
      - `+ category`
      - `- body`
  - **CreateBatchOrderAlerts**
    - 请求参数变更
      - `- incident_id`
      - `- event_content`
      - `- marked_evidence`
      - `- incident_content.type_category`
      - `- incident_content.evidence_list`
      - `- incident_content.note_list`
      - `- incident_content.attachment_list`
      - `- incident_content.description`
      - `- incident_content.incident_type.layoutId`
    - 响应参数变更
      - `* data: object<BatchOrderAlertResult> -> object<BatchOperateAlertResult>`
  - **ShowAlertRule**
    - 响应参数变更
      - `- accumulated_times`
      - `- query_type: enum value [CBSL]`
  - **UpdateAlertRule**
    - 请求参数变更
      - `- accumulated_times`
      - `- query_type: enum value [CBSL]`
    - 响应参数变更
      - `- accumulated_times`
      - `- query_type: enum value [CBSL]`
  - **CreateAlertRuleSimulation**
    - 请求参数变更
      - `- query_type: enum value [CBSL]`
  - **ShowAlertRuleTemplate**
    - 响应参数变更
      - `- query_type: enum value [CBSL]`
  - **ListPlaybooks**
    - 请求参数变更
      - `- component_id`
      - `* offset: optional -> required`
      - `* limit: optional -> required`
  - **CreatePlaybook**
    - 请求参数变更
      - `- approve_role`
      - `- user_role`
      - `- edit_role`
      - `- owner_id`
  - **ListPlaybookActions**
    - 请求参数变更
      - `* limit: optional -> required`
      - `* offset: optional -> required`
    - 响应参数变更
      - `- data.sort_order`
  - **CreatePlaybookAction**
    - 响应参数变更
      - `- data.sort_order`
  - **UpdatePlaybookAction**
    - 响应参数变更
      - `- data.sort_order`
  - **ListDataobjectRelations**
    - 请求参数变更
      - `+ order: enum value [DESC,ASC]`
      - `* condition.conditions.data: list<string> -> list<object>`
      - `* condition.conditions: list<ConditonInfo> -> list<object>`
      - `* condition.logics: list<string> -> list<object>`
    - 响应参数变更
      - `+ success`
      - `+ data.data_object`
      - `+ data.dataclass_ref`
      - `+ data.format_version`
      - `+ data.version`
      - `+ data.workspace_id`
      - `- data.dataclass_id`
      - `- data.name`
      - `- data.type`
      - `- data.content`
      - `* data: list<DataobjectInfo> -> list<DataObjectDetail>`
  - **CreateDataobjectRelations**
    - 请求参数变更
      - `* body: object<CreateRelation> -> object<CreateDataobjectRelationsRequestBody>`
  - **DeleteDataobjectRelations**
    - 请求参数变更
      - `* body: object<CreateRelation> -> object<CreateDataobjectRelationsRequestBody>`
    - 响应参数变更
      - `- total`
      - `- offset`
      - `- success`
      - `- limit`
      - `- request_id`
      - `* data: object<DataResponse> -> object<BatchOperateDataobjectResult>`
  - **ListAlerts**
    - 请求参数变更
      - `+ order: enum value [DESC,ASC]`
      - `* condition.conditions.data: list<string> -> list<object>`
      - `* condition.conditions: list<ConditonInfo> -> list<object>`
      - `* condition.logics: list<string> -> list<object>`
    - 响应参数变更
      - `* data.data_object.network_list.src_port: string -> int32`
      - `* data.data_object.sla: int32 -> string`
      - `* data.data_object.simulation: boolean -> string`
      - `* data.data_object.process.process_pid: string -> int32`
      - `* data.data_object.process.process_uid: string -> int32`
  - **DeleteAlert**
    - 请求参数变更
      - `+ batch_ids`
      - `- ids`
      - `* body: object<DeleteAlert> -> object<DeleteAlertRequestBody>`
    - 响应参数变更
      - `* data: object<DataResponse> -> object<BatchOperateAlertResult>`
  - **CreateAlert**
    - 请求参数变更
      - `+ data_object.domain_id`
      - `+ data_object.region_id`
      - `+ data_object.labels`
      - `+ data_object.creator`
      - `- data_object.label`
      - `- data_object.chop_phase`
      - `- data_object.ppdr_phase`
      - `- data_object.cteator`
      - `+ data_object.environment.cross_workspace_id`
      - `+ data_object.data_source.company_name`
      - `+ data_object.data_source.product_module`
      - `+ data_object.severity: enum value [Tips,Low,Medium,High,Fatal]`
      - `+ data_object.alert_type.category`
      - `+ data_object.alert_type.alert_type`
      - `* data_object.network_list.direction: object -> string`
      - `* data_object.network_list.src_port: string -> int32`
      - `+ data_object.network_list.src_geo.latitude`
      - `+ data_object.network_list.src_geo.longitude`
      - `+ data_object.network_list.src_geo.city_code`
      - `+ data_object.network_list.src_geo.country_code`
      - `+ data_object.network_list.dest_geo.latitude`
      - `+ data_object.network_list.dest_geo.longitude`
      - `+ data_object.network_list.dest_geo.city_code`
      - `+ data_object.network_list.dest_geo.country_code`
      - `+ data_object.resource_list.provider`
      - `+ data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `* data_object.simulation: boolean -> string`
      - `+ data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data_object.process.process_parent_name`
      - `+ data_object.process.process_parent_path`
      - `+ data_object.process.process_parent_pid`
      - `+ data_object.process.process_parent_uid`
      - `+ data_object.process.process_parent_cmdline`
      - `+ data_object.process.process_child_name`
      - `+ data_object.process.process_child_path`
      - `+ data_object.process.process_child_pid`
      - `+ data_object.process.process_child_uid`
      - `+ data_object.process.process_child_cmdline`
      - `+ data_object.process.process_launche_time`
      - `+ data_object.process.process_terminate_time`
      - `* data_object.process.process_pid: string -> int32`
      - `* data_object.process.process_uid: string -> int32`
      - `* data_object: object<CreateAlert> -> object<Alert>`
    - 响应参数变更
      - `- data.type`
      - `+ data.data_object.domain_id`
      - `+ data.data_object.region_id`
      - `+ data.data_object.labels`
      - `+ data.data_object.data_source`
      - `+ data.data_object.severity`
      - `+ data.data_object.creator`
      - `- data.data_object.datasource`
      - `- data.data_object.serverity`
      - `- data.data_object.chop_phase`
      - `- data.data_object.ppdr_phase`
      - `- data.data_object.cteator`
      - `+ data.data_object.environment.cross_workspace_id`
      - `+ data.data_object.alert_type.category`
      - `+ data.data_object.alert_type.alert_type`
      - `* data.data_object.network_list.direction: object -> string`
      - `* data.data_object.network_list.src_port: string -> int32`
      - `+ data.data_object.network_list.src_geo.latitude`
      - `+ data.data_object.network_list.src_geo.longitude`
      - `+ data.data_object.network_list.src_geo.city_code`
      - `+ data.data_object.network_list.src_geo.country_code`
      - `+ data.data_object.network_list.dest_geo.latitude`
      - `+ data.data_object.network_list.dest_geo.longitude`
      - `+ data.data_object.network_list.dest_geo.city_code`
      - `+ data.data_object.network_list.dest_geo.country_code`
      - `+ data.data_object.resource_list.provider`
      - `+ data.data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data.data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data.data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `* data.data_object.simulation: boolean -> string`
      - `+ data.data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data.data_object.process.process_parent_name`
      - `+ data.data_object.process.process_parent_path`
      - `+ data.data_object.process.process_parent_pid`
      - `+ data.data_object.process.process_parent_uid`
      - `+ data.data_object.process.process_parent_cmdline`
      - `+ data.data_object.process.process_child_name`
      - `+ data.data_object.process.process_child_path`
      - `+ data.data_object.process.process_child_pid`
      - `+ data.data_object.process.process_child_uid`
      - `+ data.data_object.process.process_child_cmdline`
      - `+ data.data_object.process.process_launche_time`
      - `+ data.data_object.process.process_terminate_time`
      - `* data.data_object.process.process_pid: string -> int32`
      - `* data.data_object.process.process_uid: string -> int32`
  - **ShowAlert**
    - 响应参数变更
      - `- data.type`
      - `+ data.data_object.domain_id`
      - `+ data.data_object.region_id`
      - `+ data.data_object.labels`
      - `+ data.data_object.data_source`
      - `+ data.data_object.severity`
      - `+ data.data_object.creator`
      - `- data.data_object.datasource`
      - `- data.data_object.serverity`
      - `- data.data_object.chop_phase`
      - `- data.data_object.ppdr_phase`
      - `- data.data_object.cteator`
      - `+ data.data_object.environment.cross_workspace_id`
      - `* data.data_object.count: string -> int32`
      - `+ data.data_object.alert_type.category`
      - `+ data.data_object.alert_type.alert_type`
      - `* data.data_object.network_list.direction: object -> string`
      - `* data.data_object.network_list.src_port: string -> int32`
      - `+ data.data_object.network_list.src_geo.latitude`
      - `+ data.data_object.network_list.src_geo.longitude`
      - `+ data.data_object.network_list.src_geo.city_code`
      - `+ data.data_object.network_list.src_geo.country_code`
      - `+ data.data_object.network_list.dest_geo.latitude`
      - `+ data.data_object.network_list.dest_geo.longitude`
      - `+ data.data_object.network_list.dest_geo.city_code`
      - `+ data.data_object.network_list.dest_geo.country_code`
      - `+ data.data_object.resource_list.provider`
      - `+ data.data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data.data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data.data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `* data.data_object.simulation: boolean -> string`
      - `+ data.data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data.data_object.process.process_parent_name`
      - `+ data.data_object.process.process_parent_path`
      - `+ data.data_object.process.process_parent_pid`
      - `+ data.data_object.process.process_parent_uid`
      - `+ data.data_object.process.process_parent_cmdline`
      - `+ data.data_object.process.process_child_name`
      - `+ data.data_object.process.process_child_path`
      - `+ data.data_object.process.process_child_pid`
      - `+ data.data_object.process.process_child_uid`
      - `+ data.data_object.process.process_child_cmdline`
      - `+ data.data_object.process.process_launche_time`
      - `+ data.data_object.process.process_terminate_time`
      - `* data.data_object.process.process_pid: string -> int32`
      - `* data.data_object.process.process_uid: string -> int32`
      - `* data.data_object: object<ShowAlertRsp> -> object<Alert>`
      - `* data: object<ShowAlertDetail> -> object<AlertDetail>`
  - **ChangeAlert**
    - 请求参数变更
      - `+ batch_ids`
      - `+ data_object.domain_id`
      - `+ data_object.region_id`
      - `+ data_object.labels`
      - `+ data_object.data_source`
      - `+ data_object.severity`
      - `+ data_object.creator`
      - `- data_object.datasource`
      - `- data_object.serverity`
      - `- data_object.chop_phase`
      - `- data_object.ppdr_phase`
      - `- data_object.cteator`
      - `+ data_object.environment.cross_workspace_id`
      - `+ data_object.alert_type.category`
      - `+ data_object.alert_type.alert_type`
      - `* data_object.network_list.direction: object -> string`
      - `* data_object.network_list.src_port: string -> int32`
      - `+ data_object.network_list.src_geo.latitude`
      - `+ data_object.network_list.src_geo.longitude`
      - `+ data_object.network_list.src_geo.city_code`
      - `+ data_object.network_list.src_geo.country_code`
      - `+ data_object.network_list.dest_geo.latitude`
      - `+ data_object.network_list.dest_geo.longitude`
      - `+ data_object.network_list.dest_geo.city_code`
      - `+ data_object.network_list.dest_geo.country_code`
      - `+ data_object.resource_list.provider`
      - `+ data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `* data_object.simulation: boolean -> string`
      - `+ data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data_object.process.process_parent_name`
      - `+ data_object.process.process_parent_path`
      - `+ data_object.process.process_parent_pid`
      - `+ data_object.process.process_parent_uid`
      - `+ data_object.process.process_parent_cmdline`
      - `+ data_object.process.process_child_name`
      - `+ data_object.process.process_child_path`
      - `+ data_object.process.process_child_pid`
      - `+ data_object.process.process_child_uid`
      - `+ data_object.process.process_child_cmdline`
      - `+ data_object.process.process_launche_time`
      - `+ data_object.process.process_terminate_time`
      - `* data_object.process.process_pid: string -> int32`
      - `* data_object.process.process_uid: string -> int32`
    - 响应参数变更
      - `- data.type`
      - `+ data.data_object.domain_id`
      - `+ data.data_object.region_id`
      - `+ data.data_object.labels`
      - `+ data.data_object.data_source`
      - `+ data.data_object.severity`
      - `+ data.data_object.creator`
      - `- data.data_object.datasource`
      - `- data.data_object.serverity`
      - `- data.data_object.chop_phase`
      - `- data.data_object.ppdr_phase`
      - `- data.data_object.cteator`
      - `+ data.data_object.environment.cross_workspace_id`
      - `+ data.data_object.alert_type.category`
      - `+ data.data_object.alert_type.alert_type`
      - `* data.data_object.network_list.direction: object -> string`
      - `* data.data_object.network_list.src_port: string -> int32`
      - `+ data.data_object.network_list.src_geo.latitude`
      - `+ data.data_object.network_list.src_geo.longitude`
      - `+ data.data_object.network_list.src_geo.city_code`
      - `+ data.data_object.network_list.src_geo.country_code`
      - `+ data.data_object.network_list.dest_geo.latitude`
      - `+ data.data_object.network_list.dest_geo.longitude`
      - `+ data.data_object.network_list.dest_geo.city_code`
      - `+ data.data_object.network_list.dest_geo.country_code`
      - `+ data.data_object.resource_list.provider`
      - `+ data.data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data.data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data.data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `* data.data_object.simulation: boolean -> string`
      - `+ data.data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data.data_object.process.process_parent_name`
      - `+ data.data_object.process.process_parent_path`
      - `+ data.data_object.process.process_parent_pid`
      - `+ data.data_object.process.process_parent_uid`
      - `+ data.data_object.process.process_parent_cmdline`
      - `+ data.data_object.process.process_child_name`
      - `+ data.data_object.process.process_child_path`
      - `+ data.data_object.process.process_child_pid`
      - `+ data.data_object.process.process_child_uid`
      - `+ data.data_object.process.process_child_cmdline`
      - `+ data.data_object.process.process_launche_time`
      - `+ data.data_object.process.process_terminate_time`
      - `* data.data_object.process.process_pid: string -> int32`
      - `* data.data_object.process.process_uid: string -> int32`
  - **ListIncidents**
    - 请求参数变更
      - `+ order: enum value [DESC,ASC]`
      - `* condition.conditions.data: list<string> -> list<object>`
      - `* condition.conditions: list<ConditonInfo> -> list<object>`
      - `* condition.logics: list<string> -> list<object>`
    - 响应参数变更
      - `+ data.dataclass_ref`
      - `+ data.format_version`
      - `+ data.id`
      - `+ data.version`
      - `- data.dataclass_id`
      - `- data.layout_id`
      - `- data.name`
      - `- data.type`
      - `- data.dataclass`
      - `+ data.data_object.domain_id`
      - `+ data.data_object.region_id`
      - `+ data.data_object.labels`
      - `+ data.data_object.data_source`
      - `+ data.data_object.severity`
      - `+ data.data_object.creator`
      - `+ data.data_object.system_alert_table`
      - `- data.data_object.datasource`
      - `- data.data_object.serverity`
      - `- data.data_object.chop_phase`
      - `- data.data_object.ppdr_phase`
      - `- data.data_object.cteator`
      - `- data.data_object.system_incident_table`
      - `+ data.data_object.environment.cross_workspace_id`
      - `+ data.data_object.incident_type.category`
      - `+ data.data_object.incident_type.incident_type`
      - `* data.data_object.network_list.direction: object -> string`
      - `* data.data_object.network_list.src_port: string -> int32`
      - `+ data.data_object.network_list.src_geo.latitude`
      - `+ data.data_object.network_list.src_geo.longitude`
      - `+ data.data_object.network_list.src_geo.city_code`
      - `+ data.data_object.network_list.src_geo.country_code`
      - `+ data.data_object.network_list.dest_geo.latitude`
      - `+ data.data_object.network_list.dest_geo.longitude`
      - `+ data.data_object.network_list.dest_geo.city_code`
      - `+ data.data_object.network_list.dest_geo.country_code`
      - `+ data.data_object.resource_list.provider`
      - `+ data.data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data.data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data.data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `+ data.data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data.data_object.process.process_parent_name`
      - `+ data.data_object.process.process_parent_path`
      - `+ data.data_object.process.process_parent_pid`
      - `+ data.data_object.process.process_parent_uid`
      - `+ data.data_object.process.process_parent_cmdline`
      - `+ data.data_object.process.process_child_name`
      - `+ data.data_object.process.process_child_path`
      - `+ data.data_object.process.process_child_pid`
      - `+ data.data_object.process.process_child_uid`
      - `+ data.data_object.process.process_child_cmdline`
      - `+ data.data_object.process.process_launche_time`
      - `+ data.data_object.process.process_terminate_time`
      - `* data.data_object.process.process_pid: string -> int32`
      - `* data.data_object.process.process_uid: string -> int32`
      - `* data: list<ListIncidentDetail> -> list<IncidentDetail>`
  - **DeleteIncident**
    - 请求参数变更
      - `+ batch_ids`
      - `- ids`
      - `* body: object<DeleteIncident> -> object<DeleteIncidentRequestBody>`
    - 响应参数变更
      - `* data: object<BatchOrderAlertResult> -> object`
  - **CreateIncident**
    - 请求参数变更
      - `+ data_object.domain_id`
      - `+ data_object.region_id`
      - `+ data_object.creator`
      - `+ data_object.system_alert_table`
      - `- data_object.serverity`
      - `- data_object.chop_phase`
      - `- data_object.ppdr_phase`
      - `- data_object.cteator`
      - `- data_object.system_incident_table`
      - `+ data_object.severity: enum value [Tips,Low,Medium,High,Fatal]`
      - `+ data_object.environment.cross_workspace_id`
      - `+ data_object.data_source.company_name`
      - `+ data_object.data_source.product_module`
      - `- data_object.incident_type.id`
      - `- data_object.incident_type.layoutId`
      - `* data_object.network_list.direction: object -> string`
      - `* data_object.network_list.src_port: string -> int32`
      - `+ data_object.network_list.src_geo.latitude`
      - `+ data_object.network_list.src_geo.longitude`
      - `+ data_object.network_list.src_geo.city_code`
      - `+ data_object.network_list.src_geo.country_code`
      - `+ data_object.network_list.dest_geo.latitude`
      - `+ data_object.network_list.dest_geo.longitude`
      - `+ data_object.network_list.dest_geo.city_code`
      - `+ data_object.network_list.dest_geo.country_code`
      - `+ data_object.resource_list.provider`
      - `+ data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `+ data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data_object.process.process_parent_name`
      - `+ data_object.process.process_parent_path`
      - `+ data_object.process.process_parent_pid`
      - `+ data_object.process.process_parent_uid`
      - `+ data_object.process.process_parent_cmdline`
      - `+ data_object.process.process_child_name`
      - `+ data_object.process.process_child_path`
      - `+ data_object.process.process_child_pid`
      - `+ data_object.process.process_child_uid`
      - `+ data_object.process.process_child_cmdline`
      - `+ data_object.process.process_launche_time`
      - `+ data_object.process.process_terminate_time`
      - `* data_object.process.process_pid: string -> int32`
      - `* data_object.process.process_uid: string -> int32`
      - `* data_object: object<CreateIncident> -> object<Incident>`
    - 响应参数变更
      - `+ data.dataclass_ref`
      - `+ data.format_version`
      - `+ data.id`
      - `+ data.version`
      - `- data.dataclass_id`
      - `- data.layout_id`
      - `- data.name`
      - `- data.type`
      - `- data.dataclass`
      - `+ data.data_object.domain_id`
      - `+ data.data_object.region_id`
      - `+ data.data_object.labels`
      - `+ data.data_object.data_source`
      - `+ data.data_object.severity`
      - `+ data.data_object.creator`
      - `+ data.data_object.system_alert_table`
      - `- data.data_object.datasource`
      - `- data.data_object.serverity`
      - `- data.data_object.chop_phase`
      - `- data.data_object.ppdr_phase`
      - `- data.data_object.cteator`
      - `- data.data_object.system_incident_table`
      - `+ data.data_object.environment.cross_workspace_id`
      - `+ data.data_object.incident_type.category`
      - `+ data.data_object.incident_type.incident_type`
      - `* data.data_object.network_list.direction: object -> string`
      - `* data.data_object.network_list.src_port: string -> int32`
      - `+ data.data_object.network_list.src_geo.latitude`
      - `+ data.data_object.network_list.src_geo.longitude`
      - `+ data.data_object.network_list.src_geo.city_code`
      - `+ data.data_object.network_list.src_geo.country_code`
      - `+ data.data_object.network_list.dest_geo.latitude`
      - `+ data.data_object.network_list.dest_geo.longitude`
      - `+ data.data_object.network_list.dest_geo.city_code`
      - `+ data.data_object.network_list.dest_geo.country_code`
      - `+ data.data_object.resource_list.provider`
      - `+ data.data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data.data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data.data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `+ data.data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data.data_object.process.process_parent_name`
      - `+ data.data_object.process.process_parent_path`
      - `+ data.data_object.process.process_parent_pid`
      - `+ data.data_object.process.process_parent_uid`
      - `+ data.data_object.process.process_parent_cmdline`
      - `+ data.data_object.process.process_child_name`
      - `+ data.data_object.process.process_child_path`
      - `+ data.data_object.process.process_child_pid`
      - `+ data.data_object.process.process_child_uid`
      - `+ data.data_object.process.process_child_cmdline`
      - `+ data.data_object.process.process_launche_time`
      - `+ data.data_object.process.process_terminate_time`
      - `* data.data_object.process.process_pid: string -> int32`
      - `* data.data_object.process.process_uid: string -> int32`
  - **ShowIncident**
    - 响应参数变更
      - `+ data.dataclass_ref`
      - `+ data.format_version`
      - `+ data.id`
      - `+ data.version`
      - `- data.dataclass_id`
      - `- data.layout_id`
      - `- data.name`
      - `- data.type`
      - `- data.dataclass`
      - `+ data.data_object.domain_id`
      - `+ data.data_object.region_id`
      - `+ data.data_object.labels`
      - `+ data.data_object.data_source`
      - `+ data.data_object.severity`
      - `+ data.data_object.creator`
      - `+ data.data_object.system_alert_table`
      - `- data.data_object.datasource`
      - `- data.data_object.serverity`
      - `- data.data_object.chop_phase`
      - `- data.data_object.ppdr_phase`
      - `- data.data_object.cteator`
      - `- data.data_object.system_incident_table`
      - `+ data.data_object.environment.cross_workspace_id`
      - `* data.data_object.count: string -> int32`
      - `+ data.data_object.incident_type.category`
      - `+ data.data_object.incident_type.incident_type`
      - `* data.data_object.network_list.direction: object -> string`
      - `* data.data_object.network_list.src_port: string -> int32`
      - `+ data.data_object.network_list.src_geo.latitude`
      - `+ data.data_object.network_list.src_geo.longitude`
      - `+ data.data_object.network_list.src_geo.city_code`
      - `+ data.data_object.network_list.src_geo.country_code`
      - `+ data.data_object.network_list.dest_geo.latitude`
      - `+ data.data_object.network_list.dest_geo.longitude`
      - `+ data.data_object.network_list.dest_geo.city_code`
      - `+ data.data_object.network_list.dest_geo.country_code`
      - `+ data.data_object.resource_list.provider`
      - `+ data.data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data.data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data.data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `+ data.data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data.data_object.process.process_parent_name`
      - `+ data.data_object.process.process_parent_path`
      - `+ data.data_object.process.process_parent_pid`
      - `+ data.data_object.process.process_parent_uid`
      - `+ data.data_object.process.process_parent_cmdline`
      - `+ data.data_object.process.process_child_name`
      - `+ data.data_object.process.process_child_path`
      - `+ data.data_object.process.process_child_pid`
      - `+ data.data_object.process.process_child_uid`
      - `+ data.data_object.process.process_child_cmdline`
      - `+ data.data_object.process.process_launche_time`
      - `+ data.data_object.process.process_terminate_time`
      - `* data.data_object.process.process_pid: string -> int32`
      - `* data.data_object.process.process_uid: string -> int32`
      - `* data.data_object: object<ShowIncident> -> object<Incident>`
      - `* data: object<ShowIncidentDetail> -> object<IncidentDetail>`
  - **ChangeIncident**
    - 请求参数变更
      - `+ batch_ids`
      - `+ data_object.domain_id`
      - `+ data_object.region_id`
      - `+ data_object.labels`
      - `+ data_object.data_source`
      - `+ data_object.severity`
      - `+ data_object.creator`
      - `+ data_object.system_alert_table`
      - `- data_object.datasource`
      - `- data_object.serverity`
      - `- data_object.chop_phase`
      - `- data_object.ppdr_phase`
      - `- data_object.cteator`
      - `- data_object.system_incident_table`
      - `+ data_object.environment.cross_workspace_id`
      - `+ data_object.incident_type.category`
      - `+ data_object.incident_type.incident_type`
      - `* data_object.network_list.direction: object -> string`
      - `* data_object.network_list.src_port: string -> int32`
      - `+ data_object.network_list.src_geo.latitude`
      - `+ data_object.network_list.src_geo.longitude`
      - `+ data_object.network_list.src_geo.city_code`
      - `+ data_object.network_list.src_geo.country_code`
      - `+ data_object.network_list.dest_geo.latitude`
      - `+ data_object.network_list.dest_geo.longitude`
      - `+ data_object.network_list.dest_geo.city_code`
      - `+ data_object.network_list.dest_geo.country_code`
      - `+ data_object.resource_list.provider`
      - `+ data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `+ data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data_object.process.process_parent_name`
      - `+ data_object.process.process_parent_path`
      - `+ data_object.process.process_parent_pid`
      - `+ data_object.process.process_parent_uid`
      - `+ data_object.process.process_parent_cmdline`
      - `+ data_object.process.process_child_name`
      - `+ data_object.process.process_child_path`
      - `+ data_object.process.process_child_pid`
      - `+ data_object.process.process_child_uid`
      - `+ data_object.process.process_child_cmdline`
      - `+ data_object.process.process_launche_time`
      - `+ data_object.process.process_terminate_time`
      - `* data_object.process.process_pid: string -> int32`
      - `* data_object.process.process_uid: string -> int32`
    - 响应参数变更
      - `+ data.dataclass_ref`
      - `+ data.format_version`
      - `+ data.id`
      - `+ data.version`
      - `- data.dataclass_id`
      - `- data.layout_id`
      - `- data.name`
      - `- data.type`
      - `- data.dataclass`
      - `+ data.data_object.domain_id`
      - `+ data.data_object.region_id`
      - `+ data.data_object.labels`
      - `+ data.data_object.data_source`
      - `+ data.data_object.severity`
      - `+ data.data_object.creator`
      - `+ data.data_object.system_alert_table`
      - `- data.data_object.datasource`
      - `- data.data_object.serverity`
      - `- data.data_object.chop_phase`
      - `- data.data_object.ppdr_phase`
      - `- data.data_object.cteator`
      - `- data.data_object.system_incident_table`
      - `+ data.data_object.environment.cross_workspace_id`
      - `+ data.data_object.incident_type.category`
      - `+ data.data_object.incident_type.incident_type`
      - `* data.data_object.network_list.direction: object -> string`
      - `* data.data_object.network_list.src_port: string -> int32`
      - `+ data.data_object.network_list.src_geo.latitude`
      - `+ data.data_object.network_list.src_geo.longitude`
      - `+ data.data_object.network_list.src_geo.city_code`
      - `+ data.data_object.network_list.src_geo.country_code`
      - `+ data.data_object.network_list.dest_geo.latitude`
      - `+ data.data_object.network_list.dest_geo.longitude`
      - `+ data.data_object.network_list.dest_geo.city_code`
      - `+ data.data_object.network_list.dest_geo.country_code`
      - `+ data.data_object.resource_list.provider`
      - `+ data.data_object.verification_state: enum value [Unknown,True_Positive,False_Positive]`
      - `+ data.data_object.handle_status: enum value [Open,Block,Closed]`
      - `+ data.data_object.ipdrr_phase: enum value [Prepartion,Detection and Analysis,Containm，Eradication& Recovery,Post-Incident-Activity]`
      - `+ data.data_object.close_reason: enum value [False detection,Resolved,Repeated,Other]`
      - `+ data.data_object.process.process_parent_name`
      - `+ data.data_object.process.process_parent_path`
      - `+ data.data_object.process.process_parent_pid`
      - `+ data.data_object.process.process_parent_uid`
      - `+ data.data_object.process.process_parent_cmdline`
      - `+ data.data_object.process.process_child_name`
      - `+ data.data_object.process.process_child_path`
      - `+ data.data_object.process.process_child_pid`
      - `+ data.data_object.process.process_child_uid`
      - `+ data.data_object.process.process_child_cmdline`
      - `+ data.data_object.process.process_launche_time`
      - `+ data.data_object.process.process_terminate_time`
      - `* data.data_object.process.process_pid: string -> int32`
      - `* data.data_object.process.process_uid: string -> int32`
  - **ListIndicators**
    - 请求参数变更
      - `- order`
      - `- from_date`
      - `- to_date`
      - `+ from_date`
      - `+ to_date`
      - `- type`
    - 响应参数变更
      - `+ data.dataclass_ref`
      - `- data.dataclass_id`
      - `- data.type`
      - `- data.layout_id`
      - `- data.dataclass`
      - `+ data.data_object.update_time`
      - `+ data.data_object.create_time`
      - `+ data.data_object.environment`
      - `+ data.data_object.data_source`
      - `+ data.data_object.first_report_time`
      - `+ data.data_object.is_deleted`
      - `+ data.data_object.last_report_time`
      - `+ data.data_object.granular_marking`
      - `+ data.data_object.name`
      - `+ data.data_object.id`
      - `+ data.data_object.project_id`
      - `+ data.data_object.revoked`
      - `+ data.data_object.status`
      - `+ data.data_object.verdict`
      - `+ data.data_object.workspace_id`
      - `+ data.data_object.confidence`
      - `+ data.data_object.indicator_type.layout_id`
      - `- data.data_object.indicator_type.layoutId`
  - **CreateIndicator**
    - 请求参数变更
      - `- name`
      - `- format_version`
      - `- type`
      - `- trigger_flag`
      - `- data_object.type`
      - `+ data_object.indicator_type.layout_id`
      - `- data_object.indicator_type.layoutId`
      - `+ data_object.data_object.update_time`
      - `+ data_object.data_object.create_time`
      - `+ data_object.data_object.environment`
      - `+ data_object.data_object.data_source`
      - `+ data_object.data_object.first_report_time`
      - `+ data_object.data_object.is_deleted`
      - `+ data_object.data_object.last_report_time`
      - `+ data_object.data_object.granular_marking`
      - `+ data_object.data_object.name`
      - `+ data_object.data_object.id`
      - `+ data_object.data_object.project_id`
      - `+ data_object.data_object.revoked`
      - `+ data_object.data_object.status`
      - `+ data_object.data_object.verdict`
      - `+ data_object.data_object.workspace_id`
      - `+ data_object.data_object.confidence`
      - `+ data_object.data_object.indicator_type.layout_id`
      - `- data_object.data_object.indicator_type.layoutId`
    - 响应参数变更
      - `+ data.dataclass_ref`
      - `- data.dataclass_id`
      - `- data.type`
      - `- data.layout_id`
      - `- data.dataclass`
      - `+ data.data_object.update_time`
      - `+ data.data_object.create_time`
      - `+ data.data_object.environment`
      - `+ data.data_object.data_source`
      - `+ data.data_object.first_report_time`
      - `+ data.data_object.is_deleted`
      - `+ data.data_object.last_report_time`
      - `+ data.data_object.granular_marking`
      - `+ data.data_object.name`
      - `+ data.data_object.id`
      - `+ data.data_object.project_id`
      - `+ data.data_object.revoked`
      - `+ data.data_object.status`
      - `+ data.data_object.verdict`
      - `+ data.data_object.workspace_id`
      - `+ data.data_object.confidence`
      - `+ data.data_object.indicator_type.layout_id`
      - `- data.data_object.indicator_type.layoutId`
  - **ShowIndicatorDetail**
    - 响应参数变更
      - `+ data.dataclass_ref`
      - `- data.dataclass_id`
      - `- data.type`
      - `- data.layout_id`
      - `- data.dataclass`
      - `+ data.data_object.update_time`
      - `+ data.data_object.create_time`
      - `+ data.data_object.environment`
      - `+ data.data_object.data_source`
      - `+ data.data_object.first_report_time`
      - `+ data.data_object.is_deleted`
      - `+ data.data_object.last_report_time`
      - `+ data.data_object.granular_marking`
      - `+ data.data_object.name`
      - `+ data.data_object.id`
      - `+ data.data_object.project_id`
      - `+ data.data_object.revoked`
      - `+ data.data_object.status`
      - `+ data.data_object.verdict`
      - `+ data.data_object.workspace_id`
      - `+ data.data_object.confidence`
      - `+ data.data_object.indicator_type.layout_id`
      - `- data.data_object.indicator_type.layoutId`
  - **UpdateIndicator**
    - 请求参数变更
      - `- trigger_flag`
      - `+ data_object.update_time`
      - `+ data_object.create_time`
      - `+ data_object.environment`
      - `+ data_object.data_source`
      - `+ data_object.first_report_time`
      - `+ data_object.is_deleted`
      - `+ data_object.last_report_time`
      - `+ data_object.granular_marking`
      - `+ data_object.name`
      - `+ data_object.id`
      - `+ data_object.project_id`
      - `+ data_object.revoked`
      - `+ data_object.status`
      - `+ data_object.verdict`
      - `+ data_object.workspace_id`
      - `+ data_object.confidence`
      - `+ data_object.indicator_type.layout_id`
      - `- data_object.indicator_type.layoutId`
    - 响应参数变更
      - `+ data.dataclass_ref`
      - `- data.dataclass_id`
      - `- data.type`
      - `- data.layout_id`
      - `- data.dataclass`
      - `+ data.data_object.update_time`
      - `+ data.data_object.create_time`
      - `+ data.data_object.environment`
      - `+ data.data_object.data_source`
      - `+ data.data_object.first_report_time`
      - `+ data.data_object.is_deleted`
      - `+ data.data_object.last_report_time`
      - `+ data.data_object.granular_marking`
      - `+ data.data_object.name`
      - `+ data.data_object.id`
      - `+ data.data_object.project_id`
      - `+ data.data_object.revoked`
      - `+ data.data_object.status`
      - `+ data.data_object.verdict`
      - `+ data.data_object.workspace_id`
      - `+ data.data_object.confidence`
      - `+ data.data_object.indicator_type.layout_id`
      - `- data.data_object.indicator_type.layoutId`
  - **ShowPlaybookMonitors**
    - 请求参数变更
      - `+ version_query_type: enum value [ALL:全部，VALID:有效的，DELETED:已删除]`
    - 响应参数变更
      - `* data.total_instance_run_num: int32 -> float`
  - **CreateAlertRule**
    - 请求参数变更
      - `- accumulated_times`
      - `- query_type: enum value [CBSL]`
    - 响应参数变更
      - `- accumulated_times`
      - `- query_type: enum value [CBSL]`
  - **ListAlertRules**
    - 响应参数变更
      - `- accumulated_times`
      - `- records.accumulated_times`
      - `- records.query_type: enum value [CBSL]`
  - **ListAlertRuleTemplates**
    - 响应参数变更
      - `- records.query_type: enum value [CBSL]`
  - **CopyPlaybookVersion**
    - 响应参数变更
      - `- data.run_mode`
      - `- data.dataobject_id`
      - `- data.actions.sort_order`
  - **CreatePlaybookRule**
    - 请求参数变更
      - `+ rule.start_type`
      - `+ rule.end_type`
      - `+ rule.end_time`
      - `+ rule.only_once`
      - `+ rule.execution_type`
      - `- rule.repeat_count`
      - `* rule.logics: object -> list<string>`
  - **ListPlaybookInstances**
    - 请求参数变更
      - `- date_type`
      - `* limit: optional -> required`
      - `* offset: optional -> required`
  - **ShowPlaybookTopology**
    - 响应参数变更
      - `- action_instances.action.sort_order`
  - **ListPlaybookVersions**
    - 请求参数变更
      - `- approve_role`
    - 响应参数变更
      - `- data.run_mode`
      - `- data.dataobject_id`
  - **CreatePlaybookVersion**
    - 请求参数变更
      - `- actions.sort_order`
    - 响应参数变更
      - `- data.run_mode`
      - `- data.dataobject_id`
      - `- data.actions.sort_order`
  - **ShowPlaybookVersion**
    - 响应参数变更
      - `- data.run_mode`
      - `- data.dataobject_id`
      - `- data.actions.sort_order`
  - **UpdatePlaybookVersion**
    - 响应参数变更
      - `- data.run_mode`
      - `- data.dataobject_id`
      - `- data.actions.sort_order`
  - **UpdatePlaybookRule**
    - 请求参数变更
      - `+ rule.start_type`
      - `+ rule.end_type`
      - `+ rule.end_time`
      - `+ rule.only_once`
      - `+ rule.execution_type`
      - `- rule.repeat_count`
      - `* rule.logics: object -> list<string>`

# 0.1.63 2023-10-26

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`ShowStackInstance`、`UpdateStackInstances`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDetailsOfApiV2**
    - 响应参数变更
      - `+ policy_functions.conditions.sys_param_name`
      - `+ policy_functions.conditions.cookie_param_name`
      - `+ policy_functions.conditions.frontend_authorizer_param_name`
      - `+ policy_functions.conditions.condition_origin: enum value [system,cookie,frontend_authorizer]`
  - **UpdateApiV2**
    - 请求参数变更
      - `+ policy_mocks.conditions.sys_param_name`
      - `+ policy_mocks.conditions.cookie_param_name`
      - `+ policy_mocks.conditions.frontend_authorizer_param_name`
      - `+ policy_mocks.conditions.condition_origin: enum value [system,cookie,frontend_authorizer]`
    - 响应参数变更
      - `+ policy_functions.conditions.sys_param_name`
      - `+ policy_functions.conditions.cookie_param_name`
      - `+ policy_functions.conditions.frontend_authorizer_param_name`
      - `+ policy_functions.conditions.condition_origin: enum value [system,cookie,frontend_authorizer]`
  - **ListApiVersionDetailV2**
    - 响应参数变更
      - `+ policy_functions.conditions.sys_param_name`
      - `+ policy_functions.conditions.cookie_param_name`
      - `+ policy_functions.conditions.frontend_authorizer_param_name`
      - `+ policy_functions.conditions.condition_origin: enum value [system,cookie,frontend_authorizer]`
  - **CreateApiV2**
    - 请求参数变更
      - `+ policy_mocks.conditions.sys_param_name`
      - `+ policy_mocks.conditions.cookie_param_name`
      - `+ policy_mocks.conditions.frontend_authorizer_param_name`
      - `+ policy_mocks.conditions.condition_origin: enum value [system,cookie,frontend_authorizer]`
    - 响应参数变更
      - `+ policy_functions.conditions.sys_param_name`
      - `+ policy_functions.conditions.cookie_param_name`
      - `+ policy_functions.conditions.frontend_authorizer_param_name`
      - `+ policy_functions.conditions.condition_origin: enum value [system,cookie,frontend_authorizer]`

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateBareMetalServers**
    - 请求参数变更
      - `+ server.extendparam.chargingMode: enum value [postPaid]`

### HuaweiCloud SDK CC

- _新增特性_
  - 支持接口`ListCentralNetworkCapabilities`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持以下接口：
    - `CreateRefreshTasks`
    - `CreatePreheatingTasks`
    - `ShowHistoryTasks`
    - `ShowHistoryTaskDetails`
    - `ShowUrlTaskInfo`
- _解决问题_
  - 无
- _特性变更_
  - **CreateRefreshTasks**
    - 请求参数变更
      - `+ refresh_task.zh_url_encode`
  - **CreatePreheatingTasks**
    - 请求参数变更
      - `+ preheating_task.zh_url_encode`

### HuaweiCloud SDK CodeArtsPipeline

- _新增特性_
  - 支持接口`CreatePipelineNew`、`RetryPipelineRun`、`AcceptManualReview`、`RejectManualReview`
- _解决问题_
  - 无
- _特性变更_
  - **ListPipelines**
    - 响应参数变更
      - `+ pipelines.latest_run.stage_status_list.id`
  - **CreatePipelineByTemplateId**
    - 请求参数变更
      - `+ variables`

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持以下接口：
    - `ShowBackgroundTaskProgress`
    - `ListCenterTask`
    - `StartInstanceResizeCheckJob`
    - `ShowBackUpInfo`
    - `CreateOrUpdateBackUpInfo`
    - `ShowExpireKeyScanInfo`
    - `ScanExpireKey`
    - `ListMigrationTaskLogs`
    - `CheckMigrationConnectivity`
    - `ExchangeInstanceIp`
    - `ExecuteCommandMobilization`
    - `LoginWebCli`
    - `UpdateMigrationTask`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DLI

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListBatches**
    - 请求参数变更
      - `+ job-name`
  - **CreateBatchJob**
    - 响应参数变更
      - `- req_body`

### HuaweiCloud SDK HSS

- _新增特性_
  - 支持以下接口：
    - `ListContainerNodes`
    - `ListBlockedIp`
    - `ChangeBlockedIp`
    - `ListIsolatedFile`
    - `ChangeIsolatedFile`
    - `ListSwrImageRepository`
    - `BatchScanSwrImage`
    - `ListImageVulnerabilities`
    - `ListVulnerabilityCve`
    - `RunImageSynchronize`
    - `ListImageRiskConfigs`
    - `ListImageRiskConfigRules`
    - `ShowImageCheckRuleDetail`
    - `ShowVulScanPolicy`
    - `ChangeVulScanPolicy`
    - `ShowVulStatics`
- _解决问题_
  - 无
- _特性变更_
  - **ListPortStatistics**
    - 请求参数变更
      - `+ port_string`
      - `+ sort_key`
      - `+ sort_dir`
    - 响应参数变更
      - `+ data_list.status`
  - **ListProtectionServer**
    - 响应参数变更
      - `+ data_list.host_source`
  - **ListHostStatus**
    - 请求参数变更
      - `+ has_intrusion`
      - `+ agent_upgradable`
  - **ListVulHosts**
    - 响应参数变更
      - `+ data_list.support_restore`
  - **ChangeVulStatus**
    - 请求参数变更
      - `+ backup_info_id`
      - `+ custom_backup_hosts`
  - **ListHostVuls**
    - 响应参数变更
      - `+ data_list.app_name`
      - `+ data_list.app_version`
      - `+ data_list.app_path`
      - `+ data_list.version`
      - `+ data_list.support_restore`
  - **ListHostProtectHistoryInfo**
    - 请求参数变更
      - `+ host_name`
      - `+ host_ip`
      - `+ file_path`
      - `+ file_operation`
  - **ListProtectionPolicy**
    - 响应参数变更
      - `+ data_list.deploy_mode`
      - `+ data_list.default_policy`
  - **ListSecurityEvents**
    - 请求参数变更
      - `+ severity_list`
      - `+ attack_tag`
      - `+ asset_value`
      - `+ tag_list`
      - `+ att_ck`
    - 响应参数变更
      - `+ data_list.description`
      - `+ data_list.event_abstract`
      - `+ data_list.tag_list`
      - `+ data_list.resource_info.container_status`
      - `+ data_list.resource_info.pod_uid`
      - `+ data_list.resource_info.pod_name`
      - `+ data_list.resource_info.namespace`
      - `+ data_list.resource_info.cluster_id`
      - `+ data_list.resource_info.cluster_name`
  - **ChangeEvent**
    - 请求参数变更
      - `+ operate_event_list.operate_detail_list.container_id`
      - `+ operate_event_list.operate_detail_list.container_name`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListProducts**
    - 请求参数变更
      - `+ product_name`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ResizeInstance**
    - 请求参数变更
      - `+ tenant_ips`
      - `+ second_tenant_subnet_id`
  - **ResizeEngineInstance**
    - 请求参数变更
      - `+ tenant_ips`
      - `+ second_tenant_subnet_id`

### HuaweiCloud SDK LakeFormation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateLakeFormationInstance**
    - 请求参数变更
      - `- order_id`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 支持接口`RunCreateAudioStreamModerationJob`、`RunCloseAudioStreamModerationJob`、`RunCreateVideoStreamModerationJob`、`RunCloseVideoStreamModerationJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeGeneralTable**
    - 请求参数变更
      - `+ with_borders`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListInstanceDiagnosis`、`ListInstancesInfoDiagnosis`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowGroup**
    - 响应参数变更
      - `+ group_desc`
  - **CreateConsumerGroupOrBatchDeleteConsumerGroup**
    - 请求参数变更
      - `+ group_desc`
  - **ListInstanceConsumerGroups**
    - 响应参数变更
      - `+ groups.group_desc`
  - **BatchUpdateConsumerGroup**
    - 请求参数变更
      - `+ groups.group_desc`

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 支持以下接口：
    - `SetHpcCacheBackend`
    - `CreateHpcCacheTask`
    - `ShowHpcCacheTask`
    - `ListHpcCacheTasks`
    - `ListFsTasks`
    - `CreateFsTask`
    - `ShowFsTask`
    - `DeleteFsTask`
    - `ListBackendTargets`
    - `CreateBackendTarget`
    - `ShowBackendTargetInfo`
    - `DeleteBackendTarget`
    - `ShowFsDirUsage`
    - `ListPermRules`
    - `CreatePermRule`
    - `ShowPermRule`
    - `UpdatePermRule`
    - `DeletePermRule`
    - `UpdateHpcShare`
- _解决问题_
  - 无
- _特性变更_
  - **ListShares**
    - 响应参数变更
      - `* shares: list<Shares> -> list<ShareInfo>`

# 0.1.62 2023-10-19

### HuaweiCloud SDK APIG

- _新增特性_
  - 支持以下接口：
    - `CheckApisV2`
    - `ShowAppBoundAppQuota`
    - `CreateOrder`
    - `CreatePrepayResize`
    - `CreatePostPayResizeOrder`
    - `ShowRestrictionOfInstanceV2`
    - `ShowDetailsOfAppAcl`
    - `UpdateAppAcl`
    - `DeleteAppAcl`
    - `ListAppQuotas`
    - `CreateAppQuota`
    - `ShowAppQuota`
    - `UpdateAppQuota`
    - `DeleteAppQuota`
    - `ListAppQuotaBoundApps`
    - `AssociateAppsForAppQuota`
    - `DisassociateAppQuotaWithApp`
    - `ListAppQuotaBindableApps`
    - `UpdateEnvironmentVariableV2`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AOM

- _新增特性_
  - 支持接口`CreateSubApp`、`UpdateSubApp`、`DeleteSubApp`
- _解决问题_
  - 无
- _特性变更_
  - **CreateApp**
    - 请求参数变更
      - `+ register_type: enum value [CONSOLE,SERVICE_DISCOVERY]`
      - `- register_type: enum value [CONSOLESERVICE_DISCOVERY]`
  - **UpdateApp**
    - 请求参数变更
      - `+ register_type: enum value [CONSOLE,SERVICE_DISCOVERY]`
      - `- register_type: enum value [CONSOLESERVICE_DISCOVERY]`
  - **CreateComponent**
    - 请求参数变更
      - `+ model_type: enum value [APPLICATION,SUB_APPLICATION]`
  - **CreateEnv**
    - 请求参数变更
      - `+ env_type: enum value [DEV,TEST,PRE,ONLINE]`
      - `+ os_type: enum value [LINUX,WINDOWS]`
      - `+ register_type: enum value [API,CONSOLE,SERVICE_DISCOVERY]`
  - **ListResourceUnderNode**
    - 请求参数变更
      - `+ ci_type: enum value [APPLICATION,SUB_APPLICATION,COMPONENT,ENVIRONMENT]`
  - **UpdateEnv**
    - 请求参数变更
      - `+ env_type: enum value [DEV,TEST,PRE,ONLINE]`
      - `+ os_type: enum value [LINUX,WINDOWS]`
      - `+ register_type: enum value [API,CONSOLE,SERVICE_DISCOVERY]`

### HuaweiCloud SDK CAE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListComponentConfigurations**
    - 响应参数变更
      - `* items.data.spec.log_paths: list<object> -> list<string>`
      - `* items.data.spec.metrics: list<object> -> list<string>`
      - `+ items.data.spec.triggers.type: enum value [cron]`
      - `+ items.data.spec.triggers.metadata.period_type`
      - `+ items.data.spec.triggers.metadata.schedulers`
      - `- items.data.spec.triggers.metadata.type: enum value [Utilization]`
      - `* items.data.spec.triggers.metadata: object -> object<ScalingTriggerMeta>`
      - `* items.data.spec.postStart.exec.command: list<object> -> list<string>`
      - `* items.data.spec.livenessProbe.exec.command: list<object> -> list<string>`
      - `* items.data.spec.items.access_control.black: list<object> -> list<string>`
      - `* items.data.spec.items.access_control.white: list<object> -> list<string>`
  - **CreateComponentConfiguration**
    - 请求参数变更
      - `* items.data.spec.log_paths: list<object> -> list<string>`
      - `* items.data.spec.metrics: list<object> -> list<string>`
      - `+ items.data.spec.triggers.type: enum value [cron]`
      - `+ items.data.spec.triggers.metadata.period_type`
      - `+ items.data.spec.triggers.metadata.schedulers`
      - `- items.data.spec.triggers.metadata.type: enum value [Utilization]`
      - `* items.data.spec.triggers.metadata: object -> object<ScalingTriggerMeta>`
      - `* items.data.spec.postStart.exec.command: list<object> -> list<string>`
      - `* items.data.spec.livenessProbe.exec.command: list<object> -> list<string>`
      - `* items.data.spec.items.access_control.black: list<object> -> list<string>`
      - `* items.data.spec.items.access_control.white: list<object> -> list<string>`

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAgent**
    - 请求参数变更
      - `* agent_id: string -> list<string>`
  - **ListVault**
    - 请求参数变更
      - `* id: string -> list<string>`

### HuaweiCloud SDK CC

- _新增特性_
  - 支持以下接口：
    - `TagCloudConnection`
    - `UntagCloudConnection`
    - `ListCloudConnectionTags`
    - `ListCloudConnectionsByTags`
    - `TagBandwidthPackage`
    - `UntagBandwidthPackage`
    - `ListBandwidthPackageTags`
    - `ListBandwidthPackagesByTags`
    - `ListCentralNetworks`
    - `CreateCentralNetwork`
    - `ShowCentralNetwork`
    - `UpdateCentralNetwork`
    - `DeleteCentralNetwork`
    - `TagCentralNetwork`
    - `UntagCentralNetwork`
    - `ListCentralNetworkTags`
    - `ListCentralNetworkPolicies`
    - `CreateCentralNetworkPolicy`
    - `ApplyCentralNetworkPolicy`
    - `DeleteCentralNetworkPolicy`
    - `ListCentralNetworkPolicyChangeSet`
    - `ListCentralNetworkConnections`
    - `UpdateCentralNetworkConnection`
    - `ListCentralNetworkGdgwAttachments`
    - `CreateCentralNetworkGdgwAttachment`
    - `ShowCentralNetworkGdgwAttachment`
    - `UpdateCentralNetworkGdgwAttachment`
    - `ListCentralNetworkAttachments`
    - `DeleteCentralNetworkAttachment`
    - `ListCentralNetworkQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `ListDomainTags`
    - `DeleteTag`
    - `BatchCreateDeleteTags`
    - `ListResourceByFilterTag`
    - `ListTags`
    - `CreateTag`
  - **ListCloudConnections**
    - 请求参数变更
      - `* id: list<string> -> list<UUIDDef>`
  - **ListCloudConnectionRoutes**
    - 请求参数变更
      - `* cloud_connection_id: list<string> -> list<UUIDDef>`
  - **ListAuthorisations**
    - 请求参数变更
      - `* id: list<string> -> list<UUIDDef>`
      - `* cloud_connection_id: list<string> -> list<UUIDDef>`
  - **ListPermissions**
    - 请求参数变更
      - `* id: list<string> -> list<UUIDDef>`
      - `* cloud_connection_id: list<string> -> list<UUIDDef>`
  - **ListCloudConnectionQuotas**
    - 请求参数变更
      - `+ cloud_connection_id`
      - `+ region_id`
  - **ListNetworkInstances**
    - 请求参数变更
      - `* id: list<string> -> list<UUIDDef>`
      - `* cloud_connection_id: list<string> -> list<UUIDDef>`
  - **ListBandwidthPackages**
    - 请求参数变更
      - `+ cloud_connection_id`
      - `* id: list<string> -> list<UUIDDef>`
  - **ListInterRegionBandwidths**
    - 请求参数变更
      - `* id: list<string> -> list<UUIDDef>`
      - `* cloud_connection_id: list<string> -> list<UUIDDef>`

### HuaweiCloud SDK CCM

- _新增特性_
  - 支持接口`CreateCertificateAuthorityOrder`
- _解决问题_
  - 无
- _特性变更_
  - **ExportCertificate**
    - 响应参数变更
      - `+ signed_and_enveloped_data`
  - **ShowCertificateAuthority**
    - 响应参数变更
      - `+ charging_mode`
      - `+ free_quota`
  - **IssueCertificateAuthorityCertificate**
    - 请求参数变更
      - `+ type`
      - `+ distinguished_name`
      - `+ key_algorithm`
      - `+ key_usages`
      - `+ crl_configuration`
  - **CreateCertificateAuthority**
    - 请求参数变更
      - `+ ca_id`
  - **ListCertificateAuthority**
    - 响应参数变更
      - `+ charging_mode`
      - `+ free_quota`
      - `+ certificate_authorities.free_quota`
      - `+ certificate_authorities.charging_mode`

### HuaweiCloud SDK CFW

- _新增特性_
  - 支持以下接口：
    - `ListDomainSets`
    - `AddDomainSet`
    - `UpdateDomainSet`
    - `DeleteDomainSet`
    - `ListFirewallList`
    - `BatchUpdateAclRuleActions`
    - `ListRuleAclTags`
    - `AddDomains`
    - `DeleteDomains`
    - `ListDomains`
    - `BatchDeleteAclRules`
    - `BatchDeleteServiceItems`
    - `BatchDeleteAddressItems`
- _解决问题_
  - 无
- _特性变更_
  - **ListFlowLogs**
    - 请求参数变更
      - `+ dst_host`
    - 响应参数变更
      - `+ data.records.dst_host`
  - **ListAccessControlLogs**
    - 请求参数变更
      - `+ dst_host`
      - `+ rule_name`
      - `+ action`
    - 响应参数变更
      - `+ data.records.src_region_id`
      - `+ data.records.src_region_name`
      - `+ data.records.dst_region_id`
      - `+ data.records.dst_region_name`
      - `+ data.records.dst_host`
  - **ListBlackWhiteLists**
    - 响应参数变更
      - `+ data.records.description`
  - **ListDomainParseDetail**
    - 请求参数变更
      - `+ address_type`
  - **UpdateDnsServers**
    - 请求参数变更
      - `+ health_check_domain_name`
  - **ListDnsServers**
    - 响应参数变更
      - `+ data.health_check_domain_name`
  - **ListAttackLogs**
    - 请求参数变更
      - `+ dst_host`
      - `+ log_type`
    - 响应参数变更
      - `+ data.records.dst_host`
      - `+ data.records.src_region_id`
      - `+ data.records.src_region_name`
      - `+ data.records.dst_region_id`
      - `+ data.records.dst_region_name`
  - **UpdateAclRule**
    - 请求参数变更
      - `+ tag`
      - `+ source.region_list_json`
      - `+ source.region_list`
      - `+ source.domain_set_id`
      - `+ source.domain_set_name`
      - `+ source.ip_address`
      - `+ source.address_group`
      - `+ source.address_group_names`
      - `+ service.custom_service`
      - `+ service.service_group`
      - `+ service.service_group_names`
  - **ListAclRules**
    - 请求参数变更
      - `+ tags_id`
      - `+ source`
      - `+ destination`
      - `+ service`
    - 响应参数变更
      - `+ data.records.tag`
      - `+ data.records.source.region_list_json`
      - `+ data.records.source.region_list`
      - `+ data.records.source.domain_set_id`
      - `+ data.records.source.domain_set_name`
      - `+ data.records.source.ip_address`
      - `+ data.records.source.address_group`
      - `+ data.records.source.address_group_names`
      - `+ data.records.service.custom_service`
      - `+ data.records.service.service_group`
      - `+ data.records.service.service_group_names`
  - **AddBlackWhiteList**
    - 请求参数变更
      - `+ description`
  - **UpdateBlackWhiteList**
    - 请求参数变更
      - `+ description`
  - **ListEipCount**
    - 响应参数变更
      - `+ data.eip_protected_self`
  - **ChangeEipStatus**
    - 响应参数变更
      - `+ data.object_id`
      - `+ data.fail_eip_id_list`
      - `- data.id`
      - `* data: object<IdObject> -> object<EIPSwitchStatusVO>`
  - **ListEastWestFirewall**
    - 响应参数变更
      - `+ data.mode`
      - `+ data.ew_vpc_route_limit`
      - `+ data.er_associated_subnet.ipv6_enable`
      - `+ data.protect_infos.protected_resource_mode`
  - **AddAclRule**
    - 请求参数变更
      - `+ rules.tag`
      - `+ rules.source.region_list_json`
      - `+ rules.source.region_list`
      - `+ rules.source.domain_set_id`
      - `+ rules.source.domain_set_name`
      - `+ rules.source.ip_address`
      - `+ rules.source.address_group`
      - `+ rules.source.address_group_names`
      - `+ rules.service.custom_service`
      - `+ rules.service.service_group`
      - `+ rules.service.service_group_names`
  - **ListEips**
    - 请求参数变更
      - `+ tags`
    - 响应参数变更
      - `+ data.records.object_id`
      - `+ data.records.tags`
      - `+ data.records.domain_id`
      - `+ data.records.owner`
      - `+ data.records.fw_domain_id`
  - **AddAddressItem**
    - 响应参数变更
      - `+ data.covered_ip`
  - **ListFirewallDetail**
    - 响应参数变更
      - `+ data.records.resource_id`
      - `+ data.records.support_url_filtering`
      - `+ data.records.flavor.session_concurrent`
      - `+ data.records.flavor.session_create`
      - `+ data.records.flavor.total_rule_count`
      - `+ data.records.flavor.used_rule_count`
      - `+ data.records.flavor.vpc_bandwith`

### HuaweiCloud SDK CodeArtsBuild

- _新增特性_
  - 支持以下接口：
    - `DownloadLogByRecordId`
    - `ShowRecordInfo`
    - `StopBuildJob`
    - `DeleteBuildJob`
    - `DisableBuildJob`
    - `ResumeBuildJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 支持接口`ShowTags`、`ParseUserBehavior`、`ShowDataSets`、`BatchSyncMetadata`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持以下接口：
    - `CreateConnectivityTest`
    - `ShowReplicationStates`
    - `ListAclAccounts`
    - `CreateAclAccount`
    - `UpdateAclAccountPassWord`
    - `ResetAclAccountPassWord`
    - `UpdateAclAccountRole`
    - `UpdateAclAccountRemark`
    - `DeleteAclAccount`
    - `ShowConfigTemplate`
    - `UpdateConfigTemplate`
    - `DeleteConfigTemplate`
- _解决问题_
  - 无
- _特性变更_
  - **ListConfigTemplates**
    - 响应参数变更
      - `+ templates`
      - `- config_templates`
  - **CreateRedislog**
    - 请求参数变更
      - `+ query_time: enum value [0,1,3,7]`
  - **ListInstances**
    - 响应参数变更
      - `+ instances.features.support_audit_log`
  - **ShowInstance**
    - 响应参数变更
      - `+ features.support_audit_log`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ReinstallServerWithCloudInit**
    - 请求参数变更
      - `+ os-reinstall.metadata.BYOL`
  - **ChangeServerOsWithCloudInit**
    - 请求参数变更
      - `+ os-change.metadata.BYOL`
  - **ChangeServerOsWithoutCloudInit**
    - 请求参数变更
      - `+ os-change.metadata.BYOL`

### HuaweiCloud SDK EG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **PutOfficialEvents**
    - 响应参数变更
      - `- failed_count`
      - `- events`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持以下接口：
    - `UpdateProxyNewConfigurations`
    - `CopyConfigurations`
    - `ListConfigurationsDifferences`
    - `ListConfigurationsInstances`
    - `ListModifyHistory`
    - `ListEnterpriseProjects`
    - `SwitchAccessControl`
    - `CreateAccessControl`
    - `DeleteScheduleTasK`
    - `ListInstanceConfigurations`
    - `ShowGaussMySqlIncrementalBackupList`
    - `UpdateBackupOffsitePolicy`
    - `CreateRestoreTables`
- _解决问题_
  - 无
- _特性变更_
  - **ListGaussMySqlDatabase**
    - 请求参数变更
      - `+ name`
      - `+ charset`

### HuaweiCloud SDK LakeFormation

- _新增特性_
  - 支持以下接口：
    - `UpdateLakeFormationInstanceDefault`
    - `UpdateLakeFormationInstanceScale`
    - `ShowAccessService`
    - `AuthorizeAccessService`
    - `ListAccessClientInfos`
    - `CreateAccessClient`
    - `ShowAccessClient`
    - `UpdateAccessClient`
    - `DeleteAccessClient`
    - `ShowAgency`
    - `CreateAgency`
    - `DeleteAgency`
    - `ListFunctionNames`
    - `BatchCheckPermission`
    - `ListPrincipals`
    - `AssociatePrincipals`
    - `RevokePrincipals`
    - `UpdatePrincipals`
    - `ShowCredential`
    - `ListConfigs`
    - `ListUsers`
    - `AssociateRoles`
    - `RevokeRoles`
    - `UpdateRoles`
    - `ListUserRoles`
    - `ListPartitionNamesWithoutLimit`
    - `DeleteAgreement`
- _解决问题_
  - 无
- _特性变更_
  - **ShowDatabase**
    - 响应参数变更
      - `+ owner_auth_source_type`
  - **UpdateDatabase**
    - 请求参数变更
      - `+ owner_auth_source_type`
    - 响应参数变更
      - `+ owner_auth_source_type`
  - **ListCatalogs**
    - 响应参数变更
      - `+ owner`
      - `+ owner_source`
      - `+ owner_type`
  - **CreateCatalog**
    - 请求参数变更
      - `+ branch_name`
      - `+ owner`
      - `+ owner_type`
      - `+ owner_source`
    - 响应参数变更
      - `+ owner`
      - `+ owner_source`
      - `+ owner_type`
  - **ShowCatalog**
    - 响应参数变更
      - `+ owner`
      - `+ owner_source`
      - `+ owner_type`
  - **UpdateCatalog**
    - 请求参数变更
      - `+ branch_name`
      - `+ owner`
      - `+ owner_type`
      - `+ owner_source`
    - 响应参数变更
      - `+ owner`
      - `+ owner_source`
      - `+ owner_type`
  - **ShowRole**
    - 响应参数变更
      - `+ role_name`
      - `+ principal_source`
      - `+ description`
      - `- role`
      - `- user_roles`
  - **UpdateRole**
    - 响应参数变更
      - `+ principal_source: enum value [AGENTTENANT]`
  - **ListLakeFormationInstances**
    - 响应参数变更
      - `+ default_instance`
      - `+ instances.default_instance`
      - `+ instances.status: enum value [RESOURCE_PREPARATION,RUNNING,RESOURCE_RELEASE,DELETED,RESOURCE_PREPARATION_FAIL,FROZEN_RELEASABLE,FROZEN_UNRELEASABLE,RECOVERING,DELETING,SCALING,SCALE_FAIL]`
  - **CreateLakeFormationInstance**
    - 请求参数变更
      - `+ order_id`
      - `+ charge_mode: enum value [prePaid]`
      - `+ specs.product_id`
    - 响应参数变更
      - `+ scale_progress`
      - `+ charge_mode`
      - `+ default_instance`
      - `+ resource_progress`
      - `+ resource_expected_duration`
      - `+ scale_expected_duration`
      - `+ status: enum value [RESOURCE_PREPARATION,RUNNING,RESOURCE_RELEASE,DELETED,RESOURCE_PREPARATION_FAIL,FROZEN_RELEASABLE,FROZEN_UNRELEASABLE,RECOVERING,DELETING,SCALING,SCALE_FAIL]`
      - `+ specs.product_id`
  - **UpdateLakeFormationInstance**
    - 响应参数变更
      - `+ default_instance`
      - `+ status: enum value [RESOURCE_PREPARATION,RUNNING,RESOURCE_RELEASE,DELETED,RESOURCE_PREPARATION_FAIL,FROZEN_RELEASABLE,FROZEN_UNRELEASABLE,RECOVERING,DELETING,SCALING,SCALE_FAIL]`
  - **ShowLakeFormationInstance**
    - 响应参数变更
      - `+ scale_progress`
      - `+ charge_mode`
      - `+ default_instance`
      - `+ resource_progress`
      - `+ resource_expected_duration`
      - `+ scale_expected_duration`
      - `+ status: enum value [RESOURCE_PREPARATION,RUNNING,RESOURCE_RELEASE,DELETED,RESOURCE_PREPARATION_FAIL,FROZEN_RELEASABLE,FROZEN_UNRELEASABLE,RECOVERING,DELETING,SCALING,SCALE_FAIL]`
      - `+ specs.product_id`
  - **ListSpecs**
    - 响应参数变更
      - `+ spec_codes.usage_value`
      - `+ spec_codes.free_usage_value`
  - **CreateDatabase**
    - 请求参数变更
      - `+ owner_auth_source_type`
    - 响应参数变更
      - `+ owner_auth_source_type`
  - **ListDatabases**
    - 响应参数变更
      - `+ owner_auth_source_type`
      - `+ databases.owner_auth_source_type`
  - **ShowFunction**
    - 响应参数变更
      - `+ owner_auth_source_type`
  - **UpdateFunction**
    - 请求参数变更
      - `+ owner_auth_source_type`
    - 响应参数变更
      - `+ owner_auth_source_type`
  - **CreateRole**
    - 响应参数变更
      - `+ principal_source: enum value [AGENTTENANT]`
  - **ListRoles**
    - 响应参数变更
      - `+ roles.principal_source: enum value [AGENTTENANT]`
  - **ListAllFunctions**
    - 响应参数变更
      - `+ owner_auth_source_type`
      - `+ functions.owner_auth_source_type`
  - **CreateFunction**
    - 请求参数变更
      - `+ owner_auth_source_type`
    - 响应参数变更
      - `+ owner_auth_source_type`
  - **ListFunctions**
    - 响应参数变更
      - `+ owner_auth_source_type`
      - `+ functions.owner_auth_source_type`
  - **CreateTable**
    - 请求参数变更
      - `+ ignore_obs_checked`
  - **UpdateTable**
    - 请求参数变更
      - `+ table.ignore_obs_checked`
  - **ShowSyncPolicy**
    - 响应参数变更
      - `+ policy_deltas.change_type`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListStructuredLogsWithTimeRange**
    - 响应参数变更
      - `+ result`
      - `- context`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListEngineProducts**
    - 响应参数变更
      - `+ products.properties.product_alias`

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`RunAudioAssessment`、`RunMultiModalAssessment`

### HuaweiCloud SDK VPC

- _新增特性_
  - 支持接口`BatchCreateSecurityGroupRules`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.61 2023-10-12

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持以下接口：
    - `ListStackSets`
    - `CreateStackSet`
    - `ShowStackSetTemplate`
    - `ListStackSetOperations`
    - `ShowStackSetMetadata`
    - `ListStackInstances`
    - `CreateStackInstance`
    - `DeleteStackInstance`
    - `DeployStackSet`
    - `DeleteStackSet`
    - `UpdateStackSet`
    - `ShowStackSetOperationMetadata`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateScalingConfig**
    - 请求参数变更
      - `+ source_scaling_configuration_id`

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateBareMetalServers**
    - 请求参数变更
      - `* server.server_tags: map<string, list<SystemTags>> -> list<SystemTags>`

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListSubCustomerNewTag`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CAE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListApplications**
    - 响应参数变更
      - `- items.annotations`
  - **CreateApplication**
    - 请求参数变更
      - `- metadata.annotations`
    - 响应参数变更
      - `- metadata.annotations`
  - **ShowApplication**
    - 响应参数变更
      - `- metadata.annotations`
  - **ListComponentConfigurations**
    - 响应参数变更
      - `+ items.data.spec`
      - `+ items.data.metadata`
      - `* items.data: object -> object<ListComponentConfigurationsResponseData>`
  - **CreateComponentConfiguration**
    - 请求参数变更
      - `+ items.data.spec`
      - `+ items.data.metadata`
      - `* items.data: object -> object<ConfigurationData>`
  - **ListComponentSnapshots**
    - 响应参数变更
      - `- items.description`
      - `+ items.context.runtime: enum value [Java17,Nodejs14,Nodejs16]`
  - **ShowComponent**
    - 响应参数变更
      - `+ spec.runtime: enum value [Java17,Nodejs14,Nodejs16]`
  - **UpdateComponent**
    - 请求参数变更
      - `+ spec.runtime: enum value [Java17,Nodejs14,Nodejs16]`
  - **CreateComponent**
    - 响应参数变更
      - `+ spec.runtime: enum value [Java17,Nodejs14,Nodejs16]`
  - **ListComponents**
    - 响应参数变更
      - `+ items.spec.runtime: enum value [Java17,Nodejs14,Nodejs16]`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowHistoryTasks**
    - 请求参数变更
      - `+ task_type`
  - **ShowUrlTaskInfo**
    - 响应参数变更
      - `+ result.mode`

### HuaweiCloud SDK CES

- _新增特性_
  - 支持以下接口：
    - `BatchUpdateNotificationMasks`
    - `BatchUpdateNotificationMaskTime`
    - `UpdateNotificationMasks`
    - `BatchDeleteNotificationMasks`
    - `ListNotificationMasks`
    - `ListNotificationMaskResources`
    - `ListOneClickAlarms`
    - `CreateOneClickAlarm`
    - `ListOneClickAlarmRules`
    - `BatchUpdateOneClickAlarmsEnabledState`
    - `BatchDeleteOneClickAlarms`
    - `UpdateOneClickAlarmNotifications`
    - `BatchUpdateOneClickAlarmPoliciesEnabledState`
    - `UpdateAlarmNotifications`
    - `ListCesTargetProjectTags`
- _解决问题_
  - 无
- _特性变更_
  - **ListAlarmHistories**
    - 响应参数变更
      - `+ alarm_histories.condition.suppress_duration: enum value [86400]`
  - **ListAgentInvocations**
    - 请求参数变更
      - `- instance_name`
      - `+ invocation_type: enum value [RETRY]`
    - 响应参数变更
      - `+ invocations.invocation_type: enum value [RETRY]`
  - **ListAgentStatus**
    - 响应参数变更
      - `+ agent_status.extensions.version`

### HuaweiCloud SDK CodeArtsDeploy

- _新增特性_
  - 支持接口`ShowExecutionParams`
- _解决问题_
  - 无
- _特性变更_
  - **ListAllApp**
    - 请求参数变更
      - `+ states`
      - `+ group_id`

### HuaweiCloud SDK Config

- _新增特性_
  - 支持以下接口：
    - `ListOrganizationConformancePacks`
    - `CreateOrganizationConformancePack`
    - `ShowOrganizationConformancePack`
    - `DeleteOrganizationConformancePack`
    - `ListOrganizationConformancePackStatuses`
    - `ShowOrganizationConformancePackDetailedStatuses`
- _解决问题_
  - 无
- _特性变更_
  - **ShowConfigurationAggregatorSourcesStatus**
    - 响应参数变更
      - `+ aggregated_source_statuses.source_name`
  - **ShowConformancePack**
    - 响应参数变更
      - `+ created_by`
  - **CreateConformancePack**
    - 响应参数变更
      - `+ created_by`
  - **ListConformancePacks**
    - 响应参数变更
      - `+ created_by`
      - `+ value.created_by`
  - **ShowAggregatePolicyStateComplianceSummary**
    - 响应参数变更
      - `+ results.group_account_name`
  - **ListAggregateComplianceByPolicyAssignment**
    - 响应参数变更
      - `+ aggregate_policy_assignments.account_name`

### HuaweiCloud SDK CSS

- _新增特性_
  - 支持以下接口：
    - `DeleteConfig`
    - `StopHotPipeline`
    - `ListCerts`
    - `ListElbs`
    - `EnableOrDisableElb`
    - `ShowElbDetail`
    - `CreateElbListener`
    - `UpdateEsListener`
    - `ListElbCerts`
    - `ListAiOps`
    - `CreateAiOps`
    - `DeleteAiOps`
    - `ListSmnTopics`
    - `UpgradeCore`
    - `ListImages`
    - `UpgradeDetail`
    - `RetryUpgradeTask`
    - `UpdateAzByInstanceType`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateShrinkNodes**
    - 请求参数变更
      - `+ migrate_data`
  - **CreateCnf**
    - 请求参数变更
      - `+ sensitive_words`
  - **UpdateCnf**
    - 请求参数变更
      - `+ sensitive_words`
  - **ShowClusterDetail**
    - 响应参数变更
      - `+ bandwidthResourceId`
      - `+ instances.resourceId`
      - `+ instances.volume.resourceIds`
      - `+ publicKibanaResp.bandwidthResourceId`
  - **ListClustersDetails**
    - 响应参数变更
      - `+ clusters.bandwidthResourceId`
      - `+ clusters.instances.resourceId`
      - `+ clusters.instances.volume.resourceIds`
      - `+ clusters.publicKibanaResp.bandwidthResourceId`

### HuaweiCloud SDK CTS

- _新增特性_
  - 支持以下接口：
    - `ListOperations`
    - `BatchCreateResourceTags`
    - `BatchDeleteResourceTags`
    - `ListUserResources`
    - `CheckObsBuckets`
    - `ListTraceResources`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DataArtsStudio

- _新增特性_
  - 支持以下接口：
    - `StopFactorySupplementDataInstance`
    - `ShowFactorySupplementData`
    - `CreateFactorySupplementDataInstance`
    - `ShowFactoryEnv`
    - `CreateFactoryEnv`
- _解决问题_
  - 无
- _特性变更_
  - **PublishApi**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ShowApplyDetail**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ShowMessageDetail**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ShowCatalogDetail**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **UpdateCatalog**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **CreateServiceCatalog**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **DeleteServiceCatalog**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **MigrateCatalog**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **MigrateApi**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **SearchIdByPath**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ShowPathById**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **PublishApiToInstance**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ExecuteApiToInstance**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **AuthorizeApiToInstance**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **AuthorizeActionApiToInstance**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **DeleteApp**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowAppInfo**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **UpdateApp**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowApisOverview**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowAppsOverview**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowApisDetail**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowAppsDetail**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **BatchApproveApply**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ListApply**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ConfirmMessage**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ListMessage**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ListAllCatalogList**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ListApiCatalogList**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ListCatalogList**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **ShowPathObjectById**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
    - 响应参数变更
      - `+ paths.name`
  - **DebugApi**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **SearchPublishInfo**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ListInstanceList**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **SearchDebugInfo**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ListApicInstances**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ListApicGroups**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **CreateApp**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ListApps**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ListApisTop**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ListAppsTop**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowApisDashboard**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowApiDashboard**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowAppsDashboard**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ListApiTopN**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ListApis**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **CreateApi**
    - 请求参数变更
      - `* workspace: optional -> required`
  - **ShowApi**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **UpdateApi**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **SearchAuthorizeApp**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`
  - **SearchBindApi**
    - 请求参数变更
      - `+ Dlm-Type`
      - `* workspace: optional -> required`

### HuaweiCloud SDK DC

- _新增特性_
  - 支持以下接口：
    - `UpdateVifPeer`
    - `DeleteVifPeer`
    - `CreateVifPeer`
    - `ShowQuotas`
    - `ListSwitchoverTestRecords`
    - `SwitchoverTest`
- _解决问题_
  - 无
- _特性变更_
  - **DeleteResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [dc-lag]`
  - **ListProjectTags**
    - 请求参数变更
      - `+ resource_type: enum value [dc-lag]`
  - **ShowResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [dc-lag]`
  - **CreateResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [dc-lag]`
  - **BatchCreateResourceTags**
    - 请求参数变更
      - `+ resource_type: enum value [dc-lag]`
  - **ShowDirectConnect**
    - 请求参数变更
      - `- limit`
      - `- marker`
    - 响应参数变更
      - `- direct_connect.type: enum value [onestop_standard,onestop_hosted]`
      - `- direct_connect.charge_mode: enum value [port]`
  - **UpdateDirectConnect**
    - 响应参数变更
      - `- direct_connect.type: enum value [onestop_standard,onestop_hosted]`
      - `- direct_connect.charge_mode: enum value [port]`
  - **ListDirectConnects**
    - 响应参数变更
      - `- direct_connects.type: enum value [onestop_standard,onestop_hosted]`
      - `- direct_connects.charge_mode: enum value [port]`
  - **ShowVirtualGateway**
    - 响应参数变更
      - `- virtual_gateway.type: enum value [default]`
  - **UpdateVirtualGateway**
    - 响应参数变更
      - `- virtual_gateway.type: enum value [default]`
  - **ListVirtualGateways**
    - 请求参数变更
      - `+ enterprise_project_id`
    - 响应参数变更
      - `- virtual_gateways.type: enum value [default]`
  - **CreateVirtualGateway**
    - 响应参数变更
      - `- virtual_gateway.type: enum value [default]`
  - **ShowVirtualInterface**
    - 响应参数变更
      - `+ virtual_interface.service_type: enum value [GDGW]`
      - `- virtual_interface.service_type: enum value [vpc,GDWW]`
  - **UpdateVirtualInterface**
    - 响应参数变更
      - `+ virtual_interface.service_type: enum value [GDGW]`
      - `- virtual_interface.service_type: enum value [vpc,GDWW]`
  - **ListVirtualInterfaces**
    - 响应参数变更
      - `+ virtual_interfaces.service_type: enum value [GDGW]`
      - `- virtual_interfaces.service_type: enum value [vpc,GDWW]`
  - **CreateVirtualInterface**
    - 请求参数变更
      - `+ virtual_interface.service_type: enum value [GDGW]`
      - `- virtual_interface.service_type: enum value [GDWW]`
    - 响应参数变更
      - `+ virtual_interface.service_type: enum value [GDGW]`
      - `- virtual_interface.service_type: enum value [vpc,GDWW]`
  - **ListTagResourceInstances**
    - 请求参数变更
      - `+ resource_type: enum value [dc-lag]`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RunOnce**
    - 响应参数变更
      - `* instanceId: string -> int64`

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持以下接口：
    - `SaveClusterDescriptionInfo`
    - `ExecuteDatabaseOmUserAction`
    - `ShowInstance`
    - `ShowDatabaseOmUserStatus`
    - `ListConfigurationsAuditRecords`
    - `ShowClusterRedistribution`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ReinstallServerWithoutCloudInit**
    - 请求参数变更
      - `+ os-reinstall.metadata.BYOL`
  - **ListFlavors**
    - 响应参数变更
      - `+ flavors.os_extra_specs.quota:vif_max_num`
      - `+ flavors.os_extra_specs.quota:sub_network_interface_max_num`
  - **ListResizeFlavors**
    - 响应参数变更
      - `+ flavors.extra_specs.quota:vif_max_num`
      - `+ flavors.extra_specs.quota:sub_network_interface_max_num`

### HuaweiCloud SDK EG

- _新增特性_
  - 支持接口`PutOfficialEvents`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowFunctionUrl`、`UpdateFunctionUrl`、`CreateFunctionUrl`、`DeleteFunctionUrl`
  - **ListAsyncInvocations**
    - 响应参数变更
      - `+ next_marker`
      - `+ count`
  - **ListActiveAsyncInvocations**
    - 响应参数变更
      - `+ next_marker`
      - `+ count`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstances**
    - 响应参数变更
      - `+ instances.datastore.complete_version`
      - `+ instances.datastore.hotfix_versions`
  - **ListInstancesDetails**
    - 响应参数变更
      - `+ instances.datastore.complete_version`
      - `+ instances.datastore.hotfix_versions`

### HuaweiCloud SDK ImageSearch

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RunAddData**
    - 请求参数变更
      - `* optional_params.category: int -> int32`
    - 响应参数变更
      - `* data.image_info.objects.category: number -> integer`
  - **RunDeleteData**
    - 响应参数变更
      - `* data.delete_info.total_num: int -> int32`
      - `* data.delete_info.delete_num: int -> int32`
  - **RunSearch**
    - 请求参数变更
      - `* optional_params.category: int -> int32`
    - 响应参数变更
      - `* data.image_info.category: number -> integer`
      - `* data.image_info.objects.category: number -> integer`
      - `* data.search_info.total_num: int -> int32`
      - `* data.search_info.return_num: int -> int32`
      - `* data.search_info.search_time: long -> int32`
  - **RunCheckData**
    - 响应参数变更
      - `* data.check_info.total_num: int -> int32`
      - `* data.check_info.return_num: int -> int32`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowJob**
    - 响应参数变更
      - `+ entities.addition_error_code`
      - `+ entities.addition_error_msg`
      - `+ entities.error_code`
      - `+ entities.error`
      - `+ entities.alarm_code`

### HuaweiCloud SDK LakeFormation

- _新增特性_
  - 支持接口`ListLakeFormationInstanceTags`、`ListQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MetaStudio

- _新增特性_
  - 支持接口`CreatePhotoDetection`、`ShowPhotoDetection`
- _解决问题_
  - 无
- _特性变更_
  - **CreateSmartLiveRoom**
    - 请求参数变更
      - `- video_config.codec: enum value [VP9]`
  - **ShowSmartLiveRoom**
    - 响应参数变更
      - `- video_config.codec: enum value [VP9]`
  - **UpdateSmartLiveRoom**
    - 请求参数变更
      - `- video_config.codec: enum value [VP9]`
    - 响应参数变更
      - `- video_config.codec: enum value [VP9]`
  - **StartSmartLive**
    - 请求参数变更
      - `- video_config.codec: enum value [VP9]`
  - **Create2DDigitalHumanVideo**
    - 请求参数变更
      - `- video_config.codec: enum value [VP9]`
  - **Show2DDigitalHumanVideo**
    - 响应参数变更
      - `- video_config.codec: enum value [VP9]`
  - **CreateVideoScripts**
    - 请求参数变更
      - `- video_config.codec: enum value [VP9]`
  - **ShowVideoScript**
    - 响应参数变更
      - `- video_config.codec: enum value [VP9]`
  - **UpdateVideoScript**
    - 请求参数变更
      - `- video_config.codec: enum value [VP9]`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeColombiaIdCard`
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeVehicleLicense**
    - 响应参数变更
      - `+ result.energy_type`
      - `+ result.front`
      - `+ result.back`
  - **RecognizeWebImage**
    - 请求参数变更
      - `+ detect_text_direction`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持以下接口：
    - `ListPostgresqlHbaInfo`
    - `ModifyPostgresqlHbaConf`
    - `AddPostgresqlHbaConf`
    - `DeletePostgresqlHbaConf`
    - `ListPostgresqlHbaInfoHistory`
- _解决问题_
  - 无
- _特性变更_
  - **UpgradeDbVersionNew**
    - 请求参数变更
      - `+ is_delayed`
      - `- delay`

### HuaweiCloud SDK ServiceStage

- _新增特性_
  - 支持以下接口：
    - `ListAuthorizations`
    - `ShowRedirectUrl`
    - `CreateOAuth`
    - `CreatePersonalAuth`
    - `CreatePasswordAuth`
    - `DeleteAuthorize`
    - `ListNamespaces`
    - `ShowProjectDetail`
    - `ListProjects`
    - `CreateProject`
    - `ListBranches`
    - `ListTags`
    - `CreateTag`
    - `DeleteTag`
    - `ListCommits`
    - `ListHooks`
    - `CreateHook`
    - `DeleteHook`
    - `ListTrees`
    - `ShowContent`
    - `UpdateFile`
    - `CreateFile`
    - `DeleteFile`
    - `ListTemplates`
    - `ListRuntimes`
    - `ListFlavors`
    - `ListEnvironments`
    - `CreateEnvironment`
    - `ShowEnvironmentDetail`
    - `ChangeEnvironment`
    - `DeleteEnvironment`
    - `ListApplications`
    - `CreateApplication`
    - `ShowApplicationDetail`
    - `ChangeApplication`
    - `DeleteApplication`
    - `ShowApplicationConfiguration`
    - `ChangeApplicationConfiguration`
    - `DeleteApplicationConfiguration`
    - `ListComponents`
    - `CreateComponent`
    - `ShowComponentDetail`
    - `ChangeComponent`
    - `DeleteComponent`
    - `ListInstances`
    - `CreateInstance`
    - `ShowInstanceDetail`
    - `ChangeInstance`
    - `DeleteInstance`
    - `UpdateInstanceAction`
    - `ListInstanceSnapshots`
    - `ShowJobDetail`
- _解决问题_
  - 无
- _特性变更_
  - **ModifyComponent**
    - 请求参数变更
      - `+ external_accesses.protocol`
      - `- external_accesses.prorocol`
  - **CreateComponent**
    - 请求参数变更
      - `+ external_accesses`
  - **ShowComponentsInApplication**
    - 响应参数变更
      - `+ components.external_accesses.protocol`
      - `- components.external_accesses.prorocol`
  - **ShowComponents**
    - 响应参数变更
      - `+ components.external_accesses.protocol`
      - `- components.external_accesses.prorocol`

### HuaweiCloud SDK SMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowCertKey**
    - 请求参数变更
      - `+ enable_ca_cert`
    - 响应参数变更
      - `+ target_mgmt_private_key`
      - `+ target_data_cert`
      - `+ target_data_private_key`
      - `+ target_mgmt_cert`
      - `+ ca`

### HuaweiCloud SDK VPC

- _新增特性_
  - 支持以下接口：
    - `ListTrafficMirrorSessions`
    - `CreateTrafficMirrorSession`
    - `ShowTrafficMirrorSession`
    - `UpdateTrafficMirrorSession`
    - `DeleteTrafficMirrorSession`
    - `RemoveSourcesFromTrafficMirrorSession`
    - `AddSourcesToTrafficMirrorSession`
    - `ListTrafficMirrorFilters`
    - `CreateTrafficMirrorFilter`
    - `ShowTrafficMirrorFilter`
    - `UpdateTrafficMirrorFilter`
    - `DeleteTrafficMirrorFilter`
    - `ListTrafficMirrorFilterRules`
    - `CreateTrafficMirrorFilterRule`
    - `ShowTrafficMirrorFilterRule`
    - `UpdateTrafficMirrorFilterRule`
    - `DeleteTrafficMirrorFilterRule`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK WorkspaceApp

- _新增特性_
  - 支持以下接口：
    - `CreateOrUpdateStoragePolicyStatement`
    - `ShowPublishableApp`
    - `UploadAppIcon`
    - `ListSessionByUserName`
    - `LogoffUserSession`
    - `ListSessionType`
- _解决问题_
  - 无
- _特性变更_
  - **ListStoragePolicyStatement**
    - 响应参数变更
      - `+ roam_actions`
      - `+ actions`
      - `+ policy_statement_id`
      - `+ items.roam_actions`
  - **UpdateAppGroup**
    - 响应参数变更
      - `+ app_type`
  - **ListAppConnection**
    - 请求参数变更
      - `* limit: required -> optional`
      - `* offset: required -> optional`
  - **ListUserConnection**
    - 请求参数变更
      - `* limit: required -> optional`
      - `* offset: required -> optional`
  - **UpdateServerGroup**
    - 请求参数变更
      - `+ storage_mount_policy`
      - `+ app_type`
      - `+ route_policy.cpu_threshold`
      - `+ route_policy.mem_threshold`
  - **ListProduct**
    - 响应参数变更
      - `+ products.expire_time`
      - `+ products.support_gpu_type`
  - **CreateAppGroup**
    - 请求参数变更
      - `+ app_type`
    - 响应参数变更
      - `+ app_type`
  - **ListAppGroup**
    - 请求参数变更
      - `+ app_type`
      - `* limit: required -> optional`
      - `* offset: required -> optional`
    - 响应参数变更
      - `+ app_type`
      - `+ items.app_type`
  - **ListPublishedApp**
    - 请求参数变更
      - `* limit: required -> optional`
      - `* offset: required -> optional`
  - **PublishApp**
    - 请求参数变更
      - `+ accounts.telephone_number`
  - **AddAppGroupAuthorization**
    - 请求参数变更
      - `+ accounts.telephone_number`
  - **ListAppGroupAuthorization**
    - 请求参数变更
      - `* limit: required -> optional`
      - `* offset: required -> optional`
  - **BatchDeleteAppGroupAuthorization**
    - 请求参数变更
      - `+ accounts.telephone_number`
  - **ListStorageAssignment**
    - 响应参数变更
      - `+ roam_actions`
      - `+ actions`
      - `+ policy_statement_id`
      - `+ items.policy_statement.roam_actions`
  - **ListServers**
    - 请求参数变更
      - `* offset: required -> optional`
      - `* limit: required -> optional`
    - 响应参数变更
      - `+ items.product_info.expire_time`
      - `+ items.product_info.support_gpu_type`
      - `+ items.vm_status: enum value [BUILD_IMAGE]`
      - `+ items.task_status: enum value [build_image]`
  - **CreateServerGroup**
    - 请求参数变更
      - `+ app_type`
      - `+ extra_session_type`
      - `+ extra_session_size`
      - `+ route_policy.cpu_threshold`
      - `+ route_policy.mem_threshold`
    - 响应参数变更
      - `+ app_type`
      - `+ extra_session_size`
      - `+ extra_session_type`
      - `+ storage_mount_policy`
      - `+ product_info.expire_time`
      - `+ product_info.support_gpu_type`
  - **ListServerGroups**
    - 请求参数变更
      - `+ app_type`
      - `* offset: required -> optional`
      - `* limit: required -> optional`
    - 响应参数变更
      - `+ app_type`
      - `+ extra_session_size`
      - `+ extra_session_type`
      - `+ storage_mount_policy`
      - `+ items.extra_session_type`
      - `+ items.extra_session_size`
      - `+ items.app_type`
      - `+ items.storage_mount_policy`
      - `+ items.product_info.expire_time`
      - `+ items.product_info.support_gpu_type`
  - **ListPolicyGroup**
    - 请求参数变更
      - `* offset: required -> optional`
      - `* limit: required -> optional`
  - **ListPolicyTemplate**
    - 请求参数变更
      - `* offset: required -> optional`
      - `* limit: required -> optional`

# 0.1.60 2023-09-19

### HuaweiCloud SDK IdentityCenterStore

- _新增特性_
  - 支持服务`IdentityCenterStore`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateBareMetalServers**
    - 请求参数变更
      - `* server.server_tags: list<SystemTags> -> map<string, list<SystemTags>>`

### HuaweiCloud SDK CAE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateComponentConfiguration**
    - 请求参数变更
      - `+ items.type: enum value [customMetric]`

### HuaweiCloud SDK CPH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`BatchMigrateCloudPhone`、`CreateCloudPhoneServer`
  - **PushShareApps**
    - 响应参数变更
      - `+ jobs`
      - `+ request_id`
  - **DeleteShareApps**
    - 响应参数变更
      - `+ jobs`
      - `+ request_id`
  - **PushShareFiles**
    - 响应参数变更
      - `+ jobs`
      - `+ request_id`
  - **ListCloudPhones**
    - 响应参数变更
      - `+ count`
  - **ShowCloudPhoneDetail**
    - 响应参数变更
      - `+ access_infos.phone_ipv6`
      - `+ access_infos.server_ipv6`
  - **ListCloudPhoneServerModels**
    - 响应参数变更
      - `+ server_models.free_size`
  - **CreateNet2CloudPhoneServer**
    - 请求参数变更
      - `+ nics.ipv6_enable`
      - `+ nics.ipv6_bandwidth`
  - **ListEncodeServers**
    - 响应参数变更
      - `+ encode_servers.encode_server_ipv6`
      - `+ encode_servers.access_infos.server_ipv6`
  - **ListCloudPhoneServers**
    - 响应参数变更
      - `+ count`

### HuaweiCloud SDK CTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateTracker**
    - 请求参数变更
      - `+ is_organization_tracker`
      - `+ management_event_selector`
  - **CreateTracker**
    - 请求参数变更
      - `+ is_organization_tracker`
      - `+ management_event_selector`
    - 响应参数变更
      - `+ is_organization_tracker`
      - `+ management_event_selector`
  - **ListTrackers**
    - 响应参数变更
      - `+ trackers.is_organization_tracker`
      - `+ trackers.management_event_selector`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持接口`ListActiveAsyncInvocations`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateInstance**
    - 请求参数变更
      - `+ availability_zone_detail`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstanceTopics**
    - 响应参数变更
      - `+ topic_max_partitions`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateTranscodingsTemplate**
    - 请求参数变更
      - `+ quality_info.bitrate_adaptive`
      - `+ quality_info.i_frame_policy`
  - **CreateTranscodingsTemplate**
    - 请求参数变更
      - `+ quality_info.bitrate_adaptive`
      - `+ quality_info.i_frame_policy`
  - **ShowTranscodingsTemplate**
    - 响应参数变更
      - `+ templates.quality_info.bitrate_adaptive`
      - `+ templates.quality_info.i_frame_policy`

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持以下接口：
    - `ShowLogConvergeConfig`
    - `ShowAdminConfig`
    - `UpdateSwitch`
    - `ShowMemberGroupAndStream`
    - `UpdateLogConvergeConfig`
- _解决问题_
  - 无
- _特性变更_
  - **ListCharts**
    - 响应参数变更
      - `+ config.can_sort`
      - `+ config.can_search`
      - `+ config.page_size`
      - `* config: object -> object<ChartConfig>`
  - **ShowNotificationTemplate**
    - 响应参数变更
      - `+ create_time`
      - `+ project_id`
      - `+ templates`
      - `+ modify_time`
      - `+ name`
      - `+ source`
      - `+ type`
      - `+ locale`
      - `+ desc`
      - `- body`
  - **DeleteTransfer**
    - 响应参数变更
      - `+ log_transfer_info.log_transfer_detail.obs_period`
      - `+ log_transfer_info.log_transfer_detail.obs_encrypted_id`
      - `+ log_transfer_info.log_transfer_detail.obs_prefix_name`
      - `+ log_transfer_info.log_transfer_detail.obs_period_unit`
      - `+ log_transfer_info.log_transfer_detail.obs_transfer_path`
      - `+ log_transfer_info.log_transfer_detail.obs_eps_id`
      - `+ log_transfer_info.log_transfer_detail.obs_bucket_name`
      - `+ log_transfer_info.log_transfer_detail.obs_encrypted_enable`
      - `+ log_transfer_info.log_transfer_detail.obs_dir_pre_fix_name`
      - `+ log_transfer_info.log_transfer_detail.dis_id`
      - `+ log_transfer_info.log_transfer_detail.dis_name`
      - `+ log_transfer_info.log_transfer_detail.kafka_id`
      - `+ log_transfer_info.log_transfer_detail.kafka_topic`
      - `+ log_transfer_info.log_transfer_detail.obs_time_zone`
      - `+ log_transfer_info.log_transfer_detail.obs_time_zone_id`
      - `+ log_transfer_info.log_transfer_detail.tags`
      - `* log_transfer_info.log_transfer_detail: object -> object<TransferDetail>`
  - **ListTransfers**
    - 响应参数变更
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_period`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_encrypted_id`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_prefix_name`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_period_unit`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_transfer_path`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_eps_id`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_bucket_name`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_encrypted_enable`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_dir_pre_fix_name`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.dis_id`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.dis_name`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.kafka_id`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.kafka_topic`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_time_zone`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.obs_time_zone_id`
      - `+ log_transfers.log_transfer_info.log_transfer_detail.tags`
      - `* log_transfers.log_transfer_info.log_transfer_detail: object -> object<TransferDetail>`
  - **UpdateTransfer**
    - 响应参数变更
      - `+ log_transfer_info.log_transfer_detail.obs_period`
      - `+ log_transfer_info.log_transfer_detail.obs_encrypted_id`
      - `+ log_transfer_info.log_transfer_detail.obs_prefix_name`
      - `+ log_transfer_info.log_transfer_detail.obs_period_unit`
      - `+ log_transfer_info.log_transfer_detail.obs_transfer_path`
      - `+ log_transfer_info.log_transfer_detail.obs_eps_id`
      - `+ log_transfer_info.log_transfer_detail.obs_bucket_name`
      - `+ log_transfer_info.log_transfer_detail.obs_encrypted_enable`
      - `+ log_transfer_info.log_transfer_detail.obs_dir_pre_fix_name`
      - `+ log_transfer_info.log_transfer_detail.dis_id`
      - `+ log_transfer_info.log_transfer_detail.dis_name`
      - `+ log_transfer_info.log_transfer_detail.kafka_id`
      - `+ log_transfer_info.log_transfer_detail.kafka_topic`
      - `+ log_transfer_info.log_transfer_detail.obs_time_zone`
      - `+ log_transfer_info.log_transfer_detail.obs_time_zone_id`
      - `+ log_transfer_info.log_transfer_detail.tags`
      - `* log_transfer_info.log_transfer_detail: object -> object<TransferDetail>`
  - **CreateTransfer**
    - 响应参数变更
      - `+ log_transfer_info.log_transfer_detail.obs_period`
      - `+ log_transfer_info.log_transfer_detail.obs_encrypted_id`
      - `+ log_transfer_info.log_transfer_detail.obs_prefix_name`
      - `+ log_transfer_info.log_transfer_detail.obs_period_unit`
      - `+ log_transfer_info.log_transfer_detail.obs_transfer_path`
      - `+ log_transfer_info.log_transfer_detail.obs_eps_id`
      - `+ log_transfer_info.log_transfer_detail.obs_bucket_name`
      - `+ log_transfer_info.log_transfer_detail.obs_encrypted_enable`
      - `+ log_transfer_info.log_transfer_detail.obs_dir_pre_fix_name`
      - `+ log_transfer_info.log_transfer_detail.dis_id`
      - `+ log_transfer_info.log_transfer_detail.dis_name`
      - `+ log_transfer_info.log_transfer_detail.kafka_id`
      - `+ log_transfer_info.log_transfer_detail.kafka_topic`
      - `+ log_transfer_info.log_transfer_detail.obs_time_zone`
      - `+ log_transfer_info.log_transfer_detail.obs_time_zone_id`
      - `+ log_transfer_info.log_transfer_detail.tags`
      - `* log_transfer_info.log_transfer_detail: object -> object<TransferDetail>`
  - **ListNotificationTemplates**
    - 响应参数变更
      - `+ create_time`
      - `+ project_id`
      - `+ templates`
      - `+ modify_time`
      - `+ name`
      - `+ source`
      - `+ type`
      - `+ locale`
      - `+ desc`
      - `- body`
  - **UpdateSqlAlarmRule**
    - 请求参数变更
      - `+ frequency.type`
      - `+ frequency.cron_expr`
      - `+ frequency.hour_of_day`
      - `+ frequency.day_of_week`
      - `+ frequency.fixed_rate`
      - `+ frequency.fixed_rate_unit`
      - `* frequency: object -> object<Frequency>`
      - `+ notification_save_rule.language`
      - `+ notification_save_rule.timezone`
      - `+ notification_save_rule.user_name`
      - `+ notification_save_rule.topics`
      - `+ notification_save_rule.template_name`
      - `* notification_save_rule: object -> object<SqlNotificationSaveRule>`
    - 响应参数变更
      - `+ frequency.type`
      - `+ frequency.cron_expr`
      - `+ frequency.hour_of_day`
      - `+ frequency.day_of_week`
      - `+ frequency.fixed_rate`
      - `+ frequency.fixed_rate_unit`
      - `* frequency: object -> object<Frequency>`
  - **CreateSqlAlarmRule**
    - 请求参数变更
      - `+ frequency.type`
      - `+ frequency.cron_expr`
      - `+ frequency.hour_of_day`
      - `+ frequency.day_of_week`
      - `+ frequency.fixed_rate`
      - `+ frequency.fixed_rate_unit`
      - `* frequency: object -> object<Frequency>`
      - `+ notification_save_rule.language`
      - `+ notification_save_rule.timezone`
      - `+ notification_save_rule.user_name`
      - `+ notification_save_rule.topics`
      - `+ notification_save_rule.template_name`
      - `* notification_save_rule: object -> object<SqlNotificationSaveRule>`
  - **ListSqlAlarmRules**
    - 响应参数变更
      - `+ sql_alarm_rules.frequency.type`
      - `+ sql_alarm_rules.frequency.cron_expr`
      - `+ sql_alarm_rules.frequency.hour_of_day`
      - `+ sql_alarm_rules.frequency.day_of_week`
      - `+ sql_alarm_rules.frequency.fixed_rate`
      - `+ sql_alarm_rules.frequency.fixed_rate_unit`
      - `* sql_alarm_rules.frequency: object -> object<Frequency>`
  - **UpdateKeywordsAlarmRule**
    - 请求参数变更
      - `+ frequency.type`
      - `+ frequency.cron_expr`
      - `+ frequency.hour_of_day`
      - `+ frequency.day_of_week`
      - `+ frequency.fixed_rate`
      - `+ frequency.fixed_rate_unit`
      - `* frequency: object -> object<Frequency>`
      - `+ notification_save_rule.language`
      - `+ notification_save_rule.timezone`
      - `+ notification_save_rule.user_name`
      - `+ notification_save_rule.topics`
      - `+ notification_save_rule.template_name`
      - `* notification_save_rule: object -> object<SqlNotificationSaveRule>`
  - **CreateKeywordsAlarmRule**
    - 请求参数变更
      - `+ notification_save_rule.language`
      - `+ notification_save_rule.timezone`
      - `+ notification_save_rule.user_name`
      - `+ notification_save_rule.topics`
      - `+ notification_save_rule.template_name`
      - `* notification_save_rule: object -> object<SqlNotificationSaveRule>`
  - **ListKeywordsAlarmRules**
    - 响应参数变更
      - `+ keywords_alarm_rules.frequency.type`
      - `+ keywords_alarm_rules.frequency.cron_expr`
      - `+ keywords_alarm_rules.frequency.hour_of_day`
      - `+ keywords_alarm_rules.frequency.day_of_week`
      - `+ keywords_alarm_rules.frequency.fixed_rate`
      - `+ keywords_alarm_rules.frequency.fixed_rate_unit`
      - `* keywords_alarm_rules.frequency: object -> object<Frequency>`

### HuaweiCloud SDK MetaStudio

- _新增特性_
  - 支持以下接口：
    - `Create2DDigitalHumanVideo`
    - `Show2DDigitalHumanVideo`
    - `Cancel2DDigitalHumanVideo`
    - `ListDigitalHumanBusinessCard`
    - `CreateDigitalHumanBusinessCard`
    - `ShowDigitalHumanBusinessCard`
    - `UpdateDigitalHumanBusinessCard`
    - `DeleteDigitalHumanBusinessCard`
    - `CreatePhotoDigitalHumanVideo`
    - `ShowPhotoDigitalHumanVideo`
    - `CancelPhotoDigitalHumanVideo`
    - `ListSmartLiveRooms`
    - `CreateSmartLiveRoom`
    - `ShowSmartLiveRoom`
    - `UpdateSmartLiveRoom`
    - `DeleteSmartLiveRoom`
    - `ListSmartLive`
    - `StartSmartLive`
    - `ShowSmartLive`
    - `StopSmartLive`
    - `ExecuteSmartLiveCommand`
    - `LiveEventReport`
    - `ListVideoScripts`
    - `CreateVideoScripts`
    - `ShowVideoScript`
    - `UpdateVideoScript`
    - `DeleteVideoScript`
- _解决问题_
  - 无
- _特性变更_
  - **CreatePictureModelingByUrlJob**
    - 请求参数变更
      - `+ X-User-Privilege`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`ExpandCluster`、`ShrinkCluster`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeVehicleCertificate`、`RecognizeAcceptanceBill`、`RecognizeRealEstateCertificate`、`RecognizeVietnamIdCard`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateConsumerGroup**
    - 请求参数变更
      - `* body: object<ConsumerGroup> -> object<UpdateConsumerGroup>`
  - **CreateRocketMqMigrationTask**
    - 响应参数变更
      - `+ task_id`

### HuaweiCloud SDK SMS

- _新增特性_
  - 支持接口`ShowPrivacyAgreements`、`CreatePrivacyAgreements`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.59 2023-09-14

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateBareMetalServers**
    - 请求参数变更
      - `+ server.nics.allowed_address_pairs`

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateVault**
    - 请求参数变更
      - `- vault.billing.promotion_info`
      - `- vault.billing.purchase_mode`
      - `- vault.billing.order_id`
  - **CreatePostPaidVault**
    - 请求参数变更
      - `- vault.billing.promotion_info`
      - `- vault.billing.purchase_mode`
      - `- vault.billing.order_id`

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateEvents**
    - 请求参数变更
      - `- detail.dimensions`
  - **ListEventDetail**
    - 响应参数变更
      - `* event_info.detail.dimensions: object<MetricsDimension> -> list<MetricsDimension>`
      - `* event_info.detail: object<EventItemDetail> -> object<ShowEventItemDetail>`

### HuaweiCloud SDK CodeArtsDeploy

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAppDetailById**
    - 响应参数变更
      - `* result.arrange_infos: object -> list<TaskV2Info>`
  - **ListNewHosts**
    - 响应参数变更
      - `+ result.permission.can_copy`
      - `- result.permission.can_connection_test`
      - `* result.permission: object<PermissionHostDetail> -> object<PermissionHostDetailNew>`
  - **ShowHostDetail**
    - 响应参数变更
      - `* result.proxy_host: string -> object<HostInfoDetail>`
      - `+ result.permission.can_copy`
      - `- result.permission.can_connection_test`
      - `* result.permission: object<PermissionHostDetail> -> object<PermissionHostDetailNew>`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAgentConfig**
    - 响应参数变更
      - `+ result.elasticsearch_enable`
      - `+ result.elasticsearch_shadow_type`
      - `+ result.elasticsearch_shadow_repository`
  - **UpdateAgentHealthStatus**
    - 响应参数变更
      - `+ result.result.elasticsearch_enable`
      - `+ result.result.elasticsearch_shadow_type`
      - `+ result.result.elasticsearch_shadow_repository`

### HuaweiCloud SDK CSMS

- _新增特性_
  - 支持以下接口：
    - `ListSecretEvents`
    - `CreateSecretEvent`
    - `ShowSecretEvent`
    - `UpdateSecretEvent`
    - `DeleteSecretEvent`
    - `ListNotificationRecords`
    - `UpdateVersion`
- _解决问题_
  - 无
- _特性变更_
  - **ListSecrets**
    - 请求参数变更
      - `+ event_name`
    - 响应参数变更
      - `+ secrets.secret_type`
      - `+ secrets.auto_rotation`
      - `+ secrets.rotation_period`
      - `+ secrets.rotation_config`
      - `+ secrets.rotation_time`
      - `+ secrets.next_rotation_time`
      - `+ secrets.event_subscriptions`
      - `+ secrets.enterprise_project_id`
  - **CreateSecret**
    - 请求参数变更
      - `+ secret_type`
      - `+ auto_rotation`
      - `+ rotation_period`
      - `+ rotation_config`
      - `+ event_subscriptions`
      - `+ enterprise_project_id`
    - 响应参数变更
      - `+ secret.secret_type`
      - `+ secret.auto_rotation`
      - `+ secret.rotation_period`
      - `+ secret.rotation_config`
      - `+ secret.rotation_time`
      - `+ secret.next_rotation_time`
      - `+ secret.event_subscriptions`
      - `+ secret.enterprise_project_id`
  - **ShowSecret**
    - 响应参数变更
      - `+ secret.secret_type`
      - `+ secret.auto_rotation`
      - `+ secret.rotation_period`
      - `+ secret.rotation_config`
      - `+ secret.rotation_time`
      - `+ secret.next_rotation_time`
      - `+ secret.event_subscriptions`
      - `+ secret.enterprise_project_id`
  - **UpdateSecret**
    - 请求参数变更
      - `+ auto_rotation`
      - `+ rotation_period`
      - `+ event_subscriptions`
    - 响应参数变更
      - `+ secret.secret_type`
      - `+ secret.auto_rotation`
      - `+ secret.rotation_period`
      - `+ secret.rotation_config`
      - `+ secret.rotation_time`
      - `+ secret.next_rotation_time`
      - `+ secret.event_subscriptions`
      - `+ secret.enterprise_project_id`
  - **UploadSecretBlob**
    - 响应参数变更
      - `+ secret.secret_type`
      - `+ secret.auto_rotation`
      - `+ secret.rotation_period`
      - `+ secret.rotation_config`
      - `+ secret.rotation_time`
      - `+ secret.next_rotation_time`
      - `+ secret.event_subscriptions`
      - `+ secret.enterprise_project_id`
  - **ListSecretVersions**
    - 响应参数变更
      - `+ version_metadatas.expire_time`
  - **CreateSecretVersion**
    - 请求参数变更
      - `+ expire_time`
    - 响应参数变更
      - `+ version_metadata.expire_time`
  - **DeleteSecretForSchedule**
    - 响应参数变更
      - `+ secret.secret_type`
      - `+ secret.auto_rotation`
      - `+ secret.rotation_period`
      - `+ secret.rotation_config`
      - `+ secret.rotation_time`
      - `+ secret.next_rotation_time`
      - `+ secret.event_subscriptions`
      - `+ secret.enterprise_project_id`
  - **RestoreSecret**
    - 响应参数变更
      - `+ secret.secret_type`
      - `+ secret.auto_rotation`
      - `+ secret.rotation_period`
      - `+ secret.rotation_config`
      - `+ secret.rotation_time`
      - `+ secret.next_rotation_time`
      - `+ secret.event_subscriptions`
      - `+ secret.enterprise_project_id`
  - **ListSecretTags**
    - 响应参数变更
      - `* sys_tags: list<TagItem> -> list<SysTag>`
  - **ListProjectSecretsTags**
    - 响应参数变更
      - `* tags: list<Tag> -> list<TagResponse>`
  - **ShowSecretVersion**
    - 响应参数变更
      - `+ version.version_metadata.expire_time`
  - **ListResourceInstances**
    - 请求参数变更
      - `* matches: list<TagItem> -> list<TagMatches>`
    - 响应参数变更
      - `- resources.sys_tags`
      - `+ resources.resource_detail.secret_type`
      - `+ resources.resource_detail.auto_rotation`
      - `+ resources.resource_detail.rotation_period`
      - `+ resources.resource_detail.rotation_config`
      - `+ resources.resource_detail.rotation_time`
      - `+ resources.resource_detail.next_rotation_time`
      - `+ resources.resource_detail.event_subscriptions`
      - `+ resources.resource_detail.enterprise_project_id`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListFlavors**
    - 响应参数变更
      - `+ flavors.replica_count`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateResource**
    - 请求参数变更
      - `+ jobRelation`
      - `+ dependPackages`
      - `+ id`
      - `+ type: enum value [pyFile]`
  - **ShowResource**
    - 响应参数变更
      - `+ jobRelation`
      - `+ dependPackages`
      - `+ id`
      - `+ type: enum value [pyFile]`
  - **UpdateResource**
    - 请求参数变更
      - `+ jobRelation`
      - `+ dependPackages`
      - `+ id`
      - `+ type: enum value [pyFile]`
  - **ListConnections**
    - 请求参数变更
      - `+ limit`
      - `+ offset`
      - `+ connectionName`
  - **ListScripts**
    - 请求参数变更
      - `+ limit`
      - `+ offset`
      - `+ scriptName`
  - **ListJobs**
    - 请求参数变更
      - `+ limit`
      - `+ offset`
      - `+ jobType`
      - `+ jobName`
  - **ListSupplementdata**
    - 请求参数变更
      - `+ sort`
      - `+ page`
      - `+ size`
      - `+ name`
      - `+ userName`
      - `+ status`
      - `+ startDate`
      - `+ endDate`
    - 响应参数变更
      - `* rows.endDate: number -> int64`
      - `* rows.startDate: number -> int64`
      - `* rows.submittedDate: number -> int64`
      - `+ rows.supplement_data_instance_time.days`
      - `+ rows.supplement_data_instance_time.time_of_day`
      - `+ rows.supplement_data_run_time.time_of_day`
      - `+ rows.supplement_data_run_time.day_of_week`
      - `+ rows.supplement_data_run_time.day_of_month`
  - **CreateSupplementdata**
    - 请求参数变更
      - `+ supplement_data_run_time`
      - `+ supplement_data_instance_time`

### HuaweiCloud SDK EVS

- _新增特性_
  - 支持接口`ModifyVolumeQoS`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持接口`ShowFunctionUrl`、`UpdateFunctionUrl`、`CreateFunctionUrl`、`DeleteFunctionUrl`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateFuncSnapshot**
    - 请求参数变更
      - `+ action: enum value [enable,disable]`
  - **CreateFunction**
    - 请求参数变更
      - `+ custom_image`
      - `+ code_type: enum value [Custom-Image-Swr]`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ListAuditLogDownloadLink`
- _解决问题_
  - 无
- _特性变更_
  - **CreateGaussMySqlInstance**
    - 响应参数变更
      - `+ instance.volume`

### HuaweiCloud SDK Image

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `RunImageDescription`
    - `RunImageSuperResolution`
    - `CreateVideoTaggingMediaTask`
    - `ShowVideoTaggingMediaTask`
    - `CreateImageHighresolutionMattingTask`
    - `ShowImageHighresolutionMattingTask`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 支持以下接口：
    - `CreateSchedule`
    - `UpdateSchedule`
    - `DeleteSchedule`
    - `ExecuteDeviceControlsSet`
    - `ExecuteDeviceControlsRelease`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`UpdateInstanceConsumerGroup`、`UpdateInstanceUser`
- _解决问题_
  - 无
- _特性变更_
  - **CreateKafkaConsumerGroup**
    - 请求参数变更
      - `+ group_desc`
  - **CreateInstanceUser**
    - 请求参数变更
      - `+ user_desc`
  - **ShowInstanceUsers**
    - 响应参数变更
      - `+ users.user_desc`
  - **ShowInstanceMessages**
    - 请求参数变更
      - `+ keyword`

### HuaweiCloud SDK KPS

- _新增特性_
  - 支持接口`ImportPrivateKey`、`ExportPrivateKey`、`BatchAssociateKeypair`、`ClearPrivateKey`
- _解决问题_
  - 无
- _特性变更_
  - **ListKeypairDetail**
    - 响应参数变更
      - `+ keypair.key_id`
      - `+ keypair.algorithm`
  - **ListFailedTask**
    - 请求参数变更
      - `* limit: string -> int32`
      - `* offset: string -> int32`
  - **AssociateKeypair**
    - 请求参数变更
      - `+ server.port`
    - 响应参数变更
      - `+ error_msg`
      - `+ error_code`
      - `+ server_id`
      - `+ status`
  - **DisassociateKeypair**
    - 响应参数变更
      - `+ error_msg`
      - `+ error_code`
      - `+ server_id`
      - `+ status`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListCharts**
    - 响应参数变更
      - `- config.can_sort`
      - `- config.can_search`
      - `- config.page_size`
  - **ShowNotificationTemplate**
    - 响应参数变更
      - `+ body`
      - `- create_time`
      - `- project_id`
      - `- templates`
      - `- modify_time`
      - `- name`
      - `- source`
      - `- type`
      - `- locale`
      - `- desc`
  - **ListLogStream**
    - 请求参数变更
      - `- tag`
    - 响应参数变更
      - `* log_streams: list<LogStream> -> list<LogStreamResBody>`
  - **ListStructuredLogsWithTimeRange**
    - 响应参数变更
      - `+ context`
      - `- body`
  - **DeleteTransfer**
    - 响应参数变更
      - `- log_transfer_info.log_transfer_detail.obs_period`
      - `- log_transfer_info.log_transfer_detail.obs_encrypted_id`
      - `- log_transfer_info.log_transfer_detail.obs_prefix_name`
      - `- log_transfer_info.log_transfer_detail.obs_period_unit`
      - `- log_transfer_info.log_transfer_detail.obs_transfer_path`
      - `- log_transfer_info.log_transfer_detail.obs_eps_id`
      - `- log_transfer_info.log_transfer_detail.obs_bucket_name`
      - `- log_transfer_info.log_transfer_detail.obs_encrypted_enable`
      - `- log_transfer_info.log_transfer_detail.obs_dir_pre_fix_name`
      - `- log_transfer_info.log_transfer_detail.dis_id`
      - `- log_transfer_info.log_transfer_detail.dis_name`
      - `- log_transfer_info.log_transfer_detail.kafka_id`
      - `- log_transfer_info.log_transfer_detail.kafka_topic`
      - `- log_transfer_info.log_transfer_detail.obs_time_zone`
      - `- log_transfer_info.log_transfer_detail.obs_time_zone_id`
      - `- log_transfer_info.log_transfer_detail.tags`
  - **ListTransfers**
    - 响应参数变更
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_period`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_encrypted_id`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_prefix_name`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_period_unit`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_transfer_path`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_eps_id`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_bucket_name`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_encrypted_enable`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_dir_pre_fix_name`
      - `- log_transfers.log_transfer_info.log_transfer_detail.dis_id`
      - `- log_transfers.log_transfer_info.log_transfer_detail.dis_name`
      - `- log_transfers.log_transfer_info.log_transfer_detail.kafka_id`
      - `- log_transfers.log_transfer_info.log_transfer_detail.kafka_topic`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_time_zone`
      - `- log_transfers.log_transfer_info.log_transfer_detail.obs_time_zone_id`
      - `- log_transfers.log_transfer_info.log_transfer_detail.tags`
  - **UpdateTransfer**
    - 响应参数变更
      - `- log_transfer_info.log_transfer_detail.obs_period`
      - `- log_transfer_info.log_transfer_detail.obs_encrypted_id`
      - `- log_transfer_info.log_transfer_detail.obs_prefix_name`
      - `- log_transfer_info.log_transfer_detail.obs_period_unit`
      - `- log_transfer_info.log_transfer_detail.obs_transfer_path`
      - `- log_transfer_info.log_transfer_detail.obs_eps_id`
      - `- log_transfer_info.log_transfer_detail.obs_bucket_name`
      - `- log_transfer_info.log_transfer_detail.obs_encrypted_enable`
      - `- log_transfer_info.log_transfer_detail.obs_dir_pre_fix_name`
      - `- log_transfer_info.log_transfer_detail.dis_id`
      - `- log_transfer_info.log_transfer_detail.dis_name`
      - `- log_transfer_info.log_transfer_detail.kafka_id`
      - `- log_transfer_info.log_transfer_detail.kafka_topic`
      - `- log_transfer_info.log_transfer_detail.obs_time_zone`
      - `- log_transfer_info.log_transfer_detail.obs_time_zone_id`
      - `- log_transfer_info.log_transfer_detail.tags`
  - **CreateTransfer**
    - 响应参数变更
      - `- log_transfer_info.log_transfer_detail.obs_period`
      - `- log_transfer_info.log_transfer_detail.obs_encrypted_id`
      - `- log_transfer_info.log_transfer_detail.obs_prefix_name`
      - `- log_transfer_info.log_transfer_detail.obs_period_unit`
      - `- log_transfer_info.log_transfer_detail.obs_transfer_path`
      - `- log_transfer_info.log_transfer_detail.obs_eps_id`
      - `- log_transfer_info.log_transfer_detail.obs_bucket_name`
      - `- log_transfer_info.log_transfer_detail.obs_encrypted_enable`
      - `- log_transfer_info.log_transfer_detail.obs_dir_pre_fix_name`
      - `- log_transfer_info.log_transfer_detail.dis_id`
      - `- log_transfer_info.log_transfer_detail.dis_name`
      - `- log_transfer_info.log_transfer_detail.kafka_id`
      - `- log_transfer_info.log_transfer_detail.kafka_topic`
      - `- log_transfer_info.log_transfer_detail.obs_time_zone`
      - `- log_transfer_info.log_transfer_detail.obs_time_zone_id`
      - `- log_transfer_info.log_transfer_detail.tags`
  - **ListNotificationTemplates**
    - 响应参数变更
      - `+ body`
      - `- create_time`
      - `- project_id`
      - `- templates`
      - `- modify_time`
      - `- name`
      - `- source`
      - `- type`
      - `- locale`
      - `- desc`
  - **UpdateSqlAlarmRule**
    - 请求参数变更
      - `- frequency.type`
      - `- frequency.cron_expr`
      - `- frequency.hour_of_day`
      - `- frequency.day_of_week`
      - `- frequency.fixed_rate`
      - `- frequency.fixed_rate_unit`
      - `- notification_save_rule.language`
      - `- notification_save_rule.timezone`
      - `- notification_save_rule.user_name`
      - `- notification_save_rule.topics`
      - `- notification_save_rule.template_name`
    - 响应参数变更
      - `- frequency.type`
      - `- frequency.cron_expr`
      - `- frequency.hour_of_day`
      - `- frequency.day_of_week`
      - `- frequency.fixed_rate`
      - `- frequency.fixed_rate_unit`
  - **CreateSqlAlarmRule**
    - 请求参数变更
      - `- frequency.type`
      - `- frequency.cron_expr`
      - `- frequency.hour_of_day`
      - `- frequency.day_of_week`
      - `- frequency.fixed_rate`
      - `- frequency.fixed_rate_unit`
      - `- notification_save_rule.language`
      - `- notification_save_rule.timezone`
      - `- notification_save_rule.user_name`
      - `- notification_save_rule.topics`
      - `- notification_save_rule.template_name`
  - **ListSqlAlarmRules**
    - 响应参数变更
      - `- sql_alarm_rules.frequency.type`
      - `- sql_alarm_rules.frequency.cron_expr`
      - `- sql_alarm_rules.frequency.hour_of_day`
      - `- sql_alarm_rules.frequency.day_of_week`
      - `- sql_alarm_rules.frequency.fixed_rate`
      - `- sql_alarm_rules.frequency.fixed_rate_unit`
  - **UpdateKeywordsAlarmRule**
    - 请求参数变更
      - `- frequency.type`
      - `- frequency.cron_expr`
      - `- frequency.hour_of_day`
      - `- frequency.day_of_week`
      - `- frequency.fixed_rate`
      - `- frequency.fixed_rate_unit`
      - `- notification_save_rule.language`
      - `- notification_save_rule.timezone`
      - `- notification_save_rule.user_name`
      - `- notification_save_rule.topics`
      - `- notification_save_rule.template_name`
  - **CreateKeywordsAlarmRule**
    - 请求参数变更
      - `- notification_save_rule.language`
      - `- notification_save_rule.timezone`
      - `- notification_save_rule.user_name`
      - `- notification_save_rule.topics`
      - `- notification_save_rule.template_name`
  - **ListKeywordsAlarmRules**
    - 响应参数变更
      - `- keywords_alarm_rules.frequency.type`
      - `- keywords_alarm_rules.frequency.cron_expr`
      - `- keywords_alarm_rules.frequency.hour_of_day`
      - `- keywords_alarm_rules.frequency.day_of_week`
      - `- keywords_alarm_rules.frequency.fixed_rate`
      - `- keywords_alarm_rules.frequency.fixed_rate_unit`

### HuaweiCloud SDK Meeting

- _新增特性_
  - 支持接口`DeleteToken`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RunCreateVideoModerationJob**
    - 请求参数变更
      - `+ biz_type`
  - **RunCreateAudioModerationJob**
    - 请求参数变更
      - `+ biz_type`
  - **RunTextModeration**
    - 请求参数变更
      - `+ biz_type`
  - **CheckImageModeration**
    - 请求参数变更
      - `+ biz_type`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`RestoreTablesNew`、`UpgradeDbVersionNew`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.58 2023-09-07

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`DeleteStackEnhanced`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListScalingActivityLogs**
    - 响应参数变更
      - `* scaling_activity_log.scaling_value: string -> int32`
  - **CreateScalingPolicy**
    - 请求参数变更
      - `+ scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **UpdateScalingPolicy**
    - 请求参数变更
      - `+ scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **ShowScalingPolicy**
    - 响应参数变更
      - `+ scaling_policy.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scaling_policy.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **ListScalingPolicies**
    - 响应参数变更
      - `+ scaling_policies.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scaling_policies.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **CreateScalingV2Policy**
    - 请求参数变更
      - `+ scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **ListAllScalingV2Policies**
    - 响应参数变更
      - `+ scaling_policies.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scaling_policies.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **UpdateScalingV2Policy**
    - 请求参数变更
      - `+ scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **ShowScalingV2Policy**
    - 响应参数变更
      - `+ scaling_policy.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scaling_policy.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **ListScalingV2Policies**
    - 响应参数变更
      - `+ scaling_policies.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
      - `- scaling_policies.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
  - **ListScalingActivityV2Logs**
    - 响应参数变更
      - `* scaling_activity_log.scaling_value: string -> int32`
  - **CreateScalingGroup**
    - 请求参数变更
      - `+ lbaas_listeners.protocol_version`
  - **ListScalingGroups**
    - 响应参数变更
      - `+ scaling_groups.lbaas_listeners.protocol_version`
  - **UpdateScalingGroup**
    - 请求参数变更
      - `+ lbaas_listeners.protocol_version`
  - **ShowScalingGroup**
    - 响应参数变更
      - `+ scaling_group.lbaas_listeners.protocol_version`

### HuaweiCloud SDK CAE

- _新增特性_
  - 支持接口`ListEips`、`UpdateEip`
- _解决问题_
  - 无
- _特性变更_
  - **ShowApplication**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Application]`
  - **CreateAgency**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Agency]`
  - **ListAgencies**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Agency]`
  - **ListEnvironments**
    - 响应参数变更
      - `+ kind`
      - `- Kind`
      - `+ api_version: enum value [v1]`
  - **CreateEnvironment**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Environment]`
  - **CreateApplication**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Application]`
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Application]`
  - **ListApplications**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Application]`
  - **ListComponentConfigurations**
    - 响应参数变更
      - `+ items.type: enum value [customMetric]`
  - **CreateComponentConfiguration**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [ComponentConfiguration]`
  - **ListComponentEvents**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [ComponentEvent]`
  - **ListComponentInstances**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [ComponentConfiguration]`
  - **ListVolumes**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Volume]`
  - **CreateVolume**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Volume]`
  - **DeleteVolume**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Component]`
  - **UpdateCertificate**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Certificate]`
  - **ListComponentSnapshots**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [ComponentSnapshot]`
  - **ShowJob**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Job]`
  - **ListDomains**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
  - **CreateDomain**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Domain]`
    - 响应参数变更
      - `+ api_version: enum value [v1]`
  - **ListCertificates**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Certificate]`
  - **CreateCertificate**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Certificate]`
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Certificate]`
  - **ListTimerRules**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [TimerRule]`
  - **CreateTimerRule**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [TimerRule]`
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [TimerRule]`
  - **UpdateTimerRule**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [TimerRule]`
  - **ShowExecutionResult**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [TimerRule]`
  - **ShowComponent**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Component]`
  - **UpdateComponent**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Component]`
  - **ExecuteAction**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Action]`
  - **CreateComponent**
    - 请求参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Component]`
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Component]`
  - **ListComponents**
    - 响应参数变更
      - `+ api_version: enum value [v1]`
      - `+ kind: enum value [Component]`

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateEvents**
    - 请求参数变更
      - `+ detail.dimensions`
  - **ListEventDetail**
    - 响应参数变更
      - `- dimensions`
      - `+ event_info.detail.dimensions`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAgentConfig**
    - 响应参数变更
      - `+ result.pulsar_enable`
      - `+ result.pulsar_shadow_topic_prefix`
      - `+ result.mock_rule_list.response_header`
      - `+ result.mock_rule_list.response_body`
      - `+ result.mock_rule_list.response_time`
      - `+ result.mock_rule_list.response_code`
  - **UpdateAgentHealthStatus**
    - 响应参数变更
      - `+ result.result.pulsar_enable`
      - `+ result.result.pulsar_shadow_topic_prefix`
      - `+ result.result.mock_rule_list.response_header`
      - `+ result.result.mock_rule_list.response_body`
      - `+ result.result.mock_rule_list.response_time`
      - `+ result.result.mock_rule_list.response_code`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **StartJob**
    - 请求参数变更
      - `+ jobParams.type`
      - `- jobParams.paramType`
  - **RunOnce**
    - 请求参数变更
      - `+ jobParams.type`
      - `- jobParams.paramType`
  - **ListJobInstances**
    - 请求参数变更
      - `+ limit`
      - `+ offset`
      - `+ minPlanTime`
      - `+ maxPlanTime`
      - `+ status`
      - `+ preciseQuery`
      - `+ jobName`
      - `+ instanceType`
    - 响应参数变更
      - `+ instances.submitTime`
      - `+ instances.instanceId`
      - `+ instances.jobId`
      - `+ instances.jobInstanceName`
      - `+ instances.instanceType`
      - `+ instances.version`
      - `+ instances.ignoreSuccess`
      - `+ instances.forceSuccess`
      - `- instances.instancesId`
      - `+ instances.status: enum value [waiting,running,success,fail,running-exception,pause,manual-stop]`
      - `* instances.planTime: int32 -> int64`
      - `* instances.startTime: int32 -> int64`
      - `* instances.endTime: int32 -> int64`
      - `* instances.executeTime: int32 -> int64`
  - **ListJobs**
    - 响应参数变更
      - `- schedule`
      - `- nodes`
      - `- basicConfig`
      - `- name`
      - `- params`
      - `- jobType`
      - `- directory`
      - `+ jobs.owner`
      - `+ jobs.priority`
      - `+ jobs.status`
      - `+ jobs.createUser`
      - `+ jobs.createTime`
      - `+ jobs.startTime`
      - `+ jobs.endTime`
      - `+ jobs.lastInstanceStatus`
      - `+ jobs.lastInstanceEndTime`
      - `+ jobs.lastUpdateTime`
      - `+ jobs.lastUpdateUser`
      - `+ jobs.path`
      - `+ jobs.singleNodeJobFlag`
      - `+ jobs.flinkJobInfo`
      - `+ jobs.alarms`
      - `- jobs.nodes`
      - `- jobs.schedule`
      - `- jobs.params`
      - `- jobs.directory`
      - `- jobs.basicConfig`
      - `- jobs.jobType: enum value [BATCH,REAL_TIME]`
      - `* jobs: list<JobInfo> -> list<JobResult>`
  - **CreateJob**
    - 请求参数变更
      - `+ logPath`
      - `+ lastUpdateUser`
      - `+ approvers`
      - `+ processType`
      - `+ targetStatus`
      - `- jobType`
      - `+ schedule.type`
      - `- schedule.scheType`
      - `+ params.type`
      - `- params.paramType`
      - `+ nodes.type`
      - `+ nodes.preNodeName`
      - `+ nodes.conditions`
      - `+ nodes.properties`
      - `- nodes.nodeType`
      - `- nodes.preNodeNames`
      - `- nodes.condition`
      - `- nodes.nodeProperties`
      - `+ nodes.failPolicy: enum value [FAIL_CHILD]`
      - `* nodes.location.x: int32 -> string`
      - `* nodes.location.y: int32 -> string`
      - `+ nodes.cronTrigger.expressionTimeZone`
      - `+ nodes.cronTrigger.period`
      - `+ nodes.cronTrigger.concurrent`
      - `* nodes.cronTrigger.dependJobs.jobs: string -> list<string>`
      - `- nodes.cronTrigger.dependJobs.dependFailPolicy: enum value [FAIL,IGNORE,SUSPEND]`
      - `* nodes.cronTrigger.dependJobs: list<DependJob> -> list<DependJobs>`
      - `* nodes.cronTrigger: object<Cron> -> object<CronTrigger>`
  - **ShowJob**
    - 响应参数变更
      - `+ logPath`
      - `+ lastUpdateUser`
      - `+ approvers`
      - `+ processType`
      - `+ targetStatus`
      - `- jobType`
      - `+ schedule.type`
      - `- schedule.scheType`
      - `+ params.type`
      - `- params.paramType`
      - `+ nodes.type`
      - `+ nodes.preNodeName`
      - `+ nodes.conditions`
      - `+ nodes.properties`
      - `- nodes.nodeType`
      - `- nodes.preNodeNames`
      - `- nodes.condition`
      - `- nodes.nodeProperties`
      - `+ nodes.failPolicy: enum value [FAIL_CHILD]`
      - `* nodes.location.x: int32 -> string`
      - `* nodes.location.y: int32 -> string`
      - `+ nodes.cronTrigger.expressionTimeZone`
      - `+ nodes.cronTrigger.period`
      - `+ nodes.cronTrigger.concurrent`
      - `* nodes.cronTrigger.dependJobs.jobs: string -> list<string>`
      - `- nodes.cronTrigger.dependJobs.dependFailPolicy: enum value [FAIL,IGNORE,SUSPEND]`
      - `* nodes.cronTrigger.dependJobs: list<DependJob> -> list<DependJobs>`
      - `* nodes.cronTrigger: object<Cron> -> object<CronTrigger>`
  - **UpdateJob**
    - 请求参数变更
      - `+ logPath`
      - `+ lastUpdateUser`
      - `+ approvers`
      - `+ processType`
      - `+ targetStatus`
      - `- jobType`
      - `+ schedule.type`
      - `- schedule.scheType`
      - `+ params.type`
      - `- params.paramType`
      - `+ nodes.type`
      - `+ nodes.preNodeName`
      - `+ nodes.conditions`
      - `+ nodes.properties`
      - `- nodes.nodeType`
      - `- nodes.preNodeNames`
      - `- nodes.condition`
      - `- nodes.nodeProperties`
      - `+ nodes.failPolicy: enum value [FAIL_CHILD]`
      - `* nodes.location.x: int32 -> string`
      - `* nodes.location.y: int32 -> string`
      - `+ nodes.cronTrigger.expressionTimeZone`
      - `+ nodes.cronTrigger.period`
      - `+ nodes.cronTrigger.concurrent`
      - `* nodes.cronTrigger.dependJobs.jobs: string -> list<string>`
      - `- nodes.cronTrigger.dependJobs.dependFailPolicy: enum value [FAIL,IGNORE,SUSPEND]`
      - `* nodes.cronTrigger.dependJobs: list<DependJob> -> list<DependJobs>`
      - `* nodes.cronTrigger: object<Cron> -> object<CronTrigger>`
  - **CreateSupplementdata**
    - 请求参数变更
      - `+ logPath`
      - `+ lastUpdateUser`
      - `+ approvers`
      - `+ processType`
      - `+ targetStatus`
      - `- jobType`
      - `+ dependJobs.processType`
      - `+ dependJobs.lastUpdateUser`
      - `+ dependJobs.logPath`
      - `+ dependJobs.targetStatus`
      - `+ dependJobs.approvers`
      - `- dependJobs.jobType`
      - `+ dependJobs.schedule.type`
      - `- dependJobs.schedule.scheType`
      - `+ dependJobs.params.type`
      - `- dependJobs.params.paramType`
      - `+ dependJobs.nodes.type`
      - `+ dependJobs.nodes.preNodeName`
      - `+ dependJobs.nodes.conditions`
      - `+ dependJobs.nodes.properties`
      - `- dependJobs.nodes.nodeType`
      - `- dependJobs.nodes.preNodeNames`
      - `- dependJobs.nodes.condition`
      - `- dependJobs.nodes.nodeProperties`
      - `+ dependJobs.nodes.failPolicy: enum value [FAIL_CHILD]`
      - `* dependJobs.nodes.location.x: int32 -> string`
      - `* dependJobs.nodes.location.y: int32 -> string`
      - `+ dependJobs.nodes.cronTrigger.expressionTimeZone`
      - `+ dependJobs.nodes.cronTrigger.period`
      - `+ dependJobs.nodes.cronTrigger.concurrent`
      - `* dependJobs.nodes.cronTrigger.dependJobs.jobs: string -> list<string>`
      - `- dependJobs.nodes.cronTrigger.dependJobs.dependFailPolicy: enum value [FAIL,IGNORE,SUSPEND]`
      - `* dependJobs.nodes.cronTrigger.dependJobs: list<DependJob> -> list<DependJobs>`
      - `* dependJobs.nodes.cronTrigger: object<Cron> -> object<CronTrigger>`

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持以下接口：
    - `StopJobAction`
    - `ShowDataProgress`
    - `UpdateDataProgress`
    - `ShowDataProcessingRulesResult`
    - `CheckDataFilter`
    - `ShowDataFilteringResult`
    - `CollectColumns`
    - `ShowColumnInfoResult`
    - `BatchStopJobsAction`
    - `ExportOperationInfo`
    - `BatchTagAction`
    - `ListProjectTags`
    - `ShowInstanceTags`
    - `UpdateStartPosition`
    - `ShowMonitorData`
    - `ShowSupportObjectType`
    - `ShowIncrementComponentsDetail`
    - `CollectDbObjectsInfo`
    - `ShowDbObjectsList`
- _解决问题_
  - 无
- _特性变更_
  - **ShowDbObjectTemplateResult**
    - 请求参数变更
      - `+ type: enum value [change]`
  - **ShowUpdateObjectSavingStatus**
    - 请求参数变更
      - `+ X-Language: enum value [en-us,zh-cn]`
  - **ShowObjectMapping**
    - 请求参数变更
      - `+ X-Language: enum value [en-us,zh-cn]`
  - **ListJobs**
    - 请求参数变更
      - `+ instance_ids`
      - `+ instance_ip`
  - **ShowDbObjectCollectionStatus**
    - 请求参数变更
      - `+ X-Language: enum value [en-us,zh-cn]`
  - **UpdateBatchAsyncJobs**
    - 请求参数变更
      - `+ jobs.type: enum value [re_create,expired_days]`
  - **UpdateJob**
    - 请求参数变更
      - `+ job.type: enum value [re_create,expired_days]`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListInstancesResourceMetrics`、`ListInstancesRecommendation`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.57 2023-08-31

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListCustomerselfResourceRecordDetails**
    - 请求参数变更
      - `+ query_type`
      - `+ bill_cycle_begin`
      - `+ bill_cycle_end`

### HuaweiCloud SDK CAE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListComponentConfigurations**
    - 响应参数变更
      - `+ items.type: enum value [apm2,log]`
  - **CreateComponentConfiguration**
    - 请求参数变更
      - `+ items.type: enum value [apm2,log]`
  - **ListVolumes**
    - 响应参数变更
      - `+ items.resource_sub_type: enum value [sfs3.0]`
  - **CreateVolume**
    - 请求参数变更
      - `+ spec.resource_sub_type: enum value [sfs3.0]`
  - **UpdateCertificate**
    - 请求参数变更
      - `- spec.policy`
  - **ListDomains**
    - 响应参数变更
      - `- items.metadata.zone_id`
      - `- items.metadata.zone_type`
  - **CreateDomain**
    - 响应参数变更
      - `- items.metadata.zone_id`
      - `- items.metadata.zone_type`
  - **ListCertificates**
    - 响应参数变更
      - `- items.spec.policy`
  - **CreateCertificate**
    - 请求参数变更
      - `- spec.policy`
    - 响应参数变更
      - `- items.spec.policy`
  - **CreateComponent**
    - 请求参数变更
      - `+ spec.runtime: enum value [Java17,Nodejs14,Nodejs16]`

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`RollbackAddonInstance`、`ResizeCluster`、`BatchCreateClusterTags`、`BatchDeleteClusterTags`
- _解决问题_
  - 无
- _特性变更_
  - **ShowAddonInstance**
    - 响应参数变更
      - `+ status.isRollbackable`
      - `+ status.previousVersion`
      - `+ status.status: enum value [rollbackFailed]`
  - **UpdateAddonInstance**
    - 响应参数变更
      - `+ status.isRollbackable`
      - `+ status.previousVersion`
      - `+ status.status: enum value [rollbackFailed]`
  - **CreateAddonInstance**
    - 响应参数变更
      - `+ status.isRollbackable`
      - `+ status.previousVersion`
      - `+ status.status: enum value [rollbackFailed]`
  - **ListAddonInstances**
    - 响应参数变更
      - `+ items.status.isRollbackable`
      - `+ items.status.previousVersion`
      - `+ items.status.status: enum value [rollbackFailed]`

### HuaweiCloud SDK CCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **IssueCertificateAuthorityCertificate**
    - 请求参数变更
      - `+ enterprise_project_id`
  - **CreateCertificateByCsr**
    - 请求参数变更
      - `+ enterprise_project_id`

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowResourceGroup**
    - 响应参数变更
      - `+ resources.event_type`
  - **ListResourceGroup**
    - 响应参数变更
      - `+ resource_groups.type`
      - `+ resource_groups.relation_ids`
      - `+ resource_groups.resources`
  - **ListEventDetail**
    - 响应参数变更
      - `+ dimensions`

### HuaweiCloud SDK CES

- _新增特性_
  - 支持以下接口：
    - `ListDashboardInfos`
    - `CreateOneDashboard`
    - `UpdateDashboard`
    - `DeleteDashboards`
    - `ListDashboardWidgets`
    - `CreateDashboardWidgets`
    - `ShowWidget`
    - `DeleteOneWidget`
    - `BatchUpdateWidgets`
- _解决问题_
  - 无
- _特性变更_
  - **ListAlarmRulePolicies**
    - 响应参数变更
      - `+ policies.extra_info`
      - `+ policies.type`
      - `* policies: list<Policy> -> list<ListPolicy>`
  - **UpdateAlarmRulePolicies**
    - 请求参数变更
      - `+ policies.type`
      - `* policies: list<Policy> -> list<UpdatePolicy>`
    - 响应参数变更
      - `+ policies.type`
      - `* policies: list<Policy> -> list<UpdatePolicy>`
  - **ListAlarmTemplates**
    - 响应参数变更
      - `- alarm_templates.association_alarm_total`
      - `- alarm_templates.policy_total`
      - `- alarm_templates.policy_statistics`
      - `- alarm_templates.association_resource_groups`
  - **ShowAlarmTemplate**
    - 响应参数变更
      - `- association_alarm_total`

### HuaweiCloud SDK CodeArtsDeploy

- _新增特性_
  - 支持以下接口：
    - `ListHostClusters`
    - `CreateHostCluster`
    - `ShowHostClusterDetail`
    - `ListNewHosts`
    - `CreateHost`
    - `ShowHostDetail`
    - `ListEnvironments`
    - `CreateEnvironment`
    - `ShowEnvironmentDetail`
    - `DeleteEnvironment`
    - `ImportHostToEnvironment`
    - `DeleteHostFromEnvironment`
    - `ListAllApp`
    - `CreateApp`
    - `ShowAppDetailById`
    - `DeleteApplication`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DAS

- _新增特性_
  - 支持接口`CreateTuning`、`ShowTuning`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowInstance**
    - 响应参数变更
      - `+ available_zones`

### HuaweiCloud SDK DGC

- _新增特性_
  - 支持接口`ListSupplementdata`、`CreateSupplementdata`、`StopSupplementdata`
- _解决问题_
  - 无
- _特性变更_
  - **ShowJob**
    - 响应参数变更
      - `* basicConfig.priority: string -> int32`
      - `* basicConfig.instanceTimeout: string -> int32`
  - **UpdateJob**
    - 请求参数变更
      - `* basicConfig.priority: string -> int32`
      - `* basicConfig.instanceTimeout: string -> int32`
  - **CreateJob**
    - 请求参数变更
      - `* basicConfig.priority: string -> int32`
      - `* basicConfig.instanceTimeout: string -> int32`
  - **ListJobs**
    - 响应参数变更
      - `* jobs.basicConfig.priority: string -> int32`
      - `* jobs.basicConfig.instanceTimeout: string -> int32`

### HuaweiCloud SDK DLF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowJob**
    - 响应参数变更
      - `* basicConfig.priority: string -> int32`
      - `* basicConfig.instanceTimeout: string -> int32`
  - **UpdateJob**
    - 请求参数变更
      - `* basicConfig.priority: string -> int32`
      - `* basicConfig.instanceTimeout: string -> int32`
  - **CreateJob**
    - 请求参数变更
      - `* basicConfig.priority: string -> int32`
      - `* basicConfig.instanceTimeout: string -> int32`
  - **ListJobs**
    - 响应参数变更
      - `* jobs.basicConfig.priority: string -> int32`
      - `* jobs.basicConfig.instanceTimeout: string -> int32`

### HuaweiCloud SDK DLI

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSqlJobs**
    - 响应参数变更
      - `+ jobs.cpu_cost`
      - `+ jobs.output_byte`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowJob**
    - 响应参数变更
      - `+ entities.server_id`
      - `+ entities.nic_id`
  - **CreateServers**
    - 请求参数变更
      - `+ server.extendparam.CB_CSBS_BACKUP`

### HuaweiCloud SDK ER

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchCreateResourceTags**
    - 请求参数变更
      - `- sys_tags`
  - **ShowStaticRoute**
    - 响应参数变更
      - `- route.attachments.priority`
  - **UpdateStaticRoute**
    - 响应参数变更
      - `- route.attachments.priority`
  - **ListStaticRoutes**
    - 响应参数变更
      - `- routes.attachments.priority`
  - **CreateStaticRoute**
    - 响应参数变更
      - `- route.attachments.priority`
  - **ListEffectiveRoutes**
    - 响应参数变更
      - `- routes.address_group_id`
      - `- routes.next_hops.priority`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **DetectFaceByFile**
    - 响应参数变更
      - `+ faces.attributes.gender`
  - **DetectFaceByFileIntl**
    - 响应参数变更
      - `+ faces.attributes.gender`
  - **DetectFaceByUrl**
    - 响应参数变更
      - `+ faces.attributes.gender`
  - **DetectFaceByUrlIntl**
    - 响应参数变更
      - `+ faces.attributes.gender`
  - **DetectFaceByBase64**
    - 响应参数变更
      - `+ faces.attributes.gender`
  - **DetectFaceByBase64Intl**
    - 响应参数变更
      - `+ faces.attributes.gender`

### HuaweiCloud SDK GES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ExportGraph2**
    - 请求参数变更
      - `+ paginate`

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **StopSimCard**
    - 请求参数变更
      - `+ iccid`
  - **ResetSimCard**
    - 请求参数变更
      - `+ iccid`
  - **ShowSimCard**
    - 请求参数变更
      - `+ iccid`
  - **EnableSimCard**
    - 请求参数变更
      - `+ iccid`
  - **ShowRealNamed**
    - 请求参数变更
      - `+ iccid`
  - **StartStopNet**
    - 请求参数变更
      - `+ iccid`
  - **SetExceedCutNet**
    - 请求参数变更
      - `+ iccid`
  - **RegisterImei**
    - 请求参数变更
      - `+ iccid`
  - **DeleteRealName**
    - 请求参数变更
      - `+ iccid`
  - **SetSpeedValue**
    - 请求参数变更
      - `+ iccid`
  - **ListSimPricePlans**
    - 请求参数变更
      - `+ iccid`
  - **BatchSetTags**
    - 请求参数变更
      - `+ iccids`
  - **BatchSetAttributes**
    - 请求参数变更
      - `+ attributes.iccid`
  - **ShowMonthUsages**
    - 请求参数变更
      - `+ iccids`
    - 响应参数变更
      - `+ month_usages.iccid`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateInstanceByEngine**
    - 请求参数变更
      - `- engine_version: enum value [1.1.0,2.7]`
  - **CreatePostPaidInstance**
    - 请求参数变更
      - `- engine_version: enum value [1.1.0,2.7]`

### HuaweiCloud SDK KooMessage

- _新增特性_
  - 支持接口`ShowTemplateVideoThumbnail`、`SetPrimaryVideoThumbnail`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`ShowMrsVersionList`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeSmartDocumentRecognizer`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RAM

- _新增特性_
  - 支持接口`ListResourceTypes`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstances**
    - 响应参数变更
      - `+ instances.public_dns_names`

### HuaweiCloud SDK ServiceStage

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ChangeResourceInEnvironment**
    - 响应参数变更
      - `+ deploy_mode`

# 0.1.56 2023-08-24

HuaweiCloud SDK APIG

- 新增特性
  - 支持以下接口：
    - ListEndpointConnections
    - AcceptOrRejectEndpointConnections
    - ListEndpointPermissions
    - AddEndpointPermissions
    - DeleteEndpointPermissions
- 解决问题
  - 无
- 特性变更
  - AssociateSignatureKeyV2
    - 响应参数变更
      - + bindings.req_method
  - ListSignatureKeysBindedToApiV2
    - 响应参数变更
      - + bindings.req_method
  - ListApisNotBoundWithSignatureKeyV2
    - 响应参数变更
      - + apis.req_method
  - ListApisBindedToSignatureKeyV2
    - 响应参数变更
      - + bindings.req_method
  - ListApisBindedToRequestThrottlingPolicyV2
    - 响应参数变更
      - + apis.req_method
  - ListApisUnbindedToRequestThrottlingPolicyV2
    - 响应参数变更
      - + apis.req_method
  - ListApiRuntimeDefinitionV2
    - 响应参数变更
      - + content_type: enum value [multipart/form-data]
      - - content_type: enum value [multipart/form-date]
  - ListApisBindedToAclPolicyV2
    - 响应参数变更
      - + apis.req_method
  - ListApisUnbindedToAclPolicyV2
    - 响应参数变更
      - + apis.req_method
  - ShowDetailsOfCustomAuthorizersV2
    - 响应参数变更
      - + network_type
  - UpdateCustomAuthorizerV2
    - 请求参数变更
      - + network_type
    - 响应参数变更
      - + network_type
  - ListInstancesV2
    - 响应参数变更
      - + instances.cbc_operation_locks
      - + instances.status: enum value [Resizing,ResizeFailed,ResizeTimeout]
      - + instances.instance_status: enum value [42,43,44]
      - + instances.spec: enum value [PLATINUM_X2,PLATINUM_X3,PLATINUM_X4,PLATINUM_X5,PLATINUM_X6,PLATINUM_X7,PLATINUM_X8]
  - CreateInstanceV2
    - 请求参数变更
      - + spec_id: enum value [PLATINUM_X2,PLATINUM_X3,PLATINUM_X4,PLATINUM_X5,PLATINUM_X6,PLATINUM_X7,PLATINUM_X8]
  - ShowDetailsOfInstanceV2
    - 响应参数变更
      - + cbc_operation_locks
      - + status: enum value [Resizing,ResizeFailed,ResizeTimeout]
      - + instance_status: enum value [42,43,44]
      - + spec: enum value [PLATINUM_X2,PLATINUM_X3,PLATINUM_X4,PLATINUM_X5,PLATINUM_X6,PLATINUM_X7,PLATINUM_X8]
  - UpdateInstanceV2
    - 响应参数变更
      - + cbc_operation_locks
      - + status: enum value [Resizing,ResizeFailed,ResizeTimeout]
      - + instance_status: enum value [42,43,44]
      - + spec: enum value [PLATINUM_X2,PLATINUM_X3,PLATINUM_X4,PLATINUM_X5,PLATINUM_X6,PLATINUM_X7,PLATINUM_X8]
  - ShowDetailsOfApiV2
    - 响应参数变更
      - + content_type: enum value [multipart/form-data]
      - - content_type: enum value [multipart/form-date]
  - UpdateApiV2
    - 请求参数变更
      - + content_type: enum value [multipart/form-data]
      - - content_type: enum value [multipart/form-date]
    - 响应参数变更
      - + content_type: enum value [multipart/form-data]
      - - content_type: enum value [multipart/form-date]
  - ListApiVersionDetailV2
    - 响应参数变更
      - + content_type: enum value [multipart/form-data]
      - - content_type: enum value [multipart/form-date]
  - CreateCustomAuthorizerV2
    - 请求参数变更
      - + network_type
    - 响应参数变更
      - + network_type
  - ListCustomAuthorizersV2
    - 响应参数变更
      - + network_type
      - + authorizer_list.network_type
  - CreateApiV2
    - 请求参数变更
      - + content_type: enum value [multipart/form-data]
      - - content_type: enum value [multipart/form-date]
    - 响应参数变更
      - + content_type: enum value [multipart/form-data]
      - - content_type: enum value [multipart/form-date]
  - ListApisV2
    - 响应参数变更
      - + apis.content_type: enum value [multipart/form-data]
      - - apis.content_type: enum value [multipart/form-date]

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RestoreBackup**
    - 请求参数变更
      - `+ restore.details`

### HuaweiCloud SDK CCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowCertificateAuthority**
    - 响应参数变更
      - `+ enterprise_project_id`
  - **ShowCertificate**
    - 响应参数变更
      - `+ enterprise_project_id`
  - **CreateCertificateAuthority**
    - 请求参数变更
      - `+ enterprise_project_id`
  - **ListCertificateAuthority**
    - 响应参数变更
      - `+ enterprise_project_id`
      - `+ certificate_authorities.enterprise_project_id`
  - **CreateCertificate**
    - 请求参数变更
      - `+ enterprise_project_id`
  - **ListCertificate**
    - 响应参数变更
      - `+ enterprise_project_id`
      - `+ certificates.enterprise_project_id`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAgentConfig**
    - 响应参数变更
      - `+ result.clickhouse_enable`
      - `+ result.clickhouse_shadow_type`
      - `+ result.clickhouse_shadow_repository`
  - **UpdateAgentHealthStatus**
    - 响应参数变更
      - `+ result.result.clickhouse_enable`
      - `+ result.result.clickhouse_shadow_type`
      - `+ result.result.clickhouse_shadow_repository`

### HuaweiCloud SDK EG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateSubscriptionTarget**
    - 请求参数变更
      - `+ detail.url`
      - `+ detail.agency_name`
      - `+ detail.invocation_http_parameters`
      - `* detail: object -> object<Detail>`
  - **UpdateSubscriptionTarget**
    - 请求参数变更
      - `+ detail.url`
      - `+ detail.agency_name`
      - `+ detail.invocation_http_parameters`
      - `* detail: object -> object<Detail>`
  - **UpdateSubscription**
    - 请求参数变更
      - `+ targets.detail.url`
      - `+ targets.detail.agency_name`
      - `+ targets.detail.invocation_http_parameters`
      - `* targets.detail: object -> object<Detail>`
  - **CreateSubscription**
    - 请求参数变更
      - `+ targets.detail.url`
      - `+ targets.detail.agency_name`
      - `+ targets.detail.invocation_http_parameters`
      - `* targets.detail: object -> object<Detail>`

### HuaweiCloud SDK ER

- _新增特性_
  - 支持接口`AcceptAttachment`、`RejectAttachment`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK HSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListJarPackageHostInfo**
    - 响应参数变更
      - `* data_list.record_time: int32 -> int64`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **GlanceListImages**
    - 响应参数变更
      - `- __root_origin`
      - `- min_disk`
      - `- __image_source_type`
      - `- container_format`
      - `- __image_size`
      - `- __support_xen_gpu_type`
      - `- protected`
      - `- __support_kvm_gpu_type`
      - `- max_ram`
      - `- id`
      - `- __isregistered`
      - `- __lazyloading`
      - `- __data_origin`
      - `- hw_firmware_type`
      - `- __os_type`
      - `- hw_vif_multiqueue_enabled`
      - `- visibility`
      - `- virtual_env_type`
      - `- __image_location`
      - `- __support_kvm`
      - `- __support_highperformance`
      - `- tags`
      - `- __backup_id`
      - `- __platform`
      - `- enterprise_project_id`
      - `- size`
      - `- __support_arm`
      - `- __support_diskintensive`
      - `- __os_version`
      - `- name`
      - `- active_at`
      - `- status`
      - `- schema`
      - `- __is_offshelved`
      - `- __support_kvm_infiniband`
      - `- created_at`
      - `- __originalimagename`
      - `- __support_agent_list`
      - `- __support_amd`
      - `- file`
      - `- updated_at`
      - `- __productcode`
      - `- checksum`
      - `- __support_fc_inject`
      - `- __description`
      - `- min_ram`
      - `- owner`
      - `- __imagetype`
      - `- __support_xen`
      - `- __is_config_init`
      - `- __account_code`
      - `- __system__cmkid`
      - `- __os_bit`
      - `- __support_xen_hana`
      - `- disk_format`
      - `- self`
      - `- __support_largememory`
      - `- __os_feature_list`
      - `- virtual_size`
      - `- __sequence_num`
      - `+ images._sys_enterprise_project_id`
      - `* images: list<GlanceShowImageResponseBody> -> list<GlanceShowImageListResponseBody>`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 支持以下接口：
    - `BatchListAppConfigsTemplates`
    - `AddAppConfigsTemplates`
    - `ShowAppConfigsTemplate`
    - `DeleteAppConfigsTemplate`
    - `AddGeneralAppConfigsTemplate`
    - `ShowModuleShadow`
    - `UpdateModuleShadow`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateModule**
    - 请求参数变更
      - `+ desired_state`

### HuaweiCloud SDK KooMessage

- _新增特性_
  - 支持以下接口：
    - `ListAimMsgTemplate`
    - `CreateAimMsgTemplate`
    - `ShowAimMsgTemplateVariable`
    - `SendAimBatchMessages`
    - `SendAimBatchDifferentMessages`
    - `DeleteAimMsgSignature`
    - `ShowAimMsgTemplateDetail`
    - `UpdateAimMsgTemplate`
    - `DeleteAimMsgTemplate`
    - `ListAimMsgSignature`
    - `AddAimMsgSignature`
    - `ListAimMsgApp`
    - `CreateSmsApp`
    - `ListAimMsgAppDetail`
    - `UpdateAimMsgApp`
    - `ShowAimMsgSignatureFileInfo`
    - `UploadAimMsgSignatureFile`
    - `ListAimMsgSignatureDetail`
    - `UpdateAimMsgSignature`
- _解决问题_
  - 无
- _特性变更_
  - **ListAimResolveDetails**
    - 请求参数变更
      - `+ task_name`
    - 响应参数变更
      - `+ resolve_details.task_name`
  - **ListResolveTasks**
    - 请求参数变更
      - `+ task_name`
    - 响应参数变更
      - `+ resolve_tasks.task_name`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateLogStream**
    - 请求参数变更
      - `* tags: object<tagsBody> -> list<tagsBody>`

### HuaweiCloud SDK NAT

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListNatGateways**
    - 响应参数变更
      - `+ nat_gateways.session_conf`
  - **CreateNatGateway**
    - 请求参数变更
      - `+ nat_gateway.session_conf`
    - 响应参数变更
      - `+ nat_gateway.session_conf`
  - **ShowNatGateway**
    - 响应参数变更
      - `+ nat_gateway.session_conf`
  - **UpdateNatGateway**
    - 请求参数变更
      - `+ nat_gateway.session_conf`
    - 响应参数变更
      - `+ nat_gateway.session_conf`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeVatInvoice**
    - 请求参数变更
      - `+ page_num`
    - 响应参数变更
      - `+ result.invoice_tag`
      - `+ result.sum_amount`
      - `+ result.sum_tax`

### HuaweiCloud SDK OSM

- _新增特性_
  - 支持接口`ShowLoginType`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RunTts**
    - 请求参数变更
      - `+ config.property: enum value [chinese_huaxiaoman_common,chinese_huaxiaofang_common,chinese_huaxiaojun_common]`

### HuaweiCloud SDK VPC

# 0.1.55 2023-08-21

### HuaweiCloud SDK EdgeSec

- _新增特性_
  - 支持边缘安全服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IEF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListEdgeNodes**
    - 请求参数变更
      - `+ group_id`
  - **UpdateDeviceTwin**
    - 响应参数变更
      - `- property_visitors.register_type`
      - `- property_visitors.access_mode`
      - `- property_visitors.register_index`
      - `- property_visitors.register_num`
      - `- property_visitors.scale_index`
      - `- property_visitors.original_datatype`
      - `- property_visitors.expected_datatype`
      - `- property_visitors.is_registerswap`
      - `- property_visitors.is_swap`
      - `- property_visitors.sample_interval`
      - `- property_visitors.data_min`
      - `- property_visitors.data_max`
      - `- property_visitors.node_id`
      - `- property_visitors.browse_name`
      - `* property_visitors: object<ValueInPropertyVisitors> -> map<string, ValueInPropertyVisitors>`
      - `- twin.excepted`
      - `- twin.actual`
      - `- twin.metadata`
      - `- twin.optional`
      - `* twin: object<ValueInTwinResponse> -> map<string, ValueInTwinResponse>`
  - **ListAppVersions**
    - 响应参数变更
      - `+ versions.configs.dns_policy`
  - **CreateAppVersions**
    - 请求参数变更
      - `+ version.configs.dns_policy`
    - 响应参数变更
      - `+ version.configs.dns_policy`
  - **ShowAppVersionDetail**
    - 响应参数变更
      - `+ version.configs.dns_policy`
  - **UpdateAppVersion**
    - 请求参数变更
      - `+ version.configs.dns_policy`
    - 响应参数变更
      - `+ version.configs.dns_policy`
  - **ListPods**
    - 请求参数变更
      - `+ plugin_instance_name`
    - 响应参数变更
      - `+ pods.configs.dns_policy`
  - **ListApps**
    - 响应参数变更
      - `+ apps.app_versions.configs.dns_policy`
  - **CreateApp**
    - 响应参数变更
      - `+ app.app_versions.configs.dns_policy`
  - **ShowAppDetail**
    - 响应参数变更
      - `+ app.app_versions.configs.dns_policy`
  - **UpdateApp**
    - 响应参数变更
      - `+ app.app_versions.configs.dns_policy`
  - **ListDeployments**
    - 响应参数变更
      - `+ deployments.template.configs.dns_policy`
  - **CreateDeployments**
    - 请求参数变更
      - `+ deployment.template.configs.dns_policy`
    - 响应参数变更
      - `+ template.configs.dns_policy`
  - **ShowDeployment**
    - 响应参数变更
      - `+ template.configs.dns_policy`
  - **UpdateDeployment**
    - 请求参数变更
      - `+ deployment.template.configs.dns_policy`
    - 响应参数变更
      - `+ template.configs.dns_policy`

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持接口`DeleteDashboard`
- _解决问题_
  - 无
- _特性变更_
  - **CreateDashBoard**
    - 响应参数变更
      - `* last_update_time: string -> int64`
      - `* useSystemTemplate: string -> boolean`
  - **CreateLogStream**
    - 请求参数变更
      - `- enterprise_project_name`
      - `- log_stream_name: enum value [lts-stream-13ci]`
      - `* ttl_in_days: string -> int32`
      - `* tags: list<tagsBody> -> object<tagsBody>`
  - **ListAccessConfig**
    - 响应参数变更
      - `+ cluster_id`
      - `+ result.cluster_id`
  - **UpdateAccessConfig**
    - 请求参数变更
      - `+ cluster_id`
    - 响应参数变更
      - `+ cluster_id`
  - **CreateAccessConfig**
    - 请求参数变更
      - `+ cluster_id`
    - 响应参数变更
      - `+ cluster_id`
  - **DeleteAccessConfig**
    - 响应参数变更
      - `+ cluster_id`
      - `+ result.cluster_id`

# 0.1.54 2023-08-17

### HuaweiCloud SDK AOS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateStack**
    - 请求参数变更
      - `+ agencies.agency_urn`
  - **GetStackMetadata**
    - 响应参数变更
      - `+ agencies.agency_urn`
  - **CreateStack**
    - 请求参数变更
      - `+ agencies.agency_urn`

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListApiRuntimeDefinitionV2**
    - 响应参数变更
      - `+ req_protocol: enum value [GRPCS]`
      - `+ backend_type: enum value [GRPC]`
  - **ShowDetailsOfApiV2**
    - 响应参数变更
      - `+ req_protocol: enum value [GRPCS]`
      - `+ backend_type: enum value [GRPC]`
      - `+ policy_https.req_protocol: enum value [GRPCS]`
      - `+ backend_api.req_protocol: enum value [GRPCS]`
  - **UpdateApiV2**
    - 请求参数变更
      - `+ req_protocol: enum value [GRPCS]`
      - `+ backend_type: enum value [GRPC]`
      - `+ policy_https.req_protocol: enum value [GRPCS]`
      - `+ backend_api.req_protocol: enum value [GRPCS]`
    - 响应参数变更
      - `+ req_protocol: enum value [GRPCS]`
      - `+ backend_type: enum value [GRPC]`
      - `+ policy_https.req_protocol: enum value [GRPCS]`
      - `+ backend_api.req_protocol: enum value [GRPCS]`
  - **ListApiVersionDetailV2**
    - 响应参数变更
      - `+ req_protocol: enum value [GRPCS]`
      - `+ backend_type: enum value [GRPC]`
      - `+ policy_https.req_protocol: enum value [GRPCS]`
      - `+ backend_api.req_protocol: enum value [GRPCS]`
  - **CreateApiV2**
    - 请求参数变更
      - `+ req_protocol: enum value [GRPCS]`
      - `+ backend_type: enum value [GRPC]`
      - `+ policy_https.req_protocol: enum value [GRPCS]`
      - `+ backend_api.req_protocol: enum value [GRPCS]`
    - 响应参数变更
      - `+ req_protocol: enum value [GRPCS]`
      - `+ backend_type: enum value [GRPC]`
      - `+ policy_https.req_protocol: enum value [GRPCS]`
      - `+ backend_api.req_protocol: enum value [GRPCS]`
  - **ListApisV2**
    - 响应参数变更
      - `+ apis.req_protocol: enum value [GRPCS]`
      - `+ apis.backend_type: enum value [GRPC]`
      - `+ apis.backend_api.req_protocol: enum value [GRPCS]`

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 支持接口`ListRtcAbnormalEvent`、`ListRtcEvent`、`ListObsBuckets`、`ListObsBucketObjects`、`UpdateObsBucketAuthority`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 支持接口`ShowUserExecuteTestCaseInfo`、`ShowTestCaseAndDefectInfo`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ShowNodesInformation`
- _解决问题_
  - 无
- _特性变更_
  - **ShowInstance**
    - 响应参数变更
      - `+ cloud_service_type_code`
      - `+ inquery_spec_code`
      - `+ cloud_resource_type_code`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateServers**
    - 请求参数变更
      - `+ server.root_volume.iops`
      - `+ server.root_volume.throughput`
      - `+ server.root_volume.volumetype: enum value [GPSSD2,ESSD2]`
      - `+ server.data_volumes.iops`
      - `+ server.data_volumes.throughput`
      - `+ server.data_volumes.volumetype: enum value [GPSSD2,ESSD2]`
  - **CreatePostPaidServers**
    - 请求参数变更
      - `+ server.data_volumes.iops`
      - `+ server.data_volumes.throughput`
      - `+ server.data_volumes.volumetype: enum value [GPSSD2,ESSD2]`
      - `+ server.root_volume.iops`
      - `+ server.root_volume.throughput`
      - `+ server.root_volume.volumetype: enum value [GPSSD2,ESSD2]`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowEnv**
    - 响应参数变更
      - `+ public_bucket_path`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ModifyGaussMysqlDns`、`CreateGaussMysqlDns`
- _解决问题_
  - 无
- _特性变更_
  - **ShowGaussMySqlInstanceInfo**
    - 响应参数变更
      - `+ instance.private_dns_names`
  - **ListGaussMySqlInstanceDetailInfo**
    - 响应参数变更
      - `+ instances.private_dns_names`

### HuaweiCloud SDK ImageSearch

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `RunCreateInstance`
    - `RunModifyPicture`
    - `RunAddPicture`
    - `RunDeletePicture`
    - `RunSearchPicture`
    - `RunCheckPicture`
    - `RunQueryInstance`
    - `RunDeleteInstance`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchRestartOrDeleteInstances**
    - 请求参数变更
      - `+ allFailure`
      - `- all_failure`
  - **CreateInstanceByEngine**
    - 请求参数变更
      - `- engine_version: enum value [2.3.0]`
  - **CreatePostPaidInstance**
    - 请求参数变更
      - `- engine_version: enum value [2.3.0]`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSnapshotConfigs**
    - 响应参数变更
      - `* body: object<LiveSnapshotConfig> -> list<LiveSnapshotConfig>`

### HuaweiCloud SDK MetaStudio

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAssetSummary**
    - 响应参数变更
      - `+ asset_list.asset_type: enum value [MUSIC]`
  - **ShowAsset**
    - 响应参数变更
      - `+ asset_type: enum value [MUSIC]`
      - `+ asset_state: enum value [BLOCK]`
      - `+ asset_extra_meta.human_model_2d_meta`
  - **UpdateDigitalAsset**
    - 请求参数变更
      - `+ asset_type: enum value [MUSIC]`
      - `+ asset_extra_meta.human_model_2d_meta`
    - 响应参数变更
      - `+ asset_type: enum value [MUSIC]`
      - `+ asset_state: enum value [BLOCK]`
      - `+ asset_extra_meta.human_model_2d_meta`
  - **CreateDigitalAsset**
    - 请求参数变更
      - `+ asset_type: enum value [MATERIAL,MUSIC]`
      - `+ asset_extra_meta.human_model_2d_meta`
  - **ListAssets**
    - 请求参数变更
      - `+ X-User-MePrivilege`
      - `+ action_editable`
    - 响应参数变更
      - `+ assets.asset_type: enum value [MUSIC]`
      - `+ assets.asset_state: enum value [BLOCK]`
      - `+ assets.asset_extra_meta.human_model_2d_meta`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateThumbnailsTask**
    - 请求参数变更
      - `+ thumbnail_para.dots_ms`
      - `+ thumbnail_para.type: enum value [DOTS_MS]`
  - **CreateTranscodingTask**
    - 请求参数变更
      - `+ thumbnail.params.dots_ms`
      - `+ thumbnail.params.type: enum value [DOTS_MS]`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchRestartOrDeleteInstances**
    - 请求参数变更
      - `+ allFailure`
      - `- all_failure`
  - **CreatePostPaidInstanceByEngine**
    - 请求参数变更
      - `- engine_version: enum value [3.7.17]`
  - **CreatePostPaidInstance**
    - 请求参数变更
      - `- engine_version: enum value [3.7.17]`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BatchDeleteInstances**
    - 请求参数变更
      - `+ allFailure`
      - `- all_failure`
  - **CreatePostPaidInstance**
    - 请求参数变更
      - `- engine_version: enum value [5.x]`

### HuaweiCloud SDK ServiceStage

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ModifyApplication**
    - 请求参数变更
      - `+ enterprise_project_id`
  - **CreateEnvironment**
    - 响应参数变更
      - `+ project_id`
      - `- resources`
  - **ShowEnvironments**
    - 响应参数变更
      - `+ environments.project_id`
  - **ShowEnvironmentInfo**
    - 响应参数变更
      - `+ project_id`
  - **ModifyEnvironment**
    - 请求参数变更
      - `- enterprise_project_id`
    - 响应参数变更
      - `+ project_id`
  - **ShowComponentInfo**
    - 响应参数变更
      - `+ affinity.az`
      - `+ affinity.node`
      - `+ affinity.component`
      - `- affinity.kind`
      - `- affinity.condition`
      - `- affinity.weight`
      - `- affinity.match_expressions`
  - **ModifyComponent**
    - 请求参数变更
      - `+ affinity.az`
      - `+ affinity.node`
      - `+ affinity.component`
      - `- affinity.kind`
      - `- affinity.condition`
      - `- affinity.weight`
      - `- affinity.match_expressions`
    - 响应参数变更
      - `- component_id`
  - **CreateComponent**
    - 请求参数变更
      - `+ affinity.az`
      - `+ affinity.node`
      - `+ affinity.component`
      - `- affinity.kind`
      - `- affinity.condition`
      - `- affinity.weight`
      - `- affinity.match_expressions`
  - **ShowComponentsInApplication**
    - 响应参数变更
      - `+ components.external_accesses`
  - **ShowComponents**
    - 响应参数变更
      - `+ components.external_accesses`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateSecurityGroupRule**
    - 请求参数变更
      - `+ security_group_rule.remote_address_group_id`
  - **NeutronCreateSecurityGroupRule**
    - 请求参数变更
      - `+ security_group_rule.remote_address_group_id`

# 0.1.53 2023-08-10

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSubCustomerBillDetail**
    - 响应参数变更
      - `+ fee_records.id`

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPrivateZones**
    - 请求参数变更
      - `* type: optional -> required`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持以下接口：
    - `DownloadData`
    - `ListDrugJob`
    - `CancelDrugJob`
    - `DeleteDrugJob`
    - `CreateOptmJob`
    - `ShowOptmJob`
    - `CreateSynthesisJob`
    - `ShowSynthesisJob`
    - `CreateDockingJob`
    - `ShowDockingJob`
    - `CreateFepJob`
    - `ShowFepJob`
    - `CreateDrugLigandSvg`
    - `CreateDrugLigandSdf`
    - `RunDrugReceptorPreprocess`
    - `ParseDrugReceptorInfo`
    - `RecognizeDrugReceptorPocket`
    - `RunDrugLigandToSmilesConversion`
    - `CreateDrugLigandInteraction2dSvg`
    - `CheckDrugLigandDifference`
    - `CreateDrugLigandPreviewTask`
    - `ShowDrugLigandPreviewTask`
    - `DeleteDrugLigandPreviewTask`
    - `CreateDrugLigandSimilarityGraphTask`
    - `ShowDrugLigandSimilarityGraphTask`
    - `DeleteDrugLigandSimilarityGraphTask`
- _解决问题_
  - 无
- _特性变更_
  - **ListComputingResourceFlavors**
    - 响应参数变更
      - `+ flavors.az`
  - **StartAutoJob**
    - 响应参数变更
      - `+ max_platform_flavor`
      - `+ app_infos.app_resource`
  - **ListComputingResources**
    - 响应参数变更
      - `+ resources.evs_resource_id`
  - **ExecuteJob**
    - 响应参数变更
      - `+ max_platform_flavor`
      - `+ app_infos.app_resource`
  - **CreateAutoJob**
    - 响应参数变更
      - `+ max_platform_flavor`
      - `+ app_infos.app_resource`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`UpdateProxyPort`、`DescribeBackupEncryptStatus`、`ModifyBackupEncryptStatus`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateProxySessionConsistence**
    - 请求参数变更
      - `+ consistence_mode`
  - **CreateGaussMySqlInstance**
    - 请求参数变更
      - `* datastore: object<MysqlDatastore> -> object<MysqlDatastoreInReq>`
    - 响应参数变更
      - `* instance.datastore: object<MysqlDatastore> -> object<MysqlDatastoreInRes>`
  - **ShowGaussMySqlBackupList**
    - 响应参数变更
      - `- backups.datastore.kernel_version`
      - `* backups.datastore: object<MysqlDatastore> -> object<MysqlDatastoreInBackup>`
  - **ShowGaussMySqlProxyList**
    - 响应参数变更
      - `+ proxy_list.proxy.consistence_mode`

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持接口`ListWorkOrders`、`ListWorkOrderDetails`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstanceConsumerGroups**
    - 响应参数变更
      - `+ groups.createdAt`
      - `+ groups.group_desc`
      - `+ groups.lag`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeMyanmarIdcard**
    - 请求参数变更
      - `+ return_translation`
    - 响应参数变更
      - `+ result.translation_info`

### HuaweiCloud SDK RAM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`SearchDistinctSharedResources`、`SearchDistinctPrincipals`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListXellogFiles`、`CreateXelLogDownload`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowOneTopic**
    - 响应参数变更
      - `+ message_type`
  - **ShowTopicStatus**
    - 响应参数变更
      - `+ max_offset`
      - `+ min_offset`
  - **ShowInstance**
    - 响应参数变更
      - `+ grpc_address`
      - `+ public_grpc_address`
  - **CreateTopicOrBatchDeleteTopic**
    - 请求参数变更
      - `+ message_type`
  - **ListRocketInstanceTopics**
    - 响应参数变更
      - `+ message_type`
      - `+ topics.message_type`
  - **ListMessages**
    - 响应参数变更
      - `* messages.reconsume_times: string -> int32`
      - `* messages.queue_id: string -> int32`
      - `* messages.queue_offset: string -> int32`
  - **ExportDlqMessage**
    - 响应参数变更
      - `* reconsume_times: string -> int32`
      - `* queue_id: string -> int32`
      - `* queue_offset: string -> int32`
  - **CreatePostPaidInstance**
    - 请求参数变更
      - `+ engine_version: enum value [5.x]`
  - **ListInstances**
    - 响应参数变更
      - `+ grpc_address`
      - `+ public_grpc_address`
      - `+ instances.grpc_address`
      - `+ instances.public_grpc_address`
  - **ShowConsumerListOrDetails**
    - 响应参数变更
      - `+ lag`
      - `+ max_offset`
      - `+ consumer_offset`

# 0.1.52 2023-08-03

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListCustomerselfResourceRecords**
    - 响应参数变更
      - `+ fee_records.id`

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListCustomerselfResourceRecords**
    - 响应参数变更
      - `+ fee_records.id`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowNode**
    - 响应参数变更
      - `- spec.extendParam.enterprise_project_id`
  - **UpdateNode**
    - 响应参数变更
      - `- spec.extendParam.enterprise_project_id`
  - **DeleteNode**
    - 响应参数变更
      - `- spec.extendParam.enterprise_project_id`
  - **CreateNode**
    - 请求参数变更
      - `- spec.extendParam.enterprise_project_id`
    - 响应参数变更
      - `- spec.extendParam.enterprise_project_id`
  - **ListNodes**
    - 响应参数变更
      - `- items.spec.extendParam.enterprise_project_id`
  - **ShowNodePool**
    - 响应参数变更
      - `- spec.nodeTemplate.extendParam.enterprise_project_id`
  - **UpdateNodePool**
    - 响应参数变更
      - `- spec.nodeTemplate.extendParam.enterprise_project_id`
  - **DeleteNodePool**
    - 响应参数变更
      - `- spec.nodeTemplate.extendParam.enterprise_project_id`
  - **CreateNodePool**
    - 请求参数变更
      - `- spec.nodeTemplate.extendParam.enterprise_project_id`
    - 响应参数变更
      - `- spec.nodeTemplate.extendParam.enterprise_project_id`
  - **ListNodePools**
    - 响应参数变更
      - `- items.spec.nodeTemplate.extendParam.enterprise_project_id`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDomainDetailByName**
    - 响应参数变更
      - `- domain.sources.weight`
      - `* domain.sources: list<SourcesConfig> -> list<SourcesDomainConfig>`
  - **ShowDomainFullConfig**
    - 响应参数变更
      - `+ configs.remark`
      - `+ configs.ip_frequency_limit`
      - `+ configs.hsts`
      - `+ configs.quic`
      - `+ configs.url_auth.inherit_config`
      - `+ configs.sources.bucket_access_key`
      - `+ configs.sources.bucket_secret_key`
      - `+ configs.sources.bucket_region`
      - `+ configs.sources.bucket_name`
      - `+ configs.request_limit_rules.priority`
      - `+ configs.request_limit_rules.match_type`
      - `+ configs.request_limit_rules.match_value`
  - **UpdateDomainFullConfig**
    - 请求参数变更
      - `+ configs.remark`
      - `+ configs.ip_frequency_limit`
      - `+ configs.hsts`
      - `+ configs.quic`
      - `+ configs.url_auth.inherit_config`
      - `+ configs.sources.bucket_access_key`
      - `+ configs.sources.bucket_secret_key`
      - `+ configs.sources.bucket_region`
      - `+ configs.sources.bucket_name`
      - `+ configs.request_limit_rules.priority`
      - `+ configs.request_limit_rules.match_type`
      - `+ configs.request_limit_rules.match_value`

### HuaweiCloud SDK Config

- _新增特性_
  - 支持以下接口：
    - `ListConformancePacks`
    - `CreateConformancePack`
    - `ShowConformancePack`
    - `DeleteConformancePack`
    - `CollectConformancePackComplianceSummary`
    - `ListConformancePackComplianceByPackId`
    - `ListConformancePackComplianceDetailsByPackId`
    - `ListConformancePackComplianceScores`
    - `ListBuiltInConformancePackTemplates`
    - `ShowBuiltInConformancePackTemplate`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAgentConfig**
    - 请求参数变更
      - `+ alias`

### HuaweiCloud SDK CTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **DeleteTracker**
    - 请求参数变更
      - `+ tracker_type: enum value [system]`

### HuaweiCloud SDK EG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDetailOfEventSource**
    - 响应参数变更
      - `+ error_info`
  - **UpdateEventSource**
    - 响应参数变更
      - `+ error_info`
  - **CreateEventSource**
    - 响应参数变更
      - `+ error_info`
  - **ListEventSources**
    - 响应参数变更
      - `+ error_info`
      - `+ items.error_info`
  - **CreateSubscriptionTarget**
    - 请求参数变更
      - `+ smn_detail`
      - `+ dead_letter_queue`
    - 响应参数变更
      - `+ dead_letter_queue`
  - **ShowDetailOfSubscriptionTarget**
    - 响应参数变更
      - `+ dead_letter_queue`
  - **UpdateSubscriptionTarget**
    - 请求参数变更
      - `+ smn_detail`
      - `+ dead_letter_queue`
    - 响应参数变更
      - `+ dead_letter_queue`
  - **ShowDetailOfConnection**
    - 响应参数变更
      - `+ error_info`
  - **UpdateConnection**
    - 响应参数变更
      - `+ error_info`
  - **UpdateEndpoint**
    - 响应参数变更
      - `+ error_info`
  - **ShowDetailOfSubscription**
    - 响应参数变更
      - `+ dead_letter_queue`
      - `+ targets.dead_letter_queue`
  - **UpdateSubscription**
    - 请求参数变更
      - `+ targets.smn_detail`
      - `+ targets.dead_letter_queue`
    - 响应参数变更
      - `+ dead_letter_queue`
      - `+ targets.dead_letter_queue`
  - **CreateConnection**
    - 响应参数变更
      - `+ error_info`
  - **ListConnections**
    - 请求参数变更
      - `+ instance_id`
    - 响应参数变更
      - `+ error_info`
      - `+ items.error_info`
  - **CreateEndpoint**
    - 响应参数变更
      - `+ error_info`
  - **ListEndpoints**
    - 请求参数变更
      - `+ subnet_id`
    - 响应参数变更
      - `+ error_info`
      - `+ items.error_info`
  - **ShowEventStreaming**
    - 响应参数变更
      - `+ source.source_kafka.seek_to: enum value [latest,earliest]`
      - `+ source.source_kafka.sasl_mechanism: enum value [SCRAM-SHA-512,PLAIN]`
  - **UpdateEventStreaming**
    - 请求参数变更
      - `+ source.source_kafka.seek_to: enum value [latest,earliest]`
      - `+ source.source_kafka.sasl_mechanism: enum value [SCRAM-SHA-512,PLAIN]`
  - **CreateSubscription**
    - 请求参数变更
      - `+ targets.smn_detail`
      - `+ targets.dead_letter_queue`
    - 响应参数变更
      - `+ dead_letter_queue`
      - `+ targets.dead_letter_queue`
  - **ListSubscriptions**
    - 响应参数变更
      - `+ dead_letter_queue`
      - `+ items.targets.dead_letter_queue`
  - **ListTriggers**
    - 响应参数变更
      - `+ dead_letter_queue`
      - `+ items.targets.dead_letter_queue`
  - **ListWorkflowTriggers**
    - 响应参数变更
      - `+ dead_letter_queue`
      - `+ items.targets.dead_letter_queue`
  - **CreateEventStreaming**
    - 请求参数变更
      - `+ source.source_kafka.seek_to: enum value [latest,earliest]`
      - `+ source.source_kafka.sasl_mechanism: enum value [SCRAM-SHA-512,PLAIN]`
  - **ListEventStreaming**
    - 响应参数变更
      - `+ source.source_kafka.seek_to: enum value [latest,earliest]`
      - `+ source.source_kafka.sasl_mechanism: enum value [SCRAM-SHA-512,PLAIN]`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ModifyGaussMySqlProxyRouteMode`
- _解决问题_
  - 无
- _特性变更_
  - **ShowGaussMySqlEngineVersion**
    - 响应参数变更
      - `+ datastores.version`
      - `+ datastores.kernel_version`
  - **CreateGaussMySqlProxy**
    - 请求参数变更
      - `+ route_mode`
  - **CreateGaussMySqlInstance**
    - 请求参数变更
      - `+ datastore.kernel_version`
    - 响应参数变更
      - `+ instance.datastore.kernel_version`
  - **ShowGaussMySqlBackupList**
    - 响应参数变更
      - `+ backups.datastore.kernel_version`
  - **ShowGaussMySqlProxyList**
    - 响应参数变更
      - `+ proxy_list.proxy.route_mode`
      - `+ proxy_list.proxy.balance_route_mode_enabled`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstances**
    - 响应参数变更
      - `+ instances.backup_used_space`
  - **ListComponentInfos**
    - 请求参数变更
      - `+ component_type`
      - `+ availability_zone_id`
    - 响应参数变更
      - `+ nodes.name`
      - `+ nodes.availability_zone_id`
      - `+ nodes.description`
      - `+ nodes.status`
      - `+ nodes.components.distributed_id`
  - **ListInstancesDetails**
    - 响应参数变更
      - `+ instances.backup_used_space`

### HuaweiCloud SDK KooMessage

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **DeleteTemplateMaterial**
    - 响应参数变更
      - `+ data`
  - **DeleteAimPersonalTemplate**
    - 响应参数变更
      - `+ data`
  - **UnfreezePub**
    - 响应参数变更
      - `+ data.pub_id`
      - `- data.data`
  - **FreezePub**
    - 响应参数变更
      - `+ data.pub_id`
      - `- data.data`
  - **ListAimResolveDetails**
    - 响应参数变更
      - `* resolve_details.resolved_status: object -> string`
  - **CreateResolveTask**
    - 请求参数变更
      - `- params.sms_params`
      - `* params: list<CreateResolveTaskParam> -> list<CreateShortChainParam>`
  - **ListAimTemplates**
    - 响应参数变更
      - `+ templates.factory_info.version`
  - **CreateVmsTemplate**
    - 请求参数变更
      - `- reminders`

### HuaweiCloud SDK MetaStudio

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListSelfPrivileges`
  - **CreateFile**
    - 响应参数变更
      - `+ file_id`
      - `+ upload_url`
  - **ListAssetSummary**
    - 响应参数变更
      - `+ asset_list.asset_type: enum value [HUMAN_MODEL_2D,BUSINESS_CARD_TEMPLET]`
  - **CreateTtsa**
    - 请求参数变更
      - `+ X-App-UserId`
      - `+ X-User-Privilege`
  - **ListTtsaJobs**
    - 请求参数变更
      - `+ X-App-UserId`
  - **ListTtsaData**
    - 响应参数变更
      - `+ motions.eyes`
      - `* motions.root: list<object> -> list<number>`
      - `* motions.joints: list<object> -> list<number>`
  - **CreatePictureModelingJob**
    - 响应参数变更
      - `+ model_asset_id`
      - `+ job_id`
  - **ListPictureModelingJobs**
    - 请求参数变更
      - `+ sort_dir: enum value [asc,desc]`
  - **DeleteAsset**
    - 请求参数变更
      - `+ mode`
  - **ShowAsset**
    - 响应参数变更
      - `+ asset_type: enum value [HUMAN_MODEL_2D,BUSINESS_CARD_TEMPLET]`
      - `+ system_properties.key: enum value [CREATED_BY_PLATFORM]`
      - `+ asset_extra_meta.voice_model_meta.tts_mode`
      - `+ asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.human_model_meta.model_properties`
  - **UpdateDigitalAsset**
    - 请求参数变更
      - `+ asset_type: enum value [HUMAN_MODEL_2D,BUSINESS_CARD_TEMPLET]`
      - `+ system_properties.key: enum value [CREATED_BY_PLATFORM]`
      - `+ asset_extra_meta.voice_model_meta.tts_mode`
      - `+ asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.human_model_meta.model_properties`
    - 响应参数变更
      - `+ asset_type: enum value [HUMAN_MODEL_2D,BUSINESS_CARD_TEMPLET]`
      - `+ system_properties.key: enum value [CREATED_BY_PLATFORM]`
      - `+ asset_extra_meta.voice_model_meta.tts_mode`
      - `+ asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.human_model_meta.model_properties`
  - **ListStyles**
    - 请求参数变更
      - `+ sort_dir: enum value [asc,desc]`
    - 响应参数变更
      - `+ styles.extra_meta.model_id`
  - **CreateVideoMotionCaptureJob**
    - 响应参数变更
      - `+ rtc_room_info`
      - `+ job_id`
  - **CreateDigitalAsset**
    - 请求参数变更
      - `+ asset_type: enum value [HUMAN_MODEL_2D,BUSINESS_CARD_TEMPLET]`
      - `+ system_properties.key: enum value [CREATED_BY_PLATFORM]`
      - `+ asset_extra_meta.voice_model_meta.tts_mode`
      - `+ asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ asset_extra_meta.human_model_meta.model_properties`
  - **ListAssets**
    - 请求参数变更
      - `+ language`
      - `- lanuage`
      - `+ sort_dir: enum value [asc,desc]`
    - 响应参数变更
      - `+ assets.asset_type: enum value [HUMAN_MODEL_2D,BUSINESS_CARD_TEMPLET]`
      - `+ assets.system_properties.key: enum value [CREATED_BY_PLATFORM]`
      - `+ assets.asset_extra_meta.voice_model_meta.tts_mode`
      - `+ assets.asset_extra_meta.voice_model_meta.external_voice_meta`
      - `+ assets.asset_extra_meta.human_model_meta.model_properties`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`ListDataConnector`、`CreateDataConnector`、`UpdateDataConnector`、`DeleteDataConnector`
- _解决问题_
  - 无
- _特性变更_
  - **CreateCluster**
    - 请求参数变更
      - `+ charge_info.period_type`
      - `+ charge_info.period_num`
      - `+ charge_info.is_auto_pay`
  - **RunJobFlow**
    - 请求参数变更
      - `+ charge_info.period_type`
      - `+ charge_info.period_num`
      - `+ charge_info.is_auto_pay`

### HuaweiCloud SDK OSM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CreateAuthorization`
  - **CreateMessages**
    - 请求参数变更
      - `- message.is_authorized`
      - `- message.authorization_content`
  - **CreateCases**
    - 请求参数变更
      - `- is_authorized`
      - `- authorization_content`
  - **ShowCaseDetail**
    - 响应参数变更
      - `- incident_detail_info.is_authorized`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowIssueV4**
    - 响应参数变更
      - `+ find_release_dev`
      - `+ release_dev`
      - `+ env`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListDatastores**
    - 请求参数变更
      - `+ database_name: enum value [MariaDB]`
  - **ListConfigurations**
    - 响应参数变更
      - `+ configurations.datastore_name: enum value [mariadb]`
  - **CreateConfiguration**
    - 请求参数变更
      - `+ datastore.type: enum value [MariaDB]`
    - 响应参数变更
      - `+ configuration.datastore_name: enum value [mariadb]`
  - **ShowConfiguration**
    - 响应参数变更
      - `+ datastore_name: enum value [mariadb]`
  - **ShowInstanceConfiguration**
    - 响应参数变更
      - `+ datastore_name: enum value [mariadb]`
  - **ListFlavors**
    - 请求参数变更
      - `+ database_name: enum value [MariaDB]`
  - **ListStorageTypes**
    - 请求参数变更
      - `+ database_name: enum value [MariaDB]`
  - **ListInstances**
    - 请求参数变更
      - `+ datastore_type: enum value [MariaDB]`
    - 响应参数变更
      - `+ instances.datastore.type: enum value [MariaDB]`
  - **CreateInstance**
    - 请求参数变更
      - `+ datastore.type: enum value [MariaDB]`
    - 响应参数变更
      - `+ instance.datastore.type: enum value [MariaDB]`
  - **CreateRestoreInstance**
    - 请求参数变更
      - `+ datastore.type: enum value [MariaDB]`
    - 响应参数变更
      - `+ instance.datastore.type: enum value [MariaDB]`
  - **ListBackups**
    - 响应参数变更
      - `+ backups.datastore.type: enum value [MariaDB]`
  - **ListOffSiteBackups**
    - 响应参数变更
      - `+ backups.datastore.type: enum value [MariaDB]`
  - **ListOffSiteInstances**
    - 响应参数变更
      - `+ offsite_backup_instances.datastore.type: enum value [MariaDB]`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListMessageTrace**
    - 请求参数变更
      - `* msg_id: optional -> required`
  - **ListMessages**
    - 请求参数变更
      - `+ key`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPorts**
    - 请求参数变更
      - `+ enable_efi`
    - 响应参数变更
      - `+ ports.enable_efi`
  - **CreatePort**
    - 响应参数变更
      - `+ port.enable_efi`
  - **ShowPort**
    - 响应参数变更
      - `+ port.enable_efi`
  - **UpdatePort**
    - 响应参数变更
      - `+ port.enable_efi`

# 0.1.51 2023-07-31

### HuaweiCloud SDK CAE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ExecuteAction**
    - 请求参数变更
      - `+ spec.build`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateIssueV4**
    - 响应参数变更
      - `+ find_release_dev`
      - `+ order`
      - `+ release_dev`
      - `+ env`
  - **ListIssuesV4**
    - 响应参数变更
      - `+ find_release_dev`
      - `+ order`
      - `+ release_dev`
      - `+ env`
      - `+ issues.order`
      - `+ issues.release_dev`
      - `+ issues.find_release_dev`
      - `+ issues.env`
  - **ListChildIssuesV4**
    - 响应参数变更
      - `+ find_release_dev`
      - `+ order`
      - `+ release_dev`
      - `+ env`
      - `+ issues.order`
      - `+ issues.release_dev`
      - `+ issues.find_release_dev`
      - `+ issues.env`

### HuaweiCloud SDK Workspace

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateDesktop**
    - 请求参数变更
      - `- security_groups.name`
      - `* security_groups: list<SecurityGroup> -> list<SecurityGroupInfo>`
  - **ShowDesktopDetail**
    - 响应参数变更
      - `- desktop.security_groups.name`
      - `* desktop.security_groups: list<SecurityGroup> -> list<SecurityGroupInfo>`
  - **ListDesktopsDetail**
    - 响应参数变更
      - `- desktops.security_groups.name`
      - `* desktops.security_groups: list<SecurityGroup> -> list<SecurityGroupInfo>`

# 0.1.50 2023-07-27

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 支持接口`ListTestCases`、`ListTestCaseHistories`、`ListBranches`、`ShowApiTestcaseHistories`
- _解决问题_
  - 无
- _特性变更_
  - **ShowPlans**
    - 响应参数变更
      - `* expire_day: string -> int32`
  - **ShowPlanList**
    - 响应参数变更
      - `* expire_day: string -> int32`

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 支持接口`CreateMergeRequestDiscussion`、`CreateMergeRequestDiscussionNote`、`ShowReviewSetting`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持以下接口：
    - `DownloadBatchCreateTemplate`
    - `ImportBatchCreateJobs`
    - `CopyJob`
    - `ShowMetering`
    - `ShowDirtyData`
    - `ShowComparePolicy`
    - `ShowHealthCompareJobList`
    - `ShowProgressData`
    - `ShowObjectMapping`
    - `ShowActions`
    - `ValidateJobName`
    - `ShowEnterpriseProject`
- _解决问题_
  - 无
- _特性变更_
  - **DownloadDbObjectTemplate**
    - 请求参数变更
      - `+ file_import_db_level`
  - **UploadDbObjectTemplate**
    - 请求参数变更
      - `+ file_import_db_level`
  - **ListAsyncJobs**
    - 响应参数变更
      - `+ jobs.status: enum value [AUTO_PARAM_VALIDATE_SUCCESS,COMMIT_SUCCESS]`
      - `- jobs.status: enum value [ASYNC_JOB_CREATING,ASYNC_JOB_CREATE_FAILED,ASYNC_JOB_COMPLETED]`
  - **CreateJob**
    - 请求参数变更
      - `+ job.node_info.base_info`
    - 响应参数变更
      - `+ is_clone_job`
      - `+ create_time`
      - `+ name`
      - `+ id`
      - `+ status`
      - `+ job.is_clone_job`
  - **BatchCreateJobsAsync**
    - 请求参数变更
      - `+ jobs.node_info.base_info`
  - **ListAsyncJobDetail**
    - 响应参数变更
      - `+ jobs.support_import_file_resp`
      - `+ jobs.instance_features`
      - `+ jobs.task_version`
      - `+ jobs.node_info.base_info`
  - **UpdateBatchAsyncJobs**
    - 请求参数变更
      - `+ jobs.type: enum value [policy]`
      - `- jobs.type: enum value [policy_config]`
      - `+ jobs.params.node_info.base_info`
  - **ShowJobDetail**
    - 请求参数变更
      - `+ type: enum value [file]`
    - 响应参数变更
      - `+ job.support_import_file_resp`
      - `+ job.instance_features`
      - `+ job.task_version`
      - `+ job.node_info.base_info`
  - **UpdateJob**
    - 请求参数变更
      - `+ job.type: enum value [policy]`
      - `- job.type: enum value [policy_config]`
      - `+ job.params.node_info.base_info`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **AttachShareBandwidth**
    - 响应参数变更
      - `+ publicip.vnic.vtep`
      - `+ publicip.vnic.vni`
      - `+ publicip.vnic.port_profile`
  - **DetachShareBandwidth**
    - 响应参数变更
      - `+ publicip.vnic.vtep`
      - `+ publicip.vnic.vni`
      - `+ publicip.vnic.port_profile`
  - **EnableNat64**
    - 响应参数变更
      - `+ publicip.vnic.vtep`
      - `+ publicip.vnic.vni`
      - `+ publicip.vnic.port_profile`
  - **DisableNat64**
    - 响应参数变更
      - `+ publicip.vnic.vtep`
      - `+ publicip.vnic.vni`
      - `+ publicip.vnic.port_profile`
  - **AttachBatchPublicIp**
    - 响应参数变更
      - `+ publicips.publicip.vnic.vtep`
      - `+ publicips.publicip.vnic.vni`
      - `+ publicips.publicip.vnic.port_profile`
  - **DetachBatchPublicIp**
    - 响应参数变更
      - `+ publicips.publicip.vnic.vtep`
      - `+ publicips.publicip.vnic.vni`
      - `+ publicips.publicip.vnic.port_profile`

### HuaweiCloud SDK ER

- _新增特性_
  - 支持以下接口：
    - `BatchCreateResourceTags`
    - `ShowQuotas`
    - `ListFlowLogs`
    - `CreateFlowLog`
    - `ShowFlowLog`
    - `UpdateFlowLog`
    - `DeleteFlowLog`
    - `EnableFlowLog`
    - `DisableFlowLog`
- _解决问题_
  - 无
- _特性变更_
  - **ListProjectTags**
    - 请求参数变更
      - `+ resource_type: enum value [ecn-attachment,connect-attachment,cfw-attachment]`
    - 响应参数变更
      - `+ tags`
  - **DeleteResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [ecn-attachment,connect-attachment,cfw-attachment]`
  - **ShowResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [ecn-attachment,connect-attachment,cfw-attachment]`
    - 响应参数变更
      - `+ tags`
  - **CreateResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [ecn-attachment,connect-attachment,cfw-attachment]`
  - **ListEnterpriseRouters**
    - 请求参数变更
      - `+ owned_by_self`
  - **ShowStaticRoute**
    - 响应参数变更
      - `+ route.attachments.priority`
  - **UpdateStaticRoute**
    - 响应参数变更
      - `+ route.attachments.priority`
  - **ListStaticRoutes**
    - 响应参数变更
      - `+ routes.attachments.priority`
  - **CreateStaticRoute**
    - 响应参数变更
      - `+ route.attachments.priority`
  - **ListEffectiveRoutes**
    - 响应参数变更
      - `+ routes.address_group_id`
      - `+ routes.next_hops.priority`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateTags**
    - 请求参数变更
      - `+ tags.key`
      - `+ tags.value`
      - `* tags: list<Kv> -> list<KvItem>`
  - **DeleteTags**
    - 请求参数变更
      - `+ tags.key`
      - `+ tags.value`
      - `* tags: list<Kv> -> list<KvItem>`
  - **ShowResInstanceInfo**
    - 请求参数变更
      - `+ matches.key`
      - `+ matches.value`
      - `* matches: list<Kv> -> list<KvItem>`
    - 响应参数变更
      - `+ resources.tags.key`
      - `+ resources.tags.value`
      - `* resources.tags: list<Kv> -> list<KvItem>`

### HuaweiCloud SDK GA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListRegions**
    - 响应参数变更
      - `+ regions`
      - `- area_regions`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持接口`ShowInstanceBiactiveRegions`
- _解决问题_
  - 无
- _特性变更_
  - **ListConfigurations**
    - 响应参数变更
      - `+ quota`
      - `+ configurations.mode`
  - **ListConfigurationTemplates**
    - 响应参数变更
      - `+ quota`
      - `+ configurations.mode`
  - **ShowInstanceConfiguration**
    - 响应参数变更
      - `+ mode`
      - `+ id`
  - **ListConfigurationDatastores**
    - 响应参数变更
      - `+ datastores.mode`
  - **ShowQuotas**
    - 请求参数变更
      - `+ datastore_type`
      - `+ mode`
  - **ListInstances**
    - 响应参数变更
      - `+ instances.datastore.whole_version`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持接口`DownloadBackup`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK HSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPorts**
    - 请求参数变更
      - `* host_id: optional -> required`
  - **ListVulnerabilities**
    - 响应参数变更
      - `+ data_list.verify_num`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 支持以下接口：
    - `ListDeviceTunnels`
    - `AddTunnel`
    - `ShowDeviceTunnel`
    - `CloseDeviceTunnel`
    - `DeleteDeviceTunnel`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateInstanceByEngine**
    - 请求参数变更
      - `+ disk_encrypted_enable`
      - `+ disk_encrypted_key`

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持接口`UpdateLogStream`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateLogGroup**
    - 请求参数变更
      - `+ tags`
  - **CreateLogGroup**
    - 请求参数变更
      - `+ tags`
  - **CreateLogStream**
    - 请求参数变更
      - `+ enterprise_project_name`
      - `+ ttl_in_days`
      - `+ tags`
      - `+ log_stream_name: enum value [lts-stream-13ci]`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowInstanceExtendProductInfo**
    - 响应参数变更
      - `+ monthly`
      - `+ hourly`
      - `- engine`
      - `- versions`
      - `- products`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ValidateConsumedMessage**
    - 请求参数变更
      - `+ engine: enum value [reliability]`
  - **ListInstances**
    - 请求参数变更
      - `+ engine: enum value [reliability]`

### HuaweiCloud SDK SMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowConfigSetting**
    - 响应参数变更
      - `* configurations: string -> list<ConfigBody>`

# 0.1.49 2023-07-20

### HuaweiCloud SDK DLI

- _新增特性_
  - 支持数据湖探索服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateIndividualStreamJob**
    - 请求参数变更
      - `- publish_param`
  - **UpdateIndividualStreamJob**
    - 请求参数变更
      - `- publish_param`
  - **CreateMixJob**
    - 请求参数变更
      - `- publish_param`

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持以下接口：
    - `AttachShareBandwidth`
    - `AttachBatchPublicIp`
    - `DetachShareBandwidth`
    - `DetachBatchPublicIp`
    - `EnableNat64`
    - `DisableNat64`
    - `ListBandwidth`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ER

- _新增特性_
  - 支持以下接口：
    - `BatchCreateResourceTags`
    - `ShowQuotas`
    - `ListFlowLogs`
    - `CreateFlowLog`
    - `ShowFlowLog`
    - `UpdateFlowLog`
    - `DeleteFlowLog`
    - `EnableFlowLog`
    - `DisableFlowLog`
- _解决问题_
  - 无
- _特性变更_
  - **ListProjectTags**
    - 请求参数变更
      - `+ resource_type: enum value [ecn-attachment,connect-attachment,cfw-attachment]`
    - 响应参数变更
      - `+ tags`
  - **DeleteResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [ecn-attachment,connect-attachment,cfw-attachment]`
  - **ShowResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [ecn-attachment,connect-attachment,cfw-attachment]`
    - 响应参数变更
      - `+ tags`
  - **CreateResourceTag**
    - 请求参数变更
      - `+ resource_type: enum value [ecn-attachment,connect-attachment,cfw-attachment]`
  - **ListEnterpriseRouters**
    - 请求参数变更
      - `+ owned_by_self`
  - **ShowStaticRoute**
    - 响应参数变更
      - `+ route.attachments.priority`
  - **UpdateStaticRoute**
    - 响应参数变更
      - `+ route.attachments.priority`
  - **ListStaticRoutes**
    - 响应参数变更
      - `+ routes.attachments.priority`
  - **CreateStaticRoute**
    - 响应参数变更
      - `+ route.attachments.priority`
  - **ListEffectiveRoutes**
    - 响应参数变更
      - `+ routes.address_group_id`
      - `+ routes.next_hops.priority`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 支持接口`DeleteBatchTask`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`ListTopicPartitions`、`ListTopicProducers`
- _解决问题_
  - 无
- _特性变更_
  - **ListProducts**
    - 请求参数变更
      - `+ engine: enum value [kafka]`
  - **UpdateInstanceTopic**
    - 请求参数变更
      - `+ topics.topic_other_configs`
      - `+ topics.topic_desc`
  - **CreateInstanceTopic**
    - 请求参数变更
      - `+ topic_other_configs`
      - `+ topic_desc`
    - 响应参数变更
      - `+ id`
  - **ListInstanceTopics**
    - 请求参数变更
      - `- offset`
      - `- limit`
    - 响应参数变更
      - `+ topics.topic_other_configs`
      - `+ topics.topic_desc`
      - `+ topics.created_at`
  - **ListInstances**
    - 请求参数变更
      - `+ engine: enum value [kafka]`
  - **ResizeEngineInstance**
    - 请求参数变更
      - `+ engine: enum value [kafka]`

### HuaweiCloud SDK MetaStudio

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreatePictureModelingJob**
    - 请求参数变更
      - `- model_asset_id`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowInstanceExtendProductInfo**
    - 请求参数变更
      - `+ engine: enum value [rabbitmq]`
    - 响应参数变更
      - `+ engine`
      - `+ versions`
      - `+ products`
      - `- monthly`
      - `- hourly`
  - **ListProducts**
    - 请求参数变更
      - `+ engine: enum value [rabbitmq]`
  - **ResizeEngineInstance**
    - 请求参数变更
      - `+ engine: enum value [rabbitmq]`
  - **ShowEngineInstanceExtendProductInfo**
    - 请求参数变更
      - `+ engine: enum value [rabbitmq]`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListInstancesSupportFastRestore`
- _解决问题_
  - 无
- _特性变更_
  - **RestoreTables**
    - 请求参数变更
      - `+ is_fast_restore`

# 0.1.48 2023-07-13

### HuaweiCloud SDK OROAS

- _新增特性_
  - 支持运筹优化算法服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AS

- _新增特性_
  - 支持接口`ListGroupScheduledTasks`、`CreateGroupScheduledTask`、`UpdateGroupScheduledTask`、`DeleteGroupScheduledTask`
- _解决问题_
  - 无
- _特性变更_
  - **CreateScalingPolicy**
    - 请求参数变更
      - `+ scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
  - **UpdateScalingPolicy**
    - 请求参数变更
      - `+ scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
  - **ShowScalingPolicy**
    - 响应参数变更
      - `+ scaling_policy.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scaling_policy.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
  - **ListScalingPolicies**
    - 响应参数变更
      - `+ scaling_policies.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scaling_policies.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
  - **CreateScalingV2Policy**
    - 请求参数变更
      - `+ scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
  - **ListAllScalingV2Policies**
    - 响应参数变更
      - `+ scaling_policies.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scaling_policies.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
  - **UpdateScalingV2Policy**
    - 请求参数变更
      - `+ scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
  - **ShowScalingV2Policy**
    - 响应参数变更
      - `+ scaling_policy.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scaling_policy.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`
  - **ListScalingV2Policies**
    - 响应参数变更
      - `+ scaling_policies.scheduled_policy.recurrence_type: enum value [DAILY,WEEKLY,MONTHLY]`
      - `- scaling_policies.scheduled_policy.recurrence_type: enum value [Daily,Weekly,Monthly]`

### HuaweiCloud SDK AntiDDoS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowNewTaskStatus**
    - 请求参数变更
      - `* task_id: optional -> required`

### HuaweiCloud SDK CAE

- _新增特性_
  - 支持以下接口：
    - `ListDomains`
    - `CreateDomain`
    - `DeleteDomain`
    - `ListCertificates`
    - `CreateCertificate`
    - `UpdateCertificate`
    - `DeleteCertificate`
    - `ListTimerRules`
    - `CreateTimerRule`
    - `UpdateTimerRule`
    - `DeleteTimerRule`
    - `ShowExecutionResult`
- _解决问题_
  - 无
- _特性变更_
  - **DeleteVolume**
    - 响应参数变更
      - `+ kind`
      - `+ api_version`
      - `+ items`
  - **ListAgencies**
    - 响应参数变更
      - `+ agencies`
      - `- items`
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Agency]`
  - **CreateAgency**
    - 请求参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Agency]`
      - `- metadata.name: enum value [cae_trust]`
  - **ListEnvironments**
    - 响应参数变更
      - `- items.type`
      - `- items.status: enum value [error]`
  - **CreateEnvironment**
    - 请求参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Environment]`
      - `- metadata.type`
    - 响应参数变更
      - `+ job_id`
      - `- metadata`
      - `- kind`
      - `- api_version`
  - **CreateApplication**
    - 请求参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Application]`
  - **ListApplications**
    - 响应参数变更
      - `+ items.annotations`
      - `+ items.created_at`
      - `+ items.updated_at`
  - **ListComponentSnapshots**
    - 响应参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [ComponentSnapshot]`
      - `+ items.context.app_id`
      - `+ items.context.available_replica`
      - `+ items.context.build`
      - `+ items.context.build_id`
      - `+ items.context.build_log_id`
      - `+ items.context.env_id`
      - `+ items.context.id`
      - `+ items.context.image_url`
      - `+ items.context.job_id`
      - `+ items.context.log_group_id`
      - `+ items.context.log_stream_id`
      - `+ items.context.name`
      - `+ items.context.operation`
      - `+ items.context.operation_status`
      - `+ items.context.replica`
      - `+ items.context.resource_limit`
      - `+ items.context.runtime`
      - `+ items.context.source`
      - `+ items.context.status`
      - `+ items.context.version`
      - `+ items.context.created_at`
      - `+ items.context.updated_at`
      - `* items.context: object -> object<ComponentSnapshotContext>`
  - **ListComponentConfigurations**
    - 响应参数变更
      - `- kind: enum value [Configuration]`
      - `+ items.type: enum value [rds,cse,env,access,scaling,volume,healthCheck,lifecycle]`
  - **CreateComponentConfiguration**
    - 请求参数变更
      - `- kind: enum value [Configuration]`
      - `+ items.type: enum value [rds,cse,env,access,scaling,volume,healthCheck,lifecycle]`
  - **ListComponentEvents**
    - 响应参数变更
      - `- kind: enum value [ComponentEvent]`
      - `+ items.involved_object_kind`
      - `- items.involved_object: enum value [Component,ComponentInstance,ComponentScaling]`
  - **ListComponentInstances**
    - 响应参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [ComponentInstance]`
  - **ListVolumes**
    - 请求参数变更
      - `+ resource_type: enum value [obs]`
    - 响应参数变更
      - `- kind: enum value [Volume]`
      - `- items.resource_type: enum value [obs]`
  - **CreateVolume**
    - 请求参数变更
      - `- kind: enum value [Volume]`
      - `- spec.resource_type: enum value [obs]`
  - **RetryJob**
    - 请求参数变更
      - `+ X-Enterprise-Project-ID`
      - `+ X-Environment-ID`
  - **ShowJob**
    - 请求参数变更
      - `+ X-Environment-ID`
    - 响应参数变更
      - `- kind: enum value [Job]`
      - `+ spec.progress`
  - **ShowComponent**
    - 响应参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Component]`
      - `- metadata.jod_id`
      - `- metadata.status`
      - `- metadata.type`
      - `+ spec.resource_limit`
      - `+ spec.build_log_id`
      - `- spec.log_strategy`
      - `+ spec.runtime: enum value [Docker,Java8,Java11,Tomcat8,Tomcat9,Python3,Nodejs8,Php7]`
  - **UpdateComponent**
    - 请求参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Component]`
      - `- metadata.created_at`
      - `- metadata.id`
      - `- metadata.jod_id`
      - `- metadata.status`
      - `- metadata.type`
      - `- metadata.updated_at`
      - `* metadata: object<Metadata> -> object<UpdateComponentRequestMetadata>`
      - `+ spec.runtime`
      - `+ spec.replica`
      - `- spec.log_strategy`
  - **ExecuteAction**
    - 请求参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Action]`
      - `* spec.source: object<Source> -> object<ActionOnComponentSource>`
  - **CreateComponent**
    - 请求参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Component]`
    - 响应参数变更
      - `- api_version: enum value [v1]`
      - `- kind: enum value [Component]`
      - `- metadata.jod_id`
      - `- metadata.status`
      - `- metadata.type`
      - `+ spec.resource_limit`
      - `- spec.access_info`
      - `- spec.build_id`
      - `- spec.image_url`
      - `- spec.job_id`
      - `- spec.log_strategy`
      - `+ spec.runtime: enum value [Docker,Java8,Java11,Tomcat8,Tomcat9,Python3,Nodejs8,Php7]`
      - `* spec: object<ComponentSpec> -> object<CreateComponentSpec>`
  - **ListComponents**
    - 响应参数变更
      - `+ items.created_at`
      - `+ items.updated_at`
      - `- items.status`
      - `+ items.spec.resource_limit`
      - `+ items.spec.build_log_id`
      - `- items.spec.log_strategy`
      - `+ items.spec.runtime: enum value [Docker,Java8,Java11,Tomcat8,Tomcat9,Python3,Nodejs8,Php7]`

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **PushTranscriberJobs**
    - 请求参数变更
      - `+ Enterprise-Project-Id`

### HuaweiCloud SDK VPC

- _新增特性_
  - 支持以下接口：
    - `ListApiVersion`
    - `NeutronListPorts`
    - `NeutronCreatePort`
    - `NeutronShowPort`
    - `NeutronUpdatePort`
    - `NeutronDeletePort`
    - `NeutronListNetworks`
    - `NeutronCreateNetwork`
    - `NeutronShowNetwork`
    - `NeutronUpdateNetwork`
    - `NeutronDeleteNetwork`
    - `NeutronListSubnets`
    - `NeutronCreateSubnet`
    - `NeutronShowSubnet`
    - `NeutronUpdateSubnet`
    - `NeutronDeleteSubnet`
    - `NeutronListRouters`
    - `NeutronCreateRouter`
    - `NeutronShowRouter`
    - `NeutronUpdateRouter`
    - `NeutronDeleteRouter`
    - `NeutronAddRouterInterface`
    - `NeutronRemoveRouterInterface`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.47 2023-07-06

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - `CodeCheck`更名为`CodeArtsCheck`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpgradeCluster**
    - 响应参数变更
      - `+ metadata`
      - `+ spec`
      - `- uid`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDomainDetailByName**
    - 响应参数变更
      - `+ domain.sources.weight`
  - **ShowDomainFullConfig**
    - 响应参数变更
      - `+ configs.business_type`
      - `+ configs.service_area`
      - `+ configs.sources.weight`
  - **UpdateDomainFullConfig**
    - 请求参数变更
      - `+ configs.business_type`
      - `+ configs.service_area`
      - `+ configs.sources.weight`

### HuaweiCloud SDK EVS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateVolume**
    - 请求参数变更
      - `+ volume.iops`
      - `+ volume.throughput`
      - `+ volume.volume_type: enum value [GPSSD2,ESSD2]`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **DeleteGaussMySqlReadonlyNode**
    - 响应参数变更
      - `+ order_id`
  - **DeleteGaussMySqlInstance**
    - 响应参数变更
      - `+ order_id`
  - **ShowSqlFilterRule**
    - 请求参数变更
      - `+ sql_type`
      - `- type`

### HuaweiCloud SDK HSS

- _新增特性_
  - 支持接口`ListJarPackageStatistics`、`ListJarPackageHostInfo`、`ListHostVuls`
- _解决问题_
  - 无
- _特性变更_
  - **ListProtectionServer**
    - 响应参数变更
      - `+ data_list.version`
      - `+ data_list.vault_id`
      - `+ data_list.vault_name`
      - `+ data_list.vault_size`
      - `+ data_list.vault_used`
      - `+ data_list.vault_allocated`
      - `+ data_list.vault_charging_mode`
      - `+ data_list.vault_status`
      - `+ data_list.backup_policy_id`
      - `+ data_list.backup_policy_name`
      - `+ data_list.backup_policy_enabled`
      - `+ data_list.resources_num`
  - **UpdateProtectionPolicy**
    - 请求参数变更
      - `+ process_whitelist`
  - **ListProtectionPolicy**
    - 请求参数变更
      - `+ protect_policy_id`
    - 响应参数变更
      - `+ data_list.process_whitelist`
  - **ListRiskConfigs**
    - 请求参数变更
      - `+ group_id`
  - **ListHostStatus**
    - 响应参数变更
      - `+ data_list.open_time`
  - **ListVulnerabilities**
    - 请求参数变更
      - `+ repair_priority`
      - `+ handle_status`
      - `+ cve_id`
      - `+ label_list`
      - `+ status`
      - `+ asset_value`
      - `+ group_name`
    - 响应参数变更
      - `+ data_list.cve_list`
      - `+ data_list.patch_url`
      - `+ data_list.repair_priority`
      - `+ data_list.hosts_num`
      - `+ data_list.repair_success_num`
      - `+ data_list.fixed_num`
      - `+ data_list.ignored_num`
  - **ListVulHosts**
    - 请求参数变更
      - `+ asset_value`
      - `+ group_name`
      - `+ handle_status`
      - `+ severity_level`
      - `+ is_affect_business`
    - 响应参数变更
      - `+ data_list.agent_id`
      - `+ data_list.app_path`
      - `+ data_list.region_name`
      - `+ data_list.public_ip`
      - `+ data_list.private_ip`
      - `+ data_list.group_id`
      - `+ data_list.group_name`
      - `+ data_list.os_type`
      - `+ data_list.asset_value`
      - `+ data_list.is_affect_business`
      - `+ data_list.first_scan_time`
      - `+ data_list.scan_time`
  - **ChangeVulStatus**
    - 请求参数变更
      - `+ remark`
      - `+ select_type`
      - `+ type`
      - `+ host_data_list`
  - **ListHostProtectHistoryInfo**
    - 响应参数变更
      - `+ data_list.file_operation`
      - `+ data_list.host_name`
      - `+ data_list.host_ip`
  - **ListHostRaspProtectHistoryInfo**
    - 响应参数变更
      - `+ data_list.host_ip`
      - `+ data_list.host_name`
  - **ListHostGroups**
    - 响应参数变更
      - `+ data_list.is_outside`
  - **StartProtection**
    - 请求参数变更
      - `+ backup_resources`
      - `+ create_protection_policy.process_whitelist`
  - **ListSecurityEvents**
    - 请求参数变更
      - `+ event_class_ids`
  - **ChangeEvent**
    - 请求参数变更
      - `+ container_name`
      - `+ container_id`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowPointTemplate**
    - 响应参数变更
      - `* : string -> binary`
  - **ShowPoints**
    - 响应参数变更
      - `* : string -> binary`
  - **DownloadAppVersion**
    - 响应参数变更
      - `* : string -> binary`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateSqlAlarmRule**
    - 请求参数变更
      - `+ notification_save_rule.template_name`
  - **CreateSqlAlarmRule**
    - 请求参数变更
      - `+ notification_save_rule.template_name`
  - **UpdateKeywordsAlarmRule**
    - 请求参数变更
      - `+ notification_save_rule.template_name`
  - **CreateKeywordsAlarmRule**
    - 请求参数变更
      - `+ notification_save_rule.template_name`
  - **ListAccessConfig**
    - 响应参数变更
      - `+ log_split`
      - `+ binary_collect`
      - `+ result.log_split`
      - `+ result.binary_collect`
      - `+ result.access_config_type: enum value [K8S_CCE]`
      - `+ result.access_config_detail.stdout`
      - `+ result.access_config_detail.stderr`
      - `+ result.access_config_detail.pathType`
      - `+ result.access_config_detail.namespaceRegex`
      - `+ result.access_config_detail.podNameRegex`
      - `+ result.access_config_detail.containerNameRegex`
      - `+ result.access_config_detail.includeLabels`
      - `+ result.access_config_detail.excludeLabels`
      - `+ result.access_config_detail.includeEnvs`
      - `+ result.access_config_detail.excludeEnvs`
      - `+ result.access_config_detail.logLabels`
      - `+ result.access_config_detail.logEnvs`
      - `+ result.access_config_detail.includeK8sLabels`
      - `+ result.access_config_detail.excludeK8sLabels`
      - `+ result.access_config_detail.logK8s`
      - `* result.access_config_detail.format.single: object<AccessConfigFormatSingle> -> object<AccessConfigFormatSingleCreate>`
      - `* result.access_config_detail.format.multi: object<AccessConfigFormatMutil> -> object<AccessConfigFormatMutilCreate>`
      - `* result.access_config_detail.format: object<AccessConfigFormat> -> object<AccessConfigFormatCreate>`
      - `* result.access_config_detail.windows_log_info: object<AccessConfigWindowsLogInfo> -> object<AccessConfigWindowsLogInfoCreate>`
      - `* result.access_config_detail: object<AccessConfigDeatil> -> object<AccessConfigDeatilCreate>`
  - **UpdateAccessConfig**
    - 请求参数变更
      - `+ log_split`
      - `+ binary_collect`
      - `+ access_config_detail.stdout`
      - `+ access_config_detail.stderr`
      - `+ access_config_detail.pathType`
      - `+ access_config_detail.namespaceRegex`
      - `+ access_config_detail.podNameRegex`
      - `+ access_config_detail.containerNameRegex`
      - `+ access_config_detail.includeLabels`
      - `+ access_config_detail.excludeLabels`
      - `+ access_config_detail.includeEnvs`
      - `+ access_config_detail.excludeEnvs`
      - `+ access_config_detail.logLabels`
      - `+ access_config_detail.logEnvs`
      - `+ access_config_detail.includeK8sLabels`
      - `+ access_config_detail.excludeK8sLabels`
      - `+ access_config_detail.logK8s`
      - `* access_config_detail.format.single: object<AccessConfigFormatSingle> -> object<AccessConfigFormatSingleCreate>`
      - `* access_config_detail.format.multi: object<AccessConfigFormatMutil> -> object<AccessConfigFormatMutilCreate>`
      - `* access_config_detail.format: object<AccessConfigFormat> -> object<AccessConfigFormatCreate>`
      - `* access_config_detail.windows_log_info: object<AccessConfigWindowsLogInfo> -> object<AccessConfigWindowsLogInfoCreate>`
      - `* access_config_detail: object<AccessConfigDeatil> -> object<AccessConfigDeatilCreate>`
    - 响应参数变更
      - `+ log_split`
      - `+ binary_collect`
      - `+ access_config_type: enum value [K8S_CCE]`
      - `+ access_config_detail.stdout`
      - `+ access_config_detail.stderr`
      - `+ access_config_detail.pathType`
      - `+ access_config_detail.namespaceRegex`
      - `+ access_config_detail.podNameRegex`
      - `+ access_config_detail.containerNameRegex`
      - `+ access_config_detail.includeLabels`
      - `+ access_config_detail.excludeLabels`
      - `+ access_config_detail.includeEnvs`
      - `+ access_config_detail.excludeEnvs`
      - `+ access_config_detail.logLabels`
      - `+ access_config_detail.logEnvs`
      - `+ access_config_detail.includeK8sLabels`
      - `+ access_config_detail.excludeK8sLabels`
      - `+ access_config_detail.logK8s`
      - `* access_config_detail.format.single: object<AccessConfigFormatSingle> -> object<AccessConfigFormatSingleCreate>`
      - `* access_config_detail.format.multi: object<AccessConfigFormatMutil> -> object<AccessConfigFormatMutilCreate>`
      - `* access_config_detail.format: object<AccessConfigFormat> -> object<AccessConfigFormatCreate>`
      - `* access_config_detail.windows_log_info: object<AccessConfigWindowsLogInfo> -> object<AccessConfigWindowsLogInfoCreate>`
      - `* access_config_detail: object<AccessConfigDeatil> -> object<AccessConfigDeatilCreate>`
  - **CreateAccessConfig**
    - 请求参数变更
      - `+ binary_collect`
      - `+ log_split`
      - `+ access_config_type: enum value [K8S_CCE]`
      - `+ access_config_detail.stdout`
      - `+ access_config_detail.stderr`
      - `+ access_config_detail.pathType`
      - `+ access_config_detail.namespaceRegex`
      - `+ access_config_detail.podNameRegex`
      - `+ access_config_detail.containerNameRegex`
      - `+ access_config_detail.includeLabels`
      - `+ access_config_detail.excludeLabels`
      - `+ access_config_detail.includeEnvs`
      - `+ access_config_detail.excludeEnvs`
      - `+ access_config_detail.logLabels`
      - `+ access_config_detail.logEnvs`
      - `+ access_config_detail.includeK8sLabels`
      - `+ access_config_detail.excludeK8sLabels`
      - `+ access_config_detail.logK8s`
    - 响应参数变更
      - `+ log_split`
      - `+ binary_collect`
      - `+ access_config_type: enum value [K8S_CCE]`
      - `+ access_config_detail.stdout`
      - `+ access_config_detail.stderr`
      - `+ access_config_detail.pathType`
      - `+ access_config_detail.namespaceRegex`
      - `+ access_config_detail.podNameRegex`
      - `+ access_config_detail.containerNameRegex`
      - `+ access_config_detail.includeLabels`
      - `+ access_config_detail.excludeLabels`
      - `+ access_config_detail.includeEnvs`
      - `+ access_config_detail.excludeEnvs`
      - `+ access_config_detail.logLabels`
      - `+ access_config_detail.logEnvs`
      - `+ access_config_detail.includeK8sLabels`
      - `+ access_config_detail.excludeK8sLabels`
      - `+ access_config_detail.logK8s`
      - `* access_config_detail.format.single: object<AccessConfigFormatSingle> -> object<AccessConfigFormatSingleCreate>`
      - `* access_config_detail.format.multi: object<AccessConfigFormatMutil> -> object<AccessConfigFormatMutilCreate>`
      - `* access_config_detail.format: object<AccessConfigFormat> -> object<AccessConfigFormatCreate>`
      - `* access_config_detail.windows_log_info: object<AccessConfigWindowsLogInfo> -> object<AccessConfigWindowsLogInfoCreate>`
      - `* access_config_detail: object<AccessConfigDeatil> -> object<AccessConfigDeatilCreate>`
  - **DeleteAccessConfig**
    - 响应参数变更
      - `+ log_split`
      - `+ binary_collect`
      - `+ result.log_split`
      - `+ result.binary_collect`
      - `+ result.access_config_type: enum value [K8S_CCE]`
      - `+ result.access_config_detail.stdout`
      - `+ result.access_config_detail.stderr`
      - `+ result.access_config_detail.pathType`
      - `+ result.access_config_detail.namespaceRegex`
      - `+ result.access_config_detail.podNameRegex`
      - `+ result.access_config_detail.containerNameRegex`
      - `+ result.access_config_detail.includeLabels`
      - `+ result.access_config_detail.excludeLabels`
      - `+ result.access_config_detail.includeEnvs`
      - `+ result.access_config_detail.excludeEnvs`
      - `+ result.access_config_detail.logLabels`
      - `+ result.access_config_detail.logEnvs`
      - `+ result.access_config_detail.includeK8sLabels`
      - `+ result.access_config_detail.excludeK8sLabels`
      - `+ result.access_config_detail.logK8s`
      - `* result.access_config_detail.format.single: object<AccessConfigFormatSingle> -> object<AccessConfigFormatSingleCreate>`
      - `* result.access_config_detail.format.multi: object<AccessConfigFormatMutil> -> object<AccessConfigFormatMutilCreate>`
      - `* result.access_config_detail.format: object<AccessConfigFormat> -> object<AccessConfigFormatCreate>`
      - `* result.access_config_detail.windows_log_info: object<AccessConfigWindowsLogInfo> -> object<AccessConfigWindowsLogInfoCreate>`
      - `* result.access_config_detail: object<AccessConfigDeatil> -> object<AccessConfigDeatilCreate>`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RunCreateVideoModerationJob**
    - 请求参数变更
      - `+ data.language`

### HuaweiCloud SDK RAM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **AssociateResourceShare**
    - 响应参数变更
      - `+ resource_share_associations.status_message`
  - **DisassociateResourceShare**
    - 响应参数变更
      - `+ resource_share_associations.status_message`
  - **SearchResourceShareAssociations**
    - 响应参数变更
      - `+ resource_share_associations.status_message`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSecurityGroupRules**
    - 请求参数变更
      - `+ remote_ip_prefix`

# 0.1.46 2023-06-29

### HuaweiCloud SDK IdentityCenter

- _新增特性_
  - 支持IAM 身份中心服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK WorkspaceApp

- _新增特性_
  - 支持云应用服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Config

- _新增特性_
  - 支持配置审计服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`ListTemplateVersions`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **DeleteTag**
    - 请求参数变更
      - `+ resource_type: enum value [cloud-connection,bandwidth-package]`
      - `+ resource_type: enum value [cc,bwp]`
  - **ListCloudConnections**
    - 请求参数变更
      - `+ used_scene`
    - 响应参数变更
      - `+ cloud_connections.tags`
  - **CreateCloudConnection**
    - 响应参数变更
      - `+ cloud_connection.tags`
  - **ShowCloudConnection**
    - 响应参数变更
      - `+ cloud_connection.tags`
  - **UpdateCloudConnection**
    - 响应参数变更
      - `+ cloud_connection.tags`
  - **ListDomainTags**
    - 请求参数变更
      - `+ resource_type: enum value [cloud-connection,bandwidth-package]`
      - `+ resource_type: enum value [cc,bwp]`
  - **BatchCreateDeleteTags**
    - 请求参数变更
      - `+ resource_type: enum value [cloud-connection,bandwidth-package]`
      - `+ resource_type: enum value [cc,bwp]`
  - **ListTags**
    - 请求参数变更
      - `+ resource_type: enum value [cloud-connection,bandwidth-package]`
      - `+ resource_type: enum value [cc,bwp]`
  - **CreateTag**
    - 请求参数变更
      - `+ resource_type: enum value [cloud-connection,bandwidth-package]`
      - `+ resource_type: enum value [cc,bwp]`
  - **ListQuotas**
    - 请求参数变更
      - `* quota_type: optional -> required`
  - **CreateBandwidthPackage**
    - 请求参数变更
      - `+ bandwidth_package.spec_code`
  - **ListResourceByFilterTag**
    - 请求参数变更
      - `+ resource_type: enum value [cloud-connection,bandwidth-package]`
      - `+ resource_type: enum value [cc,bwp]`

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ExecuteClusterSwitchover`、`ShowJobInfo`
- _解决问题_
  - 无
- _特性变更_
  - **ListConfigTemplates**
    - 响应参数变更
      - `+ config_templates.created_at`
  - **CreateInstance**
    - 请求参数变更
      - `+ template_id`

### HuaweiCloud SDK DSC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateProductOrder**
    - 请求参数变更
      - `+ charging_mode`
      - `+ cloud_service_type`
      - `+ composite_product_id`
      - `+ discount_id`
      - `+ is_auto_renew`
      - `+ period_num`
      - `+ period_type`
      - `+ product_infos`
      - `+ promotion_activity_id`
      - `+ promotion_info`
      - `+ region_id`
      - `- chargingMode`
      - `- cloudServiceType`
      - `- compositeProductId`
      - `- discountId`
      - `- isAutoRenew`
      - `- periodNum`
      - `- periodType`
      - `- productInfos`
      - `- promotionActivityId`
      - `- promotionInfo`
      - `- regionId`
  - **ShowSpecification**
    - 响应参数变更
      - `* orderInfos.productInfo: list<ProductInfoBean> -> object<ProductInfo>`
  - **ChangeRule**
    - 请求参数变更
      - `* body: object<RuleRequest> -> object<RuleChangeRequest>`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `ListDrugJob`
    - `CancelDrugJob`
    - `DeleteDrugJob`
    - `CreateOptmJob`
    - `ShowOptmJob`
    - `CreateSynthesisJob`
    - `ShowSynthesisJob`
    - `CreateDockingJob`
    - `ShowDockingJob`
    - `CreateFepJob`
    - `ShowFepJob`

### HuaweiCloud SDK GA

- _新增特性_
  - 支持以下接口：
    - `ListIpGroups`
    - `CreateIpGroup`
    - `ShowIpGroup`
    - `UpdateIpGroup`
    - `DeleteIpGroup`
    - `AddIpGroupIp`
    - `RemoveIpGroupIp`
    - `AssociateListener`
    - `DisassociateListener`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateCommand**
    - 响应参数变更
      - `+ error_msg`
      - `+ error_code`
  - **ListProperties**
    - 响应参数变更
      - `+ error_msg`
      - `+ error_code`
  - **UpdateProperties**
    - 响应参数变更
      - `+ error_msg`
      - `+ error_code`

### HuaweiCloud SDK ServiceStage

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowComponentConfigurations`、`CreateComponentConfiguration`、`ShowComponentConfiguration`、`CompareComponentConfiguration`
  - **ModifyApplicationConfiguration**
    - 响应参数变更
      - `+ environment_id`
      - `+ application_id`
      - `* configuration: list<object> -> object`
  - **ShowComponentInfo**
    - 响应参数变更
      - `- host_aliases`
      - `- dns_policy`
      - `- enable_sermant_injection`
      - `- workload_name`
      - `- workload_kind`
      - `- dns_config`
      - `- pod_labels`
      - `- security_context`
      - `- deploy_strategy.rolling_release.termination_seconds`
      - `- deploy_strategy.rolling_release.fail_strategy`
      - `- deploy_strategy.gray_release.deployment_mode`
      - `- deploy_strategy.gray_release.replica_surge_mode`
      - `- deploy_strategy.gray_release.rule_match_mode`
      - `- deploy_strategy.gray_release.rules`
      - `- liveness_probe.period_seconds`
      - `- liveness_probe.success_threshold`
      - `- liveness_probe.failure_threshold`
      - `- liveness_probe.http_headers`
  - **ModifyComponent**
    - 请求参数变更
      - `- pod_labels`
      - `- enable_sermant_injection`
      - `- host_aliases`
      - `- dns_policy`
      - `- dns_config`
      - `- security_context`
      - `- workload_kind`
      - `- deploy_strategy.rolling_release.termination_seconds`
      - `- deploy_strategy.rolling_release.fail_strategy`
      - `- deploy_strategy.gray_release.deployment_mode`
      - `- deploy_strategy.gray_release.replica_surge_mode`
      - `- deploy_strategy.gray_release.rule_match_mode`
      - `- deploy_strategy.gray_release.rules`
      - `- liveness_probe.period_seconds`
      - `- liveness_probe.success_threshold`
      - `- liveness_probe.failure_threshold`
      - `- liveness_probe.http_headers`
  - **CreateComponent**
    - 请求参数变更
      - `- workload_name`
      - `- pod_labels`
      - `- enterprise_project_id`
      - `- enable_sermant_injection`
      - `- host_aliases`
      - `- dns_policy`
      - `- dns_config`
      - `- security_context`
      - `- workload_kind`
      - `- deploy_strategy.rolling_release.termination_seconds`
      - `- deploy_strategy.rolling_release.fail_strategy`
      - `- deploy_strategy.gray_release.deployment_mode`
      - `- deploy_strategy.gray_release.replica_surge_mode`
      - `- deploy_strategy.gray_release.rule_match_mode`
      - `- deploy_strategy.gray_release.rules`
      - `- liveness_probe.period_seconds`
      - `- liveness_probe.success_threshold`
      - `- liveness_probe.failure_threshold`
      - `- liveness_probe.http_headers`
  - **ShowComponentsInApplication**
    - 响应参数变更
      - `+ components.platform_type`
  - **ShowComponents**
    - 响应参数变更
      - `+ components.platform_type`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAddressGroup**
    - 响应参数变更
      - `+ address_group.tags`
  - **UpdateAddressGroup**
    - 响应参数变更
      - `+ address_group.tags`
  - **ListAddressGroup**
    - 响应参数变更
      - `+ address_groups.tags`
  - **CreateAddressGroup**
    - 响应参数变更
      - `+ address_group.tags`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **AddOrRemoveServicePermissions**
    - 请求参数变更
      - `+ permission_type`
    - 响应参数变更
      - `+ permission_type`
  - **UpdateEndpointService**
    - 响应参数变更
      - `- cidr_type`
  - **ListServicePermissionsDetails**
    - 响应参数变更
      - `+ permissions.permission_type`
  - **BatchAddEndpointServicePermissions**
    - 请求参数变更
      - `+ permission_type`
    - 响应参数变更
      - `+ permissions.permission_type`
  - **BatchRemoveEndpointServicePermissions**
    - 响应参数变更
      - `+ permissions.permission_type`
  - **UpdateEndpointServicePermissionDesc**
    - 响应参数变更
      - `+ permissions.permission_type`
  - **CreateEndpointService**
    - 响应参数变更
      - `- cidr_type`

### HuaweiCloud SDK VSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CancelTasks**
    - 响应参数变更
      - `+ task_status: enum value [ready]`
  - **CreateTasks**
    - 响应参数变更
      - `+ task_status: enum value [ready]`
  - **ShowTasks**
    - 响应参数变更
      - `+ task_status: enum value [ready]`
  - **ListTaskHistories**
    - 响应参数变更
      - `+ data.task_status: enum value [ready]`

### HuaweiCloud SDK Workspace

- _新增特性_
  - 支持以下接口：
    - `BatchLogoffDesktops`
    - `ListDesktopsEips`
    - `ApplyDesktopsInternet`
    - `AssociateDesktopsEip`
    - `BatchDisassociateDesktopsEip`
    - `ListUnusedDesktops`
    - `ListUsedDesktopInfo`
    - `ListUserGroups`
    - `CreateUserGroup`
    - `BatchDeleteUserGroups`
    - `UpdateUserGroup`
    - `DeleteUserGroup`
    - `RunActionsOnGroup`
    - `ListUsersOfGroup`
    - `BatchCreateUsers`
    - `ResetRandomPassword`
- _解决问题_
  - 无
- _特性变更_
  - **ListDesktops**
    - 请求参数变更
      - `+ pool_id`
  - **CreateDesktop**
    - 请求参数变更
      - `+ eip`
      - `+ security_groups.name`
      - `* security_groups: list<SecurityGroupInfo> -> list<SecurityGroup>`
  - **ResizeDesktop**
    - 响应参数变更
      - `+ job_id`
      - `* jobs: list<ResizeDesktopJobResult> -> list<ResizeDesktopJobResponse>`
  - **CreateAccessPolicy**
    - 请求参数变更
      - `+ policy_objects_list.object_type: enum value [USERGROUP]`
  - **ListAccessPolicyObjects**
    - 响应参数变更
      - `+ policy_objects_list.object_type: enum value [USERGROUP]`
      - `* policy_objects_list: list<AccessPolicyObjectInfo> -> list<AccessPolicyObject>`
  - **UpdateAccessPolicyObjects**
    - 请求参数变更
      - `+ policy_objects_list.object_type: enum value [USERGROUP]`
  - **ListProducts**
    - 响应参数变更
      - `- products.package_type`
  - **CreateDesktopUser**
    - 请求参数变更
      - `+ active_type`
      - `+ user_phone`
      - `+ password`
      - `* body: object<CreateUserReq> -> object<CreateUserRequest>`
  - **ListUsers**
    - 请求参数变更
      - `+ active_type`
    - 响应参数变更
      - `+ users.user_phone`
      - `+ users.active_type`
      - `+ users.is_pre_user`
  - **UpdateUserInfo**
    - 请求参数变更
      - `+ user_phone`
      - `+ active_type`
  - **ListUserDetail**
    - 响应参数变更
      - `+ user_detail.user_phone`
      - `+ user_detail.active_type`
      - `+ user_detail.is_pre_user`
  - **ShowAssistAuthConfig**
    - 响应参数变更
      - `+ otp_config_info.apply_rule`
  - **UpdateAssistAuthMethodConfig**
    - 请求参数变更
      - `+ otp_config_info.apply_rule`
  - **ShowDesktopDetail**
    - 响应参数变更
      - `+ desktop.internet_mode`
      - `+ desktop.is_attaching_eip`
      - `+ desktop.security_groups.name`
      - `* desktop.security_groups: list<SecurityGroupInfo> -> list<SecurityGroup>`
  - **ListDesktopsDetail**
    - 请求参数变更
      - `+ pool_id`
    - 响应参数变更
      - `+ desktops.internet_mode`
      - `+ desktops.is_attaching_eip`
      - `+ desktops.security_groups.name`
      - `* desktops.security_groups: list<SecurityGroupInfo> -> list<SecurityGroup>`

# 0.1.45 2023-06-21

### HuaweiCloud SDK KooMessage

- _新增特性_
  - 支持云消息服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **DeleteGatewayResponseTypeV2**
    - 请求参数变更
      - `+ response_type: enum value [THIRD_AUTH_FAILURE,THIRD_AUTH_IDENTITIES_FAILURE,THIRD_AUTH_CONF_FAILURE]`
  - **ShowDetailsOfGatewayResponseTypeV2**
    - 请求参数变更
      - `+ response_type: enum value [THIRD_AUTH_FAILURE,THIRD_AUTH_IDENTITIES_FAILURE,THIRD_AUTH_CONF_FAILURE]`
  - **UpdateGatewayResponseTypeV2**
    - 请求参数变更
      - `+ response_type: enum value [THIRD_AUTH_FAILURE,THIRD_AUTH_IDENTITIES_FAILURE,THIRD_AUTH_CONF_FAILURE]`
  - **ShowPlugin**
    - 响应参数变更
      - `+ plugin_type: enum value [third_auth]`
  - **UpdatePlugin**
    - 请求参数变更
      - `+ plugin_type: enum value [third_auth]`
    - 响应参数变更
      - `+ plugin_type: enum value [third_auth]`
  - **CreatePlugin**
    - 请求参数变更
      - `+ plugin_type: enum value [third_auth]`
    - 响应参数变更
      - `+ plugin_type: enum value [third_auth]`
  - **ListPlugins**
    - 响应参数变更
      - `+ plugins.plugin_type: enum value [third_auth]`
  - **AttachApiToPlugin**
    - 响应参数变更
      - `+ attached_plugins.plugin_type: enum value [third_auth]`
  - **AttachPluginToApi**
    - 响应参数变更
      - `+ attached_plugins.plugin_type: enum value [third_auth]`
  - **ListApiAttachedPlugins**
    - 响应参数变更
      - `+ plugins.plugin_type: enum value [third_auth]`
  - **ListApiAttachablePlugins**
    - 响应参数变更
      - `+ plugins.plugin_type: enum value [third_auth]`
  - **ShowDetailsOfVpcChannelV2**
    - 响应参数变更
      - `+ microservice_info.cce_service_info`
      - `+ microservice_info.service_type: enum value [CCE_SERVICE]`
      - `+ microservice_info.cce_info.label_key`
      - `+ microservice_info.cce_info.label_value`
  - **UpdateVpcChannelV2**
    - 请求参数变更
      - `+ microservice_info.cce_service_info`
      - `+ microservice_info.service_type: enum value [CCE_SERVICE]`
      - `+ microservice_info.cce_info.label_key`
      - `+ microservice_info.cce_info.label_value`
    - 响应参数变更
      - `+ microservice_info.cce_service_info`
      - `+ microservice_info.service_type: enum value [CCE_SERVICE]`
      - `+ microservice_info.cce_info.label_key`
      - `+ microservice_info.cce_info.label_value`
  - **ImportMicroservice**
    - 请求参数变更
      - `+ cce_service_info`
      - `+ service_type: enum value [CCE_SERVICE]`
      - `+ cce_info.label_key`
      - `+ cce_info.label_value`
  - **CreateVpcChannelV2**
    - 请求参数变更
      - `+ microservice_info.cce_service_info`
      - `+ microservice_info.service_type: enum value [CCE_SERVICE]`
      - `+ microservice_info.cce_info.label_key`
      - `+ microservice_info.cce_info.label_value`
    - 响应参数变更
      - `+ microservice_info.cce_service_info`
      - `+ microservice_info.service_type: enum value [CCE_SERVICE]`
      - `+ microservice_info.cce_info.label_key`
      - `+ microservice_info.cce_info.label_value`
  - **ListVpcChannelsV2**
    - 响应参数变更
      - `+ vpc_channels.microservice_info.cce_service_info`
      - `+ vpc_channels.microservice_info.service_type: enum value [CCE_SERVICE]`
      - `+ vpc_channels.microservice_info.cce_info.label_key`
      - `+ vpc_channels.microservice_info.cce_info.label_value`

### HuaweiCloud SDK Classroom

- _新增特性_
  - 支持以下接口：
    - `ListPackages`
    - `ShowPackageDetail`
    - `ListExercises`
    - `ShowExerciseDetail`
    - `ExecuteExercise`
    - `ListAllDifficults`
    - `ListMyKnowledgePoints`
- _解决问题_
  - 无
- _特性变更_
  - **ApplyJudgement**
    - 请求参数变更
      - `+ runtime_type: enum value [javaScript]`

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateAutoRecord**
    - 响应参数变更
      - `- auto_record_mode`
      - `- app_id`
  - **CreateMixJob**
    - 响应参数变更
      - `+ mix_param.encode_template: enum value [1920*1080_30_4620,1920*1080_30_3150,1920*1080_15_3460,1920*1080_15_2080,1280*720_30_3420,1280*720_30_1710,1280*720_15_2260,1280*720_15_1130,640*480_30_1000,640*480_30_1500,640*480_15_500,480*360_30_490,480*360_15_320]`
      - `- mix_param.encode_template: enum value [1920x1080_30_4620,1920x1080_15_3460,1280x720_30_3420,1280x720_15_2260]`
  - **ShowMixJob**
    - 响应参数变更
      - `+ mix_param.encode_template: enum value [1920*1080_30_4620,1920*1080_30_3150,1920*1080_15_3460,1920*1080_15_2080,1280*720_30_3420,1280*720_30_1710,1280*720_15_2260,1280*720_15_1130,640*480_30_1000,640*480_30_1500,640*480_15_500,480*360_30_490,480*360_15_320]`
      - `- mix_param.encode_template: enum value [1920x1080_30_4620,1920x1080_15_3460,1280x720_30_3420,1280x720_15_2260]`
  - **UpdateMixJob**
    - 响应参数变更
      - `+ mix_param.encode_template: enum value [1920*1080_30_4620,1920*1080_30_3150,1920*1080_15_3460,1920*1080_15_2080,1280*720_30_3420,1280*720_30_1710,1280*720_15_2260,1280*720_15_1130,640*480_30_1000,640*480_30_1500,640*480_15_500,480*360_30_490,480*360_15_320]`
      - `- mix_param.encode_template: enum value [1920x1080_30_4620,1920x1080_15_3460,1280x720_30_3420,1280x720_15_2260]`

### HuaweiCloud SDK CloudTable

- _新增特性_
  - 支持以下接口：
    - `EnableComponent`
    - `ExpandClusterComponent`
    - `RebootCloudTableCluster`
    - `ShowClusterSetting`
    - `UpdateClusterSetting`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DNS

- _新增特性_
  - 支持接口`ShowDomainQuota`
- _解决问题_
  - 无
- _特性变更_
  - **ShowRecordSetWithLine**
    - 响应参数变更
      - `+ bundle`
  - **SetRecordSetsStatus**
    - 响应参数变更
      - `+ bundle`
  - **BatchUpdateRecordSetWithLine**
    - 响应参数变更
      - `+ bundle`
      - `+ recordsets.bundle`
  - **BatchDeleteRecordSetWithLine**
    - 响应参数变更
      - `+ bundle`
      - `+ recordsets.bundle`
  - **CreateRecordSetWithBatchLines**
    - 响应参数变更
      - `+ bundle`
      - `+ recordsets.bundle`

### HuaweiCloud SDK DWS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAvailableDisasterClusters**
    - 请求参数变更
      - `* primary_cluster_id: optional -> required`
      - `* standby_az_code: optional -> required`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持以下接口：
    - `ListDrugJob`
    - `CancelDrugJob`
    - `DeleteDrugJob`
    - `CreateOptmJob`
    - `ShowOptmJob`
    - `CreateSynthesisJob`
    - `ShowSynthesisJob`
    - `CreateDockingJob`
    - `ShowDockingJob`
    - `CreateFepJob`
    - `ShowFepJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateDbInstance**
    - 响应参数变更
      - `+ instance.ha.consistency_protocol`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateEdgeNode**
    - 请求参数变更
      - `+ npu_library_path`
      - `+ device_data_format`
      - `+ automatic_upgrade`
      - `+ device_data_record`
      - `+ metric_report`
      - `+ base_path.offline_cache_configs`
    - 响应参数变更
      - `+ device_data_record`
      - `+ device_data_format`
      - `+ metric_report`
      - `+ automatic_upgrade`
      - `+ base_path.offline_cache_configs`
  - **ShowEdgeNode**
    - 响应参数变更
      - `+ device_data_record`
      - `+ device_data_format`
      - `+ metric_report`
      - `+ npu_library_path`
      - `+ automatic_upgrade`
      - `+ base_path.offline_cache_configs`
  - **CreateEdgeApplicationVersion**
    - 请求参数变更
      - `+ container_settings.npu_type`
      - `+ container_settings.vnpu_template`
    - 响应参数变更
      - `+ container_settings.npu_type`
      - `+ container_settings.vnpu_template`
  - **ShowEdgeApplicationVersion**
    - 响应参数变更
      - `+ container_settings.npu_type`
      - `+ container_settings.vnpu_template`
  - **UpdateEdgeApplicationVersion**
    - 请求参数变更
      - `+ container_settings.npu_type`
      - `+ container_settings.vnpu_template`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListPredefinedTag`、`ListSimplifiedInstances`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.44 2023-06-15

### HuaweiCloud SDK CloudDeploy

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - `CloudDeploy`更名为`CodeArtsDeploy`

### HuaweiCloud SDK CCM

- _新增特性_
  - 支持以下接口：
    - `BatchCreateCaTags`
    - `BatchDeleteCaTags`
    - `ListCaTags`
    - `CreateCaTag`
    - `ListDomainCaTags`
    - `ListCaResourceInstances`
    - `CountCaResourceInstances`
    - `BatchCreateCertTags`
    - `BatchDeleteCertTags`
    - `ListCertTags`
    - `CreateCertTag`
    - `ListDomainCertTags`
    - `ListCertResourceInstances`
    - `CountCertResourceInstances`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持以下接口：
    - `ShowDatabaseAuthority`
    - `UpdateDatabaseAuthority`
    - `SyncIamUsers`
    - `ListDatabaseUsers`
    - `ShowDatabaseUser`
    - `UpdateDatabaseUserInfo`
    - `ShowDisasterProgress`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateAlarmSub**
    - 请求参数变更
      - `* enable: string -> int32`
    - 响应参数变更
      - `* enable: string -> int32`
  - **DeleteAlarmSub**
    - 响应参数变更
      - `* enable: string -> int32`
  - **ShowDisasterDetail**
    - 响应参数变更
      - `+ disaster_recovery`
      - `- start_time`
      - `- dr_type`
      - `- create_time`
      - `- name`
      - `- standby_cluster`
      - `- id`
      - `- dr_sync_period`
      - `- status`
      - `- primary_cluster`
  - **CreateAlarmSub**
    - 请求参数变更
      - `* enable: string -> int32`
    - 响应参数变更
      - `* enable: string -> int32`
  - **ListAlarmSubs**
    - 响应参数变更
      - `* alarm_subscriptions.enable: string -> int32`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListGaussMySqlErrorLog`、`ListGaussMySqlSlowLog`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateBatchTask**
    - 响应参数变更
      - `- task_progress.device_in_progress`
      - `- task_progress.rejected`
  - **ListBatchTasks**
    - 响应参数变更
      - `- batchtasks.task_progress.device_in_progress`
      - `- batchtasks.task_progress.rejected`
  - **ShowBatchTask**
    - 响应参数变更
      - `- batchtask.task_progress.device_in_progress`
      - `- batchtask.task_progress.rejected`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListClusters**
    - 响应参数变更
      - `+ clusters.eipId`
      - `+ clusters.eipAddress`
      - `+ clusters.eipv6Address`
  - **ShowClusterDetails**
    - 响应参数变更
      - `+ cluster.eipId`
      - `+ cluster.eipAddress`
      - `+ cluster.eipv6Address`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeGeneralText**
    - 请求参数变更
      - `+ single_orientation_mode`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ModifyCollation`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **SendDlqMessage**
    - 请求参数变更
      - `+ engine: enum value [reliability]`
  - **CreateRocketMqMigrationTask**
    - 请求参数变更
      - `+ type: enum value [kafka]`

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowVocabularies**
    - 请求参数变更
      - `+ offset`
      - `+ limit`

# 0.1.43 2023-06-08

### HuaweiCloud SDK iDME

- _新增特性_
  - 工业数字模型驱动引擎
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BCS

- _新增特性_
  - 支持接口`ListBcsEvents`、`ListBcsEventsStatistic`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateIndirectPartnerAccount**
    - 请求参数变更
      - `* amount: double -> bigdecimal`
  - **ListCustomerBillsMonthlyBreakDown**
    - 响应参数变更
      - `* details.past_months_amortized_amount: double -> bigdecimal`
      - `* details.amortized_cash_amount: double -> bigdecimal`
  - **ListIssuedCouponQuotas**
    - 响应参数变更
      - `* quotas.balance: double -> bigdecimal`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowNode**
    - 响应参数变更
      - `+ spec.extendParam.kube-reserved-mem`
      - `+ spec.extendParam.system-reserved-mem`
  - **UpdateNode**
    - 响应参数变更
      - `+ spec.extendParam.kube-reserved-mem`
      - `+ spec.extendParam.system-reserved-mem`
  - **DeleteNode**
    - 响应参数变更
      - `+ spec.extendParam.kube-reserved-mem`
      - `+ spec.extendParam.system-reserved-mem`
  - **CreateNode**
    - 请求参数变更
      - `+ spec.extendParam.kube-reserved-mem`
      - `+ spec.extendParam.system-reserved-mem`
    - 响应参数变更
      - `+ spec.extendParam.kube-reserved-mem`
      - `+ spec.extendParam.system-reserved-mem`
  - **ListNodes**
    - 响应参数变更
      - `+ items.spec.extendParam.kube-reserved-mem`
      - `+ items.spec.extendParam.system-reserved-mem`
  - **ShowNodePool**
    - 响应参数变更
      - `+ spec.type: enum value [pm]`
      - `+ spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `+ spec.nodeTemplate.extendParam.system-reserved-mem`
  - **UpdateNodePool**
    - 响应参数变更
      - `+ spec.type: enum value [pm]`
      - `+ spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `+ spec.nodeTemplate.extendParam.system-reserved-mem`
  - **DeleteNodePool**
    - 响应参数变更
      - `+ spec.type: enum value [pm]`
      - `+ spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `+ spec.nodeTemplate.extendParam.system-reserved-mem`
  - **CreateNodePool**
    - 请求参数变更
      - `+ spec.type: enum value [pm]`
      - `+ spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `+ spec.nodeTemplate.extendParam.system-reserved-mem`
    - 响应参数变更
      - `+ spec.type: enum value [pm]`
      - `+ spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `+ spec.nodeTemplate.extendParam.system-reserved-mem`
  - **ListNodePools**
    - 响应参数变更
      - `+ items.spec.type: enum value [pm]`
      - `+ items.spec.nodeTemplate.extendParam.kube-reserved-mem`
      - `+ items.spec.nodeTemplate.extendParam.system-reserved-mem`

### HuaweiCloud SDK CloudDeploy

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateDeploymentGroup**
    - 请求参数变更
      - `+ is_proxy_mode`

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListRecordSetsByZone**
    - 请求参数变更
      - `+ search_mode`
  - **CreateRecordSet**
    - 请求参数变更
      - `* body: object<CreateRecordSetReq> -> object<CreateRecordSetRequestBody>`
  - **CreateRecordSetWithLine**
    - 请求参数变更
      - `* body: object<CreateRecordSetWithLineReq> -> object<CreateRecordSetWithLineRequestBody>`
  - **ListPublicZones**
    - 请求参数变更
      - `+ search_mode`
  - **ListPrivateZones**
    - 请求参数变更
      - `+ search_mode`
  - **ListRecordSets**
    - 请求参数变更
      - `+ search_mode`

### HuaweiCloud SDK ECS

- _新增特性_
  - 支持接口`ChangeServerChargeMode`
- _解决问题_
  - 无
- _特性变更_
  - **CreateServers**
    - 请求参数变更
      - `+ server.nics.allowed_address_pairs`
  - **CreatePostPaidServers**
    - 请求参数变更
      - `+ server.nics.allowed_address_pairs`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListListeners**
    - 响应参数变更
      - `+ listeners.port_ranges`
  - **CreateListener**
    - 请求参数变更
      - `+ listener.port_ranges`
    - 响应参数变更
      - `+ listener.port_ranges`
  - **ShowListener**
    - 响应参数变更
      - `+ listener.port_ranges`
  - **UpdateListener**
    - 响应参数变更
      - `+ listener.port_ranges`
  - **ListPools**
    - 响应参数变更
      - `+ pools.any_port_enable`
  - **CreatePool**
    - 请求参数变更
      - `+ pool.any_port_enable`
    - 响应参数变更
      - `+ pool.any_port_enable`
  - **ShowPool**
    - 响应参数变更
      - `+ pool.any_port_enable`
  - **UpdatePool**
    - 响应参数变更
      - `+ pool.any_port_enable`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持以下接口：
    - `UpdateFuncSnapshot`
    - `ShowFuncSnapshotState`
    - `ShowResInstanceInfo`
    - `ShowProjectTagsList`
    - `CreateTags`
    - `DeleteTags`
    - `CreateVpcEndpoint`
    - `DeleteVpcEndpoint`
- _解决问题_
  - 无
- _特性变更_
  - **ListStatistics**
    - 响应参数变更
      - `* count.value: int32 -> number`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持以下接口：
    - `ListInstancesDetails`
    - `CreateDbInstance`
    - `ListParamGroupTemplates`
    - `ShowInstanceParamGroup`
    - `ListDbFlavors`
    - `ListDbBackups`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 支持接口`RetryBatchTask`、`StopBatchTask`
- _解决问题_
  - 无
- _特性变更_
  - **CreateBatchTask**
    - 响应参数变更
      - `+ task_progress.removed`
      - `+ task_progress.device_in_progress`
      - `+ task_progress.rejected`
  - **ListBatchTasks**
    - 响应参数变更
      - `+ batchtasks.task_progress.removed`
      - `+ batchtasks.task_progress.device_in_progress`
      - `+ batchtasks.task_progress.rejected`
  - **ShowBatchTask**
    - 请求参数变更
      - `+ task_detail_status`
      - `+ target`
    - 响应参数变更
      - `+ batchtask.task_progress.removed`
      - `+ batchtask.task_progress.device_in_progress`
      - `+ batchtask.task_progress.rejected`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`ListAvailableZones`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`AddIssueWorkHours`、`ListProjectWorkHoursType`
- _解决问题_
  - 无
- _特性变更_
  - **ShowProjectWorkHours**
    - 响应参数变更
      - `+ work_hours.work_hours_created_time`
      - `+ work_hours.work_hours_updated_time`
  - **ListProjectWorkHours**
    - 响应参数变更
      - `+ work_hours.work_hours_created_time`
      - `+ work_hours.work_hours_updated_time`
  - **ListIssueCustomFields**
    - 请求参数变更
      - `+ included_not_in_use`
    - 响应参数变更
      - `+ datas.create_time`
  - **ListIssuesV4**
    - 请求参数变更
      - `+ created_time_interval`
      - `+ closed_time_interval`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListInstanceTags`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 支持接口`SendDlqMessage`、`ValidateConsumedMessage`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`SendRocketMqDlqMessage`、`ValidateRocketMqConsumedMessage`
  - **CreateInstanceByEngine**
    - 请求参数变更
      - `+ product_id: enum value [c6.4u8g.cluster.small]`

### HuaweiCloud SDK TMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListResource**
    - 响应参数变更
      - `+ resources.tags`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAddressGroup**
    - 响应参数变更
      - `+ address_group.enterprise_project_id`
  - **UpdateAddressGroup**
    - 响应参数变更
      - `+ address_group.enterprise_project_id`
  - **ListAddressGroup**
    - 请求参数变更
      - `+ enterprise_project_id`
    - 响应参数变更
      - `+ address_groups.enterprise_project_id`
  - **CreateAddressGroup**
    - 请求参数变更
      - `+ address_group.enterprise_project_id`
    - 响应参数变更
      - `+ address_group.enterprise_project_id`

# 0.1.42 2023-06-01

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListMultiAccountTransferCoupons`、`ListMultiAccountRetrieveCoupons`、`ClaimEnterpriseMultiAccountCoupon`、`ReclaimEnterpriseMultiAccountCoupon`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateCustomerAccountAmount**
    - 请求参数变更
      - `* amount: double -> bigdecimal`
  - **ReclaimIndirectPartnerAccount**
    - 请求参数变更
      - `* amount: double -> bigdecimal`
  - **ReclaimToPartnerAccount**
    - 请求参数变更
      - `* amount: double -> bigdecimal`
  - **ListPartnerCouponsRecord**
    - 响应参数变更
      - `* records.operation_amount: double -> bigdecimal`
  - **ListCustomersBalancesDetail**
    - 响应参数变更
      - `* customer_balances.debt_amount: double -> bigdecimal`
      - `* customer_balances.amount: double -> bigdecimal`
  - **ShowCustomerMonthlySum**
    - 响应参数变更
      - `* consume_amount: double -> bigdecimal`
      - `* debt_amount: double -> bigdecimal`
      - `* coupon_amount: double -> bigdecimal`
      - `* flexipurchase_coupon_amount: double -> bigdecimal`
      - `* stored_value_card_amount: double -> bigdecimal`
      - `* cash_amount: double -> bigdecimal`
      - `* credit_amount: double -> bigdecimal`
      - `* writeoff_amount: double -> bigdecimal`
      - `* bill_sums.official_amount: double -> bigdecimal`
      - `* bill_sums.official_discount_amount: double -> bigdecimal`
      - `* bill_sums.truncated_amount: double -> bigdecimal`
      - `* bill_sums.consume_amount: double -> bigdecimal`
      - `* bill_sums.coupon_amount: double -> bigdecimal`
      - `* bill_sums.flexipurchase_coupon_amount: double -> bigdecimal`
      - `* bill_sums.stored_value_card_amount: double -> bigdecimal`
      - `* bill_sums.debt_amount: double -> bigdecimal`
      - `* bill_sums.writeoff_amount: double -> bigdecimal`
      - `* bill_sums.cash_amount: double -> bigdecimal`
      - `* bill_sums.credit_amount: double -> bigdecimal`
  - **UpdateCouponQuotas**
    - 请求参数变更
      - `* quota_amount: double -> bigdecimal`
  - **ListCouponQuotasRecords**
    - 响应参数变更
      - `* records.amount: double -> bigdecimal`
  - **ReclaimCouponQuotas**
    - 响应参数变更
      - `* simple_quota_infos.quota_balance: double -> bigdecimal`
  - **ShowCustomerAccountBalances**
    - 响应参数变更
      - `* debt_amount: double -> bigdecimal`
      - `* account_balances.amount: double -> bigdecimal`
      - `* account_balances.designated_amount: double -> bigdecimal`
      - `* account_balances.credit_amount: double -> bigdecimal`
  - **ListFreeResourceUsages**
    - 响应参数变更
      - `* free_resources.amount: double -> bigdecimal`
      - `* free_resources.original_amount: double -> bigdecimal`
  - **ListIssuedPartnerCoupons**
    - 响应参数变更
      - `* user_coupons.face_value: double -> bigdecimal`
      - `* user_coupons.balance: double -> bigdecimal`
  - **ListOnDemandResourceRatings**
    - 响应参数变更
      - `* amount: double -> bigdecimal`
      - `* discount_amount: double -> bigdecimal`
      - `* official_website_amount: double -> bigdecimal`
      - `* product_rating_results.amount: double -> bigdecimal`
      - `* product_rating_results.discount_amount: double -> bigdecimal`
      - `* product_rating_results.official_website_amount: double -> bigdecimal`
  - **ListSubcustomerMonthlyBills**
    - 响应参数变更
      - `* bill_sums.amount: double -> bigdecimal`
      - `* bill_sums.debt_amount: double -> bigdecimal`
      - `* bill_sums.adjustment_amount: double -> bigdecimal`
      - `* bill_sums.discount_amount: double -> bigdecimal`
      - `* bill_sums.account_details.amount: double -> bigdecimal`
  - **ListCustomerBillsMonthlyBreakDown**
    - 响应参数变更
      - `* details.usage: double -> bigdecimal`
      - `* details.free_resource_usage: double -> bigdecimal`
      - `* details.ri_usage: double -> bigdecimal`
      - `* details.consume_amount: double -> bigdecimal`
      - `* details.current_month_amortized_amount: double -> bigdecimal`
      - `* details.future_months_amortized_amount: double -> bigdecimal`
      - `* details.amortized_credit_amount: double -> bigdecimal`
      - `* details.amortized_coupon_amount: double -> bigdecimal`
      - `* details.amortized_flexipurchase_coupon_amount: double -> bigdecimal`
      - `* details.amortized_stored_value_card_amount: double -> bigdecimal`
      - `* details.amortized_bonus_amount: double -> bigdecimal`
  - **ListQuotaCoupons**
    - 响应参数变更
      - `* quotas.quota_value: double -> bigdecimal`
      - `* quotas.balance: double -> bigdecimal`
  - **ListIssuedCouponQuotas**
    - 响应参数变更
      - `* quotas.quota_value: double -> bigdecimal`

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListFreeResourceUsages**
    - 响应参数变更
      - `* free_resources.amount: double -> bigdecimal`
      - `* free_resources.original_amount: double -> bigdecimal`
  - **ListOnDemandResourceRatings**
    - 响应参数变更
      - `* amount: double -> bigdecimal`
      - `* discount_amount: double -> bigdecimal`
      - `* official_website_amount: double -> bigdecimal`
      - `* product_rating_results.amount: double -> bigdecimal`
      - `* product_rating_results.discount_amount: double -> bigdecimal`
      - `* product_rating_results.official_website_amount: double -> bigdecimal`

### HuaweiCloud SDK CBH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowNetworkConfiguration**
    - 请求参数变更
      - `+ nics.ip_address`
  - **ChangeInstanceNetwork**
    - 请求参数变更
      - `+ nics.ip_address`
  - **CreateInstance**
    - 请求参数变更
      - `+ server.nics.ip_address`

### HuaweiCloud SDK CBR

- _新增特性_
  - 支持接口`ShowSummary`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CBS

- _新增特性_
  - 支持接口`PostRequests`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowAddonInstance**
    - 响应参数变更
      - `+ metadata.alias`
      - `* metadata: object<Metadata> -> object<AddonMetadata>`
  - **UpdateAddonInstance**
    - 请求参数变更
      - `+ metadata.alias`
      - `* metadata: object<Metadata> -> object<AddonMetadata>`
    - 响应参数变更
      - `+ metadata.alias`
      - `* metadata: object<Metadata> -> object<AddonMetadata>`
  - **CreateAddonInstance**
    - 请求参数变更
      - `+ metadata.alias`
      - `* metadata: object<Metadata> -> object<AddonMetadata>`
    - 响应参数变更
      - `+ metadata.alias`
      - `* metadata: object<Metadata> -> object<AddonMetadata>`
  - **ListAddonInstances**
    - 响应参数变更
      - `+ items.metadata.alias`
      - `* items.metadata: object<Metadata> -> object<AddonMetadata>`
  - **ListAddonTemplates**
    - 响应参数变更
      - `+ items.metadata.alias`
      - `* items.metadata: object<Metadata> -> object<AddonMetadata>`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowLogs**
    - 请求参数变更
      - `+ start_time`
      - `+ end_time`
      - `- query_date`
  - **ShowDomainFullConfig**
    - 请求参数变更
      - `+ show_special_configs`
    - 响应参数变更
      - `- configs.error_code_cache.code: enum value [400,403,404,405,414,500,501,502,503,504]`
      - `+ configs.flexible_origin.back_sources.http_port`
      - `+ configs.flexible_origin.back_sources.https_port`
  - **UpdateDomainFullConfig**
    - 请求参数变更
      - `- configs.error_code_cache.code: enum value [400,403,404,405,414,500,501,502,503,504]`
      - `+ configs.flexible_origin.back_sources.http_port`
      - `+ configs.flexible_origin.back_sources.https_port`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 支持以下接口：
    - `UpdateDirectory`
    - `DeleteDirectory`
    - `ListTaskCases`
    - `CreateNewTask`
    - `CreateNewCase`
    - `ShowCase`
    - `UpdateNewCase`
    - `DeleteNewCase`
    - `ShowMergeReportLogsOutline`
    - `ShowTaskCaseAwChart`
    - `ShowMergeCaseDetail`
    - `ShowMergeTaskCase`
    - `BatchUpdateTaskStatus`
    - `CreateDirectory`
    - `DeleteNewTask`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateCase**
    - 请求参数变更
      - `+ directory_id`
      - `+ setup_contents`
      - `+ user_replicas`
      - `+ collect_log_policy`
      - `+ link_app_list`
      - `+ case_info`
      - `+ contents.content.content.rtmp_url`
      - `+ contents.content.content.flv_url`
      - `+ contents.content.content.bitrate_type`
      - `+ contents.content.content.duration`
      - `+ contents.content.content.retry_delay`
      - `+ contents.content.content.retry_time`
  - **UpdateTemp**
    - 请求参数变更
      - `+ contents.content.content.rtmp_url`
      - `+ contents.content.content.flv_url`
      - `+ contents.content.content.bitrate_type`
      - `+ contents.content.content.duration`
      - `+ contents.content.content.retry_delay`
      - `+ contents.content.content.retry_time`
  - **ShowTask**
    - 响应参数变更
      - `+ taskInfo.case_list.directory_id`
      - `+ taskInfo.case_list.setup_contents`
      - `+ taskInfo.case_list.user_replicas`
      - `+ taskInfo.case_list.collect_log_policy`
      - `+ taskInfo.case_list.link_app_list`
      - `+ taskInfo.case_list.case_info`
      - `+ taskInfo.case_list.contents.content.content.rtmp_url`
      - `+ taskInfo.case_list.contents.content.content.flv_url`
      - `+ taskInfo.case_list.contents.content.content.bitrate_type`
      - `+ taskInfo.case_list.contents.content.content.duration`
      - `+ taskInfo.case_list.contents.content.content.retry_delay`
      - `+ taskInfo.case_list.contents.content.content.retry_time`
  - **UpdateTask**
    - 请求参数变更
      - `+ case_list.directory_id`
      - `+ case_list.setup_contents`
      - `+ case_list.user_replicas`
      - `+ case_list.collect_log_policy`
      - `+ case_list.link_app_list`
      - `+ case_list.case_info`
      - `+ case_list.contents.content.content.rtmp_url`
      - `+ case_list.contents.content.content.flv_url`
      - `+ case_list.contents.content.content.bitrate_type`
      - `+ case_list.contents.content.content.duration`
      - `+ case_list.contents.content.content.retry_delay`
      - `+ case_list.contents.content.content.retry_time`
    - 响应参数变更
      - `+ taskInfo.case_list.directory_id`
      - `+ taskInfo.case_list.setup_contents`
      - `+ taskInfo.case_list.user_replicas`
      - `+ taskInfo.case_list.collect_log_policy`
      - `+ taskInfo.case_list.link_app_list`
      - `+ taskInfo.case_list.case_info`
      - `+ taskInfo.case_list.contents.content.content.rtmp_url`
      - `+ taskInfo.case_list.contents.content.content.flv_url`
      - `+ taskInfo.case_list.contents.content.content.bitrate_type`
      - `+ taskInfo.case_list.contents.content.content.duration`
      - `+ taskInfo.case_list.contents.content.content.retry_delay`
      - `+ taskInfo.case_list.contents.content.content.retry_time`
  - **UpdateTaskRelatedTestCase**
    - 响应参数变更
      - `+ taskInfo.case_list.directory_id`
      - `+ taskInfo.case_list.setup_contents`
      - `+ taskInfo.case_list.user_replicas`
      - `+ taskInfo.case_list.collect_log_policy`
      - `+ taskInfo.case_list.link_app_list`
      - `+ taskInfo.case_list.case_info`
      - `+ taskInfo.case_list.contents.content.content.rtmp_url`
      - `+ taskInfo.case_list.contents.content.content.flv_url`
      - `+ taskInfo.case_list.contents.content.content.bitrate_type`
      - `+ taskInfo.case_list.contents.content.content.duration`
      - `+ taskInfo.case_list.contents.content.content.retry_delay`
      - `+ taskInfo.case_list.contents.content.content.retry_time`

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RestorePtrRecord**
    - 请求参数变更
      - `* ptrdname: string -> object`
  - **ShowRecordSet**
    - 响应参数变更
      - `+ bundle`
  - **CreateEipRecordSet**
    - 响应参数变更
      - `+ enterprise_project_id`
  - **ShowPtrRecordSet**
    - 响应参数变更
      - `+ enterprise_project_id`
  - **ShowResourceTag**
    - 响应参数变更
      - `+ enterpriseProjectOrDefault`
  - **ListPrivateZones**
    - 请求参数变更
      - `* type: required -> optional`

### HuaweiCloud SDK EG

- _新增特性_
  - 支持以下接口：
    - `ShowDetailOfEventTrace`
    - `ShowDetailOfEvent`
    - `ListTracedEvents`
    - `PutOfficialEvents`
    - `ListEventStreaming`
    - `CreateEventStreaming`
    - `ShowEventStreaming`
    - `UpdateEventStreaming`
    - `DeleteEventStreaming`
    - `ResumeEventStreaming`
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `ListEventSchema`
    - `CreateEventSchema`
    - `ShowDetailOfEventSchema`
    - `UpdateEventSchema`
    - `DeleteEventSchema`
    - `ListEventSchemaVersions`
    - `CreateEventSchemaVersion`
    - `ShowDetailOfEventSchemaVersion`
    - `DeleteEventSchemaVersion`
    - `DiscoverEventSchemaFromData`
  - **ShowDetailOfEventSource**
    - 响应参数变更
      - `+ provider_type: enum value [PARTNER]`
  - **UpdateEventSource**
    - 响应参数变更
      - `+ provider_type: enum value [PARTNER]`
  - **ShowDetailOfChannel**
    - 响应参数变更
      - `+ provider_type: enum value [PARTNER]`
  - **UpdateChannel**
    - 响应参数变更
      - `+ provider_type: enum value [PARTNER]`
  - **UpdateSubscriptionSource**
    - 请求参数变更
      - `+ provider_type: enum value [PARTNER]`
  - **ShowDetailOfConnection**
    - 响应参数变更
      - `+ type`
      - `+ kafka_detail`
  - **UpdateConnection**
    - 响应参数变更
      - `+ type`
      - `+ kafka_detail`
  - **CreateEventSource**
    - 响应参数变更
      - `+ provider_type: enum value [PARTNER]`
  - **ListEventSources**
    - 请求参数变更
      - `+ provider_type: enum value [PARTNER]`
    - 响应参数变更
      - `+ items.provider_type: enum value [PARTNER]`
  - **CreateChannel**
    - 响应参数变更
      - `+ provider_type: enum value [PARTNER]`
  - **ListChannels**
    - 请求参数变更
      - `+ provider_type: enum value [PARTNER]`
    - 响应参数变更
      - `+ items.provider_type: enum value [PARTNER]`
  - **CreateSubscriptionTarget**
    - 请求参数变更
      - `+ kafka_detail`
  - **UpdateSubscriptionTarget**
    - 请求参数变更
      - `+ kafka_detail`
  - **CreateConnection**
    - 请求参数变更
      - `+ type`
      - `+ kafka_detail`
    - 响应参数变更
      - `+ type`
      - `+ kafka_detail`
  - **ListConnections**
    - 响应参数变更
      - `+ type`
      - `+ kafka_detail`
      - `+ items.type`
      - `+ items.kafka_detail`
  - **ListEventTarget**
    - 请求参数变更
      - `+ support_types`
  - **UpdateSubscription**
    - 请求参数变更
      - `+ sources.provider_type: enum value [PARTNER]`
      - `+ targets.kafka_detail`
  - **CreateSubscription**
    - 请求参数变更
      - `+ sources.provider_type: enum value [PARTNER]`
      - `+ targets.kafka_detail`
  - **ListSubscriptions**
    - 请求参数变更
      - `+ connection_id`

### HuaweiCloud SDK ELB

- _新增特性_
  - 支持接口`DeleteLoadBalancerForce`、`DeleteListenerForce`、`BatchUpdateMembers`
- _解决问题_
  - 无
- _特性变更_
  - **ShowQuota**
    - 响应参数变更
      - `+ quota.condition_per_policy`
      - `+ quota.listeners_per_pool`
      - `+ quota.listeners_per_loadbalancer`
      - `* quota.ipgroup_bindings: string -> int32`
      - `* quota.ipgroup_max_length: string -> int32`
  - **ShowLoadBalancer**
    - 响应参数变更
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
  - **UpdateLoadBalancer**
    - 请求参数变更
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
    - 响应参数变更
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
  - **ListListeners**
    - 请求参数变更
      - `+ protection_status`
    - 响应参数变更
      - `+ listeners.protection_status`
      - `+ listeners.protection_reason`
      - `+ listeners.gzip_enable`
  - **CreateListener**
    - 请求参数变更
      - `+ listener.protection_status`
      - `+ listener.protection_reason`
      - `+ listener.gzip_enable`
    - 响应参数变更
      - `+ listener.protection_status`
      - `+ listener.protection_reason`
      - `+ listener.gzip_enable`
  - **ShowListener**
    - 响应参数变更
      - `+ listener.protection_status`
      - `+ listener.protection_reason`
      - `+ listener.gzip_enable`
  - **UpdateListener**
    - 请求参数变更
      - `+ listener.protection_status`
      - `+ listener.protection_reason`
      - `+ listener.gzip_enable`
    - 响应参数变更
      - `+ listener.protection_status`
      - `+ listener.protection_reason`
      - `+ listener.gzip_enable`
  - **ListPools**
    - 请求参数变更
      - `+ protection_status`
    - 响应参数变更
      - `+ pools.protection_status`
      - `+ pools.protection_reason`
  - **CreatePool**
    - 请求参数变更
      - `+ pool.protection_status`
      - `+ pool.protection_reason`
    - 响应参数变更
      - `+ pool.protection_status`
      - `+ pool.protection_reason`
  - **ShowPool**
    - 响应参数变更
      - `+ pool.protection_status`
      - `+ pool.protection_reason`
  - **UpdatePool**
    - 请求参数变更
      - `+ pool.protection_status`
      - `+ pool.protection_reason`
    - 响应参数变更
      - `+ pool.protection_status`
      - `+ pool.protection_reason`
  - **UpdateMember**
    - 请求参数变更
      - `+ member.protocol_port`
  - **ListLoadBalancers**
    - 请求参数变更
      - `+ protection_status`
      - `+ global_eips`
    - 响应参数变更
      - `+ loadbalancers.protection_status`
      - `+ loadbalancers.protection_reason`
  - **CreateLoadBalancer**
    - 请求参数变更
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
    - 响应参数变更
      - `+ loadbalancer.protection_status`
      - `+ loadbalancer.protection_reason`
  - **ListL7Policies**
    - 响应参数变更
      - `+ l7policies.redirect_pools_extend_config`
      - `- l7policies.redirect_pools_config`
  - **CreateL7Policy**
    - 请求参数变更
      - `+ l7policy.redirect_pools_extend_config`
      - `- l7policy.redirect_pools_config`
    - 响应参数变更
      - `+ l7policy.redirect_pools_extend_config`
      - `- l7policy.redirect_pools_config`
  - **ShowL7Policy**
    - 响应参数变更
      - `+ l7policy.redirect_pools_extend_config`
      - `- l7policy.redirect_pools_config`
  - **UpdateL7Policy**
    - 请求参数变更
      - `+ l7policy.redirect_pools_extend_config`
      - `- l7policy.redirect_pools_config`
    - 响应参数变更
      - `+ l7policy.redirect_pools_extend_config`
      - `- l7policy.redirect_pools_config`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateGaussMySqlInstance**
    - 请求参数变更
      - `+ restore_point`

### HuaweiCloud SDK IEC

- _新增特性_
  - 支持以下接口：
    - `ListCloudImages`
    - `RegisterImage`
    - `CreateImage`
    - `ShowVolumeTypes`
    - `RebuildImage`
    - `DeleteImage`
    - `UpdateBandwidth`
    - `DeleteBandwidth`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListInstanceTopics**
    - 请求参数变更
      - `+ offset`
      - `+ limit`
  - **ListInstances**
    - 请求参数变更
      - `+ offset`
      - `+ limit`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListEditingJobs`、`CreateEditingJobs`、`DeleteEditingJobs`

### HuaweiCloud SDK Organizations

- _新增特性_
  - 支持接口`ListQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowIssueV4**
    - 响应参数变更
      - `+ story_point`
  - **SearchIssues**
    - 响应参数变更
      - `+ issue_list.due_date`
  - **ListIssueCommentsV4**
    - 响应参数变更
      - `+ comments.timestamp`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreatePostPaidInstanceByEngine**
    - 请求参数变更
      - `+ bss_param`
  - **CreatePostPaidInstance**
    - 请求参数变更
      - `+ bss_param`
  - **ListInstancesDetails**
    - 请求参数变更
      - `+ offset`
      - `+ limit`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListEngineFlavors`、`BatchDeleteManualBackup`、`DeleteJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 支持以下接口：
    - `SendRocketMqDlqMessage`
    - `ValidateRocketMqConsumedMessage`
    - `ListRocketMqMigrationTask`
    - `CreateRocketMqMigrationTask`
    - `DeleteRocketMqMigrationTask`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SMN

- _新增特性_
  - 支持以下接口：
    - `UpdateSubscription`
    - `ListLogtank`
    - `CreateLogtank`
    - `UpdateLogtank`
    - `DeleteLogtank`
- _解决问题_
  - 无
- _特性变更_
  - **ListTopicDetails**
    - 响应参数变更
      - `+ topic_id`
  - **ListTopics**
    - 请求参数变更
      - `+ topic_id`
    - 响应参数变更
      - `+ topics.topic_id`
  - **ListTopicAttributes**
    - 响应参数变更
      - `+ attributes.access_policy`
      - `+ attributes.introduction`
      - `- attributes.Version`
      - `- attributes.Id`
      - `- attributes.Statement`
  - **AddSubscription**
    - 请求参数变更
      - `+ extension`

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateAssetByFileUpload**
    - 请求参数变更
      - `+ review.interval`
      - `+ review.politics`
      - `+ review.terrorism`
      - `+ review.porn`
  - **PublishAssetFromObs**
    - 请求参数变更
      - `+ review.interval`
      - `+ review.politics`
      - `+ review.terrorism`
      - `+ review.porn`
  - **CreateAssetReviewTask**
    - 请求参数变更
      - `+ review.interval`
      - `+ review.politics`
      - `+ review.terrorism`
      - `+ review.porn`
    - 响应参数变更
      - `+ review.interval`
      - `+ review.politics`
      - `+ review.terrorism`
      - `+ review.porn`
  - **UploadMetaDataByUrl**
    - 请求参数变更
      - `+ upload_metadatas.review.interval`
      - `+ upload_metadatas.review.politics`
      - `+ upload_metadatas.review.terrorism`
      - `+ upload_metadatas.review.porn`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateVpcPeering**
    - 请求参数变更
      - `+ peering.description`

# 0.1.41 2023-05-25

### HuaweiCloud SDK CBH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowNetworkConfiguration**
    - 请求参数变更
      - `- nics.ip_address`
  - **CreateInstanceOrder**
    - 请求参数变更
      - `- product_infos.resource_size_measure_id`
      - `- product_infos.resource_size`
  - **ChangeInstanceNetwork**
    - 请求参数变更
      - `- nics.ip_address`
  - **CreateInstance**
    - 请求参数变更
      - `+ server.enterprise_project_id`
      - `- server.nics.ip_address`
      - `- server.public_ip.eip`

### HuaweiCloud SDK CBR

- _新增特性_
  - 支持以下接口：
    - `ImportCheckpoint`
    - `ListExternalVault`
    - `BatchUpdateVault`
    - `SetVaultResource`
    - `ShowMetadata`
    - `CheckAgent`
    - `ListProjects`
    - `ListDomainProjects`
    - `ShowDomain`
    - `ShowMigrateStatus`
    - `MigrateDomain`
    - `ShowStorageUsage`
    - `UpdateOrder`
    - `CreatePostPaidVault`
    - `UpdateBackup`
- _解决问题_
  - 无
- _特性变更_
  - **CreateVault**
    - 请求参数变更
      - `+ vault.threshold`
      - `+ vault.smn_notify`
      - `+ vault.backup_name_prefix`
      - `+ vault.demand_billing`
    - 响应参数变更
      - `+ vault.backup_name_prefix`
      - `+ vault.demand_billing`
      - `+ vault.cbc_delete_count`
      - `+ vault.frozen`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowCluster**
    - 响应参数变更
      - `+ metadata.alias`
  - **UpdateCluster**
    - 请求参数变更
      - `+ metadata`
    - 响应参数变更
      - `+ metadata.alias`
  - **DeleteCluster**
    - 响应参数变更
      - `+ metadata.alias`
  - **MigrateNode**
    - 请求参数变更
      - `+ spec.runtime`
    - 响应参数变更
      - `+ spec.runtime`
  - **CreateCluster**
    - 请求参数变更
      - `+ metadata.alias`
    - 响应参数变更
      - `+ metadata.alias`
  - **ListClusters**
    - 响应参数变更
      - `+ items.metadata.alias`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDomainDetailByName**
    - 响应参数变更
      - `+ domain.domain_name`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListMigrationTask**
    - 响应参数变更
      - `- task_name`
      - `- target_instance_id`
      - `- target_instance_address`
      - `- target_instance_name`
      - `- migrate_type`
      - `- created_at`
      - `- source_instance_id`
      - `- task_id`
      - `- data_source`
      - `- migration_method`
      - `- source_instance_name`
      - `- status`
  - **ListConfigTemplates**
    - 响应参数变更
      - `* template_num: number -> integer`
  - **ListInstances**
    - 响应参数变更
      - `+ instances.updated_at`
  - **ListBackgroundTask**
    - 响应参数变更
      - `- updated_at`
      - `- created_at`
      - `- status`
  - **ListFlavors**
    - 响应参数变更
      - `+ flavors.flavors_available_zones.unit`
      - `+ flavors.flavors_available_zones.available_zones`

### HuaweiCloud SDK ECS

- _新增特性_
  - 支持接口`ListFlavorSellPolicies`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPublicipsByTags**
    - 响应参数变更
      - `+ resources.resource_detail`
      - `- resources.resouce_detail`
  - **AddPublicipsIntoSharedBandwidth**
    - 响应参数变更
      - `+ bandwidth.enable_bandwidth_rules`
      - `+ bandwidth.rule_quota`
      - `+ bandwidth.bandwidth_rules`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`UpdateProxyConnectionPoolType`、`RestoreOldInstance`、`ShowBackupRestoreTime`
- _解决问题_
  - 无
- _特性变更_
  - **ShowGaussMySqlProxyList**
    - 响应参数变更
      - `+ proxy_list.proxy.connection_pool_type`
      - `+ proxy_list.proxy.switch_connection_pool_type_enabled`

### HuaweiCloud SDK IAM

- _新增特性_
  - 支持接口`AssociateRoleToAgencyOnEnterpriseProject`、`RevokeRoleFromAgencyOnEnterpriseProject`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Image

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CreateVideoObjectMaskingTask`、`ShowVideoObjectMaskingTask`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`DeleteConnector`、`CreateDeleteConnectorOrder`、`CreateKafkaConsumerGroup`、`CloseKafkaManager`
- _解决问题_
  - 无
- _特性变更_
  - **ShowInstance**
    - 响应参数变更
      - `+ kafka_manager_enable`
  - **ListInstances**
    - 响应参数变更
      - `+ kafka_manager_enable`
      - `+ instances.kafka_manager_enable`

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`BatchShowIpBelongs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MSGSMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowSignatureFile**
    - 响应参数变更
      - `+ file_desc`
  - **UpdateApp**
    - 响应参数变更
      - `- app_secret`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreatePostPaidInstanceByEngine**
    - 请求参数变更
      - `+ engine_version: enum value [3.8.35]`
  - **CreatePostPaidInstance**
    - 请求参数变更
      - `+ engine_version: enum value [3.8.35]`

### HuaweiCloud SDK RAM

- _新增特性_
  - 支持接口`ListQuota`
- _解决问题_
  - 无
- _特性变更_
  - **AssociateResourceShare**
    - 响应参数变更
      - `+ resource_share_associations.external`
  - **DisassociateResourceShare**
    - 响应参数变更
      - `+ resource_share_associations.external`
  - **SearchResourceShareAssociations**
    - 响应参数变更
      - `+ resource_share_associations.external`
  - **CreateResourceShare**
    - 请求参数变更
      - `+ allow_external_principals`
    - 响应参数变更
      - `+ resource_share.allow_external_principals`
  - **SearchResourceShares**
    - 响应参数变更
      - `+ resource_shares.allow_external_principals`
  - **UpdateResourceShare**
    - 请求参数变更
      - `+ allow_external_principals`
    - 响应参数变更
      - `+ resource_share.allow_external_principals`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSqlserverDatabases**
    - 请求参数变更
      - `+ recover_model`

### HuaweiCloud SDK RMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowPolicyAssignment**
    - 响应参数变更
      - `+ created_by`
  - **UpdatePolicyAssignment**
    - 响应参数变更
      - `+ created_by`
  - **ShowAggregatePolicyAssignmentDetail**
    - 响应参数变更
      - `+ created_by`
  - **CreatePolicyAssignments**
    - 响应参数变更
      - `+ created_by`
  - **ListPolicyAssignments**
    - 响应参数变更
      - `+ created_by`
      - `+ value.created_by`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateVpc**
    - 请求参数变更
      - `+ vpc.tags`
  - **CreateSubnet**
    - 请求参数变更
      - `+ subnet.tags`
    - **ShowAddressGroup**
    - 响应参数变更
      - `+ address_group.max_capacity`
      - `+ address_group.status`
      - `+ address_group.status_message`
  - **UpdateAddressGroup**
    - 请求参数变更
      - `+ address_group.max_capacity`
    - 响应参数变更
      - `+ address_group.max_capacity`
      - `+ address_group.status`
      - `+ address_group.status_message`
  - **ListAddressGroup**
    - 响应参数变更
      - `+ address_groups.max_capacity`
      - `+ address_groups.status`
      - `+ address_groups.status_message`
  - **CreateAddressGroup**
    - 请求参数变更
      - `+ address_group.max_capacity`
    - 响应参数变更
      - `+ address_group.max_capacity`
      - `+ address_group.status`
      - `+ address_group.status_message`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListServiceDescribeDetails**
    - 响应参数变更
      - `+ enable_policy`
  - **ListServiceDetails**
    - 响应参数变更
      - `- vip_port_id`
  - **UpdateEndpointService**
    - 请求参数变更
      - `- vip_port_id`
    - 响应参数变更
      - `- vip_port_id`
  - **ListServicePublicDetails**
    - 响应参数变更
      - `+ endpoint_services.enable_policy`
  - **CreateEndpointService**
    - 请求参数变更
      - `- vip_port_id`
    - 响应参数变更
      - `- vip_port_id`
  - **ListEndpointService**
    - 响应参数变更
      - `- endpoint_services.vip_port_id`

# 0.1.40 2023-05-18

### HuaweiCloud SDK CBR

- _新增特性_
  - 支持以下接口：
    - `AddAgentPath`
    - `ShowAgent`
    - `UpdateAgent`
    - `UnregisterAgent`
    - `ListAgent`
    - `RegisterAgent`
    - `RemoveAgentPath`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowNode**
    - 响应参数变更
      - `+ status.lastProbeTime`
  - **UpdateNode**
    - 响应参数变更
      - `+ status.lastProbeTime`
  - **DeleteNode**
    - 响应参数变更
      - `+ status.lastProbeTime`
  - **CreateNode**
    - 响应参数变更
      - `+ status.lastProbeTime`
  - **ListNodes**
    - 响应参数变更
      - `+ items.status.lastProbeTime`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateRefreshTasks**
    - 请求参数变更
      - `* refresh_task.mode: boolean -> string`
  - **ShowDomainFullConfig**
    - 响应参数变更
      - `+ configs.flexible_origin`
      - `+ configs.slice_etag_status`
      - `+ configs.origin_receive_timeout`
      - `+ configs.remote_auth`
      - `+ configs.websocket`
      - `+ configs.video_seek`
      - `+ configs.request_limit_rules`
      - `+ configs.url_auth.sign_method`
      - `+ configs.url_auth.match_type`
      - `+ configs.url_auth.key`
      - `+ configs.url_auth.backup_key`
      - `+ configs.url_auth.sign_arg`
      - `+ configs.https.expire_time`
      - `+ configs.https.certificate_type`
      - `+ configs.https.ocsp_stapling_status`
      - `+ configs.sources.obs_bucket_type`
      - `+ configs.compress.file_type`
      - `+ configs.user_agent_filter.ua_list`
  - **UpdateDomainFullConfig**
    - 请求参数变更
      - `+ configs.flexible_origin`
      - `+ configs.slice_etag_status`
      - `+ configs.origin_receive_timeout`
      - `+ configs.remote_auth`
      - `+ configs.websocket`
      - `+ configs.video_seek`
      - `+ configs.request_limit_rules`
      - `+ configs.url_auth.sign_method`
      - `+ configs.url_auth.match_type`
      - `+ configs.url_auth.backup_key`
      - `+ configs.url_auth.sign_arg`
      - `+ configs.https.certificate_type`
      - `+ configs.https.ocsp_stapling_status`
      - `+ configs.sources.obs_bucket_type`
      - `+ configs.compress.file_type`
      - `+ configs.user_agent_filter.ua_list`
  - **ShowDomainDetailByName**
    - 响应参数变更
      - `+ domain.sources.obs_bucket_type`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowInstanceStatus**
    - 请求参数变更
      - `- X-Language`
  - **StopPipelineNew**
    - 请求参数变更
      - `- X-Language`
  - **RemovePipeline**
    - 请求参数变更
      - `- X-Language`
  - **RunPipeline**
    - 请求参数变更
      - `- X-Language`
  - **BatchShowPipelinesLatestStatus**
    - 请求参数变更
      - `- X-Language`
  - **ListPipelines**
    - 请求参数变更
      - `- X-Language`
  - **DeletePipeline**
    - 请求参数变更
      - `- X-Language`
  - **StopPipelineRun**
    - 请求参数变更
      - `- X-Language`
  - **ListPipelineRuns**
    - 请求参数变更
      - `- X-Language`
  - **StartNewPipeline**
    - 请求参数变更
      - `- X-Language`
  - **BatchShowPipelinesStatus**
    - 请求参数变更
      - `- X-Language`
  - **ListPipelineSimpleInfo**
    - 请求参数变更
      - `- X-Language`
  - **ShowPipleineStatus**
    - 请求参数变更
      - `- X-Language`
  - **ListPipleineBuildResult**
    - 请求参数变更
      - `- X-Language`
  - **ListPipelineTemplates**
    - 请求参数变更
      - `- X-Language`
  - **CreatePipelineByTemplateId**
    - 请求参数变更
      - `- X-Language`
  - **ShowTemplateDetail**
    - 请求参数变更
      - `- X-Language`
  - **CreatePipelineByTemplate**
    - 请求参数变更
      - `- X-Language`
  - **ShowPipelineRunDetail**
    - 请求参数变更
      - `- X-Language`
  - **ListTemplates**
    - 请求参数变更
      - `- X-Language`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateTaskStatus**
    - 请求参数变更
      - `+ enterprise_project_id`
  - **UpdateAgentHealthStatus**
    - 响应参数变更
      - `+ result.result.kafka_enable`
      - `+ result.result.kafka_shadow_topic_prefix`
      - `+ result.result.app_log_level`
      - `+ result.result.app_log_path`
      - `+ result.result.mock_rule_list`

### HuaweiCloud SDK CSE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpgradeEngine**
    - 响应参数变更
      - `+ jobId`
      - `- job_id`
  - **RetryEngine**
    - 响应参数变更
      - `+ jobId`
      - `- job_id`
  - **DownloadKie**
    - 响应参数变更
      - `+ data.id`
  - **CreateEngine**
    - 响应参数变更
      - `+ jobId`
      - `- job_id`
  - **DeleteEngine**
    - 响应参数变更
      - `+ jobId`
      - `- job_id`

### HuaweiCloud SDK DAS

- _新增特性_
  - 支持接口`CreateShareConnections`、`CancelShareConnections`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持接口`ListAvailableDisasterClusters`、`CheckDisasterName`、`ShowDisasterDetail`、`UpdateDisasterInfo`
- _解决问题_
  - 无
- _特性变更_
  - **CreateCluster**
    - 请求参数变更
      - `+ cluster.tags`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateServers**
    - 请求参数变更
      - `+ server.root_volume.metadata`
      - `- server.root_volume.extendparam.__system__encrypted`
      - `- server.root_volume.extendparam.__system__cmkid`
      - `+ server.data_volumes.delete_on_termination`
  - **CreatePostPaidServers**
    - 请求参数变更
      - `+ server.data_volumes.delete_on_termination`
      - `+ server.root_volume.metadata`
      - `- server.root_volume.extendparam.__system__encrypted`
      - `- server.root_volume.extendparam.__system__cmkid`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`UpdateGaussMySqlDatabaseUserComment`、`UpdateGaussMySqlDatabaseComment`、`ListLtsSlowlogDetails`、`ListLtsErrorLogDetails`
- _解决问题_
  - 无
- _特性变更_
  - **ListGaussMySqlDatabaseUser**
    - 响应参数变更
      - `+ users.comment`
  - **CreateGaussMySqlDatabaseUser**
    - 请求参数变更
      - `+ users.comment`
  - **ListGaussMySqlDatabase**
    - 响应参数变更
      - `+ databases.comment`
  - **CreateGaussMySqlDatabase**
    - 请求参数变更
      - `+ databases.comment`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **BroadcastMessage**
    - 请求参数变更
      - `+ ttl`
      - `+ message_id`
  - **ShowDeviceGroup**
    - 响应参数变更
      - `+ dynamic_group_rule`
      - `+ group_type`
  - **UpdateDeviceGroup**
    - 响应参数变更
      - `+ dynamic_group_rule`
      - `+ group_type`
  - **SearchDevices**
    - 响应参数变更
      - `+ devices.groups`
  - **AddDeviceGroup**
    - 请求参数变更
      - `+ group_type`
      - `+ dynamic_group_rule`
    - 响应参数变更
      - `+ dynamic_group_rule`
      - `+ group_type`
  - **ListDeviceGroups**
    - 请求参数变更
      - `+ group_type`
      - `+ name`
    - 响应参数变更
      - `+ device_groups.group_type`
      - `* device_groups: list<DeviceGroupResponsSummery> -> list<DeviceGroupResponseSummary>`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`ListTemplates`、`SearchIssues`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.39 2023-05-11

### HuaweiCloud SDK CMS

- _新增特性_
  - 支持容量管理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`ContinueDeployStack`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowVaultResourceInstances**
    - 响应参数变更
      - `* resources.resource_detail: list<Vault> -> object<InstancesResourceDetail>`
  - **ListPolicies**
    - 响应参数变更
      - `+ policies.operation_definition.full_backup_interval`
  - **CreatePolicy**
    - 请求参数变更
      - `+ policy.operation_definition.full_backup_interval`
    - 响应参数变更
      - `+ policy.operation_definition.full_backup_interval`
  - **ShowPolicy**
    - 响应参数变更
      - `+ policy.operation_definition.full_backup_interval`
  - **UpdatePolicy**
    - 请求参数变更
      - `+ policy.operation_definition.full_backup_interval`
    - 响应参数变更
      - `+ policy.operation_definition.full_backup_interval`
  - **CreateVault**
    - 请求参数变更
      - `- vault.billing.extra_info`

### HuaweiCloud SDK CBS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ExecuteGetFramsListByImagesId**
    - 请求参数变更
      - `+ offset`
      - `+ limit`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPipelineTemplates**
    - 响应参数变更
      - `+ total`
      - `+ offset`
      - `+ templates`
      - `+ limit`
      - `- is_system`
      - `- is_show_source`
      - `- create_time`
      - `- icon`
      - `- description`
      - `- language`
      - `- domain_id`
      - `- is_collect`
      - `- update_time`
      - `- name`
      - `- manifest_version`
      - `- creator_id`
      - `- updater_id`
      - `- stages`
      - `- creator_name`
      - `- id`
      - `- region`

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ExportConnections**
    - 请求参数变更
      - `+ workspace`
  - **ExportJob**
    - 请求参数变更
      - `+ workspace`
  - **StopJob**
    - 请求参数变更
      - `+ workspace`
  - **StopJobInstance**
    - 请求参数变更
      - `+ workspace`
  - **RestoreJobInstance**
    - 请求参数变更
      - `+ workspace`
  - **CancelScript**
    - 请求参数变更
      - `+ workspace`
  - **DeleteConnction**
    - 请求参数变更
      - `+ workspace`
  - **ShowConnection**
    - 请求参数变更
      - `+ workspace`
  - **UpdateConnection**
    - 请求参数变更
      - `+ workspace`
  - **ExportJobList**
    - 请求参数变更
      - `+ workspace`
  - **ImportJob**
    - 请求参数变更
      - `+ workspace`
  - **DeleteScript**
    - 请求参数变更
      - `+ workspace`
  - **ShowScript**
    - 请求参数变更
      - `+ workspace`
  - **UpdateScript**
    - 请求参数变更
      - `+ workspace`
  - **ExecuteScript**
    - 请求参数变更
      - `+ workspace`
  - **CreateResource**
    - 请求参数变更
      - `+ workspace`
  - **DeleteResource**
    - 请求参数变更
      - `+ workspace`
  - **ShowResource**
    - 请求参数变更
      - `+ workspace`
  - **UpdateResource**
    - 请求参数变更
      - `+ workspace`
  - **CreateConnection**
    - 请求参数变更
      - `+ workspace`
  - **ListConnections**
    - 请求参数变更
      - `+ workspace`
  - **ImportConnections**
    - 请求参数变更
      - `+ workspace`
  - **ShowFileInfo**
    - 请求参数变更
      - `+ workspace`
  - **StartJob**
    - 请求参数变更
      - `+ workspace`
  - **RunOnce**
    - 请求参数变更
      - `+ workspace`
  - **ShowJobStatus**
    - 请求参数变更
      - `+ workspace`
  - **ListJobInstances**
    - 请求参数变更
      - `+ workspace`
  - **ShowJobInstance**
    - 请求参数变更
      - `+ workspace`
  - **ListSystemTasks**
    - 请求参数变更
      - `+ workspace`
  - **CreateScript**
    - 请求参数变更
      - `+ workspace`
  - **ListScripts**
    - 请求参数变更
      - `+ workspace`
  - **ListScriptResults**
    - 请求参数变更
      - `+ workspace`
  - **DeleteJob**
    - 请求参数变更
      - `+ workspace`
  - **ShowJob**
    - 请求参数变更
      - `+ workspace`
  - **UpdateJob**
    - 请求参数变更
      - `+ workspace`
  - **CreateJob**
    - 请求参数变更
      - `+ workspace`
  - **ListJobs**
    - 请求参数变更
      - `+ workspace`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowMonitoringData**
    - 响应参数变更
      - `+ results.data_guard_minitor.migration_bytes_per_second`
  - **BatchListJobDetails**
    - 响应参数变更
      - `+ results.data_transformation`
      - `+ results.tags`

### HuaweiCloud SDK ECS

- _新增特性_
  - 支持接口`NovaAttachInterface`
- _解决问题_
  - 无
- _特性变更_
  - **ReinstallServerWithoutCloudInit**
    - 请求参数变更
      - `+ os-reinstall.metadata`
  - **ChangeServerOsWithoutCloudInit**
    - 请求参数变更
      - `+ os-change.metadata`
  - **ReinstallServerWithCloudInit**
    - 请求参数变更
      - `+ os-reinstall.metadata.__system__encrypted`
      - `+ os-reinstall.metadata.__system__cmkid`
  - **ChangeServerOsWithCloudInit**
    - 请求参数变更
      - `+ os-change.metadata.__system__encrypted`
      - `+ os-change.metadata.__system__cmkid`
  - **CreateServers**
    - 请求参数变更
      - `+ server.root_volume.extendparam.__system__encrypted`
      - `+ server.root_volume.extendparam.__system__cmkid`
  - **CreatePostPaidServers**
    - 请求参数变更
      - `+ server.root_volume.extendparam.__system__encrypted`
      - `+ server.root_volume.extendparam.__system__cmkid`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持接口`CreateSynthesisTask`、`ShowSynthesisTaskResult`、`CreateCustomPropsTask`、`ShowCustomPropsTaskResult`
- _解决问题_
  - 无
- _特性变更_
  - **CreateCpiTask**
    - 请求参数变更
      - `+ custom_props`
  - **CreateGenerationTask**
    - 请求参数变更
      - `+ custom_props`
  - **CreateOptimizationTask**
    - 请求参数变更
      - `+ custom_props`
  - **ShowCpiTaskResult**
    - 响应参数变更
      - `+ task_data.custom_props`
      - `+ result.custom_props`
  - **ShowGenerationTaskResult**
    - 响应参数变更
      - `+ task_data.custom_props`
      - `+ result.initial_dataset_size`
      - `+ result.strong_constraints`
      - `+ result.weak_constraints`
      - `+ result.binding_site`
      - `+ result.custom_props`
  - **ShowOptimizationTaskResult**
    - 响应参数变更
      - `+ task_data.custom_props`
      - `+ result.strong_constraints`
      - `+ result.weak_constraints`
      - `+ result.binding_site`
      - `+ result.custom_props`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListGaussMySqlInstances**
    - 请求参数变更
      - `+ readonly_private_ip`
      - `+ proxy_ip`
    - 响应参数变更
      - `+ instances.readonly_private_ips`
      - `+ instances.proxy_ips`

### HuaweiCloud SDK Image

- _新增特性_
  - 支持接口`CreateVideoTaggingMediaTask`、`ShowVideoTaggingMediaTask`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListImages**
    - 请求参数变更
      - `+ __imagetype: enum value [market]`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateRecordIndex**
    - 响应参数变更
      - `+ width`
      - `- weight`
  - **ListAreaDetail**
    - 请求参数变更
      - `* play_domains: required -> optional`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **UpdateAomMappingRules**
    - 请求参数变更
      - `* body: object<AomMappingRequestInfo> -> object<UpdateAomMappingRequest>`

### HuaweiCloud SDK MPC

- _新增特性_
  - 支持接口`ListEditingJobs`、`CreateEditingJobs`、`DeleteEditingJobs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OSM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListOrderIncident**
    - 请求参数变更
      - `+ xCustomerId`
    - 响应参数变更
      - `* incidentInfoList.incidentTypeName: object -> string`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`UpdatePostgresqlDbUserComment`、`UpdatePostgresqlDatabase`、`DeletePostgresqlDbUser`、`DeletePostgresqlDatabase`
- _解决问题_
  - 无
- _特性变更_
  - **CreatePostgresqlDatabase**
    - 请求参数变更
      - `+ comment`
  - **CreatePostgresqlDbUser**
    - 请求参数变更
      - `+ comment`
  - **ListPostgresqlDatabases**
    - 响应参数变更
      - `+ databases.comment`
  - **ListPostgresqlDbUserPaginated**
    - 响应参数变更
      - `+ users.comment`

### HuaweiCloud SDK RMS

- _新增特性_
  - 支持接口`ShowAggregatePolicyStateComplianceSummary`、`ListAggregateComplianceByPolicyAssignment`、`ShowAggregateComplianceDetailsByPolicyAssignment`、`ShowAggregatePolicyAssignmentDetail`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 支持以下接口：
    - `ShowFsDirQuota`
    - `UpdateFsDirQuota`
    - `CreateFsDirQuota`
    - `DeleteFsDirQuota`
    - `ShowFsDir`
    - `CreateFsDir`
    - `DeleteFsDir`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK UGO

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RunSqlConversion**
    - 请求参数变更
      - `+ target_db_type: enum value [GaussDB Primary/Standby]`
  - **ConfirmTargetDbType**
    - 请求参数变更
      - `- target_db_type: enum value [RDS for MySQL,GaussDB(for MySQL),RDS for PostgreSQL]`
      - `- target_db_version: enum value [5.7,8.0,11,Enhanced Edition]`

# 0.1.38 2023-04-27

### HuaweiCloud SDK MSGSMS

- _新增特性_
  - 支持消息&短信服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListResourceUnderNode**
    - 请求参数变更
      - `+ marker`
      - `- maker`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDomainFullConfig**
    - 响应参数变更
      - `+ configs.ipv6_accelerate`
      - `+ configs.origin_range_status`

### HuaweiCloud SDK CFW

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListIpsProtectModeUsingPost**
    - 响应参数变更
      - `+ data`
      - `- object_id`
      - `- status`

### HuaweiCloud SDK CSMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListResourceInstances**
    - 响应参数变更
      - `+ resources.sys_tags`

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ResetPassword`、`UpdateInstanceBandwidth`、`ListConfigTemplates`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持以下接口：
    - `ListIamUsers`
    - `ListIamGroups`
    - `ListIamGroupUsers`
    - `ListScaleOutPolicy`
    - `CreateScaleOutPolicy`
    - `ShowScaleOutPolicy`
    - `UpdateScaleOutPolicy`
    - `DeleteScaleOutPolicy`
    - `StartScaleOutPolicy`
    - `StopScaleOutPolicy`
    - `ShowScaleInPolicy`
    - `UpdateScaleInPolicy`
    - `InstallNextflow`
    - `ShowNextflow`
    - `UninstallNextflow`
    - `CleanNextflowCache`
    - `ListNextflowVersion`
    - `ListNextflowJob`
    - `CreateNextflowJob`
    - `ShowNextflowJob`
    - `DeleteNextflowJob`
    - `RetryNextflowJob`
    - `ShowNextflowJobLog`
    - `StopNextflowJob`
    - `ShowNextflowJobReports`
    - `ListNextflowTask`
    - `ShowNextflowTaskDetail`
    - `ShowNextflowTaskLog`
    - `ListNextflowWorkflow`
    - `CreateNextflowWorkflow`
    - `ShowNextflowWorkflow`
    - `UpdateNextflowWorkflow`
    - `DeleteNextflowWorkflow`
- _解决问题_
  - 无
- _特性变更_
  - **PublishData**
    - 响应参数变更
      - `+ id`
      - `- asset_id`
  - **PublishImage**
    - 响应参数变更
      - `+ id`
      - `- asset_id`
  - **PublishApp**
    - 响应参数变更
      - `+ id`
      - `- asset_id`
  - **PublishWorkflow**
    - 响应参数变更
      - `+ id`
      - `- asset_id`
  - **ListImageTag**
    - 响应参数变更
      - `+ tags.path`
  - **CreateApp**
    - 请求参数变更
      - `* body: object<AppDto> -> object<AppReq>`
  - **UpdateApp**
    - 请求参数变更
      - `* body: object<AppDto> -> object<AppReq>`
  - **ListData**
    - 响应参数变更
      - `- path`
      - `- allowed_operate`
      - `- size`
      - `- create_time`
      - `- name`
      - `- download_url`
      - `- deletable`
      - `- type`
      - `- content`
      - `- datas.content`
      - `- datas.download_url`
      - `* datas: list<DataRsp> -> list<DataSummaryRsp>`
  - **ListObsBucketObject**
    - 请求参数变更
      - `+ search_key`

### HuaweiCloud SDK Image

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`RunQueryCustomTags`、`RunDeleteCustomTags`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateInstanceByEngine**
    - 请求参数变更
      - `+ kafka_security_protocol`
      - `+ sasl_enabled_mechanisms`
  - **ShowInstance**
    - 响应参数变更
      - `+ kafka_security_protocol`
      - `+ sasl_enabled_mechanisms: enum value [PLAIN,SCRAM-SHA-512]`
  - **CreatePostPaidInstance**
    - 请求参数变更
      - `+ kafka_security_protocol`
      - `+ sasl_enabled_mechanisms: enum value [PLAIN,SCRAM-SHA-512]`
  - **ListInstances**
    - 响应参数变更
      - `+ kafka_security_protocol`
      - `+ instances.kafka_security_protocol`
      - `+ instances.sasl_enabled_mechanisms: enum value [PLAIN,SCRAM-SHA-512]`

### HuaweiCloud SDK KMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListKeys**
    - 响应参数变更
      - `+ key_details.partition_type`
  - **ListKeyDetail**
    - 响应参数变更
      - `+ key_info.partition_type`
  - **ListRetirableGrants**
    - 响应参数变更
      - `+ total`
  - **ListKmsByTags**
    - 响应参数变更
      - `+ resources.resource_detail.partition_type`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListLogHistogram**
    - 请求参数变更
      - `+ is_iterative`
    - 响应参数变更
      - `+ isQueryComplete`
  - **Createfavorite**
    - 响应参数变更
      - `* create_time: string -> int64`
  - **ListLogStream**
    - 响应参数变更
      - `+ log_streams.is_favorite`
  - **ListLogs**
    - 请求参数变更
      - `+ is_iterative`
    - 响应参数变更
      - `+ isQueryComplete`
  - **UpdateStructTemplate**
    - 请求参数变更
      - `* rule: list<rule> -> object<rule>`
  - **CreateStructTemplate**
    - 请求参数变更
      - `* rule: list<rule> -> object<rule>`
  - **ListHistorySql**
    - 请求参数变更
      - `+ log_group_id`
      - `+ log_stream_id`

### HuaweiCloud SDK Organizations

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreatePolicy**
    - 请求参数变更
      - `+ type: enum value [tag_policy]`
  - **EnablePolicyType**
    - 请求参数变更
      - `+ policy_type: enum value [tag_policy]`
  - **DisablePolicyType**
    - 请求参数变更
      - `+ policy_type: enum value [tag_policy]`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowConsumerListOrDetails**
    - 响应参数变更
      - `* total: int64 -> int32`

### HuaweiCloud SDK SMS

- _新增特性_
  - 支持以下接口：
    - `ListApiVersion`
    - `ShowApiVersion`
    - `ShowConfig`
    - `UpdateNetworkCheckInfo`
    - `ShowConfigSetting`
    - `UploadSpecialConfigurationSetting`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.37 2023-04-20

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAllVersionByVersionId**
    - 响应参数变更
      - `+ elements.job_reference_number`
      - `+ elements.job_reference_name`
      - `- elements.reference_number`
      - `- elements.script_reference`

### HuaweiCloud SDK AOS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **GetStackTemplate**
    - 请求参数变更
      - `- executor`
  - **ListStackEvents**
    - 请求参数变更
      - `- executor`
  - **ListStackOutputs**
    - 请求参数变更
      - `- executor`
  - **DeleteStack**
    - 请求参数变更
      - `- executor`
  - **DeleteExecutionPlan**
    - 请求参数变更
      - `- executor`
  - **ApplyExecutionPlan**
    - 请求参数变更
      - `- executor`
  - **GetExecutionPlan**
    - 请求参数变更
      - `- executor`
  - **ListStackResources**
    - 请求参数变更
      - `- executor`
  - **ListExecutionPlans**
    - 请求参数变更
      - `- executor`
  - **CreateExecutionPlan**
    - 请求参数变更
      - `- executor`
  - **GetExecutionPlanMetadata**
    - 请求参数变更
      - `- executor`
  - **GetStackMetadata**
    - 请求参数变更
      - `- executor`
  - **ListStacks**
    - 请求参数变更
      - `- executor`
  - **CreateStack**
    - 请求参数变更
      - `- executor`
  - **DeployStack**
    - 请求参数变更
      - `- executor`

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListCustomerselfResourceRecordDetails**
    - 响应参数变更
      - `+ monthly_records.pre_order_id`
      - `+ monthly_records.az_code_infos`

### HuaweiCloud SDK CBS

- _新增特性_
  - 支持以下接口：
    - `ExecuteGetVideosList`
    - `ExecuteCreateVideo`
    - `ExecuteUpdateVideoById`
    - `ExecuteDeleteVideoById`
    - `ExecuteGetVideoInfoById`
    - `ExecuteUpdateVideoInfoById`
    - `ExecuteComposeVideo`
    - `ExecuteComposeVideoOndemand`
    - `ExecuteGetCharacters`
    - `ExecuteGetCharacterInfoById`
    - `ExecuteUploadPpt`
    - `ExecuteUploadImage`
    - `ExecuteUpdateImageName`
    - `ExecuteDeleteimageById`
    - `ExecuteGetImagesList`
    - `ExecutePostCreateImages`
    - `ExecuteGetFramsListByImagesId`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ShowDomainFullConfig`、`UpdateDomainFullConfig`
- _解决问题_
  - 无
- _特性变更_
  - **ShowDomainFullConfig**
    - 响应参数变更
      - `+ configs.origin_follow302_status`
      - `+ configs.cache_rules`
      - `+ configs.ip_filter`
      - `+ configs.referer`
      - `+ configs.force_redirect.redirect_code`
  - **UpdateDomainFullConfig**
    - 请求参数变更
      - `+ configs.origin_follow302_status`
      - `+ configs.cache_rules`
      - `+ configs.ip_filter`
      - `+ configs.referer`
      - `+ configs.force_redirect.redirect_code`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
  - 支持接口`ListPipelineTemplates`、`ShowPipelineTemplateDetail`、`CreatePipelineByTemplateId`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListLogsJob**
    - 请求参数变更
      - `+ start`
      - `+ limit`
  - **ShowVpcepConnection**
    - 请求参数变更
      - `+ start`
      - `+ limit`
  - **ListYmlsJob**
    - 请求参数变更
      - `+ start`
      - `+ limit`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateRedislogDownloadLink**
    - 响应参数变更
      - `+ backup_id`

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateConfiguration**
    - 响应参数变更
      - `+ configuration`
      - `- datastore_version`
      - `- created`
      - `- name`
      - `- description`
      - `- id`
      - `- datastore_name`
      - `- updated`

### HuaweiCloud SDK DLF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowJob**
    - 响应参数变更
      - `+ lastUpdateUser`
      - `+ id`
  - **UpdateJob**
    - 请求参数变更
      - `+ lastUpdateUser`
      - `+ id`
  - **CreateJob**
    - 请求参数变更
      - `+ lastUpdateUser`
      - `+ id`
  - **ListJobs**
    - 响应参数变更
      - `+ lastUpdateUser`
      - `+ id`
      - `+ jobs.id`
      - `+ jobs.lastUpdateUser`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowJobList**
    - 响应参数变更
      - `+ jobs.job_action`
      - `+ jobs.children.job_action`
  - **BatchListJobDetails**
    - 响应参数变更
      - `+ results.original_job_direction`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ImportFunction**
    - 响应参数变更
      - `+ gpu_memory`
      - `+ func_vpc.security_groups`
  - **ListFunctions**
    - 响应参数变更
      - `+ functions.gpu_memory`
      - `+ functions.is_bridge_function`
      - `+ functions.bind_bridge_funcUrns`
  - **CreateFunction**
    - 请求参数变更
      - `+ gpu_memory`
      - `+ log_config`
      - `+ network_controller`
      - `+ func_vpc.security_groups`
    - 响应参数变更
      - `+ gpu_memory`
      - `+ func_vpc.security_groups`
  - **ShowFunctionConfig**
    - 响应参数变更
      - `+ gpu_memory`
      - `+ ephemeral_storage`
      - `+ func_vpc.security_groups`
  - **UpdateFunctionConfig**
    - 请求参数变更
      - `+ gpu_memory`
      - `+ ephemeral_storage`
      - `+ log_config`
      - `+ network_controller`
      - `+ restore_hook_handler`
      - `+ restore_hook_timeout`
      - `+ func_vpc.security_groups`
    - 响应参数变更
      - `+ gpu_memory`
      - `+ ephemeral_storage`
      - `+ func_vpc.security_groups`
  - **UpdateFunctionMaxInstanceConfig**
    - 响应参数变更
      - `+ func_vpc.security_groups`
  - **CreateFunctionVersion**
    - 响应参数变更
      - `+ func_vpc.security_groups`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateDbUser**
    - 请求参数变更
      - `+ is_login_only`

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListSimCards**
    - 请求参数变更
      - `+ price_plan_id`

### HuaweiCloud SDK Image

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `CreateVideoTranslateTask`
    - `ShowVideoTranslateTask`
    - `CreateImageTranslateTask`
    - `ShowImageTranslateTask`
    - `CreateVideoCoverAnalysisTask`
    - `ShowVideoCoverAnalysisTask`
    - `CreateVideoSummarizationAnalysisTask`
    - `ShowVideoSummarizationAnalysisTask`
    - `CreateVideoShotSplitTask`
    - `ShowVideoShotSplitTask`
  - **CreateImageHighresolutionMattingTask**
    - 请求参数变更
      - `- input.data.bucket`
      - `- input.data.path`
  - **ShowImageHighresolutionMattingTask**
    - 响应参数变更
      - `- input.data.bucket`
      - `- input.data.path`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 支持接口`BroadcastMessage`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 支持以下接口：
    - `BatchListOtTemplates`
    - `AddOtTemplates`
    - `ShowOtTemplate`
    - `DeleteOtTemplate`
    - `AddGeneralOtTemplate`
    - `UpdateModuleState`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Meeting

- _新增特性_
  - 支持接口`BatchShowUserDetails`
- _解决问题_
  - 无
- _特性变更_
  - **UpdateWebinar**
    - 请求参数变更
      - `+ enableRecording`
  - **CreateWebinar**
    - 请求参数变更
      - `+ enableRecording`
    - 响应参数变更
      - `+ enableRecording`
  - **ShowWebinar**
    - 响应参数变更
      - `+ enableRecording`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeDriverLicense**
    - 响应参数变更
      - `+ result.front`
      - `+ result.back`
  - **RecognizeThailandIdcard**
    - 响应参数变更
      - `+ result.type`
      - `+ result.name_en`
      - `+ result.ref_number`
      - `+ result.confidence.name_en`
      - `+ result.confidence.ref_number`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListRestoreTimes**
    - 响应参数变更
      - `* restore_time.start_time: int32 -> int64`
      - `* restore_time.end_time: int32 -> int64`
  - **ListOffSiteRestoreTimes**
    - 响应参数变更
      - `* restore_time.start_time: int32 -> int64`
      - `* restore_time.end_time: int32 -> int64`
  - **RestoreToExistingInstance**
    - 请求参数变更
      - `* source.restore_time: int32 -> int64`
  - **RestoreExistInstance**
    - 请求参数变更
      - `* source.restore_time: int32 -> int64`
  - **CreateInstance**
    - 请求参数变更
      - `* restore_point.restore_time: int32 -> int64`
    - 响应参数变更
      - `* instance.restore_point.restore_time: int32 -> int64`
  - **CreateRestoreInstance**
    - 请求参数变更
      - `* restore_point.restore_time: int32 -> int64`
    - 响应参数变更
      - `* instance.restore_point.restore_time: int32 -> int64`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 支持接口`ShowConsumerConnections`
- _解决问题_
  - 无
- _特性变更_
  - **ShowConsumerListOrDetails**
    - 响应参数变更
      - `+ total`
      - `+ brokers`
  - **ShowUser**
    - 响应参数变更
      - `- default_group_perm: enum value [PUB,PUB|SUB]`
      - `- group_perms.perm: enum value [PUB,PUB|SUB]`
  - **UpdateUser**
    - 请求参数变更
      - `- default_group_perm: enum value [PUB,PUB|SUB]`
      - `- group_perms.perm: enum value [PUB,PUB|SUB]`
    - 响应参数变更
      - `- default_group_perm: enum value [PUB,PUB|SUB]`
      - `- group_perms.perm: enum value [PUB,PUB|SUB]`
  - **UpdateInstance**
    - 请求参数变更
      - `+ enable_publicip`
      - `+ publicip_id`
  - **CreateUser**
    - 请求参数变更
      - `- default_group_perm: enum value [PUB,PUB|SUB]`
      - `- group_perms.perm: enum value [PUB,PUB|SUB]`
    - 响应参数变更
      - `- default_group_perm: enum value [PUB,PUB|SUB]`
      - `- group_perms.perm: enum value [PUB,PUB|SUB]`
  - **ListUser**
    - 响应参数变更
      - `- users.default_group_perm: enum value [PUB,PUB|SUB]`
      - `- users.group_perms.perm: enum value [PUB,PUB|SUB]`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListVpcsByTags**
    - 响应参数变更
      - `+ resources.resource_detail`
      - `- resources.resouce_detail`
  - **ListSubnetsByTags**
    - 响应参数变更
      - `+ resources.resource_detail`
      - `- resources.resouce_detail`
  - **UpdateRouteTable**
    - 请求参数变更
      - `+ routetable.routes.add`
      - `+ routetable.routes.mod`
      - `+ routetable.routes.del`
      - `* routetable.routes: map<string, list<RouteTableRoute>> -> object<RouteTableRouteAction>`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListServiceDescribeDetails**
    - 响应参数变更
      - `+ public_border_group`
  - **ListServiceDetails**
    - 响应参数变更
      - `+ enable_policy`
      - `+ tcp_proxy: enum value [proxy_vni]`
  - **UpdateEndpointService**
    - 请求参数变更
      - `+ tcp_proxy`
    - 响应参数变更
      - `+ enable_policy`
      - `+ tcp_proxy: enum value [proxy_vni]`
  - **ListEndpointInfoDetails**
    - 响应参数变更
      - `+ endpoint_pool_id`
      - `+ public_border_group`
  - **CreateEndpointService**
    - 请求参数变更
      - `+ enable_policy`
      - `+ tcp_proxy: enum value [proxy_vni]`
    - 响应参数变更
      - `+ enable_policy`
      - `+ tcp_proxy: enum value [proxy_vni]`
  - **ListEndpointService**
    - 响应参数变更
      - `+ endpoint_services.enable_policy`
      - `+ endpoint_services.tcp_proxy: enum value [proxy_vni]`
  - **CreateEndpoint**
    - 响应参数变更
      - `+ endpoint_pool_id`
      - `+ public_border_group`
      - `+ ip`

# 0.1.36 2023-04-13

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListCustomerselfResourceRecordDetails**
    - 响应参数变更
      - `+ monthly_records.pre_order_id`
      - `+ monthly_records.az_code_infos`

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 支持接口`ShowReport`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CPTS

- _新增特性_
  - 支持接口`UpdateAgentHealthStatus`、`ShowAgentConfig`
- _解决问题_
  - 无
- _特性变更_
  - **ShowReport**
    - 响应参数变更
      - `* result.brokens.commonTimestamps: list<integer> -> list<string>`

### HuaweiCloud SDK EVS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowVolume**
    - 响应参数变更
      - `+ volume.iops`
      - `+ volume.throughput`
  - **ListVolumes**
    - 响应参数变更
      - `+ volumes.iops`
      - `+ volumes.throughput`

### HuaweiCloud SDK IES

- _新增特性_
  - 支持接口`ListRacks`、`ShowRack`、`ListStoragePools`、`ShowStoragePool`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Image

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `CreateVideoSynthesisTask`
    - `ShowVideoSynthesisTask`
    - `CreateImageToVideoTask`
    - `ShowImageToVideoTask`
    - `CreateVideoCuttingTask`
    - `ShowVideoCuttingTask`
    - `RunImageWisedesignCrop`
    - `RunImageWisedesignInpainting`
  - **RunImageTagging**
    - 响应参数变更
      - `+ result.tags.instances.bounding_box.width`
      - `+ result.tags.instances.bounding_box.height`
      - `+ result.tags.instances.bounding_box.top_left_x`
      - `+ result.tags.instances.bounding_box.top_left_y`
      - `* result.tags.instances.bounding_box: object -> object<ImageTaggingBoundingBox>`
  - **RunImageMediaTagging**
    - 响应参数变更
      - `+ result.tags.instances.bounding_box.width`
      - `+ result.tags.instances.bounding_box.height`
      - `+ result.tags.instances.bounding_box.top_left_x`
      - `+ result.tags.instances.bounding_box.top_left_y`
      - `* result.tags.instances.bounding_box: object -> object<BoundingBox>`
      - `* result.tags.instances: list<ImageTaggingInstance> -> list<ImageMediaTaggingInstance>`
  - **RunImageMediaTaggingDet**
    - 响应参数变更
      - `+ result.tags.instances.bounding_box.width`
      - `+ result.tags.instances.bounding_box.height`
      - `+ result.tags.instances.bounding_box.top_left_x`
      - `+ result.tags.instances.bounding_box.top_left_y`
      - `* result.tags.instances.bounding_box: object -> object<BoundingBox>`
  - **ShowVideoShotSplitTask**
    - 响应参数变更
      - `- state: enum value [SUCCEEDED,FAILED,RUNNING]`
  - **CreateVideoTranslateTask**
    - 请求参数变更
      - `* body: object<VideoTranslateRequestBody> -> object<CreateVideoTranslateTaskRequestBody>`
  - **CreateImageHighresolutionMattingTask**
    - 请求参数变更
      - `* input.data: list<TaskInputData> -> list<ImageHighresolutionMattingInputData>`
      - `* input: object<TaskInput> -> object<ImageHighresolutionMattingInput>`
  - **ShowImageHighresolutionMattingTask**
    - 响应参数变更
      - `* input.data: list<TaskInputData> -> list<ImageHighresolutionMattingInputData>`
      - `* input: object<TaskInput> -> object<ImageHighresolutionMattingInput>`
  - **CreateImageTranslateTask**
    - 请求参数变更
      - `* input.data: list<TaskInputData> -> list<ImageTranslateTaskInputData>`
      - `* input: object<TaskInput> -> object<ImageTranslateTaskInput>`
      - `* body: object<ImageTranslateRequestBody> -> object<CreateImageTranslateRequestBody>`
  - **ShowImageTranslateTask**
    - 响应参数变更
      - `* input.data: list<TaskInputData> -> list<ImageTranslateTaskInputData>`
      - `* input: object<TaskInput> -> object<ImageTranslateTaskInput>`
  - **CreateVideoCoverAnalysisTask**
    - 请求参数变更
      - `* input.data: list<TaskInputData> -> list<VideoCoverAnalysisTaskInputData>`
      - `* input: object<TaskInput> -> object<VideoCoverAnalysisTaskInput>`
      - `* body: object<VideoCoverAnalysisCreateTaskRequestBody> -> object<CreateVideoCoverAnalysisTaskRequestBody>`
  - **ShowVideoCoverAnalysisTask**
    - 响应参数变更
      - `* input.data: list<TaskInputData> -> list<VideoCoverAnalysisTaskInputData>`
      - `* input: object<TaskInput> -> object<VideoCoverAnalysisTaskInput>`
  - **CreateVideoSummarizationAnalysisTask**
    - 请求参数变更
      - `* input.data: list<TaskInputData> -> list<VideoSummarizationTaskInputData>`
      - `* input: object<TaskInput> -> object<VideoSummarizationTaskInput>`
      - `* body: object<VideoSummarizationCreateTaskRequestBody> -> object<CreateVideoSummarizationTaskRequestBody>`
  - **ShowVideoSummarizationAnalysisTask**
    - 响应参数变更
      - `* input.data: list<TaskInputData> -> list<VideoSummarizationTaskInputData>`
      - `* input: object<TaskInput> -> object<VideoSummarizationTaskInput>`
  - **CreateVideoObjectMaskingTask**
    - 请求参数变更
      - `* input.data: list<TaskInputData> -> list<ObjectMaskingTaskInputData>`
      - `* input: object<TaskInput> -> object<ObjectMaskingTaskInput>`
  - **ShowVideoObjectMaskingTask**
    - 响应参数变更
      - `* input.data: list<TaskInputData> -> list<ObjectMaskingTaskInputData>`
      - `* input: object<TaskInput> -> object<ObjectMaskingTaskInput>`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`BatchDeleteGroup`
- _解决问题_
  - 无
- _特性变更_
  - **ResizeEngineInstance**
    - 请求参数变更
      - `+ publicip_id`

### HuaweiCloud SDK Live

- _新增特性_
  - 支持以下接口：
    - `ListSnapshotConfigs`
    - `UpdateSnapshotConfig`
    - `CreateSnapshotConfig`
    - `DeleteSnapshotConfig`
    - `ShowDomainKeyChain`
    - `UpdateDomainKeyChain`
    - `DeleteDomainKeyChain`
    - `ShowDomainHttpsCert`
    - `UpdateDomainHttpsCert`
    - `DeleteDomainHttpsCert`
    - `UpdateObsBucketAuthorityPublic`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OSM

- _新增特性_
  - 支持接口`ListDiagnoseResources`、`ListOrderIncident`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeShortAudio**
    - 请求参数变更
      - `+ config.property: enum value [english_8k_common,english_16k_common]`
  - **CollectTranscriberJob**
    - 响应参数变更
      - `+ job_id`

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持接口`CreateCloudWafPostPaidResource`、`DeleteCloudWafPostPaidResource`
- _解决问题_
  - 无
- _特性变更_
  - **ListCustomRules**
    - 响应参数变更
      - `+ items.name`

# 0.1.35 2023-04-06

### HuaweiCloud SDK CCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ExportCertificate**
    - 请求参数变更
      - `+ password`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateRefreshTasks**
    - 请求参数变更
      - `+ refresh_task.mode`

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListAlarmHistories**
    - 响应参数变更
      - `+ alarm_histories.type: enum value [DNSHealthCheck,RESOURCE_GROUP,MULTI_INSTANCE,ALL_INSTANCE]`
  - **ListAlarmRules**
    - 响应参数变更
      - `+ alarms.type: enum value [EVENT.SYS,EVENT.CUSTOM,DNSHealthCheck,RESOURCE_GROUP,MULTI_INSTANCE,ALL_INSTANCE]`
  - **CreateAlarmRules**
    - 请求参数变更
      - `+ type: enum value [EVENT.SYS,EVENT.CUSTOM,DNSHealthCheck,RESOURCE_GROUP,MULTI_INSTANCE,ALL_INSTANCE]`

### HuaweiCloud SDK DeH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListServersDedicatedHost**
    - 响应参数变更
      - `- servers.addresses.vpc_id`
      - `* servers.addresses: object<RespAddresses> -> map<string, list<RespAddr>>`

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持接口`SendSms`、`ListSmsDetails`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Image

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CreateTextToImageTask`、`ShowTextToImageTask`、`CreateImageVariationTask`、`ShowImageVariationTask`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`RunJobFlow`
- _解决问题_
  - 无
- _特性变更_
  - **CreateCluster**
    - 请求参数变更
      - `+ log_uri`
      - `+ component_configs`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeFinancialStatement**
    - 请求参数变更
      - `+ return_rectification_matrix`
    - 响应参数变更
      - `+ result.rectification_matrix`

### HuaweiCloud SDK RAM

- _新增特性_
  - 支持接口`ListResourceShareTags`、`ListResourceSharesByTags`、`SearchResourceShareCountByTags`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.34 2023-03-30

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListRenewRateOnPeriod`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 支持接口`ListRenewRateOnPeriod`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListBackups**
    - 请求参数变更
      - `+ incremental`
  - **ListVault**
    - 响应参数变更
      - `+ vaults.billing.object_type: enum value [vmware,rds,file]`
  - **CreateVault**
    - 请求参数变更
      - `+ vault.billing.object_type: enum value [vmware,rds,file]`
    - 响应参数变更
      - `+ vault.billing.object_type: enum value [vmware,rds,file]`
  - **ShowVault**
    - 响应参数变更
      - `+ vault.billing.object_type: enum value [vmware,rds,file]`
  - **UpdateVault**
    - 响应参数变更
      - `+ vault.billing.object_type: enum value [vmware,rds,file]`
  - **ShowVaultResourceInstances**
    - 响应参数变更
      - `+ resources.resource_detail.billing.object_type: enum value [vmware,rds,file]`
  - **ListProtectable**
    - 响应参数变更
      - `+ instances.protectable.vault.billing.object_type: enum value [vmware,rds,file]`
  - **ShowProtectable**
    - 响应参数变更
      - `+ instance.protectable.vault.billing.object_type: enum value [vmware,rds,file]`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`CreateApply`、`CreateEvent`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateAutoCreatePolicy**
    - 请求参数变更
      - `+ indices`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListUsers**
    - 响应参数变更
      - `* user_list.privileges: list<string> -> string`
  - **BatchUpdateUser**
    - 响应参数变更
      - `* results.user_list.privileges: list<string> -> string`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持接口`ShowAdmetProperties`
- _解决问题_
  - 无
- _特性变更_
  - **CreateGenerationTask**
    - 请求参数变更
      - `- strong_constraints.requirement`
  - **CreateOptimizationTask**
    - 请求参数变更
      - `- strong_constraints.requirement`
  - **ShowGenerationTaskResult**
    - 响应参数变更
      - `- task_data.strong_constraints.requirement`
  - **ShowOptimizationTaskResult**
    - 响应参数变更
      - `- task_data.strong_constraints.requirement`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListOtaPackageInfo**
    - 请求参数变更
      - `- Sp-Auth-Token`
  - **CreateOtaPackage**
    - 请求参数变更
      - `- Sp-Auth-Token`
  - **DeleteOtaPackage**
    - 请求参数变更
      - `- Sp-Auth-Token`
  - **ShowOtaPackage**
    - 请求参数变更
      - `- Sp-Auth-Token`
  - **ShowRuleAction**
    - 响应参数变更
      - `+ channel_detail.http_forwarding.signature_enable`
      - `+ channel_detail.http_forwarding.token`
  - **UpdateRuleAction**
    - 请求参数变更
      - `+ channel_detail.http_forwarding.signature_enable`
      - `+ channel_detail.http_forwarding.token`
    - 响应参数变更
      - `+ channel_detail.http_forwarding.signature_enable`
      - `+ channel_detail.http_forwarding.token`
  - **CreateRuleAction**
    - 请求参数变更
      - `+ channel_detail.http_forwarding.signature_enable`
      - `+ channel_detail.http_forwarding.token`
    - 响应参数变更
      - `+ channel_detail.http_forwarding.signature_enable`
      - `+ channel_detail.http_forwarding.token`
  - **ListRuleActions**
    - 响应参数变更
      - `+ actions.channel_detail.http_forwarding.signature_enable`
      - `+ actions.channel_detail.http_forwarding.token`
  - **ShowRule**
    - 响应参数变更
      - `+ actions.device_alarm.dimension`
      - `+ condition_group.conditions.device_linkage_status_condition`
      - `+ condition_group.conditions.device_property_condition.filters.in_values`
  - **UpdateRule**
    - 请求参数变更
      - `+ actions.device_alarm.dimension`
      - `+ condition_group.conditions.device_linkage_status_condition`
      - `+ condition_group.conditions.device_property_condition.filters.in_values`
    - 响应参数变更
      - `+ actions.device_alarm.dimension`
      - `+ condition_group.conditions.device_linkage_status_condition`
      - `+ condition_group.conditions.device_property_condition.filters.in_values`
  - **CreateRule**
    - 请求参数变更
      - `+ actions.device_alarm.dimension`
      - `+ condition_group.conditions.device_linkage_status_condition`
      - `+ condition_group.conditions.device_property_condition.filters.in_values`
    - 响应参数变更
      - `+ actions.device_alarm.dimension`
      - `+ condition_group.conditions.device_linkage_status_condition`
      - `+ condition_group.conditions.device_property_condition.filters.in_values`
  - **ListRules**
    - 响应参数变更
      - `+ rules.actions.device_alarm.dimension`
      - `+ rules.condition_group.conditions.device_linkage_status_condition`
      - `+ rules.condition_group.conditions.device_property_condition.filters.in_values`

### HuaweiCloud SDK NAT

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListPrivateDnats**
    - 响应参数变更
      - `* page_info.current_count: number -> integer`
  - **ListPrivateNats**
    - 响应参数变更
      - `* page_info.current_count: number -> integer`
  - **ListPrivateSnats**
    - 响应参数变更
      - `* page_info.current_count: number -> integer`
  - **ListTransitIps**
    - 响应参数变更
      - `* page_info.current_count: number -> integer`

### HuaweiCloud SDK Organizations

- _新增特性_
  - 支持以下接口：
    - `ShowCreateAccountStatus`
    - `ShowEffectivePolicies`
    - `ListTagPolicyServices`
    - `ListTagResources`
    - `CreateTagResource`
    - `DeleteTagResource`
    - `ListResourceInstances`
    - `ShowResourceInstancesCount`
    - `ListResourceTags`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OSM

- _新增特性_
  - 支持接口`ShowConfiguration`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RAM

- _新增特性_
  - 支持接口`BatchCreateResourceShareTags`、`BatchDeleteResourceShareTags`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **PublishAssetFromObs**
    - 请求参数变更
      - `+ video_type: enum value [RMVB,WEBM]`

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持接口`ShowValueList`、`ShowPrivacyRule`、`ShowAntitamperRule`、`ShowWhiteBlackIpRule`
- _解决问题_
  - 无
- _特性变更_
  - **ShowCcRule**
    - 响应参数变更
      - `+ name`
      - `* mode: number -> int32`
  - **UpdateCcRule**
    - 请求参数变更
      - `+ name`
    - 响应参数变更
      - `+ name`
      - `* mode: number -> int32`
  - **DeleteCcRule**
    - 响应参数变更
      - `+ name`
      - `* mode: number -> int32`
  - **ShowCustomRule**
    - 响应参数变更
      - `+ time`
  - **UpdateCustomRule**
    - 响应参数变更
      - `+ time`
  - **DeleteCustomRule**
    - 响应参数变更
      - `+ time`
  - **ListAntileakageRules**
    - 响应参数变更
      - `+ items.description`
  - **CreateCcRule**
    - 请求参数变更
      - `+ name`
    - 响应参数变更
      - `+ name`
      - `* mode: number -> int32`
  - **ListCcRules**
    - 响应参数变更
      - `+ items.name`
  - **CreateCustomRule**
    - 响应参数变更
      - `+ time`
  - **ListCustomRules**
    - 响应参数变更
      - `+ items.time`

# 0.1.33 2023-03-23

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowNode**
    - 响应参数变更
      - `+ spec.extendParam.agency_name`
  - **UpdateNode**
    - 响应参数变更
      - `+ spec.extendParam.agency_name`
  - **DeleteNode**
    - 响应参数变更
      - `+ spec.extendParam.agency_name`
  - **CreateNode**
    - 请求参数变更
      - `+ spec.extendParam.agency_name`
    - 响应参数变更
      - `+ spec.extendParam.agency_name`
  - **ListNodes**
    - 响应参数变更
      - `+ items.spec.extendParam.agency_name`
  - **ShowNodePool**
    - 响应参数变更
      - `+ spec.nodeTemplate.extendParam.agency_name`
  - **UpdateNodePool**
    - 响应参数变更
      - `+ spec.nodeTemplate.extendParam.agency_name`
  - **DeleteNodePool**
    - 响应参数变更
      - `+ spec.nodeTemplate.extendParam.agency_name`
  - **CreateNodePool**
    - 请求参数变更
      - `+ spec.nodeTemplate.extendParam.agency_name`
    - 响应参数变更
      - `+ spec.nodeTemplate.extendParam.agency_name`
  - **ListNodePools**
    - 响应参数变更
      - `+ items.spec.nodeTemplate.extendParam.agency_name`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowDomainDetailByName**
    - 响应参数变更
      - `- domain.banned_reason`
      - `- domain.locked_reason`
      - `- domain.enterprise_project_id`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`UploadFilePublisher`、`ShowExtensionTestingResult`、`PublishExtension`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **CreateMigrationTask**
    - 请求参数变更
      - `+ backup_files.file_source: enum value [backup_record]`
  - **ShowMigrationTask**
    - 响应参数变更
      - `+ backup_files.file_source: enum value [backup_record]`
  - **StopMigrationTask**
    - 响应参数变更
      - `+ backup_files.file_source: enum value [backup_record]`

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowReplSetName**
    - 响应参数变更
      - `+ name`
  - **UpdateReplSetName**
    - 响应参数变更
      - `+ job_id`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持以下接口：
    - `CreateCpiTask`
    - `ShowCpiTaskResult`
    - `CreateGenerationTask`
    - `ShowGenerationTaskResult`
    - `CreateOptimizationTask`
    - `ShowOptimizationTaskResult`
    - `CreateSearchTask`
    - `ShowSearchTaskResult`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK NAT

- _新增特性_
  - 支持以下接口：
    - `ListPrivateNats`
    - `CreatePrivateNat`
    - `ShowPrivateNat`
    - `UpdatePrivateNat`
    - `DeletePrivateNat`
    - `ListPrivateDnats`
    - `CreatePrivateDnat`
    - `ShowPrivateDnat`
    - `UpdatePrivateDnat`
    - `DeletePrivateDnat`
    - `ListPrivateSnats`
    - `CreatePrivateSnat`
    - `ShowPrivateSnat`
    - `UpdatePrivateSnat`
    - `DeletePrivateSnat`
    - `ListTransitIps`
    - `CreateTransitIp`
    - `ShowTransitIp`
    - `DeleteTransitIp`
    - `ListPrivateNatsByTags`
    - `ListPrivateNatTags`
    - `ShowPrivateNatTags`
    - `CreatePrivateNatTag`
    - `BatchCreateDeletePrivateNatTags`
    - `DeletePrivateNatTag`
    - `ListTransitIpsByTags`
    - `ListTransitIpTags`
    - `ShowTransitIpTags`
    - `CreateTransitIpTag`
    - `BatchCreateDeleteTransitIpTags`
    - `DeleteTransitIpTag`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Organizations

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `ShowCreateAccountStatus`
    - `ShowEffectivePolicies`
    - `ListTagPolicyServices`
    - `ListTagResources`
    - `CreateTagResource`
    - `DeleteTagResource`
    - `ListResourceInstances`
    - `ShowResourceInstancesCount`
    - `ListResourceTags`
    - `CreateAccount`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 支持接口`ShowEngineInstanceExtendProductInfo`、`ResizeEngineInstance`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持以下接口：
    - `ListCcRules`
    - `CreateCcRule`
    - `ShowCcRule`
    - `UpdateCcRule`
    - `DeleteCcRule`
    - `ListCustomRules`
    - `CreateCustomRule`
    - `ShowCustomRule`
    - `UpdateCustomRule`
    - `DeleteCustomRule`
    - `ListAnticrawlerRules`
    - `UpdateAnticrawlerRuleType`
    - `CreateAnticrawlerRule`
    - `ShowAnticrawlerRule`
    - `UpdateAnticrawlerRule`
    - `DeleteAnticrawlerRule`
    - `ListPunishmentRules`
    - `CreatePunishmentRule`
    - `ShowPunishmentRule`
    - `UpdatePunishmentRule`
    - `DeletePunishmentRule`
    - `ListAntileakageRules`
    - `CreateAntileakageRule`
    - `ShowAntileakageRule`
    - `UpdateAntileakageRule`
    - `DeleteAntileakageRule`
    - `UpdateAntiTamperRuleRefresh`
    - `ShowGeoipRule`
    - `ShowIgnoreRule`
    - `UpdateIgnoreRule`
- _解决问题_
  - 无
- _特性变更_
  - **ListHost**
    - 响应参数变更
      - `- items.paid_type: enum value [prePaid]`
  - **DeleteHost**
    - 响应参数变更
      - `- paid_type: enum value [prePaid]`

# 0.1.32 2023-03-16

### HuaweiCloud SDK Organizations

- _新增特性_
  - 支持组织云服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 支持接口`UpdateIngressEipV2`、`AddIngressEipV2`、`RemoveIngressEipV2`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowJobInfos**
    - 响应参数变更
      - `* begin_time: date-time -> string`
      - `* end_time: date-time -> string`
      - `* entities.sub_jobs.begin_time: date-time -> string`
      - `* entities.sub_jobs.end_time: date-time -> string`

### HuaweiCloud SDK CBH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListQuotaStatus**
    - 响应参数变更
      - `* quota: string -> int32`
      - `* eip_quota: string -> int32`
  - **StopCbhInstance**
    - 请求参数变更
      - `- reboot`
  - **ShowNetworkConfiguration**
    - 请求参数变更
      - `+ server_id`
  - **CreateInstanceOrder**
    - 请求参数变更
      - `- end_time`
      - `- relative_resource_id`
      - `- product_infos.available_zone_id`
  - **ChangeInstanceNetwork**
    - 请求参数变更
      - `+ server_id`
  - **ListCbhInstance**
    - 响应参数变更
      - `- instance.is_auto_renew`
  - **CreateInstance**
    - 请求参数变更
      - `- server.image_ref`
      - `- server.name`
      - `- server.personality`
      - `- server.user_data`
      - `- server.admin_password`
      - `- server.key_name`
      - `- server.count`
      - `- server.root_volume`
      - `- server.data_volumes`
      - `- server.extend_param`
      - `- server.metadata`
      - `- server.region_id`
      - `- server.resource_spec_code`
      - `- server.end_time`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ListDomains`、`ShowDomainDetailByName`
- _解决问题_
  - 无
- _特性变更_
  - **ListDomains**
    - 请求参数变更
      - `+ show_tags`
      - `+ exact_match`
    - 响应参数变更
      - `+ domains.tags`

### HuaweiCloud SDK CPH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ShowBandwidthDetail**
    - 请求参数变更
      - `+ offset`
      - `+ limit`
  - **ListJobs**
    - 请求参数变更
      - `+ offset`
      - `+ limit`
  - **ListCloudPhoneModels**
    - 请求参数变更
      - `+ offset`
      - `+ limit`
  - **CreateCloudPhoneServer**
    - 响应参数变更
      - `+ server_ids`
  - **CreateNet2CloudPhoneServer**
    - 响应参数变更
      - `+ server_ids`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持以下接口：
    - `ListMessageStatistics`
    - `ListNotice`
    - `BatchDeleteNotice`
    - `BatchUpdateNotice`
    - `ImportUser`
    - `ListGlobalWorkflowStatistic`
- _解决问题_
  - 无
- _特性变更_
  - **ListDatabaseData**
    - 响应参数变更
      - `* objects: list<map<string, object>> -> list<map<string, string>>`
  - **ImportDatabaseData**
    - 响应参数变更
      - `- creator`
      - `- create_time`
      - `- total_count`
      - `- end_time`
      - `- name`
      - `- finish_count`
      - `- type`
      - `- failed_reason`
      - `- status`
  - **ShowDataJob**
    - 响应参数变更
      - `+ additions`
  - **UpdateJobConfig**
    - 请求参数变更
      - `- job_retain_number`
  - **ShowMessageClearRule**
    - 响应参数变更
      - `- message_retain_time`
  - **UpdateMessageClearRule**
    - 请求参数变更
      - `- message_retain_number`
      - `- message_retain_time`
  - **ShowOverview**
    - 响应参数变更
      - `+ is_arrears`
  - **UpdatePerformanceResource**
    - 请求参数变更
      - `+ schedulable`
  - **ShowEnv**
    - 响应参数变更
      - `+ enable_cold_archive`
  - **ShowUser**
    - 响应参数变更
      - `+ source`
  - **ShowAsset**
    - 响应参数变更
      - `+ versions.description`
      - `- versions.descritpion`
  - **ShowAssetVersion**
    - 响应参数变更
      - `+ version.description`
      - `- version.descritpion`
  - **CreateBackup**
    - 请求参数变更
      - `+ storage_type`
  - **ListBackup**
    - 响应参数变更
      - `+ backups.storage_type`
      - `+ backups.archive_days`
  - **ShowData**
    - 请求参数变更
      - `+ X-Need-Content`
    - 响应参数变更
      - `+ content`
  - **ListDataJob**
    - 响应参数变更
      - `- creator`
      - `- create_time`
      - `- total_count`
      - `- end_time`
      - `- name`
      - `- id`
      - `- finish_count`
      - `- type`
      - `- failed_reason`
      - `- status`
      - `+ data_jobs.additions`
  - **ShowProjectTraceData**
    - 响应参数变更
      - `- allowed_operate`
      - `- deletable`
  - **UpdateMessageReceiveConfig**
    - 请求参数变更
      - `- scope`
      - `- language`
      - `- resource_types`
  - **ShowMessageEmailConfig**
    - 响应参数变更
      - `- password`
  - **UpdateMessageEmailConfig**
    - 请求参数变更
      - `- server`
      - `- subject_prefix`
      - `- password`
      - `- user_name`
      - `- language`
      - `- email`
  - **ListUser**
    - 响应参数变更
      - `+ source`
      - `+ users.source`
  - **ShowTaskInstanceMetricData**
    - 响应参数变更
      - `- metric_name`
      - `- resource_id`
  - **ListPerformanceResourceStat**
    - 响应参数变更
      - `+ performance_resources.schedulable`
  - **ListAsset**
    - 响应参数变更
      - `+ assets.versions.description`
      - `- assets.versions.descritpion`
  - **ListStar**
    - 响应参数变更
      - `+ assets.versions.description`
      - `- assets.versions.descritpion`
  - **ListData**
    - 响应参数变更
      - `+ content`
      - `+ datas.content`
  - **ShowProjectTrace**
    - 响应参数变更
      - `- allowed_operate`
      - `- deletable`
      - `- datas.allowed_operate`
      - `- datas.deletable`
      - `* datas: list<DataRsp> -> list<TraceDataRsp>`
  - **ListComputingResources**
    - 响应参数变更
      - `+ resources.failure_reason`
      - `- resources.ip`
      - `- resources.period_num`
  - **ListDatabaseResource**
    - 响应参数变更
      - `+ resources.failure_reason`
  - **ListPerformanceResources**
    - 响应参数变更
      - `+ resources.running_job_count`
      - `+ resources.failure_reason`
      - `+ resources.schedulable`
  - **ListStorageResources**
    - 响应参数变更
      - `- resources.id`
      - `- resources.name`
  - **ExecuteJob**
    - 请求参数变更
      - `+ tasks.io_acc_type`
  - **CreateAutoJob**
    - 请求参数变更
      - `+ tasks.io_acc_type`
  - **UpdateJob**
    - 请求参数变更
      - `+ tasks.io_acc_type`
  - **ShowJob**
    - 响应参数变更
      - `+ still_running_tasks`
      - `+ task_runtime_info.sub_tasks.pod_create_time`
      - `+ task_runtime_info.sub_tasks.pod_start_time`
      - `+ task_runtime_info.sub_tasks.job_failed_times`
      - `+ tasks.io_acc_type`
  - **UpdateAutoJob**
    - 请求参数变更
      - `+ tasks.io_acc_type`
  - **ShowAutoJob**
    - 响应参数变更
      - `+ tasks.io_acc_type`
  - **ShowWorkflow**
    - 响应参数变更
      - `+ tasks.io_acc_type`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **ListBandwidthPkg**
    - 请求参数变更
      - `+ limit`
      - `+ marker`
      - `+ offset`
  - **ListCommonPools**
    - 请求参数变更
      - `+ limit`
      - `+ offset`
  - **ListShareBandwidthTypes**
    - 请求参数变更
      - `+ marker`
      - `+ offset`

### HuaweiCloud SDK IAM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **KeystoneListMappings**
    - 响应参数变更
      - `* mappings.rules.local.groups: object -> string`
  - **KeystoneShowMapping**
    - 响应参数变更
      - `* mapping.rules.local.groups: object -> string`
  - **KeystoneCreateMapping**
    - 请求参数变更
      - `* mapping.rules.local.groups: object -> string`
    - 响应参数变更
      - `* mapping.rules.local.groups: object -> string`
  - **KeystoneUpdateMapping**
    - 请求参数变更
      - `* mapping.rules.local.groups: object -> string`
    - 响应参数变更
      - `* mapping.rules.local.groups: object -> string`

### HuaweiCloud SDK Image

- _新增特性_
  - 支持接口`CreateVideoObjectMaskingTask`、`ShowVideoObjectMaskingTask`
- _解决问题_
  - 无
- _特性变更_
  - **CreateTextToImageTask**
    - 请求参数变更
      - `+ config.common.inference.image_nums`
      - `+ config.common.inference.resolution: enum value [512*768,768*512,512*512]`
  - **ShowTextToImageTask**
    - 响应参数变更
      - `+ config.common.inference.image_nums`
      - `+ config.common.inference.resolution: enum value [512*768,768*512,512*512]`
  - **CreateImageVariationTask**
    - 请求参数变更
      - `+ config.common.inference.image_nums`
      - `+ config.common.inference.resolution: enum value [512*768,768*512,512*512]`
  - **ShowImageVariationTask**
    - 响应参数变更
      - `+ config.common.inference.image_nums`
      - `+ config.common.inference.resolution: enum value [512*768,768*512,512*512]`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 支持以下接口：
    - `BatchListDcDs`
    - `CreateDs`
    - `ShowDcDs`
    - `UpdateDcDs`
    - `DeleteDcDs`
    - `SynchronizeDcConfigs`
    - `BatchListDcDevices`
    - `BatchListDcPoints`
    - `CreateDcPoint`
    - `ShowDcPoint`
    - `UpdateDcPoint`
    - `DeleteDcPoint`
    - `ImportPoints`
    - `ShowPointTemplate`
    - `ShowPoints`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IVS

- _新增特性_
  - 支持接口`DetectStandardByVideoAndIdCardImage`、`DetectStandardByVideoAndNameAndId`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - **RecognizeMvsInvoice**
    - 请求参数变更
      - `+ return_text_location`
      - `+ return_confidence`
      - `+ type`
    - 响应参数变更
      - `+ result.buyer_address`
      - `+ result.buyer_phone`
      - `+ result.licence_plate_number`
      - `+ result.registration_number`
      - `+ result.dept_motor_vehicles`
      - `+ result.auction_org_name`
      - `+ result.auction_org_address`
      - `+ result.auction_org_id`
      - `+ result.auction_org_bank_account`
      - `+ result.auction_org_phone`
      - `+ result.used_vehicle_market_name`
      - `+ result.used_vehicle_market_id`
      - `+ result.used_vehicle_market_address`
      - `+ result.used_vehicle_market_bank_account`
      - `+ result.used_vehicle_market_phone`
      - `+ result.remark`
      - `+ result.drawer_name`
      - `+ result.type`
      - `+ result.text_location`
      - `+ result.confidence`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListRecycleInstances`、`ShowRecyclePolicy`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.31 2023-03-14

### HuaweiCloud SDK Image

- _新增特性_
  - 支持接口`CreateTextToImageTask`、`ShowTextToImageTask`、`CreateImageVariationTask`、`ShowImageVariationTask`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`RunImageWisedesignColorfilter`、`RunImageWisedesignCombine`

# 0.1.30 2023-03-09

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持以下接口：
    - `ListTemplates`
    - `DeleteTemplate`
    - `ShowTemplateMetadata`
    - `UpdateTemplateMetadata`
    - `ShowTemplateVersionContent`
    - `DeleteTemplateVersion`
    - `ShowTemplateVersionMetadata`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListExecutionPlans`:
    - 响应参数`status`新增枚举值`APPLY_IN_PROGRESS`
    - 响应参数`stack_name`、`execution_plan_name`改为必填
  - 接口`GetExecutionPlan`响应参数`action`移除枚举值`IN_PLACE_UPDATE`
  - 接口`GetExecutionPlanMetadata`响应参数`stack_name`、`execution_plan_name`改为必填
  - 接口`ListStackResources`新增响应参数 `resource_attributes`
  - 接口`EstimateExecutionPlanPrice`:
    - 新增响应参数 `unsupported_message`
    - 响应参数`sale_price`类型调整 `object` -> `double`
    - 响应参数`discount`类型调整 `object` -> `double`
    - 响应参数`original_price`类型调整 `object` -> `double`
    - 响应参数`period_type`移除枚举值`WEEK`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`ShowCategoryList`、`ListPublisher`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateDevice`新增响应参数 `connection_status_update_time`、`active_time`
  - 接口`ShowDevice`新增响应参数 `connection_status_update_time`、`active_time`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeAutoClassification`新增请求参数 `extended_parameters`

# 0.1.29 2023-03-07

### HuaweiCloud SDK VOD

- _新增特性_
  - 支持接口`ModifySubtitle`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.28 2023-03-02

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateWorkflow`:
    - 移除请求参数 `trigger`
    - 响应参数`template_i18n`类型调整
  - 接口`ListWorkflow`响应参数`template_i18n`类型调整

### HuaweiCloud SDK BCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchInviteMembersToChannel`新增响应参数 `result`
  - 接口`ListNotifications`响应参数`node_count`类型调整 `string` -> `int32`
  - 接口`ListMembers`响应参数`node_count`类型调整 `string` -> `int32`
  - 接口`DownloadBlockchainCert`移除响应参数 `result`
  - 接口`DownloadBlockchainSdkConfig`移除响应参数 `result`
  - 接口`CreateBlockchainCertByUserName`移除响应参数 `result`

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListBareMetalServers`:
    - 响应参数`status`新增枚举值`HARD_REBOOT`、`DELETED`
    - 响应参数`OS-EXT-STS:vm_state`移除枚举值`suspended`
  - 接口`CreateBareMetalServers`:
    - 新增请求参数 `chargingMode`
    - 移除请求参数 `chargingmode`
  - 接口`ChangeBaremetalServerName`响应参数`OS-EXT-STS:vm_state`移除枚举值`SUSPENDED`
  - 接口`ListBareMetalServerDetails`:
    - 响应参数`status`新增枚举值`HARD_REBOOT`、`DELETED`
    - 响应参数`OS-EXT-STS:vm_state`移除枚举值`suspended`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowDomainFullConfig`:
    - 响应参数`error_code`类型调整 `string` -> `int32`
    - 响应参数`target_code`类型调整 `string` -> `int32`
  - 接口`UpdateDomainFullConfig`:
    - 请求参数`error_code`类型调整 `string` -> `int32`
    - 请求参数`target_code`类型调整 `string` -> `int32`

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCommits`新增请求参数 `page`、`per_page`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持接口`ShowReplSetName`、`UpdateReplSetName`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DRIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListV2xEdges`新增响应参数 `ip`

### HuaweiCloud SDK EG

- _新增特性_
  - 支持以下接口：
    - `CheckPutEvents`
    - `ListObsBuckets`
    - `ListWorkflowTriggers`
    - `ListPubMetrics`
    - `ListSubMetrics`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEventSource`请求参数`type`新增枚举值`ROCKETMQ`
  - 接口`CreateChannel`请求参数`name`改为必填
  - 接口`ListSubscriptions`:
    - 新增响应参数 `used`
    - 响应参数`transform`类型调整
  - 接口`CreateSubscription`:
    - 新增响应参数 `used`
    - 请求参数`transform`类型调整
    - 请求参数`detail`改为必填
    - 响应参数`transform`类型调整
  - 接口`ShowDetailOfSubscription`:
    - 新增响应参数 `used`
    - 响应参数`transform`类型调整
  - 接口`UpdateSubscription`:
    - 新增响应参数 `used`
    - 请求参数`transform`类型调整
    - 请求参数`detail`改为必填
    - 响应参数`transform`类型调整
  - 接口`CreateSubscriptionTarget`:
    - 请求参数`transform`类型调整
    - 请求参数`detail`改为必填
    - 响应参数`transform`类型调整
  - 接口`ShowDetailOfSubscriptionTarget`响应参数`transform`类型调整
  - 接口`UpdateSubscriptionTarget`:
    - 请求参数`transform`类型调整
    - 请求参数`detail`改为必填
    - 响应参数`transform`类型调整
  - 接口`CreateConnection`请求参数`vpc_id`、`subnet_id`改为必填
  - 接口`ListAgencies`请求参数`type`新增枚举值`EG_RESTORE_AGENCY`
  - 接口`CreateAgencies`请求参数`type`新增枚举值`EG_RESTORE_AGENCY`
  - 接口`ListQuotas`:
    - 请求参数`type`新增枚举值`SOURCE_ROCKETMQ`
    - 响应参数`type`新增枚举值`SOURCE_ROCKETMQ`
  - 接口`ListTriggers`:
    - 新增响应参数 `used`
    - 响应参数`transform`类型调整

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateTrigger`:
    - 新增请求参数 `event_data`
    - 新增响应参数 `trigger_id`、`trigger_type_code`、`trigger_status`、`event_data`、`last_updated_time`、`created_time`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持以下接口：
    - `ResetDbUserPassword`
    - `ModifyDbUserPrivilege`
    - `ListDbUsers`
    - `CreateDbUser`
    - `DeleteDbUser`
    - `ListInstanceDatabases`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSimCards`新增请求参数 `order_ids`

### HuaweiCloud SDK Image

- _新增特性_
  - 支持以下接口：
    - `CreateVideoSynthesisTask`
    - `ShowVideoSynthesisTask`
    - `CreateImageToVideoTask`
    - `ShowImageToVideoTask`
    - `CreateVideoCuttingTask`
    - `ShowVideoCuttingTask`
    - `CreateVideoTranslateTask`
    - `ShowVideoTranslateTask`
    - `CreateImageHighresolutionMattingTask`
    - `ShowImageHighresolutionMattingTask`
    - `CreateImageTranslateTask`
    - `ShowImageTranslateTask`
    - `CreateVideoCoverAnalysisTask`
    - `ShowVideoCoverAnalysisTask`
    - `CreateVideoSummarizationAnalysisTask`
    - `ShowVideoSummarizationAnalysisTask`
    - `CreateVideoShotSplitTask`
    - `ShowVideoShotSplitTask`
    - `RunImageWisedesignCrop`
    - `RunImageWisedesignInpainting`
    - `RunImageWisedesignColorfilter`
    - `RunImageWisedesignCombine`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 支持接口`ListOtaPackageInfo`、`CreateOtaPackage`、`ShowOtaPackage`、`DeleteOtaPackage`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstanceConsumerGroups`:
    - 新增响应参数 `groups`
    - 移除响应参数 `group_ids`、`next_offset`、`previous_offset`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`新增请求参数 `external_datasources`、`effective_days`
  - 接口`ShowAutoScalingPolicy`新增响应参数 `effective_days`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`ListSpecIssueStayTimes`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListSslCertDownloadLink`
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateConfiguration`新增响应参数 `configuration`
  - 接口`UpdatePostgresqlInstanceAlias`请求参数`alias`改为非必填
  - 接口`UpdateDatabase`请求参数`comment`改为非必填

# 0.1.27 2023-02-23

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListCustomerAccountChangeRecords`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListVault`响应参数`value`改为必填
  - 接口`CreateVault`:
    - 请求参数`value`改为必填
    - 响应参数`value`改为必填
  - 接口`ShowVault`响应参数`value`改为必填
  - 接口`UpdateVault`:
    - 请求参数`value`改为必填
    - 响应参数`value`改为必填
  - 接口`ListProtectable`响应参数`value`改为必填
  - 接口`ShowProtectable`响应参数`value`改为必填
  - 接口`ShowVaultResourceInstances`响应参数`value`改为必填

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRequest`:
    - 新增请求参数 `request_type`、`above_text`、`following_text`
    - 新增响应参数 `dispatched_task_number`
    - 请求参数`signature`改为非必填
  - 接口`ShowResult`新增响应参数 `request_type`

### HuaweiCloud SDK CPH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ChangeCloudPhoneServerModel`新增请求参数 `phone_model_name`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateLoadIkThesaurus`请求参数`mainObject`、`stopObject`、`synonymObject`改为非必填

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持接口`ShrinkInstanceNodes`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持接口`CreateClusterV2`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRoutingRules`新增请求参数 `active`
  - 接口`CreateRuleAction`新增请求参数 `mysql_forwarding`
  - 接口`ListRuleActions`新增响应参数 `mysql_forwarding`
  - 接口`UpdateRuleAction`:
    - 新增请求参数 `mysql_forwarding`
    - 新增响应参数 `mysql_forwarding`
  - 接口`ShowRuleAction`新增响应参数 `mysql_forwarding`

### HuaweiCloud SDK LakeFormation

- _新增特性_
  - 支持接口`BatchUpdateLakeFormationInstanceTags`、`CreateLakeFormationInstance`、`ListPolicy`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowLakeFormationInstance`移除响应参数 `key`、`value`
  - 接口`SetPartitionColumnStatistics`移除响应参数 `column_statistics_desc`、`column_statistics_objects`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJobExeListNew`:
    - 请求参数`job_state`新增枚举值`FAILED`、`KILLED`、`NEW`、`NEW_SAVING`、`SUBMITTED`、`ACCEPTED`、`RUNNING`、`FINISHED`, 移除枚举值`FAILED：失败`、`KILLED：已终止`、`NEW：已创建`、`NEW_SAVING：已创建保存中`、`SUBMITTED：已提交`、`ACCEPTED：已接受`、`RUNNING：运行中`、`FINISHED：已完成`
    - 请求参数`job_result`新增枚举值`FAILED`、`KILLED`、`UNDEFINED`、`SUCCEEDED`, 移除枚举值`FAILED：执行失败的作业。`、`KILLED：执行中被手动终止的作业。`、`UNDEFINED：正在执行的作业。`、`SUCCEEDED：执行成功的作业。`
  - 接口`ShowHdfsFileList`请求参数`sort_key`新增枚举值`path_suffix`、`length`、`modification_time`, 移除枚举值`path_suffix：文件或目录名称`、`length：文件大小`、`modification_time：修改时间`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`UpdateDbUserPrivilege`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数 `X-Client-Token`

### HuaweiCloud SDK SCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ImportCertificate`请求参数`certificate_chain`改为非必填
  - 接口`ExportCertificate`新增响应参数 `entire_certificate`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdatePremiumHost`新增请求参数 `flag`

# 0.1.26 2023-02-16

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateApp`移除响应参数 `id`
  - 接口`ListResourceUnderNode`:
    - 新增请求参数 `ci_ids`
    - 请求参数`ci_id`改为非必填

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`CreateLogin`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListExtensions`移除响应参数 `id`、`property_name`、`property_value`、`extension_version_id`、`created_at`、`updated_at`
  - 接口`ShowExtensionDetail`移除响应参数 `id`、`property_name`、`property_value`、`extension_version_id`、`created_at`、`updated_at`

### HuaweiCloud SDK CSS

- _新增特性_
  - 支持接口`ChangeSecurityGroup`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ChangeMode`:
    - 新增请求参数 `changeModeReq`
    - 移除请求参数 `changeModeRequestBody`
  - 接口`AddIndependentNode`:
    - 新增请求参数 `IndependentReq`
    - 移除请求参数 `IndependentRequestBody`

### HuaweiCloud SDK CTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateTracker`新增请求参数 `compress_type`、`is_sort_by_service`
  - 接口`CreateTracker`新增请求参数 `compress_type`、`is_sort_by_service`
  - 接口`ListTrackers`新增响应参数 `compress_type`、`is_sort_by_service`

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`CreateCustomTemplate`、`CreateAutoExpireScanTask`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DRIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowHistoryTrafficEvents`新增响应参数 `esn`

### HuaweiCloud SDK DWS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClusterSnapshots`新增请求参数 `project_id`、`cluster_id`、`limit`、`offset`、`sort_key`、`sort_dir`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DetectLiveByUrl`新增请求参数 `nod_threshold`
  - 接口`DetectLiveByUrlIntl`新增请求参数 `nod_threshold`
  - 接口`DetectLiveByFile`新增请求参数 `nod_threshold`
  - 接口`DetectLiveByFileIntl`新增请求参数 `nod_threshold`
  - 接口`DetectLiveByBase64`新增请求参数 `nod_threshold`
  - 接口`DetectLiveByBase64Intl`新增请求参数 `nod_threshold`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddDatabasePermission`请求参数`host`改为必填

### HuaweiCloud SDK Image

- _新增特性_
  - 支持接口`RunImageSuperResolution`、`RunRecaptureDetect`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OSM

- _新增特性_
  - 支持以下接口：
    - `CreateAuthorization`
    - `CreateFeedback`
    - `ShowDownloadAccessoryUrl`
    - `CreateQuestionInSession`
    - `CreateSession`
    - `CreateEvaluate`
    - `ListFeedbackOption`
    - `CreateQaFeedbacks`
    - `CreateAskQuestion`
    - `ShowQaPairDetail`
    - `ShowAssociatedQuestions`
    - `ShowQaPairs`
    - `ListRecommendWords`
    - `CreateQaAsk`
    - `ShowTheme`
    - `ListArticles`
    - `ListNotices`
    - `ListTools`
    - `CreateDiagnoseFeedback`
    - `ListDiagnoseItems`
    - `ListDiagnoseJob`
    - `ListDiagnoseRecords`
    - `CreateDiagnoseJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 支持接口`CreateInstanceByEngine`、`BatchCreateOrDeleteRocketmqTag`、`ShowRocketmqTags`、`ShowRocketmqProjectTags`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增请求参数 `limit`、`offset`
  - 接口`ShowConsumerListOrDetails`新增请求参数 `limit`、`offset`
  - 接口`ListConsumerGroupOfTopic`新增请求参数 `limit`、`offset`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateHost`:
    - 请求参数`tls`移除枚举值`TLS v1.3`
    - 响应参数`tls`移除枚举值`TLS v1.3`
  - 接口`ShowHost`响应参数`tls`移除枚举值`TLS v1.3`
  - 接口`CreatePremiumHost`响应参数`tls`移除枚举值`TLS v1.3`
  - 接口`ShowPremiumHost`响应参数`tls`移除枚举值`TLS v1.3`
  - 接口`UpdatePremiumHost`:
    - 请求参数`tls`移除枚举值`TLS v1.3`
    - 响应参数`tls`移除枚举值`TLS v1.3`

# 0.1.25 2023-02-09

### HuaweiCloud SDK CloudArtifact

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - `制品仓库 CloudArtifact`更名为`制品仓库 CodeArtsArtifact`

### HuaweiCloud SDK CloudBuild

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - `编译构建 CloudBuild`更名为`编译构建 CodeArts Build`

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RenewalResources`新增响应参数 `fail_resource_infos`
  - 接口`CancelResourcesSubscription`新增响应参数 `fail_resource_infos`

### HuaweiCloud SDK DBSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SwitchRiskRule`:
    - 新增响应参数 `status`
    - 移除响应参数 `result`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持接口`ListLtsSlowLogs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAsyncInvocations`新增响应参数 `error_code`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowPauseResumeStutus`:
    - 新增请求参数 `X-Auth-Token`
    - 新增响应参数 `master_instance_id`、`slave_instance_id`、`data_sync_indicators`、`rto_and_rpo_indicators`
    - 移除请求参数 `x-auth-token`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRuleActions`新增响应参数 `roma_forwarding`、`influxdb_forwarding`、`functiongraph_forwarding`、`mrs_kafka_forwarding`、`dms_rocketmq_forwarding`
  - 接口`CreateRuleAction`新增请求参数 `roma_forwarding`、`influxdb_forwarding`、`functiongraph_forwarding`、`mrs_kafka_forwarding`、`dms_rocketmq_forwarding`
  - 接口`UpdateRuleAction`:
    - 新增请求参数 `roma_forwarding`、`influxdb_forwarding`、`functiongraph_forwarding`、`mrs_kafka_forwarding`、`dms_rocketmq_forwarding`
    - 新增响应参数 `roma_forwarding`、`influxdb_forwarding`、`functiongraph_forwarding`、`mrs_kafka_forwarding`、`dms_rocketmq_forwarding`
  - 接口`ShowRuleAction`新增响应参数 `roma_forwarding`、`influxdb_forwarding`、`functiongraph_forwarding`、`mrs_kafka_forwarding`、`dms_rocketmq_forwarding`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeTollInvoice`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListErrorlogForLts`、`ListSlowlogForLts`、`ListSlowLogStatisticsForLts`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.24 2023-02-02

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`UpdateStack`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ApplyExecutionPlan`新增请求参数 `executor`
  - 接口`ListStackEvents`:
    - 新增请求参数 `filter`、`field`
    - 移除响应参数 `resource_id_key`、`resource_id_value`、`resource_name`、`resource_type`、`time`、`event_type`、`event_message`、`elapsed_seconds`
  - 接口`GetStackMetadata`响应参数`stack_name`改为必填
  - 接口`CreateStack`新增请求参数 `executor`
  - 接口`ListStackResources`:
    - 新增响应参数 `index_key`
    - 响应参数`resource_status`移除枚举值`DELETION_SKIPPED`
  - 接口`DeployStack`新增请求参数 `executor`

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RenewalResources`新增响应参数 `fail_resource_infos`
  - 接口`CancelResourcesSubscription`新增响应参数 `fail_resource_infos`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
  - 支持以下接口：
    - `RunPipeline`
    - `BatchShowPipelinesLatestStatus`
    - `ShowPipelineRunDetail`
    - `ListPipelines`
    - `DeletePipeline`
    - `StopPipelineRun`
    - `ListPipelineRuns`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`请求参数`type`新增枚举值`GaussDB`
  - 接口`ListInstances`请求参数`datastore_type`新增枚举值`GaussDB`
  - 接口`ListBackups`响应参数`type`新增枚举值`GaussDB`

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListBackPools`新增请求参数 `all_billing_cycle`
  - 接口`ListSimPools`新增请求参数 `all_billing_cycle`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateImage`请求参数`type`新增枚举值`IsoImage`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`CreateInstanceByEngine`、`ShowEngineInstanceExtendProductInfo`、`ResizeEngineInstance`、`CreateReassignmentTask`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `sasl_enabled_mechanisms`
  - 接口`CreatePostPaidInstance`请求参数`kafka_manager_user`、`kafka_manager_password`改为非必填
  - 接口`ShowInstance`新增响应参数 `sasl_enabled_mechanisms`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunTextModeration`新增请求参数 `white_glossaries`

# 0.1.23 2023-01-19

### HuaweiCloud SDK CCM

- _新增特性_
  - 支持接口`EnableCertificateAuthorityCrl`、`DisableCertificateAuthorityCrl`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCertificate`新增响应参数 `enc_cert_info`
  - 接口`ShowCertificate`新增响应参数 `enc_cert_info`
  - 接口`ExportCertificate`:
    - 新增请求参数 `is_sm_standard`
    - 新增响应参数 `enc_certificate`、`enc_private_key`、`enc_sm2_enveloped_key`

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持接口`ListMonitorIndicatorData`、`ListMonitorIndicators`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK HSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRiskConfigCheckRules`:
    - 新增响应参数 `enable_fix`、`enable_click`
    - 移除响应参数 `fix_status`、`enable_auto_fix`
  - 接口`ShowCheckRuleDetail`新增请求参数 `check_type`
  - 接口`SwitchHostsProtectStatus`请求参数`version`、`host_id_list`改为必填

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 支持接口`UploadBatchTaskFile`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowEdgeNode`新增响应参数 `virtual_ipv6_address`

# 0.1.22 2023-01-12

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListApiGroupsV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`UpdateApiGroupV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`ShowDetailsOfApiGroupV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`UpdateApiV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`ShowDetailsOfApiV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`BatchPublishOrOfflineApiV2`:
    - 新增请求参数 `group_id`
    - 请求参数`env_id`改为必填
  - 接口`ListApiVersionDetailV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`CreateInstanceV2`新增请求参数 `bandwidth_charging_mode`、`ingress_bandwidth_size`、`ingress_bandwidth_charging_mode`
  - 接口`UpdateInstanceV2`新增响应参数 `bandwidth_charging_mode`、`ingress_bandwidth_charging_mode`
  - 接口`ShowDetailsOfInstanceV2`新增响应参数 `bandwidth_charging_mode`、`ingress_bandwidth_charging_mode`
  - 接口`AddEngressEipV2`新增请求参数 `bandwidth_charging_mode`
  - 接口`UpdateEngressEipV2`新增请求参数 `bandwidth_charging_mode`
  - 接口`ImportMicroservice`请求参数`version`改为非必填
  - 接口`CreateCertificateV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`ListCertificatesV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`UpdateCertificateV2`新增响应参数 `is_has_trusted_root_ca`
  - 接口`ShowDetailsOfCertificateV2`新增响应参数 `is_has_trusted_root_ca`

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListFreeResourcesUsageRecords`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerselfResourceRecordDetails`新增响应参数 `sub_service_type_code`、`sub_service_type_name`、`sub_resource_type_code`、`sub_resource_type_name`、`sub_resource_id`、`sub_resource_name`
  - 接口`ListCustomerselfResourceRecords`新增响应参数 `sub_service_type_code`、`sub_service_type_name`、`sub_resource_type_code`、`sub_resource_type_name`、`sub_resource_id`、`sub_resource_name`
  - 接口`ListCosts`请求参数`operator`改为必填

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 支持接口`ListFreeResourcesUsageRecords`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerselfResourceRecordDetails`新增响应参数 `sub_service_type_code`、`sub_service_type_name`、`sub_resource_type_code`、`sub_resource_type_name`、`sub_resource_id`、`sub_resource_name`
  - 接口`ListCustomerselfResourceRecords`新增响应参数 `sub_service_type_code`、`sub_service_type_name`、`sub_resource_type_code`、`sub_resource_type_name`、`sub_resource_id`、`sub_resource_name`
  - 接口`ListCosts`请求参数`operator`改为必填

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAlarmHistories`:
    - 新增响应参数 `data_points`
    - 移除响应参数 `datapoints`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDbObjects`新增响应参数 `object_scope`
  - 接口`ShowDbObjectCollectionStatus`新增响应参数 `object_scope`

### HuaweiCloud SDK eiHealth

- _新增特性_
  - 支持以下接口：
    - `ExecuteAssetAction`
    - `UpdateDataPathPolicy`
    - `PublishData`
    - `PublishImage`
    - `BatchDeleteLabel`
    - `BatchDownloadResourceStatData`
    - `PublishApp`
    - `ShowTaskInstanceMetricData`
    - `BatchCancelJob`
    - `BatchRetryJob`
    - `BatchDeleteJob`
    - `PublishWorkflow`
    - `UpdateAssetVersion`
    - `DeleteAssetVersion`
    - `UpdateJob`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListRecentJob`
  - 接口`ListAsset`移除响应参数 `latest_version`
  - 接口`ShowAsset`移除响应参数 `latest_version`
  - 接口`ShowAssetVersion`新增响应参数 `failed_reason`、`labels`、`picture`
  - 接口`ListStar`移除响应参数 `latest_version`
  - 接口`CreateDatabaseData`请求参数`column`改为必填
  - 接口`UpdateDatabaseData`请求参数`column`改为必填
  - 接口`ListData`新增响应参数 `deletable`
  - 接口`ShowData`新增响应参数 `deletable`
  - 接口`ListDataJob`新增响应参数 `failed_reason`
  - 接口`UpdateProject`新增请求参数 `storage_quota`
  - 接口`ShowProject`新增响应参数 `storage_quota`
  - 接口`ShowProjectTrace`移除响应参数 `path`、`name`、`type`、`size`、`create_time`、`download_url`
  - 接口`ShowProjectTraceData`新增响应参数 `allowed_operate`、`deletable`
  - 接口`ListComputingResources`:
    - 新增响应参数 `node_labels`
    - 移除请求参数 `label`、`offset`、`limit`
  - 接口`UpdateJobConfig`移除请求参数 `job_retain_time`
  - 接口`ShowJobConfig`移除响应参数 `job_retain_time`
  - 接口`ListLabel`移除响应参数 `count`
  - 接口`BatchUpdateNodeLabel`移除请求参数 `name`
  - 接口`DeleteUser`移除请求参数 `user_id_type`
  - 接口`ShowTaskInstances`新增响应参数 `host_name`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DetectLiveByUrl`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectLiveByUrlIntl`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectLiveFaceByUrl`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectLiveByFile`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectLiveByFileIntl`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectLiveFaceByFile`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectLiveByBase64`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectLiveByBase64Intl`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectLiveFaceByBase64`新增请求参数 `Enterprise-Project-Id`
  - 接口`SearchFaceByFaceId`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectFaceByFile`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectFaceByFileIntl`新增请求参数 `Enterprise-Project-Id`
  - 接口`UpdateFace`新增请求参数 `Enterprise-Project-Id`
  - 接口`DeleteFaceByExternalImageId`新增请求参数 `Enterprise-Project-Id`
  - 接口`ShowFacesByLimit`新增请求参数 `Enterprise-Project-Id`
  - 接口`CompareFaceByFile`新增请求参数 `Enterprise-Project-Id`
  - 接口`DeleteFaceByFaceId`新增请求参数 `Enterprise-Project-Id`
  - 接口`ShowFacesByFaceId`新增请求参数 `Enterprise-Project-Id`
  - 接口`AddFacesByBase64`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectFaceByUrl`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectFaceByUrlIntl`新增请求参数 `Enterprise-Project-Id`
  - 接口`DeleteFaceSet`新增请求参数 `Enterprise-Project-Id`
  - 接口`ShowFaceSet`新增请求参数 `Enterprise-Project-Id`
  - 接口`CompareFaceByBase64`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectFaceByBase64`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectFaceByBase64Intl`新增请求参数 `Enterprise-Project-Id`
  - 接口`CreateFaceSet`新增请求参数 `Enterprise-Project-Id`
  - 接口`ShowAllFaceSets`新增请求参数 `Enterprise-Project-Id`
  - 接口`SearchFaceByFile`新增请求参数 `Enterprise-Project-Id`
  - 接口`AddFacesByUrl`新增请求参数 `Enterprise-Project-Id`
  - 接口`AddFacesByFile`新增请求参数 `Enterprise-Project-Id`
  - 接口`SearchFaceByUrl`新增请求参数 `Enterprise-Project-Id`
  - 接口`SearchFaceByBase64`新增请求参数 `Enterprise-Project-Id`
  - 接口`CompareFaceByUrl`新增请求参数 `Enterprise-Project-Id`
  - 接口`BatchDeleteFaces`新增请求参数 `Enterprise-Project-Id`

### HuaweiCloud SDK HSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateProtectionPolicy`请求参数`policy_id`、`policy_name`、`protection_mode`、`bait_protection_status`、`protection_directory`、`protection_type`、`operating_system`改为必填
  - 接口`StartProtection`请求参数`operating_system`、`ransom_protection_status`、`backup_protection_status`、`policy_id`、`pattern`、`agent_id_list`、`host_id_list`改为必填
  - 接口`StopProtection`请求参数`host_id_list`、`agent_id_list`、`close_protection_type`改为必填
  - 接口`UpdateBackupPolicyInfo`请求参数`policy_id`、`pattern`改为必填
  - 接口`ListQuotasDetail`新增响应参数 `enterprise_project_id`、`enterprise_project_name`
  - 接口`ListSecurityEvents`新增响应参数 `host_status`、`agent_status`、`protect_status`、`asset_value`、`keyword`、`hash`
  - 接口`ChangeEvent`:
    - 新增请求参数 `keyword`、`hash`
    - 移除请求参数 `description`
  - 接口`DeleteHostsGroup`请求参数`group_id`改为非必填
  - 接口`AddHostsGroup`请求参数`group_name`、`host_id_list`改为必填
  - 接口`AssociatePolicyGroup`请求参数`target_policy_group_id`改为必填
  - 接口`ChangeVulStatus`请求参数`operate_type`、`vul_id`、`host_id_list`改为必填
  - 接口`SetWtpProtectionStatusInfo`:
    - 新增请求参数 `charging_mode`
    - 移除请求参数 `payment_mode`

### HuaweiCloud SDK IVS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DetectStandardByIdCardImage`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectStandardByNameAndId`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectExtentionByNameAndId`新增请求参数 `Enterprise-Project-Id`
  - 接口`DetectExtentionByIdCardImage`新增请求参数 `Enterprise-Project-Id`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeCustomTemplate`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeGeneralTable`:
    - 新增请求参数 `return_char_location`、`return_rectification_matrix`
    - 新增响应参数 `char_list`
  - 接口`RecognizeGeneralText`新增请求参数 `language`
  - 接口`RecognizeWebImage`:
    - 新增响应参数 `extracted_data`
    - 移除响应参数 `extracted_data`、`contact_info`、`image_size`、`height`、`width`、`name`、`phone`、`province`、`city`、`district`、`detail_address`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 支持接口`ListRocketInstanceTopics`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowOneTopic`新增响应参数 `name`

# 0.1.21 2023-01-05

### HuaweiCloud SDK CPH

- _新增特性_
  - 支持以下接口：
    - `ListProjectTags`
    - `ListResourceTags`
    - `ListResourceInstances`
    - `BatchCreateTags`
    - `BatchDeleteTags`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCloudPhoneServers`新增响应参数 `enterprise_project_id`
  - 接口`ShowCloudPhoneServerDetail`新增响应参数 `enterprise_project_id`

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`ListConfigHistories`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `ListResoleRules`
    - `SetPrivateZoneProxyPattern`
    - `ShowDomainQuota`
    - `ShowRetrieval`
    - `CreateRetrieval`
    - `ShowRetrievalVerification`
    - `CreateRetrievalVerification`
    - `ListEndpoints`
    - `CreateEndpoint`
    - `ShowEndpoint`
    - `UpdateEndpoint`
    - `DeleteEndpoint`
    - `ListEndpointIpaddresses`
    - `AssociateEndpointIpaddress`
    - `DisassociateEndpointIpaddress`
    - `ListEndpointVpcs`
    - `CreateResolveRule`
    - `AssociateResolveRuleRouter`
    - `DisassociateResolveRuleRouter`
    - `ShowResoleRule`
    - `UpdateResolveRule`
    - `DeleteResolveRule`
    - `BatchDeleteZones`
    - `BatchDeletePtrRecords`
    - `BatchSetZonesStatus`
    - `BatchSetRecordSetsStatus`
    - `BatchDeleteRecordSets`
  - 接口`CreatePrivateZone`:
    - 新增请求参数 `proxy_pattern`
    - 新增响应参数 `proxy_pattern`
  - 接口`ListPrivateZones`新增响应参数 `proxy_pattern`

### HuaweiCloud SDK DRIS

- _新增特性_
  - 支持接口`BatchShowRadars`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持接口`ShowDbObjectCollectionStatus`、`ShowUpdateObjectSavingStatus`、`CollectDbObjectsAsync`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDbObjects`:
    - 新增请求参数 `db_names`
    - 新增响应参数 `status`、`id`

### HuaweiCloud SDK DWS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListNodeTypes`新增响应参数 `count`、`datastore_type`、`available_zones`、`ram`、`vcpus`、`datastores`、`volume`、`elastic_volume_specs`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFunctionTriggers`响应参数`trigger_status`新增枚举值`DISABLED`, 移除枚举值`DISABLE`
  - 接口`UpdateTrigger`请求参数`trigger_status`新增枚举值`DISABLED`, 移除枚举值`DISABLE`
  - 接口`ShowFunctionTrigger`响应参数`trigger_status`新增枚举值`DISABLED`, 移除枚举值`DISABLE`
  - 接口`CreateWorkflow`新增请求参数 `enable_stream_response`
  - 接口`UpdateWorkFlow`新增请求参数 `enable_stream_response`
  - 接口`ShowWorkFlow`新增响应参数 `enable_stream_response`

### HuaweiCloud SDK HiLens

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowNodes`新增响应参数 `deployment_num`
  - 接口`ShowNode`新增响应参数 `cluster_node_type`
  - 接口`ShowUpgradeProgress`新增请求参数 `offset`、`limit`
  - 接口`ListTasks`新增请求参数 `offset`、`limit`
  - 接口`ShowDeploymentPods`新增请求参数 `offset`、`limit`
  - 接口`ListWorkSpaces`新增请求参数 `offset`、`limit`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEdgeNode`请求参数`app_version`改为非必填
  - 接口`ShowEdgeNode`新增响应参数 `reliability_level`

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`ListUpStreamDetail`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListBandwidthDetail`:
    - 新增请求参数 `service_type`
    - 请求参数`play_domains`改为必填
  - 接口`ListDomainTrafficDetail`:
    - 新增请求参数 `service_type`
    - 请求参数`play_domains`改为必填
  - 接口`ListDomainBandwidthPeak`:
    - 新增请求参数 `service_type`
    - 请求参数`play_domains`改为必填
  - 接口`ListDomainTrafficSummary`:
    - 新增请求参数 `service_type`
    - 请求参数`play_domains`改为必填
  - 接口`ListUsersOfStream`新增请求参数 `service_type`
  - 接口`ShowUpBandwidth`新增请求参数 `type`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateSearchCriterias`:
    - 新增请求参数 `eps_id`
    - 移除请求参数 `epsId`
  - 接口`DeleteSearchCriterias`:
    - 新增请求参数 `eps_id`
    - 移除请求参数 `epsId`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunTextModeration`新增请求参数 `white_glossary_names`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPorts`:
    - 新增请求参数 `security_groups`
    - 请求参数`fixed_ips`类型调整 `string` -> `array`

# 0.1.20 2022-12-29

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListApiGroupsV2`新增响应参数 `verified_client_certificate_enabled`
  - 接口`ShowDetailsOfApiGroupV2`新增响应参数 `verified_client_certificate_enabled`
  - 接口`UpdateApiGroupV2`新增响应参数 `verified_client_certificate_enabled`
  - 接口`ShowDetailsOfApiV2`新增响应参数 `verified_client_certificate_enabled`
  - 接口`UpdateApiV2`新增响应参数 `verified_client_certificate_enabled`
  - 接口`ListApiVersionDetailV2`新增响应参数 `verified_client_certificate_enabled`
  - 接口`UpdateDomainV2`新增请求参数 `verified_client_certificate_enabled`
  - 接口`ShowDetailsOfDomainNameCertificateV2`新增响应参数 `id`、`name`、`type`、`instance_id`、`project_id`、`create_time`、`update_time`
  - 接口`ExportApiDefinitionsV2`新增请求参数 `oas_version`
  - 接口`ListVpcChannelsV2`新增响应参数 `instance_id`
  - 接口`ShowDetailsOfVpcChannelV2`新增响应参数 `instance_id`
  - 接口`UpdateVpcChannelV2`新增响应参数 `instance_id`
  - 接口`ListMetricData`请求参数`dim`新增枚举值`inbound_eip`、`outbound_eip`, 移除枚举值`inbound`、`outbound`
  - 接口`CreateInstanceV2`新增请求参数 `vpcep_service_name`
  - 接口`ShowDetailsOfInstanceV2`新增响应参数 `is_releasable`
  - 接口`UpdateInstanceV2`:
    - 新增请求参数 `vpcep_service_name`
    - 新增响应参数 `is_releasable`
  - 接口`CreateCertificateV2`新增请求参数 `trusted_root_ca`
  - 接口`BatchAssociateCertsV2`新增请求参数 `verified_client_certificate_enabled`
  - 接口`BatchDisassociateCertsV2`新增请求参数 `verified_client_certificate_enabled`
  - 接口`UpdateCertificateV2`新增请求参数 `trusted_root_ca`
  - 接口`BatchAssociateDomainsV2`新增请求参数 `verified_client_certificate_enabled`
  - 接口`BatchDisassociateDomainsV2`新增请求参数 `verified_client_certificate_enabled`
  - 接口`ListAttachedDomainsV2`新增响应参数 `verified_client_certificate_enabled`

### HuaweiCloud SDK CFW

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDnsServers`新增请求参数 `fw_instance_id`、`enterprise_project_id`
  - 接口`UpdateDnsServers`新增请求参数 `fw_instance_id`、`enterprise_project_id`
  - 接口`ListVpcProtects`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListRuleHitCount`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`DeleteAclRuleCount`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ChangeIpsSwitchUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListIpsSwitchStatusUsingGet`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListEastWestFirewall`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ChangeEwProtectStatus`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListFlowLogs`新增请求参数 `enterprise_project_id`
  - 接口`ListAccessControlLogs`新增请求参数 `enterprise_project_id`
  - 接口`ListAttackLogs`新增请求参数 `enterprise_project_id`
  - 接口`AddRuleAclUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`DeleteRuleAclUsingDelete`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`UpdateRuleAclUsingPut`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListRuleAclsUsingGet`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListRuleAclUsingPut`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`AddBlackWhiteListUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`DeleteBlackWhiteListUsingDelete`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`UpdateBlackWhiteListUsingPut`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListBlackWhiteListsUsingGet`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ChangeIpsProtectModeUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListIpsProtectModeUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListFirewallUsingGet`:
    - 新增请求参数 `enterprise_project_id`、`fw_instance_id`
    - 新增响应参数 `fw_instance_name`、`enterprise_project_id`
  - 接口`AddServiceSetUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`DeleteServiceSetUsingDelete`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListServiceSetDetails`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`UpdateServiceSetUsingPut`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`AddServiceItemsUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListServiceItemsDetails`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`DeleteServiceItemUsingDelete`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListParseDomainDetails`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`CountEips`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ChangeProtectEip`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListEipResources`:
    - 新增请求参数 `fw_instance_id`、`fw_key_word`、`eps_id`
    - 新增响应参数 `fw_instance_name`、`fw_instance_id`、`fw_enterprise_project_id`
  - 接口`DeleteAddressItemUsingDelete`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`AddAddressItemsUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListAddressItemsUsingGet`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`AddAddressSetInfoUsingPost`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListAddressSetListUsingGet`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`DeleteAddressSetInfoUsingDelete`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListAddressSetDetailUsingGet`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`UpdateAddressSetInfoUsingPut`新增请求参数 `enterprise_project_id`、`fw_instance_id`
  - 接口`ListServiceSet`新增请求参数 `enterprise_project_id`、`fw_instance_id`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClustersDetails`响应参数`size`类型调整 `string` -> `int32`
  - 接口`ShowClusterDetail`响应参数`size`类型调整 `string` -> `int32`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateBigkeyScanTask`响应参数`size`类型调整 `int32` -> `int64`
  - 接口`ShowBigkeyScanTaskDetails`响应参数`size`类型调整 `int32` -> `int64`
  - 接口`CreateHotkeyScanTask`响应参数`size`类型调整 `int32` -> `int64`
  - 接口`ShowHotkeyTaskDetails`响应参数`size`类型调整 `int32` -> `int64`

### HuaweiCloud SDK DNS

- _新增特性_
  - 支持接口`AssociateResolveRuleRouter`、`DisassociateResolveRuleRouter`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPublicZones`移除响应参数 `total_count`
  - 接口`ListPrivateZones`移除响应参数 `total_count`
  - 接口`ListRecordSetsByZone`移除响应参数 `total_count`
  - 接口`ListRecordSets`移除响应参数 `total_count`
  - 接口`BatchDeleteRecordSetWithLine`移除响应参数 `total_count`
  - 接口`BatchUpdateRecordSetWithLine`移除响应参数 `total_count`
  - 接口`ListRecordSetsWithLine`移除响应参数 `total_count`
  - 接口`CreateRecordSetWithBatchLines`移除响应参数 `total_count`
  - 接口`ShowRecordSetByZone`移除响应参数 `total_count`
  - 接口`ListPtrRecords`移除响应参数 `total_count`
  - 接口`ListCustomLine`移除响应参数 `total_count`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateJobs`新增请求参数 `charging_mode`、`period_order`
  - 接口`BatchCheckJobs`请求参数`precheck_mode`新增枚举值`forRetryJob`
  - 接口`BatchListJobDetails`新增响应参数 `period_order`、`object_infos`

### HuaweiCloud SDK DWS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClusterDetails`新增响应参数 `elb`
  - 接口`ListAlarmSubs`新增请求参数 `offset`、`limit`
  - 接口`ListEvents`新增请求参数 `offset`、`limit`
  - 接口`ListEventSubs`新增请求参数 `offset`、`limit`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ChangeBandwidthToPeriod`新增请求参数 `extendParam`
  - 接口`ChangePublicipToPeriod`新增请求参数 `extendParam`
  - 接口`ListBandwidthPkg`:
    - 新增响应参数 `tenantId`
    - 移除响应参数 `tenant_id`
  - 接口`UpdateAssociatePublicip`请求参数`associate_instance_type`、`associate_instance_id`改为必填
  - 接口`AssociatePublicips`请求参数`associate_instance_type`、`associate_instance_id`改为必填

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListLoadbalancersByTags`:
    - 移除请求参数 `without_any_tag`
    - 请求参数`values`改为必填
  - 接口`ListListenersByTags`:
    - 移除请求参数 `without_any_tag`
    - 请求参数`values`改为必填
  - 接口`ShowQuota`新增响应参数 `ipgroup_bindings`、`ipgroup_max_length`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstanceTags`新增响应参数 `total_count`

### HuaweiCloud SDK HiLens

- _新增特性_
  - 支持以下接口：
    - `ListPlatformManager`
    - `ListFirmwares`
    - `ShowDeployments`
    - `CreateDeployment`
    - `ShowDeployment`
    - `UpdateDeployment`
    - `StartAndStopDeployment`
    - `DeletePod`
    - `StartAndStopDeploymentPod`
    - `AddDeploymentNodes`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowNodes`移除响应参数 `id`、`name`、`description`、`created_at`、`cluster_id`、`cluster_node_state`、`cluster_node_type`、`firmware_name`、`upgrade_firmware_version`、`firmware_status`、`firmware_upgrade_record`、`state`、`type`、`active_status`、`cpu`、`host_ips`、`tags`
  - 接口`CreateNode`:
    - 移除请求参数 `log_group_id`
    - 请求参数`component`、`type`改为必填
  - 接口`ShowNode`:
    - 移除响应参数 `log_group_id`
    - 响应参数`firmware_upgrade_time`类型调整 `int32` -> `string`
    - 响应参数`component`、`type`改为必填
  - 接口`UpdateNode`:
    - 移除请求参数 `log_group_id`
    - 请求参数`component`、`type`改为必填
  - 接口`ShowUpgradeProgress`响应参数`status`类型调整 `string` -> `int32`
  - 接口`ShowResourceTags`:
    - 新增响应参数 `count`
    - 移除响应参数 `total`
  - 接口`ListSecrets`响应参数`count`类型调整 `string` -> `int32`
  - 接口`CreateSecret`:
    - 新增响应参数 `secret`
    - 移除响应参数 `id`、`name`、`description`、`project_id`、`created_at`、`updated_at`、`secrets`、`tags`
  - 接口`UpdateSecret`:
    - 新增响应参数 `secret`
    - 移除响应参数 `id`、`name`、`description`、`project_id`、`created_at`、`updated_at`、`secrets`、`tags`
  - 接口`ShowSkillOrderList`:
    - 响应参数`update_time`类型调整 `int32` -> `int64`
    - 响应参数`expire_time`类型调整 `int32` -> `int64`
    - 响应参数`order_limit`类型调整 `string` -> `int32`
  - 接口`CreateOrderForm`:
    - 新增响应参数 `order_id`
    - 移除请求参数 `data`、`total`
    - 移除响应参数 `data`、`total`
  - 接口`ShowSkillOrderInfo`:
    - 新增响应参数 `expiration_stop_flag`、`package_order_id`、`icon`、`commission_flag`、`product_info`、`package_id`、`measure_type`、`update_time`、`channel_limit`、`resource_step_size`、`cloud_service_type`、`developer_id`、`amount`、`format`、`resource_type`、`expire_time`、`measure_unit`、`skill_chip`、`versions`、`skill_name`、`skill_type`、`used_amount`、`charge_model`、`resource_spec_code`、`skill_id`、`skill_platform`、`order_limit`、`order_id`、`status`
    - 移除响应参数 `total`、`data`
  - 接口`DeleteDeployment`新增响应参数 `deployment_id`
  - 接口`UpdateDeploymentUsingPatch`移除响应参数 `deployment_tags`
  - 接口`ShowDeploymentPods`:
    - 新增响应参数 `start_resources`、`channel_resources`、`skill_project_id`
    - 移除响应参数 `app_docker_login`、`app_id`、`expire_time`、`image_url`、`license`、`modelKey`
    - 响应参数`host_network`、`restart_policy`、`app_url`、`name`改为必填
  - 接口`CreateWorkSpace`:
    - 新增响应参数 `workspace_id`
    - 移除响应参数 `name`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJob`新增响应参数 `sub_jobs_result`、`sub_jobs_list`
  - 接口`ShowJobProgress`新增响应参数 `sub_jobs_result`、`sub_jobs_list`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchListEdgeApps`新增响应参数 `protocol`、`edge_app_name`
  - 接口`CreateEdgeApp`新增请求参数 `edge_app_name`、`protocol`
  - 接口`ShowEdgeApp`新增响应参数 `protocol`、`edge_app_name`
  - 接口`BatchListEdgeAppVersions`新增响应参数 `name`
  - 接口`CreateEdgeApplicationVersion`新增请求参数 `supplier`、`tpl_id`
  - 接口`ShowEdgeApplicationVersion`新增响应参数 `supplier`、`tpl_id`
  - 接口`UpdateEdgeApplicationVersion`:
    - 新增请求参数 `tpl_id`
    - 新增响应参数 `name`
  - 接口`UpdateEdgeApplicationVersionState`新增响应参数 `name`
  - 接口`CreateEdgeNode`新增请求参数 `os_type`、`reliability_level`、`network_access_point`、`offline_cache_configs`、`device_auth_info`
  - 接口`ShowEdgeNode`新增响应参数 `offline_cache_configs`、`device_auth_info`

### HuaweiCloud SDK MapDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCredential`:
    - 新增请求参数 `credential`
    - 移除请求参数 `description`

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateTranscodeTemplate`请求参数`name`改为非必填
  - 接口`UpdateTemplateGroupCollection`:
    - 请求参数`collection_id`改为必填
    - 请求参数`name`、`template_group_list`改为非必填

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListVpcs`新增响应参数 `tenant_id`、`created_at`、`updated_at`
  - 接口`CreateVpc`新增响应参数 `tenant_id`、`created_at`、`updated_at`
  - 接口`ShowVpc`新增响应参数 `tenant_id`、`created_at`、`updated_at`
  - 接口`UpdateVpc`新增响应参数 `tenant_id`、`created_at`、`updated_at`
  - 接口`ListSubnets`新增响应参数 `tenant_id`、`created_at`、`updated_at`
  - 接口`CreateSubnet`新增响应参数 `tenant_id`、`created_at`、`updated_at`
  - 接口`ShowSubnet`新增响应参数 `tenant_id`、`created_at`、`updated_at`
  - 接口`ListRouteTables`新增响应参数 `created_at`、`updated_at`
  - 接口`CreateRouteTable`新增响应参数 `created_at`、`updated_at`
  - 接口`ShowRouteTable`新增响应参数 `created_at`、`updated_at`
  - 接口`UpdateRouteTable`新增响应参数 `created_at`、`updated_at`
  - 接口`AssociateRouteTable`新增响应参数 `created_at`、`updated_at`
  - 接口`DisassociateRouteTable`新增响应参数 `created_at`、`updated_at`

# 0.1.19 2022-12-26

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateWorkflow`:
    - 新增请求参数 `policy`
    - 移除请求参数 `citation_urns`、`information`、`alarm_name`
    - 移除响应参数 `param_name`、`domain_id`、`domain_name`
  - 接口`UpdateWorkflowTriggerStatus`请求参数`action`改为非必填
  - 接口`SearchWorkflowExecutionDetail`移除响应参数 `workflow_urn`、`headers`、`input`、`output`、`created_by`、`node_id`、`begin_time`、`end_time`、`function_execution_id`、`output`、`status`
  - 接口`ListAllScriptByName`:
    - 移除请求参数 `page_total`
    - 请求参数`order_by_column`改为必填
    - 请求参数`project_id`改为非必填
  - 接口`ListAllVersionByVersionId`:
    - 移除请求参数 `page_total`
    - 请求参数`order_by_column`改为必填
    - 请求参数`project_id`改为非必填
  - 接口`ListAllJobByName`:
    - 新增请求参数 `enterprise_project_id`
    - 新增响应参数 `is_latest_version`、`version_number`
    - 请求参数`order_by_column`改为必填
  - 接口`ListTemplateByJobId`请求参数`order_by_column`改为必填
  - 接口`SearchTemplateById`请求参数`share_type`改为非必填
  - 接口`ListWorkflow`:
    - 移除请求参数 `template_name_list`
    - 移除响应参数 `param_name`、`domain_id`、`domain_name`
  - 接口`ExecuteWorkflow`移除响应参数 `workflow_id`、`workflow_urn`、`status`、`headers`、`input`、`output`、`begin_time`、`end_time`、`last_update_time`、`created_by`、`execution_result_list`、`approve_user_name_list`、`project_id`、`workflow_edit_time`、`last_record_id_with_snapshot`
  - 接口`ListWorkflowExecutions`移除响应参数 `workflow_urn`、`node_id`、`begin_time`、`end_time`、`function_execution_id`、`output`、`status`
  - 接口`ListNotifiedHistories`请求参数`event_sn`改为必填
  - 接口`ShowActionRule`响应参数`type`类型调整 `string` -> `enum`
  - 接口`AddActionRule`请求参数`type`类型调整 `string` -> `enum`
  - 接口`UpdateActionRule`请求参数`type`类型调整 `string` -> `enum`
  - 接口`ListActionRule`响应参数`type`类型调整 `string` -> `enum`
  - 接口`AddEvent2alarmRule`请求参数`resource_provider`改为非必填
  - 接口`UpdateEventRule`请求参数`resource_provider`改为非必填
  - 接口`ListEvent2alarmRule`响应参数`resource_provider`改为非必填
  - 接口`CreateApp`请求参数`register_type`类型调整 `string` -> `enum`
  - 接口`UpdateApp`:
    - 移除响应参数 `aom_id`、`app_id`、`create_time`、`creator`、`description`、`display_name`、`eps_id`、`modified_time`、`modifier`、`name`、`register_type`
    - 请求参数`register_type`类型调整 `string` -> `enum`
  - 接口`UpdateComponent`移除请求参数 `model_id`、`model_type`
  - 接口`CreateEnv`:
    - 请求参数`component_id`、`os_type`改为必填
    - 请求参数`region`改为非必填
  - 接口`UpdateEnv`:
    - 请求参数`component_id`、`os_type`改为必填
    - 请求参数`region`改为非必填
  - 接口`ShowComponentByName`移除响应参数 `create_time`、`creator`、`description`、`modified_time`、`modifier`、`register_type`、`sub_app_id`

### HuaweiCloud SDK CBH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCbhInstance`:
    - 新增响应参数 `quotaDetail`、`publicip`、`expTime`、`startTime`、`endTime`、`releaseTime`、`instanceId`、`privateIp`、`taskStatus`、`vpcId`、`subnetId`、`securityGroupId`、`createinstanceStatus`、`failReason`、`instanceKey`、`orderId`、`periodNum`、`resourceId`、`alterPermit`、`publicId`、`bastionVersion`、`newBastionVersion`、`instanceStatus`、`instanceDescription`、`slaveZone`、`enterpriseProjectId`、`instanceType`、`haId`、`slaveZoneDisplay`、`webPort`、`vip`
    - 移除响应参数 `quota_detail`、`public_ip`、`exp_time`、`start_time`、`end_time`、`release_time`、`instance_id`、`private_ip`、`task_status`、`vpc_id`、`subnet_id`、`security_group_id`、`createinstance_status`、`fail_reason`、`instance_key`、`order_id`、`period_num`、`resource_id`、`alter_permit`、`public_id`、`bastion_version`、`new_bastion_version`、`instance_status`、`instance_description`
    - 响应参数`is_auto_renew`改为非必填
  - 接口`UpgradeCbhInstance`响应参数`task_id`、`order_id`改为非必填
  - 接口`ResetPassword`:
    - 新增响应参数 `request_info`
    - 移除响应参数 `code`、`description`、`task_id`、`order_id`
  - 接口`ShowAvailableZoneInfo`:
    - 新增响应参数 `availability_zone`
    - 移除响应参数 `availability_zones`
  - 接口`ResetLoginMethod`:
    - 新增响应参数 `request_info`
    - 移除响应参数 `code`、`description`、`task_id`、`order_id`
  - 接口`ChangeInstanceNetwork`:
    - 新增响应参数 `status`、`security_grp_status`、`public_eip_status`、`nics`
    - 移除响应参数 `type`、`security_groups`
    - 响应参数`firewall_status`改为必填
    - 响应参数`public_eip_statu`改为非必填
  - 接口`CreateInstance`请求参数`slave_availability_zone`改为非必填
  - 接口`ChangeInstanceOrder`:
    - 新增响应参数 `orderId`
    - 移除响应参数 `order_id`

### HuaweiCloud SDK DSC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowSpecification`:
    - 新增响应参数 `orderInfos`
    - 移除响应参数 `order_infos`
  - 接口`ListRuleGroups`新增响应参数 `is_default`

### HuaweiCloud SDK HSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPorts`新增响应参数 `laddr`
  - 接口`ListProtectionServer`新增响应参数 `os_name`、`agent_status`
  - 接口`ShowBackupPolicyInfo`移除响应参数 `associated_vaults`
  - 接口`ListSecurityEvents`新增响应参数 `os_type`、`event_details`
  - 接口`ListAlarmWhiteList`新增响应参数 `enterprise_project_name`
  - 接口`ListVulHosts`移除响应参数 `repair_necessity`

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`UpdateDomainIp6Switch`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowDomain`:
    - 新增请求参数 `enterprise_project_id`
    - 新增响应参数 `enterprise_project_id`、`is_ipv6`
  - 接口`UpdateDomain`:
    - 新增请求参数 `enterprise_project_id`
    - 新增响应参数 `enterprise_project_id`
  - 接口`CreateDomain`新增请求参数 `enterprise_project_id`

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持以下接口：
    - `Deletefavorite`
    - `ListTopnTrafficStatistics`
    - `Createfavorite`
    - `CreateTags`
    - `CreateDashboardGroup`
    - `CreateDashBoard`
    - `ListHistorySql`
    - `ListCriterias`
    - `CreateSearchCriterias`
    - `DeleteSearchCriterias`
    - `ListQueryAllSearchCriterias`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateLogStream`移除请求参数 `enterprise_project_name`、`ttl_in_days`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTranscodingTask`:
    - 移除响应参数 `ref_frames_count`
    - 响应参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`CreateTranscodingTask`:
    - 移除请求参数 `ref_frames_count`、`aspect_ratio`
    - 请求参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`ListTemplate`:
    - 移除响应参数 `ref_frames_count`
    - 响应参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`CreateTransTemplate`:
    - 移除请求参数 `ref_frames_count`
    - 请求参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`UpdateTransTemplate`:
    - 移除请求参数 `ref_frames_count`
    - 请求参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`ListTemplateGroup`:
    - 移除响应参数 `ref_frames_count`、`aspect_ratio`
    - 响应参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`CreateTemplateGroup`:
    - 移除请求参数 `ref_frames_count`、`aspect_ratio`
    - 请求参数`name`改为必填
    - 请求参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`UpdateTemplateGroup`:
    - 移除请求参数 `ref_frames_count`、`aspect_ratio`
    - 请求参数`group_id`改为必填
    - 请求参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`UpdateBucketAuthorized`移除请求参数 `project_id`
  - 接口`CreateThumbnailsTask`移除请求参数 `aspect_ratio`
  - 接口`ListEditingJob`:
    - 移除响应参数 `ref_frames_count`、`ref_frames_count`
    - 响应参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`、`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填
  - 接口`CreateEditingJob`:
    - 移除请求参数 `ref_frames_count`、`ref_frames_count`
    - 请求参数`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`、`codec`、`sample_rate`、`PVC`、`hls_interval`、`dash_interval`改为非必填

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateTranscodeTemplate`请求参数`group_id`、`name`、`bitrate`、`frame_rate`、`video_codec`、`format`、`hls_interval`改为必填
  - 接口`ListTranscodeTemplate`响应参数`bitrate`、`frame_rate`、`video_codec`、`format`、`hls_interval`改为必填
  - 接口`CreateTranscodeTemplate`请求参数`name`、`bitrate`、`frame_rate`、`video_codec`、`format`、`hls_interval`改为必填
  - 接口`UpdateTemplateGroupCollection`请求参数`name`、`template_group_list`改为必填
  - 接口`CreateTemplateGroupCollection`请求参数`name`、`template_group_list`改为必填

# 0.1.18 2022-12-22

### HuaweiCloud SDK APIG

- _新增特性_
  - 支持接口`ListProjectInstanceTags`、`ListInstanceTags`、`BatchCreateOrDeleteInstanceTags`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持以下接口：
    - `ExpandInstanceStorage`
    - `ListClusterScaleInNumbers`
    - `ListDisasterRecover`
    - `CreateDisasterRecovery`
    - `DeleteDataSource`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShrinkCluster`:
    - 新增请求参数 `clusterShrinkReq`
    - 新增响应参数 `job_id`
    - 移除请求参数 `shrink_number`、`online`、`type`、`retry`、`force_backup`、`need_agency`
  - 接口`ExecuteRedistributionCluster`:
    - 新增请求参数 `redistributionReq`
    - 移除请求参数 `redis_mode`、`parallel_jobs`
  - 接口`CreateClusterWorkload`:
    - 新增请求参数 `workload_status`
    - 新增响应参数 `workload_res_code`、`workload_res_str`
    - 移除请求参数 `workload_switch`、`max_concurrency_num`
  - 接口`ListClusterWorkload`:
    - 新增响应参数 `workload_res_code`、`workload_res_str`
    - 响应参数`workload_switch`改为必填
  - 接口`CreateWorkloadPlan`:
    - 新增请求参数 `workloadPlan`
    - 新增响应参数 `workload_res_code`、`workload_res_str`
    - 移除请求参数 `plan_id`、`plan_name`、`logical_cluster_name`
  - 接口`AddWorkloadQueue`:
    - 新增请求参数 `workload_queue`
    - 新增响应参数 `workload_res_code`、`workload_res_str`
    - 移除请求参数 `workload_queue_name`、`logical_cluster_name`、`short_query_optimize`、`short_query_concurrency_num`、`workload_resource_item_list`
  - 接口`ListWorkloadQueue`新增响应参数 `workload_res_code`、`workload_res_str`
  - 接口`DeleteWorkloadQueue`:
    - 新增响应参数 `workload_res_code`、`workload_res_str`
    - 请求参数`logical_cluster_name`改为必填
  - 接口`CopySnapshot`:
    - 新增请求参数 `linkCopyReq`
    - 新增响应参数 `snapshot_id`
    - 移除请求参数 `backup_name`、`description`
  - 接口`ListAuditLog`移除响应参数 `version`、`configure_status`
  - 接口`CreateDataSource`:
    - 新增请求参数 `extDataSourceReq`
    - 新增响应参数 `id`、`job_id`
    - 移除请求参数 `data_source_id`、`type`、`data_source_name`、`user_name`、`user_pwd`、`description`、`reboot`、`connect_info`
  - 接口`UpdateDataSource`:
    - 新增请求参数 `reconfigure`
    - 新增响应参数 `job_id`
    - 移除请求参数 `database`、`agency`
  - 接口`ListSnapshotDetails`新增响应参数 `datastore`、`cluster_name`、`bak_expected_start_time`、`bak_keep_day`、`bak_period`、`db_user`、`progress`、`backup_key`、`prior_backup_key`、`base_backup_key`、`backup_device`、`total_backup_size`、`base_backup_name`、`support_inplace_restore`、`fine_grained_backup`、`backup_level`、`fine_grained_backup_detail`、`guest_agent_version`、`cluster_status`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowWorkflowExecutionForPage`:
    - 新增响应参数 `created_by`
    - 移除响应参数 `create_by`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持接口`ModifyVolume`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAvailableFlavorInfos`新增请求参数 `offset`、`limit`

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTag`请求参数`tag_name`改为必填
  - 接口`ListProPricePlans`移除请求参数 `sim_card_id`、`partner`、`package_type`、`sim_type`
  - 接口`ListSimCards`:
    - 移除请求参数 `expire_time_duration`、`traffic_warning_threshold`、`sim_status_in`
    - 移除响应参数 `sn`、`supply_code`、`bundle_id`、`test_type`
  - 接口`StopSimCard`:
    - 移除请求参数 `price_plan_list`
    - 移除响应参数 `sim_price_plan_list`
  - 接口`ResetSimCard`:
    - 移除请求参数 `price_plan_list`
    - 移除响应参数 `sim_price_plan_list`
  - 接口`ShowSimCard`移除响应参数 `sn`、`supply_code`、`bundle_id`、`test_type`
  - 接口`ListSimPricePlans`:
    - 移除请求参数 `sim_price_plan_id`
    - 移除响应参数 `partner`、`partner_pid`

### HuaweiCloud SDK IEF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEdgeNodes`新增响应参数 `purchase_id`、`state_details`、`cert_remaining_valid_time`
  - 接口`UpdateEdgeNode`新增响应参数 `purchase_id`、`state_details`、`cert_remaining_valid_time`
  - 接口`ShowEdgeNodeDetail`新增响应参数 `purchase_id`、`state_details`、`cert_remaining_valid_time`
  - 接口`ListEdgeGroups`新增响应参数 `purchase_id`、`state_details`、`cert_remaining_valid_time`
  - 接口`UpdateEdgeGroup`新增响应参数 `purchase_id`、`state_details`、`cert_remaining_valid_time`
  - 接口`ShowEdgeGroupDetail`新增响应参数 `purchase_id`、`state_details`、`cert_remaining_valid_time`
  - 接口`UpdateEdgeGroupNodeBinding`新增响应参数 `purchase_id`、`state_details`、`cert_remaining_valid_time`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CopyImageCrossRegion`新增请求参数 `vault_id`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 支持接口`SearchDevices`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRule`新增请求参数 `device_side`
  - 接口`ListRules`新增响应参数 `device_side`
  - 接口`ShowRule`新增响应参数 `device_side`
  - 接口`UpdateRule`:
    - 新增请求参数 `device_side`
    - 新增响应参数 `device_side`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePostPaidInstance`新增请求参数 `sasl_enabled_mechanisms`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePostPaidInstance`新增请求参数 `enterprise_project_id`、`enable_acl`

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateShare`:
    - 新增请求参数 `CreateShareRequestBody`
    - 移除请求参数 `share`
  - 接口`ListShares`:
    - 请求参数`offset`类型调整 `int32` -> `int64`
    - 请求参数`limit`类型调整 `int32` -> `int64`
  - 接口`ExpandShare`:
    - 新增请求参数 `ExpandShareRequestBody`
    - 移除请求参数 `extend`
  - 接口`CreateSharedTag`:
    - 新增请求参数 `CreateSharedTagRequestBody`
    - 移除请求参数 `tag`
  - 接口`BatchAddSharedTags`:
    - 新增请求参数 `BatchAddSharedTagsRequestBody`
    - 移除请求参数 `add_shareed_tags`

# 0.1.17 2022-12-19

### HuaweiCloud SDK APM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowMonitorItemViewConfig`响应参数`function`改为必填
  - 接口`ShowTrend`请求参数`view_type`、`collector_name`、`metric_set`、`function`、`env_id`、`start_time`、`end_time`改为必填
  - 接口`ShowSumTable`请求参数`view_type`、`collector_name`、`metric_set`、`function`、`page`、`page_size`、`env_id`、`start_time`、`end_time`改为必填
  - 接口`ShowRawTable`:
    - 新增请求参数 `last_row_id`
    - 移除请求参数 `lastRowId`
    - 请求参数`function`改为必填
  - 接口`SearchAgent`:
    - 新增请求参数 `order_by_status`
    - 移除请求参数 `orderByStatus`

### HuaweiCloud SDK BCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowBlockchainDetail`新增响应参数 `order_fade_enabled`、`is_support_tc3`、`order_fade_cache`、`deploy_status`、`block_info`、`cluster_platform_type`、`status`、`status_detail`、`order_fade_enabled`
  - 接口`DeleteMemberInvite`新增响应参数 `result`
  - 接口`HandleNotification`新增响应参数 `result`
  - 接口`CreateNewBlockchain`:
    - 请求参数`node_flavor`类型调整 `string` -> `int64`
    - 请求参数`cce_flavor`类型调整 `string` -> `int64`
    - 请求参数`init_node_pwd`类型调整 `string` -> `int64`
    - 请求参数`az`类型调整 `string` -> `int64`
    - 请求参数`cluster_platform_type`类型调整 `string` -> `int64`
  - 接口`DownloadBlockchainCert`新增响应参数 `result`
  - 接口`DownloadBlockchainSdkConfig`新增响应参数 `result`
  - 接口`ListEntityMetric`新增响应参数 `filesystemUsage`
  - 接口`CreateBlockchainCertByUserName`新增响应参数 `result`

### HuaweiCloud SDK CBH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstanceOrder`请求参数`available_zone_id`改为非必填

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateVault`:
    - 请求参数`object_type`新增枚举值`workspace`
    - 响应参数`object_type`新增枚举值`workspace`
  - 接口`ListVault`响应参数`object_type`新增枚举值`workspace`
  - 接口`ShowVault`响应参数`object_type`新增枚举值`workspace`
  - 接口`UpdateVault`响应参数`object_type`新增枚举值`workspace`
  - 接口`ShowBackup`响应参数`resource_type`新增枚举值`OS::Workspace::DesktopV2`
  - 接口`ListBackups`:
    - 请求参数`resource_type`新增枚举值`OS::Workspace::DesktopV2`
    - 响应参数`resource_type`新增枚举值`OS::Workspace::DesktopV2`
  - 接口`ListProtectable`响应参数`object_type`新增枚举值`workspace`
  - 接口`ShowProtectable`响应参数`object_type`新增枚举值`workspace`
  - 接口`ShowVaultResourceInstances`响应参数`object_type`新增枚举值`workspace`

### HuaweiCloud SDK CC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateBandwidthPackage`新增请求参数 `interflow_mode`
  - 接口`ListInterRegionBandwidths`:
    - 新增响应参数 `inter_region_bandwidths`
    - 移除响应参数 `inter_region_bandwidth`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`:
    - 新增请求参数 `vpcPermissions`
    - 新增响应参数 `orderId`
    - 移除请求参数 `vpcPermisssions`
    - 移除响应参数 `ordeId`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListServersDetails`新增请求参数 `server_id`

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持接口`ShowResourcesJobDetail`、`ChangeBandwidthToPeriod`、`ChangePublicipToPeriod`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateFunction`:
    - 新增请求参数 `depend_version_list`、`func_vpc`
    - 新增响应参数 `depend_version_list`
  - 接口`UpdateFunctionCode`:
    - 新增请求参数 `depend_version_list`
    - 新增响应参数 `depend_version_list`
  - 接口`ShowFunctionCode`新增响应参数 `depend_version_list`
  - 接口`ShowFunctionConfig`新增响应参数 `depend_version_list`
  - 接口`ListReservedInstanceConfigs`:
    - 新增请求参数 `marker`、`limit`
    - 新增响应参数 `reserved_instances`
    - 移除响应参数 `reservedinstances`
  - 接口`ImportFunction`新增响应参数 `depend_version_list`
  - 接口`ListFunctionReservedInstances`:
    - 新增请求参数 `limit`
    - 移除请求参数 `maxitems`
  - 接口`ShowWorkflowExecutionForPage`:
    - 新增请求参数 `offset`、`limit`、`start_time`、`end_time`
    - 移除请求参数 `CreateWorkflowRequestBody`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SetGaussMySqlProxyWeight`移除请求参数 `id`、`weight`
  - 接口`ShowGaussMySqlJobInfo`响应参数`status`新增枚举值`Pending`
  - 接口`ListScheduleJobs`:
    - 新增响应参数 `job_status`
    - 移除响应参数 `task_status`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListConfigurationDatastores`:
    - 新增响应参数 `datastore_name`
    - 移除响应参数 `datastore_type`
  - 接口`ModifyEpsQuotas`移除请求参数 `instance`、`vcpus`、`ram`
  - 接口`ListEpsQuotas`移除响应参数 `instance`、`vcpus`、`ram`、`instance`、`vcpus`、`ram`

### HuaweiCloud SDK HSS

- _新增特性_
  - 支持以下接口：
    - `ShowAssetStatistic`
    - `ListUserStatistics`
    - `ListPortStatistics`
    - `ListProcessStatistics`
    - `ListAppStatistics`
    - `ListAutoLaunchStatistics`
    - `ListPorts`
    - `ListApps`
    - `ListAutoLaunchs`
    - `ListAppChangeHistories`
    - `ListAutoLaunchChangeHistories`
    - `ListProtectionServer`
    - `ListProtectionPolicy`
    - `UpdateProtectionPolicy`
    - `StartProtection`
    - `StopProtection`
    - `ShowBackupPolicyInfo`
    - `UpdateBackupPolicyInfo`
    - `ChangeEvent`
    - `ListAlarmWhiteList`
    - `ListHostGroups`
    - `ChangeHostsGroup`
    - `AddHostsGroup`
    - `DeleteHostsGroup`
    - `ListPolicyGroup`
    - `AssociatePolicyGroup`
    - `ListVulHosts`
    - `ChangeVulStatus`
    - `ListWtpProtectHost`
    - `SetWtpProtectionStatusInfo`
    - `SetRaspSwitch`
    - `ListHostProtectHistoryInfo`
    - `ListHostRaspProtectHistoryInfo`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRiskConfigCheckRules`新增响应参数 `fix_status`、`enable_auto_fix`、`rule_params`
  - 接口`ListSecurityEvents`新增响应参数 `extend_info`
  - 接口`ListHostStatus`:
    - 新增响应参数 `enterprise_project_id`、`agent_create_time`、`agent_update_time`、`agent_version`、`upgrade_status`、`upgrade_result_code`、`upgradable`
    - 请求参数`region`改为必填
  - 接口`ListVulnerabilities`请求参数`vul_id`改为必填

### HuaweiCloud SDK IEF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTags`响应参数`key`、`value`改为非必填
  - 接口`CreateTag`请求参数`key`、`value`改为非必填
  - 接口`ListEdgeNodes`响应参数`key`、`value`改为非必填
  - 接口`ShowEdgeNodeDetail`响应参数`key`、`value`改为非必填
  - 接口`UpdateEdgeNode`响应参数`key`、`value`改为非必填
  - 接口`CreateEdgeGroup`新增请求参数 `device_ids`
  - 接口`ListEdgeGroupCerts`:
    - 新增响应参数 `groupcerts`
    - 移除响应参数 `edge_groups`
  - 接口`ListDevices`响应参数`type`改为非必填
  - 接口`CreateDevice`请求参数`type`改为非必填
  - 接口`ShowDevice`响应参数`type`改为非必填
  - 接口`UpdateDevice`响应参数`type`改为非必填
  - 接口`ShowDeviceTwin`响应参数`type`改为非必填
  - 接口`UpdateDeviceTwin`:
    - 请求参数`twin`、`property_visitors`改为非必填
    - 响应参数`type`改为非必填
  - 接口`ListDeviceTemplates`响应参数`key`、`value`改为非必填
  - 接口`CreateDeviceTemplate`请求参数`key`、`value`改为非必填
  - 接口`ShowDeviceTemplate`响应参数`key`、`value`改为非必填
  - 接口`UpdateDeviceTemplateById`响应参数`key`、`value`改为非必填
  - 接口`ListResourceByTags`响应参数`key`、`value`改为非必填
  - 接口`BatchAddDeleteTags`请求参数`key`、`value`改为非必填
  - 接口`ListApps`响应参数`read_only`改为非必填
  - 接口`ShowAppDetail`响应参数`read_only`改为非必填
  - 接口`UpdateApp`响应参数`read_only`改为非必填
  - 接口`ListAppVersions`响应参数`read_only`改为非必填
  - 接口`CreateAppVersions`请求参数`read_only`改为非必填
  - 接口`ShowAppVersionDetail`响应参数`read_only`改为非必填
  - 接口`UpdateAppVersion`:
    - 请求参数`read_only`改为非必填
    - 响应参数`read_only`改为非必填
  - 接口`ListDeployments`响应参数`host_network`、`read_only`改为非必填
  - 接口`CreateDeployments`请求参数`host_network`、`read_only`改为非必填
  - 接口`ShowDeployment`响应参数`host_network`、`read_only`改为非必填
  - 接口`UpdateDeployment`:
    - 请求参数`replicas`、`host_network`、`read_only`改为非必填
    - 响应参数`host_network`、`read_only`改为非必填
  - 接口`ListPods`响应参数`host_network`、`read_only`改为非必填
  - 接口`CreateEncryptdatas`请求参数`description`改为非必填
  - 接口`UpdateEncryptdatas`请求参数`description`改为非必填
  - 接口`ListBatchJob`:
    - 新增响应参数 `task_total_count`、`task_success_count`、`task_failed_count`、`status_last_updated_at`、`description`
    - 移除响应参数 `task_count`、`success_count`、`failed_count`、`updated_at`
  - 接口`ShowBatchJob`:
    - 新增响应参数 `status_last_updated_at`
    - 移除响应参数 `updated_at`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRecordCallbackConfig`新增请求参数 `key`
  - 接口`UpdateRecordCallbackConfig`:
    - 新增请求参数 `key`
    - 新增响应参数 `id`、`publish_domain`、`app`、`notify_callback_url`、`notify_event_subscription`、`sign_type`、`create_time`、`update_time`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateSqlAlarmRule`新增请求参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`CreateSqlAlarmRule`新增请求参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`ListSqlAlarmRules`新增响应参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`UpdateKeywordsAlarmRule`新增请求参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`CreateKeywordsAlarmRule`新增请求参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`ListKeywordsAlarmRules`新增响应参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`ListHost`请求参数`host_id_list`改为非必填
  - 接口`UpdateStructConfig`请求参数`is_analysis`、`is_analysis`改为非必填
  - 接口`CreateStructConfig`请求参数`is_analysis`、`is_analysis`改为非必填

### HuaweiCloud SDK OSM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAuthorizations`新增响应参数 `resource_type_id`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`StartInstanceRestartAction`请求参数`restart`改为必填

### HuaweiCloud SDK Workspace

- _新增特性_
  - 支持接口`ShowQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.16 2022-12-15

### HuaweiCloud SDK APM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowMonitorItemViewConfig`响应参数`function`改为必填
  - 接口`ShowTrend`请求参数`view_type`、`collector_name`、`metric_set`、`function`、`env_id`、`start_time`、`end_time`改为必填
  - 接口`ShowSumTable`请求参数`view_type`、`collector_name`、`metric_set`、`function`、`page`、`page_size`、`env_id`、`start_time`、`end_time`改为必填
  - 接口`ShowRawTable`:
    - 新增请求参数 `last_row_id`
    - 移除请求参数 `lastRowId`
    - 请求参数`function`改为必填
  - 接口`SearchAgent`:
    - 新增请求参数 `order_by_status`
    - 移除请求参数 `orderByStatus`

### HuaweiCloud SDK BCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowBlockchainDetail`新增响应参数 `order_fade_enabled`、`is_support_tc3`、`order_fade_cache`、`deploy_status`、`block_info`、`cluster_platform_type`、`status`、`status_detail`、`order_fade_enabled`
  - 接口`DeleteMemberInvite`新增响应参数 `result`
  - 接口`HandleNotification`新增响应参数 `result`
  - 接口`CreateNewBlockchain`:
    - 请求参数`node_flavor`类型调整 `string` -> `int64`
    - 请求参数`cce_flavor`类型调整 `string` -> `int64`
    - 请求参数`init_node_pwd`类型调整 `string` -> `int64`
    - 请求参数`az`类型调整 `string` -> `int64`
    - 请求参数`cluster_platform_type`类型调整 `string` -> `int64`
  - 接口`DownloadBlockchainCert`新增响应参数 `result`
  - 接口`DownloadBlockchainSdkConfig`新增响应参数 `result`
  - 接口`ListEntityMetric`新增响应参数 `filesystemUsage`
  - 接口`CreateBlockchainCertByUserName`新增响应参数 `result`

### HuaweiCloud SDK CBH

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstanceOrder`请求参数`available_zone_id`改为非必填

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateVault`:
    - 请求参数`object_type`新增枚举值`workspace`
    - 响应参数`object_type`新增枚举值`workspace`
  - 接口`ListVault`响应参数`object_type`新增枚举值`workspace`
  - 接口`ShowVault`响应参数`object_type`新增枚举值`workspace`
  - 接口`UpdateVault`响应参数`object_type`新增枚举值`workspace`
  - 接口`ShowBackup`响应参数`resource_type`新增枚举值`OS::Workspace::DesktopV2`
  - 接口`ListBackups`:
    - 请求参数`resource_type`新增枚举值`OS::Workspace::DesktopV2`
    - 响应参数`resource_type`新增枚举值`OS::Workspace::DesktopV2`
  - 接口`ListProtectable`响应参数`object_type`新增枚举值`workspace`
  - 接口`ShowProtectable`响应参数`object_type`新增枚举值`workspace`
  - 接口`ShowVaultResourceInstances`响应参数`object_type`新增枚举值`workspace`

### HuaweiCloud SDK CC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateBandwidthPackage`新增请求参数 `interflow_mode`
  - 接口`ListInterRegionBandwidths`:
    - 新增响应参数 `inter_region_bandwidths`
    - 移除响应参数 `inter_region_bandwidth`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`:
    - 新增请求参数 `vpcPermissions`
    - 新增响应参数 `orderId`
    - 移除请求参数 `vpcPermisssions`
    - 移除响应参数 `ordeId`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListServersDetails`新增请求参数 `server_id`

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持接口`ShowResourcesJobDetail`、`ChangeBandwidthToPeriod`、`ChangePublicipToPeriod`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateFunction`:
    - 新增请求参数 `depend_version_list`、`func_vpc`
    - 新增响应参数 `depend_version_list`
  - 接口`UpdateFunctionCode`:
    - 新增请求参数 `depend_version_list`
    - 新增响应参数 `depend_version_list`
  - 接口`ShowFunctionCode`新增响应参数 `depend_version_list`
  - 接口`ShowFunctionConfig`新增响应参数 `depend_version_list`
  - 接口`ListReservedInstanceConfigs`:
    - 新增请求参数 `marker`、`limit`
    - 新增响应参数 `reserved_instances`
    - 移除响应参数 `reservedinstances`
  - 接口`ImportFunction`新增响应参数 `depend_version_list`
  - 接口`ListFunctionReservedInstances`:
    - 新增请求参数 `limit`
    - 移除请求参数 `maxitems`
  - 接口`ShowWorkflowExecutionForPage`:
    - 新增请求参数 `offset`、`limit`、`start_time`、`end_time`
    - 移除请求参数 `CreateWorkflowRequestBody`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SetGaussMySqlProxyWeight`移除请求参数 `id`、`weight`
  - 接口`ShowGaussMySqlJobInfo`响应参数`status`新增枚举值`Pending`
  - 接口`ListScheduleJobs`:
    - 新增响应参数 `job_status`
    - 移除响应参数 `task_status`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListConfigurationDatastores`:
    - 新增响应参数 `datastore_name`
    - 移除响应参数 `datastore_type`
  - 接口`ModifyEpsQuotas`移除请求参数 `instance`、`vcpus`、`ram`
  - 接口`ListEpsQuotas`移除响应参数 `instance`、`vcpus`、`ram`、`instance`、`vcpus`、`ram`

### HuaweiCloud SDK HSS

- _新增特性_
  - 支持以下接口：
    - `ShowAssetStatistic`
    - `ListUserStatistics`
    - `ListPortStatistics`
    - `ListProcessStatistics`
    - `ListAppStatistics`
    - `ListAutoLaunchStatistics`
    - `ListPorts`
    - `ListApps`
    - `ListAutoLaunchs`
    - `ListAppChangeHistories`
    - `ListAutoLaunchChangeHistories`
    - `ListProtectionServer`
    - `ListProtectionPolicy`
    - `UpdateProtectionPolicy`
    - `StartProtection`
    - `StopProtection`
    - `ShowBackupPolicyInfo`
    - `UpdateBackupPolicyInfo`
    - `ChangeEvent`
    - `ListAlarmWhiteList`
    - `ListHostGroups`
    - `ChangeHostsGroup`
    - `AddHostsGroup`
    - `DeleteHostsGroup`
    - `ListPolicyGroup`
    - `AssociatePolicyGroup`
    - `ListVulHosts`
    - `ChangeVulStatus`
    - `ListWtpProtectHost`
    - `SetWtpProtectionStatusInfo`
    - `SetRaspSwitch`
    - `ListHostProtectHistoryInfo`
    - `ListHostRaspProtectHistoryInfo`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRiskConfigCheckRules`新增响应参数 `fix_status`、`enable_auto_fix`、`rule_params`
  - 接口`ListSecurityEvents`新增响应参数 `extend_info`
  - 接口`ListHostStatus`:
    - 新增响应参数 `enterprise_project_id`、`agent_create_time`、`agent_update_time`、`agent_version`、`upgrade_status`、`upgrade_result_code`、`upgradable`
    - 请求参数`region`改为必填
  - 接口`ListVulnerabilities`请求参数`vul_id`改为必填

### HuaweiCloud SDK IEF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTags`响应参数`key`、`value`改为非必填
  - 接口`CreateTag`请求参数`key`、`value`改为非必填
  - 接口`ListEdgeNodes`响应参数`key`、`value`改为非必填
  - 接口`ShowEdgeNodeDetail`响应参数`key`、`value`改为非必填
  - 接口`UpdateEdgeNode`响应参数`key`、`value`改为非必填
  - 接口`CreateEdgeGroup`新增请求参数 `device_ids`
  - 接口`ListEdgeGroupCerts`:
    - 新增响应参数 `groupcerts`
    - 移除响应参数 `edge_groups`
  - 接口`ListDevices`响应参数`type`改为非必填
  - 接口`CreateDevice`请求参数`type`改为非必填
  - 接口`ShowDevice`响应参数`type`改为非必填
  - 接口`UpdateDevice`响应参数`type`改为非必填
  - 接口`ShowDeviceTwin`响应参数`type`改为非必填
  - 接口`UpdateDeviceTwin`:
    - 请求参数`twin`、`property_visitors`改为非必填
    - 响应参数`type`改为非必填
  - 接口`ListDeviceTemplates`响应参数`key`、`value`改为非必填
  - 接口`CreateDeviceTemplate`请求参数`key`、`value`改为非必填
  - 接口`ShowDeviceTemplate`响应参数`key`、`value`改为非必填
  - 接口`UpdateDeviceTemplateById`响应参数`key`、`value`改为非必填
  - 接口`ListResourceByTags`响应参数`key`、`value`改为非必填
  - 接口`BatchAddDeleteTags`请求参数`key`、`value`改为非必填
  - 接口`ListApps`响应参数`read_only`改为非必填
  - 接口`ShowAppDetail`响应参数`read_only`改为非必填
  - 接口`UpdateApp`响应参数`read_only`改为非必填
  - 接口`ListAppVersions`响应参数`read_only`改为非必填
  - 接口`CreateAppVersions`请求参数`read_only`改为非必填
  - 接口`ShowAppVersionDetail`响应参数`read_only`改为非必填
  - 接口`UpdateAppVersion`:
    - 请求参数`read_only`改为非必填
    - 响应参数`read_only`改为非必填
  - 接口`ListDeployments`响应参数`host_network`、`read_only`改为非必填
  - 接口`CreateDeployments`请求参数`host_network`、`read_only`改为非必填
  - 接口`ShowDeployment`响应参数`host_network`、`read_only`改为非必填
  - 接口`UpdateDeployment`:
    - 请求参数`replicas`、`host_network`、`read_only`改为非必填
    - 响应参数`host_network`、`read_only`改为非必填
  - 接口`ListPods`响应参数`host_network`、`read_only`改为非必填
  - 接口`CreateEncryptdatas`请求参数`description`改为非必填
  - 接口`UpdateEncryptdatas`请求参数`description`改为非必填
  - 接口`ListBatchJob`:
    - 新增响应参数 `task_total_count`、`task_success_count`、`task_failed_count`、`status_last_updated_at`、`description`
    - 移除响应参数 `task_count`、`success_count`、`failed_count`、`updated_at`
  - 接口`ShowBatchJob`:
    - 新增响应参数 `status_last_updated_at`
    - 移除响应参数 `updated_at`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRecordCallbackConfig`新增请求参数 `key`
  - 接口`UpdateRecordCallbackConfig`:
    - 新增请求参数 `key`
    - 新增响应参数 `id`、`publish_domain`、`app`、`notify_callback_url`、`notify_event_subscription`、`sign_type`、`create_time`、`update_time`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateSqlAlarmRule`新增请求参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`CreateSqlAlarmRule`新增请求参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`ListSqlAlarmRules`新增响应参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`UpdateKeywordsAlarmRule`新增请求参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`CreateKeywordsAlarmRule`新增请求参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`ListKeywordsAlarmRules`新增响应参数 `trigger_condition_count`、`trigger_condition_frequency`、`whether_recovery_policy`、`recovery_policy`
  - 接口`ListHost`请求参数`host_id_list`改为非必填
  - 接口`UpdateStructConfig`请求参数`is_analysis`、`is_analysis`改为非必填
  - 接口`CreateStructConfig`请求参数`is_analysis`、`is_analysis`改为非必填

### HuaweiCloud SDK OSM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAuthorizations`新增响应参数 `resource_type_id`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`StartInstanceRestartAction`请求参数`restart`改为必填

### HuaweiCloud SDK Workspace

- _新增特性_
  - 支持接口`ShowQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.15 2022-12-08

### HuaweiCloud SDK MapDS

- _新增特性_
  - 支持地图数据服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`GetExecutionPlan`、`DeleteExecutionPlan`、`DescribeExecutionPlan`、`GetStackMetadata`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListExecutionPlans`移除请求参数 `executor`
  - 接口`CreateExecutionPlan`移除请求参数 `executor`
  - 接口`ApplyExecutionPlan`移除请求参数 `executor`
  - 接口`ListStackEvents`:
    - 移除请求参数 `limit`、`marker`、`executor`
    - 移除响应参数 `next_marker`
  - 接口`ListStacks`移除请求参数 `executor`
  - 接口`CreateStack`移除请求参数 `executor`
  - 接口`GetStackTemplate`移除请求参数 `executor`
  - 接口`ListStackResources`:
    - 移除请求参数 `executor`
    - 移除响应参数 `create_time`、`update_time`
  - 接口`ListStackOutputs`:
    - 移除请求参数 `executor`、`limit`、`marker`
    - 移除响应参数 `next_marker`
  - 接口`DeployStack`移除请求参数 `executor`
  - 接口`DeleteStack`移除请求参数 `executor`

### HuaweiCloud SDK APM

- _新增特性_
  - 支持接口`SearchAgent`、`ChangeAgentStatus`、`DeleteAgent`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CBH

- _新增特性_
  - 支持以下接口：
    - `ListQuotaState`
    - `ShowNetworkConfiguration`
    - `StopCbhInstance`
    - `UpgradeCbhInstance`
    - `RestartCbhInstance`
    - `InstallInstanceEip`
    - `UninstallInstanceEip`
    - `ResetPassword`
    - `ShowAvailableZoneInfo`
    - `StartCbhInstance`
    - `SearchQuota`
    - `ResetLoginMethod`
    - `CreateInstanceOrder`
    - `ChangeInstanceNetwork`
    - `CreateInstance`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCbhInstance`:
    - 新增响应参数 `public_ip`、`is_auto_renew`
    - 移除请求参数 `X-Auth-Token`
    - 移除响应参数 `publicip`
    - 响应参数`zh_cn`、`en_us`、`exp_time`、`start_time`、`end_time`、`release_time`、`instance_id`、`private_ip`、`task_status`、`vpc_id`、`subnet_id`、`security_group_id`、`createinstance_status`、`fail_reason`、`instance_key`、`order_id`、`period_num`、`resource_id`、`public_id`、`alter_permit`、`bastion_version`、`new_bastion_version`、`instance_status`、`instance_description`改为必填

### HuaweiCloud SDK CC

- _新增特性_
  - 支持以下接口：
    - `ListDomainTags`
    - `DeleteTag`
    - `BatchCreateDeleteTags`
    - `ListResourceByFilterTag`
    - `ListTags`
    - `CreateTag`
    - `ListQuotas`
    - `ListBandwidthPackages`
    - `CreateBandwidthPackage`
    - `ShowBandwidthPackage`
    - `UpdateBandwidthPackage`
    - `DeleteBandwidthPackage`
    - `AssociateBandwidthPackage`
    - `DisassociateBandwidthPackage`
    - `ListInterRegionBandwidths`
    - `CreateInterRegionBandwidth`
    - `ShowInterRegionBandwidth`
    - `UpdateInterRegionBandwidth`
    - `DeleteInterRegionBandwidth`
    - `ListAuthorisations`
    - `CreateAuthorisation`
    - `ListPermissions`
    - `UpdateAuthorisation`
    - `DeleteAuthorisation`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateAddonInstance`请求参数`version`改为非必填
  - 接口`UpdateAddonInstance`请求参数`version`改为非必填

### HuaweiCloud SDK CFW

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIpsSwitchStatusUsingGet`:
    - 新增响应参数 `id`、`virtual_patches_status`
    - 移除响应参数 `object_id`、`virtual_patches_stauts`
  - 接口`ChangeIpsSwitchUsingPost`请求参数`ips_type`改为必填
  - 接口`ListFirewallUsingGet`移除响应参数 `fw_instance_id`、`resource_id`、`name`、`ha_type`、`charge_mode`、`service_type`、`engine_type`、`flavor`、`protect_objects`、`status`、`description`、`is_old_firewall_instance`、`support_ipv6`、`feature_toggle`

### HuaweiCloud SDK CloudDeploy

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`StartDeployTask`请求参数`type`新增枚举值`enum`

### HuaweiCloud SDK DBSS

- _新增特性_
  - 支持以下接口：
    - `ListAuditInstances`
    - `ShowAuditQuota`
    - `ListEcsSpecification`
    - `ListAvailabilityZoneInfos`
    - `ListAuditInstanceJobs`
    - `ListAuditDatabases`
    - `ListAuditOperateLogs`
    - `ListAuditRuleScopes`
    - `ListSqlInjectionRules`
    - `ListAuditRuleRisks`
    - `ShowAuditRuleRisk`
    - `ListAuditSensitiveMasks`
    - `ListProjectResourceTags`
    - `ListResourceInstanceByTag`
    - `CountResourceInstanceByTag`
    - `BatchAddResourceTag`
    - `BatchDeleteResourceTag`
    - `CreateInstancesPeriodOrder`
    - `UpdateAuditSecurityGroup`
    - `AddRdsNoAgentDatabase`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateJobsAsync`请求参数`name`、`job_type`、`engine_type`、`job_direction`、`task_type`、`net_type`、`enterprise_project_id`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`is_auto_renew`、`security_group_id`改为非必填
  - 接口`CreateJob`请求参数`name`、`job_type`、`engine_type`、`job_direction`、`task_type`、`net_type`、`enterprise_project_id`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`is_auto_renew`、`security_group_id`改为非必填
  - 接口`ListAsyncJobDetail`响应参数`name`、`job_type`、`engine_type`、`job_direction`、`task_type`、`net_type`、`enterprise_project_id`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`is_auto_renew`、`security_group_id`改为非必填
  - 接口`UpdateBatchAsyncJobs`请求参数`name`、`job_type`、`engine_type`、`job_direction`、`task_type`、`net_type`、`enterprise_project_id`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`is_auto_renew`、`security_group_id`改为非必填
  - 接口`ShowJobDetail`响应参数`name`、`job_type`、`engine_type`、`job_direction`、`task_type`、`net_type`、`enterprise_project_id`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`is_auto_renew`、`security_group_id`改为非必填
  - 接口`UpdateJob`请求参数`name`、`job_type`、`engine_type`、`job_direction`、`task_type`、`net_type`、`enterprise_project_id`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`、`is_auto_renew`、`security_group_id`改为非必填
  - 接口`ExecuteJobAction`:
    - 新增请求参数 `is_sync_re_edit`、`force_delete`
    - 请求参数`job_id`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`改为非必填
  - 接口`BatchExecuteJobActions`:
    - 新增请求参数 `is_sync_re_edit`、`force_delete`
    - 请求参数`job_id`、`ip`、`db_port`、`ssl_link`、`ssl_cert_name`、`ssl_cert_key`、`ssl_cert_check_sum`改为非必填

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持以下接口：
    - `CreateProductOrder`
    - `ShowSpecification`
    - `ShowTopics`
    - `UpdateDefaultTopic`
    - `ShowRules`
    - `ChangeRule`
    - `AddRule`
    - `DeleteRule`
    - `ListRuleGroups`
    - `AddRuleGroup`
    - `DeleteRuleGroup`
    - `AddScanJob`
    - `ListRelationDb`
    - `ListRelationTable`
    - `ListRelationColumn`
    - `ListRelationBuckets`
    - `ListRelationFile`
    - `ListDbMaskTask`
    - `ChangeDbTemplateOperation`
    - `UpdateAssetName`
    - `ListBuckets`
    - `AddBuckets`
    - `DeleteBucket`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持接口`ListHostOverview`、`ListHostDisk`、`ListHostNet`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持以下接口：
    - `ListBandwidthPkg`
    - `CountPublicIp`
    - `ShowPublicIpType`
    - `CountPublicIpInstance`
    - `BatchCreatePublicips`
    - `BatchDeletePublicIp`
    - `BatchDisassociatePublicips`
    - `CountEipAvailableResources`
- _解决问题_
  - 无
- _特性变更_
  - 接口`AssociatePublicips`请求参数`associate_instance_type`移除枚举值``
  - 接口`UpdateAssociatePublicip`请求参数`associate_instance_type`移除枚举值``

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDependency`新增响应参数 `dep_id`
  - 接口`CreateDependencyVersion`新增响应参数 `dep_id`
  - 接口`UpdateFunctionConfig`新增请求参数 `enable_dynamic_memory`、`enable_auth_in_header`
  - 接口`ShowWorkflowExecutionForPage`:
    - 新增请求参数 `offset`、`limit`
    - 新增响应参数 `total`、`size`、`executions`
    - 移除请求参数 `page`、`page_size`
    - 移除响应参数 `pager`、`hisRecords`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持以下接口：
    - `RestartGaussMySqlInstance`
    - `ShowGaussMySqlConfiguration`
    - `UpdateGaussMySqlConfiguration`
    - `DeleteGaussMySqlConfiguration`
    - `SwitchGaussMySqlConfiguration`
    - `RestartGaussMySqlNode`
    - `UpdateProxySessionConsistence`
    - `ListImmediateJobs`
    - `ListScheduleJobs`
    - `CancelScheduleTask`
    - `DeleteTaskRecord`
    - `UpgradeGaussMySqlInstanceDatabase`
    - `SwitchGaussMySqlInstanceSsl`
    - `UpdateGaussMySqlInstanceEip`
    - `CancelGaussMySqlInstanceEip`
    - `InvokeGaussMySqlInstanceSwitchOver`
    - `UpdateGaussMySqlInstanceOpsWindow`
    - `UpdateGaussMySqlInstanceSecurityGroup`
    - `UpdateGaussMySqlInstanceInternalIp`
    - `UpdateGaussMySqlInstancePort`
    - `UpdateGaussMySqlInstanceAlias`
    - `DeleteGaussMySqlBackup`
    - `CreateGaussMySqlConfiguration`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAvailableFlavorInfos`:
    - 新增请求参数 `x-auth-token`
    - 新增响应参数 `list`
    - 移除响应参数 `optional_flavor_list`
    - 响应参数`instance_id`、`instance_name`、`vcpus`、`ram`、`spec_code`、`az_status`、`region_status`、`total_count`改为必填
  - 接口`ShowSlowLogDesensitization`:
    - 新增请求参数 `x-auth-token`
    - 响应参数`desensitization_status`改为必填
  - 接口`ListProjectTags`响应参数`type`、`key`、`values`、`total_count`改为必填
  - 接口`ModifyEpsQuotas`请求参数`instance`、`vcpus`、`ram`改为必填
  - 接口`ListEpsQuotas`:
    - 移除响应参数 `enterprise_project_id`、`enterprise_project_name`、`quota`、`used`
    - 响应参数`total_count`改为必填

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ResizeInstance`新增请求参数 `publicip_id`

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持接口`ListTimeLineTrafficStatistics`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`DownloadAttachment`、`DeleteAttachment`、`ListStatusStatistic`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateIssueV4`新增响应参数 `user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`
  - 接口`ListIssuesSfV4`新增响应参数 `user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`
  - 接口`ListIssuesV4`新增响应参数 `user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`
  - 接口`UpdateIssueV4`新增响应参数 `user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`
  - 接口`ShowIssueV4`新增响应参数 `user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`
  - 接口`ListChildIssuesV4`新增响应参数 `user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`
  - 接口`ListProjectIssuesRecordsV4`新增响应参数 `user_id`、`user_num_id`
  - 接口`CreateSystemIssueV4`新增响应参数 `user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`、`user_id`、`user_num_id`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ShowSecondLevelMonitoring`、`SetSecondLevelMonitor`、`ShowAutoEnlargePolicy`、`SetAutoEnlargePolicy`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePostPaidInstance`请求参数`broker_num`改为必填
  - 接口`UpdateInstance`:
    - 新增请求参数 `enable_acl`
    - 移除请求参数 `retention_policy`
  - 接口`BatchUpdateConsumerGroup`新增响应参数 `job_id`
  - 接口`ShowGroup`:
    - 新增响应参数 `app_id`、`app_name`、`permissions`
    - 移除响应参数 `from_beginning`

### HuaweiCloud SDK TMS

- _新增特性_
  - 支持以下接口：
    - `ListResource`
    - `CreateResourceTag`
    - `DeleteResourceTag`
    - `ListTagKeys`
    - `ListTagValues`
    - `ShowResourceTag`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.14 2022-12-01

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerselfResourceRecords`新增响应参数 `formula`

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerselfResourceRecords`新增响应参数 `formula`

### HuaweiCloud SDK CFW

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIpsSwitchStatusUsingGet`:
    - 新增响应参数 `data`
    - 移除响应参数 `object_id`、`basic_defense_status`、`virtual_patches_stauts`
  - 接口`ListEastWestFirewall`:
    - 新增响应参数 `protect_infos`
    - 移除响应参数 `protected_infos`
  - 接口`ListAttackLogs`请求参数`fw_instance_id`改为非必填
  - 接口`UpdateRuleAclUsingPut`新增请求参数 `type`
  - 接口`UpdateBlackWhiteListUsingPut`新增请求参数 `list_type`、`object_id`
  - 接口`ListFirewallUsingGet`:
    - 新增响应参数 `data`
    - 移除响应参数 `fw_instance_id`、`resource_id`、`name`、`ha_type`、`charge_mode`、`service_type`、`engine_type`、`flavor`、`protect_objects`、`status`、`description`、`is_old_firewall_instance`、`support_ipv6`、`feature_toggle`
  - 接口`ListServiceSetDetails`:
    - 新增响应参数 `data`
    - 移除响应参数 `id`、`name`、`description`
  - 接口`CountEips`:
    - 新增响应参数 `data`
    - 移除响应参数 `object_id`、`eip_total`、`eip_protected`
  - 接口`ListEipResources`:
    - 新增响应参数 `data`
    - 移除响应参数 `id`、`public_ip`、`status`、`public_ipv6`、`enterprise_project_id`、`device_id`、`device_name`、`device_owner`、`associate_instance_type`
  - 接口`UpdateAddressSetInfoUsingPut`新增请求参数 `address_type`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`CreateAcceptance`、`CreateRequest`、`ShowResult`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持以下接口：
    - `CheckCluster`
    - `ShrinkCluster`
    - `ListDssPools`
    - `ListStatistics`
    - `SwitchOverCluster`
    - `CancelReadonlyCluster`
    - `UpdateMaintenanceWindow`
    - `ListClusterCn`
    - `BatchCreateClusterCn`
    - `BatchDeleteClusterCn`
    - `ListClusterConfigurations`
    - `ListClusterConfigurationsParameter`
    - `UpdateConfiguration`
    - `ListSnapshotStatistics`
    - `ExecuteRedistributionCluster`
    - `ListClusterWorkload`
    - `CreateClusterWorkload`
    - `CreateWorkloadPlan`
    - `ListWorkloadQueue`
    - `AddWorkloadQueue`
    - `DeleteWorkloadQueue`
    - `CopySnapshot`
    - `ListClusterSnapshots`
    - `DeleteSnapshotPolicy`
    - `ListSnapshotPolicy`
    - `CreateSnapshotPolicy`
    - `DeleteDisasterRecovery`
    - `StartDisasterRecovery`
    - `PauseDisasterRecovery`
    - `SwitchoverDisasterRecovery`
    - `SwitchFailoverDisaster`
    - `RestoreDisaster`
    - `ListAuditLog`
    - `ListDataSource`
    - `CreateDataSource`
    - `UpdateDataSource`
    - `ListElbs`
    - `AssociateElb`
    - `DisassociateElb`
    - `AssociateEip`
    - `DisassociateEip`
    - `UpdateClusterDns`
    - `CreateClusterDns`
    - `DeleteClusterDns`
    - `ListAlarmConfigs`
    - `ListAlarmDetail`
    - `ListAlarmStatistic`
    - `ListAlarmSubs`
    - `CreateAlarmSub`
    - `UpdateAlarmSub`
    - `DeleteAlarmSub`
    - `ListEvents`
    - `ListEventSpecs`
    - `ListEventSubs`
    - `CreateEventSub`
    - `UpdateEventSub`
    - `DeleteEventSub`
    - `ListJobDetails`
    - `ListQuotas`
    - `ListTags`
    - `ListClusterTags`
    - `BatchCreateResourceTag`
    - `BatchDeleteResourceTag`
    - `ListAvailabilityZones`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持接口`AsyncInvokeReservedFunction`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDependencies`响应参数`runtime`新增枚举值`http`
  - 接口`ListDependencyVersion`响应参数`runtime`新增枚举值`http`
  - 接口`CreateFunction`:
    - 请求参数`runtime`新增枚举值`http`
    - 响应参数`runtime`新增枚举值`http`
  - 接口`ListFunctions`响应参数`runtime`新增枚举值`http`
  - 接口`UpdateFunctionCode`响应参数`runtime`新增枚举值`http`
  - 接口`ShowFunctionCode`响应参数`runtime`新增枚举值`http`
  - 接口`UpdateFunctionConfig`:
    - 请求参数`runtime`新增枚举值`http`
    - 响应参数`runtime`新增枚举值`http`
  - 接口`ShowFunctionConfig`响应参数`runtime`新增枚举值`http`
  - 接口`UpdateFunctionMaxInstanceConfig`响应参数`runtime`新增枚举值`http`
  - 接口`ListReservedInstanceConfigs`:
    - 新增响应参数 `reservedinstances`、`page_info`、`count`
    - 移除响应参数 `function_urn`、`qualifier_type`、`qualifier_name`、`min_count`、`idle_mode`、`tactics_config`
  - 接口`CreateFunctionVersion`响应参数`runtime`新增枚举值`http`
  - 接口`ListFunctionVersions`响应参数`runtime`新增枚举值`http`
  - 接口`ImportFunction`响应参数`runtime`新增枚举值`http`

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListProPricePlans`新增请求参数 `sim_card_id`、`partner`、`package_type`、`sim_type`
  - 接口`ListSimCards`:
    - 新增请求参数 `expire_time_duration`、`traffic_warning_threshold`、`sim_status_in`
    - 新增响应参数 `sn`、`supply_code`、`bundle_id`、`test_type`
  - 接口`StopSimCard`:
    - 新增请求参数 `price_plan_list`
    - 新增响应参数 `sim_price_plan_list`
  - 接口`ResetSimCard`:
    - 新增请求参数 `price_plan_list`
    - 新增响应参数 `sim_price_plan_list`
  - 接口`ShowSimCard`新增响应参数 `sn`、`supply_code`、`bundle_id`、`test_type`
  - 接口`ListSimPools`新增响应参数 `order_id`
  - 接口`ListSimPricePlans`:
    - 新增请求参数 `sim_price_plan_id`
    - 新增响应参数 `partner`、`partner_pid`
  - 接口`CreateAttribute`新增响应参数 `id`

### HuaweiCloud SDK Image

- _新增特性_
  - 支持接口`RunQueryCustomTags`、`RunDeleteCustomTags`、`RunImageMediaTaggingDet`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunImageMediaTagging`新增请求参数 `use_default_tags`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`CopyConfiguration`、`ListInstanceParamHistories`、`ListMsdtcHosts`、`BatchAddMsdtcs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IEF

- _新增特性_
  - 支持以下接口：
    - `ListEdgeGroups`
    - `CreateEdgeGroup`
    - `ShowEdgeGroupDetail`
    - `UpdateEdgeGroup`
    - `DeleteEdgeGroup`
    - `UpdateEdgeGroupNodeBinding`
    - `ListEdgeGroupCerts`
    - `CreateEdgeGroupCert`
    - `ShowEdgeGroupCertDetail`
    - `DeleteEdgeGroupCert`
    - `ListSystemEvents`
    - `CreateSystemEvent`
    - `ShowSystemEventDetail`
    - `DeleteSystemEvent`
    - `StartSystemEvent`
    - `StopSystemEvent`
    - `ListBatchJob`
    - `CreateBatchJob`
    - `ShowBatchJob`
    - `DeleteBatchJob`
    - `StopBatchJob`
    - `ListProducts`
    - `CreateProduct`
    - `ShowProductDetail`
    - `DeleteProduct`
    - `RestoreBatchJob`
    - `RetryBatchJob`
    - `ShowQuota`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEdgeNodes`:
    - 新增响应参数 `identifier`
  - 接口`ShowEdgeNodeDetail`:
    - 新增响应参数 `identifier`
  - 接口`UpdateEdgeNode`:
    - 新增响应参数 `identifier`
  - 接口`UpdateNodeByDeviceId`:
    - 请求参数`relation`、`comment`改为非必填
    - 响应参数`relation`、`comment`改为非必填
  - 接口`ListApps`:
    - 新增响应参数 `run_as_user`
    - 响应参数`host_pid`类型调整 `string` -> `boolean`
  - 接口`ShowAppDetail`:
    - 新增响应参数 `run_as_user`
    - 响应参数`host_pid`类型调整 `string` -> `boolean`
  - 接口`UpdateApp`:
    - 新增响应参数 `run_as_user`
    - 响应参数`host_pid`类型调整 `string` -> `boolean`
  - 接口`ListAppVersions`:
    - 新增响应参数 `run_as_user`
    - 响应参数`host_pid`类型调整 `string` -> `boolean`
  - 接口`CreateAppVersions`:
    - 新增请求参数 `run_as_user`
    - 请求参数`host_pid`类型调整 `string` -> `boolean`
  - 接口`ShowAppVersionDetail`:
    - 新增响应参数 `run_as_user`
    - 响应参数`host_pid`类型调整 `string` -> `boolean`
  - 接口`UpdateAppVersion`:
    - 新增请求参数`run_as_user`
    - 新增响应参数 `run_as_user`
    - 请求参数`host_pid`类型调整 `string` -> `boolean`
    - 响应参数`host_pid`类型调整 `string` -> `boolean`
  - 接口`ListDeployments`:
    - 新增响应参数 `run_as_user`、`run_as_user`
  - 接口`CreateDeployments`新增请求参数 `run_as_user`、`run_as_user`
  - 接口`ShowDeployment`:
    - 新增响应参数 `run_as_user`、`run_as_user`
  - 接口`UpdateDeployment`:
    - 新增请求参数 `run_as_user`、`run_as_user`
    - 新增响应参数 `run_as_user`、`run_as_user`

# 0.1.13 2022-11-30

### HuaweiCloud SDK AOM

- _新增特性_
  - 支持以下接口：
    - `ListNotifiedHistories`
    - `ListMuteRule`
    - `UpdateMuteRule`
    - `AddMuteRules`
    - `DeleteMuteRules`
    - `ShowActionRule`
    - `ListActionRule`
    - `UpdateActionRule`
    - `AddActionRule`
    - `DeleteActionRule`
    - `ListEvent2alarmRule`
    - `UpdateEventRule`
    - `AddEvent2alarmRule`
    - `DeleteEvent2alarmRule`
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddAlarmRule`请求参数`statistic`类型调整 `string` -> `enum`

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateApiV2`新增请求参数 `retry_count`、`retry_count`
  - 接口`UpdateApiV2`:
    - 新增请求参数 `retry_count`、`retry_count`
    - 新增响应参数 `retry_count`、`retry_count`
  - 接口`ShowDetailsOfApiV2`新增响应参数 `retry_count`、`retry_count`
  - 接口`ListApiVersionDetailV2`新增响应参数 `retry_count`、`retry_count`
  - 接口`AssociateDomainV2`新增请求参数 `is_http_redirect_to_https`
  - 接口`UpdateDomainV2`新增请求参数 `is_http_redirect_to_https`
  - 接口`ImportMicroservice`新增请求参数 `labels`
  - 接口`ListAttachedDomainsV2`新增响应参数 `is_http_redirect_to_https`

### HuaweiCloud SDK CPH

- _新增特性_
  - 支持接口`PushFile`、`InstallApk`、`UninstallApk`、`PushShareFiles`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`DeleteShareFiles`
  - 接口`ShowBandwidthDetail`响应参数`band_widths`、`request_id`改为非必填
  - 接口`UpdateBandwidth`响应参数`request_id`改为非必填
  - 接口`ListEncodeServers`响应参数`encode_servers`、`request_id`改为非必填
  - 接口`RestartEncodeServer`响应参数`jobs`、`request_id`改为非必填
  - 接口`ListJobs`响应参数`jobs`、`request_id`改为非必填
  - 接口`ShowJob`响应参数`error_msg`、`execute_msg`、`job_id`、`end_time`、`begin_time`、`error_code`、`status`、`request_id`改为非必填
  - 接口`ListCloudPhoneImages`响应参数`phone_images`、`request_id`改为非必填
  - 接口`ListCloudPhoneModels`:
    - 新增请求参数 `status`
    - 响应参数`phone_models`、`request_id`改为非必填
  - 接口`ListCloudPhones`响应参数`phones`、`request_id`改为非必填
  - 接口`CreateCloudPhoneServer`响应参数`order_id`、`product_id`、`request_id`改为非必填
  - 接口`ImportTraffic`响应参数`jobs`、`request_id`改为非必填
  - 接口`ShowCloudPhoneDetail`响应参数`phone_name`、`server_id`、`phone_id`、`image_id`、`vnc_enable`、`phone_model_name`、`status`、`access_infos`、`property`、`metadata`、`create_time`、`update_time`、`request_id`改为非必填
  - 接口`UpdatePhoneName`响应参数`request_id`改为非必填
  - 接口`BatchMigrateCloudPhone`响应参数`jobs`、`request_id`改为非必填
  - 接口`ResetCloudPhone`响应参数`jobs`、`request_id`改为非必填
  - 接口`RestartCloudPhone`响应参数`jobs`、`request_id`改为非必填
  - 接口`BatchImportCloudPhoneData`响应参数`jobs`、`request_id`改为非必填
  - 接口`StopCloudPhone`响应参数`jobs`、`request_id`改为非必填
  - 接口`BatchExportCloudPhoneData`响应参数`jobs`、`request_id`改为非必填
  - 接口`UpdateCloudPhoneProperty`响应参数`jobs`、`request_id`改为非必填
  - 接口`RunShellCommand`响应参数`jobs`、`request_id`改为非必填
  - 接口`PushShareApps`响应参数`jobs`、`request_id`改为非必填
  - 接口`DeleteShareApps`响应参数`jobs`、`request_id`改为非必填
  - 接口`DeleteShareFiles`:
    - 新增请求参数 `DeleteShareFilesRequestBody`
    - 移除请求参数 `PushShareFilesRequestBody`
    - 响应参数`jobs`、`request_id`改为非必填
  - 接口`RunSyncCommand`响应参数`jobs`、`request_id`改为非必填
  - 接口`ListCloudPhoneServerModels`响应参数`server_models`、`request_id`改为非必填
  - 接口`ListCloudPhoneServers`响应参数`servers`、`request_id`改为非必填
  - 接口`ShowCloudPhoneServerDetail`响应参数`server_name`、`availability_zone`、`server_id`、`server_model_name`、`phone_model_name`、`keypair_name`、`status`、`vpc_id`、`cidr`、`vpc_cidr`、`subnet_id`、`subnet_cidr`、`resource_project_id`、`metadata`、`intranet_ip`、`access_ip`、`server_ip`、`public_ip`、`band_width_name`、`band_width_id`、`band_width_size`、`band_width_charge_mode`、`band_width_share_type`、`create_time`、`update_time`、`volume_name`、`volume_id`、`volume_size`、`volume_type`、`create_time`、`update_time`、`network_version`、`security_groups`、`create_time`、`update_time`、`request_id`改为非必填
  - 接口`UpdateServerName`响应参数`request_id`改为非必填
  - 接口`RestartCloudPhoneServer`响应参数`jobs`、`request_id`改为非必填
  - 接口`ChangeCloudPhoneServerModel`响应参数`order_id`、`product_id`、`request_id`改为非必填
  - 接口`UpdateKeypair`响应参数`jobs`、`request_id`改为非必填
  - 接口`ListShareFiles`响应参数`jobs`、`request_id`改为非必填
  - 接口`CreateNet2CloudPhoneServer`响应参数`order_id`、`product_id`、`request_id`改为非必填

### HuaweiCloud SDK DDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ExpandInstanceNodes`新增请求参数 `is_auto_pay`

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddReadonlyNode`请求参数`num`类型调整 `string` -> `int32`

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持接口`DisassociatePublicips`、`AssociatePublicips`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTenantMetric`新增请求参数 `metric_type`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持以下接口：
    - `ListAvailableFlavorInfos`
    - `CheckWeekPassword`
    - `ModifyPort`
    - `UpdateClientNetwork`
    - `DeleteEnlargeFailNode`
    - `ShowIpNumRequirement`
    - `ShowAutoEnlargePolicy`
    - `ShowSlowLogDesensitization`
    - `SwitchSlowlogDesensitization`
    - `ShowErrorLog`
    - `CopyConfiguration`
    - `CompareConfiguration`
    - `ListConfigurationDatastores`
    - `ShowAllInstancesBackups`
    - `CreateBack`
    - `ShowRecyclePolicy`
    - `SetRecyclePolicy`
    - `ListRecycleInstances`
    - `ShowPauseResumeStutus`
    - `PauseResumeDataSynchronization`
    - `ListProjectTags`
    - `ListEpsQuotas`
    - `ModifyEpsQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ModifyPublicIp`请求参数`public_ip`、`public_ip_id`改为非必填
  - 接口`SwitchToMaster`新增请求参数 `SwitchToMasterDisasterRecoveryBody`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持以下接口：
    - `CopyConfiguration`
    - `DeleteJob`
    - `ListEpsQuotas`
    - `ModifyEpsQuota`
    - `ListAvailableFlavors`
    - `ShowSslCertDownloadLink`
    - `ListApplicableInstances`
    - `ListAppliedHistories`
    - `ListRestorableInstances`
    - `ListTasks`
    - `ShowDeploymentForm`
    - `ShowInstanceDisk`
    - `ListProjectTags`
    - `ListInstanceTags`
    - `AddInstanceTags`
    - `ShowProjectQuotas`
    - `ListRecycleInstances`
    - `ShowConfigurationDetail`
    - `DeleteConfiguration`
    - `ListConfigurationsDiff`
    - `ShowBalanceStatus`
    - `ShowInstanceSnapshot`
    - `ValidateWeakPassword`
    - `ListHistoryOperations`
    - `ListBindedEips`
    - `ResetConfiguration`
    - `ListGaussDbDatastores`
    - `AttachEip`
    - `ListPredefinedTags`
    - `SwitchConfiguration`
    - `ValidateParaGroupName`
    - `ShowRecyclePolicy`
    - `CreateConfigurationTemplate`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowBackupPolicy`新增响应参数 `enable_standby_backup`
  - 接口`SetBackupPolicy`新增请求参数 `enable_standby_backup`

### HuaweiCloud SDK MPC

- _新增特性_
  - 支持以下接口：
    - `ListNotifySmnTopicConfig`
    - `NotifySmnTopicConfig`
    - `ListNotifyEvent`
    - `DeleteTranscodingTaskByConsole`
    - `ListStatSummary`
    - `ListAllBuckets`
    - `UpdateBucketAuthorized`
    - `ListAllObsObjList`
    - `ShowAgenciesTask`
    - `CreateAgenciesTask`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持接口`UpdateClusterName`、`ShowAutoScalingPolicy`
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateAgencyMapping`响应参数`result`改为必填

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`响应参数`read_only_by_user`类型调整 `string` -> `boolean`

### HuaweiCloud SDK ROMA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateProduct`移除响应参数 `status`
  - 接口`ListProperties`移除响应参数 `bool_false`、`bool_true`
  - 接口`CreateProperty`移除请求参数 `bool_false`、`bool_true`
  - 接口`ShowProperty`移除响应参数 `bool_false`、`bool_true`
  - 接口`UpdateProperty`:
    - 移除请求参数 `bool_false`、`bool_true`
    - 移除响应参数 `bool_false`、`bool_true`
  - 接口`ListRequestProperties`移除响应参数 `bool_false`、`bool_true`
  - 接口`CreateRequestProperty`:
    - 移除请求参数 `bool_false`、`bool_true`
    - 移除响应参数 `bool_false`、`bool_true`
  - 接口`ShowRequestProperty`移除响应参数 `bool_false`、`bool_true`
  - 接口`UpdateRequestProperty`:
    - 移除请求参数 `bool_false`、`bool_true`
    - 移除响应参数 `bool_false`、`bool_true`
  - 接口`ListResponseProperties`移除响应参数 `bool_false`、`bool_true`
  - 接口`CreateResponseProperty`移除请求参数 `bool_false`、`bool_true`
  - 接口`ShowResponseProperty`移除响应参数 `bool_false`、`bool_true`
  - 接口`UpdateResponseProperty`:
    - 移除请求参数 `bool_false`、`bool_true`
    - 移除响应参数 `bool_false`、`bool_true`
  - 接口`CreateApiV2`新增请求参数 `cookie_param_name`、`alias_urn`、`cookie_param_name`、`alias_urn`、`cookie_param_name`
  - 接口`ShowDetailsOfApiV2`新增响应参数 `cookie_param_name`、`alias_urn`、`alias_urn`、`cookie_param_name`、`cookie_param_name`
  - 接口`UpdateApiV2`:
    - 新增请求参数 `cookie_param_name`、`alias_urn`、`cookie_param_name`、`alias_urn`、`cookie_param_name`
    - 新增响应参数 `cookie_param_name`、`alias_urn`、`alias_urn`、`cookie_param_name`、`cookie_param_name`
  - 接口`ListApiVersionDetailV2`新增响应参数 `cookie_param_name`、`alias_urn`、`alias_urn`、`cookie_param_name`、`cookie_param_name`
  - 接口`ListVpcChannelsV2`新增响应参数 `microservice_info`、`type`
  - 接口`ShowDetailsOfVpcChannelV2`新增响应参数 `microservice_info`、`type`
  - 接口`UpdateVpcChannelV2`新增响应参数 `microservice_info`、`type`
  - 接口`ListProjectVpcChannelsV2`新增响应参数 `microservice_info`、`type`
  - 接口`UpdateProjectVpcChannel`新增响应参数 `microservice_info`、`type`
  - 接口`ShowDetailsOfInstanceV2`新增响应参数 `ingress_ips`

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`PushTranscriberJobs`请求参数`property`新增枚举值`chinese_8k_general`
  - 接口`RunTts`请求参数`property`新增枚举值`chinese_huaxiaoru_common`、`chinese_huaxiaohan_common`、`chinese_huaxiaoning_common`、`chinese_huaxiaozhen_common`、`english_alvin_common`、`english_amy_common`

# 0.1.12 2022-11-24

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - [Issue 58](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/58) 修复请求重试问题
- _特性变更_
  - 无

### HuaweiCloud SDK DWR

- _新增特性_
  - 支持数据工坊服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持接口`ListStackEvents`、`ListStackResources`、`DeleteStack`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK APM

- _新增特性_
  - 支持以下接口：
    - `ShowBusinessDetail`
    - `ShowSubBusinessDetail`
    - `ListAlarmData`
    - `ListAlarmNotify`
    - `SearchBusinessTopology`
    - `SearchEnvTopology`
    - `SearchTransactionConfig`
    - `ListBusinessEnv`
    - `SearchTransaction`
    - `ShowTransactionDetail`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`DeleteEnv`
  - 接口`ListAkSk`:
    - 响应参数`gmt_create`类型调整 `date` -> `string`
    - 响应参数`gmt_modify`类型调整 `date` -> `string`
  - 接口`SaveMonitorItemConfig`请求参数`monitor_item_id`、`env_id`改为必填
  - 接口`ListEnvMonitorItem`请求参数`x-business-id`
  - 接口`ShowTopologyTree`请求参数`business_id`改为非必填
  - 接口`ListEnvTags`:
    - 请求参数`business_id`改为必填
    - 响应参数`gmt_create`类型调整 `date` -> `string`
    - 响应参数`gmt_modify`类型调整 `date` -> `string`
  - 接口`ShowSpanSearch`请求参数`region`、`biz_id`改为必填
  - 接口`ShowRawTable`请求参数`view_type`、`collector_name`、`metric_set`、`title`、`table_direction`、`group_by`、`filter`、`span`、`span_field`、`page`、`page_size`、`search_word`、`instance_id`、`monitor_item_id`、`env_id`、`start_time`、`end_time`改为必填
  - 接口`ShowClobDetail`请求参数`env_id`、`clob_id`改为必填
  - 接口`ListEnvInstances`请求参数`env_id`、`page`、`page_size`改为必填
  - 接口`ShowAkSks`:
    - 响应参数`gmt_create`类型调整 `date` -> `string`
    - 响应参数`gmt_modify`类型调整 `date` -> `string`

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSubCustomerBillDetail`新增响应参数 `sub_service_type_code`、`sub_service_type_name`、`sub_resource_type_code`、`sub_resource_type_name`、`sub_resource_id`、`sub_resource_name`

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowBackup`新增响应参数 `children`
  - 接口`ListBackups`新增响应参数 `children`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`新增请求参数 `configurationsOverride`
  - 接口`ListClusters`新增响应参数 `configurationsOverride`
  - 接口`UpdateCluster`:
    - 新增请求参数 `eniNetwork`、`hostNetwork`
    - 新增响应参数 `configurationsOverride`
  - 接口`DeleteCluster`:
    - 新增请求参数 `delete_sfs30`
    - 新增响应参数 `configurationsOverride`
  - 接口`ShowCluster`新增响应参数 `configurationsOverride`
  - 接口`CreateNode`新增请求参数 `initializedConditions`
  - 接口`ListNodes`新增响应参数 `initializedConditions`
  - 接口`UpdateNode`新增响应参数 `initializedConditions`
  - 接口`DeleteNode`新增响应参数 `initializedConditions`
  - 接口`ShowNode`新增响应参数 `initializedConditions`
  - 接口`AddNode`新增请求参数 `initializedConditions`
  - 接口`ResetNode`新增请求参数 `initializedConditions`
  - 接口`CreateNodePool`新增请求参数 `customSecurityGroups`、`initializedConditions`
  - 接口`ListNodePools`新增响应参数 `customSecurityGroups`、`initializedConditions`
  - 接口`UpdateNodePool`:
    - 新增请求参数 `initializedConditions`
    - 新增响应参数 `customSecurityGroups`、`initializedConditions`
  - 接口`DeleteNodePool`新增响应参数 `customSecurityGroups`、`initializedConditions`
  - 接口`ShowNodePool`新增响应参数 `customSecurityGroups`、`initializedConditions`

### HuaweiCloud SDK CES

- _新增特性_
  - 支持以下接口：
    - `ListAlarmTemplates`
    - `CreateAlarmTemplate`
    - `BatchDeleteAlarmTemplates`
    - `ShowAlarmTemplate`
    - `UpdateAlarmTemplate`
    - `ListAlarmTemplateAssociationAlarms`
    - `ListResourceGroups`
    - `CreateResourceGroup`
    - `ShowResourceGroup`
    - `UpdateResourceGroup`
    - `ListResourceGroupsServicesResources`
    - `BatchDeleteResourceGroups`
    - `BatchCreateResources`
    - `BatchDeleteResources`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CSE

- _新增特性_
  - 支持接口`UpgradeEngine`、`RetryEngine`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CSS

- _新增特性_
  - 支持接口`UpdateInstance`、`ChangeMode`、`AddIndependentNode`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClustersDetails`新增响应参数 `totalSize`、`volume`
  - 接口`ShowClusterDetail`新增响应参数 `volume`
  - 接口`StartVpecp`请求参数`endpointWithDnsName`改为非必填

### HuaweiCloud SDK DAS

- _新增特性_
  - 支持以下接口：
    - `ShowSqlLimitSwitchStatus`
    - `ChangeSqlLimitSwitchStatus`
    - `ListSqlLimitRules`
    - `CreateSqlLimitRules`
    - `DeleteSqlLimitRules`
    - `ExportTopSqlTemplatesDetails`
    - `ShowSqlLimitJobInfo`
    - `ExportSlowSqlTemplatesDetails`
    - `ExportTopSqlTrendDetails`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowQuotas`响应参数`quotas`类型调整 `object` -> `array`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持以下接口：
    - `ListAppliedInstances`
    - `ShowConfigurationAppliedHistory`
    - `ShowConfigurationModifyHistory`
    - `CompareConfiguration`
    - `CopyConfiguration`
    - `ResetConfiguration`
    - `ListTasks`
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddReadonlyNode`新增请求参数 `is_auto_pay`

### HuaweiCloud SDK DNS

- _新增特性_
  - 支持以下接口：
    - `SetPrivateZoneProxyPattern`
    - `AssociateHealthCheck`
    - `DisassociateHealthCheck`
    - `CreateRetrieval`
    - `ShowRetrieval`
    - `CreateRetrievalVerification`
    - `ShowRetrievalVerification`
    - `CreateEndpoint`
    - `ShowEndpoint`
    - `ListEndpoints`
    - `UpdateEndpoint`
    - `DeleteEndpoint`
    - `AssociateEndpointIpaddress`
    - `ListEndpointIpaddresses`
    - `DisassociateEndpointIpaddress`
    - `ListEndpointVpcs`
    - `CreateResolveRule`
    - `ShowResoleRule`
    - `ListResoleRules`
    - `UpdateResolveRule`
    - `DeleteResolveRule`
    - `CreateLineGroup`
    - `ListLineGroups`
    - `ShowLineGroup`
    - `UpdateLineGroups`
    - `DeleteLineGroup`
    - `BatchDeleteZones`
    - `BatchDeletePtrRecords`
    - `BatchSetZonesStatus`
    - `BatchSetRecordSetsStatus`
    - `BatchDeleteRecordSets`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持以下接口：
    - `CreateDependencyVersion`
    - `ListDependencyVersion`
    - `ShowDependencyVersion`
    - `DeleteDependencyVersion`
    - `ListReservedInstanceConfigs`
    - `ListFunctionReservedInstances`
    - `ListFunctionAsMetric`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`AsyncInvokeReservedFunction`
  - 接口`UpdateFunctionConfig`新增请求参数 `custom_image`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持以下接口：
    - `ResizeColdVolume`
    - `CreateColdVolume`
    - `ModifyPublicIp`
    - `SwitchSsl`
    - `SetAutoEnlargePolicy`
    - `RestartInstance`
    - `ShowApplicableInstances`
    - `ShowModifyHistory`
    - `ShowApplyHistory`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数 `restore_info`、`port`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeHealthCode`新增响应参数 `test_interval`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ShowDomainName`、`ShowDnsName`、`ShowReplicationStatus`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ROMA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDatasourceInfo`:
    - 新增请求参数 `ssl`
    - 移除请求参数 `custom_plugin_id`
    - 请求参数`datasource_name`、`datasource_type`、`app_id`改为必填
  - 接口`ListDatasources`新增响应参数 `ssl`
  - 接口`UpdateDatasourceInfo`:
    - 新增请求参数 `ssl`
    - 新增响应参数 `ssl`
    - 移除请求参数 `custom_plugin_id`
    - 请求参数`datasource_name`、`datasource_type`、`app_id`改为必填
  - 接口`ShowDataourceDetail`新增响应参数 `ssl`
  - 接口`StartTestDatasource`:
    - 新增请求参数 `ssl`
    - 移除请求参数 `custom_plugin_id`
    - 请求参数`datasource_name`、`datasource_type`、`app_id`改为必填
  - 接口`CreateCommonTask`:
    - 移除响应参数 `created_by`
    - 请求参数`task_name`、`task_type`、`source_datasource_id`、`target_datasource_id`、`task_detail`改为必填
  - 接口`UpdateTask`:
    - 移除响应参数 `created_by`
    - 请求参数`task_name`、`task_type`、`source_datasource_id`、`target_datasource_id`、`task_detail`改为必填
  - 接口`ShowTask`移除响应参数 `created_by`
  - 接口`BatchStartOrStopTasks`请求参数`action_id`改为必填
  - 接口`CreateMultiTasks`请求参数`task_name`、`operation_types`改为必填
  - 接口`UpdateMultiTasks`请求参数`operation_types`、`repulling_snapshot`改为必填
  - 接口`CreateMultiTaskMappings`请求参数`source_datasource_id`、`target_datasource_id`改为必填

### HuaweiCloud SDK SA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ImportEvents`:
    - 新增请求参数 `dest_ip`、`cve_ids`
    - 移除请求参数 `destc_ip`
    - 请求参数`first_observed_time`类型调整 `string` -> `date-time`
    - 请求参数`last_observed_time`类型调整 `string` -> `date-time`
    - 请求参数`create_time`类型调整 `string` -> `date-time`
    - 请求参数`arrive_time`类型调整 `string` -> `date-time`
    - 请求参数`release_time`类型调整 `string` -> `date-time`
    - 请求参数`start_time`类型调整 `string` -> `date-time`
    - 请求参数`modified`类型调整 `string` -> `date-time`

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持接口`CreatePrepaidCloudWaf`、`ChangePrepaidCloudWaf`、`ShowSubscriptionInfo`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.11 2022-11-17

### HuaweiCloud SDK AOS

- _新增特性_
  - 支持应用编排服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DC

- _新增特性_
  - 支持云专线服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CFW

- _新增特性_
  - 支持云防火墙服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 支持以下接口：
    - `ListPlugins`
    - `CreatePlugin`
    - `ShowPlugin`
    - `UpdatePlugin`
    - `DeletePlugin`
    - `AttachApiToPlugin`
    - `DetachApiFromPlugin`
    - `ListPluginAttachedApis`
    - `ListPluginAttachableApis`
    - `AttachPluginToApi`
    - `DetachPluginFromApi`
    - `ListApiAttachedPlugins`
    - `ListApiAttachablePlugins`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateVpcChannelV2`新增请求参数 `type`、`microservice_labels`
  - 接口`ListVpcChannelsV2`新增响应参数 `microservice_labels`
  - 接口`ShowDetailsOfVpcChannelV2`新增响应参数 `microservice_labels`
  - 接口`UpdateVpcChannelV2`:
    - 新增请求参数 `type`、`microservice_labels`
    - 新增响应参数 `microservice_labels`
  - 接口`CreateMemberGroup`新增请求参数 `microservice_labels`
  - 接口`ListMemberGroups`新增响应参数 `microservice_labels`
  - 接口`ShowDetailsOfMemberGroup`新增响应参数 `microservice_labels`
  - 接口`UpdateMemberGroup`:
    - 新增请求参数 `microservice_labels`
    - 新增响应参数 `microservice_labels`
  - 接口`CreateInstanceV2`新增请求参数 `tags`

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DeleteServerNics`请求参数`id`改为必填

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateSubEnterpriseAccount`请求参数`sub_customer_association_type`改为必填

### HuaweiCloud SDK CloudArtifact

- _新增特性_
  - 支持接口`ShowProjectReleaseFiles`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudDeploy

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDeployTaskByTemplate`请求参数`type`新增枚举值`encrypt`
  - 接口`StartDeployTask`:
    - 新增请求参数 `trigger_source`、`key`
    - 移除请求参数 `description`、`name`、`limits`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持以下接口：
    - `AddExtensionEvaluation`
    - `AddExtensionEvaluationReply`
    - `CheckMaliciousExtensionEvaluation`
    - `DeleteEvaluationReply`
    - `DeleteEvaluation`
    - `AddExtensionStar`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ResizeInstance`请求参数`new_capacity`类型调整 `integer` -> `int32`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateJobs`请求参数`db_type`新增枚举值`gaussdbv5`、`postgresql`、`kafka`、`gaussdbv5ha`
  - 接口`BatchValidateConnections`请求参数`db_type`新增枚举值`gaussdbv5`、`kafka`、`gaussdbv5ha`
  - 接口`BatchUpdateJob`请求参数`db_type`新增枚举值`gaussdbv5`、`postgresql`、`kafka`、`gaussdbv5ha`
  - 接口`ListCompareResult`:
    - 新增响应参数 `line_compare_detail`、`content_compare_diff`、`compare_task_list`
    - 移除响应参数 `LineCompareDetail`、`ContentCompareDiff`、`CompareTaskList`
  - 接口`BatchListJobDetails`:
    - 新增响应参数 `is_multi_az`、`az_name`、`master_az`、`slave_az`、`node_role`
    - 响应参数`db_type`新增枚举值`gaussdbv5`、`postgresql`、`kafka`、`gaussdbv5ha`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RegisterServerMonitor`请求参数`monitorMetrics`类型调整 `string` -> `enum`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DisassociatePublicips`请求参数`associate_instance_type`新增枚举值`VPN`
  - 接口`AssociatePublicips`请求参数`associate_instance_type`新增枚举值`VPN`

### HuaweiCloud SDK EPS

- _新增特性_
  - 支持接口`ListProviders`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持接口`ShowInstanceRole`、`SwitchToMaster`、`SwitchToSlave`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateMessage`新增请求参数 `ttl`
  - 接口`ListCertificates`新增请求参数 `Sp-Auth-Token`、`Stage-Auth-Token`
  - 接口`AddCertificate`:
    - 新增请求参数 `Sp-Auth-Token`、`Stage-Auth-Token`、`addCertificateRequestBody`
    - 移除请求参数 `AddCertificateRequestBody`
  - 接口`DeleteCertificate`新增请求参数 `Sp-Auth-Token`、`Stage-Auth-Token`
  - 接口`CheckCertificate`:
    - 新增请求参数 `Sp-Auth-Token`、`Stage-Auth-Token`、`checkCertificateRequestBody`
    - 移除请求参数 `CheckCertificateRequestBody`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ResetPassword`请求参数`new_password`改为必填

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeGeneralTable`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeVatInvoice`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeInvoiceVerification`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeGeneralText`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeWebImage`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeHealthCode`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeQuotaInvoice`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeIdCard`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeHandwriting`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeVehicleLicense`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeTransportationLicense`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeTaxiInvoice`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeAutoClassification`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeTollInvoice`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeMvsInvoice`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeLicensePlate`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeFlightItinerary`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeBusinessLicense`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeDriverLicense`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeBusinessCard`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeTrainTicket`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeVin`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizePassport`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeBankcard`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeInsurancePolicy`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeFinancialStatement`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeQualificationCertificate`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeThailandIdcard`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeMyanmarIdcard`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeMyanmarDriverLicense`:
    - 新增请求参数 `Enterprise-Project-Id`
    - 新增响应参数 `birth`、`birth`
    - 移除响应参数 `Birth`、`Birth`
  - 接口`RecognizeChileIdCard`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeThailandLicensePlate`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeWaybillElectronic`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizePcrTestRecord`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeIdDocument`:
    - 新增请求参数 `Enterprise-Project-Id`
    - 响应参数`result`类型调整 `object` -> `object`
  - 接口`RecognizeHkIdCard`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeCambodianIdCard`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeExitEntryPermit`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeMainlandTravelPermit`新增请求参数 `Enterprise-Project-Id`
  - 接口`RecognizeMacaoIdCard`新增请求参数 `Enterprise-Project-Id`

### HuaweiCloud SDK OSM

- _新增特性_
  - 支持接口`RevokeMessage`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ROMA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDeviceGroup`移除响应参数 `app_name`
  - 接口`ShowDeviceGroupTree`移除响应参数 `total`、`app_id`、`permissions`
  - 接口`UpdateDeviceGroup`移除响应参数 `app_name`
  - 接口`CreateDevice`新增请求参数 `user_name`
  - 接口`ListDevices`:
    - 新增响应参数 `device_id`
    - 移除响应参数 `plugin_id`
  - 接口`BatchFreezeDevices`新增响应参数 `device_id`、`device_id`
  - 接口`UpdateDevice`:
    - 新增响应参数 `device_id`
    - 移除响应参数 `plugin_id`
  - 接口`ShowDevice`:
    - 新增响应参数 `device_id`
    - 响应参数`status`类型调整 `enum` -> `integer`
    - 响应参数`online_status`类型调整 `enum` -> `integer`
    - 响应参数`device_type`类型调整 `enum` -> `integer`
    - 响应参数`plugin_id`类型调整 `enum` -> `integer`
  - 接口`ListSubsets`:
    - 新增响应参数 `device_id`
    - 移除响应参数 `plugin_id`
  - 接口`ResetAuthentication`新增请求参数 `ResetAuthenticationRequestBody`
  - 接口`UpdateProduct`:
    - 新增响应参数 `status`
    - 移除响应参数 `authentication`
  - 接口`ResetProductAuthentication`新增请求参数 `ResetProductAuthenticationRequestBody`
  - 接口`CreateRule`移除响应参数 `app_name`、`sql_field`、`sql_where`、`rule_express`
  - 接口`CreateProperty`新增请求参数 `enum_dict`、`method`
  - 接口`ListProperties`:
    - 新增响应参数 `enum_dict`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`UpdateProperty`:
    - 新增请求参数 `enum_dict`
    - 新增响应参数 `enum_dict`、`method`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`ShowProperty`:
    - 新增响应参数 `enum_dict`、`method`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`CreateRequestProperty`:
    - 新增请求参数 `enum_dict`、`method`
    - 新增响应参数 `enum_dict`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`ListRequestProperties`:
    - 新增响应参数 `enum_dict`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`UpdateRequestProperty`:
    - 新增请求参数 `enum_dict`
    - 新增响应参数 `enum_dict`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`ShowRequestProperty`:
    - 新增响应参数 `enum_dict`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`CreateResponseProperty`新增请求参数 `enum_dict`、`method`
  - 接口`ListResponseProperties`:
    - 新增响应参数 `enum_dict`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`UpdateResponseProperty`:
    - 新增请求参数 `enum_dict`
    - 新增响应参数 `enum_dict`
    - 响应参数`data_type`新增枚举值`boolean`、`array`
  - 接口`ShowResponseProperty`:
    - 新增响应参数 `enum_dict`
    - 响应参数`data_type`新增枚举值`boolean`、`array`

### HuaweiCloud SDK TMS

- _新增特性_
  - 支持接口`ListProviders`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK UGO

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunSqlConversion`:
    - 请求参数`target_db_type`新增枚举值`GaussDB Centralized`, 移除枚举值`GaussDB(for openGauss)`
    - 请求参数`target_db_version`新增枚举值`2.0`, 移除枚举值`2020`
  - 接口`ConfirmTargetDbType`:
    - 请求参数`target_db_type`类型调整 `string` -> `enum`
    - 请求参数`target_db_version`类型调整 `string` -> `enum`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateSubnet`:
    - 请求参数`opt_name`新增枚举值`addresstime`
    - 响应参数`opt_name`新增枚举值`addresstime`
  - 接口`ListSubnets`响应参数`opt_name`新增枚举值`addresstime`
  - 接口`ShowSubnet`响应参数`opt_name`新增枚举值`addresstime`
  - 接口`UpdateSubnet`请求参数`opt_name`新增枚举值`addresstime`

# 0.1.10 2022-11-14

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DeleteServerNics`响应参数`job_id`改为必填
  - 接口`UpdateBaremetalServerInterfaceAttachments`:
    - 请求参数`delete_on_termination`类型调整 `string` -> `boolean`
    - 请求参数`delete_on_termination`改为必填
  - 接口`ShowServerRemoteConsole`新增响应参数 `remote_console`

### HuaweiCloud SDK CPH

- _新增特性_
  - 支持接口`DeleteShareFiles`
- _解决问题_
  - 无
- _特性变更_
  - 接口`StopCloudPhone`:
    - 新增请求参数 `StopCloudPhoneRequestBody`
    - 移除请求参数 `StopCloudPhoneReuqestBody`
  - 接口`ShowCloudPhoneServerDetail`:
    - 新增响应参数 `server_name`、`availability_zone`、`server_id`、`server_model_name`、`phone_model_name`、`keypair_name`、`status`、`vpc_id`、`cidr`、`vpc_cidr`、`subnet_id`、`subnet_cidr`、`resource_project_id`、`metadata`、`addresses`、`network_version`、`create_time`、`update_time`
    - 移除响应参数 `servers`
  - 接口`CreateNet2CloudPhoneServer`移除请求参数 `br_cidr`

### HuaweiCloud SDK CSMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSecretTags`新增响应参数 `sys_tags`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持以下接口：
    - `ShowSlowlogDesensitizationSwitch`
    - `ListRecycleInstances`
    - `CheckWeakPassword`
    - `ShowUpgradeDuration`
    - `ShowDiskUsage`
    - `ListSslCertDownloadAddress`
    - `DeleteAuditLog`
    - `ShowRecyclePolicy`
- _解决问题_
  - 无
- _特性变更_
  - 接口`SwitchSlowlogDesensitization`新增请求参数 `X-Language`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持接口`ShowWorkflowExecutionForPage`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListWorkflow`新增响应参数 `enable_stream_response`
  - 接口`UpdateWorkFlow`新增响应参数 `enable_stream_response`

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持接口`ListBackPools`、`ListBackPoolMembers`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListStreamForbidden`移除请求参数 `specify_project`
  - 接口`DeleteStreamForbidden`移除请求参数 `specify_project`
  - 接口`UpdateStreamForbidden`移除请求参数 `specify_project`
  - 接口`CreateStreamForbidden`移除请求参数 `specify_project`
  - 接口`ShowDomain`响应参数`service_area`移除枚举值`global`
  - 接口`UpdateDomain`响应参数`service_area`移除枚举值`global`
  - 接口`CreateDomain`请求参数`service_area`移除枚举值`global`
  - 接口`DeleteDomainMapping`移除请求参数 `specify_project`
  - 接口`CreateDomainMapping`移除请求参数 `specify_project`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListHosts`新增响应参数 `resource_id`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ShowPostgresqlParamValue`、`UpdatePostgresqlParameterValue`、`ListDrRelations`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`请求参数`type`新增枚举值`ESSD`
  - 接口`ListInstances`响应参数`type`新增枚举值`ESSD`
  - 接口`CreateRestoreInstance`请求参数`type`新增枚举值`ESSD`
  - 接口`CreatePostgresqlDatabase`新增请求参数 `is_revoke_public_privilege`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEndpoints`响应参数`Action`类型调整 `string` -> `array`
  - 接口`DeleteEndpointPolicy`响应参数`Action`类型调整 `string` -> `array`
  - 接口`UpdateEndpointPolicy`:
    - 请求参数`Action`类型调整 `string` -> `array`
    - 响应参数`Action`类型调整 `string` -> `array`

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持接口`MigrateCompositeHosts`、`ShowSourceIp`、`ListNoticeConfigs`、`UpdateAlertNoticeConfig`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstance`新增响应参数 `instance_name`
  - 接口`ShowLtsInfoConfig`新增响应参数 `enabled`、`ltsAttackStreamID`
  - 接口`UpdateLtsInfoConfig`:
    - 新增请求参数 `enabled`、`ltsAttackStreamID`
    - 新增响应参数 `enabled`、`ltsAttackStreamID`
    - 请求参数`enabale`改为非必填
  - 接口`ShowIpGroup`新增响应参数 `description`

# 0.1.9 2022-11-08

### HuaweiCloud SDK HSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListQuotasDetail`:
    - 新增响应参数 `on_demand_num`
    - 移除响应参数 `on_demand_numn`

### HuaweiCloud SDK Meeting

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SearchMeetings`新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`CreateMeeting`:
    - 新增请求参数 `isHostCameraOn`、`isGuestCameraOn`
    - 新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`UpdateMeeting`:
    - 新增请求参数 `isHostCameraOn`、`isGuestCameraOn`
    - 新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`ShowMeetingDetail`新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`SearchOnlineMeetings`新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`ShowOnlineMeetingDetail`新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`SearchHisMeetings`新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`ShowHisMeetingDetail`新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`CreateRecurringMeeting`:
    - 新增请求参数 `isHostCameraOn`、`isGuestCameraOn`
    - 新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`
  - 接口`UpdateRecurringMeeting`:
    - 新增请求参数 `isHostCameraOn`、`isGuestCameraOn`
    - 新增响应参数 `onlineAttendeeAmount`、`isHostCameraOn`、`isGuestCameraOn`

# 0.1.8 2022-11-03

### HuaweiCloud SDK DevSecurity

- _新增特性_
  - 支持研发安全服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GA

- _新增特性_
  - 支持全球加速服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudDeploy

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDeployTasks`新增响应参数 `id`、`release_id`、`duration`、`execution_state`、`executor_id`、`executor_nick_name`、`steps`
  - 接口`ListHosts`新增响应参数 `update_time`、`lastest_connection_time`、`connection_status`、`owner_name`、`updator_id`、`create_time`、`nick_name`、`owner_id`、`updator_name`、`connection_result`
  - 接口`ListHostGroups`:
    - 新增响应参数 `updated_by`
    - 移除响应参数 `update_by`
  - 接口`ShowDeploymentGroupDetail`:
    - 新增响应参数 `updated_by`
    - 移除响应参数 `update_by`
  - 接口`ShowDeploymentHostDetail`新增响应参数 `update_time`、`lastest_connection_time`、`connection_status`、`owner_name`、`updator_id`、`create_time`、`nick_name`、`owner_id`、`updator_name`、`connection_result`
  - 接口`ShowDeployTaskDetail`新增响应参数 `id`、`release_id`、`duration`、`execution_state`、`executor_id`、`executor_nick_name`、`steps`
  - 接口`ListDeployTaskHistoryByDate`新增响应参数 `type`
  - 接口`ShowProjectSuccessRate`:
    - 新增响应参数 `start_date`、`end_date`
    - 移除响应参数 `start_time`、`end_time`
  - 接口`ListTaskSuccessRate`:
    - 新增响应参数 `start_date`、`end_date`
    - 移除响应参数 `start_time`、`end_time`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClustersDetails`:
    - 新增响应参数 `failedReason`
    - 移除响应参数 `failed_reasons`
  - 接口`ShowClusterDetail`:
    - 新增响应参数 `failedReason`
    - 移除响应参数 `failed_reasons`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchValidateConnections`响应参数`status`新增枚举值`true`、`false`
  - 接口`BatchValidateClustersConnections`响应参数`status`新增枚举值`true`、`false`
  - 接口`BatchCheckResults`响应参数`job_direction`新增枚举值`up`、`down`、`non-dbs`
  - 接口`BatchSetSpeed`响应参数`status`新增枚举值`success`、`failed`
  - 接口`ListCompareResult`:
    - 响应参数`object_type`新增枚举值`DB`、`TABLE`、`VIEW`、`EVENT`、`ROUTINE`、`INDEX`、`TRIGGER`、`SYNONYM`、`FUNCTION`、`PROCEDURE`、`TYPE`、`RULE`、`DEFAULT_TYPE`、`PLAN_GUIDE`、`CONSTRAINT`、`FILE_GROUP`、`PARTITION_FUNCTION`、`PARTITION_SCHEME`、`TABLE_COLLATION`
    - 响应参数`object_compare_result`新增枚举值`CONSISTENT`、`INCONSISTENT`、`COMPARING`、`WAITING_FOR_COMPARISON`、`FAILED_TO_COMPARE`、`TARGET_DB_NOT_EXIT`、`CAN_NOT_COMPARE`
    - 响应参数`line_compare_result`新增枚举值`CONSISTENT`、`INCONSISTENT`、`COMPARING`、`WAITING_FOR_COMPARISON`、`FAILED_TO_COMPARE`、`TARGET_DB_NOT_EXIT`、`CAN_NOT_COMPARE`
    - 响应参数`content_compare_result`新增枚举值`CONSISTENT`、`INCONSISTENT`、`COMPARING`、`WAITING_FOR_COMPARISON`、`FAILED_TO_COMPARE`、`TARGET_DB_NOT_EXIT`、`CAN_NOT_COMPARE`
    - 响应参数`compare_task_status`新增枚举值`RUNNING`、`WAITING_FOR_RUNNING`、`SUCCESSFUL`、`FAILED`、`CANCELLED`、`TIMEOUT_INTERRUPT`、`FULL_DOING`、`INCRE_DOING`
  - 接口`BatchListProgresses`响应参数`task_mode`新增枚举值`FULL_TRANS`、`FULL_INCR_TRANS`、`INCR_TRANS`
  - 接口`BatchListJobDetails`响应参数`status`新增枚举值`CREATING`、`CREATE_FAILED`、`CONFIGURATION`、`STARTJOBING`、`WAITING_FOR_START`、`START_JOB_FAILED`、`PAUSING`、`FULL_TRANSFER_STARTED`、`FULL_TRANSFER_FAILED`、`FULL_TRANSFER_COMPLETE`、`INCRE_TRANSFER_STARTED`、`INCRE_TRANSFER_FAILED`、`RELEASE_RESOURCE_STARTED`、`RELEASE_RESOURCE_FAILED`、`RELEASE_RESOURCE_COMPLETE`、`REBUILD_NODE_STARTED`、`REBUILD_NODE_FAILED`、`CHANGE_JOB_STARTED`、`CHANGE_JOB_FAILED`、`DELETED`、`CHILD_TRANSFER_STARTING`、`CHILD_TRANSFER_STARTED`、`CHILD_TRANSFER_COMPLETE`、`CHILD_TRANSFER_FAILED`、`RELEASE_CHILD_TRANSFER_STARTED`、`RELEASE_CHILD_TRANSFER_COMPLETE`、`NODE_UPGRADE_START`、`NODE_UPGRADE_COMPLETE`、`NODE_UPGRADE_FAILED`
  - 接口`ShowJobList`:
    - 请求参数`status`新增枚举值`PAUSING`、`REBUILD_NODE_STARTED`、`REBUILD_NODE_FAILED`、`DELETED`、`NODE_UPGRADE_START`、`NODE_UPGRADE_COMPLETE`、`NODE_UPGRADE_FAILED`
    - 响应参数`status`新增枚举值`CREATING`、`CREATE_FAILED`、`CONFIGURATION`、`STARTJOBING`、`WAITING_FOR_START`、`START_JOB_FAILED`、`PAUSING`、`FULL_TRANSFER_STARTED`、`FULL_TRANSFER_FAILED`、`FULL_TRANSFER_COMPLETE`、`INCRE_TRANSFER_STARTED`、`INCRE_TRANSFER_FAILED`、`RELEASE_RESOURCE_STARTED`、`RELEASE_RESOURCE_FAILED`、`RELEASE_RESOURCE_COMPLETE`、`REBUILD_NODE_STARTED`、`REBUILD_NODE_FAILED`、`CHANGE_JOB_STARTED`、`CHANGE_JOB_FAILED`、`DELETED`、`CHILD_TRANSFER_STARTING`、`CHILD_TRANSFER_STARTED`、`CHILD_TRANSFER_COMPLETE`、`CHILD_TRANSFER_FAILED`、`RELEASE_CHILD_TRANSFER_STARTED`、`RELEASE_CHILD_TRANSFER_COMPLETE`、`NODE_UPGRADE_START`、`NODE_UPGRADE_COMPLETE`、`NODE_UPGRADE_FAILED`
    - 响应参数`db_use_type`新增枚举值`migration`、`sync`、`cloudDataGuard`
    - 响应参数`task_type`新增枚举值`FULL_TRANS`、`FULL_INCR_TRANS`、`INCR_TRANS`
    - 响应参数`status`类型调整 `string` -> `enum`
  - 接口`BatchListJobStatus`响应参数`status`新增枚举值`CREATING`、`CREATE_FAILED`、`CONFIGURATION`、`STARTJOBING`、`WAITING_FOR_START`、`START_JOB_FAILED`、`PAUSING`、`FULL_TRANSFER_STARTED`、`FULL_TRANSFER_FAILED`、`FULL_TRANSFER_COMPLETE`、`INCRE_TRANSFER_STARTED`、`INCRE_TRANSFER_FAILED`、`RELEASE_RESOURCE_STARTED`、`RELEASE_RESOURCE_FAILED`、`RELEASE_RESOURCE_COMPLETE`、`REBUILD_NODE_STARTED`、`REBUILD_NODE_FAILED`、`CHANGE_JOB_STARTED`、`CHANGE_JOB_FAILED`、`DELETED`、`CHILD_TRANSFER_STARTING`、`CHILD_TRANSFER_STARTED`、`CHILD_TRANSFER_COMPLETE`、`CHILD_TRANSFER_FAILED`、`RELEASE_CHILD_TRANSFER_STARTED`、`RELEASE_CHILD_TRANSFER_COMPLETE`、`NODE_UPGRADE_START`、`NODE_UPGRADE_COMPLETE`、`NODE_UPGRADE_FAILED`
  - 接口`BatchSwitchover`:
    - 响应参数`mongo_ha_mode`新增枚举值`Sharding`、`ReplicaSet`、`ReplicaSingle`
    - 响应参数`cluster_mode`新增枚举值`Single`、`Ha`、`GR`、`Sharding`、`Sharding4.0+`、`ReplicaSet`、`Replica`、`ReplicaSingle`、`Cluster`、`Independent`、`Combined`、`Distributed`、`NoSharding`
    - 响应参数`job_direction`新增枚举值`up`、`down`、`non-dbs`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`NovaCreateServers`请求参数`destination_type`改为非必填

### HuaweiCloud SDK EPS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`MigrateResource`新增请求参数 `region_id`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowGaussMySqlProxy`
  - 接口`CreateGaussMySqlInstance`新增请求参数 `lower_case_table_names`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeGeneralText`:
    - 新增请求参数 `character_mode`
    - 新增响应参数 `confidence`、`char_list`
  - 接口`RecognizeThailandIdcard`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持以下接口：
    - `ListInstance`
    - `CreateInstance`
    - `ShowInstance`
    - `RenameInstance`
    - `DeleteInstance`
    - `ShowLtsInfoConfig`
    - `UpdateLtsInfoConfig`
    - `ListIpGroup`
    - `CreateIpGroup`
    - `ShowIpGroup`
    - `UpdateIpGroup`
    - `DeleteIpGroup`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.7 2022-11-02

### HuaweiCloud SDK APIG

- _新增特性_
  - 支持以下接口：
    - `UpdateHealthCheck`
    - `BatchEnableMembers`
    - `BatchDisableMembers`
    - `ListMemberGroups`
    - `CreateMemberGroup`
    - `ShowDetailsOfMemberGroup`
    - `UpdateMemberGroup`
    - `DeleteMemberGroup`
    - `ListMetricData`
    - `ImportMicroservice`
    - `ListCertificatesV2`
    - `CreateCertificateV2`
    - `BatchAssociateCertsV2`
    - `BatchDisassociateCertsV2`
    - `ShowDetailsOfCertificateV2`
    - `UpdateCertificateV2`
    - `DeleteCertificateV2`
    - `BatchAssociateDomainsV2`
    - `BatchDisassociateDomainsV2`
    - `ListAttachedDomainsV2`
    - `UpdateBackendInstancesV2`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateAclStrategyV2`:
    - 请求参数`acl_type`类型调整 `string` -> `enum`
    - 请求参数`entity_type`类型调整 `string` -> `enum`
  - 接口`UpdateAclStrategyV2`:
    - 请求参数`acl_type`类型调整 `string` -> `enum`
    - 请求参数`entity_type`类型调整 `string` -> `enum`
  - 接口`ListAclPolicyBindedToApiV2`响应参数`entity_type`新增枚举值`DOMAIN_ID`
  - 接口`CreateVpcChannelV2`:
    - 新增请求参数 `member_groups`、`microservice_info`、`dict_code`
    - 请求参数`port`、`balance_strategy`、`member_type`改为必填
  - 接口`ListVpcChannelsV2`:
    - 新增请求参数 `dict_code`、`member_host`、`member_port`、`member_group_name`、`member_group_id`
    - 新增响应参数 `microservice_info`、`type`、`dict_code`、`microservice_version`、`microservice_port`
    - 响应参数`port`、`balance_strategy`、`member_type`改为必填
  - 接口`ShowDetailsOfVpcChannelV2`:
    - 新增响应参数 `microservice_info`、`type`、`dict_code`、`microservice_version`、`microservice_port`
    - 响应参数`port`、`balance_strategy`、`member_type`改为必填
  - 接口`UpdateVpcChannelV2`:
    - 新增请求参数 `member_groups`、`microservice_info`、`dict_code`
    - 新增响应参数 `microservice_info`、`type`、`dict_code`、`microservice_version`、`microservice_port`
    - 请求参数`port`、`balance_strategy`、`member_type`改为必填
    - 响应参数`port`、`balance_strategy`、`member_type`改为必填
  - 接口`ListBackendInstancesV2`新增请求参数 `member_group_name`、`member_group_id`、`precise_search`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowUrlTaskInfo`:
    - 响应参数`modify_time`类型调整 `int32` -> `int64`
    - 响应参数`create_time`类型调整 `int32` -> `int64`

### HuaweiCloud SDK ECS

- _新增特性_
  - 支持接口`UpdateServerBlockDevice`、`RegisterServerMonitor`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`UpdateTransactionSplitStatus`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowGaussMySqlProxy`新增响应参数 `transaction_split`
  - 接口`ShowGaussMySqlProxyList`新增响应参数 `transaction_split`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateScalingPolicy`:
    - 新增请求参数 `autoScalingPolicyReqV11`
    - 移除请求参数 `auto_scaling_policy_req_V11`
  - 接口`CreateCluster`:
    - 新增请求参数 `createClusterReqV11`
    - 移除请求参数 `cluster_req`
  - 接口`ListClusters`移除响应参数 `name`、`uri`、`parameters`、`nodes`、`active_master`、`fail_action`、`before_component_start`、`start_time`、`state`
  - 接口`UpdateClusterScaling`:
    - 新增请求参数 `clusterScalingReq`
    - 移除请求参数 `cluster_scaling_req`
  - 接口`ShowClusterDetails`移除响应参数 `name`、`uri`、`parameters`、`nodes`、`active_master`、`fail_action`、`before_component_start`、`start_time`、`state`
  - 接口`CreateClusterTag`:
    - 新增请求参数 `createTagReq`
    - 移除请求参数 `CreateTagRequest`
  - 接口`BatchCreateClusterTags`:
    - 新增请求参数 `batchCreateClusterTagsReq`
    - 移除请求参数 `BatchCreateClusterTagsRequest`
  - 接口`BatchDeleteClusterTags`移除请求参数 `key`、`value`
  - 接口`ListClustersByTags`:
    - 新增请求参数 `listResourceReq`
    - 移除请求参数 `ListResourceRequest`
    - 接口`CreateCluster`移除请求参数 `key`、`value`、`auto_scaling_enable`、`min_capacity`、`max_capacity`、`resources_plans`、`rules`、`exec_scripts`、`name`、`uri`、`parameters`、`nodes`、`active_master`、`fail_action`、`before_component_start`、`action_stages`、`job_type`、`job_name`、`jar_path`、`arguments`、`input`、`output`、`job_log`、`hive_script_path`、`hql`、`shutdown_cluster`、`submit_job_once_cluster_run`、`file_action`
  - 接口`CreateExecuteJob`:
    - 新增请求参数 `jobExecution`
    - 移除请求参数 `job_execution_dto`
    - 移除响应参数 `job_submit_result`
  - 接口`BatchDeleteJobs`:
    - 新增请求参数 `jobBatchDelete`
    - 移除请求参数 `job_batch_delete`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListPostgresqlExtension`、`CreatePostgresqlExtension`、`DeletePostgresqlExtension`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`PushTranscriberJobs`请求参数`property`新增枚举值`sichuan_8k_common`
  - 接口`RunTts`请求参数`property`新增枚举值`chinese_huaxiaolong_common`、`chinese_huaxiaorui_common`

# 0.1.6 2022-10-27

### HuaweiCloud SDK BMS

- _新增特性_
  - 支持接口`DeleteServerNics`、`UpdateBaremetalServerInterfaceAttachments`、`AddServerNics`、`ShowServerRemoteConsole`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowUrlTaskInfo`响应参数`id`类型调整 `int32` -> `int64`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`ListExtensions`、`ShowExtensionDetail`、`ShowExtensionEvaluation`、`ShowExtensionEvaluationStar`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListMergeRequest`:
    - 新增响应参数 `merge_requests`
    - 移除响应参数 `mergeRequests`
  - 接口`ShowMergeRequest`:
    - 新增响应参数 `approval_merge_request_approvers`、`closed_at`、`created_at`、`devcloud_source_branch`、`is_source_branch_exist`、`merge_request_assignee_list`、`merge_request_diff`、`merge_status`、`source_branch`、`target_branch`、`updated_at`
    - 移除响应参数 `approvalMergeRequestApprovers`、`closedAt`、`createdAt`、`devcloudSourceBranch`、`isSourceBranchExist`、`mergeRequestAssigneeList`、`mergeRequestDiff`、`mergeStatus`、`sourceBranch`、`targetBranch`、`updatedAt`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowInstance`:
    - 响应参数`begin_time`类型调整 `string` -> `int64`
    - 响应参数`end_time`类型调整 `string` -> `int64`
    - 响应参数`current_time`类型调整 `string` -> `int64`
    - 响应参数`next_expand_time`类型调整 `string` -> `int64`
    - 响应参数`expand_effect_time`类型调整 `string` -> `int64`
    - 响应参数`expand_interval_time`类型调整 `string` -> `int64`
  - 接口`ResizeInstance`请求参数`new_capacity`类型调整 `int32` -> `integer`
  - 接口`ListMigrationTask`新增响应参数 `target_instance_address`、`migration_method`、`task_name`、`target_instance_id`、`source_instance_name`、`target_instance_name`、`migrate_type`、`created_at`、`source_instance_id`、`task_id`、`data_source`、`status`
  - 接口`ListRedislog`:
    - 新增响应参数 `backup_id`
    - 移除响应参数 `backupId`
  - 接口`ListBackgroundTask`新增响应参数 `enable_show`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持以下接口：
    - `AddReadonlyNode`
    - `UpgradeDatabaseVersion`
    - `ShowSecondLevelMonitoringStatus`
    - `SwitchSecondLevelMonitoring`
    - `ChangeOpsWindow`
    - `SetRecyclePolicy`
    - `ExpandReplicasetNode`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListConfigurations`新增响应参数 `node_type`
  - 接口`ListInstances`新增响应参数 `patch_available`
  - 接口`ResizeInstanceVolume`新增请求参数 `node_ids`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateServers`新增请求参数 `X-Client-Token`、`batch_create_in_multi_az`
  - 接口`CreatePostPaidServers`新增请求参数 `X-Client-Token`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateListener`请求参数`tls_ciphers_policy`新增枚举值`tls-1-1`、`tls-1-2`、`tls-1-2-strict`, 移除枚举值` tls-1-1`、` tls-1-2`、` tls-1-2-strict`
  - 接口`DeleteListener`移除请求参数 `cascade`
  - 接口`DeleteLoadbalancer`移除请求参数 `cascade`
  - 接口`ListApiVersions`:
    - 新增响应参数 `versions`
    - 移除响应参数 `id`、`status`
  - 接口`CreateLoadBalancer`移除请求参数 `global_eip_ids`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持接口`ShowRestorableList`、`ListRestoreTime`、`DeleteBackup`、`RestoreExistingInstance`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK HSS

- _新增特性_
  - 支持以下接口：
    - `ListUsers`
    - `ListUserChangeHistories`
    - `ShowResourceQuotas`
    - `ListQuotasDetail`
    - `SwitchHostsProtectStatus`
    - `BatchCreateTags`
    - `DeleteResourceInstanceTag`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRiskConfigs`:
    - 新增请求参数 `check_name`
    - 移除请求参数 `check_type`
  - 接口`ShowCheckRuleDetail`:
    - 新增请求参数 `check_name`
    - 移除请求参数 `check_type`
  - 接口`ListHostStatus`:
    - 新增请求参数 `server_group`
    - 请求参数`region`改为非必填
  - 接口`ListVulnerabilities`:
    - 新增响应参数 `severity_level`
    - 响应参数`repair_necessity`类型调整 `int32` -> `string`

### HuaweiCloud SDK IEF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowEdgeNodeUpgradeDetails`
  - 接口`UpdateNodeByDeviceId`新增请求参数 `ief-instance-id`
  - 接口`ListPods`移除响应参数 `affinity`、`updated_at`
  - 接口`CreateEncryptdatas`移除响应参数 `encrypt_data`
  - 接口`ListEncryptdatas`:
    - 新增响应参数 `encrypt_datas`
    - 移除响应参数 `encrypt_data`
  - 接口`UpdateEncryptdatas`移除响应参数 `encrypt_data`
  - 接口`ListNodeEncryptdatas`:
    - 新增响应参数 `encrypt_datas`
    - 移除响应参数 `encrypt_data`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTags`请求参数`__imagetype`新增枚举值`market`
  - 接口`GlanceListImages`:
    - 请求参数`__imagetype`新增枚举值`market`
    - 响应参数`__imagetype`新增枚举值`market`
  - 接口`GlanceShowImage`响应参数`__imagetype`新增枚举值`market`
  - 接口`GlanceUpdateImage`响应参数`__imagetype`新增枚举值`market`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateMessage`新增请求参数 `properties`
  - 接口`ListDeviceMessages`新增响应参数 `properties`
  - 接口`ShowDeviceMessage`新增响应参数 `properties`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePostPaidInstance`:
    - 新增请求参数 `disk_encrypted_enable`、`disk_encrypted_key`
    - 请求参数`engine_version`新增枚举值`2.7`

### HuaweiCloud SDK Meeting

- _新增特性_
  - 支持以下接口：
    - `SearchCorpExternalDir`
    - `SetCohost`
    - `BatchHand`
    - `CancelBroadcast`
    - `AllowWaitingParticipant`
    - `MoveToWaitingRoom`
    - `AllowClientRecord`
    - `ShowLayout`
    - `SaveLayout`
    - `DeleteLayout`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunCreateVideoModerationJob`的请求参数`frame_interval`类型变更 `float` -> `integer`
  - 接口`RunQueryAudioModerationJob`的响应参数`start_time`、`end_time`类型变更 `integer` -> `float`
  - 接口`RunQueryVideoModerationJob`
    - 响应参数`time`类型变更 `integer` -> `float`
    - 响应参数`start_time`、`end_time`类型变更 `integer` -> `float`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `read_only_by_user`
  - 接口`CreateDbUser`新增请求参数 `hosts`、`databases`

### HuaweiCloud SDK SCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCertificates`响应参数`enterprise_project_id`改为必填
  - 接口`ShowCertificate`新增响应参数 `fingerprint`

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 支持接口`ChangeShareName`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.1.5 2022-09-28

### HuaweiCloud SDK APM

- _新增特性_
  - 支持接口`DeleteEnv`、`DeleteApp`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateAkSk`新增响应参数 `sk`
  - 接口`DeleteAkSk`新增响应参数 `sk`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDomains`新增响应参数 `domain_id`
  - 接口`CreateDomain`:
    - 新增请求参数 `domain_id`
    - 新增响应参数 `domain_id`
  - 接口`ShowDomainDetail`新增响应参数 `domain_id`
  - 接口`DeleteDomain`新增响应参数 `domain_id`
  - 接口`EnableDomain`新增响应参数 `domain_id`
  - 接口`DisableDomain`新增响应参数 `domain_id`
  - 接口`UpdateDomainOrigin`:
    - 新增请求参数 `domain_id`
    - 新增响应参数 `domain_id`
  - 接口`ShowDomainFullConfig`新增响应参数 `origin_range_status`、`user_agent_filter`、`origin_request_url_rewrite`、`error_code_redirect_rules`
  - 接口`UpdateDomainFullConfig`新增请求参数 `origin_range_status`、`user_agent_filter`、`origin_request_url_rewrite`、`error_code_redirect_rules`

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 支持接口`RemoveUsers`、`RemoveRoom`、`UpdateIndividualStreamJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListVariables`新增响应参数 `variable_mode`、`share_mode`
  - 接口`UpdateTask`:
    - 新增请求参数 `name`、`conditions`、`is_disabled`
    - 新增响应参数 `taskInfo`
    - 移除请求参数 `case_name`
    - 移除响应参数 `taskinfo`
    - 请求参数`check_end_length`类型调整 `string` -> `object`
    - 请求参数`check_end_str`类型调整 `string` -> `object`
    - 请求参数`check_end_type`类型调整 `string` -> `object`
  - 接口`ShowTask`:
    - 新增响应参数 `taskInfo`
    - 移除响应参数 `taskinfo`
  - 接口`ShowReport`新增响应参数 `respTimeRange`、`progressState`、`createBy`、`statusValue`
  - 接口`UpdateCase`新增请求参数 `case_id`、`name`、`case_type`、`increase_setting`、`stages`、`status`、`temp_id`、`sort`
  - 接口`UpdateTemp`:
    - 请求参数`check_end_length`类型调整 `string` -> `object`
    - 请求参数`check_end_str`类型调整 `string` -> `object`
    - 请求参数`check_end_type`类型调整 `string` -> `object`
  - 接口`UpdateTaskRelatedTestCase`:
    - 新增响应参数 `taskInfo`
    - 移除响应参数 `taskinfo`
  - 接口`ShowHistoryRunInfo`新增响应参数 `end_time`、`parallel`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowInstance`:
    - 响应参数`begin_time`类型调整 `string` -> `int64`
    - 响应参数`end_time`类型调整 `string` -> `int64`
    - 响应参数`current_time`类型调整 `string` -> `int64`
    - 响应参数`next_expand_time`类型调整 `string` -> `int64`
    - 响应参数`expand_effect_time`类型调整 `string` -> `int64`
    - 响应参数`expand_interval_time`类型调整 `string` -> `int64`
  - 接口`ResizeInstance`请求参数`new_capacity`类型调整 `int32` -> `integer`
  - 接口`ListMigrationTask`新增响应参数 `target_instance_address`、`migration_method`、`task_name`、`target_instance_id`、`source_instance_name`、`target_instance_name`、`migrate_type`、`created_at`、`source_instance_id`、`task_id`、`data_source`、`status`
  - 接口`ListRedislog`:
    - 新增响应参数 `backup_id`
    - 移除响应参数 `backupId`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`请求参数`region`改为必填

### HuaweiCloud SDK IEF

- _新增特性_
  - 支持以下接口：
    - `RestartDeploymentsPod`
    - `ShowEdgeNodeUpgradeDetails`
    - `UpgradeEdgeNode`
    - `ListEncryptdatas`
    - `CreateEncryptdatas`
    - `ShowEncryptdatas`
    - `UpdateEncryptdatas`
    - `DeleteEncryptdatas`
    - `ListNodeEncryptdatas`
    - `CreateNodeEncryptdatas`
    - `DeleteNodeEncryptdatas`
    - `ListEncryptdataNodes`
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateEdgeNode`:
    - 新增请求参数 `UpdateEdgeNodeRequestBody`
    - 移除请求参数 `UpdateEdgeNodeBody`
  - 接口`EnableDisableEdgeNodes`:
    - 新增请求参数 `EnableDisableEdgeNodesRequestBody`
    - 移除请求参数 `node`
  - 接口`ListApps`新增响应参数 `host_pid`
  - 接口`UpdateApp`新增响应参数 `host_pid`
  - 接口`ShowAppDetail`新增响应参数 `host_pid`
  - 接口`CreateAppVersions`新增请求参数 `host_pid`
  - 接口`ListAppVersions`新增响应参数 `host_pid`
  - 接口`UpdateAppVersion`:
    - 新增请求参数 `host_pid`
    - 新增响应参数 `host_pid`
  - 接口`ShowAppVersionDetail`新增响应参数 `host_pid`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTranscodingsTemplate`响应参数`width`、`height`改为非必填
  - 接口`UpdateTranscodingsTemplate`:
    - 新增请求参数 `trans_type`
    - 请求参数`width`、`height`改为非必填
  - 接口`CreateTranscodingsTemplate`:
    - 新增请求参数 `trans_type`
    - 请求参数`width`、`height`改为非必填

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeHealthCode`新增响应参数 `type`、`idcard_number`、`phone_number`、`province`、`city`、`vaccination_status`、`pcr_test_result`、`pcr_test_organization`、`pcr_test_time`、`pcr_sampling_time`、`reached_city`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 支持以下接口：
    - `UpdateEndpointServiceName`
    - `UpdateEndpointConnectionsDesc`
    - `BatchAddEndpointServicePermissions`
    - `BatchRemoveEndpointServicePermissions`
    - `UpdateEndpointServicePermissionDesc`
    - `UpdateEndpointPolicy`
    - `DeleteEndpointPolicy`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEndpointService`:
    - 新增请求参数 `description`
    - 新增响应参数 `description`
  - 接口`ListEndpointService`:
    - 新增请求参数 `public_border_group`
    - 新增响应参数 `description`、`public_border_group`
    - 响应参数`service_type`类型调整 `string` -> `enum`
    - 响应参数`server_type`类型调整 `enum` -> `string`
  - 接口`UpdateEndpointService`:
    - 新增请求参数 `description`
    - 新增响应参数 `description`
  - 接口`ListServiceDetails`:
    - 新增响应参数 `description`
    - 响应参数`service_type`类型调整 `string` -> `enum`
  - 接口`ListServiceConnections`移除响应参数 `id`、`marker_id`、`created_at`、`updated_at`、`domain_id`、`status`
  - 接口`AcceptOrRejectEndpoint`新增响应参数 `description`
  - 接口`ListServicePermissionsDetails`移除响应参数 `id`、`permission`、`created_at`
  - 接口`CreateEndpoint`:
    - 新增请求参数 `description`
    - 新增响应参数 `specification_name`、`description`、`policy_statement`、`enable_status`
  - 接口`ListEndpoints`:
    - 新增请求参数 `public_border_group`
    - 新增响应参数 `description`、`policy_statement`、`endpoint_pool_id`、`public_border_group`
  - 接口`ListEndpointInfoDetails`新增响应参数 `description`、`policy_statement`
  - 接口`ListVersionDetails`移除响应参数 `status`、`id`、`updated`、`version`、`min_version`、`links`
  - 接口`ListSpecifiedVersionDetails`移除响应参数 `status`、`id`、`updated`、`version`、`min_version`、`links`
  - 接口`ListResourceInstances`:
    - 新增请求参数 `sys_tags`、`without_any_tag`
    - 移除请求参数 `key`、`value`、`key`、`value`、`key`、`value`、`key`、`value`

# 0.1.4 2022-09-26

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDependency`新增响应参数 `version`、`last_modified`
  - 接口`ListDependencies`新增响应参数 `version`、`last_modified`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `disk_encrypted_key`、`public_management_connect_address`、`subnet_cidr`、`subnet_name`、`enable_acl`
  - 接口`ShowInstance`新增响应参数 `disk_encrypted_key`、`public_management_connect_address`、`subnet_cidr`、`subnet_name`、`enable_acl`
  - 接口`ResizeInstance`新增请求参数 `oper_type`、`new_broker_num`、`new_product_id`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 修复接口`CheckImageModeration`响应参数类型错误的问题
- _特性变更_
  - 无

# 0.1.3 2022-09-22

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateApiV2`新增请求参数 `network_type`、`alias_urn`、`network_type`、`alias_urn`
  - 接口`ShowDetailsOfApiV2`新增响应参数 `network_type`、`alias_urn`、`network_type`、`alias_urn`
  - 接口`UpdateApiV2`:
    - 新增请求参数 `network_type`、`alias_urn`、`network_type`、`alias_urn`
    - 新增响应参数 `network_type`、`alias_urn`、`network_type`、`alias_urn`
  - 接口`ListApiVersionDetailV2`新增响应参数 `network_type`、`alias_urn`、`network_type`、`alias_urn`
  - 接口`UpdateDomainV2`移除响应参数 `url_domain`、`id`、`status`、`min_ssl_version`
  - 接口`ListApisUnbindedToAclPolicyV2`新增响应参数 `req_uri`、`auth_type`
  - 接口`ListCustomAuthorizersV2`新增响应参数 `authorizer_version`、`authorizer_alias_uri`
  - 接口`CreateCustomAuthorizerV2`新增请求参数 `authorizer_version`、`authorizer_alias_uri`
  - 接口`ShowDetailsOfCustomAuthorizersV2`新增响应参数 `authorizer_version`、`authorizer_alias_uri`
  - 接口`UpdateCustomAuthorizerV2`:
    - 新增请求参数 `authorizer_version`、`authorizer_alias_uri`
    - 新增响应参数 `authorizer_version`、`authorizer_alias_uri`
  - 接口`ExportApiDefinitionsV2`请求参数`env_id`改为必填
  - 接口`ListTagsV2`:
    - 新增响应参数 `tags`
    - 移除响应参数 `responses`
  - 接口`CreateFeatureV2`移除响应参数 `id`、`name`、`enable`、`config`、`instance_id`、`update_time`
  - 接口`ShowDetailsOfInstanceV2`新增响应参数 `ingress_ip_v6`
  - 接口`UpdateInstanceV2`新增响应参数 `ingress_ip_v6`

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEvents`新增请求参数 `event_type`
  - 接口`ListEventDetail`新增响应参数 `event_type`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowAgentStatus`、`RegisterAgent`

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CheckRecord`新增响应参数 `check_end_time`
  - 接口`ShowTaskDefects`新增响应参数 `events`
  - 接口`CheckParameters`移除响应参数 `name`、`cfg_key`、`default_value`、`option_value`、`is_required`、`description`、`type`、`status`
  - 接口`CheckRulesetParameters`新增响应参数 `value`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListProjectSets`:
    - 新增响应参数 `source`
    - 移除响应参数 `status`
  - 接口`ShowTaskSet`新增响应参数 `parallel`

### HuaweiCloud SDK CTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTrackers`新增响应参数 `group_id`、`stream_id`

### HuaweiCloud SDK DDM

- _新增特性_
  - 支持接口`ResetAdministrator`、`ValidateWeakPassword`、`ResizeFlavor`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数 `admin_user_name`、`admin_user_password`
  - 接口`ShowInstance`新增响应参数 `admin_user_name`
  - 接口`ListSlowLog`新增响应参数 `host`

### HuaweiCloud SDK DevStar

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowApplicationV3`响应参数`deploy_type`新增枚举值`none`
  - 接口`ShowApplicationDependentResources`新增响应参数 `subscribe_guide`
  - 接口`ListApplicationsV6`响应参数`deploy_type`新增枚举值`none`
  - 接口`ShowApplicationReleaseRepositories`新增响应参数 `category_name`
  - 接口`ShowTemplateV3`新增响应参数 `subscribe_guide`
  - 接口`ListTemplatesV2`新增响应参数 `subscribe_guide`
  - 接口`ListTemplates`新增响应参数 `subscribe_guide`

### HuaweiCloud SDK EG

- _新增特性_
  - 支持以下接口：
    - `ListApiVersions`
    - `ListEventSchema`
    - `CreateEventSchema`
    - `ShowDetailOfEventSchema`
    - `UpdateEventSchema`
    - `DeleteEventSchema`
    - `ListEventSchemaVersions`
    - `CreateEventSchemaVersion`
    - `ShowDetailOfEventSchemaVersion`
    - `DeleteEventSchemaVersion`
    - `DiscoverEventSchemaFromData`
    - `ListConnections`
    - `CreateConnection`
    - `ShowDetailOfConnection`
    - `UpdateConnection`
    - `DeleteConnection`
    - `ListAgencies`
    - `CreateAgencies`
    - `ListTriggers`
    - `UpdateEndpoint`
    - `DeleteEndpoint`
    - `ListEndpoints`
    - `CreateEndpoint`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEventSources` 移除响应参数 `id`、`name`、`label`、`description`、`provider_type`、`event_types`、`created_time`、`updated_time`、`channel_id`、`channel_name`
  - 接口`CreateEventSource`:
    - 新增请求参数 `type`、`detail`
    - 新增响应参数 `event_types`、`type`、`detail`、`status`
  - 接口`ShowDetailOfEventSource` 新增响应参数 `event_types`、`type`、`detail`、`status`
  - 接口`UpdateEventSource`:
    - 新增请求参数 `detail`
    - 新增响应参数 `event_types`、`type`、`detail`、`status`
  - 接口`ListEventTarget`移除响应参数 `name`、`label`、`metadata`
  - 接口`ListSubscriptions`新增响应参数 `connection_id`
  - 接口`CreateSubscription`:
    - 新增请求参数 `connection_id`
    - 新增响应参数 `connection_id`
  - 接口`ShowDetailOfSubscription` 新增响应参数 `connection_id`
  - 接口`UpdateSubscription`:
    - 新增请求参数 `connection_id`
    - 新增响应参数 `connection_id`
  - 接口`CreateSubscriptionTarget`:
    - 新增请求参数 `connection_id`
    - 新增响应参数 `connection_id`
  - 接口`ShowDetailOfSubscriptionTarget` 新增响应参数 `connection_id`
  - 接口`UpdateSubscriptionTarget`:
    - 新增请求参数 `connection_id`
    - 新增响应参数 `connection_id`
  - 接口`ListQuotas`:
    - 请求参数`type`新增枚举值`CONNECTION`、`PRIVATE_ENDPOINT`、`SOURCE_RABBITMQ`
    - 响应参数`type`新增枚举值`CONNECTION`、`PRIVATE_ENDPOINT`、`SOURCE_RABBITMQ`
    - 响应参数`max`类型调整 `string` -> `int32`
    - 响应参数`min`类型调整 `string` -> `int32`
    - 响应参数`quota`类型调整 `string` -> `int32`
    - 响应参数`used`类型调整 `string` -> `int32`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePublicip`新增请求参数 `port_id`
  - 接口`CreatePrePaidPublicip`新增请求参数 `port_id`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ListGaussMySqlInstanceDetailInfo`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDeviceMessages`新增响应参数 `error_info`
  - 接口`ShowDeviceMessage`新增响应参数 `error_info`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 支持接口`CheckImageModeration`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunCreateAudioModerationJob`请求参数`url`、`categories`改为必填

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeIdCard`:
    - 新增请求参数 `detect_copy`
    - 新增响应参数 `detect_copy_result`

### HuaweiCloud SDK OMS

- _新增特性_
  - 支持以下接口：
    - `ListTaskGroup`
    - `CreateTaskGroup`
    - `ShowTaskGroup`
    - `DeleteTaskGroup`
    - `StopTaskGroup`
    - `StartTaskGroup`
    - `RetryTaskGroup`
    - `UpdateTaskGroup`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTasks`新增响应参数 `enable_metadata_migration`、`object_overwrite_mode`、`consistency_check`、`enable_requester_pays`
  - 接口`CreateTask`:
    - 新增请求参数 `enable_metadata_migration`、`object_overwrite_mode`、`consistency_check`、`enable_requester_pays`
    - 新增响应参数 `id`、`task_name`
  - 接口`ShowTask`:
    - 新增响应参数 `enable_metadata_migration`、`object_overwrite_mode`、`consistency_check`、`enable_requester_pays`
    - 请求参数`task_id`类型调整 `int64` -> `string`
  - 接口`DeleteTask`请求参数`task_id`类型调整 `int64` -> `string`
  - 接口`StopTask`请求参数`task_id`类型调整 `int64` -> `string`
  - 接口`StartTask`请求参数`task_id`类型调整 `int64` -> `string`
  - 接口`UpdateBandwidthPolicy`请求参数`task_id`类型调整 `int64` -> `string`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`DownloadImageFile`、`ListScrumProjectStatuses`、`UploadAttachments`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchRestartOrDeleteInstances`请求参数`all_failure`新增枚举值`rabbitmq`, 移除枚举值`true`、`false`

### HuaweiCloud SDK SMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTemplates`移除响应参数 `disks`
  - 接口`ShowTemplate`移除响应参数 `disks`
  - 接口`UpdateMigproject`移除请求参数 `disks`
  - 接口`ShowMigproject`移除响应参数 `disks`

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`PublishAssets`新增响应参数 `pack_type`、`pack_type`
  - 接口`UnpublishAssets`新增响应参数 `pack_type`、`pack_type`
  - 接口`ShowAssetMeta`新增响应参数 `pack_type`、`pack_type`
  - 接口`ShowAssetDetail`新增响应参数 `pack_type`、`pack_type`
  - 接口`ShowTakeOverTaskDetails`新增响应参数 `pack_type`、`pack_type`
  - 接口`ShowTakeOverAssetDetails`新增响应参数 `pack_type`、`pack_type`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DeleteIgnoreRule`新增响应参数 `rule`

# 0.1.2 2022-09-15

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerselfResourceRecordDetails`新增响应参数 `root_resource_id`、`parent_resource_id`、`trade_id`、`product_spec_desc`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ShowTags`、`CreateTags`、`BatchDeleteTags`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowUrlTaskInfo`:
    - 新增响应参数 `result`
    - 移除响应参数 `results`
  - 接口`ShowDomainFullConfig`新增响应参数 `error_code_cache`
  - 接口`UpdateDomainFullConfig`新增请求参数 `error_code_cache`

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 支持以下接口：
    - `ListFilesByQuery`
    - `ListBranchesByRepositoryId`
    - `CreateNewBranch`
    - `AssociateIssues`
    - `ListMergeRequest`
    - `ShowMergeRequest`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSubfiles`:
    - 新增响应参数 `error`、`result`、`status`
    - 移除响应参数 `trees`、`total`
  - 接口`ShowStatisticalData`:
    - 新增响应参数 `error`、`result`、`status`
    - 移除响应参数 `repoName`、`commitCount`、`repoSize`、`lastCommitTime`、`codeLines`、`branchCount`、`archiveUrl`
  - 接口`CreateCommit`请求参数`force`类型调整 `string` -> `boolean`
  - 接口`AddProtectBranchV2`:
    - 请求参数`push_access_level`类型调整 `int32` -> `enum`
    - 请求参数`merge_access_level`类型调整 `int32` -> `enum`

### HuaweiCloud SDK CSE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UploadKie`新增响应参数 `create_revision`、`update_revision`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateCloseKibana`新增请求参数 `CloseKibanaPublicReq`
  - 接口`CreateCluster`新增请求参数 `payInfo`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowInstance`新增响应参数 `enterprise_project_name`、`update_at`、`product_type`、`storage_type`、`launched_at`、`cache_mode`、`support_slow_log_flag`、`db_number`、`replica_count`、`sharding_count`、`bandwidth_info`
  - 接口`ListRedislog`新增响应参数 `backupId`
  - 接口`ShowIpWhitelist`新增响应参数 `instance_id`
  - 接口`UpdateIpWhitelist`新增请求参数 `instance_id`
  - 接口`ListBackgroundTask`新增响应参数 `updated_at`、`created_at`、`status`

### HuaweiCloud SDK EVS

- _新增特性_
  - 支持以下接口：
    - `ShowVersion`
    - `ListVersions`
    - `CinderShowVolumeTransfer`
    - `CinderDeleteVolumeTransfer`
    - `CinderListVolumeTransfers`
    - `CinderCreateVolumeTransfer`
    - `CinderAcceptVolumeTransfer`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IEF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListApps`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`UpdateApp`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`ShowAppDetail`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`ListAppVersions`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`CreateAppVersions`新增请求参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`UpdateAppVersion`:
    - 新增请求参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
    - 新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`ShowAppVersionDetail`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateLogStream`新增请求参数 `ttl_in_days`
  - 接口`ListStructuredLogsWithTimeRange`新增请求参数 `whether_to_rows`
  - 接口`UpdateStructTemplate`请求参数`isAnalysis`改为非必填
  - 接口`CreateStructTemplate`请求参数`isAnalysis`改为非必填

### HuaweiCloud SDK Moderation

- _新增特性_
  - 支持接口`RunCreateVideoModerationJob`、`RunQueryVideoModerationJob`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunCreateAudioModerationJob`请求参数`url`改为非必填

# 0.1.1 2022-09-08

### HuaweiCloud SDK IAM

- _新增特性_
  - 支持自定义认证凭据 `IamCredentials`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerselfResourceRecordDetails`新增响应参数 `root_resource_id`、`parent_resource_id`、`trade_id`、`product_spec_desc`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ShowTags`、`CreateTags`、`BatchDeleteTags`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowUrlTaskInfo`:
    - 新增响应参数 `result`
    - 移除响应参数 `results`
  - 接口`ShowDomainFullConfig`新增响应参数 `error_code_cache`
  - 接口`UpdateDomainFullConfig`新增请求参数 `error_code_cache`

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 支持以下接口：
    - `ListFilesByQuery`
    - `ListBranchesByRepositoryId`
    - `CreateNewBranch`
    - `AssociateIssues`
    - `ListMergeRequest`
    - `ShowMergeRequest`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSubfiles`:
    - 新增响应参数 `error`、`result`、`status`
    - 移除响应参数 `trees`、`total`
  - 接口`ShowStatisticalData`:
    - 新增响应参数 `error`、`result`、`status`
    - 移除响应参数 `repoName`、`commitCount`、`repoSize`、`lastCommitTime`、`codeLines`、`branchCount`、`archiveUrl`
  - 接口`CreateCommit`请求参数`force`类型调整 `string` -> `boolean`
  - 接口`AddProtectBranchV2`:
    - 请求参数`push_access_level`类型调整 `int32` -> `enum`
    - 请求参数`merge_access_level`类型调整 `int32` -> `enum`

### HuaweiCloud SDK CSE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UploadKie`新增响应参数 `create_revision`、`update_revision`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateCloseKibana`新增请求参数 `CloseKibanaPublicReq`
  - 接口`CreateCluster`新增请求参数 `payInfo`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowInstance`新增响应参数 `enterprise_project_name`、`update_at`、`product_type`、`storage_type`、`launched_at`、`cache_mode`、`support_slow_log_flag`、`db_number`、`replica_count`、`sharding_count`、`bandwidth_info`
  - 接口`ListRedislog`新增响应参数 `backupId`
  - 接口`ShowIpWhitelist`新增响应参数 `instance_id`
  - 接口`UpdateIpWhitelist`新增请求参数 `instance_id`
  - 接口`ListBackgroundTask`新增响应参数 `updated_at`、`created_at`、`status`

### HuaweiCloud SDK EVS

- _新增特性_
  - 支持以下接口：
    - `ShowVersion`
    - `ListVersions`
    - `CinderShowVolumeTransfer`
    - `CinderDeleteVolumeTransfer`
    - `CinderListVolumeTransfers`
    - `CinderCreateVolumeTransfer`
    - `CinderAcceptVolumeTransfer`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IEF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListApps`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`UpdateApp`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`ShowAppDetail`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`ListAppVersions`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`CreateAppVersions`新增请求参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`UpdateAppVersion`:
    - 新增请求参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
    - 新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`
  - 接口`ShowAppVersionDetail`新增响应参数 `cpu`、`memory`、`gpu`、`npu`、`cpu`、`memory`、`gpu`、`npu`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateLogStream`新增请求参数 `ttl_in_days`
  - 接口`ListStructuredLogsWithTimeRange`新增请求参数 `whether_to_rows`
  - 接口`UpdateStructTemplate`请求参数`isAnalysis`改为非必填
  - 接口`CreateStructTemplate`请求参数`isAnalysis`改为非必填

### HuaweiCloud SDK Moderation

- _新增特性_
  - 支持接口`RunCreateVideoModerationJob`、`RunQueryVideoModerationJob`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunCreateAudioModerationJob`请求参数`url`改为非必填

# 0.0.108 2022-09-01

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - [Issue 56](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/56) 修复RequestHandler的问题
- _特性变更_
  - 无

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerOrders`新增请求参数 `indirect_partner_id`
  - 接口`ShowCustomerOrderDetails`新增请求参数 `indirect_partner_id`
  - 接口`ListCustomerOnDemandResources`新增请求参数 `indirect_partner_id`

### HuaweiCloud SDK CC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCloudConnections`响应参数`used_scene`移除枚举值`er`
  - 接口`CreateCloudConnection`移除请求参数 `used_scene`
  - 接口`ShowCloudConnection`响应参数`used_scene`移除枚举值`er`
  - 接口`UpdateCloudConnection`响应参数`used_scene`移除枚举值`er`
  - 接口`ListNetworkInstances`响应参数`type`移除枚举值`er`
  - 接口`CreateNetworkInstance`请求参数`type`移除枚举值`er`
  - 接口`ShowNetworkInstance`响应参数`type`移除枚举值`er`
  - 接口`UpdateNetworkInstance`响应参数`type`移除枚举值`er`

### HuaweiCloud SDK CDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJobs`响应参数`id`、`type`改为非必填
  - 接口`UpdateJob`请求参数`id`、`type`改为非必填
  - 接口`CreateAndStartRandomClusterJob`请求参数`id`、`type`改为非必填
  - 接口`CreateJob`请求参数`id`、`type`改为非必填

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`新增响应参数 `ordeId`
  - 接口`ShowClusterDetail`:
    - 新增响应参数 `vpcepIp`、`elbWhiteListResp`
    - 移除响应参数 `elbWhiteList`、`action`
  - 接口`UpdateUnbindPublic`移除请求参数 `isAutoPay`
  - 接口`ListYmlsJob`:
    - 新增响应参数 `configList`
    - 移除响应参数 `configurations`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDependency`移除响应参数 `version`、`last_modified`
  - 接口`ListDependencies`移除响应参数 `version`、`last_modified`

### HuaweiCloud SDK IAM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateDomainProtectPolicy`:
    - 新增请求参数 `allow_user`、`mobile`、`admin_check`、`email`、`scene`
    - 移除响应参数 `operation_protection`
  - 接口`ShowDomainProtectPolicy`移除响应参数 `operation_protection`
  - 接口`UpdateDomainPasswordPolicy`请求参数`maximum_consecutive_identical_chars`、`minimum_password_age`、`minimum_password_length`、`number_of_recent_passwords_disallowed`、`password_not_username_or_invert`、`password_validity_period`、`password_char_combination`改为非必填
  - 接口`UpdateDomainLoginPolicy`请求参数`account_validity_period`、`custom_info_for_login`、`lockout_duration`、`login_failed_times`、`period_with_login_failures`、`session_timeout`、`show_recent_login_info`改为非必填
  - 接口`ShowDomainQuota`请求参数`type`新增枚举值`mapping`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`UpdateDbUserComment`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDbUser`新增请求参数 `comment`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListStatistics`:
    - 请求参数`hosts`类型调整 `array` -> `string`
    - 请求参数`instances`类型调整 `array` -> `string`
  - 接口`ListQpsTimeline`:
    - 请求参数`hosts`类型调整 `array` -> `string`
    - 请求参数`instances`类型调整 `array` -> `string`
  - 接口`ListBandwidthTimeline`请求参数`instances`类型调整 `array` -> `string`
  - 接口`ListTopAbnormal`:
    - 请求参数`hosts`类型调整 `array` -> `string`
    - 请求参数`instances`类型调整 `array` -> `string`
  - 接口`ListOverviewsClassification`:
    - 请求参数`hosts`类型调整 `array` -> `string`
    - 请求参数`instances`类型调整 `array` -> `string`

# 0.0.107 2022-08-29

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePostPaidServers`新增请求参数 `batch_create_in_multi_az`

### HuaweiCloud SDK IMS

- _新增特性_
  - 支持接口`ShowJobProgress`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`SetReadOnlySwitch`
- _解决问题_
  - 无
- _特性变更_
  - 接口`StartFailover`新增请求参数 `FailoverRequestBody`

# 0.0.106 2022-08-25

### HuaweiCloud SDK CDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateJob`:
    - 新增请求参数 `rows_written`、`rows_read`、`files_written`、`extended-configs`、`value`、`extended-configs`、`value`、`extended-configs`、`value`
    - 移除请求参数 `files_writte`、`values`、`values`、`values`
  - 接口`ShowJobs`:
    - 新增响应参数 `rows_written`、`rows_read`、`files_written`、`extended-configs`、`value`、`extended-configs`、`value`、`extended-configs`、`value`
    - 移除响应参数 `files_writte`、`values`、`values`、`values`
  - 接口`CreateAndStartRandomClusterJob`:
    - 新增请求参数 `rows_written`、`rows_read`、`files_written`、`extended-configs`、`value`、`extended-configs`、`value`、`extended-configs`、`value`
    - 移除请求参数 `files_writte`、`values`、`values`、`values`
  - 接口`CreateJob`:
    - 新增请求参数 `rows_written`、`rows_read`、`files_written`、`extended-configs`、`value`、`extended-configs`、`value`、`extended-configs`、`value`
    - 移除请求参数 `files_writte`、`values`、`values`、`values`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `ListMasterSlavePools`
    - `CreateMasterSlavePool`
    - `ShowMasterSlavePool`
    - `DeleteMasterSlavePool`
  - 接口`ListLoadBalancers`新增响应参数 `waf_failure_action`
  - 接口`CreateLoadBalancer`新增请求参数 `waf_failure_action`
  - 接口`ShowLoadBalancer`新增响应参数 `waf_failure_action`
  - 接口`UpdateLoadBalancer`:
    - 新增请求参数 `waf_failure_action`
    - 新增响应参数 `waf_failure_action`
    - 移除请求参数 `cloud_service_console_url`
  - 接口`ListCertificates`新增响应参数 `enc_certificate`、`enc_private_key`
  - 接口`CreateCertificate`新增请求参数 `enc_certificate`、`enc_private_key`
  - 接口`ShowCertificate`新增响应参数 `enc_certificate`、`enc_private_key`
  - 接口`UpdateCertificate`:
    - 新增请求参数 `enc_certificate`、`enc_private_key`
    - 新增响应参数 `enc_certificate`、`enc_private_key`
  - 接口`ListListeners`新增响应参数 `sni_match_algo`
  - 接口`CreateListener`新增请求参数 `sni_match_algo`
  - 接口`ShowListener`新增响应参数 `sni_match_algo`
  - 接口`UpdateListener`:
    - 新增请求参数 `sni_match_algo`
    - 新增响应参数 `sni_match_algo`
  - 接口`ListMembers`新增请求参数 `instance_id`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeWebImage`:
    - 新增请求参数 `detect_font`
    - 新增响应参数 `font_list`、`font_scores`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`SetDatabaseUserPrivilege`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CheckMd5Duplication`请求参数`size`类型调整 `int32` -> `int64`

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持接口`ListRequestTimeline`
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateGeoipRule`新增响应参数 `description`

# 0.0.105 2022-08-22

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRecordIndex`请求参数`object`类型调整 `uri` -> `string`

# 0.0.104 2022-08-18

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 支持接口`ListIndirectPartners`、`ListCosts`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSubCustomers`新增请求参数 `indirect_partner_id`
  - 接口`CreateSubCustomer`新增请求参数 `indirect_partner_id`
  - 接口`ShowSubCustomerBudget`新增请求参数 `indirect_partner_id`
  - 接口`UpdateSubCustomerBudget`新增请求参数 `indirect_partner_id`
  - 接口`FreezeSubCustomers`新增请求参数 `indirect_partner_id`
  - 接口`UnfreezeSubCustomers`新增请求参数 `indirect_partner_id`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListProjectTemplates`新增响应参数 `arch`

### HuaweiCloud SDK CPH

- _新增特性_
  - 支持云手机服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ECS

- _新增特性_
  - 支持接口`ListServersByTag`
- _解决问题_
  - 无
- _特性变更_
  - 接口`NovaCreateServers`请求参数`destination_type`改为必填

### HuaweiCloud SDK EG

- _新增特性_
  - 支持事件网格服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateFunction`:
    - 新增响应参数 `enable_dynamic_memory`、`is_stateful_function`、`enable_auth_in_header`、`custom_image`
    - 请求参数`file`、`link`改为必填
    - 响应参数`user_id`类型调整 `int32` -> `string`
    - 响应参数`user_group_id`类型调整 `int32` -> `string`
    - 响应参数`concurrent_num`改为必填
    - 响应参数`mount_share_path`改为非必填
  - 接口`ListFunctions`:
    - 新增响应参数 `fail_count`
    - 移除请求参数 `X-Auth-Token`
    - 响应参数`concurrent_num`改为必填
  - 接口`ShowFunctionCode`:
    - 移除响应参数 `id`
    - 响应参数`file`、`link`、`concurrent_num`改为必填
  - 接口`UpdateFunctionCode`:
    - 移除响应参数 `id`
    - 请求参数`file`、`link`改为必填
    - 响应参数`file`、`link`、`concurrent_num`改为必填
  - 接口`ShowFunctionConfig`:
    - 新增响应参数 `is_stateful_function`、`enable_auth_in_header`、`custom_image`
    - 移除响应参数 `id`
    - 响应参数`user_id`类型调整 `int32` -> `string`
    - 响应参数`user_group_id`类型调整 `int32` -> `string`
    - 响应参数`concurrent_num`改为必填
    - 响应参数`mount_share_path`改为非必填
  - 接口`UpdateFunctionConfig`:
    - 新增响应参数 `enable_auth_in_header`、`custom_image`
    - 移除请求参数 `X-Auth-Token`
    - 移除响应参数 `id`
    - 请求参数`user_id`类型调整 `int32` -> `string`
    - 请求参数`user_group_id`类型调整 `int32` -> `string`
    - 请求参数`concurrent_num`改为必填
    - 请求参数`mount_share_path`改为非必填
    - 响应参数`user_id`类型调整 `int32` -> `string`
    - 响应参数`user_group_id`类型调整 `int32` -> `string`
    - 响应参数`concurrent_num`改为必填
    - 响应参数`mount_share_path`改为非必填
  - 接口`UpdateFunctionMaxInstanceConfig`:
    - 移除响应参数 `id`
    - 响应参数`user_id`类型调整 `int32` -> `string`
    - 响应参数`user_group_id`类型调整 `int32` -> `string`
    - 响应参数`concurrent_num`改为必填
    - 响应参数`mount_share_path`改为非必填
  - 接口`CreateFunctionVersion`:
    - 移除响应参数 `id`
    - 响应参数`user_id`类型调整 `int32` -> `string`
    - 响应参数`user_group_id`类型调整 `int32` -> `string`
    - 响应参数`concurrent_num`改为必填
    - 响应参数`mount_share_path`改为非必填
  - 接口`ListFunctionVersions`:
    - 新增响应参数 `is_stateful_function`、`enable_auth_in_header`、`custom_image`、`reserved_instance_idle_mode`
    - 移除响应参数 `log_group_id`、`log_stream_id`
    - 响应参数`concurrent_num`改为必填
  - 接口`CreateFunctionTrigger`请求参数`trigger_type_code`新增枚举值`SMN`、`RABBITMQ`、`DEDICATEDGATEWAY`、`OPENSOURCEKAFKA`、`APIC`、`GAUSSMONGO`、`EVENTGRID`
  - 接口`ListFunctionTriggers`:
    - 响应参数`trigger_type_code`新增枚举值`RABBITMQ`、`DEDICATEDGATEWAY`、`OPENSOURCEKAFKA`、`APIC`、`GAUSSMONGO`、`EVENTGRID`
    - 响应参数`trigger_status`新增枚举值`DISABLE`, 移除枚举值`DISABLED`
  - 接口`DeleteFunctionTrigger`请求参数`trigger_type_code`新增枚举值`RABBITMQ`、`DEDICATEDGATEWAY`、`OPENSOURCEKAFKA`、`APIC`、`GAUSSMONGO`、`EVENTGRID`
  - 接口`ShowFunctionTrigger`:
    - 请求参数`trigger_type_code`新增枚举值`RABBITMQ`、`DEDICATEDGATEWAY`、`OPENSOURCEKAFKA`、`APIC`、`GAUSSMONGO`、`EVENTGRID`
    - 响应参数`trigger_type_code`新增枚举值`RABBITMQ`、`DEDICATEDGATEWAY`、`OPENSOURCEKAFKA`、`APIC`、`GAUSSMONGO`、`EVENTGRID`
    - 响应参数`trigger_status`新增枚举值`DISABLE`, 移除枚举值`DISABLED`
  - 接口`UpdateTrigger`:
    - 请求参数`trigger_status`新增枚举值`DISABLE`, 移除枚举值`DISABLED`
    - 请求参数`trigger_type_code`新增枚举值`RABBITMQ`、`DEDICATEDGATEWAY`、`OPENSOURCEKAFKA`、`APIC`、`GAUSSMONGO`、`EVENTGRID`
    - 请求参数`trigger_status`改为非必填
  - 接口`ListStatistics`响应参数`value`类型调整 `float` -> `int32`
  - 接口`UpdateFunctionReservedInstancesCount`:
    - 新增请求参数 `UpdateFunctionReservedInstancesCountRequestBody`
    - 新增响应参数 `idle_mode`、`tactics_config`
    - 移除请求参数 `UpdateFunctionReservedInstancesRequestBody`
  - 接口`CreateDependency`:
    - 新增响应参数 `version`、`last_modified`
    - 响应参数`runtime`类型调整 `enum` -> `string`
  - 接口`ListDependencies`:
    - 新增请求参数 `maxitems`、`ispublic`
    - 新增响应参数 `version`、`last_modified`
    - 响应参数`count`类型调整 `int32` -> `int64`
  - 接口`ShowDependcy`:
    - 新增响应参数 `version`、`last_modified`
    - 响应参数`runtime`类型调整 `enum` -> `string`
  - 接口`UpdateDependcy`:
    - 新增请求参数 `UpdateDependcyRequestBody`
    - 移除请求参数 `UpdateDependencyRequestBody`
    - 响应参数`runtime`类型调整 `enum` -> `string`
  - 接口`CreateEvent`移除响应参数 `content`、`last_modified`
  - 接口`UpdateEvent`移除响应参数 `content`、`last_modified`
  - 接口`ImportFunction`:
    - 新增请求参数 `package`
    - 移除请求参数 `X-Auth-Token`
    - 响应参数`concurrent_num`改为必填
  - 接口`EnableLtsLogs`新增请求参数 `X-Auth-Token`
  - 接口`ShowLtsLogDetails`新增响应参数 `group_name`
  - 接口`CancelAsyncInvocation`请求参数`request_id`改为必填
  - 接口`CreateWorkflow`:
    - 新增请求参数 `duration`
    - 请求参数`trigger_type`新增枚举值`SMN`、`APIG`、`APIG_DE`
  - 接口`ListWorkflow`:
    - 新增请求参数 `enterprise_project`、`mode`
    - 移除响应参数 `id`、`workflow_urn`、`name`、`description`、`created_time`、`updated_time`、`created_by`
  - 接口`StartWorkflowExecution`新增请求参数 `X-WorkflowRun-MergeFnParameters`
  - 接口`ListWorkflowExecutions`移除响应参数 `workflow_id`、`workflow_urn`、`execution_id`、`status`、`begin_time`、`end_time`、`last_update_time`、`created_by`
  - 接口`ShowWorkflowExecution`:
    - 新增请求参数 `X-Get-Workflow-Full-History-Data`
    - 响应参数`workflow_urn`类型调整 `string` -> `object`
  - 接口`ShowWorkFlow`:
    - 移除响应参数 `name`、`description`、`triggers`、`start`、`functions`、`states`、`constants`、`retries`、`mode`、`express_config`、`enterprise_project_id`
    - 响应参数`workflow_urn`类型调整 `string` -> `object`
    - 响应参数`id`、`workflow_urn`、`created_time`、`updated_time`、`created_by`改为必填
  - 接口`UpdateWorkFlow`:
    - 新增请求参数 `duration`
    - 请求参数`trigger_type`新增枚举值`SMN`、`APIG`、`APIG_DE`
    - 响应参数`workflow_urn`类型调整 `string` -> `object`
    - 响应参数`id`、`workflow_urn`、`name`、`description`、`created_time`、`updated_time`、`created_by`改为必填

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeIdCard`:
    - 新增请求参数 `detect_reproduce`
    - 新增响应参数 `detect_reproduce_result`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateIssueV4`:
    - 新增请求参数 `field_name`
    - 新增响应参数 `field_name`
  - 接口`ListIssuesV4`新增响应参数 `field_name`
  - 接口`UpdateIssueV4`:
    - 新增请求参数 `field_name`
    - 新增响应参数 `field_name`
  - 接口`ListChildIssuesV4`新增响应参数 `field_name`
  - 接口`CreateSystemIssueV4`:
    - 新增请求参数 `field_name`
    - 新增响应参数 `field_name`

### HuaweiCloud SDK ROMA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListNotification`新增请求参数 `limit`、`offset`

### HuaweiCloud SDK VOD

- _新增特性_
  - 支持以下接口：
    - `ListTranscodeTemplate`
    - `UpdateTranscodeTemplate`
    - `CreateTranscodeTemplate`
    - `DeleteTranscodeTemplate`
    - `ListTemplateGroupCollection`
    - `UpdateTemplateGroupCollection`
    - `CreateTemplateGroupCollection`
    - `DeleteTemplateGroupCollection`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIgnoreRule`:
    - 新增响应参数 `domain`
    - 移除响应参数 `domains`
  - 接口`ListGeoipRule`新增响应参数 `policyid`
  - 接口`UpdateGeoipRule`新增请求参数 `description`

# 0.0.103 2022-08-11

### HuaweiCloud SDK APM

- _新增特性_
  - 支持以下接口：
    - `ListOpenRegion`
    - `ListSupportedRegion`
    - `ShowTopologyTree`
    - `ShowMonitorItemViewConfig`
    - `ListEnvTags`
    - `ShowTopology`
    - `ShowEventDetail`
    - `ShowSpanSearch`
    - `ShowTraceEvents`
    - `ShowTrend`
    - `ShowSumTable`
    - `ShowRawTable`
    - `ShowClobDetail`
    - `ListEnvInstances`
    - `ShowEnvMonitorItems`
    - `ListApps`
    - `ListAppEnvs`
    - `ShowAkSks`
    - `CreateAkSk`
    - `DeleteAkSk`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListCosts`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ShowUrlTaskInfo`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowDomainFullConfig`新增响应参数 `ipv6_accelerate`
  - 接口`UpdateDomainFullConfig`新增请求参数 `ipv6_accelerate`

### HuaweiCloud SDK CSMS

- _新增特性_
  - 支持接口`UploadSecretBlob`、`DownloadSecretBlob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持以下接口：
    - `ListGaussMySqlDatabaseUser`
    - `CreateGaussMySqlDatabaseUser`
    - `DeleteGaussMySqlDatabaseUser`
    - `ResetGaussMySqlDatabasePassword`
    - `AddDatabasePermission`
    - `DeleteDatabasePermission`
    - `ListGaussMySqlDatabaseCharsets`
    - `ListGaussMySqlDatabase`
    - `CreateGaussMySqlDatabase`
    - `DeleteGaussMySqlDatabase`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListGraphs`新增响应参数 `enableHyG`、`trafficIpList`、`cryptAlgorithm`、`enableHttps`、`tags`
  - 接口`ShowGraph`新增响应参数 `enableHyG`、`trafficIpList`、`cryptAlgorithm`、`enableHttps`、`tags`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePostPaidInstance`:
    - 新增请求参数 `broker_num`
    - 请求参数`specification`新增枚举值`c6.2u4g.cluster`、`c6.4u8g.cluster`、`c6.8u16g.cluster`、`c6.12u24g.cluster`、`c6.16u32g.cluster`
    - 请求参数`partition_num`新增枚举值`250`、`500`、`1000`、`1500`、`2000`
    - 请求参数`storage_spec_code`新增枚举值`dms.physical.storage.high.v2`、`dms.physical.storage.ultra.v2`
    - 请求参数`specification`改为非必填
  - 接口`ListInstances`新增响应参数 `description`、`access_user`、`ssl_two_way_enable`、`cert_replaced`、`public_boundwidth`、`agent_enable`、`public_access_enabled`、`node_num`、`new_spec_billing_enable`、`broker_num`
  - 接口`ShowInstance`新增响应参数 `description`、`access_user`、`ssl_two_way_enable`、`cert_replaced`、`public_boundwidth`、`agent_enable`、`public_access_enabled`、`node_num`、`new_spec_billing_enable`、`broker_num`
  - 接口`ShowInstanceExtendProductInfo`请求参数`engine`改为非必填
  - 接口`ShowPartitionBeginningMessage`:
    - 新增响应参数 `offset`
    - 移除响应参数 `message_offset`
  - 接口`ShowPartitionEndMessage`:
    - 新增响应参数 `offset`
    - 移除响应参数 `message_offset`
  - 接口`ListEngineProducts`新增响应参数 `product_alias`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 支持接口`RunCreateAudioModerationJob`、`RunQueryAudioModerationJob`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunImageModeration`移除请求参数 `ad_glossaries`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeMacaoIdCard`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRestoreInstance`移除请求参数 `count`

### HuaweiCloud SDK SWR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListNamespaces`新增请求参数 `filter`
  - 接口`ListReposDetails`新增请求参数 `limit`、`offset`、`order_column`、`order_type`
  - 接口`ListRepositoryTags`新增请求参数 `filter`
  - 接口`ListSharedReposDetails`新增请求参数 `namespace`、`name`、`center`、`limit`、`offset`、`order_column`、`order_type`
  - 接口`ListRetentionHistories`:
    - 新增请求参数 `filter`
    - 移除请求参数 `offset`、`limit`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DeleteIgnoreRule`:
    - 新增响应参数 `advanced`
    - 移除响应参数 `rule`
  - 接口`CreateIgnoreRule`:
    - 新增请求参数 `advanced`
    - 新增响应参数 `advanced`
  - 接口`ListIgnoreRule`新增响应参数 `advanced`
  - 接口`ListStatistics`:
    - 移除响应参数 `host`
    - 请求参数`instances`类型调整 `string` -> `array`
    - 请求参数`hosts`类型调整 `string` -> `array`
  - 接口`ListQpsTimeline`:
    - 请求参数`instances`类型调整 `string` -> `array`
    - 请求参数`hosts`类型调整 `string` -> `array`
  - 接口`ListBandwidthTimeline`请求参数`instances`类型调整 `string` -> `array`
  - 接口`ListTopAbnormal`:
    - 请求参数`instances`类型调整 `string` -> `array`
    - 请求参数`hosts`类型调整 `string` -> `array`
  - 接口`ListOverviewsClassification`:
    - 请求参数`instances`类型调整 `string` -> `array`
    - 请求参数`hosts`类型调整 `string` -> `array`
  - 接口`ShowConsoleConfig`新增响应参数 `geoip_enable`、`load_balance_enable`、`ipv6_protection_enable`、`policy_sharing_enable`、`ip_group`、`robot_action_enable`、`http2_enable`、`timeout_config_enable`
  - 接口`CreateValueList`新增响应参数 `producer`
  - 接口`ListValueList`响应参数`type`类型调整 `string` -> `enum`
  - 接口`UpdateValueList`移除响应参数 `timestamp`
  - 接口`ListEvent`新增响应参数 `payload_location`
  - 接口`CreateHost`:
    - 新增请求参数 `web_tag`、`exclusive_ip`、`paid_type`、`lb_algorithm`、`weight`
    - 新增响应参数 `lb_algorithm`、`web_tag`、`block_page`、`extend`、`weight`、`ipv6`
  - 接口`ListHost`:
    - 新增响应参数 `region`、`web_tag`、`ipv6`
    - 移除响应参数 `timeout_config`
  - 接口`ListHostRoute`响应参数`back_protocol`类型调整 `string` -> `enum`
  - 接口`DeleteHost`新增响应参数 `web_tag`、`ipv6`
  - 接口`UpdateHost`:
    - 新增请求参数 `web_tag`、`exclusive_ip`、`paid_type`、`circuit_breaker`
    - 新增响应参数 `projectid`、`extend`、`traffic_mark`、`circuit_breaker`、`access_progress`、`weight`、`ipv6`
    - 移除请求参数 `lb_algorithm`
    - 移除响应参数 `ipv6_enable`
    - 响应参数`protocol`类型调整 `enum` -> `string`
    - 响应参数`web_tag`类型调整 `boolean` -> `string`
    - 响应参数`lb_algorithm`类型调整 `string` -> `enum`
  - 接口`ShowHost`新增响应参数 `domainid`、`projectid`、`enterprise_project_id`、`locked`、`tls`、`cipher`、`block_page`、`extend`、`traffic_mark`、`circuit_breaker`、`lb_algorithm`、`web_tag`、`flag`、`description`、`http2_enable`、`access_progress`、`weight`
  - 接口`CreatePolicy`新增响应参数 `robot_action`、`modulex_options`
  - 接口`ListPolicy`新增响应参数 `robot_action`、`modulex_options`、`hosts`
  - 接口`DeletePolicy`新增响应参数 `robot_action`、`modulex_options`
  - 接口`UpdatePolicyProtectHost`新增响应参数 `robot_action`、`modulex_options`、`hosts`
  - 接口`UpdatePolicy`:
    - 新增请求参数 `level`、`full_detection`、`robot_action`、`modulex_options`、`hosts`、`bind_host`、`extend`
    - 新增响应参数 `robot_action`、`modulex_options`
  - 接口`ShowPolicy`新增响应参数 `robot_action`、`modulex_options`、`hosts`
  - 接口`UpdatePolicyRuleStatus`请求参数`ruletype`新增枚举值`custom`、`ignore`
  - 接口`CreateWhiteblackipRule`:
    - 新增请求参数 `ip_group_id`
    - 新增响应参数 `name`、`ip_group`、`status`、`description`
    - 请求参数`addr`改为非必填
  - 接口`ListWhiteblackipRule`新增响应参数 `name`、`ip_group`
  - 接口`DeleteWhiteBlackIpRule`新增响应参数 `ip_group`
  - 接口`UpdateWhiteblackipRule`:
    - 新增请求参数 `ip_group_id`
    - 新增响应参数 `name`、`ip_group`
    - 请求参数`addr`改为非必填
  - 接口`CreatePrivacyRule`新增响应参数 `timestamp`、`status`、`description`
  - 接口`ListPrivacyRule`新增响应参数 `description`
  - 接口`UpdatePrivacyRule`新增响应参数 `timestamp`、`status`、`description`
  - 接口`CreateGeoipRule`:
    - 新增请求参数 `name`、`status`
    - 新增响应参数 `name`、`status`
  - 接口`ListGeoipRule`新增响应参数 `name`、`status`
  - 接口`DeleteGeoipRule`新增响应参数 `name`、`status`
  - 接口`UpdateGeoipRule`:
    - 新增请求参数 `name`
    - 新增响应参数 `name`
  - 接口`ListCertificates`移除响应参数 `content`、`key`
  - 接口`ListCompositeHosts`:
    - 新增响应参数 `hostid`、`web_tag`、`access_progress`、`premium_waf_instances`、`description`、`exclusive_ip`、`region`
    - 移除响应参数 `pci_dss`、`pci_3ds`、`cname`、`is_dual_az`、`ipv6`
  - 接口`ShowCompositeHost`:
    - 新增响应参数 `hostid`、`web_tag`、`access_progress`、`premium_waf_instances`、`description`、`exclusive_ip`、`region`
    - 移除响应参数 `pci_dss`、`pci_3ds`、`cname`、`is_dual_az`、`ipv6`
  - 接口`CreatePremiumHost`:
    - 新增请求参数 `block_page`、`description`、`weight`
    - 新增响应参数 `server`、`proxy`、`locked`、`timestamp`、`tls`、`cipher`、`extend`、`flag`、`description`、`enterprise_project_id`、`protect_status`、`access_status`、`block_page`
    - 响应参数`protocol`类型调整 `string` -> `enum`
  - 接口`ListPremiumHost`:
    - 新增响应参数 `extend`、`region`、`description`、`web_tag`、`hostid`
    - 移除响应参数 `mode`、`pool_ids`
  - 接口`DeletePremiumHost`:
    - 新增响应参数 `extend`、`description`、`web_tag`、`host_id`
    - 移除响应参数 `mode`、`pool_ids`
  - 接口`UpdatePremiumHost`:
    - 新增响应参数 `description`、`projectid`、`enterprise_project_id`、`web_tag`、`lb_algorithm`、`access_progress`、`weight`
    - 移除请求参数 `flag`、`extend`
    - 移除响应参数 `mode`、`pool_ids`、`project_id`、`access_code`
  - 接口`ShowPremiumHost`:
    - 新增响应参数 `description`、`projectid`、`enterprise_project_id`、`web_tag`、`access_progress`、`weight`
    - 移除响应参数 `mode`、`pool_ids`、`project_id`、`access_code`
  - 接口`UpdateCertificate`:
    - 新增请求参数 `content`、`key`
    - 请求参数`name`改为必填

# 0.0.102 2022-08-08

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复枚举类型的查询参数拼接错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持以下接口:
    - `UpdateSqlFilterControl`
    - `ShowSqlFilterControl`
    - `SetSqlFilterRule`
    - `ShowSqlFilterRule`
    - `DeleteSqlFilterRule`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.101 2022-08-02

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowGaussMySqlFlavors`新增响应参数 `flavors`
  - 接口`ShowGaussMySqlInstanceInfo`新增响应参数 `instance`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持接口`UpgradeDbVersion`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `patch_available`

### HuaweiCloud SDK Image

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunImageDescription`移除请求参数 `language`

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`CreateRecordIndex`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`CreateProjectDomain`、`ListProjectDomains`、`UpdateProjectDomain`、`CancelProjectDomain`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeShortAudio`请求参数`audio_format`新增枚举值`auto`

# 0.0.100 2022-07-28

### HuaweiCloud SDK CBS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CreateTbSession`、`ExecuteTbSession`、`DeleteTbSession`
  - 接口`CollectHotQuestions`请求参数`top`类型调整 `string` -> `int32`

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateJobs`请求参数`db_type`、`db_type`、`key`、`value`、`key`、`value`改为非必填
  - 接口`BatchSetObjects`请求参数`id`、`object_type`、`object_name`改为非必填
  - 接口`BatchUpdateJob`请求参数`name`、`alarm_to_user`、`db_type`、`db_type`、`node_type`、`engine_type`、`net_type`、`store_db_info`、`key`、`value`改为非必填
  - 接口`BatchListJobDetails`响应参数`db_type`、`db_type`、`db_type`、`db_type`改为非必填
  - 接口`BatchChangeData`请求参数`id`、`select`改为非必填

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ShowDedicatedResourceInfo`、`SetGaussMySqlProxyWeight`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowGaussMySqlProxy`新增响应参数 `proxy`、`master_node`、`readonly_nodes`
  - 接口`ShowGaussMySqlProxyList`新增响应参数 `proxy_list`
  - 接口`ShowGaussMySqlProxyFlavors`新增响应参数 `proxy_flavor_groups`
  - 接口`ShowGaussMySqlBackupList`:
    - 响应参数`status`新增枚举值`BUILDING`、`COMPLETED`、`FAILED`、`AVAILABLE`
    - 响应参数`type`新增枚举值`auto`、`manual`
    - 响应参数`backup_level`新增枚举值`0`、`1`、`2`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`:
    - 新增请求参数 `period_type`、`period_num`、`is_auto_renew`、`is_auto_pay`
    - 请求参数`charge_mode`新增枚举值`prePaid`
  - 接口`RunInstanceAction`新增请求参数 `is_auto_pay`
  - 接口`CreateRestoreInstance`:
    - 新增请求参数 `period_type`、`period_num`、`is_auto_renew`、`is_auto_pay`
    - 请求参数`charge_mode`新增枚举值`prePaid`
  - 接口`ResizeInstanceFlavor`请求参数`is_auto_pay`类型调整 `string` -> `boolean`

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持接口`ShowMonthUsages`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK KMS

- _新增特性_
  - 支持以下接口：
    - `ListKeyStores`
    - `CreateKeyStore`
    - `ShowKeyStore`
    - `DeleteKeyStore`
    - `EnableKeyStore`
    - `DisableKeyStore`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateKey`新增请求参数 `keystore_id`
  - 接口`ListKeys`新增响应参数 `keystore_id`、`key_label`
  - 接口`ListKeyDetail`新增响应参数 `keystore_id`、`key_label`
  - 接口`ListKmsByTags`新增响应参数 `keystore_id`、`key_label`

### HuaweiCloud SDK NLP

- _新增特性_
  - 支持接口`RunConstituencyParser`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.99 2022-07-21

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListLatelyApiStatisticsV2`响应参数`status`类型调整 `string` -> `int32`

### HuaweiCloud SDK CES

- _新增特性_
  - 支持以下接口：
    - `ListAlarmRules`
    - `CreateAlarmRules`
    - `BatchDeleteAlarmRules`
    - `BatchEnableAlarmRules`
    - `ListAlarmRuleResources`
    - `DeleteAlarmRuleResources`
    - `AddAlarmRuleResources`
    - `ListAlarmRulePolicies`
    - `UpdateAlarmRulePolicies`
    - `ListAgentDimensionInfo`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAlarmHistories`:
    - 新增响应参数 `datapoints`
    - 移除响应参数 `data_points`、`type`、`notification_list`、`type`、`notification_list`
    - 响应参数`status`类型调整 `string` -> `enum`
    - 响应参数`level`类型调整 `int32` -> `enum`
    - 响应参数`type`类型调整 `string` -> `enum`
    - 响应参数`period`类型调整 `integer` -> `enum`
    - 响应参数`value`类型调整 `float` -> `double`
    - 响应参数`suppress_duration`类型调整 `integer` -> `enum`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`ShowInstanceStatusInfo`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`移除请求参数 `instance_user_domain_name`、`instance_user_name`
  - 接口`CreateInstanceBy3rd`移除请求参数 `instance_user_domain_name`、`instance_user_name`

### HuaweiCloud SDK HSS

- _新增特性_
  - 支持以下接口：
    - `ListHostStatus`
    - `ListPasswordComplexity`
    - `ListRiskConfigCheckRules`
    - `ListRiskConfigHosts`
    - `ListRiskConfigs`
    - `ListSecurityEvents`
    - `ListVulnerabilities`
    - `ListWeakPasswordUsers`
    - `ShowCheckRuleDetail`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Image

- _新增特性_
  - 支持接口`RunImageDescription`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`NeutronListSecurityGroupRules`新增响应参数 `security_group_rules_links`

# 0.0.98 2022-07-14

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerBillsMonthlyBreakDown`新增响应参数 `effective_tag_pairs`、`cost_unit_pairs`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateDomainFullConfig`新增请求参数 `tls_version`
  - 接口`ShowDomainFullConfig`新增响应参数 `tls_version`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowInstance`新增响应参数 `tags`、`cpu_type`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPublicips`响应参数`create_time`类型调整 `date-time` -> `string`
  - 接口`ShowPublicip`响应参数`create_time`类型调整 `date-time` -> `string`

### HuaweiCloud SDK Image

- _新增特性_
  - 支持接口`RunImageMediaTagging`、`RunImageMainObjectDetection`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPorts`新增响应参数 `port_filter`、`ovs_hybrid_plug`
  - 接口`UpdatePort`新增响应参数 `port_filter`、`ovs_hybrid_plug`
  - 接口`ShowPort`新增响应参数 `port_filter`、`ovs_hybrid_plug`
  - 接口`CreateSecurityGroup`新增响应参数 `remote_address_group_id`
  - 接口`ListSecurityGroups`新增响应参数 `remote_address_group_id`
  - 接口`ShowSecurityGroup`新增响应参数 `remote_address_group_id`
  - 接口`ListSecurityGroupRules`新增响应参数 `remote_address_group_id`
  - 接口`ShowSecurityGroupRule`新增响应参数 `remote_address_group_id`
  - 接口`NeutronListSecurityGroups`新增响应参数 `remote_address_group_id`
  - 接口`NeutronUpdateSecurityGroup`新增响应参数 `remote_address_group_id`
  - 接口`NeutronShowSecurityGroup`新增响应参数 `remote_address_group_id`
  - 接口`NeutronListSecurityGroupRules`新增响应参数 `remote_address_group_id`
  - 接口`NeutronShowSecurityGroupRule`新增响应参数 `remote_address_group_id`

# 0.0.97 2022-07-07

### HuaweiCloud SDK APM

- _新增特性_
  - 支持接口`SearchApplication`、`SaveMonitorItemConfig`、`ListEnvMonitorItem`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`UpdateClusterEip`、`ShowClusterEndpoints`、`ShowVersion`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListNodes`:
    - 新增响应参数 `isStatic`、`privateIPv6IP`
    - 响应参数`key`、`effect`改为必填
  - 接口`CreateNode`:
    - 新增请求参数 `isStatic`
    - 请求参数`key`、`effect`改为必填
  - 接口`ShowNode`:
    - 新增响应参数 `isStatic`、`privateIPv6IP`
    - 响应参数`key`、`effect`改为必填
  - 接口`DeleteNode`:
    - 新增响应参数 `isStatic`、`privateIPv6IP`
    - 响应参数`key`、`effect`改为必填
  - 接口`UpdateNode`:
    - 新增响应参数 `isStatic`、`privateIPv6IP`
    - 响应参数`key`、`effect`改为必填
  - 接口`AddNode`请求参数`key`、`effect`改为必填
  - 接口`ResetNode`请求参数`key`、`effect`改为必填
  - 接口`ListNodePools`:
    - 新增响应参数 `isStatic`
    - 响应参数`key`、`effect`改为必填
  - 接口`CreateNodePool`:
    - 新增请求参数 `isStatic`
    - 请求参数`key`、`effect`改为必填
  - 接口`ShowNodePool`:
    - 新增响应参数 `isStatic`
    - 响应参数`key`、`effect`改为必填
  - 接口`DeleteNodePool`:
    - 新增响应参数 `isStatic`
    - 响应参数`key`、`effect`改为必填
  - 接口`UpdateNodePool`:
    - 新增响应参数 `isStatic`
    - 请求参数`key`、`effect`改为必填
    - 响应参数`key`、`effect`改为必填

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateMixJob`新增请求参数 `filling_policy`
  - 接口`UpdateMixJob`:
    - 新增请求参数 `filling_policy`
    - 新增响应参数 `filling_policy`
  - 接口`ShowMixJob`新增响应参数 `filling_policy`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClustersDetails`移除响应参数 `CREATING`
  - 接口`ShowClusterDetail`移除响应参数 `CREATING`
  - 接口`UpdateUnbindPublic`新增请求参数 `unBindPublicReq`
  - 接口`ListYmlsJob`:
    - 新增响应参数 `configurations`
    - 移除响应参数 `configList`
  - 接口`ListYmls`:
    - 新增响应参数 `configurations`
    - 移除响应参数 `configList`

### HuaweiCloud SDK CTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateNotification`:
    - 新增请求参数 `filter`
    - 新增响应参数 `filter`
  - 接口`CreateNotification`新增请求参数 `filter`
  - 接口`ListNotifications`新增响应参数 `filter`

### HuaweiCloud SDK ELB

- _新增特性_
  - 支持接口`ListMasterSlavePools`、`CreateMasterSlavePool`、`ShowMasterSlavePool`、`DeleteMasterSlavePool`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSystemSecurityPolicies`:
    - 响应参数`protocols`类型调整 `array` -> `string`
    - 响应参数`ciphers`类型调整 `array` -> `string`
  - 接口`ListQuotaDetails`新增请求参数 `X-Auth-Token`
  - 接口`ListAvailabilityZones`新增请求参数 `public_border_group`
  - 接口`CreateLoadBalancer`:
    - 新增请求参数 `id`、`global_eip_ids`
    - 移除请求参数 `min_l4_flavor_id`
    - 请求参数`X-Auth-Token`改为必填
  - 接口`ListLoadBalancers`:
    - 新增响应参数 `global_eips`、`public_border_group`
    - 移除响应参数 `min_l4_flavor_id`
    - 请求参数`X-Auth-Token`改为必填
  - 接口`UpdateLoadBalancer`:
    - 新增请求参数 `cloud_service_console_url`
    - 新增响应参数 `global_eips`、`public_border_group`
    - 移除请求参数 `min_l4_flavor_id`
    - 移除响应参数 `min_l4_flavor_id`
  - 接口`ShowLoadBalancer`:
    - 新增响应参数 `global_eips`、`public_border_group`
    - 移除响应参数 `min_l4_flavor_id`
  - 接口`ChangeLoadbalancerChargeMode`新增请求参数 `X-Auth-Token`
  - 接口`CreateCertificate`移除请求参数 `enc_certificate`、`enc_private_key`
  - 接口`ListCertificates`移除响应参数 `enc_certificate`、`enc_private_key`
  - 接口`UpdateCertificate`:
    - 移除请求参数 `enc_certificate`、`enc_private_key`
    - 移除响应参数 `enc_certificate`、`enc_private_key`
  - 接口`ShowCertificate`移除响应参数 `enc_certificate`、`enc_private_key`
  - 接口`CreateListener`新增请求参数 `quic_config`
  - 接口`ListListeners`新增响应参数 `quic_config`
  - 接口`UpdateListener`:
    - 新增请求参数 `quic_config`
    - 新增响应参数 `quic_config`
  - 接口`ShowListener`新增响应参数 `quic_config`
  - 接口`CreatePool`新增请求参数 `vpc_id`、`type`
  - 接口`ListPools`:
    - 新增请求参数 `vpc_id`、`type`
    - 新增响应参数 `created_at`、`updated_at`、`vpc_id`、`type`
  - 接口`UpdatePool`:
    - 新增请求参数 `X-Auth-Token`、`vpc_id`、`type`
    - 新增响应参数 `created_at`、`updated_at`、`vpc_id`、`type`
  - 接口`ShowPool`新增响应参数 `created_at`、`updated_at`、`vpc_id`、`type`
  - 接口`CreateMember`请求参数`project_id`类型调整 `enum` -> `string`
  - 接口`ListMembers`:
    - 新增响应参数 `status`、`loadbalancers`、`created_at`、`updated_at`
    - 移除请求参数 `instance_id`
  - 接口`UpdateMember`新增响应参数 `status`、`loadbalancers`、`created_at`、`updated_at`
  - 接口`ShowMember`新增响应参数 `status`、`loadbalancers`、`created_at`、`updated_at`
  - 接口`ListAllMembers`新增响应参数 `status`、`loadbalancers`、`created_at`、`updated_at`
  - 接口`ListHealthMonitors`新增响应参数 `created_at`、`updated_at`
  - 接口`UpdateHealthMonitor`新增响应参数 `created_at`、`updated_at`
  - 接口`ShowHealthMonitor`新增响应参数 `created_at`、`updated_at`
  - 接口`CreateL7Policy`新增请求参数 `redirect_pools_config`
  - 接口`ListL7Policies`新增响应参数 `redirect_pools_config`、`created_at`、`updated_at`
  - 接口`UpdateL7Policy`:
    - 新增请求参数 `redirect_pools_config`
    - 新增响应参数 `redirect_pools_config`、`created_at`、`updated_at`
  - 接口`ShowL7Policy`新增响应参数 `redirect_pools_config`、`created_at`、`updated_at`
  - 接口`BatchUpdatePoliciesPriority`新增请求参数 `X-Auth-Token`
  - 接口`ListL7Rules`新增响应参数 `created_at`、`updated_at`
  - 接口`UpdateL7Rule`新增响应参数 `created_at`、`updated_at`
  - 接口`ShowL7Rule`新增响应参数 `created_at`、`updated_at`
  - 接口`UpdateIpList`:
    - 移除请求参数 `name`、`ip_list`、`description`
    - 请求参数`X-Auth-Token`改为必填
  - 接口`BatchDeleteIpList`:
    - 新增请求参数 `BatchDeleteIpListRequestBody`
    - 移除请求参数 `BatchDeleteIpGroupIpListRequestBody`
    - 请求参数`X-Auth-Token`改为必填
  - 接口`BatchCreateMembers`:
    - 新增请求参数 `BatchCreateMembersRequestBody`
    - 新增响应参数 `status`
    - 移除请求参数 `BatchCreateMemberRequestBody`
  - 接口`BatchDeleteMembers`:
    - 新增请求参数 `BatchDeleteMembersRequestBody`
    - 移除请求参数 `BatchDeleteMemberRequestBody`
  - 接口`UpdateLogtank`:
    - 新增请求参数 `UpdateLogtankRequestBody`
    - 移除请求参数 `logtank`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除以下接口：
    - `CreateAccessCode`
    - `ShowProtocolMappings`
    - `UploadProtocolMappings`
    - `BatchUpdateConfigs`
    - `ShowExternalEntity`
  - 接口`CreateEdgeApp`请求参数`function_type`新增枚举值`COMPOSITE_APPLICATION`、`DATA_COLLECTION`
  - 接口`BatchListEdgeAppVersions`新增响应参数 `sdk_version`、`deploy_multi_instance`
  - 接口`CreateEdgeApplicationVersion`新增请求参数 `sdk_version`、`deploy_multi_instance`、`ext_devices`、`tcp_socket`、`period_seconds`、`failure_threshold`、`tcp_socket`、`period_seconds`、`failure_threshold`
  - 接口`UpdateEdgeApplicationVersion`:
    - 新增请求参数 `deploy_multi_instance`、`sdk_version`、`ext_devices`、`tcp_socket`、`period_seconds`、`failure_threshold`、`tcp_socket`、`period_seconds`、`failure_threshold`
    - 新增响应参数 `sdk_version`、`deploy_multi_instance`
    - 请求参数`arch`改为非必填
  - 接口`ShowEdgeApplicationVersion`新增响应参数 `deploy_multi_instance`、`sdk_version`、`tcp_socket`、`period_seconds`、`failure_threshold`、`tcp_socket`、`period_seconds`、`failure_threshold`、`ext_devices`
  - 接口`UpdateEdgeApplicationVersionState`新增响应参数 `sdk_version`、`deploy_multi_instance`
  - 接口`CreateEdgeNode`新增请求参数 `edge_node_id`、`verify_code`、`time_out`、`arch`、`base_path`、`hardware_model`
  - 接口`ShowEdgeNode`新增响应参数 `ha_config`、`hardware_model`
  - 接口`CreateInstallCmd`新增请求参数 `CreateInstallCmdRequestBody`
  - 接口`BatchListModules`:
    - 新增响应参数 `control_status`、`container_settings`
    - 响应参数`state`新增枚举值`DELETE_SUCCESS`、`STOPPED`
    - 响应参数`function_type`新增枚举值`GATEWAY_MANAGER`、`COMPOSITE_APPLICATION`、`DATA_COLLECTION`
  - 接口`CreateModule`新增请求参数 `module_name`、`container_settings`
  - 接口`UpdateModule`:
    - 新增请求参数 `module_name`、`container_settings`
    - 新增响应参数 `control_status`
    - 请求参数`app_version`改为非必填
    - 响应参数`state`新增枚举值`DELETE_SUCCESS`、`STOPPED`
    - 响应参数`function_type`新增枚举值`GATEWAY_MANAGER`、`COMPOSITE_APPLICATION`、`DATA_COLLECTION`
  - 接口`ShowModule`:
    - 新增响应参数 `control_status`、`container_settings`
    - 响应参数`state`新增枚举值`DELETE_SUCCESS`、`STOPPED`
    - 响应参数`function_type`新增枚举值`GATEWAY_MANAGER`、`COMPOSITE_APPLICATION`、`DATA_COLLECTION`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeVatInvoice`新增响应参数 `title`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`SetSensitiveSlowLog`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RMS

- _新增特性_
  - 支持接口`ListSchemas`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SIS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeShortAudio`请求参数`property`新增枚举值`chinese_16k_travel`
  - 接口`PushTranscriberJobs`请求参数`property`新增枚举值`chinese_16k_media`
  - 接口`CollectTranscriberJob`新增响应参数 `audio_duration`
  - 接口`RunTts`请求参数`property`新增枚举值`chinese_huaxiaomei_common`、`chinese_huaxiaofei_common`

# 0.0.96 2022-06-30

### HuaweiCloud SDK Core

- _新增特性_
  - 支持联邦认证
  - 支持认证信息管理
- _解决问题_
  - [Issue 53](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/53) 修复客户端发送请求出现panic的问题
- _特性变更_
  - 无

### HuaweiCloud SDK UGO

- _新增特性_
  - 支持数据库和应用迁移服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SendVerificationMessageCode`请求参数`mobile_phone`改为必填

### HuaweiCloud SDK BSSINTL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SendVerificationMessageCode`请求参数`email`改为必填

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowDomainFullConfig`新增响应参数 `cache_url_parameter_filter`
  - 接口`UpdateDomainFullConfig`新增请求参数 `cache_url_parameter_filter`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`UploadExtensionFile`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DNS

- _新增特性_
  - 支持接口`CreateRecordSetWithBatchLines`、`BatchUpdateRecordSetWithLine`、`BatchDeleteRecordSetWithLine`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRecordSetWithLine`请求参数`records`改为非必填

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持接口`UpdateFunctionMaxInstanceConfig`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `lb_port`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListComponentInfos`新增响应参数 `total_count`

### HuaweiCloud SDK IEF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTag`:
    - 新增请求参数 `CreateTagRequestBody`
    - 移除请求参数 `tag`
  - 接口`ListEdgeNodes`新增请求参数 `sort`、`state`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeHkIdCard`、`RecognizeCambodianIdCard`、`RecognizeExitEntryPermit`、`RecognizeMainlandTravelPermit`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeGeneralText`响应参数`direction`类型调整 `int32` -> `float`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`CreateProjectModule`、`ListProjectModules`、`UpdateProjectModule`、`DeleteProjectModule`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK UGO

- _新增特性_
  - 支持以下接口：
    - `ListQuotas`
    - `RunSqlConversion`
    - `ListEvaluationProjects`
    - `CreateEvaluationProject`
    - `ShowEvaluationProjectStatus`
    - `ShowEvaluationProjectDetail`
    - `ConfirmTargetDbType`
    - `DeleteEvaluationProject`
    - `ListMigrationProjects`
    - `CreateMigrationProject`
    - `ShowMigrationProjectStatus`
    - `CheckPermission`
    - `ListPermissionCheckResult`
    - `ShowMigrationProjectDetail`
    - `CommitSyntaxConversion`
    - `ListSyntaxConversionProgress`
    - `CommitVerification`
    - `ListVerificationProgress`
    - `DownloadFailureReport`
    - `DeleteMigrationProject`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`MigrateSqlStatement`
  - 接口`ListApiVersions`:
    - 新增请求参数 `X-Auth-Token`
    - 响应参数`id`、`links`、`version`、`status`、`updated`改为必填
  - 接口`ShowApiVersionInfo`:
    - 新增请求参数 `X-Auth-Token`
    - 响应参数`id`、`links`、`version`、`status`、`updated`改为必填

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持接口`ListOverviewsClassification`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListStatistics`新增响应参数 `host`
  - 接口`ListHost`新增响应参数 `timeout_config`
  - 接口`ShowHost`新增响应参数 `timeout_config`
  - 接口`UpdateHost`:
    - 新增请求参数 `block_page`、`traffic_mark`、`flag`、`extend`、`http2_enable`、`ipv6_enable`、`lb_algorithm`、`timeout_config`
    - 新增响应参数 `http2_enable`、`ipv6_enable`、`lb_algorithm`、`timeout_config`
  - 接口`UpdatePremiumHost`:
    - 新增请求参数 `mode`、`locked`、`protect_status`、`access_status`、`timestamp`、`pool_ids`、`block_page`、`traffic_mark`、`flag`、`extend`、`circuit_breaker`、`timeout_config`
    - 新增响应参数 `timeout_config`
  - 接口`ShowPremiumHost`新增响应参数 `timeout_config`

# 0.0.95 2022-06-25

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTaskDetail`新增响应参数 `cyclomatic_complexity_per_file`、`file_duplication_ratio`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持以下接口：
    - `ShowEntityConfiguration`
    - `UpdateEntityConfiguration`
    - `ShowConfigurationParameter`
    - `UpdateConfigurationParameter`
    - `DeleteConfiguration`
    - `ListConfigurations`
    - `CreateConfiguration`
    - `SwitchConfiguration`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数 `configurations`、`charge_info`
  - 接口`ResizeInstanceVolume`:
    - 新增请求参数 `is_auto_pay`
    - 新增响应参数 `order_id`
  - 接口`AddShardingNode`:
    - 新增请求参数 `is_auto_pay`
    - 新增响应参数 `order_id`
  - 接口`ResizeInstance`:
    - 新增请求参数 `is_auto_pay`
    - 新增响应参数 `order_id`
  - 接口`RestoreNewInstance`新增请求参数 `configurations`、`charge_info`

# 0.0.94 2022-06-23

### HuaweiCloud SDK CSE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UploadKie`:
    - 新增请求参数 `upload_file`
    - 移除请求参数 `UploadFile`

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`:
    - 新增请求参数 `httpsEnable`、`authorityEnable`、`adminPwd`
    - 请求参数`availability_zone`改为必填
  - 接口`ListClustersDetails`:
    - 新增响应参数 `publicKibanaResp`、`elbWhiteList`、`publicIp`、`bandwidthSize`、`diskEncrypted`、`backupAvailable`、`enterpriseProjectId`、`ip`
    - 移除响应参数 `enterprise_project_id`
  - 接口`ShowClusterDetail`:
    - 新增响应参数 `publicKibanaResp`、`elbWhiteList`、`publicIp`、`vpcId`、`subnetId`、`securityGroupId`、`bandwidthSize`、`diskEncrypted`、`backupAvailable`、`enterpriseProjectId`、`period`、`ip`
    - 移除响应参数 `enterprise_project_id`
  - 接口`ListFlavors`:
    - 新增响应参数 `type`、`availableAZ`
  - 接口`StartVpecp`:
    - 请求参数`endpointWithDnsName`类型调整 `string` -> `boolean`
  - 接口`CreateCluster`:
    - 新增请求参数 `payInfo`
    - 新增响应参数 `cluster`
    - 移除响应参数 `schema`
  - 接口`RestartCluster`:
    - 新增请求参数 `RestartClusterRequestBody`
    - 移除请求参数 `X-Auth-Token`、`RollingRestartReq`

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持接口`ListAvailableZone`、`UpdateTuningParams`
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateJobs`新增请求参数 `master_az`、`slave_az`
  - 接口`BatchSetPolicy`新增请求参数 `slot_name`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListQuotaDetails`:
    - 新增请求参数 `quota_key`
    - 移除请求参数 `type`
  - 接口`ListListeners`:
    - 新增请求参数 `loadbalancer_id`、`connection_limit`、`admin_state_up`、`http2_enable`、`enterprise_project_id`
    - 移除请求参数 `member_timeout`、`client_timeout`、`keepalive_timeout`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `subnet_id`
  - 接口`ExpandInstanceNode`新增请求参数 `subnet_id`

### HuaweiCloud SDK VSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AuthorizeDomains`:
    - 新增响应参数 `usage_notice`
    - 请求参数`auth_mode`新增枚举值`free`

# 0.0.93 2022-06-19

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复无法通过环境变量配置Region的问题
- _特性变更_
  - 无

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 修复接口`ListRangeQueryAomPromGet`的响应参数`data`类型错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeIdDocument`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RMS

- _新增特性_
  - 支持接口`RunQuery`、`CreateStoredQuery`、`ListStoredQueries`、`ShowStoredQuery`、`UpdateStoredQuery`、`DeleteStoredQuery`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CheckProductHealthy`请求参数`domain_id`、`project_id`、`region`、`company_name`、`product_name`改为必填
  - 接口`ImportEvents`:
    - 请求参数`type`、`domain_id`、`business`、`checkitem_id`、`checkpoint_id`、`spec_id`、`patch_id`改为必填
    - 请求参数`type`、`domain_id`、`product_feature`、`arrive_time`、`verification_state`改为非必填

### HuaweiCloud SDK VSS

- _新增特性_
  - 支持接口`ShowReportStatus`、`DownloadTaskReport`、`ExecuteGenerateReport`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.92 2022-06-09

### HuaweiCloud SDK CSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`:
    - 新增响应参数 `schema`
    - 移除响应参数 `id`、`name`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateFunctionCode`响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`ShowFunctionCode`响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`UpdateFunctionConfig`:
    - 新增请求参数 `domain_names`
    - 新增响应参数 `domain_names`
    - 响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`ShowFunctionConfig`响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`CreateFunctionVersion`响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`ListStatistics`新增请求参数 `option`
  - 接口`CreateDependency`响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`ListDependencies`响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`UpdateDependency`响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`ShowDependency`响应参数`runtime`类型调整 `string` -> `enum`
  - 接口`UpdateEvent`请求参数`content`改为必填
  - 接口`ListFunctionAsyncInvocations`新增请求参数 `marker`
  - 接口`BatchDeleteWorkflows`请求参数`workflow_urns`改为必填
  - 接口`CreateWorkflow`请求参数`name`、`trigger_name`、`trigger_type`、`bucket`、`events`、`prefix`、`suffix`、`start`、`name`、`operation`、`id`、`name`、`type`、`end`、`transition`、`ref_name`、`arguments`、`constants`、`name`改为必填
  - 接口`StartWorkflowExecution`请求参数`input`改为必填
  - 接口`ShowWorkflowExecution`新增响应参数 `node_name`、`execution_id`、`request_id`
  - 接口`UpdateWorkFlow`请求参数`name`、`trigger_name`、`trigger_type`、`bucket`、`events`、`prefix`、`suffix`、`start`、`name`、`operation`、`id`、`name`、`type`、`end`、`transition`、`ref_name`、`arguments`、`constants`、`name`改为必填
  - 接口`ShowWorkFlow`:
    - 新增响应参数 `lts_group_id`、`lts_stream_id`
    - 响应参数`name`、`trigger_name`、`trigger_type`、`bucket`、`events`、`prefix`、`suffix`、`start`、`name`、`operation`、`id`、`name`、`type`、`end`、`transition`、`ref_name`、`arguments`、`constants`、`name`改为必填
  - 接口`StartSyncWorkflowExecution`请求参数`input`改为必填

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持以下接口：
    - `BatchSetTags`
    - `ListTags`
    - `CreateTag`
    - `DeleteTag`
    - `BatchSetAttributes`
    - `ListAttributes`
    - `CreateAttribute`
    - `UpdateAttribute`
    - `EnableAttribute`
    - `DisableAttribute`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSimCards`新增请求参数 `tag_id`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 支持接口`RunTextModeration`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClusters`新增响应参数 `availabilityZoneId`
  - 接口`ShowClusterDetails`新增响应参数 `availabilityZoneId`
  - 接口`ListHosts`新增响应参数 `availability_zone_id`、`tags`

# 0.0.91 2022-06-02

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`新增请求参数 `instance_id`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ChangeGaussMySqlProxySpecification`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Meeting

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SearchCorpVmr`新增响应参数 `maxAudienceParties`、`expireDate`、`commercialMaxAudienceParties`

### HuaweiCloud SDK NAT

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListNatGateways`:
    - 请求参数`created_at`类型调整 `date-time` -> `string`
    - 响应参数`created_at`类型调整 `date-time` -> `string`
  - 接口`UpdateNatGateway`响应参数`created_at`类型调整 `date-time` -> `string`
  - 接口`ShowNatGateway`响应参数`created_at`类型调整 `date-time` -> `string`
  - 接口`ListNatGatewayDnatRules`:
    - 请求参数`created_at`类型调整 `date-time` -> `string`
    - 响应参数`created_at`类型调整 `date-time` -> `string`
  - 接口`UpdateNatGatewayDnatRule`响应参数`created_at`类型调整 `date-time` -> `string`
  - 接口`ShowNatGatewayDnatRule`响应参数`created_at`类型调整 `date-time` -> `string`
  - 接口`ListNatGatewaySnatRules`:
    - 请求参数`created_at`类型调整 `date-time` -> `string`
    - 响应参数`created_at`类型调整 `date-time` -> `string`
  - 接口`UpdateNatGatewaySnatRule`响应参数`created_at`类型调整 `date-time` -> `string`
  - 接口`ShowNatGatewaySnatRule`响应参数`created_at`类型调整 `date-time` -> `string`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeHealthCode`新增响应参数 `words_block_count`、`words_block_list`
  - 接口`RecognizePcrTestRecord`响应参数`confidence`类型调整 `float` -> `object`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持以下接口：
    - `ListIssueCustomFields`
    - `ListIssuesSfV4`
    - `ListProjectIssuesRecordsV4`
    - `ListWorkitemStatusRecordsV4`
    - `ListWorkitems`
    - `ShowIssuesWrokFlowConfig`
    - `ShowWorkItemWrokflowConfig`
    - `ListAssociatedIssues`
    - `ListAssociatedWikis`
    - `ListIssueAssociatedCommits`
    - `ListAssociatedTestCases`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateIssueV4`:
    - 新增请求参数 `new_custom_fields`
    - 新增响应参数 `new_custom_fields`、`new_name`
  - 接口`ListIssuesV4`:
    - 新增请求参数 `custom_fields`
    - 新增响应参数 `new_custom_fields`、`new_name`
  - 接口`ShowIssueV4`新增响应参数 `new_custom_fields`、`new_name`
  - 接口`UpdateIssueV4`:
    - 新增请求参数 `new_custom_fields`
    - 新增响应参数 `new_custom_fields`、`new_name`
  - 接口`ListChildIssuesV4`新增响应参数 `new_custom_fields`、`new_name`
  - 接口`CreateSystemIssueV4`:
    - 新增请求参数 `new_custom_fields`
    - 新增响应参数 `new_custom_fields`、`new_name`

# 0.0.90 2022-05-26

### HuaweiCloud SDK BCS

- _新增特性_
  - 支持接口`HandleUnionMemberQuitList`、`BatchRemovePeersFromChannel`、`DeleteChannel`
- _解决问题_
  - 无
- _特性变更_
  - 接口`DeleteBlockchain`新增请求参数 `is_delete_ief`、`is_delete_lightpeer`、`ief_nodes_id`
  - 接口`CreateNewBlockchain`:
    - 新增请求参数 `cluster_platform_type`
    - 移除请求参数 `user_name`、`password`
    - 请求参数`node_flavor`类型调整 `int64` -> `string`
    - 请求参数`cce_flavor`类型调整 `int64` -> `string`
    - 请求参数`init_node_pwd`类型调整 `int64` -> `string`
    - 请求参数`az`类型调整 `int64` -> `string`
    - 请求参数`cluster_platform_type`类型调整 `int64` -> `string`
  - 接口`CreateBlockchainCertByUserName`新增请求参数 `CreateBlockchainCertByUserNameRequestBody`
  - 接口`UnfreezeCert`新增请求参数 `file`
  - 接口`FreezeCert`新增请求参数 `file`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `features`、`sub_status`
  - 接口`ShowInstance`新增响应参数 `features`、`transparent_client_ip_enable`、`sub_status`
  - 接口`ResizeInstance`新增请求参数 `execute_immediately`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AttachServerVolume`新增请求参数 `volume_type`、`hw:passthrough`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持接口`ListComponentInfos`、`SwitchShard`、`ResizeInstanceFlavor`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowBackupPolicy`移除响应参数 `rate_limit`、`prefetch_block`、`filesplit_size`
  - 接口`ListDbUsers`响应参数`memberof`类型调整 `object` -> `string`

### HuaweiCloud SDK KMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ValidateSignature`:
    - 新增响应参数 `signature_valid`
    - 移除响应参数 `signature_vaild`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateAomMappingRules`:
    - 新增请求参数 `deployments_prefix`
    - 新增响应参数 `deployments_prefix`
  - 接口`UpdateAomMappingRules`:
    - 新增请求参数 `deployments_prefix`
    - 新增响应参数 `deployments_prefix`
  - 接口`ShowAomMappingRules`新增响应参数 `deployments_prefix`
  - 接口`ShowAomMappingRule`新增响应参数 `deployments_prefix`

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunCheckResult`新增响应参数 `ocr_text`、`error_code`、`error_msg`
  - 接口`RunImageBatchModeration`:
    - 新增请求参数 `moderation_rule`、`ad_categories`、`show_ocr_text`
    - 新增响应参数 `ocr_text`、`error_code`、`error_msg`
  - 接口`RunTaskSumbit`新增请求参数 `moderation_rule`、`ad_categories`、`show_ocr_text`

# 0.0.89 2022-05-19

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEnvironmentVariablesV2`新增请求参数 `group_id`
  - 接口`UpdateSignatureKeyV2`新增响应参数 `update_time`、`create_time`、`id`、`name`、`sign_type`、`sign_key`、`sign_secret`、`sign_algorithm`
  - 接口`CreateInstanceV2`新增请求参数 `loadbalancer_provider`
  - 接口`ListInstancesV2`新增响应参数 `loadbalancer_provider`
  - 接口`ShowDetailsOfInstanceV2`新增响应参数 `endpoint_service`、`endpoint_services`、`node_ips`、`publicips`、`privateips`、`loadbalancer_provider`
  - 接口`UpdateInstanceV2`新增响应参数 `endpoint_service`、`endpoint_services`、`node_ips`、`publicips`、`privateips`、`loadbalancer_provider`

### HuaweiCloud SDK CSE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEngines`移除请求参数 `X-Enterprise-Project-ID`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ResizeServer`新增请求参数 `dry_run`
  - 接口`ResizePostPaidServer`新增请求参数 `dry_run`
  - 接口`AttachServerVolume`新增请求参数 `dry_run`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClusters`:
    - 新增响应参数 `GroupName`、`NodeNum`、`NodeSize`、`NodeSpecId`、`VmProductId`、`VmSpecCode`、`NodeProductId`、`RootVolumeSize`、`RootVolumeProductId`、`RootVolumeType`、`RootVolumeResourceSpecCode`、`RootVolumeResourceType`、`DataVolumeType`、`DataVolumeCount`、`DataVolumeSize`、`DataVolumeProductId`、`DataVolumeResourceSpecCode`、`DataVolumeResourceType`、`GroupName`、`NodeNum`、`NodeSize`、`NodeSpecId`、`VmProductId`、`VmSpecCode`、`NodeProductId`、`RootVolumeSize`、`RootVolumeProductId`、`RootVolumeType`、`RootVolumeResourceSpecCode`、`RootVolumeResourceType`、`DataVolumeType`、`DataVolumeCount`、`DataVolumeSize`、`DataVolumeProductId`、`DataVolumeResourceSpecCode`、`DataVolumeResourceType`
    - 移除响应参数 `groupName`、`nodeNum`、`nodeSize`、`nodeSpecId`、`vmProductId`、`vmSpecCode`、`nodeProductId`、`rootVolumeSize`、`rootVolumeProductId`、`rootVolumeType`、`rootVolumeResourceSpecCode`、`rootVolumeResourceType`、`dataVolumeType`、`dataVolumeCount`、`dataVolumeSize`、`dataVolumeProductId`、`dataVolumeResourceSpecCode`、`dataVolumeResourceType`、`groupName`、`nodeNum`、`nodeSize`、`nodeSpecId`、`vmProductId`、`vmSpecCode`、`nodeProductId`、`rootVolumeSize`、`rootVolumeProductId`、`rootVolumeType`、`rootVolumeResourceSpecCode`、`rootVolumeResourceType`、`dataVolumeType`、`dataVolumeCount`、`dataVolumeSize`、`dataVolumeProductId`、`dataVolumeResourceSpecCode`、`dataVolumeResourceType`
  - 接口`ShowClusterDetails`:
    - 新增响应参数 `GroupName`、`NodeNum`、`NodeSize`、`NodeSpecId`、`VmProductId`、`VmSpecCode`、`NodeProductId`、`RootVolumeSize`、`RootVolumeProductId`、`RootVolumeType`、`RootVolumeResourceSpecCode`、`RootVolumeResourceType`、`DataVolumeType`、`DataVolumeCount`、`DataVolumeSize`、`DataVolumeProductId`、`DataVolumeResourceSpecCode`、`DataVolumeResourceType`、`GroupName`、`NodeNum`、`NodeSize`、`NodeSpecId`、`VmProductId`、`VmSpecCode`、`NodeProductId`、`RootVolumeSize`、`RootVolumeProductId`、`RootVolumeType`、`RootVolumeResourceSpecCode`、`RootVolumeResourceType`、`DataVolumeType`、`DataVolumeCount`、`DataVolumeSize`、`DataVolumeProductId`、`DataVolumeResourceSpecCode`、`DataVolumeResourceType`
    - 移除响应参数 `groupName`、`nodeNum`、`nodeSize`、`nodeSpecId`、`vmProductId`、`vmSpecCode`、`nodeProductId`、`rootVolumeSize`、`rootVolumeProductId`、`rootVolumeType`、`rootVolumeResourceSpecCode`、`rootVolumeResourceType`、`dataVolumeType`、`dataVolumeCount`、`dataVolumeSize`、`dataVolumeProductId`、`dataVolumeResourceSpecCode`、`dataVolumeResourceType`、`groupName`、`nodeNum`、`nodeSize`、`nodeSpecId`、`vmProductId`、`vmSpecCode`、`nodeProductId`、`rootVolumeSize`、`rootVolumeProductId`、`rootVolumeType`、`rootVolumeResourceSpecCode`、`rootVolumeResourceType`、`dataVolumeType`、`dataVolumeCount`、`dataVolumeSize`、`dataVolumeProductId`、`dataVolumeResourceSpecCode`、`dataVolumeResourceType`
  - 接口`UpdateClusterScaling` 的请求参数 `data_volume_type`、`data_volume_count`、`data_volume_size`改为必填

### HuaweiCloud SDK VOD

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UploadMetaDataByUrl`请求参数`video_type`新增枚举值`M3U8`
  - 接口`PublishAssets`新增响应参数 `sign_url`
  - 接口`UnpublishAssets`新增响应参数 `sign_url`
  - 接口`ShowAssetMeta`新增响应参数 `sign_url`
  - 接口`ShowAssetDetail`新增响应参数 `sign_url`
  - 接口`ShowTakeOverTaskDetails`新增响应参数 `sign_url`
  - 接口`ShowTakeOverAssetDetails`新增响应参数 `sign_url`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPorts`响应参数`device_owner`新增枚举值`neutron:VIP_PORT`, 移除枚举值`network:VIP_PORT`
  - 接口`UpdatePort`响应参数`device_owner`新增枚举值`neutron:VIP_PORT`, 移除枚举值`network:VIP_PORT`
  - 接口`ShowPort`响应参数`device_owner`新增枚举值`neutron:VIP_PORT`, 移除枚举值`network:VIP_PORT`

# 0.0.88 2022-05-12

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 支持接口`AddProtectBranchV2`、`AddTagV2`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateTask`:
    - 请求参数`bodys`类型调整 `string` -> `array`
    - 请求参数`data`类型调整 `string` -> `object`
    - 响应参数`bodys`类型调整 `string` -> `array`
    - 响应参数`data`类型调整 `string` -> `object`
  - 接口`ShowTask`:
    - 响应参数`bodys`类型调整 `string` -> `array`
    - 响应参数`data`类型调整 `string` -> `object`
  - 接口`UpdateCase`:
    - 请求参数`bodys`类型调整 `string` -> `array`
    - 请求参数`data`类型调整 `string` -> `object`
  - 接口`UpdateTemp`:
    - 请求参数`bodys`类型调整 `string` -> `array`
    - 请求参数`data`类型调整 `string` -> `object`

### HuaweiCloud SDK CSS

- _新增特性_
  - 支持接口`DownloadCert`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DWS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClusterDetails`新增响应参数 `nodes`
  - 接口`ListClusters`新增响应参数 `nodes`

### HuaweiCloud SDK FRS

- _新增特性_
  - 支持以下接口：
    - `DetectLiveByUrlIntl`
    - `DetectLiveByFileIntl`
    - `DetectLiveByBase64Intl`
    - `DetectFaceByFileIntl`
    - `DetectFaceByUrlIntl`
    - `DetectFaceByBase64Intl`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持以下接口：
    - `ListBackups`
    - `CreateManualBackup`
    - `DeleteManualBackup`
    - `ListRestoreTimes`
    - `CreateRestoreInstance`
    - `CreateDatabase`
    - `CreateDbUser`
    - `CreateDatabaseSchemas`
    - `AllowDbPrivileges`
    - `SetDbUserPwd`
    - `ListDatabases`
    - `ListDbUsers`
    - `ListDatabaseSchemas`
    - `ShowBackupPolicy`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSimCards`:
    - 新增请求参数 `min_used_flow`、`max_used_flow`、`min_left_flow`、`max_left_flow`
    - 移除请求参数 `min_flow`、`max_flow`

### HuaweiCloud SDK IAM

- _新增特性_
  - 支持接口`ShowDomainRoleAssignments`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunImageModeration`:
    - 新增请求参数 `show_ocr_text`
    - 新增响应参数 `ocr_text`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeHealthCode`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RestoreExistInstance`:
    - 新增请求参数 `RestoreExistingInstanceRequestBody`
    - 移除请求参数 `RestoreToExistingInstanceRequest`

### HuaweiCloud SDK RocketMQ

- _新增特性_
  - 支持分布式消息服务RocketMQ版
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.87 2022-05-05

### HuaweiCloud SDK Moderation

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunImageModeration`请求参数`image`类型调整 `byte` -> `string`

### HuaweiCloud SDK RES

- _新增特性_
  - 支持推荐系统服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.86 2022-04-28

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddOrUpdateServiceDiscoveryRules`请求参数`priority`类型调整 `string` -> `int32`
  - 接口`ListServiceDiscoveryRules`响应参数`priority`类型调整 `string` -> `int32`

### HuaweiCloud SDK CSE

- _新增特性_
  - 支持微服务引擎
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DevStar

- _新增特性_
  - 支持接口`ConfirmDeploymentJob`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDeploymentJobs`新增请求参数 `cci`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持接口`CancelAsyncInvocation`、`StartSyncWorkflowExecution`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFunctionStatistics`:
    - 响应参数`timestamp`类型调整 `int32` -> `int64`
    - 响应参数`value`类型调整 `int32` -> `double`
  - 接口`ListStatistics`:
    - 响应参数`timestamp`类型调整 `int32` -> `int64`
    - 响应参数`value`类型调整 `int32` -> `double`
  - 接口`ListFunctionAsyncInvokeConfig`新增响应参数 `enable_async_status_log`
  - 接口`ShowFunctionAsyncInvokeConfig`新增响应参数 `enable_async_status_log`
  - 接口`UpdateFunctionAsyncInvokeConfig`:
    - 新增请求参数 `enable_async_status_log`
    - 新增响应参数 `enable_async_status_log`
  - 接口`CreateWorkflow`:
    - 新增请求参数 `mode`、`express_config`
    - 请求参数`type`新增枚举值`End`
  - 接口`ShowWorkFlow`:
    - 新增响应参数 `mode`、`express_config`
    - 响应参数`type`新增枚举值`End`
  - 接口`UpdateWorkFlow`:
    - 新增请求参数 `mode`、`express_config`
    - 请求参数`type`新增枚举值`End`
  - 接口`ShowTenantMetric`:
    - 新增请求参数 `start_time`、`end_time`
    - 响应参数`timestamp`类型调整 `int32` -> `int64`
    - 响应参数`value`类型调整 `int32` -> `double`
  - 接口`ShowWorkFlowMetric`:
    - 新增请求参数 `start_time`、`end_time`
    - 响应参数`timestamp`类型调整 `int32` -> `int64`
    - 响应参数`value`类型调整 `int32` -> `double`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`ShowGaussMySqlProxyList`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateGaussMySqlReadonlyNode`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`
  - 接口`ExpandGaussMySqlInstanceVolume`移除响应参数 `size`、`order_id`
  - 接口`ListGaussMySqlDedicatedResources`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`
  - 接口`DeleteGaussMySqlProxy`新增请求参数 `CloseMysqlProxyRequest`
  - 接口`CreateGaussMySqlProxy`:
    - 新增请求参数 `proxy_name`、`proxy_mode`、`nodes_read_weight`
    - 请求参数`flavor_ref`、`node_num`改为必填
  - 接口`ExpandGaussMySqlProxy`新增请求参数 `proxy_id`
  - 接口`ListGaussMySqlErrorLog`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`
    - 请求参数`node_id`改为非必填
  - 接口`ListGaussMySqlSlowLog`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`
  - 接口`ListGaussMySqlConfigurations`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`
  - 接口`ShowGaussMySqlJobInfo`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`
  - 接口`ListInstanceTags`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`
  - 接口`ListProjectTags`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`
  - 接口`BatchTagAction`:
    - 新增请求参数 `X-Auth-Token`
    - 移除请求参数 `x-auth-token`

### HuaweiCloud SDK Meeting

- _新增特性_
  - 支持接口`ShowDepartment`
- _解决问题_
  - 无
- _特性变更_
  - 接口`SearchCorpDir`新增响应参数 `deptCodes`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 修复接口`RecognizeMyanmarDriverLicense`的响应体类型错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持接口`DeleteIgnoreRule`、`CreateIgnoreRule`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIgnoreRule`移除响应参数 `id`、`policyid`、`timestamp`、`description`、`status`、`url`、`rule`、`domain`、`url_logic`、`advanced`
  - 接口`ListValueList`新增响应参数 `producer`
  - 接口`ListEvent`新增响应参数 `process_time`、`request_body`
  - 接口`ShowEvent`:
    - 新增响应参数 `process_time`、`request_body`
    - 响应参数`headers`类型调整 `string` -> `object`
    - 响应参数`response_size`类型调整 `string` -> `int32`
    - 响应参数`response_time`类型调整 `string` -> `int64`
  - 接口`ListHost`响应参数`paid_type`类型调整 `string` -> `enum`
  - 接口`CreateHost`新增响应参数 `flag`、`http2_enable`
  - 接口`UpdateHostProtectStatus`新增响应参数 `protect_status`
  - 接口`DeleteHost`响应参数`paid_type`类型调整 `string` -> `enum`
  - 接口`UpdateHost`响应参数`protocol`类型调整 `string` -> `enum`
  - 接口`DeletePolicy`新增响应参数 `id`、`name`、`level`、`action`、`options`、`full_detection`、`hosts`、`bind_host`、`timestamp`、`extend`
  - 接口`UpdatePolicyRuleStatus`新增响应参数 `id`、`policyid`、`timestamp`、`description`、`status`
  - 接口`DeletePrivacyRule`新增响应参数 `id`、`policyid`、`timestamp`、`description`、`status`、`url`、`category`、`index`
  - 接口`DeletePremiumHost`新增响应参数 `region`

# 0.0.85 2022-04-21

### HuaweiCloud SDK AS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateScalingGroup`新增请求参数 `iam_agency_name`
  - 接口`ListScalingGroups`新增响应参数 `iam_agency_name`
  - 接口`UpdateScalingGroup`新增请求参数 `iam_agency_name`
  - 接口`ShowScalingGroup`新增响应参数 `iam_agency_name`

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListConsumeSubCustomers`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerBillsMonthlyBreakDown`新增响应参数 `sub_service_type_code`、`sub_service_type_name`、`sub_resource_type_code`、`sub_resource_type_name`、`sub_resource_id`、`sub_resource_name`

### HuaweiCloud SDK CloudDeploy

- _新增特性_
  - 支持接口`ListDeployTasks`、`ListDeployTaskHistoryByDate`、`ShowProjectSuccessRate`、`ListTaskSuccessRate`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowInstance`新增响应参数 `domain_name_info`

### HuaweiCloud SDK DWS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClusters`新增响应参数 `tags`
  - 接口`ListSnapshotDetails`移除响应参数 `cluster_id`、`size`、`name`、`description`、`finished`、`started`、`id`、`type`、`status`

### HuaweiCloud SDK IES

- _新增特性_
  - 支持智能边缘小站服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RestoreToExistingInstance`移除请求参数`restore_all_database`

# 0.0.84 2022-04-14

### HuaweiCloud SDK Core

- _新增特性_
  - 支持从实例元数据获取临时认证信息
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateDomainFullConfig`新增请求参数 `sources`、`origin_protocol`、`force_redirect`、`compress`
  - 接口`ShowDomainFullConfig`:
    - 新增响应参数 `sources`、`origin_protocol`、`force_redirect`、`compress`
    - 响应参数`certificate_source`类型调整 `string` -> `int32`

### HuaweiCloud SDK CloudBuild

- _新增特性_
  - 支持接口`ShowJobSuccessRatio`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持接口`ShowTasksRulesets`、`CheckRulesetParameters`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CSMS

- _新增特性_
  - 支持云凭据管理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateInstance`新增请求参数 `rename_commands`

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持接口`BatchSetSmn`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowQuotas`请求参数`X-Language`改为非必填
  - 接口`BatchCreateJobs`请求参数`X-Language`改为非必填
  - 接口`BatchValidateConnections`请求参数`X-Language`改为非必填
  - 接口`BatchValidateClustersConnections`请求参数`X-Language`改为非必填
  - 接口`BatchSetObjects`请求参数`X-Language`改为非必填
  - 接口`BatchCheckJobs`请求参数`X-Language`改为非必填
  - 接口`BatchCheckResults`请求参数`X-Language`改为非必填
  - 接口`BatchSetSpeed`:
    - 新增请求参数 `is_utc`
    - 请求参数`X-Language`改为非必填
  - 接口`BatchShowParams`请求参数`X-Language`改为非必填
  - 接口`UpdateParams`请求参数`X-Language`改为非必填
  - 接口`BatchStartJobs`请求参数`X-Language`改为非必填
  - 接口`BatchRestoreTask`请求参数`X-Language`改为非必填
  - 接口`BatchStopJobs`请求参数`X-Language`改为非必填
  - 接口`BatchDeleteJobs`请求参数`X-Language`改为非必填
  - 接口`BatchUpdateJob`:
    - 请求参数`X-Language`改为非必填
    - 请求参数`endpoints`、`protocol`改为非必填
  - 接口`BatchResetPassword`请求参数`X-Language`改为非必填
  - 接口`BatchSetDefiner`请求参数`X-Language`改为非必填
  - 接口`CreateCompareTask`请求参数`X-Language`改为非必填
  - 接口`ListCompareResult`请求参数`X-Language`改为非必填
  - 接口`BatchListProgresses`请求参数`X-Language`改为非必填
  - 接口`BatchListJobDetails`:
    - 新增响应参数 `is_utc`
    - 请求参数`X-Language`改为非必填
    - 响应参数`alarm_notify`类型调整 `string` -> `object`
  - 接口`ShowJobList`:
    - 请求参数`X-Language`改为非必填
    - 响应参数`billing_tag`类型调整 `string` -> `boolean`
    - 响应参数`node_newFramework`类型调整 `string` -> `boolean`
  - 接口`BatchListJobStatus`请求参数`X-Language`改为非必填
  - 接口`BatchChangeData`请求参数`X-Language`改为非必填
  - 接口`BatchSwitchover`请求参数`X-Language`改为非必填
  - 接口`BatchListRposAndRtos`请求参数`X-Language`改为非必填
  - 接口`ShowMonitoringData`请求参数`X-Language`改为非必填
  - 接口`BatchListStructProcess`请求参数`X-Language`改为非必填
  - 接口`BatchListStructDetail`请求参数`X-Language`改为非必填
  - 接口`BatchUpdateUser`请求参数`X-Language`改为非必填
  - 接口`ListUsers`请求参数`X-Language`改为非必填
  - 接口`BatchSetPolicy`:
    - 请求参数`export_snapshot`类型调整 `string` -> `boolean`
    - 请求参数`X-Language`改为非必填
    - 请求参数`conflict_policy`、`filter_ddl_policy`改为非必填

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListServersDetails`新增请求参数 `ip_eq`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateFunctionConfig`新增请求参数 `encrypted_user_data`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持接口`RestartInstance`、`ShowInstanceConfiguration`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`:
    - 移除请求参数 `solution`
    - 请求参数`sharding_num`、`coordinator_num`改为非必填
  - 接口`ListConfigurations`新增响应参数 `count`
  - 接口`ListFlavors`新增响应参数 `total`

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSimPricePlans`:
    - 响应参数`create_time`类型调整 `date` -> `date-time`
    - 响应参数`active_time`类型调整 `date` -> `date-time`
    - 响应参数`stop_time`类型调整 `date` -> `date-time`

### HuaweiCloud SDK IEC

- _新增特性_
  - 支持以下接口：
    - `ListRelatedRoutetables`
    - `ListRoutetables`
    - `CreateRoutetable`
    - `ShowRoutetable`
    - `UpdateRoutetable`
    - `DeleteRoutetable`
    - `AssociateSubnet`
    - `DisassociateSubnet`
    - `ListRoutes`
    - `CreateRoutes`
    - `UpdateRoutes`
    - `DeleteRoutes`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSecurityGroups`新增响应参数 `action`、`priority`
  - 接口`CreateSecurityGroup`新增响应参数 `action`、`priority`
  - 接口`ShowSecurityGroup`新增响应参数 `action`、`priority`
  - 接口`ListSecurityGroupRules`新增响应参数 `action`、`priority`
  - 接口`CreateSecurityGroupRule`:
    - 新增请求参数 `action`、`priority`
    - 新增响应参数 `action`、`priority`
  - 接口`ShowSecurityGroupRule`新增响应参数 `action`、`priority`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateAomMappingRules`:
    - 请求参数`target_log_group_id`、`target_log_group_name`、`target_log_stream_id`、`target_log_stream_name`改为必填
    - 请求参数`container_name`改为非必填
    - 响应参数`target_log_group_id`、`target_log_group_name`、`target_log_stream_id`、`target_log_stream_name`改为必填
    - 响应参数`container_name`改为非必填
  - 接口`ShowAomMappingRules`:
    - 响应参数`target_log_group_id`、`target_log_group_name`、`target_log_stream_id`、`target_log_stream_name`改为必填
    - 响应参数`container_name`改为非必填
  - 接口`CreateAomMappingRules`:
    - 请求参数`target_log_group_id`、`target_log_group_name`、`target_log_stream_id`、`target_log_stream_name`改为必填
    - 请求参数`container_name`改为非必填
    - 响应参数`target_log_group_id`、`target_log_group_name`、`target_log_stream_id`、`target_log_stream_name`改为必填
    - 响应参数`container_name`改为非必填
  - 接口`ShowAomMappingRule`:
    - 响应参数`target_log_group_id`、`target_log_group_name`、`target_log_stream_id`、`target_log_stream_name`改为必填
    - 响应参数`container_name`改为非必填
  - 接口`ListHost`请求参数`host_id_list`改为必填
  - 接口`ListHostGroup`请求参数`host_group_id_list`改为必填
  - 接口`ListAccessConfig`请求参数`access_config_name_list`、`host_group_name_list`、`log_group_name_list`、`log_stream_name_list`改为必填

### HuaweiCloud SDK Moderation

- _新增特性_
  - 支持接口`RunModerationAudio`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCertificates`:
    - 新增请求参数 `enterprise_project_id`、`deploy_support`
    - 新增响应参数 `enterprise_project_id`
    - 响应参数`id`、`name`、`domain`、`sans`、`signature_algorithm`、`deploy_support`、`type`、`brand`、`expire_time`、`domain_type`、`validity_period`、`status`、`domain_count`、`wildcard_count`、`description`、`total_count`改为必填
  - 接口`ImportCertificate`:
    - 新增请求参数 `enterprise_project_id`
    - 响应参数`certificate_id`改为必填
  - 接口`ShowCertificate`:
    - 新增响应参数 `enterprise_project_id`
    - 响应参数`id`、`status`、`order_id`、`name`、`type`、`brand`、`push_support`、`revoke_reason`、`signature_algorithm`、`issue_time`、`not_before`、`not_after`、`validity_period`、`validation_method`、`domain_type`、`domain`、`sans`、`domain_count`、`wildcard_count`改为必填
  - 接口`ExportCertificate`响应参数`certificate`、`certificate_chain`、`private_key`改为必填

### HuaweiCloud SDK SMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTasks`新增响应参数 `process_trace`
  - 接口`ShowTask`新增响应参数 `process_trace`
  - 接口`UpdateTask`新增请求参数 `process_trace`
  - 接口`UpdateTaskSpeed`:
    - 新增请求参数 `process_trace`、`compress_rate`
    - 请求参数`migrate_speed`改为非必填

# 0.0.83 2022-04-07

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListStoredValueCards`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListSubCustomerDiscounts`、`BatchSetSubCustomerDiscount`
  - 接口`ShowRefundOrderDetails`新增响应参数 `resource_type_name`、`service_type_name`
  - 接口`ListCustomerOrders`新增响应参数 `service_type_name`
  - 接口`ShowCustomerOrderDetails`新增响应参数 `service_type_name`、`service_type_name`
  - 接口`ListPayPerUseCustomerResources`新增响应参数 `resource_type_name`、`service_type_name`
  - 接口`ListCustomerOnDemandResources`新增响应参数 `service_type_name`、`resource_type_name`
  - 接口`ListSubcustomerMonthlyBills`新增响应参数 `cloud_service_type_name`、`resource_type_name`
  - 接口`ListCustomerselfResourceRecordDetails`新增响应参数 `cloud_service_type_name`、`resource_type_name`、`period_type`
  - 接口`ListCustomerselfResourceRecords`新增响应参数 `cloud_service_type_name`、`resource_type_name`
  - 接口`ShowCustomerMonthlySum`新增响应参数 `service_type_name`、`resource_type_name`
  - 接口`ListCustomerBillsFeeRecords`:
    - 新增请求参数 `bill_date_begin`、`bill_date_end`
    - 新增响应参数 `service_type_name`、`resource_type_name`
  - 接口`ListUsageTypes`新增响应参数 `resource_type_name`、`service_type_name`
  - 接口`ListSubCustomerBillDetail`新增响应参数 `service_type_name`、`resource_type_name`
  - 接口`ListCustomerBillsMonthlyBreakDown`新增响应参数 `service_type_name`、`resource_type_name`
  - 接口`ListFreeResourceInfos`新增响应参数 `service_type_name`
  - 接口`ListIncentiveDiscountPolicies`新增响应参数 `service_type_name`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateNodePool`移除请求参数 `kind`、`apiVersion`、`status`

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持以下接口：
    - `CreateDocWatermarkByAddress`
    - `ShowDocWatermarkByAddress`
    - `ShowImageWatermarkWithImage`
    - `CreateImageWatermarkByAddress`
    - `ShowImageWatermarkByAddress`
    - `ShowImageWatermarkWithImageByAddress`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateImageWatermark`:
    - 新增请求参数 `image_watermark`
    - 请求参数`blind_watermark`改为必填

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`StopSimCard`移除请求参数 `price_plan_list`
  - 接口`ResetSimCard`移除请求参数 `price_plan_list`

### HuaweiCloud SDK Meeting

- _新增特性_
  - 支持接口`StartMeeting`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeWaybillElectronic`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeVatInvoice`新增响应参数 `print_code`
  - 接口`RecognizeVehicleLicense`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`
  - 接口`RecognizeTaxiInvoice`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`
  - 接口`RecognizeDriverLicense`新增响应参数 `type`、`accumulated_scores`、`status`、`generation_date`、`current_time`
  - 接口`RecognizeTrainTicket`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`
  - 接口`RecognizeBankcard`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ApplyConfigurationAsync`、`UpdateInstanceConfigurationAsync`、`DeleteSqlserverDatabaseEx`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持接口`ListConfigurations`、`ListDatastores`、`ListFlavors`、`ListStorageTypes`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`:
    - 新增请求参数 `solution`
    - 请求参数`mode`新增枚举值`centralization_standard`
    - 请求参数`type`新增枚举值`ESSD`

# 0.0.82 2022-03-31

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复并发读写认证信息缓存的问题
- _特性变更_
  - 无

### HuaweiCloud SDK CC

- _新增特性_
  - 支持云连接服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`BatchShowNodesInformation`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.81 2022-03-25

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DeleteserviceDiscoveryRules`新增响应参数 `responseStatus`
  - 接口`AddOrUpdateServiceDiscoveryRules`新增响应参数 `responseStatus`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持V2版本接口：
    - `ShowDomainLocationStats`
    - `ShowDomainStats`
    - `ShowTopUrl`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DAS

- _新增特性_
  - 支持接口 `ShowSqlExplain`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRedislog`新增响应参数 `group_name`、`replication_ip`

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPublicZones`移除响应参数 `routers`
  - 接口`ShowRecordSetByZone`新增请求参数 `marker`、`limit`、`offset`、`line_id`、`tags`、`status`、`type`、`name`、`id`、`sort_key`、`sort_dir`、`search_mode`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持以下接口：
    - `ListWorkflows`
    - `CreateWorkflow`
    - `BatchDeleteWorkflows`
    - `ListWorkflowExecutions`
    - `StartWorkflowExecution`
    - `ShowWorkflowExecution`
    - `ShowWorkFlow`
    - `UpdateWorkFlow`
    - `ShowTenantMetric`
    - `ShowWorkFlowMetric`
    - `RetryWorkFlow`
    - `StopWorkFlow`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSimCards`:
    - 新增请求参数 `min_flow`、`max_flow`、`order_id`、`filter_downtime_period`
    - 响应参数`device_status_date`类型调整 `date` -> `date-time`
    - 响应参数`expire_time`类型调整 `date` -> `date-time`
  - 接口`StopSimCard`新增请求参数 `price_plan_list`
  - 接口`ResetSimCard`新增请求参数 `price_plan_list`
  - 接口`ShowSimCard`:
    - 响应参数`device_status_date`类型调整 `date` -> `date-time`
    - 响应参数`expire_time`类型调整 `date` -> `date-time`

### HuaweiCloud SDK IMS

- _新增特性_
  - 支持接口`ListVersions`、`ShowVersion`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDataImage`请求参数`os_type`改为非必填

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 支持接口 `ResetFingerprint`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeVatInvoice`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`
  - 接口`RecognizeIdCard`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`
  - 接口`RecognizeDriverLicense`:
    - 新增请求参数 `return_text_location`
    - 新增响应参数 `text_location`

### HuaweiCloud SDK VSS

- _新增特性_
  - 支持以下接口：
    - `ShowDomainSettings`
    - `UpdateDomainSettings`
    - `ListTaskHistories`
    - `ListPortResults`
    - `ListBusinessRisks`
    - `UpdateFalsePositive`
    - `CancelTasks`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDomains`新增请求参数 `domain_id`
  - 接口`ShowResults`新增响应参数 `hit_details`

# 0.0.80 2022-03-10

### HuaweiCloud SDK BCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateChannels`新增响应参数 `operation_id`
  - 接口`CreateNewBlockchain`请求参数`fabric_version`改为非必填

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DeleteAddonInstance`请求参数`cluster_id`改为非必填

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTopUrl`移除请求参数 `X-Auth-Token`
  - 接口`ShowDomainLocationStats`移除请求参数 `X-Auth-Token`
  - 接口`ShowDomainItemDetails`移除请求参数 `X-Auth-Token`
  - 接口`ShowDomainStats`移除请求参数 `X-Auth-Token`
  - 接口`ShowDomainItemLocationDetails`移除请求参数 `X-Auth-Token`

### HuaweiCloud SDK DWS

- _新增特性_
  - 支持以下接口：
    - `ListClusterDetails`
    - `DeleteCluster`
    - `ResetPassword`
    - `ListSnapshots`
    - `CreateSnapshot`
    - `RestartCluster`
    - `ListClusters`
    - `CreateCluster`
    - `RestoreCluster`
    - `ResizeCluster`
    - `ListNodeTypes`
    - `ListSnapshotDetails`
    - `DeleteSnapshot`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ELB

- _新增特性_
  - 支持以下接口：
    - `ListLogtanks`
    - `CreateLogtank`
    - `ShowLogtank`
    - `UpdateLogtank`
    - `DeleteLogtank`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`新增响应参数 `https_cps`
  - 接口`ShowFlavor`新增响应参数 `https_cps`
  - 接口`ListLoadBalancers`请求参数`X-Auth-Token`改为非必填
  - 接口`CreateLoadBalancer`请求参数`X-Auth-Token`改为非必填
  - 接口`UpdateIpList`请求参数`X-Auth-Token`改为非必填
  - 接口`BatchDeleteIpList`请求参数`X-Auth-Token`改为非必填

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSimCards`响应参数`act_date`类型调整 `date` -> `date-time`
  - 接口`ShowSimCard`响应参数`act_date`类型调整 `date` -> `date-time`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListLogs`新增响应参数 `count`

### HuaweiCloud SDK Meeting

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddDepartment`新增请求参数 `sortLevel`
  - 接口`UpdateDepartment`新增请求参数 `sortLevel`
  - 接口`ShowDeptAndChildDept`新增响应参数 `sortLevel`、`sortLevel`
  - 接口`SearchDepartmentByName`新增响应参数 `sortLevel`

# 0.0.79 2022-03-07

### HuaweiCloud SDK Core

- _新增特性_
  - 支持SK衍生认证方式
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`UpdateClusterEip`、`ShowClusterEndpoints`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CES

- _新增特性_
  - 支持以下接口 (V2)：
    - `ListAlarms`
    - `CreateAlarm`
    - `DeleteAlarm`
    - `UpdateAlarmAction`
    - `ListAlarmResources`
    - `DeleteAlarmResources`
    - `AddAlarmResources`
    - `AddResourceGroupsResourcesBatch`
    - `DeleteResourceGroupsResourcesBatch`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateJobs`请求参数`engine_type`新增枚举值`gaussdbv5`、`postgresql`
  - 接口`BatchValidateConnections`:
    - 新增请求参数 `kafka_security_config`
    - 请求参数`db_type`新增枚举值`postgresql`
  - 接口`BatchUpdateUser`:
    - 新增请求参数 `is_sync_object_privilege`
    - 新增响应参数 `no_privileges`、`parent_account`、`no_parent_account`
  - 接口`ListUsers`新增响应参数 `no_privileges`、`parent_account`、`no_parent_account`
  - 接口`BatchSetPolicy`新增请求参数 `topic_policy`、`topic`、`partition_policy`、`kafka_data_format`、`topic_name_format`、`partitions_num`、`replication_factor`、`is_fill_materialized_view`、`export_snapshot`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePrePaidPublicip`的请求参数`ip_version`类型变更： `integer` -> `enum`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShrinkInstanceNode`新增响应参数 `order_id`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEditingJob`新增响应参数 `error_code`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DownloadSlowlog`移除请求参数 `request_id`

# 0.0.78 2022-02-25

### HuaweiCloud SDK AS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAllScalingV2Policies`新增请求参数 `alarm_id`
  - 接口`CreateScalingConfig`请求参数`volume_type`新增枚举值`GPSSD`
  - 接口`ShowResourceQuota`新增响应参数 `min`
  - 接口`ShowPolicyAndInstanceQuota`新增响应参数 `min`

### HuaweiCloud SDK BMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateBaremetalServerMetadata`:
    - 请求体类型修改 `MetaData` -> `UpdateBaremetalServerMetadataReq`
    - 移除响应参数 `key`

### HuaweiCloud SDK CDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJobs`响应参数`from-connector-name`、`to-link-name`改为必填
  - 接口`UpdateJob`请求参数`from-connector-name`、`to-link-name`改为必填
  - 接口`CreateAndStartRandomClusterJob`:
    - 请求参数`from-connector-name`、`to-link-name`改为必填
    - 响应参数`progress`类型调整 `int32` -> `float`
    - 响应参数`isStopingIncrement`类型调整 `boolean` -> `string`
  - 接口`StopJob`新增响应参数 `submissions`
  - 接口`CreateJob`请求参数`from-connector-name`、`to-link-name`改为必填
  - 接口`StartJob`:
    - 响应参数`progress`类型调整 `int32` -> `float`
    - 响应参数`isStopingIncrement`类型调整 `boolean` -> `string`
  - 接口`ShowJobStatus`响应参数`progress`类型调整 `double` -> `float`
  - 接口`ShowSubmissions`响应参数`progress`类型调整 `double` -> `float`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ShowDomainLocationStats`、`ShowDomainFullConfig`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowDomainStats`:
    - 新增请求参数 `service_area`
    - 移除请求参数 `X-Auth-Token`、`country`、`district`、`isp`
    - 移除响应参数 `start_time`、`end_time`、`interval`、`action`、`stat_type`、`group_by`
  - 接口`UpdateDomainFullConfig`新增请求参数 `https`

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数 `instance_domain_id`、`instance_user_id`
  - 接口`CreateInstanceBy3rd`请求参数`instance_user_domain_name`、`instance_user_name`改为非必填

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持接口`CheckRecord`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRtcClientQosDetails`请求参数`mid`类型调整 `array` -> `string`

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 支持接口`ShowStatisticCommitV3`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListProjectSets`:
    - 新增响应参数 `CreateTime`、`UpdateTime`、`external_params`、`variables_no_file`
    - 移除响应参数 `create_time`、`update_time`、`group`
  - 接口`UpdateProject`请求参数`name`改为必填
  - 接口`ShowTask`:
    - 新增响应参数 `parallel`、`contents`、`sort`、`related_temp_running_data`
    - 移除响应参数 `content`
  - 接口`UpdateTask`:
    - 新增请求参数 `contents`、`sort`、`related_temp_running_data`
    - 新增响应参数 `parallel`、`contents`、`sort`、`related_temp_running_data`
    - 移除请求参数 `content`
    - 移除响应参数 `content`
    - 请求参数`name`改为必填
  - 接口`ShowReport`:
    - 新增响应参数 `performance`、`minNetworkTraffic`、`avgNetworkTraffic`、`maxNetworkTraffic`、`branchId`、`branchName`、`projectId`、`serviceId`
    - 移除响应参数 `progressState`、`statusValue`
    - 响应参数`averageRespTime`类型调整 `float` -> `double`
    - 响应参数`avgRecBytes`类型调整 `float` -> `double`
    - 响应参数`avgSentBytes`类型调整 `int32` -> `double`
    - 响应参数`avgTranRespTime`类型调整 `string` -> `double`
    - 响应参数`currentThreadNum`类型调整 `int32` -> `double`
    - 响应参数`errorCount`类型调整 `int32` -> `double`
    - 响应参数`errorEventsCount`类型调整 `int32` -> `double`
    - 响应参数`failedAssert`类型调整 `int32` -> `double`
    - 响应参数`failedOthers`类型调整 `int32` -> `double`
    - 响应参数`failedParsed`类型调整 `int32` -> `double`
    - 响应参数`failedRefused`类型调整 `int32` -> `double`
    - 响应参数`failedTimeout`类型调整 `int32` -> `double`
    - 响应参数`max`类型调整 `int32` -> `double`
    - 响应参数`maxRecBytes`类型调整 `int32` -> `double`
    - 响应参数`maxRespTime`类型调整 `int32` -> `double`
    - 响应参数`maxSentBytes`类型调整 `int32` -> `double`
    - 响应参数`maxTranRespTime`类型调整 `int32` -> `double`
    - 响应参数`min`类型调整 `int32` -> `double`
    - 响应参数`requests`类型调整 `int32` -> `double`
    - 响应参数`result`类型调整 `int32` -> `double`
    - 响应参数`status`类型调整 `int32` -> `double`
    - 响应参数`successCount`类型调整 `int32` -> `double`
    - 响应参数`successRate`类型调整 `int32` -> `double`
    - 响应参数`sum1xx`类型调整 `int32` -> `double`
    - 响应参数`sum2xx`类型调整 `int32` -> `double`
    - 响应参数`sum3xx`类型调整 `int32` -> `double`
    - 响应参数`sum4xx`类型调整 `int32` -> `double`
    - 响应参数`sum5xx`类型调整 `int32` -> `double`
    - 响应参数`taskStatus`类型调整 `int32` -> `double`
    - 响应参数`tp50`类型调整 `int32` -> `double`
    - 响应参数`tp75`类型调整 `int32` -> `double`
    - 响应参数`tp90`类型调整 `int32` -> `double`
    - 响应参数`tp95`类型调整 `int32` -> `double`
    - 响应参数`tp99`类型调整 `int32` -> `double`
    - 响应参数`tps`类型调整 `float` -> `double`
    - 响应参数`tranTPS`类型调整 `string` -> `double`
    - 响应参数`transactionSuccess`类型调整 `string` -> `double`
    - 响应参数`transactionalSuccessRate`类型调整 `int32` -> `double`
    - 响应参数`transactionalTps`类型调整 `int32` -> `double`
    - 响应参数`transactionalTpsSuccess`类型调整 `int32` -> `double`
    - 响应参数`transactions`类型调整 `string` -> `double`
    - 响应参数`vum`类型调整 `int32` -> `double`
    - 响应参数`avgResponseTime`类型调整 `float` -> `double`
    - 响应参数`caseRetry`类型调整 `int32` -> `double`
    - 响应参数`completeNum`类型调整 `int32` -> `double`
    - 响应参数`duration`类型调整 `int32` -> `double`
    - 响应参数`executedNum`类型调整 `int32` -> `double`
    - 响应参数`kpiCaseCount`类型调整 `int32` -> `double`
    - 响应参数`kpiCaseExecuteCount`类型调整 `int32` -> `double`
    - 响应参数`kpiCasePassCount`类型调整 `int32` -> `double`
    - 响应参数`maxUsers`类型调整 `int32` -> `double`
    - 响应参数`passNum`类型调整 `int32` -> `double`
    - 响应参数`stage`类型调整 `int32` -> `double`
    - 响应参数`totalNum`类型调整 `int32` -> `double`
  - 接口`UpdateCase`:
    - 新增请求参数 `contents`、`sort`
    - 移除请求参数 `content`
  - 接口`CreateTemp`新增请求参数 `contents`
  - 接口`UpdateTemp`:
    - 请求参数`bodys`类型调整 `array` -> `string`
    - 请求参数`name`改为必填
  - 接口`CreateVariable`新增请求参数 `is_quoted`
  - 接口`ShowHistoryRunInfo`:
    - 响应参数`run_id`类型调整 `int32` -> `double`
    - 响应参数`run_type`类型调整 `int32` -> `double`
    - 响应参数`continue_time`类型调整 `int32` -> `double`

### HuaweiCloud SDK CSS

- _新增特性_
  - 支持以下接口：
    - `UpdateFlavor`
    - `UpdateFlavorByType`
    - `UpdateShrinkNodes`
    - `UpdateShrinkCluster`
    - `ListLogsJob`
    - `ShowClusterDetail`
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateExtendCluster`移除响应参数 `id`、`instances`
  - 接口`StartConnectivityTest`移除请求参数 `status`

### HuaweiCloud SDK DDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ExpandInstanceNodes`新增请求参数 `group_id`
  - 接口`ShrinkInstanceNodes`新增请求参数 `group_id`
  - 接口`CreateDatabase`请求参数`shard_unit`改为非必填

### HuaweiCloud SDK DGC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListConnections`:
    - 新增响应参数 `type`、`description`
    - 移除响应参数 `connectionType`
    - 响应参数`total`类型调整 `string` -> `int32`
    - 响应参数`name`改为必填
  - 接口`CreateConnection`:
    - 新增请求参数 `type`、`description`
    - 移除请求参数 `connectionType`
    - 请求参数`name`改为必填
  - 接口`ShowConnection`:
    - 新增响应参数 `type`、`description`
    - 移除响应参数 `connectionType`
    - 响应参数`name`改为必填
  - 接口`UpdateConnection`:
    - 新增请求参数 `type`、`description`
    - 移除请求参数 `connectionType`
    - 请求参数`name`改为必填
  - 接口`ExecuteScript`:
    - 新增响应参数 `instanceId`
    - 移除响应参数 `jobId`
    - 请求参数`params`类型调整 `string` -> `object`

### HuaweiCloud SDK ELB

- _新增特性_
  - 支持接口`BatchCreateMembers`、`BatchDeleteMembers`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFunctions`响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
  - 接口`CreateFunction`:
    - 请求参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
    - 响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
  - 接口`ShowFunctionCode`响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
  - 接口`UpdateFunctionCode`响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
  - 接口`ShowFunctionConfig`响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
  - 接口`UpdateFunctionConfig`:
    - 请求参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
    - 响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
  - 接口`ListFunctionVersions`响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
  - 接口`CreateFunctionVersion`响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`
  - 接口`CreateDependency`请求参数`runtime`新增枚举值`Java11`、`Node.js14.18`、`Python3.9`
  - 接口`UpdateDependency`请求参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`PHP 7.3`
  - 接口`ImportFunction`响应参数`runtime`新增枚举值`Java8`、`Java11`、`Node.js6.10`、`Node.js8.10`、`Node.js10.16`、`Node.js12.13`、`Node.js14.18`、`Python2.7`、`Python3.6`、`Python3.9`、`Go1.8`、`Go1.x`、`PHP7.3`, 移除枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持接口`UpdateAuditLog`、`ShowAuditLog`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`ListSingleStreamDetail`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持接口`UpdateStructConfig`、`CreateStructConfig`、`ListStructTemplate`、`ListBreifStructTemplate`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListLogGroups`新增响应参数 `tag`
  - 接口`ListLogStream`新增响应参数 `tag`

### HuaweiCloud SDK Meeting

- _新增特性_
  - 支持华为云会议服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`BatchUpdateChildNickNames`、`ListIterationHistories`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListProjectIterationsV4`新增请求参数 `updated_time_interval`、`include_deleted`
  - 接口`ListIssuesV4`新增请求参数 `include_deleted`、`updated_time_interval`
  - 接口`ShowIssueV4`新增响应参数 `description`、`order`、`accessories`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListSlowLogFile`、`StopInstance`、`StartupInstance`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCertificates`新增响应参数 `sans`、`signature_algorithm`、`deploy_support`
  - 接口`ImportCertificate`新增请求参数 `enc_certificate`、`enc_private_key`
  - 接口`ShowCertificate`:
    - 新增响应参数 `signature_algorithm`
    - 移除响应参数 `signature_algrithm`
  - 接口`ExportCertificate`新增响应参数 `enc_certificate`、`enc_private_key`

### HuaweiCloud SDK VOD

- _新增特性_
  - 支持接口`ListDomainLogs`
- _解决问题_
  - 无
- _特性变更_
  - 接口`DeleteAssets`新增请求参数 `delete_type`

### HuaweiCloud SDK VPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`NeutronListSubnets`新增响应参数`subnetpool_id`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListWhiteblackipRule`:
    - 新增响应参数 `addr`
    - 移除响应参数 `ip`

# 0.0.77 2022-02-10

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateAlarmRule`:
    - 请求参数`statistic`类型调整 `string` -> `enum`
    - 请求参数`alarm_level`、`comparison_operator`、`evaluation_periods`、`metric_name`、`namespace`、`period`、`statistic`、`threshold`、`unit`改为非必填

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRateOnPeriodDetail`:
    - 新增请求参数 `fee_installment_mode`
    - 新增响应参数 `installment_official_website_amount`、`installment_period_type`、`installment_official_discount_amount`、`installment_amount`

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListProtectable`:
    - 响应参数`result`类型调整 `string` -> `boolean`
    - 响应参数`size`类型调整 `string` -> `int32`
  - 接口`ShowProtectable`:
    - 响应参数`result`类型调整 `string` -> `boolean`
    - 响应参数`size`类型调整 `string` -> `int32`

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`ShowVersion`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateAddonInstance`移除响应参数 `kind`、`apiVersion`、`metadata`、`spec`、`status`
  - 接口`ListNodes`新增响应参数 `isStatic`、`privateIPv6IP`
  - 接口`CreateNode`新增请求参数 `isStatic`
  - 接口`DeleteNode`新增响应参数 `isStatic`、`privateIPv6IP`
  - 接口`ShowNode`新增响应参数 `isStatic`、`privateIPv6IP`
  - 接口`UpdateNode`新增响应参数 `isStatic`、`privateIPv6IP`
  - 接口`RemoveNode`:
    - 请求参数`uid`改为必填
    - 响应参数`uid`改为必填
  - 接口`MigrateNode`:
    - 请求参数`uid`改为必填
    - 响应参数`uid`改为必填
  - 接口`ListNodePools`新增响应参数 `isStatic`
  - 接口`CreateNodePool`新增请求参数 `isStatic`
  - 接口`DeleteNodePool`新增响应参数 `isStatic`
  - 接口`ShowNodePool`新增响应参数 `isStatic`
  - 接口`UpdateNodePool`:
    - 新增请求参数 `isStatic`
    - 新增响应参数 `isStatic`

### HuaweiCloud SDK CSS

- _新增特性_
  - 支持以下接口：
    - `UpdateOndemandClusterToPeriod`
    - `UpdateClusterName`
    - `ResetPassword`
    - `StartKibanaPublic`
    - `UpdateCloseKibana`
    - `UpdateAlterKibana`
    - `UpdatePublicKibanaWhitelist`
    - `StopPublicKibanaWhitelist`
    - `CreateCnf`
    - `UpdateCnf`
    - `StartPipeline`
    - `StopPipeline`
    - `AddFavorite`
    - `StartConnectivityTest`
    - `ListTemplates`
    - `ListConfs`
    - `ListPipelines`
    - `ListActions`
    - `ShowGetConfDetail`
    - `DeleteConf`
    - `DeleteTemplate`
    - `StartLogs`
    - `StopLogs`
    - `ShowGetLogSetting`
    - `UpdateLogSetting`
    - `StartLogAutoBackupPolicy`
    - `StopLogAutoBackupPolicy`
    - `CreateLogBackup`
    - `ShowLogBackup`
    - `CreateBindPublic`
    - `UpdateUnbindPublic`
    - `UpdatePublicBandWidth`
    - `StartPublicWhitelist`
    - `StopPublicWhitelist`
    - `StartVpecp`
    - `StopVpecp`
    - `ShowVpcepConnection`
    - `UpdateVpcepConnection`
    - `UpdateVpcepWhitelist`
    - `UpdateYmls`
    - `ListYmlsJob`
    - `ListYmls`
    - `ListClustersDetails`
    - `CreateCluster`
    - `DeleteCluster`
    - `RestartCluster`
    - `UpdateExtendCluster`
    - `UpdateExtendInstanceStorage`
    - `ListFlavors`
    - `ShowClusterTag`
    - `CreateClustersTags`
    - `ListClustersTags`
    - `UpdateBatchClustersTags`
    - `DeleteClustersTags`
    - `ShowIkThesaurus`
    - `CreateLoadIkThesaurus`
    - `DeleteIkThesaurus`
    - `StartAutoSetting`
    - `UpdateSnapshotSetting`
    - `ShowAutoCreatePolicy`
    - `CreateAutoCreatePolicy`
    - `CreateSnapshot`
    - `ListSnapshots`
    - `StopSnapshot`
    - `RestoreSnapshot`
    - `DeleteSnapshot`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 支持接口`CreateOnlineMigrationTask`、`SetOnlineMigrationTask`、`BatchStopMigrationTasks`、`StopMigrationTaskSync`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DevStar

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowApplicationDependentResources`:
    - 新增请求参数 `X-Auth-Token`、`limit`、`offset`
    - 新增响应参数 `count`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateFunction`:
    - 请求参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
    - 响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
  - 接口`ListFunctions`响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
  - 接口`UpdateFunctionCode`响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
  - 接口`ShowFunctionCode`响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
  - 接口`UpdateFunctionConfig`:
    - 请求参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
    - 响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
  - 接口`ShowFunctionConfig`响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
  - 接口`CreateFunctionVersion`响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
  - 接口`ListFunctionVersions`响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`
  - 接口`CreateDependency`请求参数`runtime`新增枚举值`Go1.x`
  - 接口`UpdateDependency`请求参数`runtime`新增枚举值`Go1.x`
  - 接口`ImportFunction`响应参数`runtime`新增枚举值`Java 8`、`Node.js 6.10`、`Node.js 8.10`、`Node.js 10.16`、`Node.js 12.13`、`Python 2.7`、`Python 3.6`、`Go 1.8`、`Go 1.x`、`PHP 7.3`, 移除枚举值`Python2.7`、`Python3.6`、`Go1.8`、`Java8`、`Node.js6.10`、`Node.js8.10`、`Custom`、`PHP7.3`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowGaussMySqlInstanceInfo`新增响应参数`alias`
  - 接口`CreateGaussMySqlBackup`新增响应参数`job_id`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`响应参数`port`类型调整 `int32` -> `string`

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`ListTranscodeTaskCount`、`ListAreaDetail`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRecordData`新增请求参数 `publish_domain`
  - 接口`CreateRecordRule`移除请求参数 `plan_record_time`
  - 接口`ListRecordRules`移除响应参数 `plan_record_time`
  - 接口`ShowRecordRule`移除响应参数 `plan_record_time`
  - 接口`UpdateRecordRule`:
    - 移除请求参数 `plan_record_time`
    - 移除响应参数 `plan_record_time`
  - 接口`CreateRecordCallbackConfig`移除请求参数 `on_demand_callback_url`
  - 接口`ListRecordCallbackConfigs`移除响应参数 `on_demand_callback_url`
  - 接口`ShowRecordCallbackConfig`移除响应参数 `on_demand_callback_url`
  - 接口`UpdateRecordCallbackConfig`移除请求参数 `on_demand_callback_url`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RestoreToExistingInstance`新增请求参数`restore_all_database`
  - 接口`StartRecyclePolicy`移除请求参数`is_open_recycle_policy`

### HuaweiCloud SDK SIS

- _新增特性_
  - 支持语音交互服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.76 2022-01-25

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateApiV2`新增请求参数 `content_type`
  - 接口`ShowDetailsOfApiV2`新增响应参数 `publish_time`、`roma_app_name`、`ld_api_id`、`api_group_info`、`content_type`
  - 接口`UpdateApiV2`:
    - 新增请求参数 `content_type`
    - 新增响应参数 `publish_time`、`roma_app_name`、`ld_api_id`、`api_group_info`、`content_type`
  - 接口`ListApiRuntimeDefinitionV2`新增响应参数 `content_type`
  - 接口`ListApiVersionDetailV2`:
    - 新增响应参数 `roma_app_name`、`ld_api_id`、`api_group_info`、`content_type`
    - 响应参数`publish_time`类型调整 `date-time` -> `string`

### HuaweiCloud SDK CDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJobs`:
    - 移除响应参数 `simple`
    - 响应参数`name`、`values`改为必填
  - 接口`UpdateJob`请求参数`name`、`values`改为必填
  - 接口`CreateAndStartRandomClusterJob`请求参数`name`、`values`改为必填
  - 接口`CreateJob`请求参数`name`、`values`改为必填
  - 接口`ListClusters`移除响应参数 `config_status`

### HuaweiCloud SDK CES

- _新增特性_
  - 支持接口`ListAlarmHistories`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持接口`DeleteRuleset`、`SetDefaulTemplate`、`ShowTasklog`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTask`新增请求参数 `endpoint_id`
  - 接口`CreateRuleset`新增请求参数 `custom_attributes`
  - 接口`ListTemplateRules`:
    - 新增请求参数 `tags`
    - 新增响应参数 `rule_config_list`

### HuaweiCloud SDK DevStar

- _新增特性_
  - 支持接口`ShowRepositoryByCloudIde`、`ListTemplates`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IAM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateLoginToken`新增响应参数`session_user_id`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 支持接口`ListEngineProducts`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `dr_enable`
  - 接口`ShowInstance`新增响应参数 `dr_enable`
  - 接口`ListProducts`:
    - 新增响应参数 `Hourly`、`Monthly`
    - 移除响应参数 `hourly`、`honthly`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateSqlAlarmRule`移除请求参数 `language`
  - 接口`UpdateSqlAlarmRule`移除请求参数 `language`
  - 接口`ListSqlAlarmRules`:
    - 新增响应参数 `template_name`、`status`
    - 移除响应参数 `language`
  - 接口`CreateKeywordsAlarmRule`移除请求参数 `language`、`eps_id`、`eps_name`、`is_time_range_relative`
  - 接口`UpdateKeywordsAlarmRule`:
    - 移除请求参数 `language`、`eps_id`、`eps_name`、`is_time_range_relative`
    - 移除响应参数 `eps_id`、`eps_name`、`is_time_range_relative`
  - 接口`ListKeywordsAlarmRules`:
    - 新增响应参数 `template_name`、`status`
    - 移除响应参数 `language`、`eps_id`、`eps_name`、`is_time_range_relative`
  - 接口`ListAccessConfig`:
    - 新增请求参数 `access_config_tag_list`
    - 新增响应参数 `access_config_tag`
  - 接口`CreateAccessConfig`:
    - 新增请求参数 `access_config_tag`
    - 新增响应参数 `access_config_tag`
  - 接口`UpdateAccessConfig`:
    - 新增请求参数 `access_config_tag`
    - 新增响应参数 `access_config_tag`
  - 接口`DeleteAccessConfig`新增响应参数 `access_config_tag`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
  - 支持接口`ListEngineProducts`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.75 2022-01-17

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 修复接口`CreateNodePool`的请求体结构错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 支持接口`ListFunctionAsyncInvocations`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowGaussMySqlInstanceInfo`新增响应参数`used`
  - 接口`UpdateInstanceMonitor`新增请求参数`monitor_switch`
  - 接口`UpdateInstanceMonitor`的请求参数`period`类型修改： `string` -> `int32`
  - 接口`ShowGaussMySqlProjectQuotas`移除响应参数`mode`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListApiVersionNew`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`新增响应参数 `az_desc`

### HuaweiCloud SDK ROMA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddUserToApp`新增请求参数 `id`、`roles`
  - 接口`DownloadAssetArchive`移除响应参数 `apps`、`tasks`
  - 接口`ImportAsset`移除请求参数 `apps`、`tasks`
  - 接口`DeleteAsset`请求参数`tasks`改为必填
  - 接口`ShowMqsInstance`新增响应参数`access_user`

# 0.0.74 2022-01-10

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListOrderCouponsByOrderId`新增响应参数 `coupon_max_use_quantity`、`coupon_group`

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`ShowQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCluster`新增请求参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`ListClusters`新增响应参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`UpdateCluster`:
    - 新增请求参数 `customSan`、`containerNetwork`
    - 新增响应参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`ShowCluster`新增响应参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`DeleteCluster`新增响应参数 `customSan`、`offloadCluster`、`cidrs`、`flavor`、`faultDomain`
  - 接口`CreateNode`新增请求参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`ListNodes`新增响应参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`UpdateNode`新增响应参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`ShowNode`新增响应参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`DeleteNode`新增响应参数 `faultDomain`、`offloadNode`、`offloadNode`
  - 接口`CreateNodePool`新增请求参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
  - 接口`ListNodePools`新增响应参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
  - 接口`UpdateNodePool`:
    - 新增请求参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
    - 新增响应参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
  - 接口`ShowNodePool`新增响应参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`
  - 接口`DeleteNodePool`新增响应参数 `podSecurityGroups`、`faultDomain`、`offloadNode`、`offloadNode`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`UpdateDomainFullConfig`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`ListStacks`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListStacksByTag`

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 支持接口`ShowPlanList`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowPlans`:
    - 请求参数`limit`类型调整 `int64` -> `int32`
    - 请求参数`offset`类型调整 `int64` -> `int32`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListMigrationTask`新增响应参数 `ecs_tenant_private_ip`
  - 接口`ShowMigrationTask`新增响应参数 `ecs_tenant_private_ip`
  - 接口`StopMigrationTask`新增响应参数 `ecs_tenant_private_ip`

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`移除响应参数`lb_ip_address`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateServers`新增请求参数 `delete_on_termination`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateFunctionConfig`:
    - 新增请求参数 `is_stateful_function`
    - 新增响应参数 `is_stateful_function`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `lb_ip_address`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`请求参数`type`新增枚举值`GaussDB(for openGauss)`, 移除枚举值`GaussDB(openGauss)`
  - 接口`ListInstances`移除响应参数 `related_instance`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`GlanceListImages`响应参数`active_at`改为非必填
  - 接口`GlanceShowImage`响应参数`active_at`改为非必填
  - 接口`GlanceUpdateImage`响应参数`active_at`改为非必填

### HuaweiCloud SDK IoTAnalytics

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDatasource`请求参数`site_id`改为必填
  - 接口`UpdateDataSource`请求参数`site_id`改为必填

### HuaweiCloud SDK KPS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListKeypairs`:
    - 新增请求参数 `limit`、`marker`
    - 新增响应参数 `page_info`
  - 接口`ListFailedTask`新增请求参数 `limit`、`offset`
  - 接口`ListRunningTask`新增请求参数 `limit`、`offset`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`StartInstanceSingleToHaAction`:
    - 新增请求参数 `ad_domain_info`
    - 移除请求参数 `password`

### HuaweiCloud SDK ROMA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowMqsInstance`移除响应参数 `connect_dn`
  - 接口`ListMqsInstanceTopics`移除响应参数 `policies`
  - 接口`ShowDetailsOfAppV2`新增响应参数 `roma_app_type`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEndpointService`新增响应参数 `pool_id`
  - 接口`ListEndpointService`新增响应参数 `domain_id`
  - 接口`UpdateEndpointService`新增响应参数 `pool_id`
  - 接口`ListEndpointInfoDetails`新增响应参数 `enable_status`、`specification_name`

### HuaweiCloud SDK VSS

- _新增特性_
  - 支持接口`AuthorizeDomains`、`ShowTasks`、`CreateTasks`、`ShowResults`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDomains`响应参数`auth_status`新增枚举值`skip`

# 0.0.73 2021-12-25

### HuaweiCloud SDK CDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateLink`新增请求参数 `id`
  - 接口`ShowClusterDetail`新增响应参数 `endpointDomainName`、`isScheduleBootOff`、`failedReasons`、`components`、`createFrom`、`resourceId`、`flavorType`、`workSpaceId`、`trial`
  - 接口`UpdateJob`新增请求参数 `is_incre_job`、`flag`、`files_read`、`external_id`、`type`、`execute_start_date`、`delete_rows`、`enabled`、`bytes_written`、`id`、`is_use_sql`、`update_rows`、`group_name`、`bytes_read`、`execute_update_date`、`write_rows`、`files_writte`、`is_incrementing`、`execute_create_date`、`id`、`type`、`id`、`type`、`id`、`type`
  - 接口`ShowJobs`新增响应参数 `is_incre_job`、`flag`、`files_read`、`external_id`、`type`、`execute_start_date`、`delete_rows`、`enabled`、`bytes_written`、`id`、`is_use_sql`、`update_rows`、`group_name`、`bytes_read`、`execute_update_date`、`write_rows`、`files_writte`、`is_incrementing`、`execute_create_date`、`id`、`type`、`id`、`type`、`id`、`type`
  - 接口`CreateAndStartRandomClusterJob`:
    - 新增请求参数 `is_incre_job`、`flag`、`files_read`、`external_id`、`type`、`execute_start_date`、`delete_rows`、`enabled`、`bytes_written`、`id`、`is_use_sql`、`update_rows`、`group_name`、`bytes_read`、`execute_update_date`、`write_rows`、`files_writte`、`is_incrementing`、`execute_create_date`、`id`、`type`、`id`、`type`、`id`、`type`
    - 新增响应参数 `submissions`
    - 移除响应参数 `name`、`validation-result`
  - 接口`CreateJob`新增请求参数 `is_incre_job`、`flag`、`files_read`、`external_id`、`type`、`execute_start_date`、`delete_rows`、`enabled`、`bytes_written`、`id`、`is_use_sql`、`update_rows`、`group_name`、`bytes_read`、`execute_update_date`、`write_rows`、`files_writte`、`is_incrementing`、`execute_create_date`、`id`、`type`、`id`、`type`、`id`、`type`
  - 接口`StartJob`新增响应参数 `execute-date`
  - 接口`UpdateLink`新增请求参数 `id`
  - 接口`ShowLink`新增响应参数 `id`
  - 接口`ListClusters`:
    - 新增响应参数 `bakExpectedStartTime`、`bakKeepDay`、`createFrom`、`resourceId`、`flavorType`、`workSpaceId`、`trial`、`components`
    - 移除响应参数 `version`

### HuaweiCloud SDK CloudBuild

- _新增特性_
  - 支持接口`ShowHistoryDetails`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudTable

- _新增特性_
  - 支持表格存储服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 支持接口`ShowPlanJournals`、`ShowIssuesByPlanId`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持接口`CheckParameters`、`ListTaskParameter`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRules`新增响应参数 `rule_desc`
  - 接口`ListRulesets`新增响应参数 `is_devcloud_project_default`、`is_default_template`

### HuaweiCloud SDK CodeHub

- _新增特性_
  - 支持代码托管服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CopyInstance`新增请求参数 `backup_format`

### HuaweiCloud SDK DevStar

- _新增特性_
  - 支持以下接口：
    - `ShowApplicationV3`
    - `UpdateApplication`
    - `ShowApplicationDependentResources`
    - `DeleteApplicationV4`
    - `ShowApplicationResDeleteStatus`
    - `ListApplicationsV6`
    - `ShowDeploymentJobs`
    - `CreateDeploymentJobs`
    - `ShowPipelineLastStatusV2`
    - `ListPipelineTemplates`
    - `StartPipeline`
    - `ListProjectsV4`
    - `ShowRepositoryStatisticalDataV2`
    - `CheckRepositoryDuplicateName`
    - `ShowApplicationReleaseRepositories`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTemplateV3`新增响应参数 `dependents`
  - 接口`ListTemplatesV2`新增响应参数 `dependents`、`dependent_services`
  - 接口`ShowJobDetail`新增响应参数 `show_type`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`NovaListServerActions`移除响应参数`updated_at`

### HuaweiCloud SDK IEF

- _新增特性_
  - 支持智能边缘平台服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IoTAnalytics

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDatasource`:
    - 新增请求参数 `site_id`
    - 移除请求参数 `instance_id`
  - 接口`ShowAllDataSource`:
    - 新增响应参数 `site_id`
    - 移除响应参数 `instance_id`
  - 接口`UpdateDataSource`:
    - 新增请求参数 `site_id`
    - 新增响应参数 `site_id`
    - 移除请求参数 `instance_id`
    - 移除响应参数 `instance_id`
  - 接口`ShowDataSource`:
    - 新增响应参数 `site_id`
    - 移除响应参数 `instance_id`

### HuaweiCloud SDK Kafka

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePostPaidInstance`请求参数`partition_num`改为非必填
  - 接口`RestartManager`新增响应参数 `result`、`instance_id`
  - 接口`ListProducts`:
    - 新增响应参数 `hourly`、`honthly`
    - 移除响应参数 `Hourly`、`Monthly`

### HuaweiCloud SDK KPS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFailedTask`响应参数`task_time`类型调整 `int64` -> `string`
  - 接口`ListRunningTask`响应参数`task_time`类型调整 `int64` -> `string`

### HuaweiCloud SDK LTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCharts`新增请求参数 `offset`、`limit`
  - 接口`ListNotificationTemplates`新增请求参数 `offset`、`limit`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTranscodingTask`:
    - 新增请求参数 `auto_volume_setting`
    - 请求参数`volume`新增枚举值`original`
  - 接口`ListTranscodingTask`新增响应参数 `av_parameters`
  - 接口`CreateWatermarkTemplate`新增请求参数 `template_id`
  - 接口`ListWatermarkTemplate`新增响应参数 `template_id`
  - 接口`UpdateWatermarkTemplate`新增请求参数 `template_id`

### HuaweiCloud SDK MRS

- _新增特性_
  - 支持以下接口：
    - `CreateScalingPolicy`
    - `ShowClusterDetails`
    - `UpdateClusterScaling`
    - `ListHosts`
    - `CreateAndExecuteJob`
    - `ListExecuteJob`
    - `ShowJobExes`
    - `DeleteJobExecution`
    - `CreateCluster`
    - `ShowAgencyMapping`
    - `UpdateAgencyMapping`
    - `ShowJobExeListNew`
    - `CreateExecuteJob`
    - `ShowSingleJobExe`
    - `StopJob`
    - `ShowSqlResultWithJob`
    - `BatchDeleteJobs`
    - `ExecuteSql`
    - `ShowSqlResult`
    - `CancelSql`
    - `ShowHdfsFileList`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTasks`新增响应参数 `group_type`、`success_record_error_reason`、`skip_record_error_reason`、`save_prefix`
  - 接口`ShowTask`新增响应参数 `group_type`、`success_record_error_reason`、`skip_record_error_reason`、`save_prefix`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListApiVersion`、`ShowApiVersion`、`SearchQueryScaleComputeFlavors`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ROMA

- _新增特性_
  - 支持应用与数据集成平台服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SA

- _新增特性_
  - 支持态势感知服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.72 2021-12-17

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`ShowVersion`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowInstance`:
    - 新增响应参数 `bundle_url`、`visitor_id`、`visitor_name`、`visitor_domain_name`
    - 移除响应参数 `action_list`、`role`、`role_id`、`sub_org`
  - 接口`ListOrgInstances`:
    - 新增响应参数 `visitor_id`、`visitor_name`、`visitor_domain_name`
    - 移除响应参数 `action_list`、`role`、`role_id`、`sub_org`
  - 接口`ListInstances`:
    - 新增响应参数 `visitor_id`、`visitor_name`、`visitor_domain_name`
    - 移除响应参数 `action_list`、`role`、`role_id`、`sub_org`

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 支持接口`ListRtcAbnormalEvents`、`ListRtcAbnormalEventDimension`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEvents`:
    - 响应参数`event_count`类型调整 `string` -> `int32`
    - 响应参数`latest_occur_time`类型调整 `string` -> `int64`
  - 接口`BatchListMetricData`响应参数`variance`类型调整 `string` -> `double`
  - 接口`ListResourceGroup`响应参数`type_statistics`类型调整 `string` -> `int32`
  - 接口`ListEventDetail`:
    - 响应参数`event_users`类型调整 `string` -> `array`
    - 响应参数`event_sources`类型调整 `string` -> `array`

### HuaweiCloud SDK IoTAnalytics

- _新增特性_
  - 支持物联网数据分析服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEditingJob`新增请求参数 `input`、`edit_settings`
  - 接口`ListEditingJob`新增响应参数 `input`、`edit_settings`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持以下接口：
    - `RecognizeThailandIdcard`
    - `RecognizeMyanmarIdcard`
    - `RecognizeMyanmarDriverLicense`
    - `RecognizeChileIdCard`
    - `RecognizeThailandLicensePlate`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.71 2021-12-10

### HuaweiCloud SDK AOM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEvents`的请求参数`relation`更新枚举值为`[AND, OR, NOT]`

### HuaweiCloud SDK APM

- _新增特性_
  - 支持接口`ShowMasterAddress`、`ListAkSk`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK AS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateScalingGroup`新增请求参数 `allowed_address_pairs`
  - 接口`ListScalingGroups`新增响应参数 `allowed_address_pairs`
  - 接口`ShowScalingGroup`新增响应参数 `allowed_address_pairs`
  - 接口`UpdateScalingGroup`新增请求参数 `allowed_address_pairs`

### HuaweiCloud SDK BCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowBlockchainDetail`响应参数`node_cnt`类型调整 `int32` -> `int64`
  - 接口`ShowBlockchainFlavors`新增请求参数 `limit`、`offset`

### HuaweiCloud SDK CampusGo

- _新增特性_
  - 支持园区智能体服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CGS

- _新增特性_
  - 支持容器安全服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
  - 支持接口`ShowExtensionAuthorization`、`CreateExtensionAuthorization`、`CheckInstanceAccess`、`UpdateInstanceActivity`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRtcClientQosDetails`新增请求参数`stream_id`

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持接口`ShowTaskCmetrics`、`ListTemplateRules`、`ListTaskRuleset`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTaskDetail`新增响应参数 `is_access`、`trigger_type`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数 `readonly_domain_name`
  - 接口`ShowInstance`新增响应参数 `readonly_domain_name`

### HuaweiCloud SDK DDM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数 `time_zone`
  - 接口`RestartInstance`新增响应参数 `instanceId`、`instanceName`、`jobId`

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持接口`ShowOpenApiCalledRecords`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDocWatermark`新增请求参数 `readonly_password`
  - 接口`ShowScanJobs`:
    - 新增请求参数 `offset`
    - 移除请求参数 `page`
  - 接口`ShowScanJobResults`:
    - 新增请求参数 `offset`
    - 移除请求参数 `page`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DetectFaceByFile`移除响应参数 `landmark`、`gender`、`yaw_angle`、`roll_angle`、`pitch_angle`、`headpose`、`smile`、`skin`、`ethnic`
  - 接口`DetectFaceByUrl`移除响应参数 `landmark`、`gender`、`yaw_angle`、`roll_angle`、`pitch_angle`、`headpose`、`smile`、`skin`、`ethnic`
  - 接口`DetectFaceByBase64`移除响应参数 `landmark`、`gender`、`yaw_angle`、`roll_angle`、`pitch_angle`、`headpose`、`smile`、`skin`、`ethnic`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 支持以下接口：
    - `ListInstanceTags`
    - `ListProjectTags`
    - `BatchTagAction`
    - `ShowInstanceMonitorExtend`
    - `UpdateInstanceMonitor`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListGaussMySqlInstances`:
    - 新增请求参数 `private_ip`、`tags`
    - 新增响应参数 `tags`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`请求参数`region`改为必填

### HuaweiCloud SDK GES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CheckMetadata`

### HuaweiCloud SDK HiLens

- _新增特性_
  - 支持华为HiLens服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK KMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ImportKeyMaterial`新增请求参数 `encrypted_privatekey`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRecordRule`的请求参数`default_record_config`改为必填

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持以下接口：
    - `ListLogHistogram`
    - `ListHost`
    - `ListHostGroup`
    - `UpdateHostGroup`
    - `CreateHostGroup`
    - `DeleteHostGroup`
    - `ListAccessConfig`
    - `UpdateAccessConfig`
    - `CreateAccessConfig`
    - `DeleteAccessConfig`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`RestoreExistInstance`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SWR

- _新增特性_
  - 支持接口`ListQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateNamespaceAuth`修改请求体名称`UpdateNamespaceAuthReq` -> `UpdateNamespaceAuthRequestBody`

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListStatistics`请求参数`to`、`from`改为必填
  - 接口`ListQpsTimeline`请求参数`from`、`to`改为必填
  - 接口`ListBandwidthTimeline`请求参数`from`、`to`改为必填
  - 接口`ListTopAbnormal`请求参数`from`、`to`改为必填
  - 接口`ShowEvent`新增响应参数 `cookie`
  - 接口`CreatePremiumHost`:
    - 新增请求参数 `Content-Type`
    - 请求参数`type`类型调整 `string` -> `enum`
  - 接口`ListPremiumHost`新增请求参数 `Content-Type`
  - 接口`UpdatePremiumHost`:
    - 新增请求参数 `Content-Type`
    - 响应参数`type`类型调整 `string` -> `enum`
  - 接口`DeletePremiumHost`新增请求参数 `Content-Type`
  - 接口`ShowPremiumHost`:
    - 新增请求参数 `Content-Type`
    - 响应参数`type`类型调整 `string` -> `enum`
  - 接口`UpdatePremiumHostProtectStatus`新增请求参数 `Content-Type`

# 0.0.70 2021-11-29

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ResettingAppSecretV2`新增响应参数`roma_app_type`

### HuaweiCloud SDK BCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEntityMetric`响应参数`values `类型调整： `object` -> `array`

### HuaweiCloud SDK CBR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListBackups`新增响应参数`provider_id`

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowCluster`新增响应参数`cidrs`

### HuaweiCloud SDK DBSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`SwitchAgent`和`SwitchRiskRule`的响应参数名称调整： `status` -> `result`

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateInstance`新增请求参数`port`

### HuaweiCloud SDK DSC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowScanJobs`新增响应参数`start_time`

### HuaweiCloud SDK EVS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CinderExportToImage`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`NovaShowServer`和`NovaListServersDetails`新增响应参数`os:scheduler_hints`

### HuaweiCloud SDK ELB

- _新增特性_
  - 支持以下接口：
    - `ListApiVersions`
    - `ListSecurityPolicies`
    - `CreateSecurityPolicy`
    - `ShowSecurityPolicy`
    - `UpdateSecurityPolicy`
    - `DeleteSecurityPolicy`
    - `ListSystemSecurityPolicies`
    - `ListQuotaDetails`
    - `ChangeLoadbalancerChargeMode`
    - `BatchUpdatePoliciesPriority`
    - `UpdateIpList`
    - `BatchDeleteIpList`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ShowQuotaDefaults`
  - 接口`ListFlavors`和`ShowFlavor`新增响应参数`flavor_sold_out`和`lcu`, 移除响应参数`availability_zone_ids`
  - 接口`ShowQuota`新增响应参数`members_per_pool`
  - 接口`CreateCertificate`和`UpdateCertificate`新增请求参数`enc_certificate`和`enc_private_key`
  - 接口`ListCertificates`和`ShowCertificate`新增响应参数`enc_certificate`和`enc_private_key`
  - 接口`CreateLoadBalancer`新增请求参数`prepaid_options`、`autoscaling`、`id`
  - 接口`ListLoadBalancers`新增请求参数`elb_virsubnet_type`和`autoscaling`,新增响应参数`autoscaling`和`ip_version`
  - 接口`UpdateLoadBalancer`新增请求参数`prepaid_options`和`autoscaling`,新增响应参数`loadbalancer_id`、`order_id`、`autoscaling`、`ip_version`
  - 接口`ShowLoadBalancer`新增响应参数`autoscaling`和`ip_version`
  - 接口`ShowLoadBalancerStatus`新增响应参数`id`、`type`、`provisioning_status`
  - 接口`CreateListener`新增请求参数`security_policy_id`、`enhance_l7policy_enable`
  - 接口`ListListeners`新增请求参数`enhance_l7policy_enable`和`member_instance_id`，新增响应参数`security_policy_id`、`transparent_client_ip_enable`、`enhance_l7policy_enable`
  - 接口`UpdateListener`新增请求参数`security_policy_id`、`enhance_l7policy_enable`，新增响应参数`security_policy_id`、`transparent_client_ip_enable`、`enhance_l7policy_enable`
  - 接口`ShowListener`新增响应参数`security_policy_id`、`transparent_client_ip_enable`、`enhance_l7policy_enable`
  - 接口`ListPools`新增请求参数`listener_id`和`member_instance_id`
  - 接口`ListMembers`新增请求参数`ip_version`和`member_type`，新增响应参数`member_type`和`instance_id`
  - 接口`UpdateMember`、`ShowMember`、`ListAllMembers`新增响应参数`member_type`和`instance_id`
  - 接口`CreateL7Policy`新增请求参数`priority`、`redirect_url_config`、`fixed_response_config`、`conditions`
  - 接口`ListL7Policies`新增请求参数`priority`,新增响应参数`redirect_url_config`和`fixed_response_config`
  - 接口`UpdateL7Policy`新增请求参数`priority`、`redirect_url_config`、`fixed_response_config`、`rules`,新增响应参数`redirect_url_config`和`fixed_response_config`
  - 接口`ShowL7Policy`新增响应参数`redirect_url_config`和`fixed_response_config`
  - 接口`CreateL7Rule`新增请求参数`conditions`
  - 接口`ListL7Rules`和`ShowL7Rule`新增响应参数`conditions`
  - 接口`UpdateL7Rule`新增请求参数和响应参数`conditions`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListVersionAliases`新增响应参数`ListVersionAliases`
  - 接口`CreateDependency`和`UpdateDependency`的请求参数`name`改为必填
  - 接口`CreateEvent`的请求参数`name`和`content`改为必填

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数`enable_force_switch`

### HuaweiCloud SDK GES

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateGraph`的请求参数`graphSizeTypeIndex`类型调整： `integer` -> `string`

### HuaweiCloud SDK KMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateKey`的请求参数`key_alias`改为必填
  - 接口`EnableKey`、`CancelKeyDeletion`、`ListKeys`、`ListKeyDetail`、`ShowPublicKey`、`ListGrants`、`DeleteImportedKeyMaterial`、`EnableKeyRotation`、`DisableKeyRotation`、`ShowKeyRotationStatus`的请求参数`key_id`改为必填
  - 接口`DeleteKey`的请求参数`key_id`和`pending_days`改为必填
  - 接口`ListKeys`新增请求参数`enterprise_project_id`
  - 接口`CreateRandom`的请求参数`random_data_length`改为必填
  - 接口`CreateDatakey`和`CreateDatakeyWithoutPlaintext`的请求参数`key_id`和`datakey_length`改为必填
  - 接口`EncryptDatakey`的请求参数`key_id`、`plain_text`、`datakey_plain_length`改为必填
  - 接口`DecryptDatakey`的请求参数`key_id`、`cipher_text`、`datakey_cipher_length`改为必填
  - 接口`UpdateKeyAlias`的请求参数`key_id`和`key_alias`改为必填
  - 接口`UpdateKeyDescription`的请求参数`key_id`和`key_description`改为必填
  - 接口`CreateGrant`的请求参数`key_id`、`grantee_principal`、`operations`改为必填
  - 接口`CancelGrant`和`CancelSelfGrant`的请求参数`key_id`和`grant_id`改为必填
  - 接口`EncryptData`的请求参数`key_id`和`plain_text`改为必填
  - 接口`DecryptData`的请求参数`cipher_text`改为必填
  - 接口`CreateParametersForImport`的请求参数`key_id`和`wrapping_algorithm`改为必填
  - 接口`ImportKeyMaterial`的请求参数`key_id`、`import_token`、`encrypted_key_material`改为必填
  - 接口`UpdateKeyRotationInterval`的请求参数`key_id`和`rotation_interval`改为必填
  - 接口`CreateKmsTag`新增请求参数`sequence`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持以下接口：
    - `ChangeProxyScale`
    - `SearchQueryScaleFlavors`
    - `ShowInformationAboutDatabaseProxy`
    - `StartDatabaseProxy`
    - `StopDatabaseProxy`
    - `UpdateReadWeight`
    - `ChangeTheDelayThreshold`
    - `ShowDrReplicaStatus`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPostgresqlDatabases`的响应参数`size`类型调整： `int32` -> `int64`

### HuaweiCloud SDK SMS

- _新增特性_
  - 无
- _解决问题_
  - 修复枚举值中包含中文描述导致参数错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIgnoreRule`新增响应参数`items`
  - 接口`ListEvent`新增请求参数`attacks`
  - 接口`ShowEvent`新增响应参数`host_id`
  - 接口`UpdateHost`新增响应参数`certificatename`

# 0.0.69 2021-11-25

### HuaweiCloud SDK AOM

- _新增特性_
  - 支持应用维护管理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`CheckBackendConnectivity`

### HuaweiCloud SDK APM

- _新增特性_
  - 支持应用性能管理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BCS

- _新增特性_
  - 支持接口`DeleteMemberInvite`、`CreateBlockchainCertByUserName`、`UnfreezeCert`、`FreezeCert`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListCustomerselfResourceRecordDetails`新增请求参数`statistic_type`和响应参数`bill_date`

### HuaweiCloud SDK CBH

- _新增特性_
  - 支持云堡垒机服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowCluster`新增响应参数`platformVersion`
  - 接口`ListClusters`的请求参数`status`新增枚举值`RollingBack`和`RollbackFailed`

### HuaweiCloud SDK CloudRTC

- _新增特性_
  - 支持以下接口：
    - `ListRtcRealtimeScaleDimension`
    - `ListRtcRealtimeQuality`
    - `ListRtcRealtimeNetwork`
    - `ListRtcHistoryUsage`
    - `ListRtcHistoryScale`
    - `ListRtcHistoryQuality`
    - `ListRtcClientQosDetails`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCheck

- _新增特性_
  - 支持代码检查服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DAS

- _新增特性_
  - 无
- _解决问题_
  - 修正部分接口的请求参数`X-Language`的枚举值
- _特性变更_
  - 无

### HuaweiCloud SDK DBSS

- _新增特性_
  - 支持数据库安全服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`、`ListInstances`新增请求参数`tags`

### HuaweiCloud SDK DeH

- _新增特性_
  - 支持专属主机服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreatePrePaidPublicip`、`CreatePublicip`新增请求和响应参数`alias`
  - 接口`ShowPublicip`、`UpdatePublicip`新增响应参数`alias`

### HuaweiCloud SDK GES

- _新增特性_
  - 支持接口`ResizeGraph`、`ExpandGraph`、`UploadFromObs`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持以下接口：
    - `CreateTransfer`
    - `DeleteTransfer`
    - `UpdateTransfer`
    - `ListTransfers`
    - `ListLogStreams`
    - `RegisterDmsKafkaInstance`
    - `CreateNotificationTemplate`
    - `UpdateNotificationTemplate`
    - `ListNotificationTemplates`
    - `DeleteNotificationTemplate`
    - `ShowNotificationTemplate`
    - `ListNotificationTemplate`
    - `UpdateAlarmRuleStatus`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK KMS

- _新增特性_
  - 支持V2版本的接口
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK NLP

- _新增特性_
  - 支持自然语言处理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeHandwriting`移除响应参数`extracted_data`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListSlowlogStatistics`新增请求参数`sort`

### HuaweiCloud SDK SFSTurbo

- _新增特性_
  - 支持弹性文件服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK TMS

- _新增特性_
  - 支持接口`ShowTagQuota`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListServicePublicDetails`、`ListServiceDetails`、`ListServiceConnections`、`AcceptOrRejectEndpoint`、`ListEndpoints`、`UpdateEndpointWhite`、`ListEndpointInfoDetails`的响应参数`created_at`和`updated_at`类型调整: `datetime` -> `string`
  - 接口`CreateEndpointService`的请求参数`vpc_id`和`port_id`改为必填
  - 接口`AcceptOrRejectEndpoint`移除响应参数`error`
  - 接口`ListVersionDetails`和`ListSpecifiedVersionDetails`的响应参数`updated`类型调整: `datetime` -> `string`
  - 接口`ListResourceInstances`和`BatchAddOrRemoveResourceInstance`的请求参数`action`改为必填

### HuaweiCloud SDK WAF

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListEvent`新增请求参数`from`和`to`
  - 接口`ListWhiteblackipRule`新增请求参数`name`

# 0.0.68 2021-11-12

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复响应体类型为`array`或`map`时解析错误的问题
  - 修复请求体数据自动转义的问题
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateNode`新增请求参数`customSan`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`UpdateHttpsInfo`新增请求参数`force_redirect_https`

### HuaweiCloud SDK CloudDeploy

- _新增特性_
  - 支持部署服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListAuditLogs`的响应参数名调整： `total_count` -> `total_record`

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持接口`ShowScanJobs`和`ShowScanJobResults`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDocWatermark`新增请求参数`marked_file_password`

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListGaussMySqlConfigurations`新增请求参数`offset`和`limit`
  - 接口`ShowGaussMySqlProxy`新增响应参数`id`和`spec_code`

### HuaweiCloud SDK GES

- _新增特性_
  - 支持图引擎服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GSL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListProPricePlans`移除请求参数`sim_card_id`、`partner`、`package_type`、`sim_type`
  - 接口`ListSimCards`移除请求参数`tag_id`
  - 接口`ListSimPricePlans`移除请求参数`sim_price_plan_id`
  - 接口`StopSimCard`和`ResetSimCard`移除请求参数`price_plan_list`
  - 接口`ShowSimCard`和`ListSimCards`移除响应参数参数`sn`、`supply_code`、`bundle_id`、`test_type`
  - 接口`StopSimCard`移除响应参数`sim_price_plan_list`
  - 接口`ListSimPools`移除响应参数`order_id`
  - 接口`ListSimPricePlans`移除响应参数`partner`、`partner_pid`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJob`新增响应参数`results`

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListBandwidthDetail`、`ListUsersOfStream`新增请求参数`country`和`protocol`
  - 接口`ListDomainTrafficDetail`、`ListDomainBandwidthPeak`、`ListDomainTrafficSummary`新增请求参数`protocol`
  - 接口`ListTranscodeData`新增请求参数`stream`
  - 接口`ListHistoryStreams`新增请求参数`stream`、`start_time`、`end_time`

### HuaweiCloud SDK LTS

- _新增特性_
  - 支持以下接口：
    - `ShowStructTemplate`
    - `CreateStructTemplate`
    - `UpdateStructTemplate`
    - `DeleteStructTemplate`
    - `ShowAomMappingRules`
    - `CreateAomMappingRules`
    - `UpdateAomMappingRules`
    - `DeleteAomMappingRules`
    - `ShowAomMappingRule`
    - `ListNotificationTopics`
    - `CreateSqlAlarmRule`
    - `UpdateSqlAlarmRule`
    - `ListSqlAlarmRules`
    - `DeleteSqlAlarmRule`
    - `CreateKeywordsAlarmRule`
    - `UpdateKeywordsAlarmRule`
    - `ListKeywordsAlarmRules`
    - `DeleteKeywordsAlarmRule`
    - `ListActiveOrHistoryAlarms`
    - `DeleteActiveAlarms`
    - `ListCharts`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateEditingJob`请求参数名称调整： `const` -> `consts`
  - 接口`CreateEditingJob`移除请求参数`index`，新增请求参数`overlay_input`、`const`、`mix`
  - 接口`ListEditingJob`新增响应参数`output`
  - 接口`CreateTranscodingTask`新增请求参数`hls_init_count`和`hls_init_interval`, 请求参数`video_enhance`新增可选值`EFFICIENCY`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口（V2）`GetJobExeListNew`新增请求参数`job_id`、`user`、`queue`

### HuaweiCloud SDK OCR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeGeneralTable`新增响应参数`confidence`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 支持接口`CreateSystemIssueV4`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIrs`新增响应参数`sequence`
  - 接口`BatchUpdateIrs`新增请求参数`status_id`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`ListSlowLogsNew`、`ListErrorLogsNew`、`UpgradeDbVersion`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VAS

- _新增特性_
  - 支持视频分析服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK VPC

- _新增特性_
  - 支持接口（V3）: `AddVpcExtendCidr`、`RemoveVpcExtendCidr`、`ListVpcs`、`ShowVpc`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.67 2021-10-25

### HuaweiCloud SDK BCS

- _新增特性_
  - 支持接口`ShowBlockchainFlavors`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIndirectPartners`新增响应参数`account_manager_id`和`account_manager_name`

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowHistoryTaskDetails`新增请求参数`create_time`和响应参数`task_type`
  - 接口`ShowHistoryTasks`移除请求参数`create_time`

### HuaweiCloud SDK DNS

- _新增特性_
  - 支持接口`ShowDomainQuota`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateSharedBandwidth`新增请求参数`bandwidth_type`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AddFacesByFile`、`AddFacesByBase64`、`AddFacesByUrl`新增请求参数`single`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`的请求参数和响应参数`num`、`size`类型调整： `integer` -> `string`

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持全球SIM联接服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ImageSearch

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`RunSearchPicture`的请求参数名称调整: `isCrop` -> `is_crop`
  - 接口`RunSearchPicture`新增请求参数`box`

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJob`新增响应参数`current_task`、`image_name`、`process_percent`

### HuaweiCloud SDK IoTDA

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListDevices`新增请求参数`status`
  - 接口`CreateRuleAction`新增请求参数`file_path`

### HuaweiCloud SDK OCR

- _新增特性_
  - 新增接口`RecognizeInsurancePolicy`、`RecognizeFinancialStatement`、`RecognizeQualificationCertificate`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowBackupDownloadLink`新增响应参数`database_name`
  - 接口`ListInstances`新增响应参数`max_iops`和`expiration_time`

### HuaweiCloud SDK SDRS

- _新增特性_
  - 支持存储容灾服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.66 2021-10-19

### HuaweiCloud SDK DCS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowMigrationTask`新增响应参数`backup_id`
  - 接口`ListMigrationTask`新增以下响应参数:
    - `source_instance_name`
    - `source_instance_id`
    - `target_instance_addrs`
    - `target_instance_id`
  - 接口`ListDiagnosisTasks`的响应参数`total_usec_sum`类型调整： `integer` -> `double`

### HuaweiCloud SDK EIP

- _新增特性_
  - 支持以下接口：
    - `ListCommonPools`
    - `ListPublicBorderGroups`
    - `ListPublicipPool`
    - `ShowPublicipPool`
    - `ListShareBandwidthTypes`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListPublicips`新增请求参数`allow_share_bandwidth_type_any`
  - 接口`NeutronListFloatingIps`的请求参数`limit`的类型调整： `string` -> `integer`
  - 接口`NeutronUpdateFloatingIp`请求体的名称调整： `floatingip` -> `NeutronUpdateFloatingIpRequestBody`
  - 接口`ShowPublicip`新增响应参数`allow_share_bandwidth_types`

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTranscodingTask`新增请求参数`hls_init_count`和`hls_init_interval`

### HuaweiCloud SDK VPCEP

- _新增特性_
  - 支持VPC终端节点服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.65 2021-10-11

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复响应体为文件流时没有响应状态码的问题
- _特性变更_
  - 无

### HuaweiCloud SDK APIG

- _新增特性_
  - 支持以下接口：
    - `CreateAppCodeV2`
    - `CreateAppCodeAutoV2`
    - `ListAppCodesV2`
    - `ShowDetailsOfAppCodeV2`
    - `DeleteAppCodeV2`
    - `DebugApiV2`
    - `BatchPublishOrOfflineApiV2`
    - `ListApiVersionsV2`
    - `ChangeApiVersionV2`
    - `ListApiRuntimeDefinitionV2`
    - `ListApiVersionDetailV2`
    - `DeleteApiByVersionIdV2`
    - `ListAclStrategiesV2`
    - `BatchDeleteAclV2`
    - `CreateAclStrategyV2`
    - `ShowDetailsOfAclPolicyV2`
    - `UpdateAclStrategyV2`
    - `DeleteAclV2`
    - `BatchDeleteApiAclBindingV2`
    - `CreateApiAclBindingV2`
    - `DeleteApiAclBindingV2`
    - `ListAclPolicyBindedToApiV2`
    - `ListApisBindedToAclPolicyV2`
    - `ListApisUnbindedToAclPolicyV2`
    - `ListCustomAuthorizersV2`
    - `CreateCustomAuthorizerV2`
    - `ShowDetailsOfCustomAuthorizersV2`
    - `UpdateCustomAuthorizerV2`
    - `DeleteCustomAuthorizerV2`
    - `ExportApiDefinitionsV2`
    - `ImportApiDefinitionsV2`
    - `CreateVpcChannelV2`
    - `ListVpcChannelsV2`
    - `ShowDetailsOfVpcChannelV2`
    - `DeleteVpcChannelV2`
    - `UpdateVpcChannelV2`
    - `AddingBackendInstancesV2`
    - `ListBackendInstancesV2`
    - `DeleteBackendInstanceV2`
    - `ListLatelyApiStatisticsV2`
    - `ListLatelyGroupStatisticsV2`
    - `CreateGatewayResponseV2`
    - `ListGatewayResponsesV2`
    - `ShowDetailsOfGatewayResponseV2`
    - `UpdateGatewayResponseV2`
    - `DeleteGatewayResponseV2`
    - `ShowDetailsOfGatewayResponseTypeV2`
    - `UpdateGatewayResponseTypeV2`
    - `DeleteGatewayResponseTypeV2`
    - `ListTagsV2`
    - `CreateFeatureV2`
    - `ListFeaturesV2`
    - `CreateInstanceV2`
    - `ListInstancesV2`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListSkuInventories`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTask`新增以下响应参数:
    - `create_time`
    - `description`
    - `operate_mode`
    - `related_temp_running_data`
    - `run_status`
    - `update_time`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DisassociateServerVirtualIp`的请求参数`reverse_binding`改为非必填

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`DetectFaceByFile`、`DetectFaceByBase64`、`DetectFaceByUrl`的请求参数`attributes`可选值调整为`2,4,6,7,8,11,12,13,21`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数`instance_mode`

### HuaweiCloud SDK IoTEdge

- _新增特性_
  - 支持IoT边缘服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListClusters`新增响应参数`vpcId`，调整响应参数`start_time`类型: `string` -> `integer`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数`unchangeable_param`、`dry_run`、`count`
  - 接口`CreateRestoreInstance`新增响应参数`enable_ssl`

# 0.0.64 2021-09-29

### HuaweiCloud SDK Core

- _新增特性_
  - 无
- _解决问题_
  - 修复RequestHandler添加顺序错误的问题
- _特性变更_
  - 无

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListSubCustomerBillDetail`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListSubCustomerResFeeRecords`和`ListFreeResources`
  - 接口名调整: `ListParnterAdjustRecords` -> `ListPartnerAdjustRecords`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTemp`新增响应体

### HuaweiCloud SDK DNS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTag`的响应参数`resource_detail`类型调整： `string` -> `object`
  - 接口`CreatePrivateZone`、`UpdatePublicZone`、`DeletePublicZone`的响应参数`masters`类型调整： `string` -> `array`
  - 接口`CreatePrivateZone`和`UpdatePublicZone`的请求参数`ttl`类型调整： `string` -> `integer`
  - 接口`ListRecordSets`、`ListRecordSetsWithLine`、`ListRecordSetsByZone`的请求参数`limit`和`offset`类型调整: `string` -> `integer`
  - 接口`CreatePrivateZone`、`ListRecordSetsByZone`、`ShowRecordSetWithLine`、`ListPtrRecords`、`ListPublicZones`新增响应参数`tags`

### HuaweiCloud SDK ECS

- _新增特性_
  - 支持以下接口:
    - `ListServerTags`
    - `BatchAttachSharableVolumes`
    - `ShowServerAutoRecovery`
    - `BatchResetServersPassword`
    - `ReinstallServerWithoutCloudInit`
    - `ChangeServerOsWithoutCloudInit`
    - `BatchUpdateServersName`
    - `ShowServerPassword`
    - `AssociateServerVirtualIp`
    - `MigrateServer`
    - `ShowServerBlockDevice`
    - `DisassociateServerVirtualIp`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GaussDB

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 调整以下接口名:
    - `ShowMysqlEngineVersion` -> `ShowGaussMySqlEngineVersion`
    - `ShowMysqlFlavors` -> `ShowGaussMySqlFlavors`
    - `CreateMysqlInstance` -> `CreateGaussMySqlInstance`
    - `ShowMysqlInstanceList` -> `ListGaussMySqlInstances`
    - `DeleteMysqlInstance` -> `DeleteGaussMySqlInstance`
    - `ShowMysqlInstanceInfo` -> `ShowGaussMySqlInstanceInfo`
    - `CreateMysqlReadonlyNode` -> `CreateGaussMySqlReadonlyNode`
    - `DeleteMysqlReadonlyNode` -> `DeleteGaussMySqlReadonlyNode`
    - `ExpandMysqlInstanceVolume` -> `ExpandGaussMySqlInstanceVolume`
    - `UpdateMysqlBackupPolicy` -> `UpdateGaussMySqlBackupPolicy`
    - `UpdateMysqlInstanceName` -> `UpdateGaussMySqlInstanceName`
    - `ResetMysqlPassword` -> `ResetGaussMySqlPassword`
    - `ChangeMysqlInstanceSpecification` -> `ChangeGaussMySqlInstanceSpecification`
    - `ListDedicatedResources`  -> `ListGaussMySqlDedicatedResources`
    - `CreateMysqlProxy` -> `CreateGaussMySqlProxy`
    - `DeleteMysqlProxy` -> `DeleteGaussMySqlProxy`
    - `ShowMysqlProxy` -> `ShowGaussMySqlProxy`
    - `ShowMysqlProxyFlavors` -> `ShowGaussMySqlProxyFlavors`
    - `ExpandMysqlProxy` -> `ExpandGaussMySqlProxy`
    - `ListMysqlErrorLog` -> `ListGaussMySqlErrorLog`
    - `ListMysqlSlowLog` -> `ListGaussMySqlSlowLog`
    - `ShowMysqlProjectQuotas` -> `ShowGaussMySqlProjectQuotas`
    - `ShowMysqlQuotas` -> `ShowGaussMySqlQuotas`
    - `SetMysqlQuotas` -> `SetGaussMySqlQuotas`
    - `UpdateMysqlQuotas` -> `UpdateGaussMySqlQuotas`
    - `CreateMysqlBackup` -> `CreateGaussMySqlBackup`
    - `ShowMysqlBackupList` -> `ShowGaussMySqlBackupList`
    - `ShowMysqlBackupPolicy` -> `ShowGaussMySqlBackupPolicy`
    - `ListMysqlConfigurations` -> `ListGaussMySqlConfigurations`
    - `ShowMysqlJobInfo` -> `ShowGaussMySqlJobInfo`

### HuaweiCloud SDK ProjectMan

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListIssueRecordsV4`新增响应参数`id`和`name`
  - 接口`ListProjectIterationsV4`新增响应参数`status`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListFlavors`响应参数`group_type`新增可选值`bigmen`

# 0.0.63 2021-09-26

### HuaweiCloud SDK DRS

- _新增特性_
  - 支持接口`BatchSetPolicy`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MPC

- _新增特性_
  - 支持接口`CreateEditingJob`、`ListEditingJob`、`DeleteEditingJob`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK WAF

- _新增特性_
  - 支持下列接口：
    - `ListIgnoreRule`
    - `ListStatistics`
    - `ListQpsTimeline`
    - `ListBandwidthTimeline`
    - `ListResponseCodeTimeline`
    - `ListTopAbnormal`
    - `ShowConsoleConfig`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.62 2021-09-24

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 修复找不到接口`ListRecordContents`的问题
- _特性变更_
  - 无

# 0.0.61 2021-09-24

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListParnterAdjustRecords`和`ListFreeResourceInfos`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListSubCustomerBillDetail`

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`ShowQuotas`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Classroom

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ApplyJudgement`新增非必填的请求参数`testcases`

### HuaweiCloud SDK Cloudtest

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowTestCaseDetailV2`的请求参数`testcase_number`改为必填，移除请求参数`testcase_id`

### HuaweiCloud SDK GaussDBforopenGauss

- _新增特性_
  - 支持云数据库GaussDB(for openGauss)服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`ListRecordContents`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SWR

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListRepositoryTags`新增响应参数`domain_id`、`scanned`、`tag_type`

# 0.0.60 2021-09-16

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowCluster`新增响应参数`platformVersion`

### HuaweiCloud SDK CDN

- _新增特性_
  - 支持接口`ShowDomainStats`
- _解决问题_
  - 修复调用接口`ShowDomainItemLocationDetails`无响应数据的问题
- _特性变更_
  - 无

### HuaweiCloud SDK DDM

- _新增特性_
  - 支持接口`ListSlowLog`和`ListReadWriteRatio`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DSC

- _新增特性_
  - 支持数据安全中心服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK GSL

- _新增特性_
  - 支持全球SIM联接服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IEC

- _新增特性_
  - 支持智能边缘云服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK IMS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDataImage`新增非必填的请求参数`__support_amd`
  - 接口`ListImages`新增响应参数`__support_amd`

### HuaweiCloud SDK KMS

- _新增特性_
  - 支持接口`ShowPublicKey`、`Sign`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeInvoiceVerification`
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.59 2021-09-10

### HuaweiCloud SDK BSS

- _新增特性_
  - 支持接口`ListSubCustomerBillDetail`、`ListResourceUsageSummary`、`ListResourceUsage`
- _解决问题_
  - 无
- _特性变更_
  - 移除接口`ListResourceUsages`

### HuaweiCloud SDK CBS

- _新增特性_
  - 支持接口`CreateTbSession`、`ExecuteTbSession`、`DeleteTbSession`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 支持接口`AddNode`和`ResetNode`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CDN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateDomain`新增以下响应参数:
    - `range_status`
    - `follow_status`
    - `origin_status`
    - `auto_refresh_preheat`
  - 接口`UpdateDomainMultiCertificates`新增必填请求参数`switch`和可选请求参数`redirect_type`
  - 接口`UpdateHttpsInfo`新增必填请求参数`switch`和可选请求参数`redirect_type`
  - 接口`ShowHistoryTasks`新增可选请求参数`create_time`

### HuaweiCloud SDK DAS

- _新增特性_
  - 支持数据管理服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK DDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ShowJobDetail`新增响应参数`status`和`fail_reason`

### HuaweiCloud SDK EVS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateVolume`的请求参数`size`改为必填

### HuaweiCloud SDK IVS

- _新增特性_
  - 支持人证核身服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持接口`RecognizeInvoiceVerification`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RecognizeIdCard`新增可选请求参数`return_verification`

### HuaweiCloud SDK RDS

- _新增特性_
  - 支持接口`UpdateDatabase`
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListInstances`新增响应参数`alias`
  - 接口`CreateDatabase`新增可选请求参数`comment`

# 0.0.58 2021-08-31

### HuaweiCloud SDK CodeCraft

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCompetitionScore`、`UpdateCompetitionScore`的请求参数`score`类型调整： `string` -> `double`

### HuaweiCloud SDK CPTS

- _新增特性_
  - 支持云性能测试服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK FRS

- _新增特性_
  - 支持以下接口:
    - `DetectLiveByUrl`
    - `DetectLiveFaceByUrl`
    - `DetectLiveByFile`
    - `DetectLiveFaceByFile`
    - `DetectLiveByBase64`
    - `DetectLiveFaceByBase64`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListLiveStreamsOnline`新增响应参数`video_frame_rate`、`audio_frame_rate`、`audio_bitrate`、`resolution`

### HuaweiCloud SDK MRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ClusterScaling`的请求参数`data_volume_count`、`data_volume_size`、`data_volume_type`改为非必填
  - 接口`CreateCluster`新增请求参数`auto_create_default_security_group`

### HuaweiCloud SDK OCR

- _新增特性_
  - 支持文字识别服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK SCM

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ImportCertificate`的请求参数`name`、`certificate`、`certificate_chain`、`private_key`改为必填
  - 接口`PushCertificate`的请求参数`target_project`、`target_service`改为必填

### HuaweiCloud SDK SMN

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListTopics`新增请求参数`enterprise_project_id`、`name`、`fuzzy_name`
  - 接口`ListSubscriptions`新增请求参数`protocol`、`status`、`endpoint`

# 0.0.57 2021-08-25

### HuaweiCloud SDK CBS

- _新增特性_
  - 支持对话机器人服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CodeCraft

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateCompetitionScore`和`UpdateCompetitionScore`请求参数`score`类型调整: `float` -> `string`

### HuaweiCloud SDK DDS

- _新增特性_
  - 支持以下接口：
    - `ShowJobDetail`
    - `SwitchSlowlogDesensitization`
    - `ListFlavorInfos`
    - `UpdateInstanceRemark`
- _解决问题_
  - 无
- _特性变更_
  - 接口`RestoreInstance`、`CreateManualBackup`、`RestoreInstanceFromCollection`新增响应参数`job_id`
  - 接口`ListInstances`新增响应参数`remark`

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListServerInterfaces`新增以下响应参数:
    - `delete_on_termination`
    - `driver_mode`
    - `min_rate`
    - `multiqueue_num`
    - `pci_address`
  - 接口`ListServersDetails`新增响应参数`cpu_options`和`hypervisor`

### HuaweiCloud SDK EIP

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`BatchCreateSharedBandwidths`新增请求参数和响应参数`public_border_group`
  - 接口`AddPublicipsIntoSharedBandwidth`新增响应参数`public_border_group`

### HuaweiCloud SDK FRS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口名调整： `AuthorizeFaceRecognitionService` -> `ShowSubscribes`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ExportFunction`的新增请求参数`function_urn`和`type`
  - 接口`ListStatistics`的请求参数`filter`可选值修改为`monitor_data`、`monthly_report`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
  - 支持以下接口
    - `ListDedicatedResources`
    - `ListFlavorInfos`
    - `ListConfigurationTemplates`
    - `ListInstancesByResourceTags`
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateInstance`新增请求参数`dedicated_resource_id`
  - 接口`ListInstances`新增响应参数`dedicated_resource_id`

### HuaweiCloud SDK ImageSearch

- _新增特性_
  - 支持图像搜索服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 支持接口`RunRecord`
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK MPC

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateTransTemplate`移除请求参数`aspect_ratio`
  - 接口`CreateTranscodingTask`移除请求参数`GOP_structure`、`sr_factor`
  - 接口`CreateTemplateGroup`移除请求参数`GOP_structure`、`sr_factor`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListJobInfoDetail`的响应参数名调整: `taskDetail` -> `task_detail`
  - 接口`ListJobInfoDetail`新增响应参数`count`

# 0.0.56 2021-08-17

### HuaweiCloud SDK AntiDDoS

- _新增特性_
  - 支持流量清洗服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK CCE

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListNodePools`新增以下响应参数：
    - `annotations`
    - `updateTimestamp`
    - `creationTimestamp`
    - `creatingNode`
    - `deletingNode`
    - `conditions`
    - `enterprise_project_id`
  - 接口`ListClusters`新增响应参数`orderID`和`upgradefrom`

### HuaweiCloud SDK CloudTest

- _新增特性_
  - 支持云测服务
- _解决问题_
  - 无
- _特性变更_
  - 无

### HuaweiCloud SDK ECS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`ListServerResizeFlavors`新增响应参数`ecs:instance_architecture`
  - 接口`NovaCreateServers`新增请求参数`tenancy`和`dedicated_host_id`

### HuaweiCloud SDK ELB

- _新增特性_
  - 无
- _解决问题_
  - 修复创建LB异常的问题
- _特性变更_
  - 无

### HuaweiCloud SDK Live

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`CreateRecordCallbackConfig`移除请求参数`key`
  - 接口`ListRecordCallbackConfigs`新增响应参数`sign_type`
  - 接口`ShowDomain`新增响应参数`status_describe`和`service_area`

### HuaweiCloud SDK RDS

- _新增特性_
  - 无
- _解决问题_
  - 无
- _特性变更_
  - 接口`AllowSqlserverDbUserPrivilege`和`RevokeSqlserverDbUserPrivilege`新增请求参数`readonly`

### HuaweiCloud SDK SMS

- _新增特性_
  - 支持主机迁移服务
- _解决问题_
  - 无
- _特性变更_
  - 无

# 0.0.55 2021-08-10

### HuaweiCloud SDK AS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 修改接口`ListResourceInstances`的请求参数`key`、`value`为必填参数

### HuaweiCloud SDK CloudBuild

- _新增特性_
    - 支持编译构建服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListBandwidths`、`ShowPublicip`的响应参数名称调整：`publicip_border_group` -> `public_border_group`

### HuaweiCloud SDK EVS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListVolumes`新增请求参数`server_id`

### HuaweiCloud SDK FRS

- _新增特性_
    - 支持人脸识别服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateDeployment`移除响应参数`order_id`

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 修改接口`UpdateImage`的请求参数`value`为必填参数

### HuaweiCloud SDK Kafka

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`UpdateInstanceTopic`新增请求参数`new_partition_numbers`

### HuaweiCloud SDK LTS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListLogs`新增请求参数`highlight`
    - 接口`ListLogs`的请求参数`search_type`新增可选值`backwards`

### HuaweiCloud SDK MPC

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListTranscodingTask`移除响应参数`dynamic_range`
    - 接口`CreateTransTemplate`移除请求参数`tenant_id`
    - 接口`CreateThumbnailsTask`移除请求参数`percent`、`frame_type`
    - 接口`CreateTranscodingTask`移除请求参数`black_enhance`

### HuaweiCloud SDK MRS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`DescribeCluster`新增响应参数`start_time`、`state`

### HuaweiCloud SDK SWR

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowAccessDomain`新增以下响应参数：
        - `namespace`
        - `repository`
        - `access_domain`
        - `permit`
        - `deadline`
        - `description`
        - `creator_id`
        - `creator_name`
        - `created`
        - `updated`
        - `status`

### HuaweiCloud SDK VPC

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`NeutronListSubnets`新增请求参数`enable_dhcp`
    - 接口`NeutronListSecurityGroups`新增响应参数`security_groups_links`

# 0.0.54 2021-07-27

### HuaweiCloud SDK Classroom

- _新增特性_
    - 支持接口`ApplyJudgement`、`ShowJudgementDetail`、`ShowJudgementFile`
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.53 2021-07-26

### HuaweiCloud SDK CDN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowHistoryTasks`移除响应参数`urls`、`task_id`
    - 接口`ShowHistoryTaskDetails`移除响应参数`task_id`、`process_reason`,请求参数`process_reason`类型调整： `integer`->`string`
    - 接口`ShowTopUrl`移除请求参数`user_domain_id`、`task_id`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 支持接口`ShowPlans`
- _解决问题_
    - 无
- _特性变更_
    - 无
    
### HuaweiCloud SDK DCS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`UpdateConfigurations`新增请求参数`dcs_cluster_proxy2_node`

### HuaweiCloud SDK DDM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`UpdateUser`移除请求参数`extend_authority`

### HuaweiCloud SDK DDS

- _新增特性_
    - 支持接口`UpdateClientNetwork`
- _解决问题_
    - 无
- _特性变更_
    - 接口`SetBalancerWindow`的请求参数`start_time`、`stop_time`改为非必填
    - 接口`CreateInstance`新增请求参数`port`，新增响应参数`port`

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 支持接口`EnableLtsLogs`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowFunctionCode`新增响应参数`concurrent_num`、`id`、`encrypted_user_data`
    - 接口`ListFunctions`新增响应参数`func_vpc_id`、`encrypted_user_data`、`long_time`、`log_group_id`、`log_stream_id`、`type`，移除响应参数`version_description`、`last_modified_utc`、`dependencies`
    - 接口`UpdateVersionAlias`移除请求参数`name`、`last_modified`、`alias_urn`
    - 接口`ShowFunctionConfig`新增响应参数`encrypted_user_data`、`long_time`、`log_group_id`、`log_stream_id`、`type`，移除响应参数`version_description`、`concurrency`
    - 接口`UpdateFunctionConfig`移除请求参数`version_description`、`concurrency`、`depend_list`，新增请求参数`encrypted_user_data`、`long_time`、`log_group_id`、`log_stream_id`、`type`
    - 接口`ListFunctionVersions`移除响应参数`last_modified_utc`、`concurrency`，新增响应参数`encrypted_user_data`、`long_time`、`log_group_id`、`log_stream_id`、`type`
    - 接口`UpdateTrigger`的请求参数`size`类型调整： `string`->`integer`
    - 接口`ShowDependency`的响应参数`size`类型调整： `string`->`integer`
    - 接口`UpdateDependency`的响应参数`size`类型调整： `string`->`integer`

### HuaweiCloud SDK HSS

- _新增特性_
    - 支持接口`ListEvents`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowDomain`移除响应参数`domain_source`

### HuaweiCloud SDK MPC

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListTranscodeDetail`移除响应参数`language`
    - 接口`CreateThumbnailsTask`移除请求参数`project_id`、`tenant_project_id`、`domain_name`、`canonical_grant_id`
    - 接口`ListTranscodeDetail`新增响应参数`audit_report`
    - 接口`QueryTranscodingsTask`移除响应参数`output_url`
    - 接口`CreateTranscoding`新增请求参数`audit`，移除请求参数`special_effect`、`quality_enhance`、`template_extend`
    - 接口`ListWatermarkTemplate`移除响应参数`template_id`、`error`
    - 接口`CreateVodTranscodingTask`移除请求参数`multidrm`、`preview_duration`

### HuaweiCloud SDK VOD

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateAssetByFileUpload`的请求参数`auto_publish`类型调整： `string`->`integer`，并配置可选值为`0`、`1`

### HuaweiCloud SDK WAF

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 无
      接口`ListEvent`的响应参数`response_time`、`response_size`类型调整： `string`->`integer`

# 0.0.52 2021-07-16

### HuaweiCloud SDK AS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateScalingV2Policy`新增请求参数`description`
    - 接口`ShowScalingV2Policy`、`ShowScalingGroup`新增响应参数`description`

### HuaweiCloud SDK DCS

- _新增特性_
    - 支持更多接口：
        - `CreateDiagnosisTask`
        - `CreateRedislog`
        - `CreateRedislogDownloadLink`
        - `ListDiagnosisTasks`
        - `ListRedislog`
        - `ListSlowlog`
        - `ShowDiagnosisTaskDetails`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListInstances`新增请求参数`include_delete`

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - [Issue 40](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/40): 修复响应参数`__lazyloading`类型定义错误的问题
- _特性变更_
    - 无

# 0.0.51 2021-07-09

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决了query参数名称存在大写字母的情况下鉴权失败的问题
- _特性变更_
    - 无

### HuaweiCloud SDK BMS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口`ListBareMetalServers`的响应参数`addresses`数据结构定义错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListProtectable`新增响应参数`smn_notify`、`threshold`
    - 接口`AssociateVaultPolicy`新增请求参数`add_policy_ids`和响应参数`without_any_tag`、`smn_notify`、`threshold`

### HuaweiCloud SDK CCE

- _新增特性_
    - 支持接口`RemoveNode`、`MigrateNode`
- _解决问题_
    - 无
- _特性变更_
    - 接口`DeleteCluster`新增请求参数`tobedeleted`

### HuaweiCloud SDK CDN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowTopUrl`的请求参数`start_time`、`end_time`改为必填参数，`domain_name`新增可选值`outside_mainland_china`
    - 接口`ShowDomainItemDetails`新增请求参数`service_area`

### HuaweiCloud SDK DDM

- _新增特性_
    - 支持分布式数据库中间件服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DNS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreatePublicZone`的响应参数`masters`、`zones`类型调整：`string`->`array`

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateSharedBandwidth`、`ListBandwidths`新增响应参数`publicip_border_group`

### HuaweiCloud SDK GaussDB

- _新增特性_
    - 支持云数据库服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`GlanceCreateImageMetadata`新增响应参数`__root_origin`、`checksum`、`size`
    - 接口`GlanceAddImageMember`移除请求参数`deleted`、`deleted_at`,新增以下请求参数：
        - `__lazyloading`
        - `__os_feature_list`
        - `__root_origin`
        - `__sequence_num`
        - `__support_agent_list`
        - `__system__cmkid`
        - `active_at`
        - `hw_vif_multiqueue_enabled`
        - `max_ram`
        - `__image_location`
        - `__is_config_init`
        - `__account_code`

### HuaweiCloud SDK IoTDA

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListRules`新增响应参数`edge_node_ids`、`last_update_time`

### HuaweiCloud SDK LTS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListStructuredLogsWithTimeRange`响应参数`context`类型调整： `string`->`array`
    - 接口名称调整：
        - `UpdateLogContents`->`ListLogs`
        - `UpdateLogContents2`->`ListQueryStructuredLogs`
        - `UpdateLogContents3`->`ListStructuredLogsWithTimeRange`

### HuaweiCloud SDK MRS

- _新增特性_
    - 支持接口`ListClusters`
- _解决问题_
    - 无
- _特性变更_
    - 接口`DescribeCluster`的响应参数`clusterType`类型调整：`string`->`integer`

### HuaweiCloud SDK SWR

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowRepository`新增响应参数`domain_id`、`priority`
    - 接口`CreateRetention`新增响应参数`template`

# 0.0.50 2021-06-29

### HuaweiCloud SDK CCE

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateNodePool`、`ShowNodePool`、`UpdateNodePool`、`DeleteNodePool`新增请求参数`storage`

### HuaweiCloud SDK DRS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`BatchUpdateUser`的参数`selected`类型调整： `string`->`boolean`

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListInstances`的响应参数`port`类型调整： `string`->`integer`
    - 接口`ListInstances`的响应参数名称调整： `storage_engine`->`mode`
    - 接口`ListSlowLogs`移除响应参数`node_name`，新增响应参数`time`

### HuaweiCloud SDK MRS

- _新增特性_
    - 无
- _解决问题_
    - 修复部分参数定义错误问题
- _特性变更_
    - 接口`CreateCluster`移除请求参数`start_time`、`state`

### HuaweiCloud SDK NAT

- _新增特性_
    - 无
- _解决问题_
    - 修复接口`ListNatGateways`的请求参数`project_id`重复的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowInformationAboutDatabaseProxy`的响应参数`port`、`node_num`类型调整： `string`->`integer`

# 0.0.49 2021-06-25

### HuaweiCloud SDK APIG

- _新增特性_
    - 支持更多接口：
        - `ListGatewayResponsesV2`
        - `UpdateGatewayResponseV2`
        - `DeleteGatewayResponseV2`
        - `UpdateGatewayResponseTypeV2`
        - `DeleteGatewayResponseTypeV2`
        - `DeleteInstancesV2`
        - `UpdateInstanceV2`
        - `ListInstancesV2`
        - `RemoveEipV2`
        - `UpdateEngressEipV2`
        - `RemoveEngressEipV2`
        - `ListFeaturesV2`
        - `UpdateDomainV2`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BMS

- _新增特性_
    - 支持接口`ChangeBaremetalServerOs`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ChangeBaremetalServerName`的响应参数名称调整：`server_tags`->`sys_tags`

### HuaweiCloud SDK CDN

- _新增特性_
    - 支持接口`ShowQuota`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowHistoryTaskDetails`的请求参数`url`类型调整：`integer`->`string`

### HuaweiCloud SDK CloudRTC

- _新增特性_
    - 支持华为云实时音视频服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DRS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`BatchUpdateUser`参数`is_transfer`、`selected`类型调整：`string`->`boolean`

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`KeystoneListPermissions`新增请求参数`permission_type`、`display_name`、`catalog`、`type`

### HuaweiCloud SDK LTS

- _新增特性_
    - 支持云日志服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Meeting

- _新增特性_
    - 支持接口`InviteShare`
- _解决问题_
    - 无
- _特性变更_
    - 接口`SetMultiPicture`新增请求参数`multiPicSaveOnly`
    - 接口`SearchHisMeetings`新增响应参数`leftReason`

### HuaweiCloud SDK VOD

- _新增特性_
    - 支持视频点播服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.48 2021-06-21

### HuaweiCloud SDK BMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ChangeBaremetalServerName`新增响应参数`server_tags`、`enterprise_project_id`、`group`

### HuaweiCloud SDK CCE

- _新增特性_
    - 无
- _解决问题_
    - [Issue 22](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/22): 修正接口`ListAddonInstances`的响应参数`status`可选值
- _特性变更_
    - 无

### HuaweiCloud SDK CDN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListDomains`移除请求参数`user_domain_id`
    - 接口名称调整：`ShowRefer` -> `ShowReferer`

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowTemplateDetail`新增请求参数：
        - `template_url`
        - `create_time`
        - `last_modify_time`
        - `can_update`
        - `can_delete`
        - `need_hub`

### HuaweiCloud SDK Live

- _新增特性_
    - 新增支持接口:
        - `CreateRecordCallbackConfig`
        - `ShowRecordCallbackConfig`
        - `UpdateRecordCallbackConfig`
        - `DeleteRecordCallbackConfig`
        - `ListRecordCallbackConfigs`
        - `UpdateRecordRule`
        - `ShowRecordRule`
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：
        - `CreateRecordConfig` -> `CreateRecordRule`
        - `DeleteRecordConfig` -> `DeleteRecordRule`
        - `ListRecordConfigs` -> `ListRecordRules`
    - 移除接口：
        - `ShowTraffic`
        - `ShowBandwidth`
        - `ShowOnlineUsers`

### HuaweiCloud SDK Kafka

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowGroups`的响应参数`partitions`类型调整： `array[string]` -> `array[integer]`

### HuaweiCloud SDK RabbitMQ

- _新增特性_
    - 无
- _解决问题_
    - 修复编译失败的问题
- _特性变更_
    - 无

# 0.0.47 2021-06-10

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口`ListFreeResources`、`ListFreeResourceUsages`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CDN

- _新增特性_
    - 支持内容分发网络服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DRS

- _新增特性_
    - 支持数据复制服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 新增支持接口
        - `ImportFunction`
        - `ExportFunction`
        - `AsyncInvokeReservedFunction`
        - `DeleteReservedInstanceById`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK OSM

- _新增特性_
    - 支持工单管理服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 新增支持接口`SetBinlogClearPolicy`、`ShowBinlogClearPolicy`
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListOffSiteInstances`新增请求参数`offset`、`limit`

# 0.0.46 2021-06-04

### HuaweiCloud SDK CCE

- _新增特性_
    - 无
- _解决问题_
    - [Issue 20](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/20): 修复`extendParam`类型定义错误的问题
- _特性变更_
    - 接口`DeleteCluster`新增请求参数`tobedeleted`

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口`ShowQuotas`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
    - 新增支持接口`ListComplexQueryDevice`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK GaussDBforNoSQL

- _新增特性_
    - 支持`GaussDBforNoSQL`服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 新增支持接口`ShowQuotas`
- _解决问题_
    - 无
- _特性变更_
    - 接口`StartInstanceRestartAction`的请求参数`restart`类型调整：string -> object

# 0.0.45 2021-05-25

### HuaweiCloud SDK AS

- _新增特性_
    - 新增支持接口：
        - `ListApiVersions`
        - `ShowApiVersion`
        - `BatchProtectScalingInstances`
        - `BatchRemoveScalingInstances`
        - `CreateScalingTagInfo`
        - `BatchResumeScalingPolicies`
        - `BatchPauseScalingPolicies`
        - `PauseScalingGroup`
        - `BatchSetScalingInstancesStandby`
        - `BatchUnsetScalingInstancesStandby`
        - `ResumeScalingPolicy`
        - `PauseScalingPolicy`
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：
        - `ExecuteScalingPolicies` -> `BatchDeleteScalingPolicies`
        - `EnableOrDisableScalingGroup` -> `ResumeScalingGroup`
        - `UpdateScalingGroupInstance` -> `BatchAddScalingInstances`
        - `CompleteLifecycleAction` -> `AttachCallbackInstanceLifeCycleHook`
    - 移除接口：
        - `DeleteScalingTags`
    - `ListScalingGroups` 接口新增参数 `enterprise_project_id`
    - `ListScalingActivityV2Logs` 接口新增参数 `log_id`

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口：
        - 查询月度成本 `ListCustomerBillsMonthlyBreakDown`
        - 查询订单可用折扣 `ListOrderDiscounts`
- _解决问题_
    - 无
- _特性变更_
    - 查询客户消费记录接口 `ListSubCustomerResFeeRecords` 增加 query 参数：bill_date_begin 和 bill_date_end

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 新增支持接口：停止流水线 `StopPipelineNew`
- _解决问题_
    - 无
- _特性变更_
    - 移除接口 `StartPipeline`、`StopPipeline`

### HuaweiCloud SDK DMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：（查询项目标签）ShowProjectTags -> ShowQueueProjectTags

### HuaweiCloud SDK EPS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListEnterpriseProject`的请求参数`offset`从必填改为非必填

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 新增支持接口：
        - `ListFunctionAsyncInvokeConfig`
        - `ShowFunctionAsyncInvokeConfig`
        - `DeleteFunctionAsyncInvokeConfig`
        - `UpdateFunctionAsyncInvokeConfig`
- _解决问题_
    - 无
- _特性变更_
    - 接口`DeleteVersionAlias`、`UpdateVersionAlias`、`ShowVersionAlias`的请求参数名称调整：`name` -> `alias_name`
    - 接口`DeleteFunctionTrigger`、`UpdateTrigger`、`ShowFunctionTrigger`的请求参数名称调整：`triggerId` -> `trigger_id`

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口 `CreateUsers` 新增请求体参数和响应体参数 _access_mode_
    - 接口 `DeleteBindingDevice` 将请求体参数 _authentication_code_ 设置为必填参数

### HuaweiCloud SDK Kafka

- _新增特性_
    - 新增支持接口：
        - `CreateInstanceUser`
        - `BatchDeleteInstanceUsers`
        - `ShowInstanceUsers`
        - `ShowTopicAccessPolicy`
        - `UpdateTopicAccessPolicy`
        - `ShowKafkaTopicPartitionDiskusage`
        - `ShowInstanceMessages`
        - `ResetUserPassword`
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：
        - `ShowInstanceTags` -> `ShowKafkaTags`
        - `ShowProjectTags` -> `ShowKafkaProjectTags`
        - `BatchCreateOrDeleteInstanceTag` -> `BatchCreateOrDeleteKafkaTag`
    - 接口 `BatchCreateOrDeleteInstanceTag` 请求体名称调整：`BatchCreateOrDeleteInstanceTagRequestBody`
      -> `BatchCreateOrDeleteKafkaTagRequestBody`
    - 接口 `UpdateSinkTaskQuota` 请求体 `UpdateSinkTaskQuotaReq` 的请求参数 `sink_max_tasks` 数据类型调整：String → Integer

### HuaweiCloud SDK OMS

- _新增特性_
    - 支持对象存储迁移服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RabbitMQ

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整：
        - `BatchCreateOrDeleteInstanceTag` -> `BatchCreateOrDeleteRabbitMqTag`
        - `ShowProjectTags` -> `ShowRabbitMqProjectTags`
        - `ShowInstanceTags` -> `ShowRabbitMqTags`
    - 接口 `BatchCreateOrDeleteInstanceTag` 请求体名称调整：BatchCreateOrDeleteInstanceTagRequestBody ->
      BatchCreateOrDeleteRabbitMqTagRequestBody

# 0.0.43-rc 2021-05-14

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 解决了使用接口`NovaShowKeypair`获取秘钥，结果解析异常的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListInstances`的响应字段`type`新增结果值: `CLOUDSSD`、`LOCALSSD`
    - 接口`ListBackups`新增可选的请求参数`complete_version`
    - 将接口`ListSlowlogStatistics`的请求参数`type`从非必填改为必填

# 0.0.42-rc 2021-05-10

### HuaweiCloud SDK BMS

- _新增特性_
    - 新增支持接口`BatchCreateBaremetalServerTags`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口`MigrateAz`、`ListAz2Migrate`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EPS

- _新增特性_
    - 无
- _解决问题_
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/17): 修复`EpDetailType`枚举定义错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK KPS

- _新增特性_
    - 无
- _解决问题_
    - [Issue 19](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/19): 修复`ListKeypairs`响应体类型的问题
- _特性变更_
    - 无

### HuaweiCloud SDK MRS

- _新增特性_
    - 新增支持相关接口：
        - `BatchDeleteClusterTags`
        - `CreateClusterTag`
        - `DeleteClusterTag`
        - `ListClusterTags`
        - `ListAllTags`
        - `BatchCreateClusterTags`
        - `ListClustersByTags`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListOffSiteInstances`的响应体名称调整：`OffSiteInstanceListResponse` -> `OffSiteInstanceListResponseBody`
    - 接口`ListOffSiteInstances`的响应字段名称调整：`offsite_backup_instances` -> `offsite_backup_instance`

# 0.0.41-rc 2021-04-30

### HuaweiCloud SDK BCS

- _新增特性_
    - 新增支持查询异步操作结果的接口`ListOpRecord`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持的接口
        - 查询集群均衡设置 `ShowShardingBalancer`
        - 设置集群均衡开关 `SetBalancerSwitch`
        - 设置集群均衡活动时间窗 `SetBalancerWindow`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK HSS

- _新增特性_
    - 新增支持查询弹性云服务器状态列表的接口`ListHosts`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 查询账号配额接口`ShowDomainQuota`的请求参数`type`增加可选值：
        - `assigment_group_mp`
        - `assigment_agency_mp`
        - `assigment_group_ep`
        - `assigment_user_ep`

### HuaweiCloud SDK IoTDA

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 下线订阅管理相关接口：
        - `ListSubscriptions`
        - `CreateSubscription`
        - `UpdateSubscription`
        - `ShowSubscription`
        - `DeleteSubscription`

### HuaweiCloud SDK MPC

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`CreateMpeCallBack`新增请求参数`language`、`sky_switch`
    - 接口`CreateTranscodingTask`的请求参数`subtitle_type`可选值调整为`0`、`1`、`2`

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 获取项目详情接口`ShowProjectInfoV4`的响应体新增字段`project_code`

# 0.0.40-rc 2021-04-15

### HuaweiCloud SDK Moderation

- _新增特性_
    - 支持Moderation内容审核服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Image

- _新增特性_
    - 支持Image图像识别服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 新增支持管理数据库的接口
        - `CreateSqlserverDatabase`
        - `DeleteSqlserverDatabase`
        - `ListSqlserverDatabases`
    - 新增支持管理用户的接口
        - `CreateSqlserverDbUser`
        - `ListSqlserverDbUsers`
        - `ListAuthorizedSqlserverDbUsers`
        - `DeleteSqlserverDbUser`
        - `AllowSqlserverDbUserPrivilege`
        - `RevokeSqlserverDbUserPrivilege`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口`DeleteDatabaseUser`、`DeleteDatabaseRole`、`ShowConnectionStatistics`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ListIssuesV4`, `ListChildIssuesV4`响应体新增字段`closed_time` 、`id` 、`created_time`

### HuaweiCloud SDK VPC

- _新增特性_
    - 无
- _解决问题_
    - 修复问题，开放vpc和子网的标签
- _特性变更_
    - 无

# 0.0.39-rc 2021-03-30

### HuaweiCloud SDK Kafka

- _新增特性_
    - 无
- _解决问题_
    - 修复查询消息接口没有时间戳字段的问题
- _特性变更_
    - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口`ShowIssueV4`、`UpdateIssueV4`的响应体 `IssueResponseV4`增加`name`属性
    - 将接口`ShowProjectWorkHours`、`ListProjectWorkHours`的响应体`ShowProjectWorkHoursResponseBody`下的属性`work_time`
      修改为`work_date`

### HuaweiCloud SDK SMN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 将接口`PublishMessage`的请求参数`protocol`从必填改为非必填
    - 将接口 `PublishMessage` 请求体的 `PublishMessageRequestBody` 类属性 `subject` 由必填改为非必填

# 0.0.38-rc 2021-03-26

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决自动获取 domainId 的场景下报错校验 AK SK 校验失败的问题
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 无
- _解决问题_
    - 解决 ListLiveStreamsOnline 接口响应体反序列化失败的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 解决 ListSlowlogStatistics 接口响应体部分字段为空的问题
- _特性变更_
    - 无

### HuaweiCloud SDK SMN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 将接口 `ListSlowlogStatistics` 下的 `property` 属性从必填调整为非必填

# 0.0.37-rc 2021-03-19

### HuaweiCloud SDK Core

- _新增特性_
    - 支持文件上传
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 解决 ListFlavors 接口响应体反序列化出错的问题
- _特性变更_
    - 无

# 0.0.36-rc 2021-03-16

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 申请包周期弹性公网IP接口增加 `enterprise_project_id` 字段

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 无
- _解决问题_
    - 修复接口无法调用的问题
- _特性变更_
    - 无

# 0.0.35-rc 2021-03-15

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 当接口返回体带有 `\n` 等特殊字符时，反序列化失败问题修复
- _特性变更_
    - 当用户传入的 `endpoint` 未带协议前缀时，支持自动加上 `https` 前缀

### HuaweiCloud SDK CES

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 使用资源分组创建告警规则时，维度字段调整为非必填，即 `CreateAlarmRequestBody` 类中的属性 `metric` 对应类调整：
      `MetricInfoForAlarm` → `MetricForAlarm`
    - 将 `MetricsDimension`类中的属性资源维度(`name`)字段调整为必填字段，涉及接口：批量查询监控数据、添加监控数据、创建资源分组、更新分组

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口：
        - 根据备份恢复新实例 `RestoreNewInstance`
        - 查询实例节点会话 `ListSessions`
        - 终结实例节点会话 `DeleteSession`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 新增支持接口：查询云服务器组详情 `ShowServerGroup`
- _解决问题_
    - 无
- _特性变更_
    - 云服务器获取密码接口名调整：`ShowWindowsServerPasswordResponseBody` → `ShowServerPasswordResponseBody`
    - 云服务器清除密码接口名调整：`DeleteWindowsServerPassword` → `DeleteServerPassword`

### HuaweiCloud SDK ELB

- _新增特性_
    - 新增支持接口：查询当前租户下的后端服务器列表 `ListAllMembers`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口获取依赖包列表 `ListDependencies` 响应体的属性 `size` 类型调整：string → int64

### HuaweiCloud SDK IAM

- _新增特性_
    - 新增支持接口：
        - 查询身份提供商详情 `KeystoneShowIdentityProvider`
        - 注册身份提供商 `KeystoneCreateIdentityProvider`
        - 更新身份提供商 `KeystoneUpdateIdentityProvider`
        - 删除身份提供商 `KeystoneDeleteIdentityProvider`
        - 获取联邦认证token(OpenId Connect Id token方式) `CreateTokenWithIdToken`
- _解决问题_
    - 无
- _特性变更_
    - 下线接口：获取联邦认证unscoped token `CreateUnscopeTokenByIdpInitiated`

### HuaweiCloud SDK IMS

- _新增特性_
    - 新增支持接口：
        - 按标签查询镜像 `ListImageByTags`
        - 查询租户所有镜像标签 `ListImagesTags`
        - 查询镜像标签 `ListImageTags`
        - 添加镜像标签 `AddImageTag`
        - 删除镜像标签 `DeleteImageTag`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IoTDA

- _新增特性_
    - 支持设备接入服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 新增支持接口：
        - 创建工作项类型自定义字段 `CreateCustomfields`
        - 查询人均bug `ShowBugsPerDeveloper`
        - 查询需求按时完成率 `ShowCompletionRate`
        - 查询缺陷密度 `ShowBugDensityV2`
        - 获取项目详情 `ShowProjectInfoV4`
- _解决问题_
    - 修改接口命名错误：`ShowtIssueCompletionRate` → `ShowIssueCompletionRate`
- _特性变更_
    - `ListProjectV4` 接口响应体中的 `created_time` 和 `updated_time` 属性类型调整：string → int64

### HuaweiCloud SDK RDS

- _新增特性_
    - 支持 Postgresql 相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK SCM

- _新增特性_
    - 支持 SSL 证书管理服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.34-rc 2021-02-27

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 优化 README 文档描述及 CHANGELOG 日志格式

### HuaweiCloud SDK BMS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：`WindowsBaremetalServerCleanPwd` → `DeleteWindowsBareMetalServerPassword`
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 所有类型为 `float32` 的属性统一调整为 `float64`

### HuaweiCloud SDK CCE

- _新增特性_
    - 无
- _解决问题_
    - 修复参数 `Content-Type` 必填且返回 "Unsupported Content Type" 错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK CES

- _新增特性_
    - 无
- _解决问题_
    - 修复类 `CreateAlarmResponse` 出现循环依赖的问题
- _特性变更_
    - 无

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口：获取慢日志下载链接 `DownloadSlowlog`、获取错误日志下载链接 `DownloadErrorlog`
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DGC

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：
        - `ModifyJob` → `UpdateJob`
        - `ModifyScript` → `UpdateScript`
        - `ModifyResource` → `UpdateResource`
- _特性变更_
    - 无

### HuaweiCloud SDK DIS

- _新增特性_
    - 无
- _解决问题_
    - 修复类出现循环依赖的问题
        - `CommitCheckpointRequest`
        - `PutRecordsRequest`
        - `CreateAppRequest`
        - `UpdatePartitionCountRequest`
- _特性变更_
    - 无

### HuaweiCloud SDK ELB

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：`ListMenbers` → `ListMembers`
    - 修复接口 `ShowMember` 需要传 query 参数 `X-Auth-Token` 的问题
- _特性变更_
    - 无

### HuaweiCloud SDK EPS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：`ModifyEnterpriseProject` → `UpdateEnterpriseProject`
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 修复 `KeystoneUserResult` 类中属性名的错误定义问题：`pwd_stength` → `pwd_strength`
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口方法命名不合理的问题：
        - `DoManualBackup` → `CreateManualBackup`
        - `ModifyConfiguration` → `UpdateConfiguration`
        - `ModifyInstanceConfiguration` → `UpdateInstanceConfiguration`
    - 修复类 `CreateInstanceResponse` 和 `CreateConfigurationResponse` 出现循环依赖的问题
    - 修复接口 `CreateInstance` 不可用的问题
- _特性变更_
    - 接口 `StartInstanceAction` 请求中单机转主备场景增加 `is_auto_pay` 属性

# 0.0.33-rc 2021-02-07

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 修复请求体类型为string时请求失败的问题
- _特性变更_
    - 无

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 查询镜像支持的OS列表(ListOsVersions)接口返回体属性 `os_bit` 数据类型调整：string → int32

### HuaweiCloud SDK Live

- _新增特性_
    - 新增支持接口：获取直播播放日志(ListLiveSampleLogs)、创建直播域名(CreateDomain)、删除直播域名(DeleteDomain)、修改直播域名(UpdateDomain)、查询直播域名(
      ShowDomain)、创建直播域名映射关系(CreateDomainMapping)、删除直播域名映射关系(DeleteDomainMapping)
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK MRS

- _新增特性_
    - 支持MapReduce服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.32-rc 2021-01-30

### HuaweiCloud SDK DDS

- _新增特性_
    - 新增支持接口：查询当前支持的API版本信息列表、查询指定API版本信息、设置审计日志策略、查询审计日志策略、获取审计日志列表
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整: UpdateAutoTerminateTimeServer → UpdateServerAutoTerminateTime

### HuaweiCloud SDK EVS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建云硬盘接口支持指定专属存储池ID
    - 查询配额相关接口属性 `allocated` 类型调整: string → int

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 查询IAM用户详情接口响应体增加属性`access_mode`

# 0.0.31-rc 2021-01-25

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 设置默认`DefaultTimeout`为120秒

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口：查询订单可用折扣
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DNS

- _新增特性_
    - 支持云解析服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 新增支持接口：修改云服务器云主机销毁时间
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.30-rc 2021-01-15

### HuaweiCloud SDK Core

- _新增特性_
    - Region管理支持使用`ValueOf`方法获取`region`信息
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DGC

- _新增特性_
    - 新增支持接口：作业相关接口（增删改查）、脚本相关接口（增删改查）、资源相关接口（增删改查）
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建/查询用户接口响应字段 `is_domain_owner` 类型调整：string → boolean

### HuaweiCloud SDK Live

- _新增特性_
    - 新增支持接口：查询直播中的流信息
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 新增支持接口：查询跨区域备份策略、设置跨区域备份策略、获取跨区域备份列表、查询跨区域可恢复时间段、查询跨区域备份实例列表、查询API版本列表、查询指定的API版本信息
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK SWR

- _新增特性_
    - 支持容器镜像服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.29-beta 2020-12-31

### HuaweiCloud SDK BMS

- _新增特性_
    - 无
- _解决问题_
    - 修复查询裸金属服务器详情接口问题
    - 修复查询裸金属服务器详情列表接口问题
    - 完善查询裸金属服务器网卡信息接口字段
- _特性变更_
    - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
    - 新增支持接口：查询当前账号访问权限（ShowAccountStatus）
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DCS

- _新增特性_
    - 无
- _解决问题_
    - 修正接口返回数据类型避免反序列化失败：
        - 查询所有实例列表接口：响应参数`Tags`类型调整 Tag → ResourceTag
        - 查询慢日志接口：响应参数`duration`类型调整 int32 → string
        - 查询大key分析详情接口：响应参数`db`类型调整 int32 → string
        - 查询热key分析详情接口：响应参数`db`类型调整 int32 → string
- _特性变更_
    - 无

### HuaweiCloud SDK DGC

- _新增特性_
    - 支持数据湖治理中心服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DIS

- _新增特性_
    - 支持数据接入服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建实例接口响应类型调整

### HuaweiCloud SDK SMN

- _新增特性_
    - 无
- _解决问题_
    - 修正消息发布接口请求参数：Object → Map<String, String>
- _特性变更_
    - 无

# 0.0.28-beta 2020-12-28

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 调用临时AK/SK返回异常问题修复
- _特性变更_
    - 无

### HuaweiCloud SDK DCS

- _新增特性_
    - 无
- _解决问题_
    - 修改缓存接口port类型为integer
- _特性变更_
    - 无

# 0.0.27-beta 2020-12-25

### HuaweiCloud SDK DCS

- _新增特性_
    - 无
- _解决问题_
    - 修复SDK在Mac操作系统上无法正常编译的问题
- _特性变更_
    - 接口ListInstances请求Query参数名称调整：id → instance_id

### HuaweiCloud SDK DDS

- _新增特性_
    - 支持文档数据库服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Kafka

- _新增特性_
    - 无
- _解决问题_
    - 修复SDK在Mac操作系统上无法正常编译的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RabbitMQ

- _新增特性_
    - 无
- _解决问题_
    - 修复SDK在Mac操作系统上无法正常编译的问题
- _特性变更_
    - 无

### HuaweiCloud SDK RMS

- _新增特性_
    - 新增支持接口：资源查询相关接口、Region查询相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.26-beta 2020-12-23

### HuaweiCloud SDK Core

- _新增特性_
    - 支持Region管理，用户新建客户端时可以通过{Service}Region传入，无需自己配置endpoint。配置Region后，支持自动回填ProjectId/DomainId
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口：查询用量单位列表
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CES

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - ShowMetricData接口字段更新

### HuaweiCloud SDK Live

- _新增特性_
    - 新增支持接口：查询播放画像信息
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK MPC

- _新增特性_
    - 新增支持接口：视频增强模板相关接口，声道合并任务相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK RDS

- _新增特性_
    - 支持云数据库服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK SMN

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 消息模板名称中字段类型调整

# 0.0.25-beta 2020-12-15

### HuaweiCloud SDK CCE

- _新增特性_
    - 支持云容器引擎服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ELB

- _新增特性_
    - 无
- _解决问题_
    - 创建监听器接口返回为空问题修复
    - 证书列表查询接口返回非列表问题修复
- _特性变更_
    - 无

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 新增支持接口：更新函数预留实例个数
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK NAT

- _新增特性_
    - 无
- _解决问题_
    - 修复批量创建DNAT规则失败的问题
- _特性变更_
    - 无

# 0.0.24-beta 2020-12-04

### HuaweiCloud SDK SMN

- _新增特性_
    - 支持消息通知服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.23-beta 2020-11-30

### HuaweiCloud SDK BCS

- _新增特性_
    - 支持区块链服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BMS

- _新增特性_
    - 支持裸金属服务器
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 新增支持接口：查询使用量列表，设置/取消包周期资源到期转按需
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 新增支持接口：迁移资源接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CES

- _新增特性_
    - 新增支持接口：
        - 查询事件监控列表
        - 查询某一个事件监控详情
        - 创建资源分组
        - 更新资源分组
        - 删除资源分组
        - 查询所有资源分组
        - 修改告警规则
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK DCS

- _新增特性_
    - 支持分布式缓存服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - [Issue 21](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/21) 开放查询SSH密钥详情接口

### HuaweiCloud SDK FunctionGraph

- _新增特性_
    - 支持函数工作流服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 新增支持接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - Live服务客户端名字修正：LiveAPIClient → LiveClient

# 0.0.22-beta 2020-11-17

### HuaweiCloud SDK AS

- _新增特性_
    - 无
- _解决问题_
    - [Issue 8](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/8) Fix the problem that creating scaling
      policy failed.
- _特性变更_
    - 无

### HuaweiCloud SDK DMS

- _新增特性_
    - 支持分布式消息服务，提供Kafka专享版和RabbitMQ专享版
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建虚拟机接口（按需和包周期）增加 `dry_run` 属性，表示是否预检此次请求

### HuaweiCloud SDK NAT

- _新增特性_
    - 支持NAT网关服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK VPC

- _新增特性_
    - 支持网络ACL相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.21-beta 2020-11-11

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 同步[Pull requests #11](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/pull/11)修改
- _特性变更_
    - 无

# HuaweiCloud SDK CBR

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建存储库接口(CreateVault)新增存储库turbo类型
    - 修改策略接口(UpdatePolicy)删除多余字段

### HuaweiCloud SDK CES

- _新增特性_
    - 新增接口响应示例，调整字段描述
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 生成客户端文件的名字调整：devcloudpipeline_client → cloudPipeline_client
    - 生成Meta文件的名字调整：devcloudpipeline_meta → cloudPipeline_meta

### HuaweiCloud SDK DevStar

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 修改接口参数，调整示例代码

# 0.0.20-beta 2020-11-02

### HuaweiCloud SDK CES

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 创建告警规则接口增加资源分组字段
    - 查询告警历史接口响应体metadata修改，只返回total字段

### HuaweiCloud SDK ELB

- _新增特性_
    - 无
- _解决问题_
    - 创建转发规则接口参数名错误问题修复
- _特性变更_
    - 无

# 0.0.19-beta 2020-10-31

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - Fix: query参数中包含枚举变量时请求失败
    - [Issue 7](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/7) 解决json.Marshal变成{}对象的问题
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 新增支持接口：TAG标签相关接口
- _解决问题_
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/17) 修复ListBackups接口响应体问题
- _特性变更_
    - 无

### HuaweiCloud SDK CTS

- _新增特性_
    - 新增支持接口：ListQuotas（查询租户追踪器配额信息）
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EPS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整，原有的`*EP`接口展开为`*EnterpriseProject`，涉及接口：
        1. ListEP → ListEnterpriseProject
        2. CreateEP → CreateEnterpriseProject
        3. ShowEP → ShowEnterpriseProject
        4. ModifyEP → ModifyEnterpriseProject
        5. EnableEP → EnableEnterpriseProject
        6. ShowEPQuota → ShowEnterpriseProjectQuota
        7. ShowResourceBindEP → ShowResourceBindEnterpriseProject
        8. DisableEP → DisableEnterpriseProject

### HuaweiCloud SDK Iam

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口名称调整，涉及接口：
        1. KeystoneCreateUserTokenByPasswordAndMFA → KeystoneCreateUserTokenByPasswordAndMfa
        2. CreateUnscopeTokenByIDPInitiated → CreateUnscopeTokenByIdpInitiated

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 新增支持接口：迭代信息、用户信息、项目成员、项目信息、项目指标、项目统计等相关方法的接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.18-beta 2020-10-20

### HuaweiCloud SDK ELB

- _新增特性_
    - 增加v2版本接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.17-beta 2020-10-14

### HuaweiCloud SDK BSS

- _新增特性_
    - 伙伴中心支持导出产品目录价
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 新增直播V2接口
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.16-beta 2020-10-12

### HuaweiCloud SDK CTS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 删除v1版本废弃接口

### HuaweiCloud SDK DevStar

- _新增特性_
    - 无
- _解决问题_
    - 服务客户端认证类型调整为全局认证，即GlobalCredentials
- _特性变更_
    - 无

# 0.0.15-beta 2020-09-30

### HuaweiCloud SDK AS

- _新增特性_
    - 支持弹性伸缩服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.14-beta 2020-09-24

### HuaweiCloud SDK BSS

- _新增特性_
    - 无
- _解决问题_
    - 修复BssClient类无法加载的问题
- _特性变更_
    - 无

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 接口ListPublicips调整，企业项目ID不支持多值查询

# 0.0.13-beta 2020-09-18

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 修正接口参数类型
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 运营能力开放接口更新

### HuaweiCloud SDK EIP

- _新增特性_
    - 无
- _解决问题_
    - 修复查询弹性公网IP不支持传入多个值的问题
- _特性变更_
    - 无

### HuaweiCloud SDK ELB

- _新增特性_
    - 无
- _解决问题_
    - 修复部分接口参数错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK IMS

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 调整接口描述

### HuaweiCloud SDK Live

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 修改禁播参数描述

# 0.0.12.1-beta 2020-09-16

### HuaweiCloud SDK ECS

- _新增特性_
    - 无
- _解决问题_
    - 修复接口参数类型错误
- _特性变更_
    - 无

### HuaweiCloud SDK All

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 可选参数类型为list时，参数类型变更为指针，例如：[]string 将变更为 *[]string

# 0.0.12-beta 2020-09-11

### HuaweiCloud SDK KPS

- _新增特性_
    - 支持KPS服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EVS

- _新增特性_
    - 支持EVS服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 无
- _解决问题_
    - 修复返回值类型定义错误的问题
- _特性变更_
    - 无

# 0.0.11-beta 2020-09-09

### HuaweiCloud SDK CBR

- _新增特性_
    - 无
- _解决问题_
    - 修复资源相关接口类型错误的问题
- _特性变更_
    - 无

### HuaweiCloud SDK VPC

- _新增特性_
    - 无
- _解决问题_
    - 修复安全组相关类型错误的问题
- _特性变更_
    - 无

# 0.0.10-beta 2020-09-04

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 修复yaml中没有定义format的整型枚举参数无法生成枚举代码的问题
- _特性变更_
    - 调整Http请求头的User-Agent信息

# 0.0.9-beta 2020-08-28

### HuaweiCloud SDK CloudPipeline

- _新增特性_
    - 支持流水线服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EIP

- _新增特性_
    - 新增支持接口：弹性公网IP标签相关接口和共享带宽相关接口
- _解决问题_
    - 批量创建带宽接口，修改billing_info字段
- _特性变更_
    - 无

### HuaweiCloud SDK IAM

- _新增特性_
    - 新增支持接口：console一致性相关接口
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ProjectMan

- _新增特性_
    - 支持项目管理服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK VPC

- _新增特性_
    - 新增支持接口：安全组、安全组规则和可用IP数
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.8-beta 2020-08-25

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 无
- _特性变更_
    - 可选枚举变量类型变更为指针类型。

### HuaweiCloud SDK ELB

- _新增特性_
    - 支持弹性负载均衡服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.7-beta 2020-08-20

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决部分枚举类型变量无法正常读取的问题。
- _特性变更_
    - 以_开头的枚举类型变量名称添加前缀 'E'。

# 0.0.6-beta 2020-08-20

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决当结构体包含枚举类型变量场景下部分依赖丢失的问题。
- _特性变更_
    - 无

# 0.0.5-beta 2020-08-19

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - 解决枚举类型反序列化失败的问题。
    - 解决部分场景下云服务异常响应反序列化失败问题。
- _特性变更_
    - 无

# 0.0.4-beta 2020-08-18

### HuaweiCloud SDK Core

- _新增特性_
    - 无
- _解决问题_
    - Go 基础类型默认值序列化丢失问题修复
- _特性变更_
    - 无

# 0.0.3-beta 2020-08-14

### HuaweiCloud SDK APIG

- _新增特性_
    - 支持API网关
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK BSS

- _新增特性_
    - 支持能力开放服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 0.0.2-beta 2020-08-11

### HuaweiCloud SDK Core

- _新增特性_
    - 支持临时AK/SK认证模式
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CBR

- _新增特性_
    - 支持云备份服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK CloudIDE

- _新增特性_
    - 支持云测服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK ECS

- _新增特性_
    - 支持弹性云服务器服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EIP

- _新增特性_
    - 支持弹性公网IP服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK EVS

- _新增特性_
    - 支持云硬盘服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK IMS

- _新增特性_
    - 支持镜像服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK Live

- _新增特性_
    - 支持视频直播服务
- _解决问题_
    - 无
- _特性变更_
    - 无

### HuaweiCloud SDK MPC

- _新增特性_
    - 支持媒体转码服务
- _解决问题_
    - 无
- _特性变更_
    - 无

# 3.0.1-beta 2020-07-30

### 首次发布

- _已支持服务_
    - Classroom
    - 云审计服务（CTS）
    - 模板引擎（DevStar）
    - 企业管理服务（EPS）
    - 统一身份认证服务（IAM）
    - 标签管理服务（TMS）
    - 虚拟私有云（VPC）
