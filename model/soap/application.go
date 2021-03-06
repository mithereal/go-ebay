// Code generated by gowsdl DO NOT EDIT.

package ebay

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

//
// Indicates whether the error is a severe error (causing the request to fail)
// or an informational error (a warning) that should be communicated to the
// user.
//
type AckValue string

const (

	//
	// The request was processed successfully, but something occurred that may
	// affect your application or the user.
	//
	AckValueSuccess AckValue = "Success"

	//
	// The request that triggered the error was not processed successfully.
	// When a serious application-level error occurs, the error is returned
	// instead of the business data.
	//
	AckValueFailure AckValue = "Failure"

	//
	// The request that triggered the error was processed successfully but with some warnings.
	//
	AckValueWarning AckValue = "Warning"

	//
	// The request that triggered the error was processed successfully but with some warnings.
	//
	AckValuePartialFailure AckValue = "PartialFailure"
)

//
// Indicates whether the error is a severe error (causing the request to fail)
// or an informational error (a warning) that should be communicated to the
// user.
//
type ErrorSeverity string

const (

	//
	// The request that triggered the error was not processed successfully.
	// When a serious application-level error occurs, the error is returned
	// instead of the business data.
	//
	ErrorSeverityError ErrorSeverity = "Error"

	//
	// The request was processed successfully, but something occurred that may
	// affect your application or the user. For example, eBay may have changed a
	// value the user sent in. In this case, eBay returns a normal, successful
	// response and also returns the warning.
	//
	//ErrorSeverityWarning ErrorSeverity = "Warning"
)

//
// There are three categories of error: request errors, application errors,
// and system errors. Request and application errors are caused primarily by
// invalid data passed in the request. System errors are caused primarily by
// application failures and cannot be corrected by changing request values.
//
type ErrorCategory string

const (

	//
	// Indicates that an error has occurred on the eBay system side, such as a
	// database or server down. An application can retry the request as-is a
	// reasonable number of times (eBay recommends twice). If the error
	// persists, contact Developer Technical Support. Once the problem has been
	// resolved, the request may be resent in its original form.
	//
	ErrorCategorySystem ErrorCategory = "System"

	//
	// An error occurred due to a problem with the request, such as missing or
	// invalid fields. The problem must be corrected before the request can be
	// made again. If the problem is due to something in the application (such
	// as a missing required field), the application must be changed. Once the
	// problem in the appilcation or data is resolved, resend the corrected
	// request to eBay.
	//
	ErrorCategoryApplication ErrorCategory = "Application"

	//
	// An error occurred due to a problem with the request, such as invalid or
	// missing data. The problem must be corrected before the request can be
	// made again. If the problem is a result of end-user data, the application
	// must alert the end-user to the problem and provide the means for the
	// end-user to correct the data. Once the problem in the data is resolved,
	// resend the request to eBay with the corrected data.
	//
	ErrorCategoryRequest ErrorCategory = "Request"
)

//
// Container for available types of charges that your application's subscribers can incur.
//
type OpeneBaySubscriptionChargeType string

const (

	//
	// The subscription is on a free trial basis.
	//
	OpeneBaySubscriptionChargeTypeFreeTrial OpeneBaySubscriptionChargeType = "FreeTrial"

	//
	// Non-recurring charge.
	//
	OpeneBaySubscriptionChargeTypeNRC OpeneBaySubscriptionChargeType = "NRC"

	//
	// Non-recurring setup charge.
	//
	OpeneBaySubscriptionChargeTypeNRCSetup OpeneBaySubscriptionChargeType = "NRCSetup"

	//
	// Recurring charge for a plan when the plan charges are not to be pro-rated at the
	// end of subscription. Not applicable for usage.
	//
	OpeneBaySubscriptionChargeTypeRecurring OpeneBaySubscriptionChargeType = "Recurring"

	//
	// Recurring charge for a plan when the plan charges are to be pro-rated at the
	// end of subscription. Not applicable for usage.
	//
	OpeneBaySubscriptionChargeTypeRecurringProRateEnd OpeneBaySubscriptionChargeType = "RecurringProRateEnd"

	//
	// Indicates a plan for which the subscriber will pay usage charges.
	//
	OpeneBaySubscriptionChargeTypeUsage OpeneBaySubscriptionChargeType = "Usage"

	//
	// A free plan that will not have usage or other charges.
	//
	OpeneBaySubscriptionChargeTypeFree OpeneBaySubscriptionChargeType = "Free"

	//
	// Billable charge that is not included in the plan and is not subject to
	// eBay service fee. For example, shipping charges.
	//
	OpeneBaySubscriptionChargeTypeNonPlanUsage OpeneBaySubscriptionChargeType = "NonPlanUsage"
)

