package treasuryprimesandbox

import "time"

type Campaign struct {
	Candidate    string
	Sweep        float64
	SweepCeiling float64
	Checking     float64
	HighYield    float64

	Contribution []*ContributionRecord
}

type ContributionRecord struct {
	CommitteeID                        string               `csv:"committee_id"`
	CommitteeName                      string               `csv:"committee_name"`
	ReportYear                         int                  `csv:"report_year"`
	ReportType                         string               `csv:"report_type"`
	ImageNumber                        string               `csv:"image_number"`
	LineNumber                         string               `csv:"line_number"`
	TransactionID                      string               `csv:"transaction_id"`
	FileNumber                         int                  `csv:"file_number"`
	EntityType                         string               `csv:"entity_type"`
	EntityTypeDesc                     string               `csv:"entity_type_desc"`
	UnusedContbrID                     string               `csv:"unused_contbr_id"`
	ContributorPrefix                  string               `csv:"contributor_prefix"`
	ContributorName                    string               `csv:"contributor_name"`
	RecipientCommitteeType             string               `csv:"recipient_committee_type"`
	RecipientCommitteeOrgType          string               `csv:"recipient_committee_org_type"`
	RecipientCommitteeDesignation      string               `csv:"recipient_committee_designation"`
	ContributorFirstName               string               `csv:"contributor_first_name"`
	ContributorMiddleName              string               `csv:"contributor_middle_name"`
	ContributorLastName                string               `csv:"contributor_last_name"`
	ContributorSuffix                  string               `csv:"contributor_suffix"`
	ContributorStreet1                 string               `csv:"contributor_street_1"`
	ContributorStreet2                 string               `csv:"contributor_street_2"`
	ContributorCity                    string               `csv:"contributor_city"`
	ContributorState                   string               `csv:"contributor_state"`
	ContributorZip                     string               `csv:"contributor_zip"`
	ContributorEmployer                string               `csv:"contributor_employer"`
	ContributorOccupation              string               `csv:"contributor_occupation"`
	ContributorID                      string               `csv:"contributor_id"`
	ReceiptType                        string               `csv:"receipt_type"`
	ReceiptTypeDesc                    string               `csv:"receipt_type_desc"`
	ReceiptTypeFull                    string               `csv:"receipt_type_full"`
	MemoCode                           string               `csv:"memo_code"`
	MemoCodeFull                       string               `csv:"memo_code_full"`
	ContributionReceiptDate            ContributionDateTime `csv:"contribution_receipt_date"`
	ContributionReceiptAmount          float64              `csv:"contribution_receipt_amount"`
	ContributorAggregateYTD            float64              `csv:"contributor_aggregate_ytd"`
	CandidateID                        string               `csv:"candidate_id"`
	CandidateName                      string               `csv:"candidate_name"`
	CandidateFirstName                 string               `csv:"candidate_first_name"`
	CandidateLastName                  string               `csv:"candidate_last_name"`
	CandidateMiddleName                string               `csv:"candidate_middle_name"`
	CandidatePrefix                    string               `csv:"candidate_prefix"`
	CandidateSuffix                    string               `csv:"candidate_suffix"`
	CandidateOffice                    string               `csv:"candidate_office"`
	CandidateOfficeFull                string               `csv:"candidate_office_full"`
	CandidateOfficeState               string               `csv:"candidate_office_state"`
	CandidateOfficeStateFull           string               `csv:"candidate_office_state_full"`
	CandidateOfficeDistrict            string               `csv:"candidate_office_district"`
	ConduitCommitteeID                 string               `csv:"conduit_committee_id"`
	ConduitCommitteeName               string               `csv:"conduit_committee_name"`
	ConduitCommitteeStreet1            string               `csv:"conduit_committee_street1"`
	ConduitCommitteeStreet2            string               `csv:"conduit_committee_street2"`
	ConduitCommitteeCity               string               `csv:"conduit_committee_city"`
	ConduitCommitteeState              string               `csv:"conduit_committee_state"`
	ConduitCommitteeZip                string               `csv:"conduit_committee_zip"`
	DonorCommitteeName                 string               `csv:"donor_committee_name"`
	NationalCommitteeNonFederalAccount string               `csv:"national_committee_nonfederal_account"`
	ElectionType                       string               `csv:"election_type"`
	ElectionTypeFull                   string               `csv:"election_type_full"`
	FECElectionTypeDesc                string               `csv:"fec_election_type_desc"`
	FECElectionYear                    int                  `csv:"fec_election_year"`
	AmendmentIndicator                 string               `csv:"amendment_indicator"`
	AmendmentIndicatorDesc             string               `csv:"amendment_indicator_desc"`
	ScheduleTypeFull                   string               `csv:"schedule_type_full"`
	LoadDate                           string               `csv:"load_date"`
	OriginalSubID                      string               `csv:"original_sub_id"`
	BackReferenceTransactionID         string               `csv:"back_reference_transaction_id"`
	BackReferenceScheduleName          string               `csv:"back_reference_schedule_name"`
	FilingForm                         string               `csv:"filing_form"`
	LinkID                             string               `csv:"link_id"`
	IsIndividual                       bool                 `csv:"is_individual"`
	MemoText                           string               `csv:"memo_text"`
	TwoYearTransactionPeriod           int                  `csv:"two_year_transaction_period"`
	ScheduleType                       string               `csv:"schedule_type"`
	IncreasedLimit                     bool                 `csv:"increased_limit"`
	SubID                              string               `csv:"sub_id"`
	PDFURL                             string               `csv:"pdf_url"`
	LineNumberLabel                    string               `csv:"line_number_label"`
}

type ContributionDateTime struct {
	time.Time
}

// Convert the internal date as CSV string
func (date *ContributionDateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02"), nil
}

// You could also use the standard Stringer interface
func (date *ContributionDateTime) String() string {
	return date.String() // Redundant, just for example
}

// Convert the CSV string as internal date
func (date *ContributionDateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02", csv)
	return err
}
