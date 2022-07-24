Create schema mnemonicsdb;

use mnemonicsdb; 

create table mnems (
	 MnemonicID varchar(255) not null primary key
    ,MemoryTip varchar(255) not null
    ,ItemToRemember varchar(255) not null
    ,ImageID    varchar(255)   
    ,CreatedByID   varchar(255) not null
    ,CreatedDate   DateTime not null
);

create table user (
	UserID varchar(255) not null primary key,
    NickName varchar(255) not null,
    DateJoined DateTime not null,
    MnemonicCount int
);

create table collection(
     CollectionID varchar(255) not null primary key,
     CreatedByID varchar(255) not null,

);

create table mnemCollectionMap(
    CollectionID varchar(255) not null,
    MnemonicID varchar(255) not null
);

create table upvoteLog(
    EntityType enum('mnemonic', 'collection') not null,
    actorIsRegistered bool not null,
    actorID   string null,
    ipAddress  string null,
    Date DateTime not null

);

create table viewLog(
    EntityType enum('mnemonic', 'collection', 'user') not null,
    actorIsRegistered bool not null,
    actorID   string null,
    ipAddress  string null,
    Date DateTime not null
)

insert into mnems values ( "12", "blah", "item", null, "user1", "today!")