//
// Container for details about the type of billing record.
//
type OpeneBayBillingRecordTypeType string

const (

	//
	// Indicates that the current billing record represents a charge for a subscription.
	//
	OpeneBayBillingRecordTypeTypeSubscriptionCharge OpeneBayBillingRecordTypeType = "SubscriptionCharge"

	//
	// Indicates that the current billing record represents a one-time charge.
	//
	OpeneBayBillingRecordTypeTypeOneTimeCharge OpeneBayBillingRecordTypeType = "OneTimeCharge"

	//
	// Indicates that the current billing record represents a usage charge.
	//
	OpeneBayBillingRecordTypeTypeUsageCharge OpeneBayBillingRecordTypeType = "UsageCharge"

	//
	// Indicates that the current billing record represents a statement.
	//
	OpeneBayBillingRecordTypeTypeStatement OpeneBayBillingRecordTypeType = "Statement"
)

//
// Container for the state of the credit.
//
type OpeneBayCreditLifeCycleStateType string

const (

	//
	// Indicates that the credit has been submitted.
	//
	OpeneBayCreditLifeCycleStateTypeSubmitted OpeneBayCreditLifeCycleStateType = "Submitted"

	//
	// Indicates that the credit transaction is pending.
	//
	OpeneBayCreditLifeCycleStateTypePending OpeneBayCreditLifeCycleStateType = "Pending"

	//
	// Indicates that the credit transaction has been processed.
	//
	OpeneBayCreditLifeCycleStateTypeProcessed OpeneBayCreditLifeCycleStateType = "Processed"

	//
	// Indicates that the credit transaction has failed to be processed.
	//
	OpeneBayCreditLifeCycleStateTypeFailed OpeneBayCreditLifeCycleStateType = "Failed"
)

//
// Container for explanation of the credit.
//
type OpeneBayCreditReasonType string

const (

	//
	// Explains why the credit has been applied to the subscriber's account.
	//
	OpeneBayCreditReasonTypeStatementCredit OpeneBayCreditReasonType = "StatementCredit"
)

type AddUsageRequest OpeneBayAddUsageRequestType

type AddUsageResponse OpeneBayAddUsageResponseType

type GetBillingStatementsRequest OpeneBayGetBillingStatementsRequestType

type GetBillingStatementsResponse OpeneBayGetBillingStatementsResponseType

type GetBillingRecordsRequest OpeneBayGetBillingRecordsRequestType

type GetBillingRecordsResponse OpeneBayGetBillingRecordsResponseType

type AddCreditRequest OpeneBayAddCreditRequestType

type AddCreditResponse OpeneBayAddCreditResponseType

type GetCreditsRequest OpeneBayGetCreditsRequestType

type GetCreditsResponse OpeneBayGetCreditsResponseType

type SetBillingStartDateRequest OpeneBaySetBillingStartDateRequestType

type SetBillingStartDateResponse OpeneBaySetBillingStartDateResponseType

type Amount struct {
	Value float64

	//
	// Currency in which the monetary amount is specified. A three-letter ID, such as
	// USD, CAD, DEM. Currently, USD is the only available value.
	//
	CurrencyId string `xml:"currencyId,attr,omitempty"`
}

type TimeRange struct {

	//
	// Specifies the earliest (oldest) date to be used in a date range.
	//
	TimeFrom time.Time `xml:"timeFrom,omitempty"`

	//
	// Specifies the latest (most recent) date to be used in a date range.
	//
	TimeTo time.Time `xml:"timeTo,omitempty"`
}

type BaseServiceRequest struct {
}

type BaseServiceResponse struct {

	//
	// Indicates whether there are any errors or warnings associated with the
	// processing of the request.
	//
	Ack *AckValue `xml:"ack,omitempty"`

	//
	// Information for an error or warning that occurred when eBay processed the
	// request.
	//
	ErrorMessage *ErrorMessage `xml:"errorMessage,omitempty"`

	//
	// Open eBay version.
	//
	Version string `xml:"version,omitempty"`

	//
	// The date and time when eBay processed the request. The time zone of this value
	// is GMT and the format is the ISO 8601 date and time format (YYYY-MM-
	// DDTHH:MM:SS.SSSZ). See Time Values in the eBay Web Services guide for
	// information about this time format and converting to and from the GMT time zone.
	//
	Timestamp time.Time `xml:"timestamp,omitempty"`
}

type ErrorMessage struct {

	//
	// ActivityProfile about a single error.
	//
	Error []*ErrorData `xml:"error,omitempty"`
}

