import { IonicModule } from '@ionic/angular';
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { SwiperPage } from './swiper.page';
import { ExploreContainerComponentModule } from '../explore-container/explore-container.module';

import {SwiperPageRoutingModule} from './swiper-routing.module'

@NgModule({
  imports: [
    IonicModule,
    CommonModule,
    FormsModule,
    ExploreContainerComponentModule,
    SwiperPageRoutingModule
  ],
  declarations: [SwiperPage]
})
export class SwiperPageModule {}