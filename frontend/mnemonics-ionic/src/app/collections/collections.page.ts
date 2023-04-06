import { Component } from '@angular/core';

@Component({
  selector: 'app-collections',
  templateUrl: 'collections.page.html',
})
export class CollectionsPage {
  public categories = [
    {id:1,name:'Medical', collections:12},
    {id:2,name:'Literary', collections:2},
    {id:3,name:'Math', collections:5},
    {id:4,name:'Biology', collections:1},
    {id:5,name:'Body Language', collections:0}
]
  loggedInUserID = localStorage.getItem('loggedInUserId')
  view: string = 'Search'
  constructor() {}

  showSearch(){
    this.view = 'Search'

  }

  showBrowse(){
    this.view = 'Browse'
  }

  Search(event: any){
    console.log('search', event.target)

  }

}