type ErrorData struct {

	//
	// A unique code that identifies the particular error condition that occurred.
	// Your application can use error codes as identifiers in your customized
	// error-handling algorithms.
	//
	ErrorId int64 `xml:"errorId,omitempty"`

	//
	// Name of the domain upon which the error occurred.
	// <dl>
	// <lh>Domains include:</lh>
	// <dt>
	// Marketplace
	// </dt>
	// <dd>
	// A business or validation error occurred for the UserProfile Service.
	// </dd>
	// <dt>
	// SOA
	// </dt>
	// <dd>
	// An exception occurred in the Service Oriented Architecture (SOA) framework.
	// </dd>
	// </dl>
	//
	Domain string `xml:"domain,omitempty"`

	//
	// Indicates whether the error caused the request to fail (Error) or not
	// (Warning).
	// <br><br>
	// If the request fails and the source of the problem is within the application
	// (such as a missing required element), please change the application before you
	// retry the request. If the problem is due to end-user input data, please alert
	// the end-user to the problem and provide the means for them to correct the data.
	// Once the problem in the application or data is resolved, you can attempt to re-
	// send the request to eBay.
	// <br><br>
	// If the source of the problem is on eBay's side, you can retry the request as-is
	// a reasonable number of times (eBay recommends twice). If the error persists,
	// contact Developer Technical Support. Once the problem has been resolved, the
	// request may be resent in its original form.
	// <br><br>
	// When a warning occurs, the error is returned in addition to the business data.
	// In this case, you do not need to retry the request (as the original request was
	// successful). However, depending on the cause or nature of the warning, you
	// might need to contact either the end user or eBay to effect a long term
	// solution to the problem to prevent it from reoccurring in the future.
	//
	Severity *ErrorSeverity `xml:"severity,omitempty"`

	//
	// There are three categories of errors: request errors, application errors, and
	// system errors.
	//
	Category *ErrorCategory `xml:"category,omitempty"`

	//
	// A detailed description of the condition that resulted in the error.
	//
	Message string `xml:"message,omitempty"`

	//
	// Name of the subdomain upon which the error occurred. Subdomains include the
	// following: UserProfile (in which the error is specific to the UserProfile
	// service) and MarketplaceCommon (in which the error is common to all
	// Marketplace
	// services).
	//
	Subdomain string `xml:"subdomain,omitempty"`

	//
	// Unique identifier for an exception associated with an error.
	//
	ExceptionId string `xml:"exceptionId,omitempty"`

	//
	// Some warning and error messages return one or more variables that contain
	// contextual information about the error. This is often the field or value that
	// triggered the error.
	//
	Parameter []*ErrorParameter `xml:"parameter,omitempty"`
}

type ErrorParameter struct {
	Value string

	//
	// The name of the parameter in the list of parameter types returned
	// within the error type.
	//
	Name string `xml:"name,attr,omitempty"`
}

type OpeneBayBillingStatementType struct {

	//
	// Unique identifier for a billing statement, created by eBay at statement time.
	//
	StatementId string `xml:"statementId,omitempty"`

	//
	// Time a statement was generated. Subscribers will see the date only.
	//
	StatementTime time.Time `xml:"statementTime,omitempty"`

	//
	// Time at which full payment was made.
	//
	ClosedTime time.Time `xml:"closedTime,omitempty"`

	//
	// Previous statement balance, plus new charges (RC, NRC, Usage, Prebill Usage
	// Adjustments, Discounts, Unit Credits and Taxes) and payments.
	//
	StatementTotal *Amount `xml:"statementTotal,omitempty"`

	//
	// Net adjustments applied to this statement. These are post-bill adjustments only. A
	// negative value denotes a credit and a positive value denotes a debit. Changes such as
	// payments or adjustments are made against statements.
	//
	AdjustmentTotal *Amount `xml:"adjustmentTotal,omitempty"`

	//
	// Total payments already applied to the statement.
	//
	PaymentTotal *Amount `xml:"paymentTotal,omitempty"`

	//
	// Total new charges (combination of postitive and negative charges). newChargeTotal
	// minus adjustmentTotal lets you know the available balance to credit for the
	// statement.
	//
	NewChargeTotal *Amount `xml:"newChargeTotal,omitempty"`

	//
	// Total new negative charges for statement. Includes pre-billed adjustments. Not
	// changed after statement creation.
	//
	NewCreditTotal *Amount `xml:"newCreditTotal,omitempty"`

	//
	// Sum of of new charges plus adjustments plus payments.
	//
	BalanceDue *Amount `xml:"balanceDue,omitempty"`
}

