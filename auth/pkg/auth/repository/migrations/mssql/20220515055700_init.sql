-- +migrate Up
CREATE DATABASE [auth]

-- +migrate Up
IF  NOT EXISTS (SELECT * FROM sys.objects
    WHERE object_id = OBJECT_ID(N'[dbo].[users]') AND type in (N'U'))

BEGIN
    CREATE TABLE [dbo].[users] (
        [id] INT IDENTITY(1,1) NOT NULL CONSTRAINT pk_users_id PRIMARY KEY,
        [public_id] VARCHAR(40) NOT NULL,
        [name] VARCHAR(250) NOT NULL,
        [password] VARCHAR(50) NOT NULL,
        [role] VARCHAR(50) NOT NULL
    );

    CREATE UNIQUE INDEX ix_users_publicId ON [dbo].[users] (public_id ASC);
END

-- +migrate Down
DROP DATABASE [auth]

-- +migrate Down
DROP TABLE [users]