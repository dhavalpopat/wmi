// Copyright 2019 (c) Microsoft Corporation.
// Licensed under the MIT license.

//
// Author:
//      Auto Generated on 6/13/2024 using wmigen
//      Source Eventlog_Type
//////////////////////////////////////////////
package wmi

// Eventlog_Type
type Eventlog_Type int

const (
	// ISCSI_ERR_TDI_CONNECT_FAILED enum
	Eventlog_Type_ISCSI_ERR_TDI_CONNECT_FAILED Eventlog_Type = 1
	// ISCSI_ERR_INSUFFICIENT_SESSION_RESOURCES enum
	Eventlog_Type_ISCSI_ERR_INSUFFICIENT_SESSION_RESOURCES Eventlog_Type = 2
	// ISCSI_ERR_INVALID_COMMAND_SEQUENCE_NUMBER enum
	Eventlog_Type_ISCSI_ERR_INVALID_COMMAND_SEQUENCE_NUMBER Eventlog_Type = 3
	// ISCSI_ERR_INVALID_BURST_LENGTH enum
	Eventlog_Type_ISCSI_ERR_INVALID_BURST_LENGTH Eventlog_Type = 4
	// ISCSI_ERR_SETUP_NETWORK_NODE enum
	Eventlog_Type_ISCSI_ERR_SETUP_NETWORK_NODE Eventlog_Type = 5
	// ISCSI_ERR_INSUFFICIENT_CONNECTION_RESOURCES enum
	Eventlog_Type_ISCSI_ERR_INSUFFICIENT_CONNECTION_RESOURCES Eventlog_Type = 6
	// ISCSI_ERR_SEND_FAILED enum
	Eventlog_Type_ISCSI_ERR_SEND_FAILED Eventlog_Type = 7
	// ISCSI_ERR_ISCSI_REQUEST_TIMEOUT enum
	Eventlog_Type_ISCSI_ERR_ISCSI_REQUEST_TIMEOUT Eventlog_Type = 8
	// ISCSI_ERR_SCSI_REQUEST_TIMEOUT enum
	Eventlog_Type_ISCSI_ERR_SCSI_REQUEST_TIMEOUT Eventlog_Type = 9
	// ISCSI_ERR_LOGIN_FAILED enum
	Eventlog_Type_ISCSI_ERR_LOGIN_FAILED Eventlog_Type = 10
	// ISCSI_ERR_LOGIN_PDU_ERROR enum
	Eventlog_Type_ISCSI_ERR_LOGIN_PDU_ERROR Eventlog_Type = 11
	// ISCSI_ERR_INVALID_LOGIN_REDIRECT_DATA enum
	Eventlog_Type_ISCSI_ERR_INVALID_LOGIN_REDIRECT_DATA Eventlog_Type = 12
	// ISCSI_ERR_INVALID_AUTHMETHOD enum
	Eventlog_Type_ISCSI_ERR_INVALID_AUTHMETHOD Eventlog_Type = 13
	// ISCSI_ERR_INVALID_CHAP_ALGORITHM enum
	Eventlog_Type_ISCSI_ERR_INVALID_CHAP_ALGORITHM Eventlog_Type = 14
	// ISCSI_ERR_INVALID_CHAP_CHALLENGE enum
	Eventlog_Type_ISCSI_ERR_INVALID_CHAP_CHALLENGE Eventlog_Type = 15
	// ISCSI_ERR_INVALID_KEY_DURING_CHAP enum
	Eventlog_Type_ISCSI_ERR_INVALID_KEY_DURING_CHAP Eventlog_Type = 16
	// ISCSI_ERR_INVALID_CHAP_RESPONSE enum
	Eventlog_Type_ISCSI_ERR_INVALID_CHAP_RESPONSE Eventlog_Type = 17
	// ISCSI_ERR_HEADER_DIGEST_NEEDED enum
	Eventlog_Type_ISCSI_ERR_HEADER_DIGEST_NEEDED Eventlog_Type = 18
	// ISCSI_ERR_HEADER_DATA_NEEDED enum
	Eventlog_Type_ISCSI_ERR_HEADER_DATA_NEEDED Eventlog_Type = 19
	// ISCSI_ERR_CONNECTION_LOST enum
	Eventlog_Type_ISCSI_ERR_CONNECTION_LOST Eventlog_Type = 20
	// ISCSI_ERR_INVALID_DATA_SEGMENT_LENGTH enum
	Eventlog_Type_ISCSI_ERR_INVALID_DATA_SEGMENT_LENGTH Eventlog_Type = 21
	// ISCSI_ERR_HEADER_DIGEST_ERROR enum
	Eventlog_Type_ISCSI_ERR_HEADER_DIGEST_ERROR Eventlog_Type = 22
	// ISCSI_ERR_ISCSI_PDU_ERROR enum
	Eventlog_Type_ISCSI_ERR_ISCSI_PDU_ERROR Eventlog_Type = 23
	// ISCSI_ERR_UNKNOWN_ISCSI_OPCODE enum
	Eventlog_Type_ISCSI_ERR_UNKNOWN_ISCSI_OPCODE Eventlog_Type = 24
	// ISCSI_ERR_DATA_DIGEST_ERROR enum
	Eventlog_Type_ISCSI_ERR_DATA_DIGEST_ERROR Eventlog_Type = 25
	// ISCSI_ERR_EXCESS_DATA_SENT enum
	Eventlog_Type_ISCSI_ERR_EXCESS_DATA_SENT Eventlog_Type = 26
	// ISCSI_ERR_UNEXPECTED_PDU enum
	Eventlog_Type_ISCSI_ERR_UNEXPECTED_PDU Eventlog_Type = 27
	// ISCSI_ERR_INVALID_RTT_PDU enum
	Eventlog_Type_ISCSI_ERR_INVALID_RTT_PDU Eventlog_Type = 28
	// ISCSI_ERR_ISCSI_PDU_REJECTED enum
	Eventlog_Type_ISCSI_ERR_ISCSI_PDU_REJECTED Eventlog_Type = 29
	// ISCSI_ERR_INSUFFICIENT_WORKITEM_RESOURCES enum
	Eventlog_Type_ISCSI_ERR_INSUFFICIENT_WORKITEM_RESOURCES Eventlog_Type = 30
	// ISCSI_ERR_INSUFFICIENT_REQ_PACKET_RESOURCES enum
	Eventlog_Type_ISCSI_ERR_INSUFFICIENT_REQ_PACKET_RESOURCES Eventlog_Type = 31
	// ISCSI_WRN_RECEIVED_ASYNC_LOGOUT enum
	Eventlog_Type_ISCSI_WRN_RECEIVED_ASYNC_LOGOUT Eventlog_Type = 32
	// ISCSI_ERR_INVALID_CHAP_CHALLENGE_SIZE enum
	Eventlog_Type_ISCSI_ERR_INVALID_CHAP_CHALLENGE_SIZE Eventlog_Type = 33
	// ISCSI_INFO_RECONNECTED_TO_TARGET enum
	Eventlog_Type_ISCSI_INFO_RECONNECTED_TO_TARGET Eventlog_Type = 34
	// ISCSI_ERR_INVALID_TARGET_CHAP_SECRET enum
	Eventlog_Type_ISCSI_ERR_INVALID_TARGET_CHAP_SECRET Eventlog_Type = 35
	// ISCSI_ERR_INVALID_INITIATOR_CHAP_SECRET enum
	Eventlog_Type_ISCSI_ERR_INVALID_INITIATOR_CHAP_SECRET Eventlog_Type = 36
	// ISCSI_ERR_FIPS_NOT_AVAILABLE enum
	Eventlog_Type_ISCSI_ERR_FIPS_NOT_AVAILABLE Eventlog_Type = 37
	// ISCSI_ERR_CHAP_NOT_OFFERED enum
	Eventlog_Type_ISCSI_ERR_CHAP_NOT_OFFERED Eventlog_Type = 38
	// ISCSI_ERR_DEVICE_RESET enum
	Eventlog_Type_ISCSI_ERR_DEVICE_RESET Eventlog_Type = 39
	// ISCSI_ERR_CHAP_OFFERED enum
	Eventlog_Type_ISCSI_ERR_CHAP_OFFERED Eventlog_Type = 40
	// ISCSI_ERR_AUTH_METHOD_NOT_OFFERED enum
	Eventlog_Type_ISCSI_ERR_AUTH_METHOD_NOT_OFFERED Eventlog_Type = 41
	// ISCSI_ERR_INVALID_STATUS_SEQ_NUM enum
	Eventlog_Type_ISCSI_ERR_INVALID_STATUS_SEQ_NUM Eventlog_Type = 42
)