type OpeneBayBillingRecordType struct {

	//
	// Identifies a combination of application and subscriber. Also known as EID.
	//
	BillingAccountId string `xml:"billingAccountId,omitempty"`

	//
	// Specifies whether the record is a one-time charge, a statement, a subscription
	// charge, or a usage charge. For these four different kinds of records, the values
	// transmitted for other elements, such as taxAmount and recordTime, represent elements
	// appropriate for the record type. For example, for recordTime, if the recordType is
	// Usage then recordTime represents the transaction date; but if recordType is Statement,
	// recordTime represents the time of the statement's creation. See the specific values
	// in each element.
	//
	RecordType *OpeneBayBillingRecordTypeType `xml:"recordType,omitempty"`

	//
	// Unique identifier for the record, assigned by eBay. Contains state and location
	// information about the record. An example of a recordId: 1111:222:333:444:x.
	// Retrieve a recordID using getBillingRecords.
	//
	RecordId string `xml:"recordId,omitempty"`

	//
	// Specifies whether or not a statement has been presented to the subscriber for the
	// record.
	//
	Billed bool `xml:"billed,omitempty"`

	//
	// Unique identifier for a billing statement, created by eBay at statement time.
	//
	StatementId string `xml:"statementId,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, recordTime represents the transaction date,
	// equivalent to the statement date.
	// <br><br>
	// If recordType is OneTimeCharge, recordTime represents the time of the transaction,
	// equivalent to the statement date.
	// <br><br>
	// If recordType is Usage, recordTime represents the time of the transaction,
	// equivalent to the statement date.
	// <br><br>
	// If recordType is Statement, recordTime represents the statement date.
	//
	RecordTime time.Time `xml:"recordTime,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, recordTime2 represents the beginning of the time range
	// for which the subscriber is being charged.
	// <br><br>
	// If recordType is OneTimeCharge, recordTime2 represents the effective date of the charge.
	// <br><br>
	// If recordType is Usage, recordTime2 represents the date the charge was calculated
	// by your application.
	// <br><br>
	// If recordType is Statement, recordTime2 represents the statement preparation date.
	//
	RecordTime2 time.Time `xml:"recordTime2,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, amount represents the total amount due for the
	// subscription to two or fewer digits after the decimal point in the subscriber's currency.
	// <br><br>
	// If recordType is OneTimeCharge, amount represents the amount of the charge to two or fewer digits after the decimal
	// point in the subscriber's currency.
	// <br><br>
	// If recordType is Usage, amount represents the total of the usage charge to two or fewer digits after the decimal
	// point in the subscriber's currency.
	// <br><br>
	// If recordType is Statement, amount represents the amount available to credit at
	// the time of the statement, to two or fewer digits after the decimal
	// point, in the subscriber's currency.
	//
	Amount *Amount `xml:"amount,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, amount2 represents the amount for the subscription
	// in the biller's currency.
	// <br><br>
	// If recordType is OneTimeCharge, amount2 represents the amount of the charge in
	// the biller's currency.
	// <br><br>
	// If recordType is Usage, amount2 represents the amount charged for the usage in
	// the biller's currency.
	// <br><br>
	// If recordType is Statement, amount2 represents the statement total, in the biller's currency.
	//
	Amount2 *Amount `xml:"amount2,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, discountAmount is the amount of a discount applied
	// to the charges in the billing record, to two or fewer digits after the decimal
	// point, in the billing currency.
	// <br><br>
	// If recordType is OneTimeCharge, discountAmount is the amount of the discount
	// applied to the charges in the billing record, to two or fewer digits after the decimal
	// point, in the billing currency.
	// <br><br>
	// If recordType is Usage, discountAmount is the amount of the discount applied
	// to the charges in the billing record, to two or fewer digits after the decimal
	// point, in the billing currency.
	// <br><br>
	// If recordType is Statement, discountAmount is not returned because no further
	// adjustments can be made once the statement has been created.
	//
	DiscountAmount *Amount `xml:"discountAmount,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, taxAmount is the amount of the tax to two or fewer
	// digits after the decimal point, in the billing currency.
	// <br><br>
	// If recordType is OneTimeCharge, taxAmount is the amount of the tax to two or fewer digits after the decimal
	// point, in the billing currency.
	// <br><br>
	// If recordType is Usage, taxAmount is the amount of the tax to two or fewer digits after the decimal
	// point, in the billing currency.
	// <br><br>
	// If recordType is Statement, taxAmount is the total amount of any tax included in
	// the statement.
	//
	TaxAmount *Amount `xml:"taxAmount,omitempty"`

	//
	// Returns zero if no tax rate, or the tax rate. For euros, Luxembourg rate of 15%
	// will be assumed. This field is for future use.
	//
	TaxRate float64 `xml:"taxRate,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, recordDescription contains a text description of the subscription.
	// <br><br>
	// If recordType is OneTimeCharge, recordDescription contains a text description of the charge.
	// <br><br>
	// If recordType is Usage, recordDescription contains a text description of the usage.
	// <br><br>
	// If recordType is Statement, recordDescription contains a text description of the statement.
	//
	RecordDescription string `xml:"recordDescription,omitempty"`

	//
	// Returns true if adjustments such as discounts, credits, and usage can be applied
	// to the record.
	//
	Adjustable bool `xml:"adjustable,omitempty"`
}

type OpeneBayCreditType struct {

	//
	// Unique identifier assigned to the credit request by eBay.
	//
	CreditId int64 `xml:"creditId,omitempty"`

	//
	// The eBay user ID of the subscriber.
	//
	UserName string `xml:"userName,omitempty"`

	//
	// Total amount of the credit applied to the subscriber's account.
	//
	CreditAmount *Amount `xml:"creditAmount,omitempty"`

	//
	// The date assigned to the credit request by the application.
	//
	CreditTime time.Time `xml:"creditTime,omitempty"`

	//
	// A unique identifier for a credit.
	//
	ExternalCreditId string `xml:"externalCreditId,omitempty"`

	//
	// Explains why the credit has been applied to the subscriber's account.
	//
	CreditReason *OpeneBayCreditReasonType `xml:"creditReason,omitempty"`

	//
	// Describes the credit applied to the subscriber's account.
	//
	CreditDescription string `xml:"creditDescription,omitempty"`

	//
	// A unique identifier for each combination of an application and a subscriber.
	//
	BillingAccountId string `xml:"billingAccountId,omitempty"`

	//
	// Unique identifier for the billing statement in which the credit was applied to the
	// subscriber's account, created by eBay at statement time.
	//
	StatementId string `xml:"statementId,omitempty"`

	//
	// Specifies whether the record is a one-time charge, a statement, a subscription
	// charge, or a usage charge. For these four different kinds of records, the values
	// transmitted for other elements, such as taxAmount and recordTime, represent elements
	// appropriate for the record type. For example, for recordTime, if the recordType is
	// Usage then recordTime represents the transaction date; but if recordType is Statement,
	// recordTime represents the time of the statement's creation. See the specific values
	// in each element.
	//
	RecordType *OpeneBayBillingRecordTypeType `xml:"recordType,omitempty"`

	//
	// Unique identifier for the record, assigned by eBay. Contains state and location
	// information about the record. An example of a recordId: 1111:222:333:444:x.
	//
	RecordId string `xml:"recordId,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, recordTime represents the statement date.
	// <br><br>
	// If recordType is OneTimeCharge, recordTime represents the time of the transaction.
	// <br><br>
	// If recordType is Usage, recordTime represents the time of the transaction.
	// <br><br>
	// If recordType is Statement, recordTime represents the statement time.
	//
	RecordTime time.Time `xml:"recordTime,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, amount2 represents the credit amount in the
	// billing currency.
	// <br><br>
	// If recordType is OneTimeCharge, amount2 represents the the credit amount in the
	// billing currency.
	// <br><br>
	// If recordType is Usage, amount2 represents the credit amount in the billing
	// currency.
	// <br><br>
	// If recordType is Statement, amount2 represents the statement total.
	//
	Amount2 *Amount `xml:"amount2,omitempty"`

	//
	// The recordType determines what this value signifies.
	// <br><br>
	// If recordType is Subscription, recordDescription contains a text description of the subscription.
	// <br><br>
	// If recordType is OneTimeCharge, recordDescription contains a text description of the charge.
	// <br><br>
	// If recordType is Usage, recordDescription contains a text description of the usage.
	// <br><br>
	// If recordType is Statement, recordDescription contains a text description of the statement.
	//
	RecordDescription string `xml:"recordDescription,omitempty"`

	//
	// Time the credit was processed by eBay.
	//
	LastProcessedTime time.Time `xml:"lastProcessedTime,omitempty"`

	//
	// The lifecycle state of the credit can be submitted, pending, processed, or failed.
	//
	State *OpeneBayCreditLifeCycleStateType `xml:"state,omitempty"`

	//
	// When a credit is manually processed, may include a note from eBay to the
	// application such as, "Your credit request was processed successfully."
	//
	ProcessingNote string `xml:"processingNote,omitempty"`
}

type OpeneBayAddUsageRequestType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services addUsageRequest"`

	*BaseServiceRequest

	//
	// A planId is assigned by eBay to each subscription plan.
	// Your application often uses this value in calls
	// to the Open eBay Application Integration Service.
	// Note that the planId is diffrent from the externalPlanId.
	// In contrast to the planId, the externalPlanId is the
	// value that you provide,
	// as a "Developer Plan ID," when you create a subscription plan.
	//
	PlanId int64 `xml:"planId,omitempty"`

	//
	// For a particular application, the subscriptionId
	// uniquely identifies a user's subscription.
	//
	SubscriptionId int64 `xml:"subscriptionId,omitempty"`

	//
	// The eBay user ID of the subscriber.
	//
	UserName string `xml:"userName,omitempty"`

	//
	// The transaction date assigned to the transaction by the application.
	//
	TransactionTime time.Time `xml:"transactionTime,omitempty"`

	//
	// The transaction reference number assigned by the application.
	//
	ExternalTransactionId string `xml:"externalTransactionId,omitempty"`

	//
	// Memo that will be visible on the subscriber statement.
	//
	Memo string `xml:"memo,omitempty"`

	//
	// Charge amount that eBay will display to the subscriber on behalf of your
	// application. Charge amount must have two or fewer digits after the decimal
	// point and 12 or fewer digits before the decimal point.
	//
	ChargeAmount *Amount `xml:"chargeAmount,omitempty"`

	//
	// Identifier of the usage charge. This is an action code set up in your application's billing plan.
	//
	ChargeType *OpeneBaySubscriptionChargeType `xml:"chargeType,omitempty"`

	//
	// Set to true when you require immediate payment for usage that is not included in plan.
	//
	ImmediatePayment bool `xml:"immediatePayment,omitempty"`
}

