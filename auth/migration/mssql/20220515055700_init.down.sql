IF EXISTS(
        SELECT *
        FROM sys.databases
        WHERE name = 'auth'
) BEGIN EXEC ('DROP DATABASE auth')