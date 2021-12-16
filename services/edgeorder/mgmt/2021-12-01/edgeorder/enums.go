package edgeorder

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ActionStatusEnum enumerates the values for action status enum.
type ActionStatusEnum string

const (
	// ActionStatusEnumAllowed Allowed flag.
	ActionStatusEnumAllowed ActionStatusEnum = "Allowed"
	// ActionStatusEnumNotAllowed Not Allowed flag.
	ActionStatusEnumNotAllowed ActionStatusEnum = "NotAllowed"
)

// PossibleActionStatusEnumValues returns an array of possible values for the ActionStatusEnum const type.
func PossibleActionStatusEnumValues() []ActionStatusEnum {
	return []ActionStatusEnum{ActionStatusEnumAllowed, ActionStatusEnumNotAllowed}
}

// ActionType enumerates the values for action type.
type ActionType string

const (
	// ActionTypeInternal ...
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns an array of possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{ActionTypeInternal}
}

// AddressType enumerates the values for address type.
type AddressType string

const (
	// AddressTypeCommercial Commercial Address.
	AddressTypeCommercial AddressType = "Commercial"
	// AddressTypeNone Address type not known.
	AddressTypeNone AddressType = "None"
	// AddressTypeResidential Residential Address.
	AddressTypeResidential AddressType = "Residential"
)

// PossibleAddressTypeValues returns an array of possible values for the AddressType const type.
func PossibleAddressTypeValues() []AddressType {
	return []AddressType{AddressTypeCommercial, AddressTypeNone, AddressTypeResidential}
}

// AddressValidationStatus enumerates the values for address validation status.
type AddressValidationStatus string

const (
	// AddressValidationStatusAmbiguous Address provided is ambiguous, please choose one of the alternate
	// addresses returned.
	AddressValidationStatusAmbiguous AddressValidationStatus = "Ambiguous"
	// AddressValidationStatusInvalid Address provided is invalid or not supported.
	AddressValidationStatusInvalid AddressValidationStatus = "Invalid"
	// AddressValidationStatusValid Address provided is valid.
	AddressValidationStatusValid AddressValidationStatus = "Valid"
)

// PossibleAddressValidationStatusValues returns an array of possible values for the AddressValidationStatus const type.
func PossibleAddressValidationStatusValues() []AddressValidationStatus {
	return []AddressValidationStatus{AddressValidationStatusAmbiguous, AddressValidationStatusInvalid, AddressValidationStatusValid}
}

// AvailabilityStage enumerates the values for availability stage.
type AvailabilityStage string

const (
	// AvailabilityStageAvailable Product is available.
	AvailabilityStageAvailable AvailabilityStage = "Available"
	// AvailabilityStageComingSoon Product is coming soon.
	AvailabilityStageComingSoon AvailabilityStage = "ComingSoon"
	// AvailabilityStageDeprecated Product is deprecated.
	AvailabilityStageDeprecated AvailabilityStage = "Deprecated"
	// AvailabilityStagePreview Product is in preview.
	AvailabilityStagePreview AvailabilityStage = "Preview"
	// AvailabilityStageSignup Product is available only on signup.
	AvailabilityStageSignup AvailabilityStage = "Signup"
	// AvailabilityStageUnavailable Product is not available.
	AvailabilityStageUnavailable AvailabilityStage = "Unavailable"
)

// PossibleAvailabilityStageValues returns an array of possible values for the AvailabilityStage const type.
func PossibleAvailabilityStageValues() []AvailabilityStage {
	return []AvailabilityStage{AvailabilityStageAvailable, AvailabilityStageComingSoon, AvailabilityStageDeprecated, AvailabilityStagePreview, AvailabilityStageSignup, AvailabilityStageUnavailable}
}

// BillingType enumerates the values for billing type.
type BillingType string

const (
	// BillingTypeMeterDetails ...
	BillingTypeMeterDetails BillingType = "MeterDetails"
	// BillingTypePav2 ...
	BillingTypePav2 BillingType = "Pav2"
	// BillingTypePurchase ...
	BillingTypePurchase BillingType = "Purchase"
)

