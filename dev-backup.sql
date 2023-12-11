CREATE TABLE public.additional_item (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    price bigint NOT NULL,
    description character varying(255) NOT NULL,
    enabled boolean NOT NULL,
    additional_items_group_id bigint NOT NULL
);


ALTER TABLE public.additional_item OWNER TO cardap;

--
-- TOC entry 197 (class 1259 OID 37985)
-- Name: additional_item_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.additional_item_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.additional_item_id_seq OWNER TO cardap;

--
-- TOC entry 3058 (class 0 OID 0)
-- Dependencies: 197
-- Name: additional_item_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.additional_item_id_seq OWNED BY public.additional_item.id;


--
-- TOC entry 224 (class 1259 OID 38129)
-- Name: additional_items_group; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.additional_items_group (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    description character varying(255),
    "order" integer NOT NULL,
    minimum_items integer NOT NULL,
    maximum_items integer NOT NULL,
    enabled boolean NOT NULL,
    company_id bigint NOT NULL
);


ALTER TABLE public.additional_items_group OWNER TO cardap;

--
-- TOC entry 223 (class 1259 OID 38127)
-- Name: additional_items_group_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.additional_items_group_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.additional_items_group_id_seq OWNER TO cardap;

--
-- TOC entry 3059 (class 0 OID 0)
-- Dependencies: 223
-- Name: additional_items_group_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.additional_items_group_id_seq OWNED BY public.additional_items_group.id;


--
-- TOC entry 200 (class 1259 OID 37999)
-- Name: address; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.address (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    address character varying(255) NOT NULL,
    area character varying(255),
    city character varying(255),
    state character varying(255),
    postal_code character varying(255),
    complement character varying(255),
    number character varying(10) NOT NULL,
    latitude character varying(255),
    longitude character varying(255),
    company_id bigint NOT NULL
);


ALTER TABLE public.address OWNER TO cardap;

--
-- TOC entry 199 (class 1259 OID 37997)
-- Name: address_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.address_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.address_id_seq OWNER TO cardap;

--
-- TOC entry 3060 (class 0 OID 0)
-- Dependencies: 199
-- Name: address_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.address_id_seq OWNED BY public.address.id;


--
-- TOC entry 222 (class 1259 OID 38121)
-- Name: cardap_table; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.cardap_table (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    company_id bigint NOT NULL
);


ALTER TABLE public.cardap_table OWNER TO cardap;

--
-- TOC entry 221 (class 1259 OID 38119)
-- Name: cardap_table_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.cardap_table_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cardap_table_id_seq OWNER TO cardap;

--
-- TOC entry 3061 (class 0 OID 0)
-- Dependencies: 221
-- Name: cardap_table_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.cardap_table_id_seq OWNED BY public.cardap_table.id;


--
-- TOC entry 220 (class 1259 OID 38109)
-- Name: cardap_user; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.cardap_user (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    login character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    company_id bigint
);


ALTER TABLE public.cardap_user OWNER TO cardap;

--
-- TOC entry 219 (class 1259 OID 38107)
-- Name: cardap_user_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.cardap_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cardap_user_id_seq OWNER TO cardap;

--
-- TOC entry 3062 (class 0 OID 0)
-- Dependencies: 219
-- Name: cardap_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.cardap_user_id_seq OWNED BY public.cardap_user.id;


--
-- TOC entry 202 (class 1259 OID 38011)
-- Name: category; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.category (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    description character varying(255),
    "order" integer NOT NULL,
    enabled boolean NOT NULL,
    menu_id integer
);


ALTER TABLE public.category OWNER TO cardap;

--
-- TOC entry 201 (class 1259 OID 38009)
-- Name: category_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.category_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.category_id_seq OWNER TO cardap;

--
-- TOC entry 3063 (class 0 OID 0)
-- Dependencies: 201
-- Name: category_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.category_id_seq OWNED BY public.category.id;


--
-- TOC entry 218 (class 1259 OID 38099)
-- Name: clients; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.clients (
    phone bigint NOT NULL,
    name character varying(255) NOT NULL,
    address character varying(255),
    area character varying(255),
    city character varying(255),
    state character varying(255),
    postal_code character varying(255),
    complement character varying(255),
    number character varying(10),
    landmark character varying(255)
);


ALTER TABLE public.clients OWNER TO cardap;

