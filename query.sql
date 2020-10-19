
CREATE SEQUENCE public.category_category_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE public.category
(
    category_id bigint NOT NULL DEFAULT nextval('category_category_id_seq'::regclass),
    category_name VARCHAR(255),
    CONSTRAINT category_id_pkey PRIMARY KEY (category_id)
)

WITH (
    OIDS=FALSE
);

INSERT INTO "public"."category" ("category_id", "category_name") VALUES ('3', 'Shirt');
INSERT INTO "public"."category" ("category_id", "category_name") VALUES ('2', 'Computers');
INSERT INTO "public"."category" ("category_id", "category_name") VALUES ('1', 'Electronics');
INSERT INTO "public"."category" ("category_id", "category_name") VALUES ('4', 'Accessory');
INSERT INTO "public"."category" ("category_id", "category_name") VALUES ('5', 'Phone');


CREATE SEQUENCE public.product_product_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE public.product
(
    product_id bigint NOT NULL DEFAULT nextval('product_product_id_seq'::regclass),
    product_name VARCHAR(255),
    CONSTRAINT product_id_pkey PRIMARY KEY (product_id)
)

WITH (
    OIDS=FALSE
);

INSERT INTO "public"."product" ("product_id", "product_name") VALUES ('1', 'Leather Bag');
INSERT INTO "public"."product" ("product_id", "product_name") VALUES ('3', 'IPhone SE');
INSERT INTO "public"."product" ("product_id", "product_name") VALUES ('4', 'Asus Zenphone');
INSERT INTO "public"."product" ("product_id", "product_name") VALUES ('5', 'Monitor HP 30 Inch');
INSERT INTO "public"."product" ("product_id", "product_name") VALUES ('6', 'Army Shirt');
INSERT INTO "public"."product" ("product_id", "product_name") VALUES ('2', 'Sport Shoes');

CREATE SEQUENCE public.product_category_product_category_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE public.product_category
(
    product_category_id bigint NOT NULL DEFAULT nextval('product_category_product_category_id_seq'::regclass),
    category_id bigint NOT NULL,
		product_id bigint NOT NULL,
    CONSTRAINT product_category_id_pkey PRIMARY KEY (product_category_id)
)

WITH (
    OIDS=FALSE
);

INSERT INTO "public"."product_category" ("product_category_id", "category_id", "product_id") VALUES ('1', '4', '1');
INSERT INTO "public"."product_category" ("product_category_id", "category_id", "product_id") VALUES ('2', '4', '2');
INSERT INTO "public"."product_category" ("product_category_id", "category_id", "product_id") VALUES ('3', '5', '3');
INSERT INTO "public"."product_category" ("product_category_id", "category_id", "product_id") VALUES ('4', '5', '4');
INSERT INTO "public"."product_category" ("product_category_id", "category_id", "product_id") VALUES ('5', '2', '5');

CREATE SEQUENCE public.user_cart_user_cart_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE public.user_cart
(
    user_cart_id bigint NOT NULL DEFAULT nextval('user_cart_user_cart_id_seq'::regclass),
    user_id bigint NOT NULL,
		is_active BOOLEAN DEFAULT TRUE,
    is_checkout BOOLEAN DEFAULT FALSE,
    is_paid BOOLEAN DEFAULT FALSE,
    CONSTRAINT user_cart_id_pkey PRIMARY KEY (user_cart_id)
)

WITH (
    OIDS=FALSE
);

CREATE SEQUENCE public.user_cart_detail_product_user_cart_detail_product_id_seq
    INCREMENT 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    START 1
    CACHE 1;

CREATE TABLE public.user_cart_detail_product
(
    user_cart_detail_product_id bigint NOT NULL DEFAULT nextval('user_cart_detail_product_user_cart_detail_product_id_seq'::regclass),
    user_cart_id bigint NOT NULL,
    product_id bigint NOT NULL,
		total bigint NOT NULL,
    CONSTRAINT user_cart_detail_product_id_pkey PRIMARY KEY (user_cart_detail_product_id)
)

WITH (
    OIDS=FALSE
);
