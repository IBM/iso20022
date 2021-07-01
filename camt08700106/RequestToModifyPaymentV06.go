package camt08700106

import (
	"encoding/xml"
)

type Message struct{
	XMLName xml.Name `xml:"urn:worldwire Message"`
	Head  *BusinessApplicationHeaderV01 `xml:"AppHdr" json:"head"`
	Body *RequestToModifyPaymentV06 `xml:"ReqToModfyPmt" json:"body"`
}

func (d *Message) AddMessage() *RequestToModifyPaymentV06 {
	d.Body = new(RequestToModifyPaymentV06)
	return d.Body
}

// Scope
// The RequestToModifyPayment message is sent by a case creator/case assigner to a case assignee.
// This message is used to request the modification of characteristics of an original payment instruction.
// Usage
// The RequestToModifyPayment message must be answered with a:
// - ResolutionOfInvestigation message with a positive final outcome when the case assignee can perform the requested modification;
// - ResolutionOfInvestigation message with a negative final outcome when the case assignee may perform the requested modification but fails to do so (too late, irrevocable instruction, one requested element cannot be modified);
// - RejectCaseAssignment message when the case assignee is unable or not authorised to perform the requested modification;
// - NotificationOfCaseAssignment message to indicate whether the case assignee will take on the case himself or reassign the case to a subsequent party in the payment processing chain.
// The RequestToModifyPayment message covers one and only one original instruction at a time. If several original payment instructions need to be modified, then multiple RequestToModifyPayment messages must be sent.
// The RequestToModifyPayment message can be sent to request the modification of one or several elements of the original payment instruction. If many elements need to be modified, it is recommended to cancel the original payment instruction and initiate a new one.
// The RequestToModifyPayment must be processed on an all or nothing basis. If one of the elements to be modified cannot be altered, the assignment must be rejected in full by means of a negative ResolutionOfInvestigation message. (See section on ResolutionOfInvestigation for more details).
// The RequestToModifyPayment message must never be sent to request the modification of the currency of the original payment instruction. If the currency is wrong, use a CustomerPaymentCancellationRequest or a FIToFIPaymentCancellationRequest message to cancel it and issue and a new payment instruction.
// The RequestToModifyPayment message may be forwarded to subsequent case assignee(s).
// When a RequestToModifyPayment message is used to decrease the amount of the original payment instruction, the modification will trigger a return of funds from the case assignee to the case creator. The assignee may indicate, within the ResolutionOfInvestigation message, the amount to be returned, the date it is or will be returned and the channel through which the return will be done.
// The RequestToModifyPayment message must never be sent to request the increase of the amount of the original payment instruction. To increase the amount in a payment, the debtor can do one of the following:
// - Cancel the first payment using a CustomerPaymentCancellationRequest or a FIToFIPaymentCancellationRequest message and make a new payment with a higher and correct amount;
// - Simply send a second payment with the supplementary amount.
// Depending on the requested modification(s) and the processing stage of the original payment instruction, the processing of a request to modify payment case may end with one of the following:
// - an AdditionalPaymentInformation message sent to the creditor of the original payment instruction;
// - a DebitAuthorisationRequest message sent to the creditor of the original payment instruction;
// - a CustomerPaymentCancellationRequest or a FIToFIPaymentCancellationRequest message sent to a subsequent case assignee.
// The RequestToModifyPayment message can be sent to correct characteristics of an original payment instruction following receipt of an UnableToApply message. In this scenario, the case identification will remain the same.
// The RequestToModifyPayment message has the following main characteristics:
// The case creator assigns a unique case identification. This information will be passed unchanged to all subsequent case assignee(s).
// Lowering the amount of an original payment instruction for which cover is provided by a separate instruction will systematically mean the modification of the whole transaction, including the cover. The case assignee performing the amount modification must initiate the return of funds in excess to the case creator.
// The modification of the agent's or agents' information on an original payment instruction for which cover is provided by a separate instruction will systematically mean the whole transaction is modified, that is the cover is executed through the agent(s) mentioned in the RequestToModifyPayment message. The cover payment must not be modified separately.
// The modification of a payment instruction can be initiated by either the debtor or any subsequent agent in the payment processing chain.
// The case creator provides the information to be modified in line with agreements made with the case assignee. If the case assignee needs in turn to assign the case to a subsequent case assignee, the requested modification(s) must be in line with the agreement made with the next case assignee and a NotificationOfCaseAssignment message must be sent to the case assigner. Otherwise, the request to modify payment case must be rejected (by means of a negative ResolutionOfInvestigation message).
type RequestToModifyPaymentV06 struct {
	XMLName xml.Name
	
		// Identifies the assignment of an investigation case from an assigner to an assignee.
	// Usage: The assigner must be the sender of this confirmation and the assignee must be the receiver.
	Assgnmt *CaseAssignment5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.06 Assgnmt" json:"assgnmt" `
	
		// Identifies the investigation case.
	Case *Case5 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.06 Case,omitempty" json:"case" `
	
		// Identifies the payment transaction to be modified.
	Undrlyg *UnderlyingTransaction5Choice `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.06 Undrlyg" json:"undrlyg" `
	
		// Identifies the list of modifications requested.
	Mod *RequestedModification8 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.06 Mod" json:"mod" `
	
		// Further information related to the processing of the investigation that may need to be acted upon by the assignee.
	InstrForAssgne *InstructionForAssignee1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.06 InstrForAssgne,omitempty" json:"instr_for_assgne" `
	
		// Additional information that cannot be captured in the structured elements and/or any other specific block.
	SplmtryData []*SupplementaryData1 `xml:"urn:iso:std:iso:20022:tech:xsd:camt.087.001.06 SplmtryData,omitempty" json:"splmtry_data" `
	
}


func (r *RequestToModifyPaymentV06) AddAssgnmt() *CaseAssignment5 {
	r.Assgnmt = new(CaseAssignment5)
	return r.Assgnmt
}

func (r *RequestToModifyPaymentV06) AddCase() *Case5 {
	r.Case = new(Case5)
	return r.Case
}

func (r *RequestToModifyPaymentV06) AddUndrlyg() *UnderlyingTransaction5Choice {
	r.Undrlyg = new(UnderlyingTransaction5Choice)
	return r.Undrlyg
}

func (r *RequestToModifyPaymentV06) AddMod() *RequestedModification8 {
	r.Mod = new(RequestedModification8)
	return r.Mod
}

func (r *RequestToModifyPaymentV06) AddInstrForAssgne() *InstructionForAssignee1 {
	r.InstrForAssgne = new(InstructionForAssignee1)
	return r.InstrForAssgne
}

func (r *RequestToModifyPaymentV06) AddSplmtryData() *SupplementaryData1 {
	newValue := new (SupplementaryData1)
	r.SplmtryData = append(r.SplmtryData, newValue)
	return newValue
}

