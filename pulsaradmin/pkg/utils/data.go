// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package utils

// ClusterData information on a cluster
type ClusterData struct {
	Name                           string   `json:"-"`
	ServiceURL                     string   `json:"serviceUrl"`
	ServiceURLTls                  string   `json:"serviceUrlTls"`
	BrokerServiceURL               string   `json:"brokerServiceUrl"`
	BrokerServiceURLTls            string   `json:"brokerServiceUrlTls"`
	PeerClusterNames               []string `json:"peerClusterNames"`
	AuthenticationPlugin           string   `json:"authenticationPlugin"`
	AuthenticationParameters       string   `json:"authenticationParameters"`
	BrokerClientTrustCertsFilePath string   `json:"brokerClientTrustCertsFilePath"`
	BrokerClientTLSEnabled         bool     `json:"brokerClientTlsEnabled"`
}

// FunctionData information for a Pulsar Function
type FunctionData struct {
	UpdateAuthData       bool `json:"updateAuthData"`
	RetainOrdering       bool `json:"retainOrdering"`
	Watch                bool `json:"watch"`
	AutoAck              bool `json:"autoAck"`
	Parallelism          int  `json:"parallelism"`
	WindowLengthCount    int  `json:"windowLengthCount"`
	SlidingIntervalCount int  `json:"slidingIntervalCount"`
	MaxMessageRetries    int  `json:"maxMessageRetries"`

	TimeoutMs                 int64   `json:"timeoutMs"`
	SlidingIntervalDurationMs int64   `json:"slidingIntervalDurationMs"`
	WindowLengthDurationMs    int64   `json:"windowLengthDurationMs"`
	RAM                       int64   `json:"ram"`
	Disk                      int64   `json:"disk"`
	CPU                       float64 `json:"cpu"`
	SubsName                  string  `json:"subsName"`
	DeadLetterTopic           string  `json:"deadLetterTopic"`
	Key                       string  `json:"key"`
	State                     string  `json:"state"`
	TriggerValue              string  `json:"triggerValue"`
	TriggerFile               string  `json:"triggerFile"`
	Topic                     string  `json:"topic"`

	UserCodeFile                 string          `json:"-"`
	FQFN                         string          `json:"fqfn"`
	Tenant                       string          `json:"tenant"`
	Namespace                    string          `json:"namespace"`
	FuncName                     string          `json:"functionName"`
	InstanceID                   string          `json:"instance_id"`
	ClassName                    string          `json:"className"`
	FunctionType                 string          `json:"functionType"`
	CleanupSubscription          bool            `json:"cleanupSubscription"`
	Jar                          string          `json:"jarFile"`
	Py                           string          `json:"pyFile"`
	Go                           string          `json:"goFile"`
	Inputs                       string          `json:"inputs"`
	TopicsPattern                string          `json:"topicsPattern"`
	Output                       string          `json:"output"`
	ProducerConfig               string          `json:"producerConfig"`
	LogTopic                     string          `json:"logTopic"`
	SchemaType                   string          `json:"schemaType"`
	CustomSerDeInputs            string          `json:"customSerdeInputString"`
	CustomSchemaInput            string          `json:"customSchemaInputString"`
	CustomSchemaOutput           string          `json:"customSchemaOutputString"`
	InputSpecs                   string          `json:"inputSpecs"`
	InputTypeClassName           string          `json:"inputTypeClassName"`
	OutputSerDeClassName         string          `json:"outputSerdeClassName"`
	OutputTypeClassName          string          `json:"outputTypeClassName"`
	FunctionConfigFile           string          `json:"fnConfigFile"`
	ProcessingGuarantees         string          `json:"processingGuarantees"`
	UserConfig                   string          `json:"userConfigString"`
	RetainKeyOrdering            bool            `json:"retainKeyOrdering"`
	BatchBuilder                 string          `json:"batchBuilder"`
	ForwardSourceMessageProperty bool            `json:"forwardSourceMessageProperty"`
	SubsPosition                 string          `json:"subsPosition"`
	SkipToLatest                 bool            `json:"skipToLatest"`
	CustomRuntimeOptions         string          `json:"customRuntimeOptions"`
	Secrets                      string          `json:"secretsString"`
	DestinationFile              string          `json:"destinationFile"`
	Path                         string          `json:"path"`
	RuntimeFlags                 string          `json:"runtimeFlags,omitempty"`
	FuncConf                     *FunctionConfig `json:"-"`
}

