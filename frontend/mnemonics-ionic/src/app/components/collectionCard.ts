import { CommonModule } from "@angular/common";
import { Component, Input } from "@angular/core";
import { IonicModule} from '@ionic/angular';




export class CollectionInfo{
    link: string
    name: string
    mnemonics: number
    stars: number
    views: number

    constructor(
        link: string, 
        name: string, 
        mnemonics: number,
        stars: number,
        views: number){
        
        this.link = link
        this.name = name
        this.mnemonics = mnemonics
        this.stars = stars
        this.views = views
    }

}


@Component({
    selector:'app-collectionCard',
    standalone:true,
    templateUrl:"./collectionCard.html",
    imports:[
        IonicModule,

    ]
})
export class CollectionCard {
    @Input() collectionInfo: CollectionInfo = new CollectionInfo('/', 'Collection Name', 0, 0, 0)
    constructor(){}

}