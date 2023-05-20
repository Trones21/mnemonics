import { CommonModule } from "@angular/common";
import { Component, Inject, Injectable, Input } from "@angular/core";
import { IonicModule } from '@ionic/angular';
import { IonDatetime } from "@ionic/angular";
import { MnemonicService } from "../mnemonic/mnemonic.service";



//Collection Ranks must be computed client side (due to planned db schema -- a single mnemonic can be in many collections,
//& I dont plan to keep track of collection level ranks in DB)

// @Injectable()
// export class Mnemonic{
//     link: String
//     name: String
//     stars: Number
//     views: Number
//     globalViewRank: Number
//     globalStarRank: Number
//     collectionViewRank?: Number
//     collectionStarRank?: Number
//     constructor(
//         link: String, 
//         name: String, 
//         stars: Number,
//         views: Number,
//         globalViewRank:Number,
//         globalStarRank: Number){
        
//         this.link = link
//         this.name = name
//         this.stars = stars
//         this.views = views
//         this.globalViewRank = globalViewRank
//         this.globalStarRank = globalStarRank
//     }

// }

@Component({
    selector:'app-mnemonicCard',
    standalone:true,
    templateUrl:"./mnemonicCard.html",
    imports:[
        IonicModule,
        CommonModule
    ]
})
export class MnemonicCard {
    @Input() name: String = '';
    @Input() views: Number = 0;
}

export interface Mnemonic{
    link: String
    name: String
    stars: Number
    views: Number
    globalViewRank: Number
    globalStarRank: Number
    collectionViewRank?: Number
    collectionStarRank?: Number
}


// export class MnemonicCard {
//     mnemonic: Mnemonic | null = null;
//     constructor(){
//     }
//     ngOnInit(mnemonic: Mnemonic){
//         this.mnemonic = mnemonic;
//     }
// }