// PossibleBillingTypeValues returns an array of possible values for the BillingType const type.
func PossibleBillingTypeValues() []BillingType {
	return []BillingType{BillingTypeMeterDetails, BillingTypePav2, BillingTypePurchase}
}

// ChargingType enumerates the values for charging type.
type ChargingType string

const (
	// ChargingTypePerDevice Per device charging type.
	ChargingTypePerDevice ChargingType = "PerDevice"
	// ChargingTypePerOrder Per order charging type.
	ChargingTypePerOrder ChargingType = "PerOrder"
)

// PossibleChargingTypeValues returns an array of possible values for the ChargingType const type.
func PossibleChargingTypeValues() []ChargingType {
	return []ChargingType{ChargingTypePerDevice, ChargingTypePerOrder}
}

// CreatedByType enumerates the values for created by type.
type CreatedByType string

const (
	// CreatedByTypeApplication ...
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey ...
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity ...
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser ...
	CreatedByTypeUser CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns an array of possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{CreatedByTypeApplication, CreatedByTypeKey, CreatedByTypeManagedIdentity, CreatedByTypeUser}
}

// DescriptionType enumerates the values for description type.
type DescriptionType string

const (
	// DescriptionTypeBase Base description.
	DescriptionTypeBase DescriptionType = "Base"
)

// PossibleDescriptionTypeValues returns an array of possible values for the DescriptionType const type.
func PossibleDescriptionTypeValues() []DescriptionType {
	return []DescriptionType{DescriptionTypeBase}
}

// DisabledReason enumerates the values for disabled reason.
type DisabledReason string

const (
	// DisabledReasonCountry Not available in the requested country.
	DisabledReasonCountry DisabledReason = "Country"
	// DisabledReasonFeature Required features are not enabled.
	DisabledReasonFeature DisabledReason = "Feature"
	// DisabledReasonNone Not disabled.
	DisabledReasonNone DisabledReason = "None"
	// DisabledReasonNoSubscriptionInfo Subscription has not registered to Microsoft.DataBox and Service does
	// not have the subscription notification.
	DisabledReasonNoSubscriptionInfo DisabledReason = "NoSubscriptionInfo"
	// DisabledReasonNotAvailable The product is not yet available.
	DisabledReasonNotAvailable DisabledReason = "NotAvailable"
	// DisabledReasonOfferType Subscription does not have required offer types.
	DisabledReasonOfferType DisabledReason = "OfferType"
	// DisabledReasonOutOfStock The product is out of stock.
	DisabledReasonOutOfStock DisabledReason = "OutOfStock"
	// DisabledReasonRegion Not available to push data to the requested Azure region.
	DisabledReasonRegion DisabledReason = "Region"
)

// PossibleDisabledReasonValues returns an array of possible values for the DisabledReason const type.
func PossibleDisabledReasonValues() []DisabledReason {
	return []DisabledReason{DisabledReasonCountry, DisabledReasonFeature, DisabledReasonNone, DisabledReasonNoSubscriptionInfo, DisabledReasonNotAvailable, DisabledReasonOfferType, DisabledReasonOutOfStock, DisabledReasonRegion}
}

// DoubleEncryptionStatus enumerates the values for double encryption status.
type DoubleEncryptionStatus string

const (
	// DoubleEncryptionStatusDisabled Double encryption is disabled
	DoubleEncryptionStatusDisabled DoubleEncryptionStatus = "Disabled"
	// DoubleEncryptionStatusEnabled Double encryption is enabled
	DoubleEncryptionStatusEnabled DoubleEncryptionStatus = "Enabled"
)

// PossibleDoubleEncryptionStatusValues returns an array of possible values for the DoubleEncryptionStatus const type.
func PossibleDoubleEncryptionStatusValues() []DoubleEncryptionStatus {
	return []DoubleEncryptionStatus{DoubleEncryptionStatusDisabled, DoubleEncryptionStatusEnabled}
}

