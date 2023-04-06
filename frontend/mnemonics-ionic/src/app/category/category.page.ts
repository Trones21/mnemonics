import { Component, OnInit } from "@angular/core"
import { CategoriesService } from "./categories.service"
import { CollectionService } from "../collections/collection.service"
import {CollectionCard} from "../components/collectionCard"
import { CollectionInfo } from "../components/collectionCard"
import { CommonModule, NgFor } from "@angular/common"

@Component({
    selector: 'app-category',
    templateUrl: 'category.page.html',
    standalone: true,
    imports:[
        CollectionCard,
        CommonModule,
        NgFor
    ]
  })

//This is currently unused -- switch to it when I switch to routing. 
//The collections page currently "navigates" to the category page using a switch statements
export class CategoryPage implements OnInit{
    // category = {
    //     name:"Math",
    //     id: 3,
    //     collections:[
    //         {id:'123', name: "Algebra"},
    //         {id:'abc', name: 'Geometry'} 
    //     ]
    // }
    category: CategoryDetails = {id:'', name:'', collectionCount: 0, collectionCards: new Array<CollectionInfo>}
    constructor(private categoriesService: CategoriesService, private collectionService: CollectionService){
    }

    ngOnInit(){
        // this.categoriesService.getCategoryByID('1').then( c => {
        //         c.collectionCards = new Array<CollectionInfo> 
        //         for(let col of c.collections){
        //             let cInfo = new CollectionInfo(
        //                 col.id,
        //                 col.link,
        //                 col.name,
        //                 col.mnemonicsCount,
        //                 col.starCount,
        //                 col.viewCount
        //                 )
        //             c.collectionCards.push(cInfo)
        //         }
        //         this.category = (c as CategoryDetails)
        // })
    }

}

export interface CategoryDetails{
    id: string,
    name: string,
    collectionCount: number,
    collectionCards: Array<CollectionInfo>
  }