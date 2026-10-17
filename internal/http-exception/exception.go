package http_exception

type HTTPExceptionBody struct {
	Message    string `json:"message"`
	Error      string `json:"error"`
	StatusCode int    `json:"status_code"`
}

type IHTTPException interface {
	ResponseCode() int
	ResponseBody() HTTPExceptionBody
	Error() string
}

type HTTPStatusCode = string

const (
	HTTPContinueStatus                      HTTPStatusCode = "CONTINUE"                        // 100
	HTTPSwitchingProtocolsStatus            HTTPStatusCode = "SWITCHING_PROTOCOLS"             // 101
	HTTPProcessingStatus                    HTTPStatusCode = "PROCESSING"                      // 102
	HTTPEarlyHintsStatus                    HTTPStatusCode = "EARLY_HINTS"                     // 103
	HTTPOkStatus                            HTTPStatusCode = "OK"                              // 200
	HTTPCreatedStatus                       HTTPStatusCode = "CREATED"                         // 201
	HTTPAcceptedStatus                      HTTPStatusCode = "ACCEPTED"                        // 202
	HTTPNonAuthoritativeInfoStatus          HTTPStatusCode = "NON_AUTHORITATIVE_INFO"          // 203
	HTTPNoContentStatus                     HTTPStatusCode = "NO_CONTENT"                      // 204
	HTTPResetContentStatus                  HTTPStatusCode = "RESET_CONTENT"                   // 205
	HTTPPartialContentStatus                HTTPStatusCode = "PARTIAL_CONTENT"                 // 206
	HTTPMultiStatusStatus                   HTTPStatusCode = "MULTI_STATUS"                    // 207
	HTTPAlreadyReportedStatus               HTTPStatusCode = "ALREADY_REPORTED"                // 208
	HTTPIMUsedStatus                        HTTPStatusCode = "IM_USED"                         // 226
	HTTPMultipleChoicesStatus               HTTPStatusCode = "MULTIPLE_CHOICES"                // 300
	HTTPMovedPermanentlyStatus              HTTPStatusCode = "MOVED_PERMANENTLY"               // 301
	HTTPFoundStatus                         HTTPStatusCode = "FOUND"                           // 302
	HTTPSeeOtherStatus                      HTTPStatusCode = "SEE_OTHER"                       // 303
	HTTPNotModifiedStatus                   HTTPStatusCode = "NOT_MODIFIED"                    // 304
	HTTPTemporaryRedirectStatus             HTTPStatusCode = "TEMPORARY_REDIRECT"              // 307
	HTTPPermanentRedirectStatus             HTTPStatusCode = "PERMANENT_REDIRECT"              // 308
	HTTPBadRequestStatus                    HTTPStatusCode = "BAD_REQUEST"                     // 400
	HTTPUnauthorizedStatus                  HTTPStatusCode = "UNAUTHORIZED"                    // 401
	HTTPPaymentRequiredStatus               HTTPStatusCode = "PAYMENT_REQUIRED"                // 402
	HTTPForbiddenStatus                     HTTPStatusCode = "FORBIDDEN"                       // 403
	HTTPNotFoundStatus                      HTTPStatusCode = "NOT_FOUND"                       // 404
	HTTPMethodNotAllowedStatus              HTTPStatusCode = "METHOD_NOT_ALLOWED"              // 405
	HTTPNotAcceptableStatus                 HTTPStatusCode = "NOT_ACCEPTABLE"                  // 406
	HTTPProxyAuthRequiredStatus             HTTPStatusCode = "PROXY_AUTH_REQUIRED"             // 407
	HTTPRequestTimeoutStatus                HTTPStatusCode = "REQUEST_TIMEOUT"                 // 408
	HTTPConflictStatus                      HTTPStatusCode = "CONFLICT"                        // 409
	HTTPGoneStatus                          HTTPStatusCode = "GONE"                            // 410
	HTTPLengthRequiredStatus                HTTPStatusCode = "LENGTH_REQUIRED"                 // 411
	HTTPPreconditionFailedStatus            HTTPStatusCode = "PRECONDITION_FAILED"             // 412
	HTTPPayloadTooLargeStatus               HTTPStatusCode = "PAYLOAD_TOO_LARGE"               // 413
	HTTPURITooLongStatus                    HTTPStatusCode = "URI_TOO_LONG"                    // 414
	HTTPUnsupportedMediaTypeStatus          HTTPStatusCode = "UNSUPPORTED_MEDIA_TYPE"          // 415
	HTTPRangeNotSatisfiableStatus           HTTPStatusCode = "RANGE_NOT_SATISFIABLE"           // 416
	HTTPExpectationFailedStatus             HTTPStatusCode = "EXPECTATION_FAILED"              // 417
	HTTPImATeapotStatus                     HTTPStatusCode = "IM_A_TEAPOT"                     // 418
	HTTPMisdirectedRequestStatus            HTTPStatusCode = "MISDIRECTED_REQUEST"             // 421
	HTTPUnprocessableEntityStatus           HTTPStatusCode = "UNPROCESSABLE_ENTITY"            // 422
	HTTPLockedStatus                        HTTPStatusCode = "LOCKED"                          // 423
	HTTPFailedDependencyStatus              HTTPStatusCode = "FAILED_DEPENDENCY"               // 424
	HTTPTooEarlyStatus                      HTTPStatusCode = "TOO_EARLY"                       // 425
	HTTPUpgradeRequiredStatus               HTTPStatusCode = "UPGRADE_REQUIRED"                // 426
	HTTPPreconditionRequiredStatus          HTTPStatusCode = "PRECONDITION_REQUIRED"           // 428
	HTTPTooManyRequestsStatus               HTTPStatusCode = "TOO_MANY_REQUESTS"               // 429
	HTTPRequestHeaderFieldsTooLargeStatus   HTTPStatusCode = "REQUEST_HEADER_FIELDS_TOO_LARGE" // 431
	HTTPUnavailableForLegalReasonsStatus    HTTPStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"   // 451
	HTTPInternalServerErrorStatus           HTTPStatusCode = "INTERNAL_SERVER_ERROR"           // 500
	HTTPNotImplementedStatus                HTTPStatusCode = "NOT_IMPLEMENTED"                 // 501
	HTTPBadGatewayStatus                    HTTPStatusCode = "BAD_GATEWAY"                     // 502
	HTTPServiceUnavailableStatus            HTTPStatusCode = "SERVICE_UNAVAILABLE"             // 503
	HTTPGatewayTimeoutStatus                HTTPStatusCode = "GATEWAY_TIMEOUT"                 // 504
	HTTPHTTPVersionNotSupportedStatus       HTTPStatusCode = "HTTP_VERSION_NOT_SUPPORTED"      // 505
	HTTPVariantAlsoNegotiatesStatus         HTTPStatusCode = "VARIANT_ALSO_NEGOTIATES"         // 506
	HTTPInsufficientStorageStatus           HTTPStatusCode = "INSUFFICIENT_STORAGE"            // 507
	HTTPLoopDetectedStatus                  HTTPStatusCode = "LOOP_DETECTED"                   // 508
	HTTPNotExtendedStatus                   HTTPStatusCode = "NOT_EXTENDED"                    // 510
	HTTPNetworkAuthenticationRequiredStatus HTTPStatusCode = "NETWORK_AUTHENTICATION_REQUIRED" // 511
)

