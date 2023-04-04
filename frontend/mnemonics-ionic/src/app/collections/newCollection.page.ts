import { Component, isStandalone } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { IonicModule } from '@ionic/angular';

@Component({
  selector: 'app-newCollection',
  standalone: true,
  templateUrl:'newCollection.page.html',
  imports:[
    IonicModule,
    FormsModule,
    ReactiveFormsModule
  ]
})
export class newCollectionPage {

  constructor() {}

}