--
-- TOC entry 206 (class 1259 OID 38033)
-- Name: company; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.company (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    description character varying(255),
    image_path character varying(255),
    phone_number character varying(255),
    company_code character varying(255),
    whats_app character varying(255),
    has_delivery boolean,
    has_withdrawn boolean,
    theme character varying(255),
    user_id bigint
);


ALTER TABLE public.company OWNER TO cardap;

--
-- TOC entry 205 (class 1259 OID 38031)
-- Name: company_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.company_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.company_id_seq OWNER TO cardap;

--
-- TOC entry 3064 (class 0 OID 0)
-- Dependencies: 205
-- Name: company_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.company_id_seq OWNED BY public.company.id;


--
-- TOC entry 204 (class 1259 OID 38026)
-- Name: company_payment_type; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.company_payment_type (
    company_id integer NOT NULL,
    payment_type_id integer NOT NULL
);


ALTER TABLE public.company_payment_type OWNER TO cardap;

--
-- TOC entry 203 (class 1259 OID 38021)
-- Name: company_section; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.company_section (
    company_id integer NOT NULL,
    section_id integer NOT NULL
);


ALTER TABLE public.company_section OWNER TO cardap;

--
-- TOC entry 208 (class 1259 OID 38045)
-- Name: menu; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.menu (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    description character varying(255),
    enabled boolean,
    minimum_order bigint,
    company_id bigint
);


ALTER TABLE public.menu OWNER TO cardap;

--
-- TOC entry 207 (class 1259 OID 38043)
-- Name: menu_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.menu_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.menu_id_seq OWNER TO cardap;

--
-- TOC entry 3065 (class 0 OID 0)
-- Dependencies: 207
-- Name: menu_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.menu_id_seq OWNED BY public.menu.id;


--
-- TOC entry 196 (class 1259 OID 37980)
-- Name: migrations; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.migrations (
    id character varying(255) NOT NULL
);


ALTER TABLE public.migrations OWNER TO cardap;

--
-- TOC entry 210 (class 1259 OID 38054)
-- Name: opening_hours; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.opening_hours (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    start_day character varying(255) NOT NULL,
    end_day character varying(255) NOT NULL,
    start_time time(4) without time zone NOT NULL,
    end_time time(4) without time zone NOT NULL,
    company_id bigint NOT NULL
);


ALTER TABLE public.opening_hours OWNER TO cardap;

--
-- TOC entry 209 (class 1259 OID 38052)
-- Name: opening_hours_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.opening_hours_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.opening_hours_id_seq OWNER TO cardap;

--
-- TOC entry 3066 (class 0 OID 0)
-- Dependencies: 209
-- Name: opening_hours_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.opening_hours_id_seq OWNED BY public.opening_hours.id;


--
-- TOC entry 212 (class 1259 OID 38066)
-- Name: payment_type; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.payment_type (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL
);


ALTER TABLE public.payment_type OWNER TO cardap;

--
-- TOC entry 211 (class 1259 OID 38064)
-- Name: payment_type_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.payment_type_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.payment_type_id_seq OWNER TO cardap;

--
-- TOC entry 3067 (class 0 OID 0)
-- Dependencies: 211
-- Name: payment_type_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.payment_type_id_seq OWNED BY public.payment_type.id;


--
-- TOC entry 215 (class 1259 OID 38080)
-- Name: product; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.product (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    description character varying(255),
    price bigint,
    "order" integer,
    enabled boolean,
    image character varying(255),
    adults_only boolean,
    category_id bigint NOT NULL
);


ALTER TABLE public.product OWNER TO cardap;

--
-- TOC entry 213 (class 1259 OID 38073)
-- Name: product_additional_group; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.product_additional_group (
    product_id integer NOT NULL,
    additional_items_group_id integer NOT NULL
);


ALTER TABLE public.product_additional_group OWNER TO cardap;

--
-- TOC entry 214 (class 1259 OID 38078)
-- Name: product_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.product_id_seq OWNER TO cardap;

--
-- TOC entry 3068 (class 0 OID 0)
-- Dependencies: 214
-- Name: product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.product_id_seq OWNED BY public.product.id;


--
-- TOC entry 217 (class 1259 OID 38092)
-- Name: section; Type: TABLE; Schema: public; Owner: cardap
--

CREATE TABLE public.section (
    id integer NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL
);


ALTER TABLE public.section OWNER TO cardap;

