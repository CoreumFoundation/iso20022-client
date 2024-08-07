package server

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/CoreumFoundation/iso20022-client/iso20022/docs"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

type Handler struct {
	Parser       processes.Parser
	MessageQueue processes.MessageQueue
}

// Send godoc
// @Summary Send a message
// @Description Send an ISO20022 message
// @Tags Messaging
// @Accept application/xml
// @Produce json
// @Router /send [post]
// @Param body body string true "ISO20022 message in xml format" SchemaExample(<?xml version="1.0" encoding="UTF-8"?>\r\n<FIToFICstmrCdtTrf>\r\n\t<GrpHdr>\r\n\t\t<MsgId>BBBB/150928-CT/EUR/912</MsgId>\r\n\t\t<CreDtTm>2015-09-28T16:01:00</CreDtTm>\r\n\t\t<NbOfTxs>2</NbOfTxs>\r\n\t\t<TtlIntrBkSttlmAmt Ccy="EUR">504500</TtlIntrBkSttlmAmt>\r\n\t\t<IntrBkSttlmDt>2015-09-29</IntrBkSttlmDt>\r\n\t\t<SttlmInf>\r\n\t\t\t<SttlmMtd>INDA</SttlmMtd>\r\n\t\t\t<SttlmAcct>\r\n\t\t\t\t<Id>\r\n\t\t\t\t\t<Othr>\r\n\t\t\t\t\t\t<Id>29314569847</Id>\r\n\t\t\t\t\t</Othr>\r\n\t\t\t\t</Id>\r\n\t\t\t</SttlmAcct>\r\n\t\t</SttlmInf>\r\n\t\t<InstgAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>BBBBUS33</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</InstgAgt>\r\n\t\t<InstdAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>EEEEDEFF</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</InstdAgt>\r\n\t</GrpHdr>\r\n\t<CdtTrfTxInf>\r\n\t\t<PmtId>\r\n\t\t\t<InstrId>BBBB/150928-CCT/EUR/912/1</InstrId>\r\n\t\t\t<EndToEndId>ABC/ABC-13679/2015-09-15</EndToEndId>\r\n\t\t\t<TxId>BBBB/150928-CCT/EUR/912/1</TxId>\r\n\t\t</PmtId>\r\n\t\t<PmtTpInf>\r\n\t\t\t<InstrPrty>NORM</InstrPrty>\r\n\t\t</PmtTpInf>\r\n\t\t<IntrBkSttlmAmt Ccy="EUR">499500</IntrBkSttlmAmt>\r\n\t\t<InstdAmt Ccy="EUR">500000</InstdAmt>\r\n\t\t<ChrgBr>CRED</ChrgBr>\r\n\t\t<ChrgsInf>\r\n\t\t\t<Amt Ccy="EUR">500</Amt>\r\n\t\t\t<Agt>\r\n\t\t\t\t<FinInstnId>\r\n\t\t\t\t\t<BICFI>BBBBUS33</BICFI>\r\n\t\t\t\t</FinInstnId>\r\n\t\t\t</Agt>\r\n\t\t</ChrgsInf>\r\n\t\t<Dbtr>\r\n\t\t\t<Nm>ABC Corporation</Nm>\r\n\t\t\t<PstlAdr>\r\n\t\t\t\t<StrtNm>Times Square</StrtNm>\r\n\t\t\t\t<BldgNb>7</BldgNb>\r\n\t\t\t\t<PstCd>NY 10036</PstCd>\r\n\t\t\t\t<TwnNm>New York</TwnNm>\r\n\t\t\t\t<Ctry>US</Ctry>\r\n\t\t\t</PstlAdr>\r\n\t\t</Dbtr>\r\n\t\t<DbtrAcct>\r\n\t\t\t<Id>\r\n\t\t\t\t<Othr>\r\n\t\t\t\t\t<Id>00125574999</Id>\r\n\t\t\t\t</Othr>\r\n\t\t\t</Id>\r\n\t\t</DbtrAcct>\r\n\t\t<DbtrAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>BBBBUS33</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</DbtrAgt>\r\n\t\t<CdtrAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>DDDDBEBB</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</CdtrAgt>\r\n\t\t<Cdtr>\r\n\t\t\t<Nm>GHI Semiconductors</Nm>\r\n\t\t\t<PstlAdr>\r\n\t\t\t\t<StrtNm>Avenue Brugmann</StrtNm>\r\n\t\t\t\t<BldgNb>415</BldgNb>\r\n\t\t\t\t<PstCd>1180</PstCd>\r\n\t\t\t\t<TwnNm>Brussels</TwnNm>\r\n\t\t\t\t<Ctry>BE</Ctry>\r\n\t\t\t</PstlAdr>\r\n\t\t</Cdtr>\r\n\t\t<CdtrAcct>\r\n\t\t\t<Id>\r\n\t\t\t\t<IBAN>BE30001216371411</IBAN>\r\n\t\t\t</Id>\r\n\t\t</CdtrAcct>\r\n\t\t<Purp>\r\n\t\t\t<Cd>GDDS</Cd>\r\n\t\t</Purp>\r\n\t\t<RmtInf>\r\n\t\t\t<Strd>\r\n\t\t\t\t<RfrdDocInf>\r\n\t\t\t\t\t<Tp>\r\n\t\t\t\t\t\t<CdOrPrtry>\r\n\t\t\t\t\t\t\t<Cd>CINV</Cd>\r\n\t\t\t\t\t\t</CdOrPrtry>\r\n\t\t\t\t\t</Tp>\r\n\t\t\t\t\t<Nb>ABC-13679</Nb>\r\n\t\t\t\t\t<RltdDt>\r\n\t\t\t\t\t\t<Tp>\r\n\t\t\t\t\t\t\t<Cd>INDA</Cd>\r\n\t\t\t\t\t\t</Tp>\r\n\t\t\t\t\t\t<Dt>2015-09-08</Dt>\r\n\t\t\t\t\t</RltdDt>\r\n\t\t\t\t</RfrdDocInf>\r\n\t\t\t</Strd>\r\n\t\t</RmtInf>\r\n\t</CdtTrfTxInf>\r\n\t<CdtTrfTxInf>\r\n\t\t<PmtId>\r\n\t\t\t<InstrId>BBBB/150928-CCT/EUR/912/2</InstrId>\r\n\t\t\t<EndToEndId>BBBB/150928-ZZ/JO/164794</EndToEndId>\r\n\t\t\t<TxId>BBBB/150928-CCT/EUR/912/2</TxId>\r\n\t\t</PmtId>\r\n\t\t<PmtTpInf>\r\n\t\t\t<InstrPrty>NORM</InstrPrty>\r\n\t\t</PmtTpInf>\r\n\t\t<IntrBkSttlmAmt Ccy="EUR">5000</IntrBkSttlmAmt>\r\n\t\t<ChrgBr>SHAR</ChrgBr>\r\n\t\t<Dbtr>\r\n\t\t\t<Nm>Mr. Jones</Nm>\r\n\t\t\t<PstlAdr>\r\n\t\t\t\t<StrtNm>16th Street</StrtNm>\r\n\t\t\t\t<BldgNb>30</BldgNb>\r\n\t\t\t\t<PstCd>NY10023</PstCd>\r\n\t\t\t\t<TwnNm>New York</TwnNm>\r\n\t\t\t\t<Ctry>US</Ctry>\r\n\t\t\t</PstlAdr>\r\n\t\t</Dbtr>\r\n\t\t<DbtrAcct>\r\n\t\t\t<Id>\r\n\t\t\t\t<Othr>\r\n\t\t\t\t\t<Id>00125583145</Id>\r\n\t\t\t\t</Othr>\r\n\t\t\t</Id>\r\n\t\t</DbtrAcct>\r\n\t\t<DbtrAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>BBBBUS33</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</DbtrAgt>\r\n\t\t<CdtrAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>EEEEDEFF</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</CdtrAgt>\r\n\t\t<Cdtr>\r\n\t\t\t<Nm>ZZ Insurances</Nm>\r\n\t\t\t<PstlAdr>\r\n\t\t\t\t<StrtNm>Friedrich-Ebert-Anlage</StrtNm>\r\n\t\t\t\t<BldgNb>2-14</BldgNb>\r\n\t\t\t\t<PstCd>D-60 325</PstCd>\r\n\t\t\t\t<TwnNm>Frankfurt am Main</TwnNm>\r\n\t\t\t\t<Ctry>DE</Ctry>\r\n\t\t\t\t<AdrLine>City Haus 1 10th Floor</AdrLine>\r\n\t\t\t</PstlAdr>\r\n\t\t</Cdtr>\r\n\t\t<CdtrAcct>\r\n\t\t\t<Id>\r\n\t\t\t\t<IBAN>DE89370400440532014000</IBAN>\r\n\t\t\t</Id>\r\n\t\t</CdtrAcct>\r\n\t\t<RmtInf>\r\n\t\t\t<Ustrd>Contract ZZ/JO/164794</Ustrd>\r\n\t\t</RmtInf>\r\n\t</CdtTrfTxInf>\r\n</FIToFICstmrCdtTrf>\r\n)
// @Success 201 {object} StandardResponse{data=MessageSendResponse}
// @Failure 400 {object} StandardResponse{message=string}
// @Failure 500 "Something bad happened"
func (h *Handler) Send(c *gin.Context) {
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(c.Request.Body)

	message, err := io.ReadAll(c.Request.Body)
	if err != nil {
		resp := GetFailResponseFromErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	metadata, err := h.Parser.ExtractMetadataFromIsoMessage(message)
	if err != nil {
		resp := GetFailResponseFromErrors(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, resp)
		return
	}

	// TODO: Check for duplicate messages by ID
	fmt.Printf("Got message with ID : %s\n", metadata.ID)

	resp := GetSuccessResponse(MessageSendResponse{MessageID: metadata.ID})
	c.JSON(http.StatusCreated, resp)

	go h.MessageQueue.PushToSend(metadata.ID, message)
}

// Receive godoc
// @Summary Receive a message
// @Description Tries to receive an ISO20022 message if there is any
// @Tags Messaging
// @Produce application/xml
// @Router /receive [get]
// @x-code-samples [{"lang":"xml","source":"<?xml version=\"1.0\" encoding=\"UTF-8\"?>\r\n<FIToFICstmrCdtTrf>\r\n\t<GrpHdr>\r\n\t\t<MsgId>BBBB/150928-CT/EUR/912</MsgId>\r\n\t\t<CreDtTm>2015-09-28T16:01:00</CreDtTm>\r\n\t\t<NbOfTxs>2</NbOfTxs>\r\n\t\t<TtlIntrBkSttlmAmt Ccy=\"EUR\">504500</TtlIntrBkSttlmAmt>\r\n\t\t<IntrBkSttlmDt>2015-09-29</IntrBkSttlmDt>\r\n\t\t<SttlmInf>\r\n\t\t\t<SttlmMtd>INDA</SttlmMtd>\r\n\t\t\t<SttlmAcct>\r\n\t\t\t\t<Id>\r\n\t\t\t\t\t<Othr>\r\n\t\t\t\t\t\t<Id>29314569847</Id>\r\n\t\t\t\t\t</Othr>\r\n\t\t\t\t</Id>\r\n\t\t\t</SttlmAcct>\r\n\t\t</SttlmInf>\r\n\t\t<InstgAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>BBBBUS33</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</InstgAgt>\r\n\t\t<InstdAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>EEEEDEFF</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</InstdAgt>\r\n\t</GrpHdr>\r\n\t<CdtTrfTxInf>\r\n\t\t<PmtId>\r\n\t\t\t<InstrId>BBBB/150928-CCT/EUR/912/1</InstrId>\r\n\t\t\t<EndToEndId>ABC/ABC-13679/2015-09-15</EndToEndId>\r\n\t\t\t<TxId>BBBB/150928-CCT/EUR/912/1</TxId>\r\n\t\t</PmtId>\r\n\t\t<PmtTpInf>\r\n\t\t\t<InstrPrty>NORM</InstrPrty>\r\n\t\t</PmtTpInf>\r\n\t\t<IntrBkSttlmAmt Ccy=\"EUR\">499500</IntrBkSttlmAmt>\r\n\t\t<InstdAmt Ccy=\"EUR\">500000</InstdAmt>\r\n\t\t<ChrgBr>CRED</ChrgBr>\r\n\t\t<ChrgsInf>\r\n\t\t\t<Amt Ccy=\"EUR\">500</Amt>\r\n\t\t\t<Agt>\r\n\t\t\t\t<FinInstnId>\r\n\t\t\t\t\t<BICFI>BBBBUS33</BICFI>\r\n\t\t\t\t</FinInstnId>\r\n\t\t\t</Agt>\r\n\t\t</ChrgsInf>\r\n\t\t<Dbtr>\r\n\t\t\t<Nm>ABC Corporation</Nm>\r\n\t\t\t<PstlAdr>\r\n\t\t\t\t<StrtNm>Times Square</StrtNm>\r\n\t\t\t\t<BldgNb>7</BldgNb>\r\n\t\t\t\t<PstCd>NY 10036</PstCd>\r\n\t\t\t\t<TwnNm>New York</TwnNm>\r\n\t\t\t\t<Ctry>US</Ctry>\r\n\t\t\t</PstlAdr>\r\n\t\t</Dbtr>\r\n\t\t<DbtrAcct>\r\n\t\t\t<Id>\r\n\t\t\t\t<Othr>\r\n\t\t\t\t\t<Id>00125574999</Id>\r\n\t\t\t\t</Othr>\r\n\t\t\t</Id>\r\n\t\t</DbtrAcct>\r\n\t\t<DbtrAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>BBBBUS33</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</DbtrAgt>\r\n\t\t<CdtrAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>DDDDBEBB</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</CdtrAgt>\r\n\t\t<Cdtr>\r\n\t\t\t<Nm>GHI Semiconductors</Nm>\r\n\t\t\t<PstlAdr>\r\n\t\t\t\t<StrtNm>Avenue Brugmann</StrtNm>\r\n\t\t\t\t<BldgNb>415</BldgNb>\r\n\t\t\t\t<PstCd>1180</PstCd>\r\n\t\t\t\t<TwnNm>Brussels</TwnNm>\r\n\t\t\t\t<Ctry>BE</Ctry>\r\n\t\t\t</PstlAdr>\r\n\t\t</Cdtr>\r\n\t\t<CdtrAcct>\r\n\t\t\t<Id>\r\n\t\t\t\t<IBAN>BE30001216371411</IBAN>\r\n\t\t\t</Id>\r\n\t\t</CdtrAcct>\r\n\t\t<Purp>\r\n\t\t\t<Cd>GDDS</Cd>\r\n\t\t</Purp>\r\n\t\t<RmtInf>\r\n\t\t\t<Strd>\r\n\t\t\t\t<RfrdDocInf>\r\n\t\t\t\t\t<Tp>\r\n\t\t\t\t\t\t<CdOrPrtry>\r\n\t\t\t\t\t\t\t<Cd>CINV</Cd>\r\n\t\t\t\t\t\t</CdOrPrtry>\r\n\t\t\t\t\t</Tp>\r\n\t\t\t\t\t<Nb>ABC-13679</Nb>\r\n\t\t\t\t\t<RltdDt>\r\n\t\t\t\t\t\t<Tp>\r\n\t\t\t\t\t\t\t<Cd>INDA</Cd>\r\n\t\t\t\t\t\t</Tp>\r\n\t\t\t\t\t\t<Dt>2015-09-08</Dt>\r\n\t\t\t\t\t</RltdDt>\r\n\t\t\t\t</RfrdDocInf>\r\n\t\t\t</Strd>\r\n\t\t</RmtInf>\r\n\t</CdtTrfTxInf>\r\n\t<CdtTrfTxInf>\r\n\t\t<PmtId>\r\n\t\t\t<InstrId>BBBB/150928-CCT/EUR/912/2</InstrId>\r\n\t\t\t<EndToEndId>BBBB/150928-ZZ/JO/164794</EndToEndId>\r\n\t\t\t<TxId>BBBB/150928-CCT/EUR/912/2</TxId>\r\n\t\t</PmtId>\r\n\t\t<PmtTpInf>\r\n\t\t\t<InstrPrty>NORM</InstrPrty>\r\n\t\t</PmtTpInf>\r\n\t\t<IntrBkSttlmAmt Ccy=\"EUR\">5000</IntrBkSttlmAmt>\r\n\t\t<ChrgBr>SHAR</ChrgBr>\r\n\t\t<Dbtr>\r\n\t\t\t<Nm>Mr. Jones</Nm>\r\n\t\t\t<PstlAdr>\r\n\t\t\t\t<StrtNm>16th Street</StrtNm>\r\n\t\t\t\t<BldgNb>30</BldgNb>\r\n\t\t\t\t<PstCd>NY10023</PstCd>\r\n\t\t\t\t<TwnNm>New York</TwnNm>\r\n\t\t\t\t<Ctry>US</Ctry>\r\n\t\t\t</PstlAdr>\r\n\t\t</Dbtr>\r\n\t\t<DbtrAcct>\r\n\t\t\t<Id>\r\n\t\t\t\t<Othr>\r\n\t\t\t\t\t<Id>00125583145</Id>\r\n\t\t\t\t</Othr>\r\n\t\t\t</Id>\r\n\t\t</DbtrAcct>\r\n\t\t<DbtrAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>BBBBUS33</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</DbtrAgt>\r\n\t\t<CdtrAgt>\r\n\t\t\t<FinInstnId>\r\n\t\t\t\t<BICFI>EEEEDEFF</BICFI>\r\n\t\t\t</FinInstnId>\r\n\t\t</CdtrAgt>\r\n\t\t<Cdtr>\r\n\t\t\t<Nm>ZZ Insurances</Nm>\r\n\t\t\t<PstlAdr>\r\n\t\t\t\t<StrtNm>Friedrich-Ebert-Anlage</StrtNm>\r\n\t\t\t\t<BldgNb>2-14</BldgNb>\r\n\t\t\t\t<PstCd>D-60 325</PstCd>\r\n\t\t\t\t<TwnNm>Frankfurt am Main</TwnNm>\r\n\t\t\t\t<Ctry>DE</Ctry>\r\n\t\t\t\t<AdrLine>City Haus 1 10th Floor</AdrLine>\r\n\t\t\t</PstlAdr>\r\n\t\t</Cdtr>\r\n\t\t<CdtrAcct>\r\n\t\t\t<Id>\r\n\t\t\t\t<IBAN>DE89370400440532014000</IBAN>\r\n\t\t\t</Id>\r\n\t\t</CdtrAcct>\r\n\t\t<RmtInf>\r\n\t\t\t<Ustrd>Contract ZZ/JO/164794</Ustrd>\r\n\t\t</RmtInf>\r\n\t</CdtTrfTxInf>\r\n</FIToFICstmrCdtTrf>\r\n"}]
// @Success 200
// @Response 204
func (h *Handler) Receive(c *gin.Context) {
	message, ok := h.MessageQueue.PopFromReceive()
	if ok {
		c.Data(http.StatusOK, "application/xml", message)
	} else {
		c.Status(http.StatusNoContent)
	}
}

func (h *Handler) Doc(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", docs.SwaggerTemplate)
}

func (h *Handler) Swagger(c *gin.Context) {
	c.Data(http.StatusOK, "application/json", docs.SwaggerJson)
}