type OpeneBayAddUsageResponseType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services addUsageResponse"`

	*BaseServiceResponse

	//
	// eBay assigns this ID to the transaction and returns it to you.
	//
	TransactionId int64 `xml:"transactionId,omitempty"`
}

type OpeneBayGetBillingStatementsRequestType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services getBillingStatementsRequest"`

	*BaseServiceRequest

	//
	// A planId is assigned by eBay to each subscription plan.
	// Your application often uses this value in calls
	// to the Open eBay Application Integration Service.
	// Note that the planId is diffrent from the externalPlanId.
	// In contrast to the planId, the externalPlanId is the
	// value that you provide,
	// as a "Developer Plan ID," when you create a subscription plan.
	//
	PlanId int64 `xml:"planId,omitempty"`

	//
	// For a particular application, the subscriptionId
	// uniquely identifies a user's subscription.
	//
	SubscriptionId int64 `xml:"subscriptionId,omitempty"`

	//
	// The eBay user ID of the subscriber.
	//
	UserName string `xml:"userName,omitempty"`
}

type OpeneBayGetBillingStatementsResponseType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services getBillingStatementsResponse"`

	*BaseServiceResponse

	//
	// Billing statement details.
	//
	Statement []*OpeneBayBillingStatementType `xml:"statement,omitempty"`
}