// Failure Domain information
type FailureDomainData struct {
	ClusterName string   `json:"-"`
	DomainName  string   `json:"-"`
	BrokerList  []string `json:"brokers"`
}

type FailureDomainMap map[string]FailureDomainData

// Tenant args
type TenantData struct {
	Name            string   `json:"-"`
	AdminRoles      []string `json:"adminRoles"`
	AllowedClusters []string `json:"allowedClusters"`
}

type SourceData struct {
	Tenant                   string  `json:"tenant,omitempty"`
	Namespace                string  `json:"namespace,omitempty"`
	Name                     string  `json:"name,omitempty"`
	SourceType               string  `json:"sourceType,omitempty"`
	ProcessingGuarantees     string  `json:"processingGuarantees,omitempty"`
	DestinationTopicName     string  `json:"destinationTopicName,omitempty"`
	ProducerConfig           string  `json:"producerConfig,omitempty"`
	BatchBuilder             string  `json:"batchBuilder,omitempty"`
	DeserializationClassName string  `json:"deserializationClassName,omitempty"`
	SchemaType               string  `json:"schemaType,omitempty"`
	Parallelism              int     `json:"parallelism,omitempty"`
	Archive                  string  `json:"archive,omitempty"`
	ClassName                string  `json:"className,omitempty"`
	SourceConfigFile         string  `json:"sourceConfigFile,omitempty"`
	CPU                      float64 `json:"cpu,omitempty"`
	RAM                      int64   `json:"ram,omitempty"`
	Disk                     int64   `json:"disk,omitempty"`
	SourceConfigString       string  `json:"sourceConfigString,omitempty"`
	BatchSourceConfigString  string  `json:"batchSourceConfigString,omitempty"`
	CustomRuntimeOptions     string  `json:"customRuntimeOptions,omitempty"`
	Secrets                  string  `json:"secretsString,omitempty"`

	SourceConf *SourceConfig `json:"-,omitempty"`
	InstanceID string        `json:"instanceId,omitempty"`

	UpdateAuthData bool   `json:"updateAuthData,omitempty"`
	RuntimeFlags   string `json:"runtimeFlags,omitempty"`
}

type SinkData struct {
	UpdateAuthData               bool        `json:"updateAuthData,omitempty"`
	RetainOrdering               bool        `json:"retainOrdering,omitempty"`
	AutoAck                      bool        `json:"autoAck,omitempty"`
	Parallelism                  int         `json:"parallelism,omitempty"`
	RAM                          int64       `json:"ram,omitempty"`
	Disk                         int64       `json:"disk,omitempty"`
	TimeoutMs                    int64       `json:"timeoutMs,omitempty"`
	CPU                          float64     `json:"cpu,omitempty"`
	Tenant                       string      `json:"tenant,omitempty"`
	Namespace                    string      `json:"namespace,omitempty"`
	Name                         string      `json:"name,omitempty"`
	SinkType                     string      `json:"sinkType,omitempty"`
	CleanupSubscription          bool        `json:"cleanupSubscription"`
	Inputs                       string      `json:"inputs,omitempty"`
	TopicsPattern                string      `json:"topicsPattern,omitempty"`
	SubsName                     string      `json:"subsName,omitempty"`
	SubsPosition                 string      `json:"subsPosition,omitempty"`
	CustomSerdeInputString       string      `json:"customSerdeInputString,omitempty"`
	CustomSchemaInputString      string      `json:"customSchemaInputString,omitempty"`
	InputSpecs                   string      `json:"inputSpecs,omitempty"`
	MaxMessageRetries            int         `json:"maxMessageRetries,omitempty"`
	DeadLetterTopic              string      `json:"deadLetterTopic,omitempty"`
	ProcessingGuarantees         string      `json:"processingGuarantees,omitempty"`
	RetainKeyOrdering            bool        `json:"retainKeyOrdering,omitempty"`
	Archive                      string      `json:"archive,omitempty"`
	ClassName                    string      `json:"className,omitempty"`
	SinkConfigFile               string      `json:"sinkConfigFile,omitempty"`
	SinkConfigString             string      `json:"sinkConfigString,omitempty"`
	NegativeAckRedeliveryDelayMs int64       `json:"negativeAckRedeliveryDelayMs,omitempty"`
	CustomRuntimeOptions         string      `json:"customRuntimeOptions,omitempty"`
	Secrets                      string      `json:"secretsString,omitempty"`
	InstanceID                   string      `json:"instanceId,omitempty"`
	TransformFunction            string      `json:"transformFunction,omitempty"`
	TransformFunctionClassName   string      `json:"transformFunctionClassName,omitempty"`
	TransformFunctionConfig      string      `json:"transformFunctionConfig,omitempty"`
	RuntimeFlags                 string      `json:"runtimeFlags,omitempty"`
	SinkConf                     *SinkConfig `json:"-,omitempty"`
}

