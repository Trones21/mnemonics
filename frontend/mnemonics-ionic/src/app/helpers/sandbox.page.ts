import { CommonModule } from "@angular/common";
import { Component, Input } from "@angular/core";
import { IonicModule} from '@ionic/angular';
import { SingleCollectionPage } from "../collections/singleCollection.page";
import { MnemonicCard } from "../components/mnemonicCard";

@Component({
    selector: 'app-sandbox',
    standalone: true,
    template: `
   <!-- <app-mnemonicCard name='Test'> -->
    <!-- <button (click)="refresh()">Refresh</button> -->
    <app-single-collection id={{str}}></app-single-collection>,
    `,
    imports: [
        IonicModule,
        SingleCollectionPage,
        MnemonicCard
    ]
})
export class Sandbox {
    str = 'sdfsdf'
    ids = ['1' ,'2', '435']
    currentId = 0;
    constructor(){}


    refresh(){
        console.log('refresh called')
        if(this.currentId === this.ids.length - 1){
            this.str = this.ids[0]
        }else{
            this.str = this.ids[this.currentId + 1]
        }
    }

}