type OpeneBayGetBillingRecordsRequestType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services getBillingRecordsRequest"`

	*BaseServiceRequest

	//
	// A planId is assigned by eBay to each subscription plan.
	// Your application often uses this value in calls
	// to the Open eBay Application Integration Service.
	// Note that the planId is diffrent from the externalPlanId.
	// In contrast to the planId, the externalPlanId is the
	// value that you provide,
	// as a "Developer Plan ID," when you create a subscription plan.
	//
	PlanId int64 `xml:"planId,omitempty"`

	//
	// For a particular application, the subscriptionId
	// uniquely identifies a user's subscription.
	//
	SubscriptionId int64 `xml:"subscriptionId,omitempty"`

	//
	// The eBay user ID of the subscriber.
	//
	UserName string `xml:"userName,omitempty"`

	//
	// Specifies whether the record is a one-time charge, a statement, a
	// subscription charge, or a usage charge.
	//
	RecordType *OpeneBayBillingRecordTypeType `xml:"recordType,omitempty"`

	//
	// Unique identifier for a billing statement, created by eBay at statement time. Obtain this value from
	// the response to the getBillingStatements request.
	//
	StatementId string `xml:"statementId,omitempty"`

	//
	// Requested time range for records.
	//
	RecordTimeRange *TimeRange `xml:"recordTimeRange,omitempty"`
}

type OpeneBayGetBillingRecordsResponseType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services getBillingRecordsResponse"`

	*BaseServiceResponse

	//
	// Billing account transaction record.
	//
	Record []*OpeneBayBillingRecordType `xml:"record,omitempty"`
}