// ImageType enumerates the values for image type.
type ImageType string

const (
	// ImageTypeBulletImage Bullet image.
	ImageTypeBulletImage ImageType = "BulletImage"
	// ImageTypeGenericImage Generic image.
	ImageTypeGenericImage ImageType = "GenericImage"
	// ImageTypeMainImage Main image.
	ImageTypeMainImage ImageType = "MainImage"
)

// PossibleImageTypeValues returns an array of possible values for the ImageType const type.
func PossibleImageTypeValues() []ImageType {
	return []ImageType{ImageTypeBulletImage, ImageTypeGenericImage, ImageTypeMainImage}
}

// LengthHeightUnit enumerates the values for length height unit.
type LengthHeightUnit string

const (
	// LengthHeightUnitCM Centimeter.
	LengthHeightUnitCM LengthHeightUnit = "CM"
	// LengthHeightUnitIN Inch, applicable for West US.
	LengthHeightUnitIN LengthHeightUnit = "IN"
)

// PossibleLengthHeightUnitValues returns an array of possible values for the LengthHeightUnit const type.
func PossibleLengthHeightUnitValues() []LengthHeightUnit {
	return []LengthHeightUnit{LengthHeightUnitCM, LengthHeightUnitIN}
}

// LinkType enumerates the values for link type.
type LinkType string

const (
	// LinkTypeDocumentation Link to product documentation
	LinkTypeDocumentation LinkType = "Documentation"
	// LinkTypeGeneric Generic link.
	LinkTypeGeneric LinkType = "Generic"
	// LinkTypeKnowMore Link to know more
	LinkTypeKnowMore LinkType = "KnowMore"
	// LinkTypeSignUp Link to sign up for products
	LinkTypeSignUp LinkType = "SignUp"
	// LinkTypeSpecification Link to product specification.
	LinkTypeSpecification LinkType = "Specification"
	// LinkTypeTermsAndConditions Terms and conditions link.
	LinkTypeTermsAndConditions LinkType = "TermsAndConditions"
)

// PossibleLinkTypeValues returns an array of possible values for the LinkType const type.
func PossibleLinkTypeValues() []LinkType {
	return []LinkType{LinkTypeDocumentation, LinkTypeGeneric, LinkTypeKnowMore, LinkTypeSignUp, LinkTypeSpecification, LinkTypeTermsAndConditions}
}

// MeteringType enumerates the values for metering type.
type MeteringType string

const (
	// MeteringTypeAdhoc Adhoc billing.
	MeteringTypeAdhoc MeteringType = "Adhoc"
	// MeteringTypeOneTime One time billing.
	MeteringTypeOneTime MeteringType = "OneTime"
	// MeteringTypeRecurring Recurring billing.
	MeteringTypeRecurring MeteringType = "Recurring"
)

// PossibleMeteringTypeValues returns an array of possible values for the MeteringType const type.
func PossibleMeteringTypeValues() []MeteringType {
	return []MeteringType{MeteringTypeAdhoc, MeteringTypeOneTime, MeteringTypeRecurring}
}

// NotificationStageName enumerates the values for notification stage name.
type NotificationStageName string

const (
	// NotificationStageNameDelivered Notification at order item delivered to customer.
	NotificationStageNameDelivered NotificationStageName = "Delivered"
	// NotificationStageNameShipped Notification at order item shipped from microsoft datacenter.
	NotificationStageNameShipped NotificationStageName = "Shipped"
)

// PossibleNotificationStageNameValues returns an array of possible values for the NotificationStageName const type.
func PossibleNotificationStageNameValues() []NotificationStageName {
	return []NotificationStageName{NotificationStageNameDelivered, NotificationStageNameShipped}
}

// OrderItemCancellationEnum enumerates the values for order item cancellation enum.
type OrderItemCancellationEnum string

