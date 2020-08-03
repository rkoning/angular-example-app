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
      const comment = this.form.value as Comment;
      comment.parentId = this.threadId;
      comment.id = undefined;
      this.threadService.addComment(this.threadId, comment).subscribe((res) => {
        this.snackBar.open('Comment posted!', 'Ok', { duration: 3000 });
      }, (err) => {
        console.log(err);
        this.snackBar.open('Error posting comment, try again later', 'Ok', { duration: 3000 });
      });
    } else {
      this.snackBar.open('Invalid comment', undefined, {duration: 3000});
    }
  }
}
