import { AfterViewInit, Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { Accommodation } from 'src/app/model/accommodation';
import { Address } from 'src/app/model/address';
import { Benefit } from 'src/app/model/benefit';
import { AccommodationService } from 'src/app/services/accommodation/accommodation.service';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';

@Component({
  selector: 'app-view',
  templateUrl: './view.component.html',
  styleUrls: ['./view.component.css'],
})
export class ViewComponent implements OnInit {
  public accommodations: Accommodation[] = [];
  public startDate: any;  
  public endDate: any;


  constructor(
    private accommodationService: AccommodationService,
    public dialog: MatDialog,
    private router: Router
  ) {}

  ngOnInit() {
    this.accommodationService.getAccommodations().subscribe((res: any) => {
      console.log(res);
      this.accommodations = res.accommodations;
      console.log(this.accommodations);
    });
  }

  reserveAccommodation(acc: Accommodation) {
    this.router.navigate(['/accommodation-details'], {
      queryParams: {
        accommodationID: acc.id,
        startDate: this.startDate,
        endDate: this.endDate,
      },
    });
  }

  createAccommodation() {}
}
