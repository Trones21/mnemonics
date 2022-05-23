Create schema mnemonicsdb;

use mnemonicsdb; 

create table mnems (
	MnemonicID varchar(255) not null primary key
    ,MemoryTip varchar(255) not null
    ,ItemToRemember varchar(255) not null
,ImageID    varchar(255)   
,CreatedByID   varchar(255) not null
,CreatedDate   varchar(255) not null
);

create table user (
	UserID varchar(255) not null primary key,
    NickName varchar(255) not null,
    DateJoined varchar(255) not null,
    MnemonicCount int
);

insert into mnems values ( "12", "blah", "item", null, "user1", "today!")