type OpeneBayAddCreditRequestType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services addCreditRequest"`

	*BaseServiceRequest

	//
	// A planId is assigned by eBay to each subscription plan.
	// Your application often uses this value in calls
	// to the Open eBay Application Integration Service.
	// Note that the planId is diffrent from the externalPlanId.
	// In contrast to the planId, the externalPlanId is the
	// value that you provide,
	// as a "Developer Plan ID," when you create a subscription plan.
	//
	PlanId int64 `xml:"planId,omitempty"`

	//
	// For a particular application, the subscriptionId
	// uniquely identifies a user's subscription.
	//
	SubscriptionId int64 `xml:"subscriptionId,omitempty"`

	//
	// The eBay user ID of the subscriber.
	//
	UserName string `xml:"userName,omitempty"`

	//
	// Amount to be credited to the subscriber.
	// The amount is in the billing currency (the currency specified in the
	// original charge).
	//
	CreditAmount *Amount `xml:"creditAmount,omitempty"`

	//
	// The date assigned to the credit request by the application.
	//
	CreditTime time.Time `xml:"creditTime,omitempty"`

	//
	// A transaction reference number that you can assign to the credit. This number
	// will be returned in the getCredit response.
	//
	ExternalCreditId string `xml:"externalCreditId,omitempty"`

	//
	// Explains why the credit has been applied to the subscriber's account.
	//
	CreditReason *OpeneBayCreditReasonType `xml:"creditReason,omitempty"`

	//
	// Describes the credit applied to the subscriber's account.
	//
	CreditDescription string `xml:"creditDescription,omitempty"`

	//
	// A unique identifier for each combination of an application and a subscriber.
	//
	BillingAccountId string `xml:"billingAccountId,omitempty"`

	//
	// Unique identifier for a billing statement, created by eBay at statement time.	Returned in the
	// getBillingRecords response.
	//
	StatementId string `xml:"statementId,omitempty"`

	//
	// Specifies whether the record is a one-time charge, a statement, a subscription
	// charge, or a usage charge. For these four different kinds of records, the values
	// transmitted for other elements, such as taxAmount and recordTime, represent elements
	// appropriate for the record type. For example, for recordTime, if the recordType is
	// Usage then recordTime represents the transaction date; but if recordType is Statement,
	// recordTime represents the time of the statement's creation. See the specific values
	// in each element. Returned in the getBillingRecords response.
	//
	RecordType *OpeneBayBillingRecordTypeType `xml:"recordType,omitempty"`

	//
	// Unique identifier for the record, assigned by eBay. Contains state and location
	// information about the record. An example of a recordId: 1111:222:333:444:x.
	// Returned in the getBillingRecords response.
	//
	RecordId string `xml:"recordId,omitempty"`

	//
	// The recordType determines what this value signifies. Returned in the getBillingRecords response.
	// <br><br>
	// If recordType is Subscription, recordTime represents the beginning of the time range, the time
	// from, that ends at the time to given in recordTime2.
	// <br><br>
	// If recordType is OneTimeCharge, recordTime represents the time of the transaction.
	// <br><br>
	// If recordType is Usage, recordTime represents the time of the transaction.
	// <br><br>
	// If recordType is Statement, recordTime represents the statement time.
	//
	RecordTime time.Time `xml:"recordTime,omitempty"`

	//
	// The recordType determines what this value signifies. Returned in the getBillingRecords response.
	// eBay uses this value to ensure that the credit is matched to the correct original charge.
	// <br><br>
	// If recordType is Subscription, amount2 represents the subscription amount, in the billing currency,
	// against which the credit is to be applied.
	// <br><br>
	// If recordType is OneTimeCharge, amount2 represents the charge amount, in the billing currency,
	// against which the credit is to be applied.
	// <br><br>
	// If recordType is Usage, amount2 represents the amount of the usage charge, in the billing currency,
	// against which the credit is to be applied.
	// <br><br>
	// If recordType is Statement, amount2 represents the statement total, in the billing currency,
	// against which the credit is to be applied.
	//
	Amount2 *Amount `xml:"amount2,omitempty"`

	//
	// The recordType determines what this value signifies. Returned in the getBillingRecords response.
	// <br><br>
	// If recordType is Subscription, recordDescription contains a text description of the subscription.
	// <br><br>
	// If recordType is OneTimeCharge, recordDescription contains a text description of the charge.
	// <br><br>
	// If recordType is Usage, recordDescription contains a text description of the usage.
	// <br><br>
	// If recordType is Statement, recordDescription contains a text description of the statement.
	//
	RecordDescription string `xml:"recordDescription,omitempty"`
}

type OpeneBayAddCreditResponseType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services addCreditResponse"`

	*BaseServiceResponse

	//
	// Unique identifier assigned to the credit request by eBay.
	//
	CreditId int64 `xml:"creditId,omitempty"`
}

