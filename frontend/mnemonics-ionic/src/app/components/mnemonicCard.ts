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
    @Input() mnemonic!: Mnemonic;
}

export interface Mnemonic{
    link: string
    name: string
    stars: number
    views: number
    hint: string
    answer: string
    globalViewRank: number
    globalStarRank: number
    collectionViewRank?: number
    collectionStarRank?: number
    createdAt: IonDatetime
    updatedAt: IonDatetime
}
