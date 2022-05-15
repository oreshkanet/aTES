-- +migrate Up
CREATE DATABASE [task_tracker]

-- +migrate Up
IF  NOT EXISTS (SELECT * FROM sys.objects
    WHERE object_id = OBJECT_ID(N'[dbo].[users]') AND type in (N'U'))

BEGIN
    CREATE TABLE [dbo].[users] (
        [id] INT IDENTITY(1,1) NOT NULL CONSTRAINT pk_users_id PRIMARY KEY,
        [public_id] VARCHAR(40) NOT NULL,
        [name] VARCHAR(250) NOT NULL,
        [role] VARCHAR(50) NOT NULL
    );

    CREATE UNIQUE INDEX ix_users_publicId ON [dbo].[users] (public_id ASC);
END

IF  NOT EXISTS (SELECT * FROM sys.objects
                WHERE object_id = OBJECT_ID(N'[dbo].[users]') AND type in (N'U'))

BEGIN
    CREATE TABLE [dbo].[tasks] (
       [id] INT IDENTITY(1,1) NOT NULL CONSTRAINT pk_tasks_id PRIMARY KEY,
       [public_id] VARCHAR(40) NOT NULL,
       [title] VARCHAR(250) NOT NULL,
       [description] VARCHAR(1024) NOT NULL,
       [assigned_user] VARCHAR(40) NOT NULL,
    );

    CREATE UNIQUE INDEX ix_users_publicId ON [dbo].[users] (public_id ASC);
END

-- +migrate Down
DROP DATABASE [task_tracker]

-- +migrate Down
DROP TABLE [users]