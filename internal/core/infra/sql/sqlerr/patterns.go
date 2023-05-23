package sqlerr

var (
	notFoundPattern               = "record not found"
	uniqueConstraintPattern       = "violates unique constraint"
	foreignKeyConstraintPattern   = "violates foreign key constraint"
	notNullConstraintPattern      = "violates not-null constraint"
	checkConstraintPattern        = "violates check constraint"
	contextDeadlinePattern        = "context deadline exceeded"
	tooManyClientsPattern         = "sorry, too many clients already"
	noRowsAffectedPattern         = "no rows affected"
	failedToConnectPattern        = "failed to connect to"
	columnDoesNotExistPattern     = "SQLSTATE 42703"
	inputSyntaxPattern            = "SQLSTATE 22P02"
	missingWhereConditionsPattern = "WHERE conditions required"
	tooLongValuePattern           = "SQLSTATE 22001"
)