const (
	// OrderItemCancellationEnumCancellable Order item can be cancelled without fee.
	OrderItemCancellationEnumCancellable OrderItemCancellationEnum = "Cancellable"
	// OrderItemCancellationEnumCancellableWithFee Order item can be cancelled with fee.
	OrderItemCancellationEnumCancellableWithFee OrderItemCancellationEnum = "CancellableWithFee"
	// OrderItemCancellationEnumNotCancellable Order item not cancellable.
	OrderItemCancellationEnumNotCancellable OrderItemCancellationEnum = "NotCancellable"
)

// PossibleOrderItemCancellationEnumValues returns an array of possible values for the OrderItemCancellationEnum const type.
func PossibleOrderItemCancellationEnumValues() []OrderItemCancellationEnum {
	return []OrderItemCancellationEnum{OrderItemCancellationEnumCancellable, OrderItemCancellationEnumCancellableWithFee, OrderItemCancellationEnumNotCancellable}
}

// OrderItemReturnEnum enumerates the values for order item return enum.
type OrderItemReturnEnum string

const (
	// OrderItemReturnEnumNotReturnable Order item not returnable.
	OrderItemReturnEnumNotReturnable OrderItemReturnEnum = "NotReturnable"
	// OrderItemReturnEnumReturnable Order item can be returned without fee.
	OrderItemReturnEnumReturnable OrderItemReturnEnum = "Returnable"
	// OrderItemReturnEnumReturnableWithFee Order item can be returned with fee.
	OrderItemReturnEnumReturnableWithFee OrderItemReturnEnum = "ReturnableWithFee"
)

// PossibleOrderItemReturnEnumValues returns an array of possible values for the OrderItemReturnEnum const type.
func PossibleOrderItemReturnEnumValues() []OrderItemReturnEnum {
	return []OrderItemReturnEnum{OrderItemReturnEnumNotReturnable, OrderItemReturnEnumReturnable, OrderItemReturnEnumReturnableWithFee}
}

// OrderItemType enumerates the values for order item type.
type OrderItemType string

const (
	// OrderItemTypePurchase Purchase OrderItem.
	OrderItemTypePurchase OrderItemType = "Purchase"
	// OrderItemTypeRental Rental OrderItem.
	OrderItemTypeRental OrderItemType = "Rental"
)

// PossibleOrderItemTypeValues returns an array of possible values for the OrderItemType const type.
func PossibleOrderItemTypeValues() []OrderItemType {
	return []OrderItemType{OrderItemTypePurchase, OrderItemTypeRental}
}

// Origin enumerates the values for origin.
type Origin string

const (
	// OriginSystem ...
	OriginSystem Origin = "system"
	// OriginUser ...
	OriginUser Origin = "user"
	// OriginUsersystem ...
	OriginUsersystem Origin = "user,system"
)

// PossibleOriginValues returns an array of possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{OriginSystem, OriginUser, OriginUsersystem}
}

// StageName enumerates the values for stage name.
type StageName string

const (
	// StageNameCancelled Order has been cancelled.
	StageNameCancelled StageName = "Cancelled"
	// StageNameConfirmed Order is confirmed
	StageNameConfirmed StageName = "Confirmed"
	// StageNameDelivered Order is delivered to customer
	StageNameDelivered StageName = "Delivered"
	// StageNameInReview Order is currently in draft mode and can still be cancelled
	StageNameInReview StageName = "InReview"
	// StageNameInUse Order is in use at customer site
	StageNameInUse StageName = "InUse"
	// StageNamePlaced Currently in draft mode and can still be cancelled
	StageNamePlaced StageName = "Placed"
	// StageNameReadyToShip Order is ready to ship
	StageNameReadyToShip StageName = "ReadyToShip"
	// StageNameReturnCompleted Return has now completed.
	StageNameReturnCompleted StageName = "ReturnCompleted"
	// StageNameReturnedToMicrosoft Order has been received back to microsoft.
	StageNameReturnedToMicrosoft StageName = "ReturnedToMicrosoft"
	// StageNameReturnInitiated Return has been initiated by customer.
	StageNameReturnInitiated StageName = "ReturnInitiated"
	// StageNameReturnPickedUp Order is in transit from customer to microsoft.
	StageNameReturnPickedUp StageName = "ReturnPickedUp"
	// StageNameShipped Order is in transit to customer
	StageNameShipped StageName = "Shipped"
)

