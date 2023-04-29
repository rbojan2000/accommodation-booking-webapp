import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NavbarComponent } from './modules/navbar/navbar.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { MatCardModule } from '@angular/material/card';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ToastrModule } from 'ngx-toastr';
import { ViewComponent } from './modules/view/view.component';
import { ReserveAccommodationComponent } from './modules/reserve-accommodation/reserve-accommodation.component';
import { MatDialogModule } from '@angular/material/dialog';
import { ReservationDetailsComponent } from './modules/reservation-details/reservation-details.component';

@NgModule({
  declarations: [
    AppComponent,
    NavbarComponent,
    ViewComponent,
    ReserveAccommodationComponent,
    ReservationDetailsComponent,
  ],
  imports: [
    BrowserAnimationsModule,
    ToastrModule.forRoot({
      positionClass: 'toast-top-right',
      preventDuplicates: true,
    }),
    FormsModule,
    BrowserModule,
    NgbModule,
    AppRoutingModule,
    MatCardModule,
    MatDialogModule,
    HttpClientModule,
  ],
  providers: [],

  bootstrap: [AppComponent],
})
export class AppModule {}
