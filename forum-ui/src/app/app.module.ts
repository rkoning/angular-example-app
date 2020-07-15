import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppRoutingModule } from './app-routing.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { MatCardModule } from '@angular/material/card';
import { MatDialogModule } from '@angular/material/dialog';
import { MatInputModule } from '@angular/material/input';
import { MatSnackBar } from '@angular/material/snack-bar';
import { MatButtonModule } from '@angular/material/button';
import { MatListModule } from '@angular/material/list';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';

import { AppComponent } from './app.component';
import { ForumComponent } from './components/forum/forum.component';
import { ThreadComponent } from './components/thread/thread.component';
import { ThreadFormComponent } from './components/thread-form/thread-form.component';
import { CommentComponent } from './components/comment/comment.component';
import { ThreadCardComponent } from './components/thread-card/thread-card.component';
import { CommentFormComponent } from './components/comment-form/comment-form.component';

@NgModule({
  declarations: [
    AppComponent,
    ForumComponent,
    ThreadComponent,
    CommentComponent,
    ThreadFormComponent,
    ThreadCardComponent,
    CommentFormComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatCardModule,
    MatDialogModule,
    MatInputModule,
    MatListModule,
    MatProgressSpinnerModule,
    ReactiveFormsModule,
    HttpClientModule,
  ],
  providers: [
    MatSnackBar,
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
