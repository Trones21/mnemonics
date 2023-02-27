import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { TabsPage } from './tabs.page';

const routes: Routes = [
  {
    path: '',
    component: TabsPage,
    children: [
      {
        path: 'swiper',
        loadChildren: () => import('../swiper/swiper.module').then(m => m.SwiperPageModule)
      },
      {
        path: 'collections',
        loadChildren: () => import('../collections/collections.module').then(m => m.CollectionsPageModule)
      },
      {
        path: 'profile',
        loadChildren: () => import('../profile/profile.module').then(m => m.ProfilePageModule)
      },
      {
        path: '',
        redirectTo: '/swiper',  //Basically which tab show be selected when first loading / empty path
        pathMatch: 'full'
      }
    ]
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
})
export class TabsPageRoutingModule {}
