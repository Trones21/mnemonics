import { CommonModule } from "@angular/common";
import { Component, Input } from "@angular/core";
import { IonicModule } from '@ionic/angular';
import { IonDatetime } from "@ionic/angular";



export class Mnemonic{
    link: string
    name: string
    stars: number
    views: number
    globalViewRank: number
    globalStarRank: number
    //Collection Ranks must be computed client side (due to planned db schema -- a single mnemonic can be in many collections,
    //& I dont plan to keep track of collection level ranks in DB)
    collectionViewRank?: number
    collectionStarRank?: number
    constructor(
        link: string, 
        name: string, 
        stars: number,
        views: number,
        globalViewRank:number,
        globalStarRank: number){
        
        this.link = link
        this.name = name
        this.stars = stars
        this.views = views
        this.globalViewRank = globalViewRank
        this.globalStarRank = globalStarRank
    }

}


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
    @Input() mnemonic: Mnemonic = new Mnemonic('/', 'Mnemonic Name', 0, 0, 47, 52)
    constructor(){}



}