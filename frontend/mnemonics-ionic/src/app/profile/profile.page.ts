import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { CollectionCard } from '../components/collectionCard';
import { CollectionInfo } from '../components/collectionCard';
import { ProfileService } from './profile.service';
import { ActivatedRoute, Params, Router } from '@angular/router';

@Component({
  selector: 'app-profile',
  templateUrl: 'profile.page.html'
})

export class ProfilePage implements OnInit{

  profile: any 
  myUserId = '6425fe608d3ae9661a039848'//localStorage.getItem('loggedInUserID')
  
  profileService: ProfileService
  constructor(profileService: ProfileService, private route: ActivatedRoute) {
    this.profileService = profileService
    this.route = route
  }

  ngOnInit(){
    let idParam = this.route.snapshot.params['id']
    console.log(idParam)
    this.profileService.getProfile(idParam).then(p => this.profile = p)
  }

  //Endpoints 
  //logo = api/collections/<id>/logo
  //collection items =   api/collections/<id>/all
  //stats = api/profile/<id>/stats
  //profile = api/profile/<id>
  
 
}

//ToDo later if I need to, may remove the TypeError
interface Profile{
username: string,
userID: string,

}