--
-- TOC entry 216 (class 1259 OID 38090)
-- Name: section_id_seq; Type: SEQUENCE; Schema: public; Owner: cardap
--

CREATE SEQUENCE public.section_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.section_id_seq OWNER TO cardap;

--
-- TOC entry 3069 (class 0 OID 0)
-- Dependencies: 216
-- Name: section_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: cardap
--

ALTER SEQUENCE public.section_id_seq OWNED BY public.section.id;


--
-- TOC entry 2835 (class 2604 OID 37990)
-- Name: additional_item id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.additional_item ALTER COLUMN id SET DEFAULT nextval('public.additional_item_id_seq'::regclass);


--
-- TOC entry 2846 (class 2604 OID 38132)
-- Name: additional_items_group id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.additional_items_group ALTER COLUMN id SET DEFAULT nextval('public.additional_items_group_id_seq'::regclass);


--
-- TOC entry 2836 (class 2604 OID 38002)
-- Name: address id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.address ALTER COLUMN id SET DEFAULT nextval('public.address_id_seq'::regclass);


--
-- TOC entry 2845 (class 2604 OID 38124)
-- Name: cardap_table id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.cardap_table ALTER COLUMN id SET DEFAULT nextval('public.cardap_table_id_seq'::regclass);


--
-- TOC entry 2844 (class 2604 OID 38112)
-- Name: cardap_user id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.cardap_user ALTER COLUMN id SET DEFAULT nextval('public.cardap_user_id_seq'::regclass);


--
-- TOC entry 2837 (class 2604 OID 38014)
-- Name: category id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.category ALTER COLUMN id SET DEFAULT nextval('public.category_id_seq'::regclass);


--
-- TOC entry 2838 (class 2604 OID 38036)
-- Name: company id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.company ALTER COLUMN id SET DEFAULT nextval('public.company_id_seq'::regclass);


--
-- TOC entry 2839 (class 2604 OID 38048)
-- Name: menu id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.menu ALTER COLUMN id SET DEFAULT nextval('public.menu_id_seq'::regclass);


--
-- TOC entry 2840 (class 2604 OID 38057)
-- Name: opening_hours id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.opening_hours ALTER COLUMN id SET DEFAULT nextval('public.opening_hours_id_seq'::regclass);


--
-- TOC entry 2841 (class 2604 OID 38069)
-- Name: payment_type id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.payment_type ALTER COLUMN id SET DEFAULT nextval('public.payment_type_id_seq'::regclass);


--
-- TOC entry 2842 (class 2604 OID 38083)
-- Name: product id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.product ALTER COLUMN id SET DEFAULT nextval('public.product_id_seq'::regclass);


--
-- TOC entry 2843 (class 2604 OID 38095)
-- Name: section id; Type: DEFAULT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.section ALTER COLUMN id SET DEFAULT nextval('public.section_id_seq'::regclass);


--
-- TOC entry 3025 (class 0 OID 37987)
-- Dependencies: 198
-- Data for Name: additional_item; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.additional_item (id, created_at, updated_at, deleted_at, name, price, description, enabled, additional_items_group_id) VALUES (1, '2020-08-06 13:44:24.690475+00', '2020-08-06 13:44:24.690475+00', NULL, 'Queijo extra', 100, 'Queijo mussarela', true, 1);
INSERT INTO public.additional_item (id, created_at, updated_at, deleted_at, name, price, description, enabled, additional_items_group_id) VALUES (2, '2020-08-06 13:44:24.697377+00', '2020-08-06 13:44:24.697377+00', NULL, 'Presunto extra', 100, 'Apresuntado do mais barato', false, 1);
INSERT INTO public.additional_item (id, created_at, updated_at, deleted_at, name, price, description, enabled, additional_items_group_id) VALUES (3, '2020-08-06 13:44:24.786096+00', '2020-08-06 13:44:24.786096+00', NULL, 'Gelo', 100, 'Mais gelo', true, 2);
INSERT INTO public.additional_item (id, created_at, updated_at, deleted_at, name, price, description, enabled, additional_items_group_id) VALUES (4, '2020-08-06 13:44:24.793823+00', '2020-08-06 13:44:24.793823+00', NULL, 'Limão', 100, 'Limãozinho pra coquinha', false, 2);
INSERT INTO public.additional_item (id, created_at, updated_at, deleted_at, name, price, description, enabled, additional_items_group_id) VALUES (5, '2020-08-06 13:44:24.803716+00', '2020-08-06 13:44:24.803716+00', NULL, 'Açúcar', 100, 'Aqui até o açúcar é cobrado', false, 2);