type HTTPException struct {
	statusCode HTTPStatusCode
	error      string
}

var statusCodes = map[HTTPStatusCode]int{
	HTTPContinueStatus:                      100,
	HTTPSwitchingProtocolsStatus:            101,
	HTTPProcessingStatus:                    102,
	HTTPEarlyHintsStatus:                    103,
	HTTPOkStatus:                            200,
	HTTPCreatedStatus:                       201,
	HTTPAcceptedStatus:                      202,
	HTTPNonAuthoritativeInfoStatus:          203,
	HTTPNoContentStatus:                     204,
	HTTPResetContentStatus:                  205,
	HTTPPartialContentStatus:                206,
	HTTPMultiStatusStatus:                   207,
	HTTPAlreadyReportedStatus:               208,
	HTTPIMUsedStatus:                        226,
	HTTPMultipleChoicesStatus:               300,
	HTTPMovedPermanentlyStatus:              301,
	HTTPFoundStatus:                         302,
	HTTPSeeOtherStatus:                      303,
	HTTPNotModifiedStatus:                   304,
	HTTPTemporaryRedirectStatus:             307,
	HTTPPermanentRedirectStatus:             308,
	HTTPBadRequestStatus:                    400,
	HTTPUnauthorizedStatus:                  401,
	HTTPPaymentRequiredStatus:               402,
	HTTPForbiddenStatus:                     403,
	HTTPNotFoundStatus:                      404,
	HTTPMethodNotAllowedStatus:              405,
	HTTPNotAcceptableStatus:                 406,
	HTTPProxyAuthRequiredStatus:             407,
	HTTPRequestTimeoutStatus:                408,
	HTTPConflictStatus:                      409,
	HTTPGoneStatus:                          410,
	HTTPLengthRequiredStatus:                411,
	HTTPPreconditionFailedStatus:            412,
	HTTPPayloadTooLargeStatus:               413,
	HTTPURITooLongStatus:                    414,
	HTTPUnsupportedMediaTypeStatus:          415,
	HTTPRangeNotSatisfiableStatus:           416,
	HTTPExpectationFailedStatus:             417,
	HTTPImATeapotStatus:                     418,
	HTTPMisdirectedRequestStatus:            421,
	HTTPUnprocessableEntityStatus:           422,
	HTTPLockedStatus:                        423,
	HTTPFailedDependencyStatus:              424,
	HTTPTooEarlyStatus:                      425,
	HTTPUpgradeRequiredStatus:               426,
	HTTPPreconditionRequiredStatus:          428,
	HTTPTooManyRequestsStatus:               429,
	HTTPRequestHeaderFieldsTooLargeStatus:   431,
	HTTPUnavailableForLegalReasonsStatus:    451,
	HTTPInternalServerErrorStatus:           500,
	HTTPNotImplementedStatus:                501,
	HTTPBadGatewayStatus:                    502,
	HTTPServiceUnavailableStatus:            503,
	HTTPGatewayTimeoutStatus:                504,
	HTTPHTTPVersionNotSupportedStatus:       505,
	HTTPVariantAlsoNegotiatesStatus:         506,
	HTTPInsufficientStorageStatus:           507,
	HTTPLoopDetectedStatus:                  508,
	HTTPNotExtendedStatus:                   510,
	HTTPNetworkAuthenticationRequiredStatus: 511,
}

func (H HTTPException) ResponseCode() int {
	return statusCodes[H.statusCode]
}

func (H HTTPException) ResponseBody() HTTPExceptionBody {
	return HTTPExceptionBody{
		Message:    H.error,
		Error:      H.statusCode,
		StatusCode: statusCodes[H.statusCode],
	}
}

func (H HTTPException) Error() string {
	return H.error
}

func New(statusCode HTTPStatusCode, err string) HTTPException {
	return HTTPException{
		statusCode: statusCode,
		error:      err,
	}
}
