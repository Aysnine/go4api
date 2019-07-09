
var setUpStartUnixNano = 0;
var setUpStart = "";
var setUpEndUnixNano = 0;
var setUpEnd = "";

var normalStartUnixNano = 1562658959340926000;
var normalStart = "2019-07-09 15:55:59.340926";
var normalEndUnixNano = 1562658959590055000;
var normalEnd = "2019-07-09 15:55:59.590055";

var tearDownStartUnixNano = 0;
var tearDownStart = "";
var tearDownEndUnixNano = 0;
var tearDownEnd = "";

var gStartUnixNano = 1562658959340926000;
var gStart = "2019-07-09 15:55:59.340926";
var gEndUnixNano = 1562658959590055000;
var gEnd = "2019-07-09 15:55:59.590055";

var tcResults = [{"TcName":"resapi-dummy-ec-001","IfGlobalSetUpTearDown":"RegularCases","Priority":"530","ParentTestCase":"root","SetUpResult":"NoSetUp","SetUpTestMessages":[],"Path":"","Method":"","JsonFilePath":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/1dummy-ec copy.json","CsvFile":"","CsvRow":"","MutationArea":"","MutationCategory":"","MutationRule":"","MutationInfo":{"CurrValue":null,"FieldPath":null,"FieldSubType":"","FieldType":"","MutatedValue":null},"HttpResult":"NoHttp","ActualStatusCode":999,"StartTime":"2019-07-09 15:55:59.340926","EndTime":"2019-07-09 15:55:59.340931","HttpTestMessages":null,"StartTimeUnixNano":1562658959340926000,"EndTimeUnixNano":1562658959340931000,"DurationUnixNano":5000,"DurationUnixMillis":0,"TearDownResult":"NoTearDown","TearDownTestMessages":[],"TestResult":"Success","HttpUrl":"","CaseOrigin":{"resapi-dummy-ec-001":{"ifGlobalSetUpTestCase":false,"ifGlobalTearDownTestCase":false,"inputs":null,"outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"outputs":null,"parentTestCase":"root","priority":"530","request":null,"response":null,"session":null,"setUp":null,"tearDown":null}},"GlobalVariables":{},"Session":{},"LocalVariables":{},"ActualHeader":{},"ActualBody":null},{"TcName":"ec-sql-0002","IfGlobalSetUpTearDown":"RegularCases","Priority":"531","ParentTestCase":"resapi-dummy-ec-001","SetUpResult":"Success","SetUpTestMessages":[[{"AssertionResults":"Success","ReponsePart":"sql","FieldName":"$(sql).#","AssertionKey":"GreaterOrEquals","ActualValue":0,"ExpValue":0}]],"Path":"","Method":"","JsonFilePath":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/2-1sql-update-coupon-ec.json","CsvFile":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/2-1sql-update-coupon-ec_dt1.csv","CsvRow":"3","MutationArea":"","MutationCategory":"","MutationRule":"","MutationInfo":{"CurrValue":null,"FieldPath":null,"FieldSubType":"","FieldType":"","MutatedValue":null},"HttpResult":"NoHttp","ActualStatusCode":999,"StartTime":"2019-07-09 15:55:59.49778","EndTime":"2019-07-09 15:55:59.497785","HttpTestMessages":null,"StartTimeUnixNano":1562658959497780000,"EndTimeUnixNano":1562658959497785000,"DurationUnixNano":5000,"DurationUnixMillis":0,"TearDownResult":"NoTearDown","TearDownTestMessages":[],"TestResult":"Success","HttpUrl":"","CaseOrigin":{"ec-sql-0002":{"ifGlobalSetUpTestCase":false,"ifGlobalTearDownTestCase":false,"inputs":null,"outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"outputs":null,"parentTestCase":"resapi-dummy-ec-001","priority":"531","request":null,"response":null,"session":null,"setUp":[{"cmd":"update reservation.coupon set selling_city_id = 'EC' where COUPON_ID = '931130000112';","cmdResponse":{"$(sql).#":{"GreaterOrEquals":0}},"cmdSource":"master","cmdType":"sql","outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"session":null}],"tearDown":null}},"GlobalVariables":{},"Session":{},"LocalVariables":{},"ActualHeader":{},"ActualBody":null},{"TcName":"ec-sql-0003","IfGlobalSetUpTearDown":"RegularCases","Priority":"531","ParentTestCase":"resapi-dummy-ec-001","SetUpResult":"Success","SetUpTestMessages":[[{"AssertionResults":"Success","ReponsePart":"sql","FieldName":"$(sql).#","AssertionKey":"GreaterOrEquals","ActualValue":0,"ExpValue":0}]],"Path":"","Method":"","JsonFilePath":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/2-1sql-update-coupon-ec.json","CsvFile":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/2-1sql-update-coupon-ec_dt1.csv","CsvRow":"4","MutationArea":"","MutationCategory":"","MutationRule":"","MutationInfo":{"CurrValue":null,"FieldPath":null,"FieldSubType":"","FieldType":"","MutatedValue":null},"HttpResult":"NoHttp","ActualStatusCode":999,"StartTime":"2019-07-09 15:55:59.587893","EndTime":"2019-07-09 15:55:59.587897","HttpTestMessages":null,"StartTimeUnixNano":1562658959587893000,"EndTimeUnixNano":1562658959587897000,"DurationUnixNano":4000,"DurationUnixMillis":0,"TearDownResult":"NoTearDown","TearDownTestMessages":[],"TestResult":"Success","HttpUrl":"","CaseOrigin":{"ec-sql-0003":{"ifGlobalSetUpTestCase":false,"ifGlobalTearDownTestCase":false,"inputs":null,"outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"outputs":null,"parentTestCase":"resapi-dummy-ec-001","priority":"531","request":null,"response":null,"session":null,"setUp":[{"cmd":"update reservation.coupon set selling_city_id = 'EC' where COUPON_ID = '931130000113';","cmdResponse":{"$(sql).#":{"GreaterOrEquals":0}},"cmdSource":"master","cmdType":"sql","outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"session":null}],"tearDown":null}},"GlobalVariables":{},"Session":{},"LocalVariables":{},"ActualHeader":{},"ActualBody":null},{"TcName":"ec-sql-0004","IfGlobalSetUpTearDown":"RegularCases","Priority":"531","ParentTestCase":"resapi-dummy-ec-001","SetUpResult":"Success","SetUpTestMessages":[[{"AssertionResults":"Success","ReponsePart":"sql","FieldName":"$(sql).#","AssertionKey":"GreaterOrEquals","ActualValue":0,"ExpValue":0}]],"Path":"","Method":"","JsonFilePath":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/2-1sql-update-coupon-ec.json","CsvFile":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/2-1sql-update-coupon-ec_dt1.csv","CsvRow":"5","MutationArea":"","MutationCategory":"","MutationRule":"","MutationInfo":{"CurrValue":null,"FieldPath":null,"FieldSubType":"","FieldType":"","MutatedValue":null},"HttpResult":"NoHttp","ActualStatusCode":999,"StartTime":"2019-07-09 15:55:59.589991","EndTime":"2019-07-09 15:55:59.589994","HttpTestMessages":null,"StartTimeUnixNano":1562658959589991000,"EndTimeUnixNano":1562658959589994000,"DurationUnixNano":3000,"DurationUnixMillis":0,"TearDownResult":"NoTearDown","TearDownTestMessages":[],"TestResult":"Success","HttpUrl":"","CaseOrigin":{"ec-sql-0004":{"ifGlobalSetUpTestCase":false,"ifGlobalTearDownTestCase":false,"inputs":null,"outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"outputs":null,"parentTestCase":"resapi-dummy-ec-001","priority":"531","request":null,"response":null,"session":null,"setUp":[{"cmd":"update reservation.coupon set selling_city_id = 'EC' where COUPON_ID = '931130000114';","cmdResponse":{"$(sql).#":{"GreaterOrEquals":0}},"cmdSource":"master","cmdType":"sql","outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"session":null}],"tearDown":null}},"GlobalVariables":{},"Session":{},"LocalVariables":{},"ActualHeader":{},"ActualBody":null},{"TcName":"ec-sql-0001","IfGlobalSetUpTearDown":"RegularCases","Priority":"531","ParentTestCase":"resapi-dummy-ec-001","SetUpResult":"Success","SetUpTestMessages":[[{"AssertionResults":"Success","ReponsePart":"sql","FieldName":"$(sql).#","AssertionKey":"GreaterOrEquals","ActualValue":0,"ExpValue":0}]],"Path":"","Method":"","JsonFilePath":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/2-1sql-update-coupon-ec.json","CsvFile":"/Users/pingzhu/StarbucksResAPI/scenarios/b1/2-1sql-update-coupon-ec_dt1.csv","CsvRow":"2","MutationArea":"","MutationCategory":"","MutationRule":"","MutationInfo":{"CurrValue":null,"FieldPath":null,"FieldSubType":"","FieldType":"","MutatedValue":null},"HttpResult":"NoHttp","ActualStatusCode":999,"StartTime":"2019-07-09 15:55:59.590053","EndTime":"2019-07-09 15:55:59.590055","HttpTestMessages":null,"StartTimeUnixNano":1562658959590053000,"EndTimeUnixNano":1562658959590055000,"DurationUnixNano":2000,"DurationUnixMillis":0,"TearDownResult":"NoTearDown","TearDownTestMessages":[],"TestResult":"Success","HttpUrl":"","CaseOrigin":{"ec-sql-0001":{"ifGlobalSetUpTestCase":false,"ifGlobalTearDownTestCase":false,"inputs":null,"outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"outputs":null,"parentTestCase":"resapi-dummy-ec-001","priority":"531","request":null,"response":null,"session":null,"setUp":[{"cmd":"update reservation.coupon set selling_city_id = 'EC' where COUPON_ID = '931130000111';","cmdResponse":{"$(sql).#":{"GreaterOrEquals":0}},"cmdSource":"master","cmdType":"sql","outFiles":null,"outGlobalVariables":null,"outLocalVariables":null,"session":null}],"tearDown":null}},"GlobalVariables":{},"Session":{},"LocalVariables":{},"ActualHeader":{},"ActualBody":null}]