--
-- TOC entry 3051 (class 0 OID 38129)
-- Dependencies: 224
-- Data for Name: additional_items_group; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.additional_items_group (id, created_at, updated_at, deleted_at, name, description, "order", minimum_items, maximum_items, enabled, company_id) VALUES (1, '2020-08-06 13:44:24.683098+00', '2020-08-06 13:44:25.107871+00', NULL, 'Opcionais de Hamburguer', 'Só para os gordoes que so o lanche nao basta', 0, 1, 10, true, 1);
INSERT INTO public.additional_items_group (id, created_at, updated_at, deleted_at, name, description, "order", minimum_items, maximum_items, enabled, company_id) VALUES (2, '2020-08-06 13:44:24.778616+00', '2020-08-06 13:44:25.131116+00', NULL, 'Opcionais de Bebidas', 'Além de comer, beba até passar mal', 0, 1, 10, true, 1);


--
-- TOC entry 3027 (class 0 OID 37999)
-- Dependencies: 200
-- Data for Name: address; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.address (id, created_at, updated_at, deleted_at, address, area, city, state, postal_code, complement, number, latitude, longitude, company_id) VALUES (1, '2020-08-06 13:44:24.546432+00', '2020-08-06 13:44:24.992509+00', NULL, 'Rua Rio Grande do Sul', 'Vila Sao Geraldo', 'Taubate', 'SP', '12012080', 'aaa', '98', '', '', 1);


--
-- TOC entry 3049 (class 0 OID 38121)
-- Dependencies: 222
-- Data for Name: cardap_table; Type: TABLE DATA; Schema: public; Owner: cardap
--



--
-- TOC entry 3047 (class 0 OID 38109)
-- Dependencies: 220
-- Data for Name: cardap_user; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.cardap_user (id, created_at, updated_at, deleted_at, name, login, email, password, company_id) VALUES (1, '2020-07-04 23:31:53.756+00', '2020-07-04 23:31:53.734+00', NULL, 'cardap Admin', 'cardap-admin', 'falecom@cardap.in', 'edbf177c911e9bb38bbc9ab1102a8e81', NULL);
INSERT INTO public.cardap_user (id, created_at, updated_at, deleted_at, name, login, email, password, company_id) VALUES (2, '2020-08-06 13:44:24.342871+00', '2020-08-06 13:44:24.342871+00', NULL, 'Vitor Nunes Benedito Mamutao', 'mamute', 'mamute@lesadao.com', '202cb962ac59075b964b07152d234b70', 1);


--
-- TOC entry 3029 (class 0 OID 38011)
-- Dependencies: 202
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.category (id, created_at, updated_at, deleted_at, name, description, "order", enabled, menu_id) VALUES (1, '2020-08-06 13:44:25.0731+00', '2020-08-06 13:44:25.0731+00', NULL, 'Lanches', 'Lanches Nervosos', 0, true, 1);
INSERT INTO public.category (id, created_at, updated_at, deleted_at, name, description, "order", enabled, menu_id) VALUES (2, '2020-08-06 13:44:25.118897+00', '2020-08-06 13:44:25.118897+00', NULL, 'Bebidas', 'Bebidas que destroem seu estomago', 1, false, 1);


--
-- TOC entry 3045 (class 0 OID 38099)
-- Dependencies: 218
-- Data for Name: clients; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.clients (phone, name, address, area, city, state, postal_code, complement, number, landmark) VALUES (12997861092, 'Vitor Nunes', 'Rua General Carlos Mamute, n 24', 'Jardim Satélite', 'São José dos Campos', 'SP', '12345-000', 'Zona Sul', '', 'Próximo a Av. Perseu');


--
-- TOC entry 3033 (class 0 OID 38033)
-- Dependencies: 206
-- Data for Name: company; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.company (id, created_at, updated_at, deleted_at, name, description, image_path, phone_number, company_code, whats_app, has_delivery, has_withdrawn, theme, user_id) VALUES (1, '2020-08-06 13:44:24.325782+00', '2020-08-06 13:44:24.952647+00', NULL, 'Bar do Zé', 'aaa', 'mala1.png', '12997861092', 'bar-do-ze', '12997861092', true, true, 'green', 0);


