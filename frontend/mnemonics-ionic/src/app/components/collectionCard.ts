import { CommonModule } from "@angular/common";
import { Component, Input } from "@angular/core";
import { IonicModule} from '@ionic/angular';

@Component({
    selector:'app-collectionCard',
    standalone:true,
    templateUrl:"./collectionCard.html",
    imports:[
        IonicModule,

    ]
})
export class CollectionCard {
    @Input() collectionInfo: CollectionInfo | null = null
    constructor(){}

}

export interface CollectionInfo{
    id: string
    link: string
    name: string
    mnemonicCount: number
    //instead of stars, maybe I could use a voteBalance, then let ppl upvote/downvote... or maybe keep both? 
    //starCount could actually be "Favorited"
    starCount: number
    viewCount: number
}