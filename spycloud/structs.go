package spycloud

type ResponseResults struct {
	Cursor  string         `json:"cursor"`
	Hits    int            `json:"hits"`
	Results []BreachRecord `json:"results"`
}

type BreachRecord struct {
	DocumentId                    string   `json:"document_id"`
	SourceId                      int      `json:"source_id"`
	SpycloudPublishDate           string   `json:"spycloud_publish_date"`
	Severity                      int      `json:"severity"`
	ActiveInvestor                string   `json:"active_investor"`
	Address1                      string   `json:"address_1"`
	Address2                      string   `json:"address_2"`
	Age                           string   `json:"age"`
	Birthplace                    string   `json:"birthplace"`
	BuysOnline                    string   `json:"buys_online"`
	CatOwner                      string   `json:"cat_owner"`
	ChristianFamily               string   `json:"christian_family"`
	City                          string   `json:"city"`
	CompanyName                   string   `json:"company_name"`
	Country                       string   `json:"country"`
	County                        string   `json:"county"`
	CreditRating                  string   `json:"credit_rating"`
	DateOfDeath                   string   `json:"date_of_death"`
	DeviceModel                   string   `json:"device_model"`
	Dob                           string   `json:"dob"`
	DogOwner                      string   `json:"dog_owner"`
	DriversLicense                string   `json:"drivers_license"`
	EcFirstName                   string   `json:"ec_first_name"`
	EcLastName                    string   `json:"ec_last_name"`
	EcPhone                       string   `json:"ec_phone"`
	EcPostalCode                  string   `json:"ec_postal_code"`
	EcRelation                    string   `json:"ec_relation"`
	Education                     string   `json:"education"`
	EstimatedIncome               string   `json:"estimated_income"`
	EthnicGroup                   string   `json:"ethnic_group"`
	Ethnicity                     string   `json:"ethnicity"`
	Fax                           string   `json:"fax"`
	FirstName                     string   `json:"first_name"`
	Gender                        string   `json:"gender"`
	Geolocation                   string   `json:"geolocation"`
	Grandchildren                 string   `json:"grandchildren"`
	HasAirConditioning            string   `json:"has_air_conditioning"`
	HasAmexCard                   string   `json:"has_amex_card"`
	HasCreditCards                string   `json:"has_credit_cards"`
	HasChildren                   string   `json:"has_children"`
	HasDiscoverCard               string   `json:"has_discover_card"`
	HasMasterCard                 string   `json:"has_master_card"`
	HasPets                       string   `json:"has_pets"`
	HasSwimmingPool               string   `json:"has_swimming_pool"`
	HasVisaCard                   string   `json:"has_visa_card"`
	HobbiesAndInterests           string   `json:"hobbies_and_interests"`
	HomePage                      string   `json:"home_page"`
	HomeBuiltYear                 string   `json:"home_built_year"`
	HomePurchaseDate              string   `json:"home_purchase_date"`
	HomePurchasePrice             string   `json:"home_purchase_price"`
	HomeTransactionType           string   `json:"home_transaction_type"`
	HomeValue                     string   `json:"home_value"`
	Industry                      string   `json:"industry"`
	Isp                           string   `json:"isp"`
	InvestmentsPersonal           string   `json:"investments_personal"`
	IPAddresses                   []string `json:"ip_addresses"`
	IsSmoker                      string   `json:"is_smoker"`
	Language                      string   `json:"language"`
	LastName                      string   `json:"last_name"`
	MaritalStatus                 string   `json:"marital_status"`
	MiddleName                    string   `json:"middle_name"`
	MortgageAmount                string   `json:"mortgage_amount"`
	MortgageLenderName            string   `json:"mortgage_lender_name"`
	MortgageLoanType              string   `json:"mortgage_loan_type"`
	MortgageRate                  string   `json:"mortgage_rate"`
	NameSuffix                    string   `json:"name_suffix"`
	NationalId                    string   `json:"national_id"`
	NetWorth                      string   `json:"net_worth"`
	NumberChildren                string   `json:"number_children"`
	PassportCountry               string   `json:"passport_country"`
	PassportExpDate               string   `json:"passport_exp_date"`
	PassportIssueDate             string   `json:"passport_issue_date"`
	PassportNumber                string   `json:"passport_number"`
	PayableTo                     string   `json:"payable_to"`
	Phone                         string   `json:"phone"`
	PoliticalAffiliation          string   `json:"political_affiliation"`
	PostalCode                    string   `json:"postal_code"`
	Religion                      string   `json:"religion"`
	ResidenceLengthYears          string   `json:"residence_length_years"`
	SewerType                     string   `json:"sewer_type"`
	SingleParent                  string   `json:"single_parent"`
	SocialSecurityNumber          string   `json:"social_security_number"`
	SsnLastFour                   string   `json:"ssn_last_four"`
	State                         string   `json:"state"`
	Timezone                      string   `json:"timezone"`
	Title                         string   `json:"title"`
	VehicleIdentificationNumber   string   `json:"vehicle_identification_number"`
	VehicleMake                   string   `json:"vehicle_make"`
	VehicleModel                  string   `json:"vehicle_model"`
	VoterId                       string   `json:"voter_id"`
	VoterRegistrationDate         string   `json:"voter_registration_date"`
	WaterType                     string   `json:"water_type"`
	SocialAim                     []string `json:"social_aim"`
	SocialAboutme                 []string `json:"social_aboutme"`
	SocialAgnellist               []string `json:"social_agnellist"`
	SocialBehance                 []string `json:"social_behance"`
	SocialCrunchbase              []string `json:"social_crunchbase"`
	SocialDribble                 []string `json:"social_dribble"`
	SocialFacebook                []string `json:"social_facebook"`
	SocialFlickr                  []string `json:"social_flickr"`
	SocialFoursquare              []string `json:"social_foursquare"`
	SocialGithub                  []string `json:"social_github"`
	SocialGitlab                  []string `json:"social_gitlab"`
	SocialGravatar                []string `json:"social_gravatar"`
	SocialGoogle                  []string `json:"social_google"`
	SocialIcq                     []string `json:"social_icq"`
	SocialIndeed                  []string `json:"social_indeed"`
	SocialInstagram               []string `json:"social_instagram"`
	SocialKlout                   []string `json:"social_klout"`
	SocialLindin                  []string `json:"social_lindin"`
	SocialMedium                  []string `json:"social_medium"`
	SocialMeetup                  []string `json:"social_meetup"`
	SocialMsn                     []string `json:"social_msn"`
	SocialMyspace                 []string `json:"social_myspace"`
	SocialOther                   []string `json:"social_other"`
	SocialPintrest                []string `json:"social_pintrest"`
	SocialQuora                   []string `json:"social_quora"`
	SocialReddit                  []string `json:"social_reddit"`
	SocialSkype                   []string `json:"social_skype"`
	SocialSoundcloud              []string `json:"social_soundcloud"`
	SocialStackoverflow           []string `json:"social_stackoverflow"`
	SocialSteam                   []string `json:"social_steam"`
	SocialTelegram                []string `json:"social_telegram"`
	SocialTwitter                 []string `json:"social_twitter"`
	SocialViemo                   []string `json:"social_viemo"`
	SocialWeibo                   []string `json:"social_weibo"`
	SocialWhatsapp                []string `json:"social_whatsapp"`
	SocialWordpress               []string `json:"social_wordpress"`
	SocialXing                    []string `json:"social_xing"`
	SocialYahoo                   []string `json:"social_yahoo"`
	SocialYoutube                 []string `json:"social_youtube"`
	BankNumber                    string   `json:"bank_number"`
	CcBin                         string   `json:"cc_bin"`
	CcCode                        string   `json:"cc_code"`
	CcExpiration                  string   `json:"cc_expiration"`
	CcLastFour                    string   `json:"cc_last_four"`
	CcNumber                      string   `json:"cc_number"`
	CcType                        string   `json:"cc_type"`
	Taxid                         string   `json:"taxid"`
	BotnetId                      string   `json:"botnet_id"`
	FromCookiesData               string   `json:"from_cookies_data"`
	FromPostData                  string   `json:"from_post_data"`
	InfectedMachineId             string   `json:"infected_machine_id"`
	InfectedTime                  string   `json:"infected_time"`
	LogonServer                   string   `json:"logon_server"`
	SystemInstallDate             string   `json:"system_install_date"`
	SystemModel                   string   `json:"system_model"`
	TargetDomain                  string   `json:"target_domain"`
	TargetUrl                     string   `json:"target_url"`
	UserAgent                     string   `json:"user_agent"`
	UserBrowser                   string   `json:"user_browser"`
	UserHostname                  string   `json:"user_hostname"`
	UserOs                        string   `json:"user_os"`
	UserSysDomain                 string   `json:"user_sys_domain"`
	UserSysRegisteredOrganization string   `json:"user_sys_registered_organization"`
	UserSysRegisteredOwner        string   `json:"user_sys_registered_owner"`
	AccountCaption                string   `json:"account_caption"`
	AccountImageUrl               string   `json:"account_image_url"`
	AccountLastActivityTime       string   `json:"account_last_activity_time"`
	AccountLoginTime              string   `json:"account_login_time"`
	AccountModificationTime       string   `json:"account_modification_time"`
	AccountNickname               string   `json:"account_nickname"`
	AccountNotes                  string   `json:"account_notes"`
	AccountPasswordDate           string   `json:"account_password_date"`
	AccountSecret                 string   `json:"account_secret"`
	AccountSecretQuestion         string   `json:"account_secret_question"`
	AccountSignupTime             string   `json:"account_signup_time"`
	AccountStatus                 string   `json:"account_status"`
	AccountTitle                  string   `json:"account_title"`
	AccountType                   string   `json:"account_type"`
	BackupEmail                   string   `json:"backup_email"`
	BackupEmailUsername           string   `json:"backup_email_username"`
	Domain                        string   `json:"domain"`
	Email                         string   `json:"email"`
	EmailDomain                   string   `json:"email_domain"`
	EmailUsername                 string   `json:"email_username"`
	NumPosts                      int      `json:"num_posts"`
	Password                      string   `json:"password"`
	PasswordPlaintext             string   `json:"password_plaintext"`
	PasswordType                  string   `json:"password_type"`
	Salt                          string   `json:"salt"`
	Service                       string   `json:"service"`
	ServiceExpiration             string   `json:"service_expiration"`
	Username                      string   `json:"username"`
	EmailStatus                   string   `json:"email_status"`
	CrmLastActivity               string   `json:"crm_last_activity"`
	CrmContactCreated             string   `json:"crm_contact_created"`
	CompanyWebSite                string   `json:"company_web_site"`
	CompanyRevenue                string   `json:"company_revenue"`
	Employees                     string   `json:"employees"`
	JobTitle                      string   `json:"job_title"`
	JonLevel                      string   `json:"jon_level"`
	JobStartDate                  string   `json:"job_start_date"`
	LinkedinNumberConnections     int      `json:"linkedin_number_connections"`
	NacisCode                     string   `json:"nacis_code"`
	SicCode                       string   `json:"sic_code"`
	HealthInsuranceId             string   `json:"health_insurance_id"`
	HealthInsuranceProvider       string   `json:"health_insurance_provider"`
	Desc                          string   `json:"desc"`
	Guid                          string   `json:"guid"`
	PastbinKey                    string   `json:"pastbin_key"`
	RecordAdditionalDate          string   `json:"record_additional_date"`
	RecordCrackedDate             string   `json:"record_cracked_date"`
	RecordModificationDate        string   `json:"record_modification_date"`
	SourceFile                    string   `json:"source_file"`
}