// PossibleStageNameValues returns an array of possible values for the StageName const type.
func PossibleStageNameValues() []StageName {
	return []StageName{StageNameCancelled, StageNameConfirmed, StageNameDelivered, StageNameInReview, StageNameInUse, StageNamePlaced, StageNameReadyToShip, StageNameReturnCompleted, StageNameReturnedToMicrosoft, StageNameReturnInitiated, StageNameReturnPickedUp, StageNameShipped}
}

// StageStatus enumerates the values for stage status.
type StageStatus string

const (
	// StageStatusCancelled Stage has been cancelled.
	StageStatusCancelled StageStatus = "Cancelled"
	// StageStatusCancelling Stage is cancelling.
	StageStatusCancelling StageStatus = "Cancelling"
	// StageStatusFailed Stage has failed.
	StageStatusFailed StageStatus = "Failed"
	// StageStatusInProgress Stage is in progress.
	StageStatusInProgress StageStatus = "InProgress"
	// StageStatusNone No status available yet.
	StageStatusNone StageStatus = "None"
	// StageStatusSucceeded Stage has succeeded.
	StageStatusSucceeded StageStatus = "Succeeded"
)

// PossibleStageStatusValues returns an array of possible values for the StageStatus const type.
func PossibleStageStatusValues() []StageStatus {
	return []StageStatus{StageStatusCancelled, StageStatusCancelling, StageStatusFailed, StageStatusInProgress, StageStatusNone, StageStatusSucceeded}
}

// SupportedFilterTypes enumerates the values for supported filter types.
type SupportedFilterTypes string

const (
	// SupportedFilterTypesDoubleEncryptionStatus Double encryption status
	SupportedFilterTypesDoubleEncryptionStatus SupportedFilterTypes = "DoubleEncryptionStatus"
	// SupportedFilterTypesShipToCountries Ship to country
	SupportedFilterTypesShipToCountries SupportedFilterTypes = "ShipToCountries"
)

// PossibleSupportedFilterTypesValues returns an array of possible values for the SupportedFilterTypes const type.
func PossibleSupportedFilterTypesValues() []SupportedFilterTypes {
	return []SupportedFilterTypes{SupportedFilterTypesDoubleEncryptionStatus, SupportedFilterTypesShipToCountries}
}

// TransportShipmentTypes enumerates the values for transport shipment types.
type TransportShipmentTypes string

const (
	// TransportShipmentTypesCustomerManaged Shipment Logistics is handled by the customer.
	TransportShipmentTypesCustomerManaged TransportShipmentTypes = "CustomerManaged"
	// TransportShipmentTypesMicrosoftManaged Shipment Logistics is handled by Microsoft.
	TransportShipmentTypesMicrosoftManaged TransportShipmentTypes = "MicrosoftManaged"
)

// PossibleTransportShipmentTypesValues returns an array of possible values for the TransportShipmentTypes const type.
func PossibleTransportShipmentTypesValues() []TransportShipmentTypes {
	return []TransportShipmentTypes{TransportShipmentTypesCustomerManaged, TransportShipmentTypesMicrosoftManaged}
}

// WeightMeasurementUnit enumerates the values for weight measurement unit.
type WeightMeasurementUnit string

const (
	// WeightMeasurementUnitKGS Kilograms.
	WeightMeasurementUnitKGS WeightMeasurementUnit = "KGS"
	// WeightMeasurementUnitLBS Pounds.
	WeightMeasurementUnitLBS WeightMeasurementUnit = "LBS"
)

// PossibleWeightMeasurementUnitValues returns an array of possible values for the WeightMeasurementUnit const type.
func PossibleWeightMeasurementUnitValues() []WeightMeasurementUnit {
	return []WeightMeasurementUnit{WeightMeasurementUnitKGS, WeightMeasurementUnitLBS}
}