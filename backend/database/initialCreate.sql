drop schema mnemonicsdb;
Create schema mnemonicsdb;
use mnemonicsdb;
create table mnems (
    MnemonicID varchar(255) not null primary key,
    MemoryTip varchar(255) not null,
    ItemToRemember varchar(255) not null,
    ImageID varchar(255),
    CreatedByID varchar(255) not null,
    CreatedDate DateTime not null
);
create table user (
    UserID varchar(255) not null primary key,
    NickName varchar(255) not null unique,
    DateJoined DateTime not null,
    EmailAddress varchar(255) not null unique,
    MnemonicCount int,
    Constraint CHK_Email CHECK (EmailAddress like '%@%.%')
);
create table collection(
    CollectionID varchar(255) not null primary key,
    CreatedByID varchar(255) not null
);
create table mnemCollectionMap(
    CollectionID varchar(255) not null,
    MnemonicID varchar(255) not null
);
create table upvoteLog(
    EntityType enum('mnemonic', 'collection') not null,
    actorIsRegistered bool not null,
    actorID varchar(255) null,
    ipAddress varchar(255) null,
    Date DateTime not null
);
create table viewLog(
    EntityType enum('mnemonic', 'collection', 'user') not null,
    actorIsRegistered bool not null,
    actorID varchar(255) null,
    ipAddress varchar(255) null,
    Date DateTime not null
)
insert into mnems
values ("12", "blah", "item", null, "user1", "20220301")
insert into user
values("user1", "Tommy", "20220101", 1)