// Topic data
type PartitionedTopicMetadata struct {
	Partitions int `json:"partitions"`
}

type ManagedLedgerInfoLedgerInfo struct {
	LedgerID             int64  `json:"ledgerId"`
	Entries              int64  `json:"entries"`
	Size                 int64  `json:"size"`
	Timestamp            int64  `json:"timestamp"`
	Offloaded            bool   `json:"isOffloaded"`
	OffloadedContextUUID string `json:"offloadedContextUuid"`
}

type ManagedLedgerInfo struct {
	Version            int                           `json:"version"`
	CreationDate       string                        `json:"creationDate"`
	ModificationData   string                        `json:"modificationData"`
	Ledgers            []ManagedLedgerInfoLedgerInfo `json:"ledgers"`
	TerminatedPosition PositionInfo                  `json:"terminatedPosition"`
	Cursors            map[string]CursorInfo         `json:"cursors"`
}

type NamespacesData struct {
	Enable                         bool     `json:"enable"`
	Unload                         bool     `json:"unload"`
	NumBundles                     int      `json:"numBundles"`
	BookkeeperEnsemble             int      `json:"bookkeeperEnsemble"`
	BookkeeperWriteQuorum          int      `json:"bookkeeperWriteQuorum"`
	MessageTTL                     int      `json:"messageTTL"`
	BookkeeperAckQuorum            int      `json:"bookkeeperAckQuorum"`
	ManagedLedgerMaxMarkDeleteRate float64  `json:"managedLedgerMaxMarkDeleteRate"`
	ClusterIDs                     string   `json:"clusterIds"`
	RetentionTimeStr               string   `json:"retentionTimeStr"`
	LimitStr                       string   `json:"limitStr"`
	LimitTime                      int64    `json:"limitTime"`
	PolicyStr                      string   `json:"policyStr"`
	BacklogQuotaType               string   `json:"backlogQuotaType"`
	AntiAffinityGroup              string   `json:"antiAffinityGroup"`
	Tenant                         string   `json:"tenant"`
	Cluster                        string   `json:"cluster"`
	Bundle                         string   `json:"bundle"`
	Clusters                       []string `json:"clusters"`
}

type TopicStats struct {
	BacklogSize         int64                        `json:"backlogSize"`
	MsgCounterIn        int64                        `json:"msgInCounter"`
	MsgCounterOut       int64                        `json:"msgOutCounter"`
	MsgRateIn           float64                      `json:"msgRateIn"`
	MsgRateOut          float64                      `json:"msgRateOut"`
	MsgThroughputIn     float64                      `json:"msgThroughputIn"`
	MsgThroughputOut    float64                      `json:"msgThroughputOut"`
	AverageMsgSize      float64                      `json:"averageMsgSize"`
	StorageSize         int64                        `json:"storageSize"`
	Publishers          []PublisherStats             `json:"publishers"`
	Subscriptions       map[string]SubscriptionStats `json:"subscriptions"`
	Replication         map[string]ReplicatorStats   `json:"replication"`
	DeDuplicationStatus string                       `json:"deduplicationStatus"`
}

