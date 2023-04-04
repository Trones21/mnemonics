import { IonicModule } from '@ionic/angular';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CollectionsPage } from './collections.page';
import { ExploreContainerComponentModule } from '../explore-container/explore-container.module';

import { CollectionsPageRoutingModule } from './collections-routing.module';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    ExploreContainerComponentModule,
    CollectionsPageRoutingModule
  ],
  declarations: [CollectionsPage]
})
export class CollectionsPageModule {
  
}
