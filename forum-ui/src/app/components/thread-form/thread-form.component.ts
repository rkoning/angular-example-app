import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Thread } from '@models/thread.model';
import { ThreadService } from '@services/thread.service';
import {MatSnackBar} from '@angular/material/snack-bar';
import { MatDialog } from '@angular/material/dialog';
import { ThreadRoutesService } from '@routes/thread-routes.service';

@Component({
  selector: 'app-thread-form',
  templateUrl: './thread-form.component.html',
  styleUrls: ['./thread-form.component.scss']
})
export class ThreadFormComponent implements OnInit {

  threadForm = new FormGroup({
    userId: new FormControl('', Validators.required),
    title: new FormControl('', Validators.required),
    text: new FormControl('', Validators.required),
  });

  constructor(private threadService: ThreadService, private threadRoutes: ThreadRoutesService,
              private snackBar: MatSnackBar, private dialog: MatDialog) { }

  ngOnInit(): void {

  }

  public addThread = (): void => {
    if (this.threadForm.valid) {
      const thread = this.threadForm.value as Thread;
      this.threadRoutes.add(thread).subscribe((res) => {
        this.threadService.addThread(res.data);
        this.snackBar.open('Thread posted!', 'Ok', { duration: 3000 });
      }, (err) => {
        console.log(err);
        this.snackBar.open('Error posting thread, try again later', 'Ok', { duration: 3000 });
      });
      this.dialog.closeAll();
    } else {
      this.snackBar.open('Invalid thread', 'Ok', { duration: 3000 });
    }
  }
}