type ProducerAccessMode string

const (
	ProduceModeShared               ProducerAccessMode = "Shared"
	ProduceModeExclusive                               = "Exclusive"
	ProduceModeExclusiveWithFencing                    = "ExclusiveWithFencing"
	ProduceModeWaitForExclusive                        = "WaitForExclusive"
)

type PublisherStats struct {
	AccessModel               ProducerAccessMode `json:"accessMode"`
	ProducerID                int64              `json:"producerId"`
	MsgRateIn                 float64            `json:"msgRateIn"`
	MsgThroughputIn           float64            `json:"msgThroughputIn"`
	AverageMsgSize            float64            `json:"averageMsgSize"`
	ChunkedMessageRate        float64            `json:"chunkedMessageRate"`
	IsSupportsPartialProducer bool               `json:"supportsPartialProducer"`
	ProducerName              string             `json:"producerName"`
	Address                   string             `json:"address"`
	ConnectedSince            string             `json:"connectedSince"`
	ClientVersion             string             `json:"clientVersion"`
	Metadata                  map[string]string  `json:"metadata"`
}

type SubscriptionStats struct {
	BlockedSubscriptionOnUnackedMsgs          bool              `json:"blockedSubscriptionOnUnackedMsgs"`
	IsReplicated                              bool              `json:"isReplicated"`
	LastConsumedFlowTimestamp                 int64             `json:"lastConsumedFlowTimestamp"`
	LastConsumedTimestamp                     int64             `json:"lastConsumedTimestamp"`
	LastAckedTimestamp                        int64             `json:"lastAckedTimestamp"`
	MsgRateOut                                float64           `json:"msgRateOut"`
	MsgThroughputOut                          float64           `json:"msgThroughputOut"`
	MsgRateRedeliver                          float64           `json:"msgRateRedeliver"`
	MsgRateExpired                            float64           `json:"msgRateExpired"`
	MsgBacklog                                int64             `json:"msgBacklog"`
	MsgBacklogNoDelayed                       int64             `json:"msgBacklogNoDelayed"`
	MsgDelayed                                int64             `json:"msgDelayed"`
	UnAckedMessages                           int64             `json:"unackedMessages"`
	SubType                                   string            `json:"type"`
	ActiveConsumerName                        string            `json:"activeConsumerName"`
	BytesOutCounter                           int64             `json:"bytesOutCounter"`
	MsgOutCounter                             int64             `json:"msgOutCounter"`
	MessageAckRate                            float64           `json:"messageAckRate"`
	ChunkedMessageRate                        float64           `json:"chunkedMessageRate"`
	BacklogSize                               int64             `json:"backlogSize"`
	EarliestMsgPublishTimeInBacklog           int64             `json:"earliestMsgPublishTimeInBacklog"`
	TotalMsgExpired                           int64             `json:"totalMsgExpired"`
	LastExpireTimestamp                       int64             `json:"lastExpireTimestamp"`
	LastMarkDeleteAdvancedTimestamp           int64             `json:"lastMarkDeleteAdvancedTimestamp"`
	Consumers                                 []ConsumerStats   `json:"consumers"`
	IsDurable                                 bool              `json:"isDurable"`
	AllowOutOfOrderDelivery                   bool              `json:"allowOutOfOrderDelivery"`
	ConsumersAfterMarkDeletePosition          map[string]string `json:"consumersAfterMarkDeletePosition"`
	NonContiguousDeletedMessagesRanges        int               `json:"nonContiguousDeletedMessagesRanges"`
	NonContiguousDeletedMessagesRangesSrzSize int               `json:"nonContiguousDeletedMessagesRangesSerializedSize"`
	DelayedMessageIndexSizeInBytes            int64             `json:"delayedMessageIndexSizeInBytes"`
	SubscriptionProperties                    map[string]string `json:"subscriptionProperties"`
	FilterProcessedMsgCount                   int64             `json:"filterProcessedMsgCount"`
	FilterAcceptedMsgCount                    int64             `json:"filterAcceptedMsgCount"`
	FilterRejectedMsgCount                    int64             `json:"filterRejectedMsgCount"`
	FilterRescheduledMsgCount                 int64             `json:"filterRescheduledMsgCount"`
}