--
-- TOC entry 3031 (class 0 OID 38026)
-- Dependencies: 204
-- Data for Name: company_payment_type; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.company_payment_type (company_id, payment_type_id) VALUES (1, 11);
INSERT INTO public.company_payment_type (company_id, payment_type_id) VALUES (1, 7);
INSERT INTO public.company_payment_type (company_id, payment_type_id) VALUES (1, 5);
INSERT INTO public.company_payment_type (company_id, payment_type_id) VALUES (1, 2);


--
-- TOC entry 3030 (class 0 OID 38021)
-- Dependencies: 203
-- Data for Name: company_section; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.company_section (company_id, section_id) VALUES (1, 5);
INSERT INTO public.company_section (company_id, section_id) VALUES (1, 2);


--
-- TOC entry 3035 (class 0 OID 38045)
-- Dependencies: 208
-- Data for Name: menu; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.menu (id, created_at, updated_at, deleted_at, description, enabled, minimum_order, company_id) VALUES (1, '2020-08-06 13:44:25.063922+00', '2020-08-06 13:44:25.063922+00', NULL, 'Menu de Hambuguers', true, 10000, 1);


--
-- TOC entry 3023 (class 0 OID 37980)
-- Dependencies: 196
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.migrations (id) VALUES ('1');
INSERT INTO public.migrations (id) VALUES ('2');
INSERT INTO public.migrations (id) VALUES ('3');
INSERT INTO public.migrations (id) VALUES ('4');
INSERT INTO public.migrations (id) VALUES ('5');
INSERT INTO public.migrations (id) VALUES ('6');
INSERT INTO public.migrations (id) VALUES ('7');
INSERT INTO public.migrations (id) VALUES ('8');
INSERT INTO public.migrations (id) VALUES ('9');


--
-- TOC entry 3037 (class 0 OID 38054)
-- Dependencies: 210
-- Data for Name: opening_hours; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.opening_hours (id, created_at, updated_at, deleted_at, start_day, end_day, start_time, end_time, company_id) VALUES (2, '2020-08-06 13:44:24.60594+00', '2020-08-06 13:44:25.051924+00', NULL, 'THURSDAY', 'SATURDAY', '08:00:00', '23:00:00', 1);
INSERT INTO public.opening_hours (id, created_at, updated_at, deleted_at, start_day, end_day, start_time, end_time, company_id) VALUES (1, '2020-08-06 13:44:24.598549+00', '2020-08-06 13:44:25.058283+00', NULL, 'SUNDAY', 'WEDNESDAY', '08:00:00', '19:00:00', 1);


--
-- TOC entry 3039 (class 0 OID 38066)
-- Dependencies: 212
-- Data for Name: payment_type; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (1, NULL, NULL, NULL, 'Mastercard Crédito');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (3, NULL, NULL, NULL, 'Visa Crédito');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (4, NULL, NULL, NULL, 'Visa Débito');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (6, NULL, NULL, NULL, 'Elo Débito');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (8, NULL, NULL, NULL, 'Alelo VA - Voucher');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (9, NULL, NULL, NULL, 'Sodexo VR - Voucher');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (10, NULL, NULL, NULL, 'Sodexo VA - Voucher');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (12, NULL, NULL, NULL, 'Picpay');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (13, NULL, NULL, NULL, 'MercadoPago');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (11, NULL, '2020-08-06 13:44:24.999185+00', NULL, 'Dinheiro');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (7, NULL, '2020-08-06 13:44:25.013098+00', NULL, 'Alelo VR - Voucher');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (5, NULL, '2020-08-06 13:44:25.026911+00', NULL, 'Elo Crédito');
INSERT INTO public.payment_type (id, created_at, updated_at, deleted_at, name) VALUES (2, NULL, '2020-08-06 13:44:25.038453+00', NULL, 'Mastercard Débito');


