import { Injectable } from '@angular/core';
import { DataService } from '../data.service';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {
  baseUrl: string
  constructor(dataService: DataService,) { 
     this.baseUrl = dataService.baseUrl;
  }

  async getProfile(profileId: string){
    let res = await fetch(`${this.baseUrl}/profile/${profileId}`)
    let obj = await res.json();
    console.log(obj.profile)
    return obj.profile;
  }

  // getProfileRxjs(){

  // }

  createProfile(){

  }

  updateProfile(profile: any){
    fetch( `${this.baseUrl}/profile/${profile.Id}`,
      {method:'POST',
       body:profile
    })
  }
}
