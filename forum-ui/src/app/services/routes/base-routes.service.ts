import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class BaseRoutesService<T> {

  public apiRoot = 'localhost:80';
  public baseRoute: string;

  constructor(protected http: HttpClient) { }

  public get(itemId: string): Observable<T> {
    return this.http.get<T>(`${this.apiRoot}/${this.baseRoute}/${itemId}`);
  }

  public list(): Observable<T[]> {
    return this.http.get<T[]>(`${this.apiRoot}/${this.baseRoute}`);
  }

  public add(jsonBody: Partial<T>): Observable<T> {
    return this.http.post<T>(`${this.apiRoot}/${this.baseRoute}`, jsonBody);
  }

  public update(jsonBody: Partial<T>): Observable<T> {
    return this.http.put<T>(`${this.apiRoot}/${this.baseRoute}`, jsonBody);
  }

  public delete(itemId: string): Observable<T> {
    return this.http.delete<T>(`${this.apiRoot}/${this.baseRoute}/${itemId}`)
  }
}