--
-- TOC entry 3042 (class 0 OID 38080)
-- Dependencies: 215
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.product (id, created_at, updated_at, deleted_at, name, description, price, "order", enabled, image, adults_only, category_id) VALUES (1, '2020-08-06 13:44:25.080039+00', '2020-08-06 13:44:25.080039+00', NULL, 'Hamburgao nervosíssimo', 'Coma um e cague ate as tripas', 20000, 0, true, 'mala2.png', false, 1);
INSERT INTO public.product (id, created_at, updated_at, deleted_at, name, description, price, "order", enabled, image, adults_only, category_id) VALUES (2, '2020-08-06 13:44:25.101546+00', '2020-08-06 13:44:25.101546+00', NULL, 'Hamburgao caro pra carai', 'Seu lanche gourmet sem nada mas caro pra cacete pra postar no insta', 30000, 0, true, 'mala3.png', false, 1);
INSERT INTO public.product (id, created_at, updated_at, deleted_at, name, description, price, "order", enabled, image, adults_only, category_id) VALUES (3, '2020-08-06 13:44:25.124902+00', '2020-08-06 13:44:25.124902+00', NULL, 'Cachacinca Marvada', 'A Verdadeira Cachaça dos pingaiada de bar', 1000, 0, true, 'mala4.png', true, 2);
INSERT INTO public.product (id, created_at, updated_at, deleted_at, name, description, price, "order", enabled, image, adults_only, category_id) VALUES (4, '2020-08-06 13:44:25.143983+00', '2020-08-06 13:44:25.143983+00', NULL, 'Suco de Batata Doce', 'A bebida dos crosfiteiros', 30000, 1, true, 'mala5.png', true, 2);


--
-- TOC entry 3040 (class 0 OID 38073)
-- Dependencies: 213
-- Data for Name: product_additional_group; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.product_additional_group (product_id, additional_items_group_id) VALUES (1, 1);
INSERT INTO public.product_additional_group (product_id, additional_items_group_id) VALUES (2, 1);
INSERT INTO public.product_additional_group (product_id, additional_items_group_id) VALUES (3, 2);


--
-- TOC entry 3044 (class 0 OID 38092)
-- Dependencies: 217
-- Data for Name: section; Type: TABLE DATA; Schema: public; Owner: cardap
--

INSERT INTO public.section (id, created_at, updated_at, deleted_at, name) VALUES (1, NULL, NULL, NULL, 'Alimentação');
INSERT INTO public.section (id, created_at, updated_at, deleted_at, name) VALUES (3, NULL, NULL, NULL, 'Vestuário');
INSERT INTO public.section (id, created_at, updated_at, deleted_at, name) VALUES (4, NULL, NULL, NULL, 'Artesanato');
INSERT INTO public.section (id, created_at, updated_at, deleted_at, name) VALUES (6, NULL, NULL, NULL, 'Bebidas');
INSERT INTO public.section (id, created_at, updated_at, deleted_at, name) VALUES (7, NULL, NULL, NULL, 'Pizzas');
INSERT INTO public.section (id, created_at, updated_at, deleted_at, name) VALUES (8, NULL, NULL, NULL, 'Lanches Árabes');
INSERT INTO public.section (id, created_at, updated_at, deleted_at, name) VALUES (5, NULL, '2020-08-06 13:44:24.962456+00', NULL, 'Hamburguers');
INSERT INTO public.section (id, created_at, updated_at, deleted_at, name) VALUES (2, NULL, '2020-08-06 13:44:24.976931+00', NULL, 'Lanches');


--
-- TOC entry 3070 (class 0 OID 0)
-- Dependencies: 197
-- Name: additional_item_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.additional_item_id_seq', 5, true);


--
-- TOC entry 3071 (class 0 OID 0)
-- Dependencies: 223
-- Name: additional_items_group_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.additional_items_group_id_seq', 2, true);


--
-- TOC entry 3072 (class 0 OID 0)
-- Dependencies: 199
-- Name: address_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.address_id_seq', 1, false);


--
-- TOC entry 3073 (class 0 OID 0)
-- Dependencies: 221
-- Name: cardap_table_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.cardap_table_id_seq', 1, false);


--
-- TOC entry 3074 (class 0 OID 0)
-- Dependencies: 219
-- Name: cardap_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.cardap_user_id_seq', 2, true);


--
-- TOC entry 3075 (class 0 OID 0)
-- Dependencies: 201
-- Name: category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.category_id_seq', 2, true);


--
-- TOC entry 3076 (class 0 OID 0)
-- Dependencies: 205
-- Name: company_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.company_id_seq', 1, true);


--
-- TOC entry 3077 (class 0 OID 0)
-- Dependencies: 207
-- Name: menu_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.menu_id_seq', 1, true);