type ConsumerStats struct {
	BlockedConsumerOnUnAckedMsgs bool              `json:"blockedConsumerOnUnackedMsgs"`
	AvailablePermits             int               `json:"availablePermits"`
	UnAckedMessages              int               `json:"unackedMessages"`
	MsgRateOut                   float64           `json:"msgRateOut"`
	MsgThroughputOut             float64           `json:"msgThroughputOut"`
	MsgRateRedeliver             float64           `json:"msgRateRedeliver"`
	ConsumerName                 string            `json:"consumerName"`
	BytesOutCounter              int64             `json:"bytesOutCounter"`
	MsgOutCounter                int64             `json:"msgOutCounter"`
	MessageAckRate               float64           `json:"messageAckRate"`
	ChunkedMessageRate           float64           `json:"chunkedMessageRate"`
	AvgMessagesPerEntry          int               `json:"avgMessagesPerEntry"`
	Address                      string            `json:"address"`
	ConnectedSince               string            `json:"connectedSince"`
	ClientVersion                string            `json:"clientVersion"`
	LastAckedTimestamp           int64             `json:"lastAckedTimestamp"`
	LastConsumedTimestamp        int64             `json:"lastConsumedTimestamp"`
	LastConsumedFlowTimestamp    int64             `json:"lastConsumedFlowTimestamp"`
	Metadata                     map[string]string `json:"metadata"`
}

type ReplicatorStats struct {
	Connected                 bool    `json:"connected"`
	MsgRateIn                 float64 `json:"msgRateIn"`
	MsgRateOut                float64 `json:"msgRateOut"`
	MsgThroughputIn           float64 `json:"msgThroughputIn"`
	MsgThroughputOut          float64 `json:"msgThroughputOut"`
	MsgRateExpired            float64 `json:"msgRateExpired"`
	ReplicationBacklog        int64   `json:"replicationBacklog"`
	ReplicationDelayInSeconds int64   `json:"replicationDelayInSeconds"`
	InboundConnection         string  `json:"inboundConnection"`
	InboundConnectedSince     string  `json:"inboundConnectedSince"`
	OutboundConnection        string  `json:"outboundConnection"`
	OutboundConnectedSince    string  `json:"outboundConnectedSince"`
}

type PersistentTopicInternalStats struct {
	WaitingCursorsCount                int                    `json:"waitingCursorsCount"`
	PendingAddEntriesCount             int                    `json:"pendingAddEntriesCount"`
	EntriesAddedCounter                int64                  `json:"entriesAddedCounter"`
	NumberOfEntries                    int64                  `json:"numberOfEntries"`
	TotalSize                          int64                  `json:"totalSize"`
	CurrentLedgerEntries               int64                  `json:"currentLedgerEntries"`
	CurrentLedgerSize                  int64                  `json:"currentLedgerSize"`
	LastLedgerCreatedTimestamp         string                 `json:"lastLedgerCreatedTimestamp"`
	LastLedgerCreationFailureTimestamp string                 `json:"lastLedgerCreationFailureTimestamp"`
	LastConfirmedEntry                 string                 `json:"lastConfirmedEntry"`
	State                              string                 `json:"state"`
	Ledgers                            []LedgerInfo           `json:"ledgers"`
	Cursors                            map[string]CursorStats `json:"cursors"`
	SchemaLedgers                      []SchemaLedger         `json:"schemaLedgers"`
	CompactedLedger                    CompactedLedger        `json:"compactedLedger"`
}

type LedgerInfo struct {
	LedgerID        int64  `json:"ledgerId"`
	Entries         int64  `json:"entries"`
	Size            int64  `json:"size"`
	Timestamp       int64  `json:"timestamp"`
	Offloaded       bool   `json:"offloaded"`
	MetaData        string `json:"metadata"`
	UnderReplicated bool   `json:"underReplicated"`
}

