import { CommonModule } from '@angular/common';
import { AfterViewInit, Component, ElementRef, QueryList, ViewChildren } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { GestureController, IonCard, IonicModule, Platform } from '@ionic/angular';
import { ExploreContainerComponentModule } from '../explore-container/explore-container.module';

@Component({
  selector: 'app-practice-swiper',
  standalone:true,
  template: `<ion-content class="ion-padding" [fullscreen]="true">
  <ion-header collapse="condense">
    <ion-toolbar>
      <ion-title size="large">Swiper</ion-title>
    </ion-toolbar>
  </ion-header>

  <ion-card *ngFor=" let m of mnemonics" [style.position]="'absolute'" [style.width]="'75%'" [style.height]="'75%'">
    <ion-card-title>{{m.MemoryTip}}</ion-card-title>
    <ion-card-content>{{m.ItemToRemember}}</ion-card-content>
  </ion-card>

</ion-content>`,
  imports:[
    IonicModule,
    CommonModule,
    FormsModule,
    ExploreContainerComponentModule,
  ]

})
export class SwiperPage implements AfterViewInit{

  mnemonics = JSON.parse(`[
    {
      "MnemonicID": 1,
      "MemoryTip": "Please Excuse My Dear Aunt Sally",
      "ItemToRemember": "Maths Order of Operations - Parentheses, Exponents ...",
      "ImageID": ""
    },
    {
      "MnemonicID": 2,
      "MemoryTip": "In fourteen hundred ninety-two, Columbus sailed the ocean blue",
      "ItemToRemember": "Date - That Columbus landed in America in 1492",
      "ImageID": ""
    },
    {
      "MnemonicID": 3,
      "MemoryTip": "Blah",
      "ItemToRemember": "sdgnjfdgfdgd dfdgdg",
      "ImageID": ""
    }
  ]`)
  
  @ViewChildren(IonCard, {read:ElementRef}) cards!:QueryList<ElementRef>;
  longpressActive = false;
  
  constructor(private gestureCtrl: GestureController, private plt: Platform) {
  }

  ngAfterViewInit(){
    const cardArray = this.cards.toArray();
    this.swipe(cardArray)
  }

  swipe(cardArray: any){
    for(let i= 0; i < cardArray.length; i++){
      const card = cardArray[i]
      const gesture = this.gestureCtrl.create({
        el: card.nativeElement,
        gestureName: 'swipe',
        onStart: ev => {
          
        },
        onMove: ev => {
          card.nativeElement.style.transform = `translateX(${ev.deltaX}px) rotate(${ev.deltaX / 10}deg)`; 
        },
        onEnd: ev => {
          card.nativeElement.style.transition = '.5s ease-out'

          //Fly out Right 
          if(ev.deltaX > 150){
            card.nativeElement.style.transform = `translateX(${+this.plt.width() * 2}px) rotate(${ev.deltaX / 2}deg)`

          //Fly out Left
          } else if( ev.deltaX < -150) {
            card.nativeElement.style.transform = `translateX(-${+this.plt.width() * 2}px) rotate(${ev.deltaX / 2}deg)`

          }else{
            card.nativeElement.style.transform = ''
          }

        },

        
      });
    
      gesture.enable();
    }
  }

}