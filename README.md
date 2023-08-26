// developed by Dariel Vega Bernal

This is not the final document, it is a work in progress.

# Introduction

# API explanation

- Get test

    https://www.hackerrank.com/x/api/v3/tests?limit=<number, must be more than zero>&offset=<number>

    This api returns a list of tests, the limit is the number of tests that you want to get and the offset is the number of tests that you want to skip.

    example: https://www.hackerrank.com/x/api/v3/tests?limit=2&offset=0

    In that example we are getting the first two tests.

    example: https://www.hackerrank.com/x/api/v3/tests?limit=2&offset=2

    In that example we are getting the third and fourth tests.

- Get candidates

    https://www.hackerrank.com/x/api/v3/tests/{test_id obligatory}/candidates?limit={number, must be more than zero}&offset={number}

    This api returns a list of candidates, the limit is the number of candidates that you want to get and the offset is the number of candidates that you want to skip.
    
    it have the same example as the Get Test

- Get pdf

    https://www.hackerrank.com/x/tests/{test_id obligatory}/candidates/{candidate id}/pdf

    candidate id example: "id": "1234556",

    " status
    This indicates the current status of the candidate. Can be between -1 to 7. -1 indicates that candidate is currently invited, while 7 indicates that candidates has completed the test and a report has been generated for the candidate. All other intermediate states are internal and should not be used. "

    when the user status is 7, this will have a pdf

