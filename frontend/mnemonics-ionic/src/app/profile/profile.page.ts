import { Component } from '@angular/core';
import { CollectionCard } from '../components/collectionCard';
import { CollectionInfo } from '../components/collectionCard';

@Component({
  selector: 'app-profile',
  templateUrl: 'profile.page.html',
})

// class Profile{
//   username:string
//   userID: string
//   mnemonicCount: number

// }

export class ProfilePage {

  constructor() {}

  //Endpoints 
  //logo = api/collections/<id>/logo
  //collection items =   api/collections/<id>/all
  //stats = api/profile/<id>/stats
  //profile = api/profile/<id>
  myUserId = "6425fe608d3ae9661a03984"
  profile = {
    username: "Trones",
    userID: "6425fe608d3ae9661a039848",
    mnemonicCount: 17, 
    badges: [{
      name: "Top 10% of Contributors",
      logoPath: '/badges/top10'
    }
    ],
    myCollections:
    [
      { 
        link: "/collections/6425fe608d3ae9661a039848-1",
        name: "myCreations",
        mnemonics: 17,
        stars: 51,
        views: 216,
        createdAt: "2011-09-04T06:31:12.212Z",
        updatedAt: "2011-09-05T06:31:12.212Z"
      }
    ],
    favoritedCollections: [
      { 
      link: "/collections/6425fe608d3ae9661a039848",
      name: "Medical",
      id: "6425fe604e175c2a4a8265e9",
      mnemonics: 10,
      stars: 51,
      views: 216,
      createdAt: "2011-09-04T06:31:12.212Z",
      updatedAt: "2011-09-05T06:31:12.212Z"
     },
     { 
     link: "/collections/6425fe608d3ae9661a039848",  
     name: "Math",
     id: "6425fe604e175c2a4a8265e9",
     mnemonics: 2,
     stars: 12,
     views: 199,
     createdAt: "2011-10-04T06:31:12.212Z",
     updatedAt: "2011-11-05T06:31:12.212Z"
    },
    ]
  } 

}