--
-- TOC entry 3078 (class 0 OID 0)
-- Dependencies: 209
-- Name: opening_hours_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.opening_hours_id_seq', 2, true);


--
-- TOC entry 3079 (class 0 OID 0)
-- Dependencies: 211
-- Name: payment_type_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.payment_type_id_seq', 13, true);


--
-- TOC entry 3080 (class 0 OID 0)
-- Dependencies: 214
-- Name: product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.product_id_seq', 4, true);


--
-- TOC entry 3081 (class 0 OID 0)
-- Dependencies: 216
-- Name: section_id_seq; Type: SEQUENCE SET; Schema: public; Owner: cardap
--

SELECT pg_catalog.setval('public.section_id_seq', 8, true);


--
-- TOC entry 2850 (class 2606 OID 37995)
-- Name: additional_item additional_item_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.additional_item
    ADD CONSTRAINT additional_item_pkey PRIMARY KEY (id);


--
-- TOC entry 2890 (class 2606 OID 38137)
-- Name: additional_items_group additional_items_group_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.additional_items_group
    ADD CONSTRAINT additional_items_group_pkey PRIMARY KEY (id);


--
-- TOC entry 2853 (class 2606 OID 38007)
-- Name: address address_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.address
    ADD CONSTRAINT address_pkey PRIMARY KEY (id);


--
-- TOC entry 2888 (class 2606 OID 38126)
-- Name: cardap_table cardap_table_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.cardap_table
    ADD CONSTRAINT cardap_table_pkey PRIMARY KEY (id);


--
-- TOC entry 2885 (class 2606 OID 38117)
-- Name: cardap_user cardap_user_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.cardap_user
    ADD CONSTRAINT cardap_user_pkey PRIMARY KEY (id);


--
-- TOC entry 2856 (class 2606 OID 38019)
-- Name: category category_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_pkey PRIMARY KEY (id);


--
-- TOC entry 2883 (class 2606 OID 38106)
-- Name: clients clients_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.clients
    ADD CONSTRAINT clients_pkey PRIMARY KEY (phone);


--
-- TOC entry 2861 (class 2606 OID 38030)
-- Name: company_payment_type company_payment_type_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.company_payment_type
    ADD CONSTRAINT company_payment_type_pkey PRIMARY KEY (company_id, payment_type_id);


--
-- TOC entry 2863 (class 2606 OID 38041)
-- Name: company company_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.company
    ADD CONSTRAINT company_pkey PRIMARY KEY (id);


--
-- TOC entry 2859 (class 2606 OID 38025)
-- Name: company_section company_section_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.company_section
    ADD CONSTRAINT company_section_pkey PRIMARY KEY (company_id, section_id);


--
-- TOC entry 2867 (class 2606 OID 38050)
-- Name: menu menu_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.menu
    ADD CONSTRAINT menu_pkey PRIMARY KEY (id);


--
-- TOC entry 2848 (class 2606 OID 37984)
-- Name: migrations migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.migrations
    ADD CONSTRAINT migrations_pkey PRIMARY KEY (id);


--
-- TOC entry 2870 (class 2606 OID 38062)
-- Name: opening_hours opening_hours_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.opening_hours
    ADD CONSTRAINT opening_hours_pkey PRIMARY KEY (id);


--
-- TOC entry 2873 (class 2606 OID 38071)
-- Name: payment_type payment_type_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.payment_type
    ADD CONSTRAINT payment_type_pkey PRIMARY KEY (id);


--
-- TOC entry 2875 (class 2606 OID 38077)
-- Name: product_additional_group product_additional_group_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.product_additional_group
    ADD CONSTRAINT product_additional_group_pkey PRIMARY KEY (product_id, additional_items_group_id);


--
-- TOC entry 2878 (class 2606 OID 38088)
-- Name: product product_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT product_pkey PRIMARY KEY (id);


--
-- TOC entry 2881 (class 2606 OID 38097)
-- Name: section section_pkey; Type: CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.section
    ADD CONSTRAINT section_pkey PRIMARY KEY (id);


--
-- TOC entry 2851 (class 1259 OID 37996)
-- Name: idx_additional_item_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_additional_item_deleted_at ON public.additional_item USING btree (deleted_at);


--
-- TOC entry 2891 (class 1259 OID 38138)
-- Name: idx_additional_items_group_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_additional_items_group_deleted_at ON public.additional_items_group USING btree (deleted_at);


