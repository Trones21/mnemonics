import { CommonModule } from "@angular/common";
import { Component, Input } from "@angular/core";
import { RouterLink } from "@angular/router";
import { IonicModule} from '@ionic/angular';
import { SwiperPage } from "../swiper/swiper.page";

@Component({
    selector:'app-practice',
    standalone: true,
    template:`
    <ion-header>Practice</ion-header>
    note -- not the actual UI, just trying to get a sense of how this will work 
    <ion-item>
    <ion-label>Activity</ion-label>
    <ion-select placeholder="Select One">
      <ion-select-option value="swiper">Swiper</ion-select-option>
      <ion-select-option value="matcher">Matcher</ion-select-option>
    </ion-select>
   </ion-item>
   <ion-item>
    <ion-label>Collection</ion-label>
    <ion-select placeholder="Select One">
      <ion-select-option value="">Swiper</ion-select-option>
      <ion-select-option value="matcher">Matcher</ion-select-option>
    </ion-select>
   </ion-item>
    
    `,
    imports:[
        IonicModule,
        SwiperPage
    ]
})
export class PracticeSetupPage{
    constructor(){}
}