type CursorInfo struct {
	Version                   int                `json:"version"`
	CreationDate              string             `json:"creationDate"`
	ModificationDate          string             `json:"modificationDate"`
	CursorsLedgerID           int64              `json:"cursorsLedgerId"`
	MarkDelete                PositionInfo       `json:"markDelete"`
	IndividualDeletedMessages []MessageRangeInfo `json:"individualDeletedMessages"`
	Properties                map[string]int64
}

type PositionInfo struct {
	LedgerID int64 `json:"ledgerId"`
	EntryID  int64 `json:"entryId"`
}

type MessageRangeInfo struct {
	From      PositionInfo `json:"from"`
	To        PositionInfo `json:"to"`
	Offloaded bool         `json:"offloaded"`
}

type CursorStats struct {
	MarkDeletePosition                       string           `json:"markDeletePosition"`
	ReadPosition                             string           `json:"readPosition"`
	WaitingReadOp                            bool             `json:"waitingReadOp"`
	PendingReadOps                           int              `json:"pendingReadOps"`
	MessagesConsumedCounter                  int64            `json:"messagesConsumedCounter"`
	CursorLedger                             int64            `json:"cursorLedger"`
	CursorLedgerLastEntry                    int64            `json:"cursorLedgerLastEntry"`
	IndividuallyDeletedMessages              string           `json:"individuallyDeletedMessages"`
	LastLedgerWitchTimestamp                 string           `json:"lastLedgerWitchTimestamp"`
	State                                    string           `json:"state"`
	NumberOfEntriesSinceFirstNotAckedMessage int64            `json:"numberOfEntriesSinceFirstNotAckedMessage"`
	TotalNonContiguousDeletedMessagesRange   int              `json:"totalNonContiguousDeletedMessagesRange"`
	Properties                               map[string]int64 `json:"properties"`
}

type PartitionedTopicStats struct {
	MsgRateIn           float64                      `json:"msgRateIn"`
	MsgRateOut          float64                      `json:"msgRateOut"`
	MsgThroughputIn     float64                      `json:"msgThroughputIn"`
	MsgThroughputOut    float64                      `json:"msgThroughputOut"`
	AverageMsgSize      float64                      `json:"averageMsgSize"`
	StorageSize         int64                        `json:"storageSize"`
	Publishers          []PublisherStats             `json:"publishers"`
	Subscriptions       map[string]SubscriptionStats `json:"subscriptions"`
	Replication         map[string]ReplicatorStats   `json:"replication"`
	DeDuplicationStatus string                       `json:"deduplicationStatus"`
	Metadata            PartitionedTopicMetadata     `json:"metadata"`
	Partitions          map[string]TopicStats        `json:"partitions"`
}

type SchemaData struct {
	Version         int64  `json:"version"`
	Filename        string `json:"filename"`
	Jar             string `json:"jar"`
	Type            string `json:"type"`
	Classname       string `json:"classname"`
	AlwaysAllowNull bool   `json:"alwaysAllowNull"`
	DryRun          bool   `json:"dryRun"`
}

type LookupData struct {
	BrokerURL    string `json:"brokerUrl"`
	BrokerURLTLS string `json:"brokerUrlTls"`
	HTTPURL      string `json:"httpUrl"`
	HTTPURLTLS   string `json:"httpUrlTls"`
}

type NsIsolationPoliciesData struct {
	Namespaces                 []string `json:"namespaces"`
	Primary                    []string `json:"primary"`
	Secondary                  []string `json:"secondary"`
	AutoFailoverPolicyTypeName string   `json:"autoFailoverPolicyTypeName"`
	AutoFailoverPolicyParams   string   `json:"autoFailoverPolicyParams"`
}

type BrokerData struct {
	URL         string `json:"brokerUrl"`
	ConfigName  string `json:"configName"`
	ConfigValue string `json:"configValue"`
}

type BrokerStatsData struct {
	Indent bool `json:"indent"`
}