--
-- TOC entry 2854 (class 1259 OID 38008)
-- Name: idx_address_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_address_deleted_at ON public.address USING btree (deleted_at);


--
-- TOC entry 2886 (class 1259 OID 38118)
-- Name: idx_cardap_user_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_cardap_user_deleted_at ON public.cardap_user USING btree (deleted_at);


--
-- TOC entry 2857 (class 1259 OID 38020)
-- Name: idx_category_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_category_deleted_at ON public.category USING btree (deleted_at);


--
-- TOC entry 2864 (class 1259 OID 38042)
-- Name: idx_company_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_company_deleted_at ON public.company USING btree (deleted_at);


--
-- TOC entry 2865 (class 1259 OID 38051)
-- Name: idx_menu_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_menu_deleted_at ON public.menu USING btree (deleted_at);


--
-- TOC entry 2868 (class 1259 OID 38063)
-- Name: idx_opening_hours_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_opening_hours_deleted_at ON public.opening_hours USING btree (deleted_at);


--
-- TOC entry 2871 (class 1259 OID 38072)
-- Name: idx_payment_type_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_payment_type_deleted_at ON public.payment_type USING btree (deleted_at);


--
-- TOC entry 2876 (class 1259 OID 38089)
-- Name: idx_product_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_product_deleted_at ON public.product USING btree (deleted_at);


--
-- TOC entry 2879 (class 1259 OID 38098)
-- Name: idx_section_deleted_at; Type: INDEX; Schema: public; Owner: cardap
--

CREATE INDEX idx_section_deleted_at ON public.section USING btree (deleted_at);


--
-- TOC entry 2892 (class 2606 OID 38169)
-- Name: additional_item additional_item_group_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.additional_item
    ADD CONSTRAINT additional_item_group_fk FOREIGN KEY (additional_items_group_id) REFERENCES public.additional_items_group(id) ON DELETE CASCADE;


--
-- TOC entry 2893 (class 2606 OID 38174)
-- Name: address address_company_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.address
    ADD CONSTRAINT address_company_fk FOREIGN KEY (company_id) REFERENCES public.company(id);


--
-- TOC entry 2894 (class 2606 OID 38149)
-- Name: category category_menu_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.category
    ADD CONSTRAINT category_menu_fk FOREIGN KEY (menu_id) REFERENCES public.menu(id) ON DELETE CASCADE;


--
-- TOC entry 2899 (class 2606 OID 38154)
-- Name: product category_product_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.product
    ADD CONSTRAINT category_product_fk FOREIGN KEY (category_id) REFERENCES public.category(id) ON DELETE CASCADE;


--
-- TOC entry 2897 (class 2606 OID 38179)
-- Name: company_payment_type company_payment_type_company_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.company_payment_type
    ADD CONSTRAINT company_payment_type_company_fk FOREIGN KEY (company_id) REFERENCES public.company(id);


--
-- TOC entry 2898 (class 2606 OID 38184)
-- Name: company_payment_type company_payment_type_payment_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.company_payment_type
    ADD CONSTRAINT company_payment_type_payment_fk FOREIGN KEY (payment_type_id) REFERENCES public.payment_type(id);


--
-- TOC entry 2896 (class 2606 OID 38144)
-- Name: company_section company_section_c_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.company_section
    ADD CONSTRAINT company_section_c_fk FOREIGN KEY (company_id) REFERENCES public.company(id);


--
-- TOC entry 2895 (class 2606 OID 38139)
-- Name: company_section company_section_s_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.company_section
    ADD CONSTRAINT company_section_s_fk FOREIGN KEY (section_id) REFERENCES public.section(id);


--
-- TOC entry 2901 (class 2606 OID 38164)
-- Name: additional_items_group group_company_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.additional_items_group
    ADD CONSTRAINT group_company_fk FOREIGN KEY (company_id) REFERENCES public.company(id);


--
-- TOC entry 2900 (class 2606 OID 38159)
-- Name: cardap_table table_company_fk; Type: FK CONSTRAINT; Schema: public; Owner: cardap
--

ALTER TABLE ONLY public.cardap_table
    ADD CONSTRAINT table_company_fk FOREIGN KEY (company_id) REFERENCES public.company(id);


-- Completed on 2020-08-06 10:48:16 -03

--
-- PostgreSQL database dump complete
--

