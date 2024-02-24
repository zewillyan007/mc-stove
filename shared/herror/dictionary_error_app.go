package herror

var DictionaryErrorApp *ErrorDictionary = &ErrorDictionary{
	Revision:    "1.0.0",
	DigitsClass: 2,
	DigitsError: 4,
	ErrorClass: map[string]*ErrorClass{
		"C01": {
			Type: "Access",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "E-mail not found"},
				"E0002": {Cod: "0002", Msg: "Failed to generate qrcode"},
				"E0003": {Cod: "0003", Msg: "Token not found"},
				"E0004": {Cod: "0004", Msg: "User is blocked"},
				"E0005": {Cod: "0005", Msg: "User is suspended"},
				"E0006": {Cod: "0006", Msg: "User does not have permission for this endpoint and action"},
				"E0007": {Cod: "0007", Msg: "User session does not match token id_session"},
				"E0008": {Cod: "0008", Msg: "User session is invalid"},
				"E0009": {Cod: "0009", Msg: "User session not found"},
				"E0010": {Cod: "0010", Msg: "Not allowed to work on this houer"},
				"E0011": {Cod: "0011", Msg: "Not allowed to work on this weekday"},
				"E0012": {Cod: "0012", Msg: "Failed to create session"},
				"E0013": {Cod: "0013", Msg: "User is not null"},
				"E0014": {Cod: "0014", Msg: "Token is not null"},
				"E0015": {Cod: "0015", Msg: "Password is not null"},
				"E0016": {Cod: "0016", Msg: "Token Unauthorized"},
				"E0017": {Cod: "0017", Msg: "Reactivate Failed"},
				"E0018": {Cod: "0018", Msg: "Registration Expired"},
				"E0019": {Cod: "0019", Msg: "Hash not reported"},
				"E0020": {Cod: "0020", Msg: "Remove access failed"},
				"E0021": {Cod: "0021", Msg: "Documents type not found"},
				"E0022": {Cod: "0022", Msg: "Document Name is not null"},
				"E0023": {Cod: "0023", Msg: "Document Mnemonic is not null"},
				"E0024": {Cod: "0024", Msg: "Document require Date is not null"},
				"E0025": {Cod: "0025", Msg: "Document require Organ is not null"},
				"E0026": {Cod: "0026", Msg: "General Jobs type not found"},
				"E0027": {Cod: "0027", Msg: "Status is not null"},
				"E0028": {Cod: "0028", Msg: "Local TZ type not found"},
				"E0029": {Cod: "0029", Msg: "Local TZ Name cannot be null"},
				"E0030": {Cod: "0030", Msg: "Local TZ Active cannot be null"},
				"E0031": {Cod: "0031", Msg: "Local TZ Creation Date Time cannot be null"},
				"E0032": {Cod: "0032", Msg: "Get menu tree not found"},
				"E0033": {Cod: "0033", Msg: "Model map not found"},
				"E0034": {Cod: "0034", Msg: "Model map Onwer cannot be null"},
				"E0035": {Cod: "0035", Msg: "Model map Entity cannot be null"},
				"E0036": {Cod: "0036", Msg: "Model map Active cannot be null"},
				"E0037": {Cod: "0037", Msg: "Person address not found"},
				"E0038": {Cod: "0038", Msg: "Id Person is not null"},
				"E0039": {Cod: "0039", Msg: "Principal is not null"},
				"E0040": {Cod: "0040", Msg: "Active is not null"},
				"E0041": {Cod: "0041", Msg: "Person note not found"},
				"E0042": {Cod: "0042", Msg: "Id Person is not 0"},
				"E0043": {Cod: "0043", Msg: "Person type not found"},
				"E0044": {Cod: "0044", Msg: "Person type Name is not null"},
				"E0045": {Cod: "0045", Msg: "Person type Mnemonic is not null"},
				"E0046": {Cod: "0046", Msg: "Profile duplicate"},
				"E0047": {Cod: "0047", Msg: "Profile not found"},
				"E0048": {Cod: "0048", Msg: "Profile Name cannot be null"},
				"E0049": {Cod: "0049", Msg: "Profile Active cannot be null"},
				"E0050": {Cod: "0050", Msg: "Create Profile failed"},
				"E0051": {Cod: "0051", Msg: "Save Profile failed"},
				"E0052": {Cod: "0052", Msg: "Remove Profile failed"},
				"E0053": {Cod: "0053", Msg: "Get user permissions failed"},
				"E0054": {Cod: "0054", Msg: "User Get profile permissions failed"},
				"E0055": {Cod: "0055", Msg: "User Get menu tree failed"},
				"E0056": {Cod: "0056", Msg: "User Get menu failed"},
				"E0057": {Cod: "0057", Msg: "User Id Person Type is not null"},
				"E0058": {Cod: "0058", Msg: "User Id user not given"},
				"E0059": {Cod: "0059", Msg: "User type not found"},
				"E0060": {Cod: "0060", Msg: "Filter users movement not found"},
				"E0061": {Cod: "0061", Msg: "Filter users association not found"},
				"E0062": {Cod: "0062", Msg: "User all info not found"},
				"E0063": {Cod: "0063", Msg: "User all last actions not found"},
				"E0064": {Cod: "0064", Msg: "User all not found"},
				"E0065": {Cod: "0065", Msg: "Name is not null"},
				"E0066": {Cod: "0066", Msg: "Create user all failed"},
				"E0067": {Cod: "0067", Msg: "Principal cannot be null"},
				"E0068": {Cod: "0068", Msg: "Email invalid"},
				"E0069": {Cod: "0069", Msg: "Active cannot be null"},
				"E0070": {Cod: "0070", Msg: "Email is already in use"},
				"E0071": {Cod: "0071", Msg: "ShortName is already in use"},
				"E0072": {Cod: "0072", Msg: "Id Person cannot be null"},
				"E0073": {Cod: "0073", Msg: "Active cannot be null"},
				"E0074": {Cod: "0074", Msg: "Document cannot be null"},
				"E0075": {Cod: "0075", Msg: "Document Type cannot be null"},
				"E0076": {Cod: "0076", Msg: "Workschedule not found"},
				"E0077": {Cod: "0077", Msg: "Workschedule Name cannot be null"},
				"E0078": {Cod: "0078", Msg: "Workschedule Day 1 cannot be null"},
				"E0079": {Cod: "0079", Msg: "Workschedule Day 2 cannot be null"},
				"E0080": {Cod: "0080", Msg: "Workschedule Day 3 cannot be null"},
				"E0081": {Cod: "0081", Msg: "Workschedule Day 4 cannot be null"},
				"E0082": {Cod: "0082", Msg: "Workschedule Day 5 cannot be null"},
				"E0083": {Cod: "0083", Msg: "Workschedule Day 6 cannot be null"},
				"E0084": {Cod: "0084", Msg: "Workschedule Day 7 cannot be null"},
				"E0085": {Cod: "0085", Msg: "Workschedule StartTime cannot be null"},
				"E0086": {Cod: "0086", Msg: "Workschedule Active cannot be null"},
				"E0087": {Cod: "0087", Msg: "Workschedule EndTime cannot be null"},
				"E0088": {Cod: "0088", Msg: "Workschedule Creation Date Time cannot be null"},
				"E0089": {Cod: "0089", Msg: "Person Responsibile Id Person is not null"},
				"E0090": {Cod: "0090", Msg: "Person Responsibile Id Person_Legal_Respons is not null"},
				"E0091": {Cod: "0091", Msg: "Person Responsibile Active is not null"},
				"E0092": {Cod: "0092", Msg: "Person Tribute Id Tribute Type is not null"},
				"E0093": {Cod: "0093", Msg: "Person Tribute Id Person is not null"},
				"E0094": {Cod: "0094", Msg: "Person Tribute Active is not null"},
				"E0095": {Cod: "0095", Msg: "Profile Menu Id Profile Type is not null"},
				"E0096": {Cod: "0096", Msg: "Profile Menu menu is not null"},
				"E0097": {Cod: "0097", Msg: "Profile Menu Access permission not found"},
				"E0098": {Cod: "0098", Msg: "User Profile Menu Id Profile Menu cannot be null"},
				"E0099": {Cod: "0099", Msg: "User Profile Menu Id User Type cannot be null"},
				"E0100": {Cod: "0100", Msg: "User Sale Network Id user cannot be null"},
				"E0101": {Cod: "0101", Msg: "User Sale Network Id sale network cannot be null"},
				"E0102": {Cod: "0102", Msg: "Menu Sort is not null"},
				"E0103": {Cod: "0103", Msg: "Menu Name is not null"},
				"E0104": {Cod: "0104", Msg: "Menu Active is not null"},
				"E0105": {Cod: "0105", Msg: "Module Active cannot be null"},
				"E0106": {Cod: "0106", Msg: "User account reactivated"},
				"E0107": {Cod: "0107", Msg: "An error occurred in the duplicate check"},
				"E0108": {Cod: "0108", Msg: "Workschedule's name already exists"},
				"E0109": {Cod: "0109", Msg: "Profile in use"},
				"E0110": {Cod: "0110", Msg: "Scan error"},
				"E0111": {Cod: "0111", Msg: "Get empty permissions failed"},
				"E0112": {Cod: "0112", Msg: "Profile status cannot be null"},
				"E0113": {Cod: "0113", Msg: "Work schedule in use, cannot be removed"},
				"E0114": {Cod: "0114", Msg: "Check date fields - invalid format"},
				"E0115": {Cod: "0115", Msg: "End time must be greater than the start time of the suspension"},
				"E0116": {Cod: "0116", Msg: "Cannot register minors"},
				"E0117": {Cod: "0117", Msg: "Validate birth date (cannot be null for PHYSICAL person type)"},
				"E0118": {Cod: "0118", Msg: "The user is already registered"},
			},
		},
		"C02": {
			Type: "Lottery",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Description cannot be null"},
				"E0002": {Cod: "0002", Msg: "IdEdition cannot be null"},
				"E0003": {Cod: "0003", Msg: "Type cannot be null"},
				"E0004": {Cod: "0004", Msg: "Number cannot be zero"},
				"E0005": {Cod: "0005", Msg: "Description cannot be null"},
				"E0006": {Cod: "0006", Msg: "Short Description cannot be null"},
				"E0007": {Cod: "0007", Msg: "Prorate Award cannot be null"},
				"E0008": {Cod: "0008", Msg: "IdLorrery cannot be null"},
				"E0009": {Cod: "0009", Msg: "Number cannot be null"},
				"E0010": {Cod: "0010", Msg: "DrawsQty cannot be null"},
				"E0011": {Cod: "0011", Msg: "Status cannot be null"},
				"E0012": {Cod: "0012", Msg: "ResultStatus cannot be null"},
				"E0013": {Cod: "0013", Msg: "Redeemable cannot be null"},
				"E0014": {Cod: "0014", Msg: "AuthorizationLink cannot be null"},
				"E0015": {Cod: "0015", Msg: "IdEdition cannot be null"},
				"E0016": {Cod: "0016", Msg: "IdModality cannot be null"},
				"E0017": {Cod: "0017", Msg: "IdEdition cannot be null"},
				"E0018": {Cod: "0018", Msg: "IdProduct cannot be null"},
				"E0019": {Cod: "0019", Msg: "IdSaleNetwork cannot be null"},
				"E0020": {Cod: "0020", Msg: "Status cannot be null"},
				"E0021": {Cod: "0021", Msg: "Lottery can only be linked with product status being RELEASED"},
				"E0022": {Cod: "0022", Msg: "Status new can only be updated to canceled or released"},
				"E0023": {Cod: "0023", Msg: "Status released can only be updated to suspended"},
				"E0024": {Cod: "0024", Msg: "Status suspended can only be updated to released"},
				"E0025": {Cod: "0025", Msg: "Status canceled cannot be updated"},
				"E0026": {Cod: "0026", Msg: "Product Name cannot be null"},
				"E0027": {Cod: "0027", Msg: "Status cannot be null"},
				"E0028": {Cod: "0028", Msg: "Type cannot be null"},
				"E0029": {Cod: "0029", Msg: "Status cannot be edited"},
				"E0030": {Cod: "0030", Msg: "Name cannot be null"},
				"E0031": {Cod: "0031", Msg: "ShortName cannot be null"},
				"E0032": {Cod: "0032", Msg: "Status cannot be null"},
				"E0033": {Cod: "0033", Msg: "MaxAwardValue cannot be null"},
				"E0034": {Cod: "0034", Msg: "IdModality cannot be null"},
				"E0035": {Cod: "0035", Msg: "IdSaleChannel cannot be null"},
				"E0036": {Cod: "0036", Msg: "IdModality cannot be null"},
				"E0037": {Cod: "0037", Msg: "Id Draw cannot be null"},
				"E0038": {Cod: "0038", Msg: "Status can only be updated to reopen"},
				"E0039": {Cod: "0039", Msg: "Ticket not found"},
				"E0040": {Cod: "0040", Msg: "Status new can only be updated to canceled"},
				"E0041": {Cod: "0041", Msg: "Status released can only be updated to suspended"},
				"E0042": {Cod: "0042", Msg: "Status suspended can only be updated to released"},
				"E0043": {Cod: "0043", Msg: "Status cannot be updated"},
				"E0044": {Cod: "0044", Msg: "Status reopen can only be updated to closed"},
				"E0045": {Cod: "0045", Msg: "Edition can be created only if lottery is released"},
				"E0046": {Cod: "0046", Msg: "There is a lottery of the same base with the same name"},
				"E0047": {Cod: "0047", Msg: "informed period less than the edition period"},
				"E0048": {Cod: "0048", Msg: "Max Award Value cannot be null"},
				"E0049": {Cod: "0049", Msg: "Insertion date cannot be earlier than the draw date"},
				"E0050": {Cod: "0050", Msg: "EndDateTime must be after StartDateTime"},
				"E0051": {Cod: "0051", Msg: "DrawDateTime must be after EndDateTime"},
				"E0052": {Cod: "0052", Msg: "Name is already in use"},
				"E0053": {Cod: "0053", Msg: "No company sent"},
				"E0054": {Cod: "0054", Msg: "Status is not closed"},
				"E0055": {Cod: "0055", Msg: "Unexpected error"},
				"E0056": {Cod: "0056", Msg: "The max award value of the prize is not valid"},
				"E0057": {Cod: "0057", Msg: "EndDateTime must be after time now"},
				"E0058": {Cod: "0058", Msg: "Name or ShortName is already in use"},
				"E0059": {Cod: "0059", Msg: "Edition number in use"},
				"E0060": {Cod: "0060", Msg: "Awards cannot be empty"},
			},
		},
		"C03": {
			Type: "Network",
			List: map[string]*ErrorItem{
				//Allocation
				"E0001": {Cod: "0001", Msg: "Name cannot be null"},
				"E0002": {Cod: "0002", Msg: "Initial cannot be null"},
				"E0003": {Cod: "0003", Msg: "Lock cannot be null"},
				"E0004": {Cod: "0004", Msg: "Association cannot be null"},
				"E0005": {Cod: "0005", Msg: "Bond cannot be null"},
				"E0006": {Cod: "0006", Msg: "Client Permission cannot be null"},
				"E0007": {Cod: "0007", Msg: "There is already an allocation as initial"},
				"E0008": {Cod: "0008", Msg: "Terminals need to be in the same allocation, company and regional"},
				//Allocation_History
				"E0009": {Cod: "0009", Msg: "Id terminal cannot be null"},
				"E0010": {Cod: "0010", Msg: "Id allocation cannot be null"},
				//Application
				"E0011": {Cod: "0011", Msg: "Name cannot be null"},
				"E0012": {Cod: "0012", Msg: "ShortName cannot be null"},
				"E0013": {Cod: "0013", Msg: "Title cannot be null"},
				"E0014": {Cod: "0014", Msg: "Description cannot be null"},
				"E0015": {Cod: "0015", Msg: "Status cannot be null"},
				"E0016": {Cod: "0016", Msg: "File in invalid format"},
				"E0017": {Cod: "0017", Msg: "Data cannot be null or invalid informations"},
				//Chip
				"E0018": {Cod: "0018", Msg: "ICCID cannot be null"},
				"E0019": {Cod: "0019", Msg: "Id phone operator cannot be null"},
				"E0020": {Cod: "0020", Msg: "Status cannot be null"},
				"E0021": {Cod: "0021", Msg: "File in invalid format"},
				"E0022": {Cod: "0022", Msg: "File in invalid format, could not delete file"},
				"E0023": {Cod: "0023", Msg: "Phone operator not found. Check that the ICCID entered is correct"},
				"E0024": {Cod: "0024", Msg: "Chip not found"},
				//City
				"E0025": {Cod: "0025", Msg: "Name cannot be null"},
				//Code_Control
				"E0026": {Cod: "0026", Msg: "Code cannot be null"},
				"E0027": {Cod: "0027", Msg: "Status cannot be null"},
				//Invoice
				"E0028": {Cod: "0028", Msg: "Invoice number cannot be null"},
				"E0029": {Cod: "0029", Msg: "Warranty days cannot be null"},
				//Movement
				"E0030": {Cod: "0030", Msg: "Transition cannot be null"},
				"E0031": {Cod: "0031", Msg: "Allocation initial cannot be null"},
				"E0032": {Cod: "0032", Msg: "Allocation final cannot be null"},
				"E0033": {Cod: "0033", Msg: "User cannot be null"},
				//Phone operator
				"E0034": {Cod: "0034", Msg: "Name cannot be null"},
				"E0035": {Cod: "0035", Msg: "Status cannot be null"},
				//Sale Network detail
				"E0036": {Cod: "0036", Msg: "Allow Reopen cannot be null"},
				"E0037": {Cod: "0037", Msg: "Note cannot be null"},
				//Sale Network
				"E0038": {Cod: "0038", Msg: "Id Hierarchical Structure cannot be null"},
				"E0039": {Cod: "0039", Msg: "Id Person cannot be null"},
				"E0040": {Cod: "0040", Msg: "Code cannot be null"},
				"E0041": {Cod: "0041", Msg: "Status cannot be null"},
				"E0042": {Cod: "0042", Msg: "No parent was found for the given id"},
				"E0043": {Cod: "0043", Msg: "Incorrect current passwords"},
				"E0044": {Cod: "0044", Msg: "Not move your Regional"},
				"E0045": {Cod: "0045", Msg: "Accounting Limit: Incorrect Value"},
				"E0046": {Cod: "0046", Msg: "Activity Start Time: format incorrect or out range"},
				"E0047": {Cod: "0047", Msg: "Activity End Time: format incorrect or out range"},
				"E0048": {Cod: "0048", Msg: "Preferred Collection Time: Format Incorrect or Out Range"},
				"E0049": {Cod: "0049", Msg: "Name is already in use"},
				"E0050": {Cod: "0050", Msg: "Short Name is already in use"},
				"E0051": {Cod: "0051", Msg: "Social Name is already in use"},
				"E0052": {Cod: "0052", Msg: "Id Parent cannot be null"},
				//Sale network relationship
				"E0053": {Cod: "0053", Msg: "Social Name is already in use"},
				"E0054": {Cod: "0054", Msg: "Id Parent cannot be null"},
				//Software profile
				"E0055": {Cod: "0055", Msg: "Name cannot be null"},
				"E0056": {Cod: "0056", Msg: "Active cannot be null"},
				"E0057": {Cod: "0057", Msg: "Software Profile cannot be null in Company"},
				//State
				"E0058": {Cod: "0058", Msg: "Name cannot be null"},
				//Terminal detail
				"E0059": {Cod: "0059", Msg: "Administrative cannot be null"},
				"E0060": {Cod: "0060", Msg: "Allows Training Mode cannot be null"},
				"E0061": {Cod: "0061", Msg: "Allows Delay Clock cannot be null"},
				//Terminal
				"E0062": {Cod: "0062", Msg: "Code cannot be null"},
				"E0063": {Cod: "0063", Msg: "Locked cannot be null"},
				"E0064": {Cod: "0064", Msg: "Code field Cannot Be Changed"},
				"E0065": {Cod: "0065", Msg: "Company field Cannot Be Changed"},
				"E0066": {Cod: "0066", Msg: "Regional field Cannot Be Changed"},
				"E0067": {Cod: "0067", Msg: "Allocation field Cannot Be Changed"},
				"E0068": {Cod: "0068", Msg: "Model field Cannot Be Changed"},
				"E0069": {Cod: "0069", Msg: "Software field Cannot Be Changed"},
				"E0070": {Cod: "0070", Msg: "Software version field Cannot Be Changed"},
				"E0071": {Cod: "0071", Msg: "Batch lock field Cannot Be Changed"},
				"E0072": {Cod: "0072", Msg: "Lock field Cannot Be Changed"},
				"E0073": {Cod: "0073", Msg: "Lock reason field Cannot Be Changed"},
				"E0074": {Cod: "0074", Msg: "Key field Cannot Be Changed"},
				"E0075": {Cod: "0075", Msg: "Serial field Cannot Be Changed"},
				"E0076": {Cod: "0076", Msg: "Type field Cannot Be Changed"},
				"E0077": {Cod: "0077", Msg: "Invoice id field Cannot Be Changed"},
				"E0078": {Cod: "0078", Msg: "lock invalid"},
				"E0079": {Cod: "0079", Msg: "Unlock invalid"},
				//Terminal model
				"E0080": {Cod: "0080", Msg: "Name cannot be null"},
				"E0081": {Cod: "0081", Msg: "Serial type cannot be null"},
				"E0082": {Cod: "0082", Msg: "Status cannot be null"},
				//Terminal note
				"E0083": {Cod: "0083", Msg: "Id terminal cannot be null"},
				//Terminal operator
				"E0084": {Cod: "0084", Msg: "Terminal identifier cannot be null"},
				"E0085": {Cod: "0085", Msg: "Sale network identifier cannot be null"},
				"E0086": {Cod: "0086", Msg: "Some operator is already linked to the terminal"},
				//Transition
				"E0087": {Cod: "0087", Msg: "Name cannot be null"},
				"E0088": {Cod: "0088", Msg: "Allocation initial cannot be null"},
				"E0089": {Cod: "0089", Msg: "Allocation final cannot be null"},
				"E0090": {Cod: "0090", Msg: "Customer receipt cannot be null"},
				"E0091": {Cod: "0091", Msg: "Maintenance protocol cannot be null"},
				"E0092": {Cod: "0092", Msg: "Delivery term cannot be null"},
				"E0093": {Cod: "0093", Msg: "Closure receipt cannot be null"},
				"E0094": {Cod: "0094", Msg: "Attach document cannot be null"},
				"E0095": {Cod: "0095", Msg: "Send email customer cannot be null"},
				"E0096": {Cod: "0096", Msg: "Serial list cannot be null"},
				"E0097": {Cod: "0097", Msg: "Contract cannot be null"},
				"E0098": {Cod: "0098", Msg: "Can perform transition cannot be null"},
				"E0099": {Cod: "0099", Msg: "Status cannot be null"},
				"E0100": {Cod: "0100", Msg: "The transition is not valid"},
				//errors in service
				"E0101": {Cod: "0101", Msg: "File type is invalid, only pdf is allowed"},
				"E0102": {Cod: "0102", Msg: "Terminal blocked"},
				"E0103": {Cod: "0103", Msg: "Seller is invalid"},
				"E0104": {Cod: "0104", Msg: "Code, name or ShortName is already in use"},
				//"E0105": {Cod: "0105", Msg: "Accounting limit cannot be null in Company"},
				"E0106": {Cod: "0106", Msg: "The application cannot be null in Company"},
				"E0107": {Cod: "0107", Msg: "Company not informed"},
				"E0108": {Cod: "0108", Msg: "Regional not informed"},
				"E0110": {Cod: "0110", Msg: "Chip already registered"},
				"E0111": {Cod: "0111", Msg: "Model's name already registered"},
				"E0112": {Cod: "0112", Msg: "Allocation's name already registered"},
				"E0113": {Cod: "0113", Msg: "Sale Limit: Incorrect Value"},
				"E0114": {Cod: "0114", Msg: "Award payment limit: Incorrect Value"},
				"E0115": {Cod: "0115", Msg: "File type is invalid, only csv is allowed"},
				"E0116": {Cod: "0116", Msg: "User does not have permission for this operation"},
				"E0117": {Cod: "0117", Msg: "Special max award payment limit: Incorrect Value"},
				"E0118": {Cod: "0118", Msg: "Negative paid value"},
				"E0119": {Cod: "0119", Msg: "Paid value larger than balance"},
				"E0120": {Cod: "0120", Msg: "Zero paid value and non zero balance"},
				"E0121": {Cod: "0121", Msg: "No such file or directory"},
				"E0122": {Cod: "0122", Msg: "The row does not have the right amount of columns"},
				"E0123": {Cod: "0123", Msg: "Not found delimiter ';' "},
				"E0124": {Cod: "0124", Msg: "Error opening file"},
				"E0125": {Cod: "0125", Msg: "File not found"},
			},
		},
		"C04": {
			Type: "Report",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Unable to load data"},
				"E0002": {Cod: "0002", Msg: "Scan error"},
				"E0003": {Cod: "0003", Msg: "Error checking permission"},
				"E0004": {Cod: "0004", Msg: "Negative paid value"},
				"E0005": {Cod: "0005", Msg: "Paid value larger than balance"},
				"E0006": {Cod: "0006", Msg: "An error occurred while sending payment details"},
				"E0007": {Cod: "0007", Msg: "Rendering of accounts performed successfully, there was an error in returning to the user"},
				"E0008": {Cod: "0008", Msg: "Slipe Size is required. Check if the input was sent correctly"},
				"E0009": {Cod: "0009", Msg: "Bobbin size is required. Check if the input was sent correctly"},
				"E0010": {Cod: "0010", Msg: "The quantity of bobbin per box is required. Check if the input was sent correctly"},
				"E0011": {Cod: "0011", Msg: "No reports found, check your permissions"},
				"E0012": {Cod: "0012", Msg: "Check required fields"},
				"E0013": {Cod: "0013", Msg: "Check if sent parameters are valid"},
				"E0014": {Cod: "0014", Msg: "Check if fields for sorting are valid"},
				"E0015": {Cod: "0015", Msg: "(Code) OR (User) OR (Company) OR (Regional) OR (Initial Allocation) OR (Final Allocation)"},
				"E0016": {Cod: "0016", Msg: "Edition interval is invalid"},
				"E0017": {Cod: "0017", Msg: "You need to inform who you are filtering by"},
				"E0018": {Cod: "0018", Msg: "Value filtered by is invalid"},
				"E0019": {Cod: "0019", Msg: "Invalid parameters: check current edition and reference edition parameters"},
				"E0020": {Cod: "0020", Msg: "Value limit not informed."},
			},
		},
		"C05": {
			Type: "Dispatch",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: ""},
			},
		},
		"C06": {
			Type: "Database",
			List: map[string]*ErrorItem{
				"E0001": {Cod: "0001", Msg: "Instruction error"},
			},
		},
	},
}
