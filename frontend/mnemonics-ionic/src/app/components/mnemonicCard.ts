import { CommonModule } from "@angular/common";
import { Component, Inject, Injectable, Input } from "@angular/core";
import { IonicModule } from '@ionic/angular';
import { IonDatetime } from "@ionic/angular";
import { MnemonicService } from "../mnemonic/mnemonic.service";


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
    // @Input() name: string = '';
    // @Input() views: number = 0;
    @Input() m!: Mnemonic;
}

export interface Mnemonic{
    link: string
    name: string
    stars: number
    views: number
    globalViewRank: number
    globalStarRank: number
    collectionViewRank?: number
    collectionStarRank?: number
}


// export class MnemonicCard {
//     mnemonic: Mnemonic | null = null;
//     constructor(){
//     }
//     ngOnInit(mnemonic: Mnemonic){
//         this.mnemonic = mnemonic;
//     }
// }

