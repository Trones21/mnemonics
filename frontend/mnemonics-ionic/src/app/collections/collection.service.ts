import { Injectable } from '@angular/core';
import { DataService } from '../data.service';

//ToDo Next: Begin creating entity services to inject data into pages/components 
@Injectable({
  providedIn: 'root'
})
export class CollectionService {
  baseUrl: string
  constructor(dataService: DataService,) { 
     this.baseUrl = dataService.baseUrl;
  }

  async getCollectionsInCategory(categoryId: string){
    let res = await fetch(`${this.baseUrl}/collections?categoryId=${categoryId}`)
    let obj = await res.json();
    console.log(obj.child)
    return obj.child;
  }


  // async getCollectionByID(collectionId: string){
  //   let res = await fetch(`${this.baseUrl}/collections/`)
  //   //this should also return the full details of all the mnemonics, ...but a collection could have like 100 mnemonics 
  //   let obj = await res.json();
  //   console.log(obj.child)
  //   return obj.child;
  // }



}