type ResourceQuotaData struct {
	Names        string `json:"names"`
	Bundle       string `json:"bundle"`
	MsgRateIn    int64  `json:"msgRateIn"`
	MsgRateOut   int64  `json:"msgRateOut"`
	BandwidthIn  int64  `json:"bandwidthIn"`
	BandwidthOut int64  `json:"bandwidthOut"`
	Memory       int64  `json:"memory"`
	Dynamic      bool   `json:"dynamic"`
}

type PersistenceData struct {
	BookkeeperEnsemble             int64   `json:"bookkeeperEnsemble"`
	BookkeeperWriteQuorum          int64   `json:"bookkeeperWriteQuorum"`
	BookkeeperAckQuorum            int64   `json:"bookkeeperAckQuorum"`
	ManagedLedgerMaxMarkDeleteRate float64 `json:"managedLedgerMaxMarkDeleteRate"`
}

type DelayedDeliveryCmdData struct {
	Enable                 bool   `json:"enable"`
	Disable                bool   `json:"disable"`
	DelayedDeliveryTimeStr string `json:"delayedDeliveryTimeStr"`
}

type DelayedDeliveryData struct {
	TickTime float64 `json:"tickTime"`
	Active   bool    `json:"active"`
	// MaxDelayInMillis is optional and was added for enhanced delayed delivery support
	// Default value 0 means no maximum delay limit (backward compatible)
	MaxDelayInMillis int64 `json:"maxDelayInMillis,omitempty"`
}

// NewDelayedDeliveryData creates a DelayedDeliveryData with backward compatible defaults
func NewDelayedDeliveryData(tickTime float64, active bool) *DelayedDeliveryData {
	return &DelayedDeliveryData{
		TickTime: tickTime,
		Active:   active,
		// MaxDelayInMillis is left as 0 (no limit) for backward compatibility
	}
}

// NewDelayedDeliveryDataWithMaxDelay creates a DelayedDeliveryData with max delay limit
func NewDelayedDeliveryDataWithMaxDelay(tickTime float64, active bool, maxDelayMs int64) *DelayedDeliveryData {
	return &DelayedDeliveryData{
		TickTime:         tickTime,
		Active:           active,
		MaxDelayInMillis: maxDelayMs,
	}
}

type DispatchRateData struct {
	DispatchThrottlingRateInMsg  int64 `json:"dispatchThrottlingRateInMsg"`
	DispatchThrottlingRateInByte int64 `json:"dispatchThrottlingRateInByte"`
	RatePeriodInSecond           int64 `json:"ratePeriodInSecond"`
	RelativeToPublishRate        bool  `json:"relativeToPublishRate"`
}

type PublishRateData struct {
	PublishThrottlingRateInMsg  int64 `json:"publishThrottlingRateInMsg"`
	PublishThrottlingRateInByte int64 `json:"publishThrottlingRateInByte"`
}

type SchemaLedger struct {
	LedgerID    int64 `json:"ledgerId"`
	Entries     int64 `json:"entries"`
	Size        int64 `json:"size"`
	Timestamp   int64 `json:"timestamp"`
	IsOffloaded bool  `json:"isOffloaded"`
}

type CompactedLedger struct {
	LedgerID        int64 `json:"ledgerId"`
	Entries         int64 `json:"entries"`
	Size            int64 `json:"size"`
	Offloaded       bool  `json:"offloaded"`
	UnderReplicated bool  `json:"underReplicated"`
}

type GetStatsOptions struct {
	GetPreciseBacklog        bool `json:"get_precise_backlog"`
	SubscriptionBacklogSize  bool `json:"subscription_backlog_size"`
	GetEarliestTimeInBacklog bool `json:"get_earliest_time_in_backlog"`
	ExcludePublishers        bool `json:"exclude_publishers"`
	ExcludeConsumers         bool `json:"exclude_consumers"`
}

type BrokerInfo struct {
	BrokerID   string `json:"brokerId"`
	ServiceURL string `json:"serviceUrl"`
}

type TopicVersion string

const (
	TopicVersionV1 TopicVersion = "V1"
	TopicVersionV2 TopicVersion = "V2"
)

func (t TopicVersion) String() string {
	return string(t)
}
