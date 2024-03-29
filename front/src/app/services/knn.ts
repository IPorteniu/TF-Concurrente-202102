import { Injectable } from '@angular/core';
import {HttpClient} from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class KnnService {
  constructor(private http: HttpClient) { }

  postKnn(datos: any) {
    return this.http.post(`http://localhost:9080/api/agregar`, datos);
  }
}
