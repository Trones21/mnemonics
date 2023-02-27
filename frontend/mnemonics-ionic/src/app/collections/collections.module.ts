import { IonicModule } from '@ionic/angular';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { CollectionsPage } from './collections.page';
import { ExploreContainerComponentModule } from '../explore-container/explore-container.module';

import { CollectionsPageRoutingModule } from './collections-routing.module';

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    ExploreContainerComponentModule,
    CollectionsPageRoutingModule
  ],
  declarations: [CollectionsPage]
})
export class CollectionsPageModule {}