type OpeneBayGetCreditsRequestType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services getCreditsRequest"`

	*BaseServiceRequest

	//
	// The eBay user ID of the subscriber.
	//
	UserName string `xml:"userName,omitempty"`

	//
	// Unique identifier assigned to the credit request by eBay.
	//
	CreditId int64 `xml:"creditId,omitempty"`

	//
	// The lifecycle state of the credit can be submitted, pending, processed, or failed.
	//
	State *OpeneBayCreditLifeCycleStateType `xml:"state,omitempty"`

	//
	// Time range filter for the requested credits.
	//
	CreditTimeRange *TimeRange `xml:"creditTimeRange,omitempty"`
}

type OpeneBayGetCreditsResponseType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services getCreditsResponse"`

	*BaseServiceResponse

	//
	// Details about the requested credits.
	//
	Credit []*OpeneBayCreditType `xml:"credit,omitempty"`
}

type OpeneBaySetBillingStartDateRequestType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services setBillingStartDateRequest"`

	*BaseServiceRequest

	//
	// A planId is assigned by eBay to each subscription plan.
	// Your application often uses this value in calls
	// to the Open eBay Application Integration Service.
	// Note that the planId is diffrent from the externalPlanId.
	// In contrast to the planId, the externalPlanId is the
	// value that you provide,
	// as a "Developer Plan ID," when you create a subscription plan.
	//
	PlanId int64 `xml:"planId,omitempty"`

	//
	// For a particular application, the subscriptionId
	// uniquely identifies a user's subscription.
	//
	SubscriptionId int64 `xml:"subscriptionId,omitempty"`

	//
	// The eBay user ID of the subscriber.
	//
	UserName string `xml:"userName,omitempty"`

	//
	// Date from which a subscription begins to incur subscription charges.
	//
	BillingStartDate time.Time `xml:"billingStartDate,omitempty"`
}

type OpeneBaySetBillingStartDateResponseType struct {
	XMLName xml.Name `xml:"http://www.ebay.com/marketplace/services setBillingStartDateResponse"`

	*BaseServiceResponse
}

type OpeneBayApplicationIntegrationServicePort interface {
	AddUsage(request *OpeneBayAddUsageRequestType) (*OpeneBayAddUsageResponseType, error)

	GetBillingStatements(request *OpeneBayGetBillingStatementsRequestType) (*OpeneBayGetBillingStatementsResponseType, error)

	GetBillingRecords(request *OpeneBayGetBillingRecordsRequestType) (*OpeneBayGetBillingRecordsResponseType, error)

	AddCredit(request *OpeneBayAddCreditRequestType) (*OpeneBayAddCreditResponseType, error)

	GetCredits(request *OpeneBayGetCreditsRequestType) (*OpeneBayGetCreditsResponseType, error)

	SetBillingStartDate(request *OpeneBaySetBillingStartDateRequestType) (*OpeneBaySetBillingStartDateResponseType, error)
}

type openeBayApplicationIntegrationServicePort struct {
	client *soap.Client
}

func NewOpeneBayApplicationIntegrationServicePort(client *soap.Client) OpeneBayApplicationIntegrationServicePort {
	return &openeBayApplicationIntegrationServicePort{
		client: client,
	}
}

func (service *openeBayApplicationIntegrationServicePort) AddUsage(request *OpeneBayAddUsageRequestType) (*OpeneBayAddUsageResponseType, error) {
	response := new(OpeneBayAddUsageResponseType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *openeBayApplicationIntegrationServicePort) GetBillingStatements(request *OpeneBayGetBillingStatementsRequestType) (*OpeneBayGetBillingStatementsResponseType, error) {
	response := new(OpeneBayGetBillingStatementsResponseType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *openeBayApplicationIntegrationServicePort) GetBillingRecords(request *OpeneBayGetBillingRecordsRequestType) (*OpeneBayGetBillingRecordsResponseType, error) {
	response := new(OpeneBayGetBillingRecordsResponseType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *openeBayApplicationIntegrationServicePort) AddCredit(request *OpeneBayAddCreditRequestType) (*OpeneBayAddCreditResponseType, error) {
	response := new(OpeneBayAddCreditResponseType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *openeBayApplicationIntegrationServicePort) GetCredits(request *OpeneBayGetCreditsRequestType) (*OpeneBayGetCreditsResponseType, error) {
	response := new(OpeneBayGetCreditsResponseType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *openeBayApplicationIntegrationServicePort) SetBillingStartDate(request *OpeneBaySetBillingStartDateRequestType) (*OpeneBaySetBillingStartDateResponseType, error) {
	response := new(OpeneBaySetBillingStartDateResponseType)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
