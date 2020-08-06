import { Component, OnInit, Input } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { MatSnackBar } from '@angular/material/snack-bar';
import { Comment } from '@models/comment.model';
import { ThreadRoutesService } from '@routes/thread-routes.service';
import { ThreadService } from '@services/thread.service';

@Component({
  selector: 'app-comment-form',
  templateUrl: './comment-form.component.html',
  styleUrls: ['./comment-form.component.scss']
})
export class CommentFormComponent implements OnInit {
  @Input() threadId: string;

  public form = new FormGroup({
    userId: new FormControl('', Validators.required),
    text: new FormControl('', Validators.required),
  });


  constructor(private threadService: ThreadService, private threadRoutes: ThreadRoutesService, private snackBar: MatSnackBar) { }

  ngOnInit(): void {
  }

  public addComment = (): void => {
    if (this.form.valid) {

    } else {
      this.snackBar.open('Invalid comment', undefined, {duration: 3000});
    }
  }
}
