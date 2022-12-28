insert into "public.materials"(material_name, cost_per_gram) VALUES ('Гранит', 20.10);
insert into "public.materials"(material_name, cost_per_gram) VALUES ('Железо', 0.23);
insert into "public.materials"(material_name, cost_per_gram) VALUES ('Дерево', 4.20);
insert into "public.materials"(material_name, cost_per_gram) VALUES ('Картон', 2.35);

insert into "public.materials"(material_name, cost_per_gram) VALUES ('тест', 321.12);

insert into "public.details"(detail_name, weight, material_name) VALUES ('Брусок', 50.12, 'Дерево');
insert into "public.details"(detail_name, weight, material_name) VALUES ('Коробка', 100.32, 'Картон');
insert into "public.details"(detail_name, weight, material_name) VALUES ('Маленькая коробка', 101, 'Картон');
insert into "public.details"(detail_name, weight, material_name) VALUES ('Обработанный гранит', 1000, 'Гранит');
insert into "public.details"(detail_name, weight, material_name) VALUES ('Шуруп', 5.21, 'Железо');

insert into "public.details"(detail_name, weight, material_name) VALUES ('тест', 321.12, 'тест');

insert into "public.products"(product_name) VALUES ('Деревянная лавка');
insert into "public.products"(product_name) VALUES ('Коробка шурупов');

insert into "public.products"(product_name) VALUES ('тест');

insert into "public.product_composition"(product_number, detail_name, details_amount) VALUES (1, 'Брусок', 20);
insert into "public.product_composition"(product_number, detail_name, details_amount) VALUES (1, 'Шуруп', 40);
insert into "public.product_composition"(product_number, detail_name, details_amount) VALUES (2, 'Коробка', 1);
insert into "public.product_composition"(product_number, detail_name, details_amount) VALUES (2, 'Шуруп', 200);

insert into "public.product_composition"(product_number, detail_name, details_amount) VALUES (3, 'Обработанный гранит', 50);
insert into "public.product_composition"(product_number, detail_name, details_amount) VALUES (3, 'Брусок', 10);

select * from "public.details";
select * from "public.materials";
select * from "public.product_composition";

select detail_name, details_amount from "public.product_composition" where product_number =1;




CREATE TABLE "public.product_composition" (
	"product_number" integer NOT NULL,
	"detail_name" varchar(50) NOT NULL,
	"details_amount" integer NOT NULL,
	CONSTRAINT "product_composition_pk" PRIMARY KEY ("product_number","detail_name")
) WITH (
  OIDS=FALSE
);
CREATE TABLE "public.products" (
	"product_number" integer NOT NULL,
	"product_name" varchar(50) NOT NULL,
	CONSTRAINT "products_pk" PRIMARY KEY ("product_number")
) WITH (
  OIDS=FALSE
);
CREATE TABLE "public.details" (
	"detail_name" varchar(50) NOT NULL,
	"weight" FLOAT NOT NULL,
	"material_name" varchar(50) NOT NULL,
	CONSTRAINT "details_pk" PRIMARY KEY ("detail_name")
) WITH (
  OIDS=FALSE
);
CREATE TABLE "public.materials" (
	"material_name" varchar(50) NOT NULL,
	"cost_per_gram" float NOT NULL,
	CONSTRAINT "materials_pk" PRIMARY KEY ("material_name")
) WITH (
  OIDS=FALSE
);
ALTER TABLE "public.product_composition" ADD CONSTRAINT "product_composition_fk0" FOREIGN KEY ("product_number") REFERENCES "public.products"("product_number");
ALTER TABLE "public.product_composition" ADD CONSTRAINT "product_composition_fk1" FOREIGN KEY ("detail_name") REFERENCES "public.details"("detail_name");
ALTER TABLE "public.details" ADD CONSTRAINT "details_fk0" FOREIGN KEY ("material_name") REFERENCES "public.materials"("material_name");