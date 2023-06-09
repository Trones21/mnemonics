import { Injectable } from '@angular/core';
import { DataService } from '../data.service';
import { CollectionInfo } from '../components/collectionCard';
import { Collection } from './singleCollection.page';

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
    console.log(obj)
    return obj;
  }


  async getCollectionByID(collectionId: string, includeMnemonicDetails: boolean): Promise<Collection>{
      let res = await fetch(`${this.baseUrl}/collections/details/${collectionId}?includeMnemonicDetails=${includeMnemonicDetails}`)
      let obj = await res.json();
      let collection = (obj.collection as Collection);
      return collection
  }



}
